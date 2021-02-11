// SPDX-License-Identifier: MIT

package admindb

import (
	"database/sql/driver"
	"fmt"

	refs "go.mindeco.de/ssb-refs"
)

// It's important to wrap all the model generated types into these since we don't want the admindb interfaces to depend on them.

// User holds all the information an authenticated user of the site has.
type User struct {
	ID   int64
	Name string
}

// DBFeedRef wraps a feed reference and implements the SQL marshaling interfaces.
type DBFeedRef struct{ refs.FeedRef }

// Scan implements https://pkg.go.dev/database/sql#Scanner to read strings into feed references.
func (r *DBFeedRef) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		return fmt.Errorf("unexpected type: %T", src)
	}

	fr, err := refs.ParseFeedRef(str)
	if err != nil {
		return err
	}

	r.FeedRef = *fr
	return nil
}

// Value returns feed references as strings to the database.
// https://pkg.go.dev/database/sql/driver#Valuer
func (r DBFeedRef) Value() (driver.Value, error) {
	return driver.Value(r.Ref()), nil
}
