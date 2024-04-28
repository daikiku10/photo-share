// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package internal

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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Photos is an object representing the database table.
type Photos struct { // 主キー
	ID string `boil:"id" json:"id" toml:"id" yaml:"id"`
	// 作成日時
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// 写真タイトル
	Title string `boil:"title" json:"title" toml:"title" yaml:"title"`
	// 説明
	Description string `boil:"description" json:"description" toml:"description" yaml:"description"`
	// 写真URL
	ImageUrl string `boil:"imageUrl" json:"imageUrl" toml:"imageUrl" yaml:"imageUrl"`
	// 投稿者ID
	AuthorId string `boil:"authorId" json:"authorId" toml:"authorId" yaml:"authorId"`
	// カテゴリID
	CategoryId string `boil:"categoryId" json:"categoryId" toml:"categoryId" yaml:"categoryId"`

	R *photosR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L photosL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PhotosColumns = struct {
	ID          string
	CreatedAt   string
	Title       string
	Description string
	ImageUrl    string
	AuthorId    string
	CategoryId  string
}{
	ID:          "id",
	CreatedAt:   "created_at",
	Title:       "title",
	Description: "description",
	ImageUrl:    "imageUrl",
	AuthorId:    "authorId",
	CategoryId:  "categoryId",
}

var PhotosTableColumns = struct {
	ID          string
	CreatedAt   string
	Title       string
	Description string
	ImageUrl    string
	AuthorId    string
	CategoryId  string
}{
	ID:          "photos.id",
	CreatedAt:   "photos.created_at",
	Title:       "photos.title",
	Description: "photos.description",
	ImageUrl:    "photos.imageUrl",
	AuthorId:    "photos.authorId",
	CategoryId:  "photos.categoryId",
}

// Generated where

var PhotosWhere = struct {
	ID          whereHelperstring
	CreatedAt   whereHelpertime_Time
	Title       whereHelperstring
	Description whereHelperstring
	ImageUrl    whereHelperstring
	AuthorId    whereHelperstring
	CategoryId  whereHelperstring
}{
	ID:          whereHelperstring{field: "`photos`.`id`"},
	CreatedAt:   whereHelpertime_Time{field: "`photos`.`created_at`"},
	Title:       whereHelperstring{field: "`photos`.`title`"},
	Description: whereHelperstring{field: "`photos`.`description`"},
	ImageUrl:    whereHelperstring{field: "`photos`.`imageUrl`"},
	AuthorId:    whereHelperstring{field: "`photos`.`authorId`"},
	CategoryId:  whereHelperstring{field: "`photos`.`categoryId`"},
}

// PhotosRels is where relationship names are stored.
var PhotosRels = struct {
}{}

// photosR is where relationships are stored.
type photosR struct {
}

// NewStruct creates a new relationship struct
func (*photosR) NewStruct() *photosR {
	return &photosR{}
}

// photosL is where Load methods for each relationship are stored.
type photosL struct{}

var (
	photosAllColumns            = []string{"id", "created_at", "title", "description", "imageUrl", "authorId", "categoryId"}
	photosColumnsWithoutDefault = []string{"id", "title", "description", "imageUrl", "authorId", "categoryId"}
	photosColumnsWithDefault    = []string{"created_at"}
	photosPrimaryKeyColumns     = []string{"id"}
	photosGeneratedColumns      = []string{}
)

type (
	// PhotosSlice is an alias for a slice of pointers to Photos.
	// This should almost always be used instead of []Photos.
	PhotosSlice []*Photos
	// PhotosHook is the signature for custom Photos hook methods
	PhotosHook func(context.Context, boil.ContextExecutor, *Photos) error

	photosQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	photosType                 = reflect.TypeOf(&Photos{})
	photosMapping              = queries.MakeStructMapping(photosType)
	photosPrimaryKeyMapping, _ = queries.BindMapping(photosType, photosMapping, photosPrimaryKeyColumns)
	photosInsertCacheMut       sync.RWMutex
	photosInsertCache          = make(map[string]insertCache)
	photosUpdateCacheMut       sync.RWMutex
	photosUpdateCache          = make(map[string]updateCache)
	photosUpsertCacheMut       sync.RWMutex
	photosUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var photosAfterSelectMu sync.Mutex
var photosAfterSelectHooks []PhotosHook

var photosBeforeInsertMu sync.Mutex
var photosBeforeInsertHooks []PhotosHook
var photosAfterInsertMu sync.Mutex
var photosAfterInsertHooks []PhotosHook

var photosBeforeUpdateMu sync.Mutex
var photosBeforeUpdateHooks []PhotosHook
var photosAfterUpdateMu sync.Mutex
var photosAfterUpdateHooks []PhotosHook

var photosBeforeDeleteMu sync.Mutex
var photosBeforeDeleteHooks []PhotosHook
var photosAfterDeleteMu sync.Mutex
var photosAfterDeleteHooks []PhotosHook

var photosBeforeUpsertMu sync.Mutex
var photosBeforeUpsertHooks []PhotosHook
var photosAfterUpsertMu sync.Mutex
var photosAfterUpsertHooks []PhotosHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Photos) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photosAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Photos) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photosBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Photos) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photosAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Photos) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photosBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Photos) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photosAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Photos) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photosBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Photos) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photosAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Photos) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photosBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Photos) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photosAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPhotosHook registers your hook function for all future operations.
func AddPhotosHook(hookPoint boil.HookPoint, photosHook PhotosHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		photosAfterSelectMu.Lock()
		photosAfterSelectHooks = append(photosAfterSelectHooks, photosHook)
		photosAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		photosBeforeInsertMu.Lock()
		photosBeforeInsertHooks = append(photosBeforeInsertHooks, photosHook)
		photosBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		photosAfterInsertMu.Lock()
		photosAfterInsertHooks = append(photosAfterInsertHooks, photosHook)
		photosAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		photosBeforeUpdateMu.Lock()
		photosBeforeUpdateHooks = append(photosBeforeUpdateHooks, photosHook)
		photosBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		photosAfterUpdateMu.Lock()
		photosAfterUpdateHooks = append(photosAfterUpdateHooks, photosHook)
		photosAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		photosBeforeDeleteMu.Lock()
		photosBeforeDeleteHooks = append(photosBeforeDeleteHooks, photosHook)
		photosBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		photosAfterDeleteMu.Lock()
		photosAfterDeleteHooks = append(photosAfterDeleteHooks, photosHook)
		photosAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		photosBeforeUpsertMu.Lock()
		photosBeforeUpsertHooks = append(photosBeforeUpsertHooks, photosHook)
		photosBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		photosAfterUpsertMu.Lock()
		photosAfterUpsertHooks = append(photosAfterUpsertHooks, photosHook)
		photosAfterUpsertMu.Unlock()
	}
}

