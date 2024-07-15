package database_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/adminvoras/commons-lib/pkg/database"
	voraserrors "github.com/adminvoras/commons-lib/pkg/errors"
)

func Test_clientBuilder_Build(t *testing.T) {
	type fields struct {
		driverName      string
		charset         string
		host            string
		dbName          string
		username        string
		password        string
		maxIdleConns    int
		maxOpenConns    int
		initialPing     bool
		connMaxLifetime time.Duration
	}

	tests := []struct {
		name      string
		fields    fields
		wantedErr error
	}{
		{
			name: "Database client is successfully created",
			fields: fields{
				driverName:      "mysql",
				charset:         "utf8",
				host:            "anyhost",
				dbName:          "dbname",
				username:        "username",
				password:        "password",
				maxIdleConns:    1,
				maxOpenConns:    1,
				connMaxLifetime: 1 * time.Millisecond,
				initialPing:     false,
			},
			wantedErr: nil,
		},
		{
			name: "Database client is not created when the driver is not imported",
			fields: fields{
				driverName:      "sqlserver",
				charset:         "utf8",
				host:            "anyhost",
				dbName:          "dbname",
				username:        "username",
				password:        "password",
				maxIdleConns:    1,
				maxOpenConns:    1,
				connMaxLifetime: 1 * time.Millisecond,
				initialPing:     false,
			},
			wantedErr: voraserrors.New(errors.New(`sql: unknown driver "sqlserver" (forgotten import?)`), "error connecting to sqlserver database"),
		},
		{
			name: "Database client is not created when the database host is empty",
			fields: fields{
				driverName:      "sqlserver",
				charset:         "utf8",
				host:            "",
				dbName:          "dbname",
				username:        "username",
				password:        "password",
				maxIdleConns:    1,
				maxOpenConns:    1,
				connMaxLifetime: 1 * time.Millisecond,
				initialPing:     false,
			},
			wantedErr: voraserrors.New(nil, "database host cannot be empty"),
		},
		{
			name: "Database client is not created when the database name is empty",
			fields: fields{
				driverName:      "sqlserver",
				charset:         "utf8",
				host:            "host",
				dbName:          "",
				username:        "username",
				password:        "password",
				maxIdleConns:    1,
				maxOpenConns:    1,
				connMaxLifetime: 1 * time.Millisecond,
				initialPing:     false,
			},
			wantedErr: voraserrors.New(nil, "database name cannot be empty"),
		},
		{
			name: "Database client is not created when the database username is empty",
			fields: fields{
				driverName:      "sqlserver",
				charset:         "utf8",
				host:            "host",
				dbName:          "dbname",
				username:        "",
				password:        "password",
				maxIdleConns:    1,
				maxOpenConns:    1,
				connMaxLifetime: 1 * time.Millisecond,
				initialPing:     false,
			},
			wantedErr: voraserrors.New(nil, "database username cannot be empty"),
		},
		{
			name: "Database client is not created when the database password is empty",
			fields: fields{
				driverName:      "sqlserver",
				charset:         "utf8",
				host:            "host",
				dbName:          "dbname",
				username:        "username",
				password:        "",
				maxIdleConns:    1,
				maxOpenConns:    1,
				connMaxLifetime: 1 * time.Millisecond,
				initialPing:     false,
			},
			wantedErr: voraserrors.New(nil, "database password cannot be empty"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := database.NewClientBuilder().
				WithDriverName(tt.fields.driverName).
				WithCharset(tt.fields.charset).
				WithHost(tt.fields.host).
				WithDBName(tt.fields.dbName).
				WithUsername(tt.fields.username).
				WithPassword(tt.fields.password).
				WithMaxIdleConns(tt.fields.maxIdleConns).
				WithMaxOpenConns(tt.fields.maxOpenConns).
				WithConnMaxLifetime(tt.fields.connMaxLifetime).
				WithInitialPing(tt.fields.initialPing)

			got, err := builder.Build()

			if tt.wantedErr != nil {
				assert.Equal(t, tt.wantedErr, err, "Error is not the expected building database client")

				return
			}

			assert.Nil(t, err, "Unexpected error building database client")
			assert.NotNil(t, got, "Database client should be not nil")
		})
	}
}
