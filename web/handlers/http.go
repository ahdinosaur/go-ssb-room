// SPDX-License-Identifier: MIT

package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"go.mindeco.de/http/auth"
	"go.mindeco.de/http/render"

	"github.com/ssb-ngi-pointer/go-ssb-room/admindb"
	"github.com/ssb-ngi-pointer/go-ssb-room/internal/repo"
	"github.com/ssb-ngi-pointer/go-ssb-room/roomstate"
	"github.com/ssb-ngi-pointer/go-ssb-room/web"
	"github.com/ssb-ngi-pointer/go-ssb-room/web/handlers/admin"
	roomsAuth "github.com/ssb-ngi-pointer/go-ssb-room/web/handlers/auth"
	"github.com/ssb-ngi-pointer/go-ssb-room/web/handlers/news"
	"github.com/ssb-ngi-pointer/go-ssb-room/web/i18n"
	"github.com/ssb-ngi-pointer/go-ssb-room/web/router"
)

// New initializes the whole web stack for rooms, with all the sub-modules and routing.
func New(
	m *mux.Router,
	repo repo.Interface,
	roomState *roomstate.Manager,
	as admindb.AuthWithSSBService,
	fs admindb.AuthFallbackService,
) (http.Handler, error) {
	if m == nil {
		m = router.CompleteApp()
	}

	locHelper, err := i18n.New(repo)
	if err != nil {
		return nil, err
	}

	var a *auth.Handler

	r, err := render.New(web.Templates,
		render.BaseTemplates("/base.tmpl"),
		render.AddTemplates(concatTemplates(
			[]string{
				"/landing/index.tmpl",
				"/landing/about.tmpl",
				"/error.tmpl",
			},
			news.HTMLTemplates,
			roomsAuth.HTMLTemplates,
			admin.HTMLTemplates,
		)...),
		render.FuncMap(web.TemplateFuncs(m)),
		// TODO: move these to the i18n helper pkg
		render.InjectTemplateFunc("i18npl", func(r *http.Request) interface{} {
			loc := localizerFromRequest(locHelper, r)
			return loc.LocalizePlurals
		}),
		render.InjectTemplateFunc("i18n", func(r *http.Request) interface{} {
			loc := localizerFromRequest(locHelper, r)
			return loc.LocalizeSimple
		}),
		render.InjectTemplateFunc("is_logged_in", func(r *http.Request) interface{} {
			no := func() *admindb.User { return nil }

			v, err := a.AuthenticateRequest(r)
			if err != nil {
				return no
			}
			uid, ok := v.(int64)
			if !ok {
				// TODO: hook up logging
				fmt.Fprintf(os.Stderr, "warning: not the expected ID type: %T\n", v)
				return no
			}

			user, err := fs.GetByID(r.Context(), uid)
			if err != nil {
				return no
			}

			yes := func() *admindb.User { return user }
			return yes
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("web Handler: failed to create renderer: %w", err)
	}

	cookieCodec, err := web.LoadOrCreateCookieSecrets(repo)
	if err != nil {
		return nil, err
	}

	store := &sessions.CookieStore{
		Codecs: cookieCodec,
		Options: &sessions.Options{
			Path:   "/",
			MaxAge: 2 * 60 * 60, // two hours in seconds
		},
	}

	authErrH := func(rw http.ResponseWriter, req *http.Request, err error, code int) {
		msg := err.Error()

		// localize some specific error messages
		if err == auth.ErrBadLogin {
			msg = localizerFromRequest(locHelper, req).LocalizeSimple("AuthErrorBadLogin")
		}

		r.HTML("/error.tmpl", func(rw http.ResponseWriter, req *http.Request) (interface{}, error) {
			return errorTemplateData{
				Err: msg,
				// TODO: localize?
				Status:     http.StatusText(code),
				StatusCode: code,
			}, nil
		}).ServeHTTP(rw, req)
	}

	notAuthorizedH := r.HTML("/error.tmpl", func(rw http.ResponseWriter, req *http.Request) (interface{}, error) {
		statusCode := http.StatusUnauthorized
		rw.WriteHeader(statusCode)
		return errorTemplateData{
			statusCode,
			"Unauthorized",
			"you are not authorized to access the requested site",
		}, nil
	})

	a, err = auth.NewHandler(fs,
		auth.SetStore(store),
		auth.SetErrorHandler(authErrH),
		auth.SetNotAuthorizedHandler(notAuthorizedH),
	)
	if err != nil {
		return nil, fmt.Errorf("web Handler: failed to init fallback auth system: %w", err)
	}

	// hookup handlers to the router
	roomsAuth.Handler(m, r, a)

	adminRouter := m.PathPrefix("/admin").Subrouter()
	adminRouter.Use(a.Authenticate)

	// we dont strip path here because it somehow fucks with the middleware setup
	adminRouter.PathPrefix("/").Handler(admin.Handler(adminRouter, r, roomState))

	m.PathPrefix("/news").Handler(http.StripPrefix("/news", news.Handler(m, r)))

	m.Get(router.CompleteIndex).Handler(r.StaticHTML("/landing/index.tmpl"))
	m.Get(router.CompleteAbout).Handler(r.StaticHTML("/landing/about.tmpl"))

	m.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(web.Assets)))

	m.NotFoundHandler = r.HTML("/error.tmpl", func(rw http.ResponseWriter, req *http.Request) (interface{}, error) {
		rw.WriteHeader(http.StatusNotFound)
		return errorTemplateData{http.StatusNotFound, "Not Found", "the requested page wasnt found.."}, nil
	})

	if web.Production {
		return m, nil
	}

	return r.GetReloader()(m), nil
}

// utils

type errorTemplateData struct {
	StatusCode int
	Status     string
	Err        string
}

func concatTemplates(lst ...[]string) []string {
	var catted []string

	for _, tpls := range lst {
		for _, t := range tpls {
			catted = append(catted, t)
		}

	}
	return catted
}

func localizerFromRequest(helper *i18n.Helper, r *http.Request) *i18n.Localizer {
	lang := r.FormValue("lang")
	accept := r.Header.Get("Accept-Language")
	return helper.NewLocalizer(lang, accept)
}
