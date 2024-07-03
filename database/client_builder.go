package database

import (
	"fmt"
	"time"

	// MySQL driver.
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	voraserror "github.com/adminvoras/commons-lib/errors"
)

const (
	defaultDriverName      = "mysql"
	defaultCharset         = "utf8"
	defaultMaxOpenConns    = 350
	defaultMaxIdleConns    = 100
	defaultConnMaxLifetime = 100 * time.Millisecond
	datasourceConnection   = "%s:%s@tcp(%s)/%s?parseTime=%v&charset=%s"
)

// ClientBuilder the database client builder interface.
type ClientBuilder interface {
	WithDriverName(name string) ClientBuilder
	WithCharset(charset string) ClientBuilder
	WithHost(host string) ClientBuilder
	WithDBName(name string) ClientBuilder
	WithUsername(username string) ClientBuilder
	WithPassword(password string) ClientBuilder
	WithMaxIdleConns(maxIdleConns int) ClientBuilder
	WithMaxOpenConns(maxOpenConns int) ClientBuilder
	WithConnMaxLifetime(connMaxLifetime time.Duration) ClientBuilder
	WithInitialPing(initialPing bool) ClientBuilder
	Build() (Client, error)
}

// clientBuilder the database client builder.
type clientBuilder struct {
	driverName      string
	charset         string
	host            string
	dbName          string
	username        string
	password        string
	maxIdleConns    int
	maxOpenConns    int
	connMaxLifetime time.Duration
	initialPing     bool
}

// NewClientBuilder creates a new database client builder with default settings.
func NewClientBuilder() ClientBuilder {
	builder := &clientBuilder{
		driverName:      defaultDriverName,
		charset:         defaultCharset,
		maxIdleConns:    defaultMaxIdleConns,
		maxOpenConns:    defaultMaxOpenConns,
		connMaxLifetime: defaultConnMaxLifetime,
		initialPing:     true,
	}

	return builder
}

func (builder *clientBuilder) WithDriverName(name string) ClientBuilder {
	builder.driverName = name

	return builder
}

func (builder *clientBuilder) WithCharset(charset string) ClientBuilder {
	builder.charset = charset

	return builder
}

func (builder *clientBuilder) WithHost(host string) ClientBuilder {
	builder.host = host

	return builder
}

func (builder *clientBuilder) WithDBName(name string) ClientBuilder {
	builder.dbName = name

	return builder
}

func (builder *clientBuilder) WithUsername(username string) ClientBuilder {
	builder.username = username

	return builder
}

func (builder *clientBuilder) WithPassword(password string) ClientBuilder {
	builder.password = password

	return builder
}

func (builder *clientBuilder) WithMaxIdleConns(maxIdleConns int) ClientBuilder {
	builder.maxIdleConns = maxIdleConns

	return builder
}

func (builder *clientBuilder) WithMaxOpenConns(maxOpenConns int) ClientBuilder {
	builder.maxOpenConns = maxOpenConns

	return builder
}

func (builder *clientBuilder) WithConnMaxLifetime(connMaxLifetime time.Duration) ClientBuilder {
	builder.connMaxLifetime = connMaxLifetime

	return builder
}

func (builder *clientBuilder) WithInitialPing(initialPing bool) ClientBuilder {
	builder.initialPing = initialPing

	return builder
}

func (builder *clientBuilder) Build() (Client, error) {
	if builder.host == "" {
		return nil, voraserror.New(nil, "database host cannot be empty")
	}

	if builder.dbName == "" {
		return nil, voraserror.New(nil, "database name cannot be empty")
	}

	if builder.username == "" {
		return nil, voraserror.New(nil, "database username cannot be empty")
	}

	if builder.password == "" {
		return nil, voraserror.New(nil, "database password cannot be empty")
	}

	db, err := sqlx.Open(builder.driverName, fmt.Sprintf(datasourceConnection, builder.username, builder.password,
		builder.host, builder.dbName, true, builder.charset))

	errorMessage := fmt.Sprintf("error connecting to %v database", builder.driverName)

	if err != nil {
		return nil, voraserror.New(err, errorMessage)
	}

	db.SetMaxIdleConns(builder.maxIdleConns)
	db.SetMaxOpenConns(builder.maxOpenConns)
	db.SetConnMaxLifetime(builder.connMaxLifetime)

	if builder.initialPing {
		if err = db.Ping(); err != nil {
			return nil, voraserror.New(err, "%s: ping has failed")
		}
	}

	return db, err
}
