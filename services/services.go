package services

import (
	"fmt"

	"github.com/sulton0011/api-gateway/config"
	pb "github.com/sulton0011/api-gateway/genproto/task"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	TaskService() pb.TaskServiceClient
}

type serviceManager struct {
	TasksService pb.TaskServiceClient
}

func (s *serviceManager) TaskService() pb.TaskServiceClient {
	return s.TasksService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connTask, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.TaskServiceHost, conf.TaskServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		TasksService: pb.NewTaskServiceClient(connTask),
	}

	return serviceManager, nil
}
