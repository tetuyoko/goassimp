package models

type ConvertLog struct {
	Template
	UUID string `sql:"not null;unique_index:idx_name_uuid"`
	Path string `sql:"not null;`
}
