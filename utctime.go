package utctime

import (
	"fmt"
	"strings"
	"time"
)

const defaultFormat = time.RFC3339
const MilisecondsFormat = "2006-01-02T15:04:05.000Z"
const DateTimeZeroFormat = "2006-01-02T15:04:05Z"

var layouts = []string{
	defaultFormat,
	"2006-01-02T15:04Z",       // ISO 8601 UTC
	DateTimeZeroFormat,        // ISO 8601 UTC
	MilisecondsFormat,         // ISO 8601 UTC
	"2006-01-02T15:04:05",     // ISO 8601 UTC
	"2006-01-02 15:04",        // Custom UTC
	"2006-01-02 15:04:05",     // Custom UTC
	"2006-01-02 15:04:05.000", // Custom UTC
	"2006-01-02",
}

type UTCTime struct {
	time.Time
}

func Now() UTCTime {
	return UTCTime{
		Time: time.Now(),
	}
}

func (utc UTCTime) String() string {
	return utc.Format(defaultFormat)
}

func (utc UTCTime) Stringf(format string) string {
	return utc.Format(format)
}

func (utc *UTCTime) UnmarshalJSON(data []byte) error {
	timeString := strings.Trim(string(data), `"`)
	parsed, err := ParseToUTCTime(timeString)
	if err != nil {
		return err
	}

	*utc = parsed
	return nil
}

func (utc UTCTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, utc.Stringf(DateTimeZeroFormat))), nil
}

func ParseToUTCTime(timeString string) (utc UTCTime, err error) {
	var parsed time.Time
	for _, layout := range layouts {
		parsed, err = time.Parse(layout, timeString)
		if err == nil {
			utc.Time = parsed.UTC()
			return
		}
	}
	return utc, fmt.Errorf("invalid date format: %s", timeString)
}

func ParseToUTCInLocation(layout string, timeString string, location string) (utc UTCTime, err error) {
	var parsed time.Time
	loc, _ := time.LoadLocation(location)
	parsed, err = time.ParseInLocation(layout, timeString, loc)
	if err == nil {
		utc.Time = parsed.UTC()
		return
	}
	return utc, fmt.Errorf("invalid date format: %s", timeString)
}

func (utc UTCTime) StringMiliseconds() string {
	return utc.Format(MilisecondsFormat)
}

func (utc UTCTime) ToSPTimeZone() UTCTime {
	loc, _ := time.LoadLocation("America/Sao_Paulo")

	return UTCTime{
		Time: utc.Time.In(loc),
	}
}
