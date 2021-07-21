package microservice

import (
	"context"
	"go-server/domain"

	pb "go-server/gen/pb-go/weather"

	"github.com/sirupsen/logrus"
)

const (
	address = "localhost:123"
)

type WeatherRepo struct {
	weatherGRPC pb.WeatherClient
}

type WeatherClient struct {
	client pb.Weather_QueryClient
}

func NewWeatherRepo(weather pb.WeatherClient) domain.WeatherRepository {
	return &WeatherRepo{
		weatherGRPC: weather,
	}
}

func (w *WeatherRepo) GetStreamByLocation(ctx context.Context, location string) (domain.StreamWeather, error) {
	client, err := w.weatherGRPC.Query(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &WeatherClient{
		client: client,
	}, nil
}

func (wc *WeatherClient) Send(w *domain.Weather) error {
	if err := wc.client.Send(&pb.QueryRequest{Location: w.Location}); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (wc *WeatherClient) Recv() (*domain.Weather, error) {
	weather, err := wc.client.Recv()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &domain.Weather{
		Location: weather.GetLocation(),
		Weather:  string(weather.GetWeather()),
	}, nil
}
