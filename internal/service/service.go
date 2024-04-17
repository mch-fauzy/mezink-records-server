package service

import (
	"github.com/mezink-records-server/internal/model"
	"github.com/mezink-records-server/internal/model/dto"
	"github.com/mezink-records-server/internal/repository"
	"github.com/rs/zerolog/log"
)

type ServiceImpl struct {
	Repository repository.Repository
}

func ProvideService(repository repository.Repository) *ServiceImpl {
	return &ServiceImpl{
		Repository: repository,
	}
}

type Service interface {
	GetRecordList(req dto.Request) (dto.Response, error)
}

func (s *ServiceImpl) GetRecordList(req dto.Request) (dto.Response, error) {
	code := 1
	msg := "Failed"

	dateFilterValue := req.ToInterface()
	recordList, err := s.Repository.ResolveRecordList(model.Filter{
		FilterFields: []model.FilterField{
			{
				Field:    model.RecordDBFieldName.CreatedAt,
				Operator: model.OperatorRange,
				Value:    dateFilterValue,
			},
		},
	})

	if err != nil {
		log.Error().Err(err).Msg("[GetRecordList] Failed to retrieve record list")
		return dto.Response{}, err
	}

	code = 0
	msg = "Success"
	result := dto.BuildRecordListResponse(code, msg, recordList, req)

	return result, nil
}
