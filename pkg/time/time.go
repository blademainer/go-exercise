package time

import "time"

// Rfc3399Mills Rfc3399的毫秒格式
const Rfc3399Mills = "2006-01-02T15:04:05.999Z07:00"

// NowTimeString 返回Rfc3399Mills格式时间
func NowTimeString() string {
	return GetFixedZoneTimeString(time.Now())
}

// ParseTime 解析Rfc3399Mills格式时间
func ParseTime(s string) (time.Time, error) {
	t, err := time.Parse(Rfc3399Mills, s)
	return t, err
}

// FixedZoneTimeFromTimeAndLayout 解析字符串为东八区日期。
func FixedZoneTimeFromTimeAndLayout(s string, layout string) (string, error) {
	parseTime, err := time.Parse(layout, s)
	if err != nil {
		return "", err
	}
	timeString := GetFixedZoneTimeString(parseTime)
	return timeString, nil
}

// FixedZoneTime 解析字符串为东八区日期。
//
// 日期串必须是 2006-01-02T15:04:05.999Z07:00 格式
func FixedZoneTime(s string) (string, error) {
	return FixedZoneTimeFromTimeAndLayout(s, Rfc3399Mills)
}

// GetFixedZoneTimeString 获取固定东八区的时间
func GetFixedZoneTimeString(t time.Time) string {
	// 默认东八区
	zone := time.FixedZone("UTC+8", 8*60*60)
	s := t.In(zone).Format(Rfc3399Mills)
	return s
}

// DurationOfTime 计算两个日期字符串的间隔时间
func DurationOfTime(start string, end string) (d time.Duration, err error) {
	u, err := ParseTime(start)
	if err != nil {
		return
	}
	c, err := ParseTime(end)
	if err != nil {
		return
	}
	d = c.Sub(u)
	return
}