// One returns a single photos record from the query.
func (q photosQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Photos, error) {
	o := &Photos{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "internal: failed to execute a one query for photos")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Photos records from the query.
func (q photosQuery) All(ctx context.Context, exec boil.ContextExecutor) (PhotosSlice, error) {
	var o []*Photos

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "internal: failed to assign all query results to Photos slice")
	}

	if len(photosAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Photos records in the query.
func (q photosQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "internal: failed to count photos rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q photosQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "internal: failed to check if photos exists")
	}

	return count > 0, nil
}

// PluralPhotos retrieves all the records using an executor.
func PluralPhotos(mods ...qm.QueryMod) photosQuery {
	mods = append(mods, qm.From("`photos`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`photos`.*"})
	}

	return photosQuery{q}
}

// FindPhotos retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPhotos(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Photos, error) {
	photosObj := &Photos{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `photos` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, photosObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "internal: unable to select from photos")
	}

	if err = photosObj.doAfterSelectHooks(ctx, exec); err != nil {
		return photosObj, err
	}

	return photosObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Photos) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("internal: no photos provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(photosColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	photosInsertCacheMut.RLock()
	cache, cached := photosInsertCache[key]
	photosInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			photosAllColumns,
			photosColumnsWithDefault,
			photosColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(photosType, photosMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(photosType, photosMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `photos` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `photos` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `photos` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, photosPrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "internal: unable to insert into photos")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "internal: unable to populate default values for photos")
	}

CacheNoHooks:
	if !cached {
		photosInsertCacheMut.Lock()
		photosInsertCache[key] = cache
		photosInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Photos.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Photos) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	photosUpdateCacheMut.RLock()
	cache, cached := photosUpdateCache[key]
	photosUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			photosAllColumns,
			photosPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("internal: unable to update photos, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `photos` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, photosPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(photosType, photosMapping, append(wl, photosPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "internal: unable to update photos row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "internal: failed to get rows affected by update for photos")
	}

	if !cached {
		photosUpdateCacheMut.Lock()
		photosUpdateCache[key] = cache
		photosUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q photosQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "internal: unable to update all for photos")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "internal: unable to retrieve rows affected for photos")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PhotosSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("internal: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), photosPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `photos` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, photosPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "internal: unable to update all in photos slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "internal: unable to retrieve rows affected all in update all photos")
	}
	return rowsAff, nil
}

var mySQLPhotosUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Photos) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("internal: no photos provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(photosColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPhotosUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	photosUpsertCacheMut.RLock()
	cache, cached := photosUpsertCache[key]
	photosUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			photosAllColumns,
			photosColumnsWithDefault,
			photosColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			photosAllColumns,
			photosPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("internal: unable to upsert photos, could not build update column list")
		}

		ret := strmangle.SetComplement(photosAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`photos`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `photos` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(photosType, photosMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(photosType, photosMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "internal: unable to upsert for photos")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(photosType, photosMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "internal: unable to retrieve unique values for photos")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "internal: unable to populate default values for photos")
	}

CacheNoHooks:
	if !cached {
		photosUpsertCacheMut.Lock()
		photosUpsertCache[key] = cache
		photosUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Photos record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Photos) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("internal: no Photos provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), photosPrimaryKeyMapping)
	sql := "DELETE FROM `photos` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "internal: unable to delete from photos")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "internal: failed to get rows affected by delete for photos")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q photosQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("internal: no photosQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "internal: unable to delete all from photos")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "internal: failed to get rows affected by deleteall for photos")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PhotosSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(photosBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), photosPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `photos` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, photosPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "internal: unable to delete all from photos slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "internal: failed to get rows affected by deleteall for photos")
	}

	if len(photosAfterDeleteHooks) != 0 {
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
func (o *Photos) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPhotos(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhotosSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PhotosSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), photosPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `photos`.* FROM `photos` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, photosPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "internal: unable to reload all in PhotosSlice")
	}

	*o = slice

	return nil
}

// PhotosExists checks if the Photos row exists.
func PhotosExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `photos` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "internal: unable to check if photos exists")
	}

	return exists, nil
}

// Exists checks if the Photos row exists.
func (o *Photos) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PhotosExists(ctx, exec, o.ID)
}
