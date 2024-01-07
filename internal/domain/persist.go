package domain

import "time"

type Persist interface {
	Store(key string, time time.Time, data interface{}) error
	Info(key string) []Row
}

type Row struct {
	Time time.Time   `json:"time"`
	Data interface{} `json:"data"`
}
