package repository

import (
	"fmt"
	"strings"

	"github.com/mezink-records-server/infras"
	"github.com/mezink-records-server/internal/model"
	"github.com/mezink-records-server/shared/failure"
	"github.com/rs/zerolog/log"
)

const (
	selectRecordQuery = `
	SELECT
		id,
		marks,
		created_at
	FROM
		records
	`
)

type RepositoryMySql struct {
	DB *infras.MySqlConn
}

func ProvideRepositoryMySql(db *infras.MySqlConn) *RepositoryMySql {
	return &RepositoryMySql{
		DB: db,
	}
}

type Repository interface {
	ResolveRecordList(filter model.Filter) (model.RecordList, error)
}

func (r *RepositoryMySql) ResolveRecordList(filter model.Filter) (model.RecordList, error) {
	query, args, err := composeFilterQuery(selectRecordQuery, filter)
	if err != nil {
		return model.RecordList{}, err
	}

	var recordList model.RecordList
	err = r.DB.Read.Select(&recordList, query, args...)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[ResolveRecordList] Failed to get record")
		err = failure.InternalError(err)
		return model.RecordList{}, err
	}

	return recordList, nil
}

func composeFilterQuery(queryStr string, filter model.Filter) (string, []interface{}, error) {
	var (
		err  error
		args []interface{}
	)

	query := fmt.Sprintf(queryStr)

	if len(filter.FilterFields) > 0 {
		var (
			filterQueries []string
			filterArgs    []interface{}
		)

		for _, filterField := range filter.FilterFields {
			switch filterField.Operator {
			case model.OperatorRange:
				// Input must be []interface{}
				valueArray, ok := filterField.Value.([]interface{})
				if !ok && len(valueArray) != 2 {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return query, args, err
				}
				filterQueries = append(filterQueries, fmt.Sprintf("`%s` BETWEEN ? AND ?", filterField.Field))
				filterArgs = append(filterArgs, valueArray...)
			}
		}

		query += fmt.Sprintf(" WHERE %s", strings.Join(filterQueries, " AND "))
		args = append(args, filterArgs...)
	}

	return query, args, err
}
