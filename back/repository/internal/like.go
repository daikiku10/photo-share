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

// Like is an object representing the database table.
type Like struct { // 主キー
	ID string `boil:"id" json:"id" toml:"id" yaml:"id"`
	// 作成日時
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// ユーザーID
	UserId string `boil:"userId" json:"userId" toml:"userId" yaml:"userId"`
	// 写真ID
	PhotoId string `boil:"photoId" json:"photoId" toml:"photoId" yaml:"photoId"`

	R *likeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L likeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LikeColumns = struct {
	ID        string
	CreatedAt string
	UserId    string
	PhotoId   string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UserId:    "userId",
	PhotoId:   "photoId",
}

var LikeTableColumns = struct {
	ID        string
	CreatedAt string
	UserId    string
	PhotoId   string
}{
	ID:        "like.id",
	CreatedAt: "like.created_at",
	UserId:    "like.userId",
	PhotoId:   "like.photoId",
}

// Generated where

var LikeWhere = struct {
	ID        whereHelperstring
	CreatedAt whereHelpertime_Time
	UserId    whereHelperstring
	PhotoId   whereHelperstring
}{
	ID:        whereHelperstring{field: "`like`.`id`"},
	CreatedAt: whereHelpertime_Time{field: "`like`.`created_at`"},
	UserId:    whereHelperstring{field: "`like`.`userId`"},
	PhotoId:   whereHelperstring{field: "`like`.`photoId`"},
}

// LikeRels is where relationship names are stored.
var LikeRels = struct {
}{}

// likeR is where relationships are stored.
type likeR struct {
}

// NewStruct creates a new relationship struct
func (*likeR) NewStruct() *likeR {
	return &likeR{}
}

// likeL is where Load methods for each relationship are stored.
type likeL struct{}

var (
	likeAllColumns            = []string{"id", "created_at", "userId", "photoId"}
	likeColumnsWithoutDefault = []string{"id", "userId", "photoId"}
	likeColumnsWithDefault    = []string{"created_at"}
	likePrimaryKeyColumns     = []string{"id"}
	likeGeneratedColumns      = []string{}
)

type (
	// LikeSlice is an alias for a slice of pointers to Like.
	// This should almost always be used instead of []Like.
	LikeSlice []*Like
	// LikeHook is the signature for custom Like hook methods
	LikeHook func(context.Context, boil.ContextExecutor, *Like) error

	likeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	likeType                 = reflect.TypeOf(&Like{})
	likeMapping              = queries.MakeStructMapping(likeType)
	likePrimaryKeyMapping, _ = queries.BindMapping(likeType, likeMapping, likePrimaryKeyColumns)
	likeInsertCacheMut       sync.RWMutex
	likeInsertCache          = make(map[string]insertCache)
	likeUpdateCacheMut       sync.RWMutex
	likeUpdateCache          = make(map[string]updateCache)
	likeUpsertCacheMut       sync.RWMutex
	likeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var likeAfterSelectMu sync.Mutex
var likeAfterSelectHooks []LikeHook

var likeBeforeInsertMu sync.Mutex
var likeBeforeInsertHooks []LikeHook
var likeAfterInsertMu sync.Mutex
var likeAfterInsertHooks []LikeHook

var likeBeforeUpdateMu sync.Mutex
var likeBeforeUpdateHooks []LikeHook
var likeAfterUpdateMu sync.Mutex
var likeAfterUpdateHooks []LikeHook

var likeBeforeDeleteMu sync.Mutex
var likeBeforeDeleteHooks []LikeHook
var likeAfterDeleteMu sync.Mutex
var likeAfterDeleteHooks []LikeHook

var likeBeforeUpsertMu sync.Mutex
var likeBeforeUpsertHooks []LikeHook
var likeAfterUpsertMu sync.Mutex
var likeAfterUpsertHooks []LikeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Like) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range likeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Like) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range likeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Like) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range likeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Like) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range likeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Like) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range likeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Like) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range likeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Like) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range likeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Like) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range likeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Like) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range likeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLikeHook registers your hook function for all future operations.
func AddLikeHook(hookPoint boil.HookPoint, likeHook LikeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		likeAfterSelectMu.Lock()
		likeAfterSelectHooks = append(likeAfterSelectHooks, likeHook)
		likeAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		likeBeforeInsertMu.Lock()
		likeBeforeInsertHooks = append(likeBeforeInsertHooks, likeHook)
		likeBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		likeAfterInsertMu.Lock()
		likeAfterInsertHooks = append(likeAfterInsertHooks, likeHook)
		likeAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		likeBeforeUpdateMu.Lock()
		likeBeforeUpdateHooks = append(likeBeforeUpdateHooks, likeHook)
		likeBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		likeAfterUpdateMu.Lock()
		likeAfterUpdateHooks = append(likeAfterUpdateHooks, likeHook)
		likeAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		likeBeforeDeleteMu.Lock()
		likeBeforeDeleteHooks = append(likeBeforeDeleteHooks, likeHook)
		likeBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		likeAfterDeleteMu.Lock()
		likeAfterDeleteHooks = append(likeAfterDeleteHooks, likeHook)
		likeAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		likeBeforeUpsertMu.Lock()
		likeBeforeUpsertHooks = append(likeBeforeUpsertHooks, likeHook)
		likeBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		likeAfterUpsertMu.Lock()
		likeAfterUpsertHooks = append(likeAfterUpsertHooks, likeHook)
		likeAfterUpsertMu.Unlock()
	}
}

