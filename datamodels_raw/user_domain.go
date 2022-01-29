// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package datamodels_raw

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UserDomain is an object representing the database table.
type UserDomain struct {
	UserID    string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	DomainID  string    `boil:"domain_id" json:"domain_id" toml:"domain_id" yaml:"domain_id"`
	Admin     bool      `boil:"admin" json:"admin" toml:"admin" yaml:"admin"`
	Email     bool      `boil:"email" json:"email" toml:"email" yaml:"email"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *userDomainR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userDomainL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserDomainColumns = struct {
	UserID    string
	DomainID  string
	Admin     string
	Email     string
	CreatedAt string
	UpdatedAt string
}{
	UserID:    "user_id",
	DomainID:  "domain_id",
	Admin:     "admin",
	Email:     "email",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var UserDomainTableColumns = struct {
	UserID    string
	DomainID  string
	Admin     string
	Email     string
	CreatedAt string
	UpdatedAt string
}{
	UserID:    "user_domain.user_id",
	DomainID:  "user_domain.domain_id",
	Admin:     "user_domain.admin",
	Email:     "user_domain.email",
	CreatedAt: "user_domain.created_at",
	UpdatedAt: "user_domain.updated_at",
}

// Generated where

var UserDomainWhere = struct {
	UserID    whereHelperstring
	DomainID  whereHelperstring
	Admin     whereHelperbool
	Email     whereHelperbool
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	UserID:    whereHelperstring{field: "\"user_domain\".\"user_id\""},
	DomainID:  whereHelperstring{field: "\"user_domain\".\"domain_id\""},
	Admin:     whereHelperbool{field: "\"user_domain\".\"admin\""},
	Email:     whereHelperbool{field: "\"user_domain\".\"email\""},
	CreatedAt: whereHelpernull_Time{field: "\"user_domain\".\"created_at\""},
	UpdatedAt: whereHelpernull_Time{field: "\"user_domain\".\"updated_at\""},
}

// UserDomainRels is where relationship names are stored.
var UserDomainRels = struct {
	Domain string
	User   string
}{
	Domain: "Domain",
	User:   "User",
}

// userDomainR is where relationships are stored.
type userDomainR struct {
	Domain *Domain `boil:"Domain" json:"Domain" toml:"Domain" yaml:"Domain"`
	User   *User   `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*userDomainR) NewStruct() *userDomainR {
	return &userDomainR{}
}

// userDomainL is where Load methods for each relationship are stored.
type userDomainL struct{}

var (
	userDomainAllColumns            = []string{"user_id", "domain_id", "admin", "email", "created_at", "updated_at"}
	userDomainColumnsWithoutDefault = []string{"user_id", "domain_id"}
	userDomainColumnsWithDefault    = []string{"admin", "email", "created_at", "updated_at"}
	userDomainPrimaryKeyColumns     = []string{"user_id", "domain_id"}
	userDomainGeneratedColumns      = []string{}
)

type (
	// UserDomainSlice is an alias for a slice of pointers to UserDomain.
	// This should almost always be used instead of []UserDomain.
	UserDomainSlice []*UserDomain
	// UserDomainHook is the signature for custom UserDomain hook methods
	UserDomainHook func(context.Context, boil.ContextExecutor, *UserDomain) error

	userDomainQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userDomainType                 = reflect.TypeOf(&UserDomain{})
	userDomainMapping              = queries.MakeStructMapping(userDomainType)
	userDomainPrimaryKeyMapping, _ = queries.BindMapping(userDomainType, userDomainMapping, userDomainPrimaryKeyColumns)
	userDomainInsertCacheMut       sync.RWMutex
	userDomainInsertCache          = make(map[string]insertCache)
	userDomainUpdateCacheMut       sync.RWMutex
	userDomainUpdateCache          = make(map[string]updateCache)
	userDomainUpsertCacheMut       sync.RWMutex
	userDomainUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userDomainAfterSelectHooks []UserDomainHook

var userDomainBeforeInsertHooks []UserDomainHook
var userDomainAfterInsertHooks []UserDomainHook

var userDomainBeforeUpdateHooks []UserDomainHook
var userDomainAfterUpdateHooks []UserDomainHook

var userDomainBeforeDeleteHooks []UserDomainHook
var userDomainAfterDeleteHooks []UserDomainHook

var userDomainBeforeUpsertHooks []UserDomainHook
var userDomainAfterUpsertHooks []UserDomainHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserDomain) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userDomainAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserDomain) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userDomainBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserDomain) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userDomainAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserDomain) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userDomainBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserDomain) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userDomainAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserDomain) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userDomainBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserDomain) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userDomainAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserDomain) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userDomainBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserDomain) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userDomainAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserDomainHook registers your hook function for all future operations.
func AddUserDomainHook(hookPoint boil.HookPoint, userDomainHook UserDomainHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userDomainAfterSelectHooks = append(userDomainAfterSelectHooks, userDomainHook)
	case boil.BeforeInsertHook:
		userDomainBeforeInsertHooks = append(userDomainBeforeInsertHooks, userDomainHook)
	case boil.AfterInsertHook:
		userDomainAfterInsertHooks = append(userDomainAfterInsertHooks, userDomainHook)
	case boil.BeforeUpdateHook:
		userDomainBeforeUpdateHooks = append(userDomainBeforeUpdateHooks, userDomainHook)
	case boil.AfterUpdateHook:
		userDomainAfterUpdateHooks = append(userDomainAfterUpdateHooks, userDomainHook)
	case boil.BeforeDeleteHook:
		userDomainBeforeDeleteHooks = append(userDomainBeforeDeleteHooks, userDomainHook)
	case boil.AfterDeleteHook:
		userDomainAfterDeleteHooks = append(userDomainAfterDeleteHooks, userDomainHook)
	case boil.BeforeUpsertHook:
		userDomainBeforeUpsertHooks = append(userDomainBeforeUpsertHooks, userDomainHook)
	case boil.AfterUpsertHook:
		userDomainAfterUpsertHooks = append(userDomainAfterUpsertHooks, userDomainHook)
	}
}

