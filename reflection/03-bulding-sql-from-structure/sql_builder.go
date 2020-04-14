package main

import (
	"errors"
	"fmt"
	"strings"
)

type SqlBuilder struct {
	fields, values []string
	table          string
}

func (b *SqlBuilder) Fields(fields ...string) *SqlBuilder {
	for _, v := range fields {
		b.fields = append(b.fields, v)
	}
	return b
}

func (b *SqlBuilder) Values(values ...string) *SqlBuilder {
	for _, v := range values {
		b.values = append(b.values, v)
	}
	return b
}

func (b *SqlBuilder) Into(table string) *SqlBuilder {
	b.table = table
	return b
}

func (b *SqlBuilder) Build() (string, error) {

	if b.table == "" {
		return "", errors.New("table is not specified")
	}

	return fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES(%s)",
		b.table,
		strings.Join(b.fields, ","),
		strings.Join(b.values, ","),
	), nil

}
