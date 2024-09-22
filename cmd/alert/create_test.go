package alert

import (
	"encoding/json"
	"io"
	"os"
	"reflect"
	"testing"

	alertApi "github.com/sq325/grafanaApi/pkg/apis/alert"
)

func Test_datasourcesFromAlert(t *testing.T) {
	file, err := os.Open("/Users/sunquan/Documents/GoProjects/grafanaApi/pkg/apis/alert/api_test.go")
	if err != nil {
		t.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	bys, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("无法读取文件: %v", err)
	}

	var alerts alertApi.ProvisionedAlertRules

	err = json.Unmarshal(bys, &alerts)
	if err != nil {
		t.Fatalf("无法解析 JSON: %v", err)
	}
	type args struct {
		alert alertApi.ProvisionedAlertRule
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "Valid alert with single datasource",
			args: args{
				alert: alerts[0],
			},
			want: [][]string{
				{"datasource-uid-1", "datasource-type-1"},
			},
			wantErr: false,
		},
		{
			name: "Valid alert with multiple datasources",
			args: args{
				alert: alerts[1],
			},
			want: [][]string{
				{"datasource-uid-1", "datasource-type-1"},
				{"datasource-uid-2", "datasource-type-2"},
			},
			wantErr: false,
		},
		{
			name: "Alert with no datasources",
			args: args{
				alert: alerts[2],
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Alert with datasource UID '__expr__'",
			args: args{
				alert: alerts[3],
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := datasourcesFromAlert(tt.args.alert)
			if (err != nil) != tt.wantErr {
				t.Errorf("datasourcesFromAlert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("datasourcesFromAlert() = %v, want %v", got, tt.want)
			}
		})
	}
}
