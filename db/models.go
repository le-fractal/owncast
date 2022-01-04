// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type ApAcceptedActivity struct {
	ID        int32
	Iri       string
	Actor     string
	Type      string
	Timestamp time.Time
}

type ApFollower struct {
	Iri        string
	Inbox      string
	Name       sql.NullString
	Username   string
	Image      sql.NullString
	Request    string
	CreatedAt  sql.NullTime
	ApprovedAt sql.NullTime
	DisabledAt sql.NullTime
}

type ApOutbox struct {
	Iri              string
	Value            []byte
	Type             string
	CreatedAt        sql.NullTime
	LiveNotification sql.NullBool
}
