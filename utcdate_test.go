package utctime

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestParseToUTCDate(t *testing.T) {
	dt, _ := time.Parse("2006-01-02", "2022-10-13")

	type args struct {
		timeString string
	}
	tests := []struct {
		name    string
		args    args
		wantUtc UTCDate
		wantErr bool
	}{
		{
			name: "cast date",
			args: args{timeString: "2022-10-13"},
			wantUtc: UTCDate{
				Time: dt,
			},
		},
		{
			name: "cast date time gmt",
			args: args{timeString: "2022-10-13T23:00:00-03:00"},
			wantUtc: UTCDate{
				Time: dt.Add(time.Hour * 24),
			},
		},
		{
			name: "cast date time gmt",
			args: args{timeString: "2022-10-13T23:00:00Z"},
			wantUtc: UTCDate{
				Time: dt,
			},
		},
		{
			name:    "cast date time gmt",
			args:    args{timeString: "2022-10"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUtc, err := ParseToUTCDate(tt.args.timeString)
			fmt.Printf("%v", gotUtc)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToUTCDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUtc, tt.wantUtc) {
				t.Errorf("ParseToUTCDate() = %v, want %v", gotUtc, tt.wantUtc)
			}
		})
	}
}

func TestUTCDate_UnmarshalJSON(t *testing.T) {
	dt, _ := time.Parse("2006-01-02", "2022-10-13")

	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    UTCDate
	}{
		{
			name: "success",
			args: args{data: []byte("2022-10-13")},
			want: UTCDate{
				Time: dt,
			},
		},
		{
			name:    "error",
			args:    args{data: []byte("2022-10-")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utc := &UTCDate{}
			if err := utc.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UTCDate.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if *utc != tt.want {
				t.Errorf("UTCDate.UnmarshalJSON() error = %v, wantErr %v", utc, tt.want)
			}
		})
	}
}

func TestUTCDate_MarshalJSON(t *testing.T) {
	dt, _ := time.Parse("2006-01-02", "2022-10-13")
	dt1, _ := time.Parse(time.RFC3339, "2022-10-13T23:00:00-03:00")
	tests := []struct {
		name    string
		utc     UTCDate
		want    []byte
		wantErr bool
	}{
		{
			name: "success",
			utc: UTCDate{
				Time: dt,
			},
			want: []byte(`"2022-10-13"`),
		},
		{
			name: "success",
			utc: UTCDate{
				Time: dt1,
			},
			want: []byte(`"2022-10-13"`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.utc.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("UTCDate.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UTCDate.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}
