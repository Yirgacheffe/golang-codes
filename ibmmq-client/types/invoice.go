package types

import "time"

type InoviceEvent struct {
	Id        int
	EventId   string
	ReqId     string
	EventType string
	EventData string
	CreatedOn time.Time
}

type Invoice struct {
	Id          int
	ReqId       string
	SerialNbr   string
	InvoiceCode string
	InvoiceTime string
	Status      string
	CreatedOn   time.Time
	UpdatedOn   time.Time
}
