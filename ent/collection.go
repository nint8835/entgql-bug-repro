// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ot *OtherTestQuery) CollectFields(ctx context.Context, satisfies ...string) *OtherTestQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		ot = ot.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return ot
}

func (ot *OtherTestQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *OtherTestQuery {
	return ot
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TestQuery) CollectFields(ctx context.Context, satisfies ...string) *TestQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return t
}

func (t *TestQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *TestQuery {
	return t
}
