package grpc

import (
	"context"
	"go-server/domain"

	pb "go-server/gen/pb-go/gen"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DigimonHandler struct {
	DigimonUsecase domain.DigimonUsecase
	DietUsecase    domain.DietUsecase
	WeatherUsecase domain.WeatherUseCase
	pb.UnimplementedDigimonServer
}

func NewDigimonHandler(s *grpc.Server, digimonUsecase domain.DigimonUsecase, dietUsecase domain.DietUsecase, weatherUsecase domain.WeatherUseCase) {
	handler := &DigimonHandler{
		DigimonUsecase: digimonUsecase,
		DietUsecase:    dietUsecase,
		WeatherUsecase: weatherUsecase,
	}

	pb.RegisterDigimonServer(s, handler)
}

func (d *DigimonHandler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	aDigimon := domain.Digimon{
		Name: req.GetName(),
	}
	if err := d.DigimonUsecase.Store(ctx, &aDigimon); err != nil {
		logrus.Error(err)
		return nil, status.Errorf(codes.Internal, "Internal error. Store failed")
	}
	return &pb.CreateResponse{
		Id:     aDigimon.ID,
		Name:   aDigimon.Name,
		Status: aDigimon.Status,
	}, nil
}

func (d *DigimonHandler) QueryStream(req *pb.QueryRequest, srv pb.Digimon_QueryStreamServer) error {
	weatherClient, err := d.WeatherUsecase.GetStreamByLocation(context.Background(), "A")
	if err != nil {
		logrus.Error(err)
		return err
	}
	for {
		if err := weatherClient.Send(&domain.Weather{
			Location: "A",
		}); err != nil {
			logrus.Error(err)
			return err
		}

		aWeather, err := weatherClient.Recv()
		if err != nil {
			logrus.Error(err)
			return err
		}

		anDigimon, err := d.DigimonUsecase.GetByID(context.Background(), req.GetId())
		if err != nil {
			logrus.Error(err)
			return err
		}
		srv.Send(&pb.QueryResponse{
			Id:       anDigimon.ID,
			Name:     anDigimon.Name,
			Status:   anDigimon.Status,
			Location: aWeather.Location,
			Weather:  aWeather.Weather,
		})

	}
}

func (d *DigimonHandler) Foster(ctx context.Context, req *pb.FosterRequest) (*pb.FosterResponse, error) {
	if err := d.DietUsecase.Store(ctx, &domain.Diet{
		UserID: req.GetId(),
		Name:   req.GetFood().GetName(),
	}); err != nil {
		logrus.Error(err)
		return nil, status.Errorf(codes.Internal, "Internal error. Store failed")
	}

	if err := d.DigimonUsecase.UpdateStatus(ctx, &domain.Digimon{
		ID:     req.GetId(),
		Status: "good",
	}); err != nil {
		logrus.Error(err)
		return nil, status.Errorf(codes.Internal, "Internal error. Store failed")
	}

	return &pb.FosterResponse{}, nil
}
