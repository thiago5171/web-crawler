package postgres

import (
	"backend_template/src/core/domain/errors"
	"backend_template/src/infra"
	"backend_template/src/infra/repository"
	"strings"
)

var logger = infra.Logger().With().Str("port", "postgres").Logger()

func defaultExecQuery(sqlQuery string, args ...interface{}) errors.Error {
	result, err := repository.ExecQuery(sqlQuery, args...)
	if err != nil {
		return err
	} else if rowsAff, err := result.RowsAffected(); err != nil {
		return errors.NewInternal(err)
	} else if rowsAff == 0 {
		if strings.Contains(strings.ToLower(sqlQuery), "update") {
			return errors.NewFromString("no entries were found to be updated")
		} else if strings.Contains(strings.ToLower(sqlQuery), "delete") {
			return errors.NewFromString("no entries were found to be deleted")
		}
		return errors.NewUnexpected()
	}
	return nil
}

func defaultTxExecQuery(tx *repository.SQLTransaction, sqlQuery string, args ...interface{}) errors.Error {
	result, err := tx.ExecQuery(sqlQuery, args...)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	} else if rowsAff, err := result.RowsAffected(); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return errors.NewInternal(err)
	} else if rowsAff == 0 {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return errors.NewUnexpected()
	}
	return nil
}

func txQueryRowReturningID(tx *repository.SQLTransaction, sqlQuery string, args ...interface{}) (string, errors.Error) {
	row := tx.QueryRow(sqlQuery, args...)
	if err := row.Err(); err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return "", rollBackErr
		}
		return "", repository.TranslateError(row.Err())
	}
	var strUUID = ""
	scanErr := row.Scan(&strUUID)
	if scanErr != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return "", rollBackErr
		}
		return "", repository.TranslateError(scanErr)
	}
	return strUUID, nil
}