// One returns a single userDomain record from the query.
func (q userDomainQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserDomain, error) {
	o := &UserDomain{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "datamodels_raw: failed to execute a one query for user_domain")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserDomain records from the query.
func (q userDomainQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserDomainSlice, error) {
	var o []*UserDomain

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "datamodels_raw: failed to assign all query results to UserDomain slice")
	}

	if len(userDomainAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserDomain records in the query.
func (q userDomainQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to count user_domain rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userDomainQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "datamodels_raw: failed to check if user_domain exists")
	}

	return count > 0, nil
}

// Domain pointed to by the foreign key.
func (o *UserDomain) Domain(mods ...qm.QueryMod) domainQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.DomainID),
	}

	queryMods = append(queryMods, mods...)

	query := Domains(queryMods...)
	queries.SetFrom(query.Query, "\"domains\"")

	return query
}

// User pointed to by the foreign key.
func (o *UserDomain) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadDomain allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userDomainL) LoadDomain(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserDomain interface{}, mods queries.Applicator) error {
	var slice []*UserDomain
	var object *UserDomain

	if singular {
		object = maybeUserDomain.(*UserDomain)
	} else {
		slice = *maybeUserDomain.(*[]*UserDomain)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userDomainR{}
		}
		args = append(args, object.DomainID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userDomainR{}
			}

			for _, a := range args {
				if a == obj.DomainID {
					continue Outer
				}
			}

			args = append(args, obj.DomainID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`domains`),
		qm.WhereIn(`domains.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Domain")
	}

	var resultSlice []*Domain
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Domain")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for domains")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for domains")
	}

	if len(userDomainAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Domain = foreign
		if foreign.R == nil {
			foreign.R = &domainR{}
		}
		foreign.R.UserDomains = append(foreign.R.UserDomains, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.DomainID == foreign.ID {
				local.R.Domain = foreign
				if foreign.R == nil {
					foreign.R = &domainR{}
				}
				foreign.R.UserDomains = append(foreign.R.UserDomains, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userDomainL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserDomain interface{}, mods queries.Applicator) error {
	var slice []*UserDomain
	var object *UserDomain

	if singular {
		object = maybeUserDomain.(*UserDomain)
	} else {
		slice = *maybeUserDomain.(*[]*UserDomain)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userDomainR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userDomainR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userDomainAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserDomains = append(foreign.R.UserDomains, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserDomains = append(foreign.R.UserDomains, local)
				break
			}
		}
	}

	return nil
}

// SetDomain of the userDomain to the related item.
// Sets o.R.Domain to related.
// Adds o to related.R.UserDomains.
func (o *UserDomain) SetDomain(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Domain) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_domain\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"domain_id"}),
		strmangle.WhereClause("\"", "\"", 2, userDomainPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.DomainID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DomainID = related.ID
	if o.R == nil {
		o.R = &userDomainR{
			Domain: related,
		}
	} else {
		o.R.Domain = related
	}

	if related.R == nil {
		related.R = &domainR{
			UserDomains: UserDomainSlice{o},
		}
	} else {
		related.R.UserDomains = append(related.R.UserDomains, o)
	}

	return nil
}

// SetUser of the userDomain to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserDomains.
func (o *UserDomain) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_domain\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userDomainPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.DomainID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &userDomainR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserDomains: UserDomainSlice{o},
		}
	} else {
		related.R.UserDomains = append(related.R.UserDomains, o)
	}

	return nil
}

// UserDomains retrieves all the records using an executor.
func UserDomains(mods ...qm.QueryMod) userDomainQuery {
	mods = append(mods, qm.From("\"user_domain\""))
	return userDomainQuery{NewQuery(mods...)}
}

// FindUserDomain retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserDomain(ctx context.Context, exec boil.ContextExecutor, userID string, domainID string, selectCols ...string) (*UserDomain, error) {
	userDomainObj := &UserDomain{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_domain\" where \"user_id\"=$1 AND \"domain_id\"=$2", sel,
	)

	q := queries.Raw(query, userID, domainID)

	err := q.Bind(ctx, exec, userDomainObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "datamodels_raw: unable to select from user_domain")
	}

	if err = userDomainObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userDomainObj, err
	}

	return userDomainObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserDomain) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("datamodels_raw: no user_domain provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userDomainColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userDomainInsertCacheMut.RLock()
	cache, cached := userDomainInsertCache[key]
	userDomainInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userDomainAllColumns,
			userDomainColumnsWithDefault,
			userDomainColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userDomainType, userDomainMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userDomainType, userDomainMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_domain\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_domain\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "datamodels_raw: unable to insert into user_domain")
	}

	if !cached {
		userDomainInsertCacheMut.Lock()
		userDomainInsertCache[key] = cache
		userDomainInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserDomain.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserDomain) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userDomainUpdateCacheMut.RLock()
	cache, cached := userDomainUpdateCache[key]
	userDomainUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userDomainAllColumns,
			userDomainPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("datamodels_raw: unable to update user_domain, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_domain\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userDomainPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userDomainType, userDomainMapping, append(wl, userDomainPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to update user_domain row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to get rows affected by update for user_domain")
	}

	if !cached {
		userDomainUpdateCacheMut.Lock()
		userDomainUpdateCache[key] = cache
		userDomainUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userDomainQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to update all for user_domain")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to retrieve rows affected for user_domain")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserDomainSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("datamodels_raw: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_domain\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userDomainPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to update all in userDomain slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to retrieve rows affected all in update all userDomain")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserDomain) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("datamodels_raw: no user_domain provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userDomainColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	userDomainUpsertCacheMut.RLock()
	cache, cached := userDomainUpsertCache[key]
	userDomainUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userDomainAllColumns,
			userDomainColumnsWithDefault,
			userDomainColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userDomainAllColumns,
			userDomainPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("datamodels_raw: unable to upsert user_domain, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userDomainPrimaryKeyColumns))
			copy(conflict, userDomainPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_domain\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userDomainType, userDomainMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userDomainType, userDomainMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "datamodels_raw: unable to upsert user_domain")
	}

	if !cached {
		userDomainUpsertCacheMut.Lock()
		userDomainUpsertCache[key] = cache
		userDomainUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserDomain record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserDomain) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("datamodels_raw: no UserDomain provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userDomainPrimaryKeyMapping)
	sql := "DELETE FROM \"user_domain\" WHERE \"user_id\"=$1 AND \"domain_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to delete from user_domain")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to get rows affected by delete for user_domain")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userDomainQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("datamodels_raw: no userDomainQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to delete all from user_domain")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to get rows affected by deleteall for user_domain")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserDomainSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userDomainBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_domain\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userDomainPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to delete all from userDomain slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to get rows affected by deleteall for user_domain")
	}

	if len(userDomainAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserDomain) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserDomain(ctx, exec, o.UserID, o.DomainID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserDomainSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserDomainSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_domain\".* FROM \"user_domain\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userDomainPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "datamodels_raw: unable to reload all in UserDomainSlice")
	}

	*o = slice

	return nil
}

// UserDomainExists checks if the UserDomain row exists.
func UserDomainExists(ctx context.Context, exec boil.ContextExecutor, userID string, domainID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_domain\" where \"user_id\"=$1 AND \"domain_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, userID, domainID)
	}
	row := exec.QueryRowContext(ctx, sql, userID, domainID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "datamodels_raw: unable to check if user_domain exists")
	}

	return exists, nil
}
