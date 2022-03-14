package json_util_test

import (
	"testing"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/verify/json_util"
)

func TestSortJsonByKeys(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		json string
		want string
	}{
		{
			name: "empty",
			json: `{}`,
			want: `{}`,
		},
		{
			name: "abc",
			json: `{"a":1,"b":2,"c":3}`,
			want: `{"a":1,"b":2,"c":3}`,
		},
		{
			name: "cba",
			json: `{"c":1,"b":2,"a":3}`,
			want: `{"a":3,"b":2,"c":1}`,
		},
		{
			name: "nested",
			json: `{"d":{"c":1,"a":2},"b":3}`,
			want: `{"b":3,"d":{"c":2,"a":1}}`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := json_util.SortJsonByKeys([]byte(tt.json))
			if err != nil {
				t.Errorf("SortJsonByKeys() error = %v", err)

				return
			}

			if string(got) != tt.want {
				t.Errorf("SortJsonByKeys() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
