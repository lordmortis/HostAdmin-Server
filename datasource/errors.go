package datasource

import (
	"errors"
)

var ErrUUIDParse = errors.New("unable to parse UUID")
var ErrNoResults = errors.New("no results from query")
var ErrNoDb = errors.New("no database connection")
var ErrNoUpdate = errors.New("model not updated")
var ErrNotInDb = errors.New("model not in database")