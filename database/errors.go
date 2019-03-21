package database

import "github.com/lib/pq"

const (
	notNullViolationCode = pq.ErrorCode("23502")
	uniqueViolationCode  = pq.ErrorCode("23505")
)
