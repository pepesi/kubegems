package model

type Quota struct {
	ID      uint
	RelKind string
	RelName string
	RelID   uint
	Datas   map[string]interface{}
}

type QuotaData struct{}
