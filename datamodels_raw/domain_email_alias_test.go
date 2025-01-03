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

func testDomainEmailAliases(t *testing.T) {
	t.Parallel()

	query := DomainEmailAliases()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDomainEmailAliasesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainEmailAlias{}
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
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

	count, err := DomainEmailAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDomainEmailAliasesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainEmailAlias{}
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DomainEmailAliases().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DomainEmailAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDomainEmailAliasesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainEmailAlias{}
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DomainEmailAliasSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DomainEmailAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDomainEmailAliasesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainEmailAlias{}
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DomainEmailAliasExists(ctx, tx, o.DomainID, o.Address)
	if err != nil {
		t.Errorf("Unable to check if DomainEmailAlias exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DomainEmailAliasExists to return true, but got false.")
	}
}

func testDomainEmailAliasesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainEmailAlias{}
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	domainEmailAliasFound, err := FindDomainEmailAlias(ctx, tx, o.DomainID, o.Address)
	if err != nil {
		t.Error(err)
	}

	if domainEmailAliasFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDomainEmailAliasesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainEmailAlias{}
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DomainEmailAliases().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDomainEmailAliasesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainEmailAlias{}
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DomainEmailAliases().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDomainEmailAliasesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	domainEmailAliasOne := &DomainEmailAlias{}
	domainEmailAliasTwo := &DomainEmailAlias{}
	if err = randomize.Struct(seed, domainEmailAliasOne, domainEmailAliasDBTypes, false, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}
	if err = randomize.Struct(seed, domainEmailAliasTwo, domainEmailAliasDBTypes, false, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = domainEmailAliasOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = domainEmailAliasTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DomainEmailAliases().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDomainEmailAliasesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	domainEmailAliasOne := &DomainEmailAlias{}
	domainEmailAliasTwo := &DomainEmailAlias{}
	if err = randomize.Struct(seed, domainEmailAliasOne, domainEmailAliasDBTypes, false, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}
	if err = randomize.Struct(seed, domainEmailAliasTwo, domainEmailAliasDBTypes, false, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = domainEmailAliasOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = domainEmailAliasTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DomainEmailAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func domainEmailAliasBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DomainEmailAlias) error {
	*o = DomainEmailAlias{}
	return nil
}

func domainEmailAliasAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DomainEmailAlias) error {
	*o = DomainEmailAlias{}
	return nil
}

func domainEmailAliasAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DomainEmailAlias) error {
	*o = DomainEmailAlias{}
	return nil
}

func domainEmailAliasBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DomainEmailAlias) error {
	*o = DomainEmailAlias{}
	return nil
}

func domainEmailAliasAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DomainEmailAlias) error {
	*o = DomainEmailAlias{}
	return nil
}

func domainEmailAliasBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DomainEmailAlias) error {
	*o = DomainEmailAlias{}
	return nil
}

func domainEmailAliasAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DomainEmailAlias) error {
	*o = DomainEmailAlias{}
	return nil
}

func domainEmailAliasBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DomainEmailAlias) error {
	*o = DomainEmailAlias{}
	return nil
}

func domainEmailAliasAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DomainEmailAlias) error {
	*o = DomainEmailAlias{}
	return nil
}

func testDomainEmailAliasesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DomainEmailAlias{}
	o := &DomainEmailAlias{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias object: %s", err)
	}

	AddDomainEmailAliasHook(boil.BeforeInsertHook, domainEmailAliasBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	domainEmailAliasBeforeInsertHooks = []DomainEmailAliasHook{}

	AddDomainEmailAliasHook(boil.AfterInsertHook, domainEmailAliasAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	domainEmailAliasAfterInsertHooks = []DomainEmailAliasHook{}

	AddDomainEmailAliasHook(boil.AfterSelectHook, domainEmailAliasAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	domainEmailAliasAfterSelectHooks = []DomainEmailAliasHook{}

	AddDomainEmailAliasHook(boil.BeforeUpdateHook, domainEmailAliasBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	domainEmailAliasBeforeUpdateHooks = []DomainEmailAliasHook{}

	AddDomainEmailAliasHook(boil.AfterUpdateHook, domainEmailAliasAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	domainEmailAliasAfterUpdateHooks = []DomainEmailAliasHook{}

	AddDomainEmailAliasHook(boil.BeforeDeleteHook, domainEmailAliasBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	domainEmailAliasBeforeDeleteHooks = []DomainEmailAliasHook{}

	AddDomainEmailAliasHook(boil.AfterDeleteHook, domainEmailAliasAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	domainEmailAliasAfterDeleteHooks = []DomainEmailAliasHook{}

	AddDomainEmailAliasHook(boil.BeforeUpsertHook, domainEmailAliasBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	domainEmailAliasBeforeUpsertHooks = []DomainEmailAliasHook{}

	AddDomainEmailAliasHook(boil.AfterUpsertHook, domainEmailAliasAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	domainEmailAliasAfterUpsertHooks = []DomainEmailAliasHook{}
}

func testDomainEmailAliasesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainEmailAlias{}
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DomainEmailAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDomainEmailAliasesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainEmailAlias{}
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(domainEmailAliasColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DomainEmailAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDomainEmailAliasToOneDomainUsingDomain(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DomainEmailAlias
	var foreign Domain

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, domainEmailAliasDBTypes, false, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
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

	slice := DomainEmailAliasSlice{&local}
	if err = local.L.LoadDomain(ctx, tx, false, (*[]*DomainEmailAlias)(&slice), nil); err != nil {
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

func testDomainEmailAliasToOneSetOpDomainUsingDomain(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DomainEmailAlias
	var b, c Domain

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, domainEmailAliasDBTypes, false, strmangle.SetComplement(domainEmailAliasPrimaryKeyColumns, domainEmailAliasColumnsWithoutDefault)...); err != nil {
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

		if x.R.DomainEmailAliases[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DomainID != x.ID {
			t.Error("foreign key was wrong value", a.DomainID)
		}

		if exists, err := DomainEmailAliasExists(ctx, tx, a.DomainID, a.Address); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testDomainEmailAliasesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainEmailAlias{}
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
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

func testDomainEmailAliasesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainEmailAlias{}
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DomainEmailAliasSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDomainEmailAliasesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainEmailAlias{}
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DomainEmailAliases().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	domainEmailAliasDBTypes = map[string]string{`DomainID`: `uuid`, `Address`: `character varying`, `Destinations`: `ARRAYtext`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                       = bytes.MinRead
)

func testDomainEmailAliasesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(domainEmailAliasPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(domainEmailAliasAllColumns) == len(domainEmailAliasPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DomainEmailAlias{}
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DomainEmailAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDomainEmailAliasesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(domainEmailAliasAllColumns) == len(domainEmailAliasPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DomainEmailAlias{}
	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DomainEmailAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, domainEmailAliasDBTypes, true, domainEmailAliasPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(domainEmailAliasAllColumns, domainEmailAliasPrimaryKeyColumns) {
		fields = domainEmailAliasAllColumns
	} else {
		fields = strmangle.SetComplement(
			domainEmailAliasAllColumns,
			domainEmailAliasPrimaryKeyColumns,
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

	slice := DomainEmailAliasSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDomainEmailAliasesUpsert(t *testing.T) {
	t.Parallel()

	if len(domainEmailAliasAllColumns) == len(domainEmailAliasPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DomainEmailAlias{}
	if err = randomize.Struct(seed, &o, domainEmailAliasDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DomainEmailAlias: %s", err)
	}

	count, err := DomainEmailAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, domainEmailAliasDBTypes, false, domainEmailAliasPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DomainEmailAlias struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DomainEmailAlias: %s", err)
	}

	count, err = DomainEmailAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
