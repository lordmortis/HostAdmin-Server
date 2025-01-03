// Code generated by SQLBoiler 4.12.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package datamodels_raw

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testUserDomains(t *testing.T) {
	t.Parallel()

	query := UserDomains()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserDomainsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserDomain{}
	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserDomainsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserDomain{}
	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserDomains().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserDomainsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserDomain{}
	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserDomainSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserDomainsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserDomain{}
	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserDomainExists(ctx, tx, o.UserID, o.DomainID)
	if err != nil {
		t.Errorf("Unable to check if UserDomain exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserDomainExists to return true, but got false.")
	}
}

func testUserDomainsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserDomain{}
	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userDomainFound, err := FindUserDomain(ctx, tx, o.UserID, o.DomainID)
	if err != nil {
		t.Error(err)
	}

	if userDomainFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserDomainsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserDomain{}
	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserDomains().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserDomainsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserDomain{}
	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserDomains().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserDomainsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userDomainOne := &UserDomain{}
	userDomainTwo := &UserDomain{}
	if err = randomize.Struct(seed, userDomainOne, userDomainDBTypes, false, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}
	if err = randomize.Struct(seed, userDomainTwo, userDomainDBTypes, false, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userDomainOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userDomainTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserDomains().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserDomainsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userDomainOne := &UserDomain{}
	userDomainTwo := &UserDomain{}
	if err = randomize.Struct(seed, userDomainOne, userDomainDBTypes, false, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}
	if err = randomize.Struct(seed, userDomainTwo, userDomainDBTypes, false, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userDomainOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userDomainTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userDomainBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserDomain) error {
	*o = UserDomain{}
	return nil
}

func userDomainAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserDomain) error {
	*o = UserDomain{}
	return nil
}

func userDomainAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserDomain) error {
	*o = UserDomain{}
	return nil
}

func userDomainBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserDomain) error {
	*o = UserDomain{}
	return nil
}

func userDomainAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserDomain) error {
	*o = UserDomain{}
	return nil
}

func userDomainBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserDomain) error {
	*o = UserDomain{}
	return nil
}

func userDomainAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserDomain) error {
	*o = UserDomain{}
	return nil
}

func userDomainBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserDomain) error {
	*o = UserDomain{}
	return nil
}

func userDomainAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserDomain) error {
	*o = UserDomain{}
	return nil
}

func testUserDomainsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserDomain{}
	o := &UserDomain{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userDomainDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserDomain object: %s", err)
	}

	AddUserDomainHook(boil.BeforeInsertHook, userDomainBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userDomainBeforeInsertHooks = []UserDomainHook{}

	AddUserDomainHook(boil.AfterInsertHook, userDomainAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userDomainAfterInsertHooks = []UserDomainHook{}

	AddUserDomainHook(boil.AfterSelectHook, userDomainAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userDomainAfterSelectHooks = []UserDomainHook{}

	AddUserDomainHook(boil.BeforeUpdateHook, userDomainBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userDomainBeforeUpdateHooks = []UserDomainHook{}

	AddUserDomainHook(boil.AfterUpdateHook, userDomainAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userDomainAfterUpdateHooks = []UserDomainHook{}

	AddUserDomainHook(boil.BeforeDeleteHook, userDomainBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userDomainBeforeDeleteHooks = []UserDomainHook{}

	AddUserDomainHook(boil.AfterDeleteHook, userDomainAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userDomainAfterDeleteHooks = []UserDomainHook{}

	AddUserDomainHook(boil.BeforeUpsertHook, userDomainBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userDomainBeforeUpsertHooks = []UserDomainHook{}

	AddUserDomainHook(boil.AfterUpsertHook, userDomainAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userDomainAfterUpsertHooks = []UserDomainHook{}
}

func testUserDomainsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserDomain{}
	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserDomainsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserDomain{}
	if err = randomize.Struct(seed, o, userDomainDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userDomainColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserDomainToOneDomainUsingDomain(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserDomain
	var foreign Domain

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userDomainDBTypes, false, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, domainDBTypes, false, domainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Domain struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.DomainID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Domain().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UserDomainSlice{&local}
	if err = local.L.LoadDomain(ctx, tx, false, (*[]*UserDomain)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Domain == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Domain = nil
	if err = local.L.LoadDomain(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Domain == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUserDomainToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserDomain
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userDomainDBTypes, false, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UserDomainSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*UserDomain)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUserDomainToOneSetOpDomainUsingDomain(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserDomain
	var b, c Domain

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userDomainDBTypes, false, strmangle.SetComplement(userDomainPrimaryKeyColumns, userDomainColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, domainDBTypes, false, strmangle.SetComplement(domainPrimaryKeyColumns, domainColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, domainDBTypes, false, strmangle.SetComplement(domainPrimaryKeyColumns, domainColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Domain{&b, &c} {
		err = a.SetDomain(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Domain != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserDomains[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DomainID != x.ID {
			t.Error("foreign key was wrong value", a.DomainID)
		}

		if exists, err := UserDomainExists(ctx, tx, a.UserID, a.DomainID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testUserDomainToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserDomain
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userDomainDBTypes, false, strmangle.SetComplement(userDomainPrimaryKeyColumns, userDomainColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserDomains[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		if exists, err := UserDomainExists(ctx, tx, a.UserID, a.DomainID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testUserDomainsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserDomain{}
	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserDomainsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserDomain{}
	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserDomainSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserDomainsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserDomain{}
	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserDomains().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userDomainDBTypes = map[string]string{`UserID`: `uuid`, `DomainID`: `uuid`, `Admin`: `boolean`, `Email`: `boolean`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                 = bytes.MinRead
)

func testUserDomainsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userDomainPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userDomainAllColumns) == len(userDomainPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserDomain{}
	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserDomainsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userDomainAllColumns) == len(userDomainPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserDomain{}
	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userDomainDBTypes, true, userDomainPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userDomainAllColumns, userDomainPrimaryKeyColumns) {
		fields = userDomainAllColumns
	} else {
		fields = strmangle.SetComplement(
			userDomainAllColumns,
			userDomainPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := UserDomainSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserDomainsUpsert(t *testing.T) {
	t.Parallel()

	if len(userDomainAllColumns) == len(userDomainPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserDomain{}
	if err = randomize.Struct(seed, &o, userDomainDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserDomain: %s", err)
	}

	count, err := UserDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userDomainDBTypes, false, userDomainPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserDomain struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserDomain: %s", err)
	}

	count, err = UserDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
