package verify_test

import (
	"testing"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/verify"
)

func TestSignature(t *testing.T) {
	t.Parallel()

	type args struct {
		jsonBytes   []byte
		address     string
		instanceUrl string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				jsonBytes:   []byte(`{}`),
				address:     "",
				instanceUrl: "",
			},
			want: false,
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got, _ := verify.Signature(tt.args.jsonBytes, tt.args.address, tt.args.instanceUrl); got != tt.want {
				t.Errorf("Signature() = %v, want %v", got, tt.want)
			}
		})
	}
}
