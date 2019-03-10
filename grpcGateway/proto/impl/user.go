package proto_impl

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/xshoji/go-sample-box/grpcGateway/proto"
	"sync"
)

type UserService struct{}

var UserRepositoryMap = make(map[int64]*proto.User)
var idSequence int64 = 1
var mu sync.Mutex

func createResponse(status int64, message string) *proto.UserResponse {
	return &proto.UserResponse{
		Status:  status,
		Message: message,
	}
}

func createOkResponse() *proto.UserResponse {
	return createResponse(200, "OK")
}

func createNotFoundResponse() *proto.UserResponse {
	return createResponse(500, "Not found")
}

func (*UserService) Create(context context.Context, user *proto.User) (*proto.UserResponse, error) {
	mu.Lock()
	defer mu.Unlock()
	user.Id = idSequence
	UserRepositoryMap[idSequence] = user
	idSequence++
	response := createOkResponse()
	response.Users = []*proto.User{user}
	return response, nil
}

func (*UserService) Read(ctx context.Context, in *proto.SimpleRequest) (*proto.UserResponse, error) {
	foundUser, ok := UserRepositoryMap[in.Id]
	if ok != true {
		return createNotFoundResponse(), nil
	}
	response := createOkResponse()
	response.Users = []*proto.User{foundUser}
	return response, nil
}

func (*UserService) ReadAll(ctx context.Context, in *empty.Empty) (*proto.UserResponse, error) {
	users := make([]*proto.User, 0, len(UserRepositoryMap))
	for _, v := range UserRepositoryMap {
		users = append(users, v)
	}
	response := createOkResponse()
	response.Users = users

	return response, nil
}

func (*UserService) Update(context context.Context, user *proto.User) (*proto.UserResponse, error) {
	foundUser, ok := UserRepositoryMap[user.Id]
	if ok != true {
		return createNotFoundResponse(), nil
	}

	fmt.Println(user)

	if user.Name != "" {
		foundUser.Name = user.Name
	}
	if user.GetAgeOptional() != nil {
		foundUser.AgeOptional = user.GetAgeOptional()
	}
	response := createOkResponse()
	response.Users = []*proto.User{foundUser}
	return response, nil
}

func (*UserService) Delete(context context.Context, in *proto.SimpleRequest) (*proto.UserResponse, error) {
	foundUser, ok := UserRepositoryMap[in.Id]
	if ok {
		delete(UserRepositoryMap, in.Id)
	} else {
		return createNotFoundResponse(), nil
	}

	response := createOkResponse()
	response.Users = []*proto.User{foundUser}
	return response, nil
}
