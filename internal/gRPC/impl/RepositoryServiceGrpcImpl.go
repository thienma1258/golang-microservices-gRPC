package impl

import (
	"context"
	"golang-microservices-gRPC/internal/gRPC/domain"
	"golang-microservices-gRPC/internal/gRPC/service"
	"log"
	"strconv"
)

type RepositoryServiceGrpcImpl struct {
}

//NewRepositoryServiceGrpcImpl returns the pointer to the implementation.
func NewRepositoryServiceGrpcImpl() *RepositoryServiceGrpcImpl {
	return &RepositoryServiceGrpcImpl{}
}

//Add function implementation of gRPC Service.
func (serviceImpl *RepositoryServiceGrpcImpl) Add(ctx context.Context, in *domain.Repository) (*service.AddRepositoryResponse, error) {
	log.Println("Received request for adding repository with id " + strconv.FormatInt(in.Id, 10))

	//Logic to persist to database or storage.
	log.Println("Repository persisted to the storage")

	return &service.AddRepositoryResponse{
		AddedRepository: in,
		Error:           nil,
	}, nil
}