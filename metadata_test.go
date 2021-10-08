package erc721_test

import (
	"context"
	"math/big"
	"testing"
)

func TestTokenURI(t *testing.T) {

	type args struct {
		tokenId *big.Int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				tokenId: big.NewInt(120410001),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			got, err := e.TokenURI(ctx, tt.args.tokenId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ERC721.TokenURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
