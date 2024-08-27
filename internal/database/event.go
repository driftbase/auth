package database

import "time"

type Event struct {
	Id                string    `db:"id" json:"id"`
	TenantId          string    `db:"tenant_id" json:"tenant_id"`
	AggregateType     string    `db:"aggregate_type" json:"aggregate_type"`
	AggregateVersion  int       `db:"aggregate_version" json:"aggregate_version"`
	AggregateId       string    `db:"aggregate_id" json:"aggregate_id"`
	AggregateSequence int64     `db:"aggregate_sequence" json:"aggregate_sequence"`
	EventType         string    `db:"event_type" json:"event_type"`
	EventVersion      int       `db:"event_version" json:"event_version"`
	Data              []byte    `db:"data" json:"data"`
	Creator           string    `db:"creator" json:"creator"`
	Owner             string    `db:"owner" json:"owner"`
	CorrelationId     *string   `db:"correlation_id" json:"correlation_id"`
	CausationId       *string   `db:"causation_id" json:"causation_id"`
	Position          float64   `db:"position" json:"position"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
}