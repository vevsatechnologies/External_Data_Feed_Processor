// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testHistoricData(t *testing.T) {
	t.Parallel()

	query := HistoricData(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testHistoricDataDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	historicDatum := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = historicDatum.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := HistoricData(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHistoricDataQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	historicDatum := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = HistoricData(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := HistoricData(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHistoricDataSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	historicDatum := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := HistoricDatumSlice{historicDatum}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := HistoricData(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testHistoricDataExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	historicDatum := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := HistoricDatumExists(tx, historicDatum.ID)
	if err != nil {
		t.Errorf("Unable to check if HistoricDatum exists: %s", err)
	}
	if !e {
		t.Errorf("Expected HistoricDatumExistsG to return true, but got false.")
	}
}
func testHistoricDataFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	historicDatum := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Insert(tx); err != nil {
		t.Error(err)
	}

	historicDatumFound, err := FindHistoricDatum(tx, historicDatum.ID)
	if err != nil {
		t.Error(err)
	}

	if historicDatumFound == nil {
		t.Error("want a record, got nil")
	}
}
func testHistoricDataBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	historicDatum := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = HistoricData(tx).Bind(historicDatum); err != nil {
		t.Error(err)
	}
}

func testHistoricDataOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	historicDatum := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := HistoricData(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testHistoricDataAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	historicDatumOne := &HistoricDatum{}
	historicDatumTwo := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatumOne, historicDatumDBTypes, false, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}
	if err = randomize.Struct(seed, historicDatumTwo, historicDatumDBTypes, false, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatumOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = historicDatumTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := HistoricData(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testHistoricDataCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	historicDatumOne := &HistoricDatum{}
	historicDatumTwo := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatumOne, historicDatumDBTypes, false, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}
	if err = randomize.Struct(seed, historicDatumTwo, historicDatumDBTypes, false, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatumOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = historicDatumTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := HistoricData(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func historicDatumBeforeInsertHook(e boil.Executor, o *HistoricDatum) error {
	*o = HistoricDatum{}
	return nil
}

func historicDatumAfterInsertHook(e boil.Executor, o *HistoricDatum) error {
	*o = HistoricDatum{}
	return nil
}

func historicDatumAfterSelectHook(e boil.Executor, o *HistoricDatum) error {
	*o = HistoricDatum{}
	return nil
}

func historicDatumBeforeUpdateHook(e boil.Executor, o *HistoricDatum) error {
	*o = HistoricDatum{}
	return nil
}

func historicDatumAfterUpdateHook(e boil.Executor, o *HistoricDatum) error {
	*o = HistoricDatum{}
	return nil
}

func historicDatumBeforeDeleteHook(e boil.Executor, o *HistoricDatum) error {
	*o = HistoricDatum{}
	return nil
}

func historicDatumAfterDeleteHook(e boil.Executor, o *HistoricDatum) error {
	*o = HistoricDatum{}
	return nil
}

func historicDatumBeforeUpsertHook(e boil.Executor, o *HistoricDatum) error {
	*o = HistoricDatum{}
	return nil
}

func historicDatumAfterUpsertHook(e boil.Executor, o *HistoricDatum) error {
	*o = HistoricDatum{}
	return nil
}

func testHistoricDataHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &HistoricDatum{}
	o := &HistoricDatum{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, historicDatumDBTypes, false); err != nil {
		t.Errorf("Unable to randomize HistoricDatum object: %s", err)
	}

	AddHistoricDatumHook(boil.BeforeInsertHook, historicDatumBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	historicDatumBeforeInsertHooks = []HistoricDatumHook{}

	AddHistoricDatumHook(boil.AfterInsertHook, historicDatumAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	historicDatumAfterInsertHooks = []HistoricDatumHook{}

	AddHistoricDatumHook(boil.AfterSelectHook, historicDatumAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	historicDatumAfterSelectHooks = []HistoricDatumHook{}

	AddHistoricDatumHook(boil.BeforeUpdateHook, historicDatumBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	historicDatumBeforeUpdateHooks = []HistoricDatumHook{}

	AddHistoricDatumHook(boil.AfterUpdateHook, historicDatumAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	historicDatumAfterUpdateHooks = []HistoricDatumHook{}

	AddHistoricDatumHook(boil.BeforeDeleteHook, historicDatumBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	historicDatumBeforeDeleteHooks = []HistoricDatumHook{}

	AddHistoricDatumHook(boil.AfterDeleteHook, historicDatumAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	historicDatumAfterDeleteHooks = []HistoricDatumHook{}

	AddHistoricDatumHook(boil.BeforeUpsertHook, historicDatumBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	historicDatumBeforeUpsertHooks = []HistoricDatumHook{}

	AddHistoricDatumHook(boil.AfterUpsertHook, historicDatumAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	historicDatumAfterUpsertHooks = []HistoricDatumHook{}
}
func testHistoricDataInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	historicDatum := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := HistoricData(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHistoricDataInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	historicDatum := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Insert(tx, historicDatumColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := HistoricData(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHistoricDataReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	historicDatum := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = historicDatum.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testHistoricDataReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	historicDatum := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := HistoricDatumSlice{historicDatum}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testHistoricDataSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	historicDatum := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := HistoricData(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	historicDatumDBTypes = map[string]string{`Exchangeid`: `integer`, `FillType`: `character varying`, `Globaltradeid`: `character varying`, `ID`: `integer`, `OrderType`: `character varying`, `Price`: `character varying`, `Quantity`: `character varying`, `Timest`: `timestamp without time zone`, `Total`: `character varying`, `Tradeid`: `character varying`}
	_                    = bytes.MinRead
)

func testHistoricDataUpdate(t *testing.T) {
	t.Parallel()

	if len(historicDatumColumns) == len(historicDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	historicDatum := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := HistoricData(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	if err = historicDatum.Update(tx); err != nil {
		t.Error(err)
	}
}

func testHistoricDataSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(historicDatumColumns) == len(historicDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	historicDatum := &HistoricDatum{}
	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := HistoricData(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, historicDatum, historicDatumDBTypes, true, historicDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(historicDatumColumns, historicDatumPrimaryKeyColumns) {
		fields = historicDatumColumns
	} else {
		fields = strmangle.SetComplement(
			historicDatumColumns,
			historicDatumPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(historicDatum))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := HistoricDatumSlice{historicDatum}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testHistoricDataUpsert(t *testing.T) {
	t.Parallel()

	if len(historicDatumColumns) == len(historicDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	historicDatum := HistoricDatum{}
	if err = randomize.Struct(seed, &historicDatum, historicDatumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = historicDatum.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert HistoricDatum: %s", err)
	}

	count, err := HistoricData(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &historicDatum, historicDatumDBTypes, false, historicDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize HistoricDatum struct: %s", err)
	}

	if err = historicDatum.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert HistoricDatum: %s", err)
	}

	count, err = HistoricData(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
