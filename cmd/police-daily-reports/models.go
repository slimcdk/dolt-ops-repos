package main

import (
	"time"
)

type ReportStructError struct {
	Message string `json:"Message"`
}

type ReportStruct struct {
	ID                 string `gorm:"" json:"Id"`
	DistrictName       string `gorm:"" json:"DistrictName"`
	ArticleType        string `gorm:"" json:"ArticleType"`
	PublishDate        string `gorm:"" json:"PublishDate"`
	Link               string `gorm:"" json:"Link"`
	ListDate           string `gorm:"" json:"ListDate"`
	Headline           string `gorm:"" json:"Headline"`
	Manchet            string `gorm:"" json:"Manchet"`
	Article            string `gorm:"" json:"Article"`
	ToolTip            string `gorm:"" json:"ToolTip"`
	Image              string `gorm:"" json:"Image"`
	ImageDescription   string `gorm:"" json:"ImageDescription"`
	PhotographerText   string `gorm:"" json:"PhotographerText"`
	NoPhoto            bool   `gorm:"" json:"NoPhoto"`
	DisplayDescription bool   `gorm:"" json:"DisplayDescription"`
}

type ReportStructSuccess struct {
	NewsList          []ReportStruct `json:"NewsList"`
	TotalNumberOfNews uint           `json:"TotalNumberOfNews"`
}

type QueryOptions struct {
	ItemID        string    `url:"itemId"`
	IsNewsList    bool      `url:"isNewsList"`
	NewsType      string    `url:"newsType"`
	DistrictQuery string    `url:"districtQuery"`
	FromDate      time.Time `url:"fromDate"`
	ToDate        time.Time `url:"toDate"`
	Page          int       `url:"page"`
	PageSize      int32     `url:"pageSize"`
}
