package dto

import (
	"time"

	"gitlab.com/techpotion/leadershack2022/api/entity"
)

type GetPointsRequestQueryDTO struct {
	Region          string    `form:"region"          validate:"required"`
	DateTimeFrom    time.Time `form:"datetime_from"   validate:"required"`
	DateTimeTo      time.Time `form:"datetime_to"     validate:"required"`
	XMin            *float64  `form:"x_min"           validate:"gte=-180,lte=180"`
	Ymin            *float64  `form:"y_min"           validate:"gte=-90,lte=90"`
	XMax            *float64  `form:"x_max"           validate:"gte=-180,lte=180"`
	YMax            *float64  `form:"y_max"           validate:"gte=-90,lte=90"`
	Limit           int       `form:"limit"           validate:"required,gt=0,lte=10000"`
	Offest          *int      `form:"offset"          validate:"required,gte=0"`
	UrgencyCategory *string   `form:"urgency_category"`
}

type GetPointsResponseDTO struct {
	Points  []entity.HCSPoint `json:"points"`
	Count   int               `json:"count"`
	Success bool              `json:"success"`
	Error   string            `json:"error,omitempty"`
}
