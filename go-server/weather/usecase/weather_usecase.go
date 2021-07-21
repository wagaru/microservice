package usecase

import (
	"context"
	"go-server/domain"
)

type WeatherUseCase struct {
	weatherRepo domain.WeatherRepository
}

func NewWeatherUsecase(repo domain.WeatherRepository) domain.WeatherUseCase {
	return &WeatherUseCase{
		weatherRepo: repo,
	}
}

func (w *WeatherUseCase) GetStreamByLocation(ctx context.Context, location string) (domain.StreamWeather, error) {
	return w.weatherRepo.GetStreamByLocation(ctx, location)
}
