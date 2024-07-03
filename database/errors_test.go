package database_test

//
//import (
//	"context"
//	"errors"
//	"testing"
//
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/adminvoras/commons-lib/database"
//
//	"github.com/jmoiron/sqlx"
//)
//
//func TestIsNoRowsError(t *testing.T) {
//	type args struct {
//		err error
//	}
//
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		{
//			name: "No rows error is successfully detected",
//			args: args{
//				err: errors.New("sql: no rows in result set"),
//			},
//			want: true,
//		},
//		{
//			name: "No rows error is not detected when the error is different",
//			args: args{
//				err: errors.New("some error"),
//			},
//			want: false,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got := database.IsNoRowsError(tt.args.err)
//
//			assert.Equals(t, "Error is not the expected", tt.want, got)
//		})
//	}
//}
//
//func TestFinishTransaction(t *testing.T) {
//
//	type args struct {
//		ctx context.Context
//		err error
//	}
//
//	tests := []struct {
//		name         string
//		args         args
//		enableCommit bool
//	}{
//		{
//			name: "Database transaction successfully finished",
//			args: args{
//				ctx: reconcontext.Background(),
//			},
//			enableCommit: true,
//		},
//		{
//			name: "Database transaction cannot be finished when the commit returns an error",
//			args: args{
//				ctx: reconcontext.Background(),
//			},
//			enableCommit: false,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			db, mock, err := sqlmock.New()
//			assert.Nil(t, "Unexpected error creating mock database", err)
//
//			mock.ExpectBegin()
//
//			if tt.enableCommit {
//				mock.ExpectCommit()
//				mock.ExpectRollback()
//			}
//
//			sqlTx, err := db.BeginTx(tt.args.ctx, nil)
//			assert.Nil(t, "Unexpected error creating mock transaction", err)
//
//			tx := &sqlx.Tx{
//				Tx: sqlTx,
//			}
//
//			database.FinishTransaction(tt.args.ctx, tx, tt.args.err)
//		})
//	}
//}
