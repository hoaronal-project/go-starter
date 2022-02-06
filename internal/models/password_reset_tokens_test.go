// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

func testPasswordResetTokens(t *testing.T) {
	t.Parallel()

	query := PasswordResetTokens()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPasswordResetTokensDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasswordResetToken{}
	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
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

	count, err := PasswordResetTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPasswordResetTokensQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasswordResetToken{}
	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PasswordResetTokens().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PasswordResetTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPasswordResetTokensSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasswordResetToken{}
	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PasswordResetTokenSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PasswordResetTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPasswordResetTokensExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasswordResetToken{}
	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PasswordResetTokenExists(ctx, tx, o.Token)
	if err != nil {
		t.Errorf("Unable to check if PasswordResetToken exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PasswordResetTokenExists to return true, but got false.")
	}
}

func testPasswordResetTokensFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasswordResetToken{}
	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	passwordResetTokenFound, err := FindPasswordResetToken(ctx, tx, o.Token)
	if err != nil {
		t.Error(err)
	}

	if passwordResetTokenFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPasswordResetTokensBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasswordResetToken{}
	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PasswordResetTokens().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPasswordResetTokensOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasswordResetToken{}
	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PasswordResetTokens().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPasswordResetTokensAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	passwordResetTokenOne := &PasswordResetToken{}
	passwordResetTokenTwo := &PasswordResetToken{}
	if err = randomize.Struct(seed, passwordResetTokenOne, passwordResetTokenDBTypes, false, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}
	if err = randomize.Struct(seed, passwordResetTokenTwo, passwordResetTokenDBTypes, false, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = passwordResetTokenOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = passwordResetTokenTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PasswordResetTokens().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPasswordResetTokensCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	passwordResetTokenOne := &PasswordResetToken{}
	passwordResetTokenTwo := &PasswordResetToken{}
	if err = randomize.Struct(seed, passwordResetTokenOne, passwordResetTokenDBTypes, false, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}
	if err = randomize.Struct(seed, passwordResetTokenTwo, passwordResetTokenDBTypes, false, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = passwordResetTokenOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = passwordResetTokenTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PasswordResetTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testPasswordResetTokensInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasswordResetToken{}
	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PasswordResetTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPasswordResetTokensInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasswordResetToken{}
	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(passwordResetTokenColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PasswordResetTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPasswordResetTokenToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local PasswordResetToken
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, passwordResetTokenDBTypes, false, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
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

	slice := PasswordResetTokenSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*PasswordResetToken)(&slice), nil); err != nil {
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

func testPasswordResetTokenToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PasswordResetToken
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, passwordResetTokenDBTypes, false, strmangle.SetComplement(passwordResetTokenPrimaryKeyColumns, passwordResetTokenColumnsWithoutDefault)...); err != nil {
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

		if x.R.PasswordResetTokens[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testPasswordResetTokensReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasswordResetToken{}
	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
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

func testPasswordResetTokensReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasswordResetToken{}
	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PasswordResetTokenSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPasswordResetTokensSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasswordResetToken{}
	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PasswordResetTokens().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	passwordResetTokenDBTypes = map[string]string{`Token`: `uuid`, `ValidUntil`: `timestamp with time zone`, `UserID`: `uuid`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                         = bytes.MinRead
)

func testPasswordResetTokensUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(passwordResetTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(passwordResetTokenAllColumns) == len(passwordResetTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PasswordResetToken{}
	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PasswordResetTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPasswordResetTokensSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(passwordResetTokenAllColumns) == len(passwordResetTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PasswordResetToken{}
	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PasswordResetTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, passwordResetTokenDBTypes, true, passwordResetTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(passwordResetTokenAllColumns, passwordResetTokenPrimaryKeyColumns) {
		fields = passwordResetTokenAllColumns
	} else {
		fields = strmangle.SetComplement(
			passwordResetTokenAllColumns,
			passwordResetTokenPrimaryKeyColumns,
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

	slice := PasswordResetTokenSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPasswordResetTokensUpsert(t *testing.T) {
	t.Parallel()

	if len(passwordResetTokenAllColumns) == len(passwordResetTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PasswordResetToken{}
	if err = randomize.Struct(seed, &o, passwordResetTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PasswordResetToken: %s", err)
	}

	count, err := PasswordResetTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, passwordResetTokenDBTypes, false, passwordResetTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PasswordResetToken struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PasswordResetToken: %s", err)
	}

	count, err = PasswordResetTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
