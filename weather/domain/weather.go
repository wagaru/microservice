package domain

import "context"

type WeatherEnum int32

const (
	SUNNY WeatherEnum = iota
	CLOUDY
)

type Weather struct {
	Location string
	Weather  WeatherEnum
}

type WeatherRepository interface {
	GetByLocation(ctx context.Context, location string) (*Weather, error)
}

type WeatherUseCase interface {
	GetByLocation(ctx context.Context, location string) (*Weather, error)
}
