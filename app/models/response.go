package models

type Error struct {
	TraceId string `json:"trace_id"`
	Message string `json:"message"`
}
