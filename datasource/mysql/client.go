package mysql

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/micro/go-micro/util/log"
)

// safe guard
const selectLimit = 10000

func (h *Handler) ExecRawSql(sql string, values ...interface{}) (int64, error) {
	log.Infof("exec raw sql %s with values %v\n", sql, values)
	db, err := h.GetDb()
	if err != nil {
		return 0, err
	}
	if db == nil {
		return 0, errors.New("nil *gorm.DB")
	}

	dbExec := db.Exec(sql, values)
	return dbExec.RowsAffected, dbExec.Error
}

func (h *Handler) SelectRows(sql string, values ...interface{}) (*sql.Rows, error) {
	log.Infof("exec raw sql %s with values %v\n", sql, values)
	db, err := h.GetDb()
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, errors.New("nil *gorm.DB")
	}

	dbRun := db
	if !strings.Contains(strings.ToUpper(sql), " LIMIT ") {
		dbRun = db.Limit(selectLimit)
	}
	return dbRun.Raw(sql, values).Rows()
}
