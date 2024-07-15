package date_test

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/adminvoras/commons-lib/pkg/date"
	"github.com/stretchr/testify/assert"
)

func TestNumTZDate_UnmarshalJSON(t *testing.T) {
	type testJSON struct {
		Date date.NumTZDate `json:"date"`
	}

	tests := []struct {
		name      string
		json      string
		wantedErr error
		wanted    string
	}{
		{
			name:      "Should unmarshal JSON successfully when the date is not UTC (-3)",
			json:      `{"date" : "2020-01-21T13:24:32.123-0300"}`,
			wantedErr: nil,
			wanted:    "2020-01-21 16:24:32.123 +0000 UTC",
		},
		{
			name:      "Should unmarshal JSON successfully when the date is not UTC (+2)",
			json:      `{"date" : "2020-01-21T13:24:32.123+0200"}`,
			wantedErr: nil,
			wanted:    "2020-01-21 11:24:32.123 +0000 UTC",
		},
		{
			name:      "Should unmarshal JSON successfully when the date is already in UTC",
			json:      `{"date" : "2020-08-01T18:29:08.000+0000"}`,
			wantedErr: nil,
			wanted:    "2020-08-01 18:29:08 +0000 UTC",
		},
		{
			name: "Should return ParseError when format is not the correct",
			json: `{"date" : "2020-01-21T13:24:32"}`,
			wantedErr: &time.ParseError{
				Layout:     "2006-01-02T15:04:05.999-0700",
				Value:      "2020-01-21T13:24:32",
				LayoutElem: "-0700",
				ValueElem:  "",
				Message:    "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &testJSON{}
			err := json.Unmarshal(bytes.NewBufferString(tt.json).Bytes(), res)

			if tt.wantedErr != nil {
				assert.Equal(t, tt.wantedErr, err, "Error is not the expected")

				return
			}

			assert.Nil(t, err, "Error should be nil")
			assert.Equal(t, tt.wanted, res.Date.String(), "Date is not the expected")
		})
	}
}

func TestNumTZDate_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		date      date.NumTZDate
		wantedErr error
		wanted    string
	}{
		{
			name: "Should marshal date successfully when the date is UTC-8",
			date: date.NumTZDate{
				Time: time.Date(2020, 8, 5, 17, 5, 58, 497*int(time.Millisecond), time.FixedZone("UTC-8", -8*60*60)),
			},
			wantedErr: nil,
			wanted:    `"2020-08-05T17:05:58.497-0800"`,
		},
		{
			name: "Should marshal date successfully when the date is UTC-3",
			date: date.NumTZDate{
				Time: time.Date(2020, 9, 15, 18, 5, 58, 497*int(time.Millisecond), time.FixedZone("UTC-3", -3*60*60)),
			},
			wantedErr: nil,
			wanted:    `"2020-09-15T18:05:58.497-0300"`,
		},
		{
			name: "Should marshal date successfully when the date is UTC+2",
			date: date.NumTZDate{
				Time: time.Date(2020, 10, 30, 19, 5, 58, 497*int(time.Millisecond), time.FixedZone("UTC+2", 2*60*60)),
			},
			wantedErr: nil,
			wanted:    `"2020-10-30T19:05:58.497+0200"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, err := tt.date.MarshalJSON()

			if tt.wantedErr != nil {
				assert.Equal(t, tt.wantedErr, err, "Error is not the expected")

				return
			}

			assert.Nil(t, err, "Error should be nil")
			assert.Equal(t, tt.wanted, string(bytes), "Date is not the expected")
		})
	}
}
