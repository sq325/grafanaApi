package datasource

import (
	"fmt"
	"testing"
)

func Test_api_GetAll(t *testing.T) {
	api := NewApi("localhost"+":"+"3000", "glsa_KJQIl68ZSu1UPIJDltd5DVVlJUHhpjYb_9b4b0fbe")

	tests := []struct {
		name string
		want DataSources
	}{
		{
			name: "Valid Response",
			want: DataSources{ /* Add expected DataSources here */ },
		},
		{
			name: "Empty Response",
			want: DataSources{},
		},
		{
			name: "Invalid Token",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := api.GetAll()
			for i, ds := range got {
				fmt.Println(i, ds)
			}
		})
	}
}
