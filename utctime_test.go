package utctime

import (
	"reflect"
	"testing"
	"time"
)

func TestParseToUTCInLocation(t *testing.T) {
	timeTest := time.Date(2022, 05, 12, 17, 59, 24, 0, time.UTC)
	type args struct {
		timeString string
		location   string
		layout     string
	}
	tests := []struct {
		name    string
		args    args
		wantUtc UTCTime
		wantErr bool
	}{
		{
			name: "UTC TIME TEST",
			args: args{
				layout:     "2006-01-02T15:04:05Z",
				timeString: "2022-05-12T14:59:24Z",
				location:   "America/Sao_Paulo",
			},
			wantUtc: UTCTime{
				Time: timeTest,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUtc, err := ParseToUTCInLocation(tt.args.layout, tt.args.timeString, tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToUTCInLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUtc, tt.wantUtc) {
				t.Errorf("ParseToUTCInLocation() = %v, want %v", gotUtc, tt.wantUtc)
			}
		})
	}
}
