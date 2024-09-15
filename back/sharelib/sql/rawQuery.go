package sql

import "strings"

type RawQuery struct {
	Query string
	Args  []any
}

func NewRawQuery(query string) *RawQuery {
	return &RawQuery{
		Query: query,
		Args:  make([]any, 0, 180), // argsの容量を事前に確保しておく
	}
}

// Add スペースと新しいクエリを結合する
func Add(raw *RawQuery, query string) {
	raw.Query += " " + query
}

// AddSingleParam スペースと新しいクエリを結合し、単体のparamをArgsに追加する
func AddSingleParam[T any](raw *RawQuery, query string, param T) {
	raw.Query += " " + query
	raw.Args = append(raw.Args, param)
}

// AddMultiParams スペースと新しいクエリを結合し、複数のparamsをArgsに追加する
func AddMultiParams[T any](raw *RawQuery, query string, params []T) {
	raw.Query += " " + query
	raw.Args = append(raw.Args, toAny[T](params)...)
}

// toAny 受け取った配列をAny配列に変換する
func toAny[T any](s []T) []any {
	a := make([]any, len(s))
	for i, v := range s {
		a[i] = v
	}
	return a
}

// InCondition SQLクエリのIN条件に使用するプレイスホルダーを生成します
func InCondition[T any](params []T) string {
	placeholders := make([]string, len(params))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	return "(" + strings.Join(placeholders, ",") + ")"
}

// Combine 複数のRawQueryを組み合わせる
func Combine(rawQueries []*RawQuery) (string, []any) {
	rawSql := ""
	args := make([]any, 0, len(rawQueries)*5) // argsの容量を事前に確保しておく(5は適当)
	for _, rawQuery := range rawQueries {
		rawSql += " " + rawQuery.Query
		args = append(args, rawQuery.Args...)
	}
	return rawSql, args
}

// CombineForUnionALL 複数のRawQueryにUNION ALLを付与して組み合わせる
func CombineForUnionALL(rawQueries []*RawQuery) (string, []any) {
	rawSql := ""
	args := make([]any, 0, len(rawQueries)*5) // argsの容量を事前に確保しておく(5は適当)
	for i, rawQuery := range rawQueries {
		if i > 0 {
			rawSql += " UNION ALL"
		}
		rawSql += " " + rawQuery.Query
		args = append(args, rawQuery.Args...)
	}
	return rawSql, args
}
