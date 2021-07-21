package grpc

import (
	"context"
	"errors"
	"io"
	"weather/domain"

	pb "weather/gen/pb-go/weather"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type WeatherHandler struct {
	WeatherUsecase domain.WeatherUseCase
	pb.UnimplementedWeatherServer
}

func NewWeatherHandler(s *grpc.Server, weatherUsecase domain.WeatherUseCase) {
	handler := &WeatherHandler{
		WeatherUsecase: weatherUsecase,
	}

	pb.RegisterWeatherServer(s, handler)
}

func mappingWeatherEnum(weather domain.WeatherEnum) (pb.QueryResponse_Weather, error) {
	switch weather {
	case domain.SUNNY:
		return pb.QueryResponse_SUNNY, nil
	case domain.CLOUDY:
		return pb.QueryResponse_CLOUDY, nil
	default:
		return pb.QueryResponse_SUNNY, errors.New("This weather does not exists.")
	}
}

func (w *WeatherHandler) Query(srv pb.Weather_QueryServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			logrus.Info("client send end stream...")
			return nil
		}
		if err != nil {
			logrus.Error(err)
			return err
		}
		location := req.GetLocation()
		res, err := w.WeatherUsecase.GetByLocation(context.Background(), location)
		if err != nil {
			logrus.Error(err)
			return err
		}
		weatherEnum, err := mappingWeatherEnum(res.Weather)
		if err != nil {
			logrus.Error(err)
			return err
		}
		srv.Send(
			&pb.QueryResponse{
				Location: res.Location,
				Weather:  weatherEnum,
			})
	}
}
