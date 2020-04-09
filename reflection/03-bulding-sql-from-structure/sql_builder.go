package main

import "fmt"

type SqlBuilder struct {
	fields []string
	values []interface{}
	table  string
}

func (b *SqlBuilder) Fields(fields ...string) *SqlBuilder {
	for _, v := range fields {
		b.fields = append(b.fields, v)
	}
	return b
}

func (b *SqlBuilder) Values(values ...interface{}) *SqlBuilder {
	for _, v := range values {
		b.values = append(b.values, v)
	}
	return b
}

func (b *SqlBuilder) Into(table string) *SqlBuilder {
	b.table = table
	return b
}

func (b *SqlBuilder) Build() string {
	fmt.Println(b)
	return ""
}
