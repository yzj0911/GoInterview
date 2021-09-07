package xorm

import "time"

type Report struct {
	Date       string  `json:"date" xorm:"DATE NOT NULL COMMENT('日期')"`
	Impression float64 `json:"impression" xorm:"INT(11) NOT NULL DEFAULT 0 COMMENT('展现')"`
	Click      float64 `json:"click" xorm:"INT(11) NOT NULL DEFAULT 0 COMMENT('点击')"`
	Cost       float64 `json:"cost" xorm:"FLOAT NOT NULL DEFAULT 0 COMMENT('消费')"`
	Ctr        float64 `json:"ctr" xorm:"FLOAT NOT NULL DEFAULT 0 COMMENT('点击率')"`
	Cpm        float64 `json:"cpm" xorm:"FLOAT NOT NULL DEFAULT 0 COMMENT('均价')"`
	Pv         float64 `json:"pv" xorm:"INT(11) NOT NULL DEFAULT 0 COMMENT('页面浏览量')"`
	Uv         float64 `json:"uv" xorm:"INT(11) NOT NULL DEFAULT 0 COMMENT('访客数量')"`
}

type Time time.Time

type ReportBaidu struct {
	Id        int64 `json:"id" xorm:"BIGINT(20) PK AUTOINCR COMMENT('ID')"`
	Report    `xorm:"extends"`
	Ar        float64 `json:"ar" xorm:"FLOAT NOT NULL DEFAULT 0 COMMENT('抵达率')"`
	CreatedAt Time    `json:"created_at" xorm:"DATETIME NOT NULL CREATED COMMENT('创建时间')"`
}

type ReportGoogle struct {
	Id        int64 `json:"id" xorm:"BIGINT(20) PK AUTOINCR COMMENT('ID')"`
	Report    `xorm:"extends"`
	CreatedAt Time `json:"created_at" xorm:"DATETIME NOT NULL CREATED COMMENT('创建时间')"`
}
