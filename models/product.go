package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"type:varchar(300)" json:"nama_product"`
	Deskripsi   string `gorm:"type:text" json:"deskripsi"`
	gorm.Model
}

type ApiLog struct {
	ID                   uuid.UUID `gorm:"type:uuid;;primaryKey" json:"id"`
	Duration             int64     `gorm:"type:int8" json:"duration"`
	EndPoint             string    `gorm:"type:varchar(50)" json:"end_point"`
	HttpStatus           string    `gorm:"type:varchar(5)" json:"http_status"`
	Method               string    `gorm:"type:varchar(10)" json:"method"`
	TsStart              string    `gorm:"type:varchar(20)" json:"ts_start"`
	TsEnd                string    `gorm:"type:varchar(20)" json:"ts_end"`
	UserAgent            string    `gorm:"type:varchar(30)" json:"user_agent"`
	ReqID                string    `gorm:"type:varchar(50)" json:"req_id"`
	Origin               string    `gorm:"type:varchar(100)" json:"origin"`
	ReqSize              int64     `gorm:"type:int8" json:"req_size"`
	InstanceName         string    `gorm:"type:varchar(20)" json:"instance_name"`
	ClientIP             string    `gorm:"type:varchar(20)" json:"client_ip"`
	UpstreamElapsed      int64     `gorm:"type:int8" json:"upstream_elapsed"`
	RemoteHost           string    `gorm:"type:varchar(50)" json:"remote_host"`
	UserName             string    `gorm:"type:varchar(50)" json:"user_name"`
	Params               string    `gorm:"type:varchar(50)" json:"params"`
	ResSize              int64     `gorm:"type:int8" json:"res_size"`
	SendToUpstream       int64     `gorm:"type:int8" json:"send_to_upstream"`
	UpstreamResponseCode *string   `gorm:"type:varchar(50)" json:"upstream_response_code"`
	Col1                 *string   `gorm:"type:varchar(50)" json:"col1"`
	Col10                *string   `gorm:"type:varchar(50)" json:"col10"`
	Col2                 *string   `gorm:"type:varchar(50)" json:"col2"`
	Col3                 *string   `gorm:"type:varchar(50)" json:"col3"`
	Col4                 *string   `gorm:"type:varchar(50)" json:"col4"`
	Col5                 *string   `gorm:"type:varchar(50)" json:"col5"`
	Col6                 *string   `gorm:"type:varchar(50)" json:"col6"`
	Col7                 *string   `gorm:"type:varchar(50)" json:"col7"`
	Col8                 *string   `gorm:"type:varchar(50)" json:"col8"`
	Col9                 *string   `gorm:"type:varchar(50)" json:"col9"`
}
