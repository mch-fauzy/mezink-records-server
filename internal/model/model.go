package model

import "time"

type Record struct {
	Id        int64     `db:"id"`
	Marks     []uint8   `db:"marks"`
	CreatedAt time.Time `db:"created_at"`
}

type RecordList []*Record

type recordDBFieldName struct {
	CreatedAt string
	Marks     string
}

var RecordDBFieldName = recordDBFieldName{
	CreatedAt: "created_at",
	Marks:     "SUM(marks)",
}
