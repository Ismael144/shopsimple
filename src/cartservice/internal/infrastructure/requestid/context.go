package requestid

import "context"

type ctxKey struct{}

func key() ctxKey { return ctxKey{} }

func With(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ctxKey{}, id)
}

func From(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(ctxKey{}).(string)
	return id, ok 
}