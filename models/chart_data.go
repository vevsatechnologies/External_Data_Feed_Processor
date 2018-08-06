// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
	"gopkg.in/volatiletech/null.v6"
)

// ChartDatum is an object representing the database table.
type ChartDatum struct {
	ID              int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Exchangeid      null.Int    `boil:"exchangeid" json:"exchangeid,omitempty" toml:"exchangeid" yaml:"exchangeid,omitempty"`
	Date            null.Time   `boil:"date" json:"date,omitempty" toml:"date" yaml:"date,omitempty"`
	High            null.String `boil:"high" json:"high,omitempty" toml:"high" yaml:"high,omitempty"`
	Low             null.String `boil:"low" json:"low,omitempty" toml:"low" yaml:"low,omitempty"`
	Open            null.String `boil:"open" json:"open,omitempty" toml:"open" yaml:"open,omitempty"`
	Close           null.String `boil:"close" json:"close,omitempty" toml:"close" yaml:"close,omitempty"`
	Volume          null.String `boil:"volume" json:"volume,omitempty" toml:"volume" yaml:"volume,omitempty"`
	Quotevolume     null.String `boil:"quotevolume" json:"quotevolume,omitempty" toml:"quotevolume" yaml:"quotevolume,omitempty"`
	Weightedaverage null.String `boil:"weightedaverage" json:"weightedaverage,omitempty" toml:"weightedaverage" yaml:"weightedaverage,omitempty"`

	R *chartDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L chartDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChartDatumColumns = struct {
	ID              string
	Exchangeid      string
	Date            string
	High            string
	Low             string
	Open            string
	Close           string
	Volume          string
	Quotevolume     string
	Weightedaverage string
}{
	ID:              "id",
	Exchangeid:      "exchangeid",
	Date:            "date",
	High:            "high",
	Low:             "low",
	Open:            "open",
	Close:           "close",
	Volume:          "volume",
	Quotevolume:     "quotevolume",
	Weightedaverage: "weightedaverage",
}

// chartDatumR is where relationships are stored.
type chartDatumR struct {
}

// chartDatumL is where Load methods for each relationship are stored.
type chartDatumL struct{}

var (
	chartDatumColumns               = []string{"id", "exchangeid", "date", "high", "low", "open", "close", "volume", "quotevolume", "weightedaverage"}
	chartDatumColumnsWithoutDefault = []string{"exchangeid", "date", "high", "low", "open", "close", "volume", "quotevolume", "weightedaverage"}
	chartDatumColumnsWithDefault    = []string{"id"}
	chartDatumPrimaryKeyColumns     = []string{"id"}
)

type (
	// ChartDatumSlice is an alias for a slice of pointers to ChartDatum.
	// This should generally be used opposed to []ChartDatum.
	ChartDatumSlice []*ChartDatum
	// ChartDatumHook is the signature for custom ChartDatum hook methods
	ChartDatumHook func(boil.Executor, *ChartDatum) error

	chartDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	chartDatumType                 = reflect.TypeOf(&ChartDatum{})
	chartDatumMapping              = queries.MakeStructMapping(chartDatumType)
	chartDatumPrimaryKeyMapping, _ = queries.BindMapping(chartDatumType, chartDatumMapping, chartDatumPrimaryKeyColumns)
	chartDatumInsertCacheMut       sync.RWMutex
	chartDatumInsertCache          = make(map[string]insertCache)
	chartDatumUpdateCacheMut       sync.RWMutex
	chartDatumUpdateCache          = make(map[string]updateCache)
	chartDatumUpsertCacheMut       sync.RWMutex
	chartDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var chartDatumBeforeInsertHooks []ChartDatumHook
var chartDatumBeforeUpdateHooks []ChartDatumHook
var chartDatumBeforeDeleteHooks []ChartDatumHook
var chartDatumBeforeUpsertHooks []ChartDatumHook

var chartDatumAfterInsertHooks []ChartDatumHook
var chartDatumAfterSelectHooks []ChartDatumHook
var chartDatumAfterUpdateHooks []ChartDatumHook
var chartDatumAfterDeleteHooks []ChartDatumHook
var chartDatumAfterUpsertHooks []ChartDatumHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ChartDatum) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range chartDatumBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ChartDatum) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range chartDatumBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ChartDatum) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range chartDatumBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ChartDatum) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range chartDatumBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ChartDatum) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range chartDatumAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ChartDatum) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range chartDatumAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ChartDatum) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range chartDatumAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ChartDatum) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range chartDatumAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ChartDatum) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range chartDatumAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChartDatumHook registers your hook function for all future operations.
func AddChartDatumHook(hookPoint boil.HookPoint, chartDatumHook ChartDatumHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		chartDatumBeforeInsertHooks = append(chartDatumBeforeInsertHooks, chartDatumHook)
	case boil.BeforeUpdateHook:
		chartDatumBeforeUpdateHooks = append(chartDatumBeforeUpdateHooks, chartDatumHook)
	case boil.BeforeDeleteHook:
		chartDatumBeforeDeleteHooks = append(chartDatumBeforeDeleteHooks, chartDatumHook)
	case boil.BeforeUpsertHook:
		chartDatumBeforeUpsertHooks = append(chartDatumBeforeUpsertHooks, chartDatumHook)
	case boil.AfterInsertHook:
		chartDatumAfterInsertHooks = append(chartDatumAfterInsertHooks, chartDatumHook)
	case boil.AfterSelectHook:
		chartDatumAfterSelectHooks = append(chartDatumAfterSelectHooks, chartDatumHook)
	case boil.AfterUpdateHook:
		chartDatumAfterUpdateHooks = append(chartDatumAfterUpdateHooks, chartDatumHook)
	case boil.AfterDeleteHook:
		chartDatumAfterDeleteHooks = append(chartDatumAfterDeleteHooks, chartDatumHook)
	case boil.AfterUpsertHook:
		chartDatumAfterUpsertHooks = append(chartDatumAfterUpsertHooks, chartDatumHook)
	}
}

