// SPDX-License-Identifier: MIT

// Package roomdb implements all the persisted database needs of the room server.
// This includes authentication, allow/deny list managment, invite and alias creation and also the notice content for the CMS.
//
// The interfaces defined here are implemented twice. Once in SQLite for production and once as mocks for testing, generated by counterfeiter (https://github.com/maxbrunsfeld/counterfeiter).
//
// See the package documentation of roomdb/sqlite for how to update it.
// It's important not to use the types generated by sqlboiler (sqlite/models) in the argument and return values of the interfaces here.
// This would leak details of the internal implementation of the roomdb/sqlite package and we want to have full control over how these interfaces can be used.
package roomdb

import (
	"context"

	"go.mindeco.de/http/auth"
	refs "go.mindeco.de/ssb-refs"
)

// AuthFallbackService allows password authentication which might be helpful for scenarios
// where one lost access to his ssb device or key.
type AuthFallbackService interface {

	// Check receives the username and password (in clear) and checks them accordingly.
	// If it's a valid combination it returns the user ID, or an error if they are not.
	auth.Auther

	Create(_ context.Context, memberID int64, login string, password []byte) error
	// GetByID(context.Context, int64) (User, error)
	// ListAll()
	// ListByMember()
	// Remove(pwid)
}

// needed?! not sure we need to hold the challanges
// AuthWithSSBService defines functions needed for the challange/response system of sign-in with ssb
type AuthWithSSBService interface{}

// MembersService stores and retreives the list of internal users (members, mods and admins).
type MembersService interface {
	// Add adds a new member
	Add(_ context.Context, nickName string, pubKey refs.FeedRef, r Role) (int64, error)

	// GetByID returns the member if it exists
	GetByID(context.Context, int64) (Member, error)

	// GetByFeed returns the member if it exists
	GetByFeed(context.Context, refs.FeedRef) (Member, error)

	// List returns a list of all the members.
	List(context.Context) ([]Member, error)

	// RemoveFeed removes the feed from the list.
	RemoveFeed(context.Context, refs.FeedRef) error

	// RemoveID removes the feed for the ID from the list.
	RemoveID(context.Context, int64) error

	// SetRole
	SetRole(context.Context, int64, Role) error
}

// DeniedKeysService changes the lists of public keys that are not allowed to get into the room
type DeniedKeysService interface {
	// Add adds the feed to the list, together with a comment for other members
	Add(ctx context.Context, ref refs.FeedRef, comment string) error

	// HasFeed returns true if a feed is on the list.
	HasFeed(context.Context, refs.FeedRef) bool

	// HasFeed returns true if a feed is on the list.
	HasID(context.Context, int64) bool

	// GetByID returns the list entry for that ID or an error
	GetByID(context.Context, int64) (ListEntry, error)

	// List returns a list of all the feeds.
	List(context.Context) ([]ListEntry, error)

	// RemoveFeed removes the feed from the list.
	RemoveFeed(context.Context, refs.FeedRef) error

	// RemoveID removes the feed for the ID from the list.
	RemoveID(context.Context, int64) error
}

// AliasesService manages alias handle registration and lookup
type AliasesService interface {
	// Resolve returns all the relevant information for that alias or an error if it doesnt exist
	Resolve(context.Context, string) (Alias, error)

	// GetByID returns the alias for that ID or an error
	GetByID(context.Context, int64) (Alias, error)

	// List returns a list of all registerd aliases
	List(ctx context.Context) ([]Alias, error)

	// Register receives an alias and signature for it. Validation needs to happen before this.
	Register(ctx context.Context, alias string, userFeed refs.FeedRef, signature []byte) error

	// Revoke removes an alias from the system
	Revoke(ctx context.Context, alias string) error
}

// InvitesService manages creation and consumption of invite tokens for joining the room.
type InvitesService interface {
	// Create creates a new invite for a new member. It returns the token or an error.
	// createdBy is user ID of the admin or moderator who created it.
	// aliasSuggestion is optional (empty string is fine) but can be used to disambiguate open invites. (See https://github.com/ssb-ngi-pointer/rooms2/issues/21)
	Create(ctx context.Context, createdBy int64, aliasSuggestion string) (string, error)

	// Consume checks if the passed token is still valid.
	// If it is it adds newMember to the members of the room and invalidates the token.
	// If the token isn't valid, it returns an error.
	Consume(ctx context.Context, token string, newMember refs.FeedRef) (Invite, error)

	// GetByToken returns the Invite if one for that token exists, or an error
	GetByToken(ctx context.Context, token string) (Invite, error)

	// GetByToken returns the Invite if one for that ID exists, or an error
	GetByID(ctx context.Context, id int64) (Invite, error)

	// List returns a list of all the valid invites
	List(ctx context.Context) ([]Invite, error)

	// Revoke removes a active invite and invalidates it for future use.
	Revoke(ctx context.Context, id int64) error
}

// PinnedNoticesService allows an admin to assign Notices to specific placeholder pages.
// like updates, privacy policy, code of conduct
type PinnedNoticesService interface {
	// List returns a list of all the pinned notices with their corresponding notices and languages
	List(context.Context) (PinnedNotices, error)

	// Set assigns a fixed page name to an page ID and a language to allow for multiple translated versions of the same page.
	Set(ctx context.Context, name PinnedNoticeName, id int64) error

	// Get returns a single notice for a name and a language
	Get(ctx context.Context, name PinnedNoticeName, language string) (*Notice, error)
}

// NoticesService is the low level store to manage single notices
type NoticesService interface {
	// GetByID returns the page for that ID or an error
	GetByID(context.Context, int64) (Notice, error)

	// Save updates the passed page or creates it if it's ID is zero
	Save(context.Context, *Notice) error

	// RemoveID removes the page for that ID.
	RemoveID(context.Context, int64) error
}

// for tests we use generated mocks from these interfaces created with https://github.com/maxbrunsfeld/counterfeiter

//go:generate counterfeiter -o mockdb/aliases.go . AliasesService

//go:generate counterfeiter -o mockdb/auth.go . AuthWithSSBService

//go:generate counterfeiter -o mockdb/auth_fallback.go . AuthFallbackService

//go:generate counterfeiter -o mockdb/denied.go . DeniedKeysService

//go:generate counterfeiter -o mockdb/fixed_pages.go . PinnedNoticesService

//go:generate counterfeiter -o mockdb/invites.go . InvitesService

//go:generate counterfeiter -o mockdb/members.go . MembersService

//go:generate counterfeiter -o mockdb/pages.go . NoticesService
