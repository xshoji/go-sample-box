package proto_impl

import (
	"context"
	"github.com/xshoji/go-sample-box/grpcGateway/proto"
	"sync"
)

type UserService struct{}

var UserRepositoryMap = make(map[int64]proto.User)
var idSequence int64 = 1
var mu sync.Mutex

func (*UserService) Create(context context.Context, user *proto.User) (*proto.User, error) {
	mu.Lock()
	defer mu.Unlock()
	user.Id = idSequence
	UserRepositoryMap[idSequence] = *user
	idSequence++
	return user, nil
}

func (*UserService) Read(ctx context.Context, in *proto.SimpleRequest) (*proto.User, error) {
	foundUser, ok := UserRepositoryMap[in.Id]
	if ok != true {
		foundUser = proto.User{Id: 0, Name: "UnknownUser", Age: 0}
	}
	return &foundUser, nil
}

func (*UserService) Update(context context.Context, user *proto.User) (*proto.User, error) {
	UserRepositoryMap[user.Id] = *user
	return user, nil
}

func (*UserService) Delete(context context.Context, in *proto.SimpleRequest) (*proto.User, error) {
	foundUser, ok := UserRepositoryMap[in.Id]
	if ok {
		delete(UserRepositoryMap, in.Id)
	}
	if ok != true {
		foundUser = proto.User{Id: 0, Name: "UnknownUser", Age: 0}
	}
	return &foundUser, nil
}
