package log_test

import (
	"sort"
	"testing"

	"github.com/adminvoras/commons-lib/pkg/log"
	"github.com/stretchr/testify/assert"
)

func Test_log_getMessage(t *testing.T) {
	type fields struct {
		requestID string
	}

	type args struct {
		message string
		args    []interface{}
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "Without args",
			fields: fields{requestID: "123"},
			args: args{
				message: "the log message",
				args:    nil,
			},
			want: "the log message",
		},
		{
			name:   "with one string argument",
			fields: fields{requestID: "123"},
			args: args{
				message: "with one string %v message",
				args:    []interface{}{"argument"},
			},
			want: "with one string argument message",
		},
		{
			name:   "with one int argument",
			fields: fields{requestID: "123"},
			args: args{
				message: "with %v int argument message",
				args:    []interface{}{1},
			},
			want: "with 1 int argument message",
		},
		{
			name:   "with one string argument and one int",
			fields: fields{requestID: "123"},
			args: args{
				message: "with one string %s and %d int value message",
				args:    []interface{}{"argument", 1},
			},
			want: "with one string argument and 1 int value message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			theLogger := log.NewLogger(tt.fields.requestID)
			got := theLogger.GetMessage(tt.args.message, tt.args.args...)

			assert.Equal(t, tt.want, got, "Logger is not the expected")
		})
	}
}

func TestDefaultLogger(t *testing.T) {
	logger := log.DefaultLogger()

	assert.True(t, len(logger.GetRequestID()) > 0, "Request ID cannot be empty")
}

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name      string
		requestID string
		wantedID  string
	}{
		{
			name:      "New Logger with request ID",
			requestID: "the-request-id",
			wantedID:  "the-request-id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := log.NewLogger(tt.requestID)

			assert.Equal(t, tt.wantedID, got.GetRequestID(), "Unexpected request id")
		})
	}
}

func Test_log_GetTags(t *testing.T) {
	type sourceTest struct{}
	type args struct {
		source interface{}
		tags   map[string]string
	}
	tests := []struct {
		name          string
		args          args
		previousCalls int
		want          []string
	}{
		{
			name: "Should return tags with sequence id 1",
			args: args{
				source: sourceTest{},
				tags:   nil,
			},
			previousCalls: 0,
			want: []string{
				"Request_ID:requestID",
				"Class:log_test.sourceTest",
				"Sequence_ID:1",
			},
		},
		{
			name: "Should work correctly when tags parameter is empty",
			args: args{
				source: sourceTest{},
				tags:   map[string]string{},
			},
			previousCalls: 19,
			want: []string{
				"Request_ID:requestID",
				"Class:log_test.sourceTest",
				"Sequence_ID:20",
			},
		},
		{
			name: "Should return 4 tags when tags parameter has 1 element",
			args: args{
				source: sourceTest{},
				tags: map[string]string{
					"key1": "value1",
				},
			},
			previousCalls: 3,
			want: []string{
				"key1:value1",
				"Request_ID:requestID",
				"Class:log_test.sourceTest",
				"Sequence_ID:4",
			},
		},
		{
			name: "Should return 5 tags when tags parameter has 2 element",
			args: args{
				source: sourceTest{},
				tags: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
			previousCalls: 0,
			want: []string{
				"key1:value1",
				"key2:value2",
				"Request_ID:requestID",
				"Class:log_test.sourceTest",
				"Sequence_ID:1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			theLogger := log.NewLogger("requestID")

			for i := 0; i < tt.previousCalls; i++ {
				theLogger.GetTags(tt.args.source, tt.args.tags)
			}

			got := theLogger.GetTags(tt.args.source, tt.args.tags)

			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i] > tt.want[j]
			})

			sort.Slice(got, func(i, j int) bool {
				return got[i] > got[j]
			})

			assert.Equal(t, tt.want, got, "Unexpected tag")
		})
	}
}
