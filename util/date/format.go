package date

import "time"

const (
	DATE_TIME_FORMAT = Format("2006-01-02 15:04:05")
	DATE_FORMAT = Format("2006-01-02")
)

type Format string

func (format Format) Format() string{
	return time.Now().Format(string(format))
}
