package utctime

import (
	"fmt"
	"strings"
	"time"
)

const dateformat = "2006-01-02"

type UTCDate struct {
	time.Time
}

func (utc UTCDate) String() string {
	return utc.Format(dateformat)
}

func (utc UTCDate) Stringf(format string) string {
	return utc.Format(format)
}

func ParseToUTCDate(timeString string) (utc UTCDate, err error) {
	var parsed time.Time
	for _, layout := range layouts {
		parsed, err = time.Parse(layout, timeString)
		if err == nil {
			utc.Time = parsed.UTC().Truncate(time.Hour * 24)
			return
		}
	}
	return utc, fmt.Errorf("invalid date format: %s", timeString)
}

func (utc *UTCDate) UnmarshalJSON(data []byte) error {
	timeString := strings.Trim(string(data), `"`)
	parsed, err := ParseToUTCDate(timeString)
	if err != nil {
		return err
	}

	*utc = parsed
	return nil
}

func (utc UTCDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, utc.Stringf(dateformat))), nil
}
func (utc UTCDate) BeforeOrEqual(otherDate UTCDate) bool {
	return utc.Before(otherDate.Time) || utc.Equal(otherDate.Time)
}

func (utc UTCDate) AfterOrEqual(otherDate UTCDate) bool {
	return utc.After(otherDate.Time) || utc.Equal(otherDate.Time)
}
