package database

import (
	"context"
	"strings"

	"github.com/adminvoras/commons-lib/log"
	"github.com/jmoiron/sqlx"
)

const noRowsMessage = "no rows in result set"

func IsNoRowsError(err error) bool {
	return strings.Contains(err.Error(), noRowsMessage)
}

func FinishTransaction(ctx context.Context, tx *sqlx.Tx, err error) {
	logger := log.DefaultLogger()

	if err != nil {
		if err = tx.Rollback(); err != nil {
			logger.Error(nil, nil, err, "Error rollbacking database transaction changes")
		}

		return
	}

	if err = tx.Commit(); err != nil {
		logger.Error(nil, nil, err, "Error committing database transaction changes")

		FinishTransaction(ctx, tx, err)
	}
}
