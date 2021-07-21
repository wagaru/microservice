package fake

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"weather/domain"
)

type weatherRepo struct {
}

func NewFakeWeatherRepository() domain.WeatherRepository {
	return &weatherRepo{}
}

func createRandomWeather() domain.WeatherEnum {
	rand.Seed(time.Now().UnixNano())
	switch randomNum := rand.Intn(2); randomNum {
	case 0:
		return domain.SUNNY
	case 1:
		return domain.CLOUDY
	default:
		return domain.SUNNY
	}
}

func (w *weatherRepo) GetByLocation(ctx context.Context, location string) (*domain.Weather, error) {
	switch location {
	case "A":
		return &domain.Weather{
			Location: "A",
			Weather:  createRandomWeather(),
		}, nil
	default:
		return nil, errors.New("This location does not exist.")
	}
}
