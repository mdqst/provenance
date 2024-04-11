package sanction

import (
	"testing"

	"github.com/stretchr/testify/assert"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestKeyContainsModuleName(t *testing.T) {
	assert.Contains(t, bypassKey, ModuleName, "bypassKey")
}

func TestBypassFuncsSDKContext(t *testing.T) {
	tests := []struct {
		name string
		ctx  sdk.Context
		exp  bool
	}{
		{
			name: "brand new mostly empty context",
			ctx:  sdk.NewContext(nil, cmtproto.Header{}, false, nil),
			exp:  false,
		},
		{
			name: "context with bypass",
			ctx:  WithBypass(sdk.NewContext(nil, cmtproto.Header{}, false, nil)),
			exp:  true,
		},
		{
			name: "context with bypass on one that originally was without it",
			ctx:  WithBypass(WithoutBypass(sdk.NewContext(nil, cmtproto.Header{}, false, nil))),
			exp:  true,
		},
		{
			name: "context with bypass twice",
			ctx:  WithBypass(WithBypass(sdk.NewContext(nil, cmtproto.Header{}, false, nil))),
			exp:  true,
		},
		{
			name: "context without bypass",
			ctx:  WithoutBypass(sdk.NewContext(nil, cmtproto.Header{}, false, nil)),
			exp:  false,
		},
		{
			name: "context without bypass on one that originally had it",
			ctx:  WithoutBypass(WithBypass(sdk.NewContext(nil, cmtproto.Header{}, false, nil))),
			exp:  false,
		},
		{
			name: "context without bypass twice",
			ctx:  WithoutBypass(WithoutBypass(sdk.NewContext(nil, cmtproto.Header{}, false, nil))),
			exp:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := HasBypass(tc.ctx)
			assert.Equal(t, tc.exp, actual, "HasBypass")
		})
	}
}

func TestBypassFuncsDoNotModifyProvided(t *testing.T) {
	origCtx := sdk.NewContext(nil, cmtproto.Header{}, false, nil)
	assert.False(t, HasBypass(origCtx), "HasBypass(origCtx)")
	afterWith := WithBypass(origCtx)
	assert.True(t, HasBypass(afterWith), "HasBypass(afterWith)")
	assert.False(t, HasBypass(origCtx), "HasBypass(origCtx) after giving it to WithBypass")
	afterWithout := WithoutBypass(afterWith)
	assert.False(t, HasBypass(afterWithout), "HasBypass(afterWithout)")
	assert.True(t, HasBypass(afterWith), "HasBypass(afterWith) after giving it to WithoutBypass")
	assert.False(t, HasBypass(origCtx), "HasBypass(origCtx) after giving afterWith to WithoutBypass")
}
