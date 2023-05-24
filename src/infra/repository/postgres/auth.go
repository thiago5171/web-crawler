package postgres

import (
	"backend_template/src/core/domain/account"
	"backend_template/src/core/domain/credentials"
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/role"
	"backend_template/src/core/interfaces/adapters"
	"backend_template/src/core/utils"
	"backend_template/src/infra/repository"
	"backend_template/src/infra/repository/postgres/query"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type authPostgresRepository struct{}

func NewAuthPostgresRepository() adapters.AuthAdapter {
	return &authPostgresRepository{}
}

func (r *authPostgresRepository) Login(credentials credentials.Credentials) (account.Account, errors.Error) {
	rows, err := repository.Queryx(query.Account().Select().ByCredentials(), credentials.Email())
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, errors.NewFromString("email and/or password are incorrect")
	}
	var id, password string
	scanErr := rows.Scan(&id, &password)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, errors.NewFromString("email and/or password are incorrect")
		}
		logger.Error().Msg(scanErr.Error())
		return nil, errors.NewUnexpected()
	}
	if err := comparePasswords(password, credentials.Password()); err != nil {
		return nil, err
	}
	account, err := r.getAccountByCustomQuery(query.Account().Select().ByID(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	return account, nil
}

func (r *authPostgresRepository) getAccountByCustomQuery(query string, args ...interface{}) (account.Account, errors.Error) {
	rows, queryErr := repository.Queryx(query, args...)
	if queryErr != nil {
		return nil, queryErr
	}
	if !rows.Next() {
		return nil, errors.NewFromString("account not found")
	}
	var serializedAccount = map[string]interface{}{}
	scanErr := rows.MapScan(serializedAccount)
	if scanErr != nil {
		return nil, errors.NewUnexpected()
	}
	account, convErr := newAccountFromMapRows(serializedAccount)
	if convErr != nil {
		return nil, convErr
	}
	return account, nil
}

func comparePasswords(current, confirmation string) errors.Error {
	if !utils.PasswordIsValid(current, confirmation) {
		return errors.NewFromString("invalid password")
	}
	return nil
}

func newRoleFromMapRows(data map[string]interface{}) (role.Role, errors.Error) {
	var err error
	var id uuid.UUID
	var name = fmt.Sprint(data["name"])
	var code = fmt.Sprint(data["code"])
	id, err = uuid.Parse(string(data["id"].([]uint8)))
	if err != nil {
		return nil, errors.NewUnexpected()
	}
	return role.New(&id, name, code)
}