// One returns a single like record from the query.
func (q likeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Like, error) {
	o := &Like{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "internal: failed to execute a one query for like")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Like records from the query.
func (q likeQuery) All(ctx context.Context, exec boil.ContextExecutor) (LikeSlice, error) {
	var o []*Like

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "internal: failed to assign all query results to Like slice")
	}

	if len(likeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Like records in the query.
func (q likeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "internal: failed to count like rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q likeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "internal: failed to check if like exists")
	}

	return count > 0, nil
}

// Likes retrieves all the records using an executor.
func Likes(mods ...qm.QueryMod) likeQuery {
	mods = append(mods, qm.From("`like`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`like`.*"})
	}

	return likeQuery{q}
}

// FindLike retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLike(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Like, error) {
	likeObj := &Like{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `like` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, likeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "internal: unable to select from like")
	}

	if err = likeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return likeObj, err
	}

	return likeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Like) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("internal: no like provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(likeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	likeInsertCacheMut.RLock()
	cache, cached := likeInsertCache[key]
	likeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			likeAllColumns,
			likeColumnsWithDefault,
			likeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(likeType, likeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(likeType, likeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `like` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `like` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `like` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, likePrimaryKeyColumns))
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
		return errors.Wrap(err, "internal: unable to insert into like")
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
		return errors.Wrap(err, "internal: unable to populate default values for like")
	}

CacheNoHooks:
	if !cached {
		likeInsertCacheMut.Lock()
		likeInsertCache[key] = cache
		likeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Like.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Like) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	likeUpdateCacheMut.RLock()
	cache, cached := likeUpdateCache[key]
	likeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			likeAllColumns,
			likePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("internal: unable to update like, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `like` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, likePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(likeType, likeMapping, append(wl, likePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "internal: unable to update like row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "internal: failed to get rows affected by update for like")
	}

	if !cached {
		likeUpdateCacheMut.Lock()
		likeUpdateCache[key] = cache
		likeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q likeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "internal: unable to update all for like")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "internal: unable to retrieve rows affected for like")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LikeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), likePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `like` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, likePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "internal: unable to update all in like slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "internal: unable to retrieve rows affected all in update all like")
	}
	return rowsAff, nil
}

var mySQLLikeUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Like) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("internal: no like provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(likeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLLikeUniqueColumns, o)

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

	likeUpsertCacheMut.RLock()
	cache, cached := likeUpsertCache[key]
	likeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			likeAllColumns,
			likeColumnsWithDefault,
			likeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			likeAllColumns,
			likePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("internal: unable to upsert like, could not build update column list")
		}

		ret := strmangle.SetComplement(likeAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`like`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `like` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(likeType, likeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(likeType, likeMapping, ret)
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
		return errors.Wrap(err, "internal: unable to upsert for like")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(likeType, likeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "internal: unable to retrieve unique values for like")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "internal: unable to populate default values for like")
	}

CacheNoHooks:
	if !cached {
		likeUpsertCacheMut.Lock()
		likeUpsertCache[key] = cache
		likeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Like record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Like) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("internal: no Like provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), likePrimaryKeyMapping)
	sql := "DELETE FROM `like` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "internal: unable to delete from like")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "internal: failed to get rows affected by delete for like")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q likeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("internal: no likeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "internal: unable to delete all from like")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "internal: failed to get rows affected by deleteall for like")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LikeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(likeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), likePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `like` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, likePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "internal: unable to delete all from like slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "internal: failed to get rows affected by deleteall for like")
	}

	if len(likeAfterDeleteHooks) != 0 {
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
func (o *Like) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLike(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LikeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LikeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), likePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `like`.* FROM `like` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, likePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "internal: unable to reload all in LikeSlice")
	}

	*o = slice

	return nil
}

// LikeExists checks if the Like row exists.
func LikeExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `like` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "internal: unable to check if like exists")
	}

	return exists, nil
}

// Exists checks if the Like row exists.
func (o *Like) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return LikeExists(ctx, exec, o.ID)
}
