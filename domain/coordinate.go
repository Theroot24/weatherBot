package domain

type Coordinate struct {
	Lat float32 `json:"lat" validate:"required,gte=-90,lte=90"`
	Lon float32 `json:"lon" validate:"required,gte=-180,lte=180"`
}
