package merchant

import "context"

type key int

// userKey is the key for user.User values in Contexts. It is
// unexported; clients use user.NewContext and user.FromContext
// instead of using this key directly.
var merchantIdKey key

// NewContext returns a new Context that carries value u.
func NewContext(ctx context.Context, merchantId int64) context.Context {
	return context.WithValue(ctx, merchantIdKey, merchantId)
}

// FromContext returns the User value stored in ctx, if any.
func FromContext(ctx context.Context) (int64, bool) {
	u, ok := ctx.Value(merchantIdKey).(int64)
	return u, ok
}
