package time

import (
	"fmt"
	"testing"
)

func TestNowTimeString(t *testing.T) {
	timeString := NowTimeString()
	fmt.Println(timeString)
	time, err := ParseTime(timeString)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(time)
}

func TestFixedZoneTime(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"",
			args{
				s: "2020-05-11T11:46:13.409+00:00",
			},
			"2020-05-11T19:46:13.409+08:00",
			false,
		},
		{
			"",
			args{
				s: "2020-05-11T11:46:13.409-08:00",
			},
			"2020-05-12T03:46:13.409+08:00",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FixedZoneTime(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("FixedZoneTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FixedZoneTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDurationOfTime(t *testing.T) {

	d, err := DurationOfTime("2020-05-11T11:46:13.409-08:00", "2020-05-11T11:46:03.409-08:00")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(d)
}