// OneP returns a single chartDatum record from the query, and panics on error.
func (q chartDatumQuery) OneP() *ChartDatum {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single chartDatum record from the query.
func (q chartDatumQuery) One() (*ChartDatum, error) {
	o := &ChartDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for chart_data")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all ChartDatum records from the query, and panics on error.
func (q chartDatumQuery) AllP() ChartDatumSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all ChartDatum records from the query.
func (q chartDatumQuery) All() (ChartDatumSlice, error) {
	var o []*ChartDatum

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ChartDatum slice")
	}

	if len(chartDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all ChartDatum records in the query, and panics on error.
func (q chartDatumQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all ChartDatum records in the query.
func (q chartDatumQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count chart_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q chartDatumQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q chartDatumQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if chart_data exists")
	}

	return count > 0, nil
}

// ChartDataG retrieves all records.
func ChartDataG(mods ...qm.QueryMod) chartDatumQuery {
	return ChartData(boil.GetDB(), mods...)
}

// ChartData retrieves all the records using an executor.
func ChartData(exec boil.Executor, mods ...qm.QueryMod) chartDatumQuery {
	mods = append(mods, qm.From("\"chart_data\""))
	return chartDatumQuery{NewQuery(exec, mods...)}
}

// FindChartDatumG retrieves a single record by ID.
func FindChartDatumG(id int, selectCols ...string) (*ChartDatum, error) {
	return FindChartDatum(boil.GetDB(), id, selectCols...)
}

// FindChartDatumGP retrieves a single record by ID, and panics on error.
func FindChartDatumGP(id int, selectCols ...string) *ChartDatum {
	retobj, err := FindChartDatum(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindChartDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChartDatum(exec boil.Executor, id int, selectCols ...string) (*ChartDatum, error) {
	chartDatumObj := &ChartDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"chart_data\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(chartDatumObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from chart_data")
	}

	return chartDatumObj, nil
}

// FindChartDatumP retrieves a single record by ID with an executor, and panics on error.
func FindChartDatumP(exec boil.Executor, id int, selectCols ...string) *ChartDatum {
	retobj, err := FindChartDatum(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ChartDatum) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *ChartDatum) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *ChartDatum) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *ChartDatum) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no chart_data provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(chartDatumColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	chartDatumInsertCacheMut.RLock()
	cache, cached := chartDatumInsertCache[key]
	chartDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			chartDatumColumns,
			chartDatumColumnsWithDefault,
			chartDatumColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(chartDatumType, chartDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(chartDatumType, chartDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"chart_data\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"chart_data\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into chart_data")
	}

	if !cached {
		chartDatumInsertCacheMut.Lock()
		chartDatumInsertCache[key] = cache
		chartDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single ChartDatum record. See Update for
// whitelist behavior description.
func (o *ChartDatum) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single ChartDatum record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *ChartDatum) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the ChartDatum, and panics on error.
// See Update for whitelist behavior description.
func (o *ChartDatum) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the ChartDatum.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *ChartDatum) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	chartDatumUpdateCacheMut.RLock()
	cache, cached := chartDatumUpdateCache[key]
	chartDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			chartDatumColumns,
			chartDatumPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update chart_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"chart_data\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, chartDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(chartDatumType, chartDatumMapping, append(wl, chartDatumPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update chart_data row")
	}

	if !cached {
		chartDatumUpdateCacheMut.Lock()
		chartDatumUpdateCache[key] = cache
		chartDatumUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q chartDatumQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q chartDatumQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for chart_data")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ChartDatumSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ChartDatumSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ChartDatumSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChartDatumSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chartDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"chart_data\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, chartDatumPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in chartDatum slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ChartDatum) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *ChartDatum) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *ChartDatum) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *ChartDatum) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no chart_data provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(chartDatumColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
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
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	chartDatumUpsertCacheMut.RLock()
	cache, cached := chartDatumUpsertCache[key]
	chartDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			chartDatumColumns,
			chartDatumColumnsWithDefault,
			chartDatumColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			chartDatumColumns,
			chartDatumPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert chart_data, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(chartDatumPrimaryKeyColumns))
			copy(conflict, chartDatumPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"chart_data\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(chartDatumType, chartDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(chartDatumType, chartDatumMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert chart_data")
	}

	if !cached {
		chartDatumUpsertCacheMut.Lock()
		chartDatumUpsertCache[key] = cache
		chartDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single ChartDatum record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *ChartDatum) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single ChartDatum record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ChartDatum) DeleteG() error {
	if o == nil {
		return errors.New("models: no ChartDatum provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single ChartDatum record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *ChartDatum) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single ChartDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ChartDatum) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no ChartDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), chartDatumPrimaryKeyMapping)
	sql := "DELETE FROM \"chart_data\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from chart_data")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q chartDatumQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q chartDatumQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no chartDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from chart_data")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o ChartDatumSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o ChartDatumSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no ChartDatum slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ChartDatumSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChartDatumSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no ChartDatum slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(chartDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chartDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"chart_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, chartDatumPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from chartDatum slice")
	}

	if len(chartDatumAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *ChartDatum) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *ChartDatum) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ChartDatum) ReloadG() error {
	if o == nil {
		return errors.New("models: no ChartDatum provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ChartDatum) Reload(exec boil.Executor) error {
	ret, err := FindChartDatum(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ChartDatumSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ChartDatumSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChartDatumSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ChartDatumSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChartDatumSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	chartData := ChartDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chartDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"chart_data\".* FROM \"chart_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, chartDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&chartData)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChartDatumSlice")
	}

	*o = chartData

	return nil
}

// ChartDatumExists checks if the ChartDatum row exists.
func ChartDatumExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"chart_data\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if chart_data exists")
	}

	return exists, nil
}

// ChartDatumExistsG checks if the ChartDatum row exists.
func ChartDatumExistsG(id int) (bool, error) {
	return ChartDatumExists(boil.GetDB(), id)
}

// ChartDatumExistsGP checks if the ChartDatum row exists. Panics on error.
func ChartDatumExistsGP(id int) bool {
	e, err := ChartDatumExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ChartDatumExistsP checks if the ChartDatum row exists. Panics on error.
func ChartDatumExistsP(exec boil.Executor, id int) bool {
	e, err := ChartDatumExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
