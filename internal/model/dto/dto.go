package dto

import (
	"encoding/json"
	"regexp"
	"time"

	"github.com/mezink-records-server/internal/model"
	"github.com/mezink-records-server/shared/failure"
)

const (
	notDigitPattern = "^[0-9]+$"
)

type Request struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

func (req Request) ToInterface() []interface{} {
	return []interface{}{
		req.StartDate,
		req.EndDate,
	}
}

func (req Request) Validate() error {
	// check if date is empty or not in correct format
	dateRegex := regexp.MustCompile(`^(?:\d{4}-\d{2}-\d{2})?$`)

	if !dateRegex.MatchString(req.StartDate) {
		return failure.BadRequestFromString("Start Date must be in YYYY-MM-DD format")
	}

	if !dateRegex.MatchString(req.EndDate) {
		return failure.BadRequestFromString("End Date must be in YYYY-MM-DD format")
	}

	if req.MinCount < 1 {
		return failure.BadRequestFromString("Min Count must be at least 1")
	}

	if req.MaxCount < 1 {
		return failure.BadRequestFromString("Max Count must be at least 1")
	}

	return nil
}

type RecordResponse struct {
	Id         int64     `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalMarks int       `json:"totalMarks"`
}

type Response struct {
	Code    int              `json:"code"`
	Message string           `json:"msg"`
	Records []RecordResponse `json:"records"`
}

func BuildRecordListResponse(code int, msg string, recordList model.RecordList, request Request) Response {
	return Response{
		Code:    code,
		Message: msg,
		Records: BuildRecordList(recordList, request),
	}
}

func NewStudentEnrollmentResponse(record model.Record) RecordResponse {
	totalMarks := calculateTotalMarks(record.Marks)
	// log.Printf("totalMarks: %v", totalMarks)
	return RecordResponse{
		Id:         record.Id,
		CreatedAt:  record.CreatedAt,
		TotalMarks: totalMarks,
	}
}

func BuildRecordList(recordList model.RecordList, request Request) []RecordResponse {
	results := []RecordResponse{}
	for _, record := range recordList {
		totalMarks := calculateTotalMarks(record.Marks)

		if totalMarks >= request.MinCount && totalMarks <= request.MaxCount {
			results = append(results, NewStudentEnrollmentResponse(*record))
		}
	}
	return results
}

func calculateTotalMarks(marks []uint8) int {
	total := 0

	var marksArray []int
	json.Unmarshal([]uint8(marks), &marksArray)
	for _, mark := range marksArray {
		total += mark
	}

	return total
}
