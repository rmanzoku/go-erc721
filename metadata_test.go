package erc721_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/rmanzoku/go-erc721/bind"
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
			opts := *&bind.CallOpts{Context: context.TODO()}

			got, err := e.TokenURI(&opts, tt.args.tokenId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ERC721.TokenURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
