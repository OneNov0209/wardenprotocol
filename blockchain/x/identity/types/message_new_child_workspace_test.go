package types

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/qredo/fusionchain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgNewChildWorkspace_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgNewChildWorkspace
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgNewChildWorkspace{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgNewChildWorkspace{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
