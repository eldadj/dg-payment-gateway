package merchant

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := context.Background()
	t.Run("value not in context", func(t *testing.T) {
		_, valid := FromContext(ctx)
		assert.False(t, valid)
	})
	t.Run("set value in context", func(t *testing.T) {
		merchantId := int64(2)
		newCtx := NewContext(ctx, merchantId)
		//the context passed should have been modified
		assert.NotEqual(t, ctx, newCtx)
		//check if set in returned context
		ctxMerchantId, valid := FromContext(newCtx)
		assert.True(t, valid)
		assert.Equal(t, merchantId, ctxMerchantId)
		//original context shouldn't have the value
		_, valid = FromContext(ctx)
		assert.False(t, valid)
	})
}
