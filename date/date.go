package date

import (
	"encoding/json"
	"strings"
	"time"
)

// nolint: unused, deadcode, varcheck // copied from golang time package
const (
	stdISO8601TZ             = "2006-01-02T15:04:05.999-Z0700"     // "Z0700"  // prints Z for UTC
	stdISO8601SecondsTZ      = "2006-01-02T15:04:05.999-Z070000"   // "Z070000"
	stdISO8601ShortTZ        = "2006-01-02T15:04:05.999-Z07"       // "Z07"
	stdISO8601ColonTZ        = "2006-01-02T15:04:05.999-Z07:00"    // "Z07:00" // prints Z for UTC
	stdISO8601ColonSecondsTZ = "2006-01-02T15:04:05.999-Z07:00:00" // "Z07:00:00"
	stdNumTZ                 = "2006-01-02T15:04:05.999-0700"      // "-0700"  // always numeric
	stdNumSecondsTz          = "2006-01-02T15:04:05.999-070000"    // "-070000"
	stdNumShortTZ            = "2006-01-02T15:04:05.999-07"        // "-07"    // always numeric
	stdNumColonTZ            = "2006-01-02T15:04:05.999-07:00"     // "-07:00" // always numeric
	stdNumColonSecondsTZ     = "2006-01-02T15:04:05.999-07:00:00"  // "-07:00:00"
)

type JsonDate interface {
	UnmarshalJSON(data []byte) error
	MarshalJSON() ([]byte, error)
	setTime(time time.Time)
	getTime() time.Time
}

func UnmarshalJSON(date JsonDate, layout string, data []byte) error {
	stringDate := strings.Trim(strings.Trim(string(data), `"`), `\"`)

	parsedDate, err := time.Parse(layout, stringDate)
	if err != nil {
		return err
	}

	date.setTime(parsedDate.UTC())

	return nil
}

func MarshalJSON(date JsonDate, layout string) ([]byte, error) {
	dateString := ""

	if date != nil && !date.getTime().IsZero() {
		dateString = date.getTime().Format(layout)
	}

	return json.Marshal(dateString)
}

// NumTZDate is the date struct for '2006-01-02T15:04:05.999-0700' format.
type NumTZDate struct {
	time.Time
}

func (date *NumTZDate) UnmarshalJSON(data []byte) error {
	return UnmarshalJSON(date, stdNumTZ, data)
}

func (date *NumTZDate) MarshalJSON() ([]byte, error) {
	return MarshalJSON(date, stdNumTZ)
}

func (date *NumTZDate) setTime(time time.Time) {
	date.Time = time
}

func (date *NumTZDate) getTime() time.Time {
	return date.Time
}
