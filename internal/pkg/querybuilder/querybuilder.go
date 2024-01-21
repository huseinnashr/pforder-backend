package querybuilder

import (
	"strconv"
	"strings"
)

type QueryBuilder struct {
	query string
	args  []interface{}
	count int
}

func New(query string) *QueryBuilder {
	return &QueryBuilder{query: query}
}

func (q *QueryBuilder) AddQuery(query string, arg ...interface{}) {
	query = strings.ReplaceAll(query, "?", "?$?")
	parts := strings.Split(query, "?")

	for i, part := range parts {
		if part == "$" {
			q.count++
			parts[i] = part + strconv.Itoa(q.count)
		}
	}

	q.query += " " + strings.Join(parts, "")
	q.args = append(q.args, arg...)
}

func (q *QueryBuilder) AddString(query string) {
	q.query += " " + query
}

func (q *QueryBuilder) GetQuery() string {
	return q.query
}

func (q *QueryBuilder) GetArgs() []interface{} {
	return q.args
}
