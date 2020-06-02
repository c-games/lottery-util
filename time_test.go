package util

import (
	"reflect"
	"testing"
	"time"
)

func TestGetTimeByParam(t *testing.T) {

	//要用now ，否則日期不對
	now := time.Now()

	dateStr0 := now.Format(TimeDateLayoutString) + " " + "10:10:00"
	d0, _ := time.ParseInLocation(TimeLayoutString, dateStr0, GetDefaultLocation())

	dateStr1 := now.Format(TimeDateLayoutString) + " " + "23:01:01"
	d1, _ := time.ParseInLocation(TimeLayoutString, dateStr1, GetDefaultLocation())

	dateStr2 := now.Format(TimeDateLayoutString) + " " + "09:59:59"
	d2, _ := time.ParseInLocation(TimeLayoutString, dateStr2, GetDefaultLocation())

	type args struct {
		hour string
		min  string
		sec  string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			"0",
			args{"10", "10", "00"},
			d0,
			false,
		},
		{
			"1",
			args{"23", "01", "01"},
			d1,
			false,
		},
		{
			"2",
			args{"09", "59", "59"},
			d2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := GetTimeByParam(time.Now(), tt.args.hour, tt.args.min, tt.args.sec)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTimeByParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTimeByParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTodayTimeByBeginOrEnd(t *testing.T) {

	now := time.Now()

	dateStr0 := now.Format(TimeDateLayoutString) + " " + "09:10:00"
	d0, _ := time.ParseInLocation(TimeLayoutString, dateStr0, GetDefaultLocation())

	dateStr1 := now.Format(TimeDateLayoutString) + " " + "23:01:00"
	d1, _ := time.ParseInLocation(TimeLayoutString, dateStr1, GetDefaultLocation())

	dateStr2 := now.Format(TimeDateLayoutString) + " " + "00:00:00"
	d2, _ := time.ParseInLocation(TimeLayoutString, dateStr2, GetDefaultLocation())

	type args struct {
		time          time.Time
		beginOrEndStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			"0",
			args{now, "09:10:00"},
			d0,
			false,
		},
		{
			"1",
			args{now, "23:01:00"},
			d1,
			false,
		},
		{
			"2",
			args{now, "00:00:00"},
			d2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTimeByHourString(tt.args.time, tt.args.beginOrEndStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTodayTimeByBeginOrEnd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTodayTimeByBeginOrEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

//潤2月=2月有29日
//1.逢4的倍數閏，例如：西元1996、2012、2016年等，為4的倍數，故為閏年。
//2.逢100的倍數不閏，例如：西元1800、1900、2100年，為100的倍數，當年不閏年。
//3.逢400的倍數閏，例如：西元1600、2000、2400年，為400的倍數，有閏年。
//4.逢4000的倍數不閏，例如：西元4000、8000年，不閏年。

func TestSubDate(t *testing.T) {

	begin0, _ := GetTimeWithDefaultLocation("2019-06-01 00:00:00")
	end0, _ := GetTimeWithDefaultLocation("2019-06-02 00:00:00")

	begin1, _ := GetTimeWithDefaultLocation("2019-01-01 05:10:30")
	end1, _ := GetTimeWithDefaultLocation("2019-01-30 23:00:00")

	//4的倍數，潤2月
	//2
	begin2, _ := GetTimeWithDefaultLocation("2016-02-28 05:10:30")
	end2, _ := GetTimeWithDefaultLocation("2016-03-01 23:00:00")

	//100的倍數，不潤
	//1
	begin3, _ := GetTimeWithDefaultLocation("2100-02-28 05:10:30")
	end3, _ := GetTimeWithDefaultLocation("2100-03-01 23:00:00")

	//400的倍數，潤2月
	//2
	begin4, _ := GetTimeWithDefaultLocation("2000-02-28 05:10:30")
	end4, _ := GetTimeWithDefaultLocation("2000-03-01 23:00:00")

	type args struct {
		begin time.Time
		end   time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"0",
			args{begin0, end0},
			1,
			false,
		},
		{
			"1",
			args{begin1, end1},
			29,
			false,
		},
		{
			"2",
			args{begin2, end2},
			2,
			false,
		},
		{
			"3",
			args{begin3, end3},
			1,
			false,
		},
		{
			"4",
			args{begin4, end4},
			2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SubDate(tt.args.begin, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SubDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualDate(t *testing.T) {

	begin0, _ := GetTimeWithDefaultLocation("2019-06-01 00:00:00")
	end0, _ := GetTimeWithDefaultLocation("2019-06-01 00:00:00")

	begin1, _ := GetTimeWithDefaultLocation("2019-01-01 05:10:30")
	end1, _ := GetTimeWithDefaultLocation("2019-01-30 23:00:00")

	type args struct {
		date1 time.Time
		date2 time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"0",
			args{begin0, end0},
			true,
			false,
		},
		{
			"1",
			args{begin1, end1},
			false,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EqualDate(tt.args.date1, tt.args.date2)
			if (err != nil) != tt.wantErr {
				t.Errorf("EqualDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EqualDate() = %v, want %v", got, tt.want)
			}
		})
	}
}





func TestParseWithLayoutString(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"0",
			args{GetNowWithDefaultLocation()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			 got := ParseWithLayoutString(tt.args.t)
				t.Logf("ParseWithLayoutString() = %v", got)
			t.Log(GetNowWithDefaultLocation().String())
		})
	}
}

func TestNow(t *testing.T){
	n:=GetNowWithDefaultLocation()
	t.Logf("GetNowWithDefaultLocation() %s",n.String())
}

func TestIsBetween(t *testing.T) {

	b0,_:=GetTimeByTimeString("2019-10-10 00:00:10")
	e0,_:=GetTimeByTimeString("2019-10-11 23:33:00")
	t0,_:=GetTimeByTimeString("2019-10-11 11:33:00")

	if r0,_:=IsBetween(t0,b0,e0);!r0{
		t.Errorf("want %v got %v",!r0,r0)
	}

	b1,_:=GetTimeByTimeString("2019-10-10 00:00:10")
	e1,_:=GetTimeByTimeString("2019-10-11 23:33:00")
	t1,_:=GetTimeByTimeString("2019-10-11 23:44:00")

	if r0,_:=IsBetween(t1,b1,e1);r0{
		t.Errorf("want %v got %v",!r0,r0)
	}
}