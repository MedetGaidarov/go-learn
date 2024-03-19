package grpc

import (
	"context"
	"errors"

	ssov1 "github.com/MedetGaidarov/go-protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverApi struct {
	ssov1.UnimplementedAuthServer // Хитрая штука, о ней ниже
	auth                          Auth
}

// Тот самый интерфейс, котрый мы передавали в grpcApp
type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appId int,
	) (token string, err error)
	RegiserNewUser(
		ctx context.Context,
		email string,
		password string,

	) (userId int64, err error)
}

func Register(gRPCServer *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPCServer, &serverApi{auth: auth})
}

func (s *serverApi) Login(
	ctx context.Context,
	in *ssov1.LoginRequest,
) (*ssov1.LoginRespone, error) {
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	if in.GetAppId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	token, err := s.auth.LoginAuth(ctx, in.Email, in.Password, int(in.GetAppId()))

	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		}
		return nil, status.Error(codes.InvalidArgument, "invalid email or password")
	}
	return &ssov1.LoginResponse{Token: token}, nil

}

func (s *serverApi) Register(
	ctx context.Context,
	in *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	uid, error := s.auth.RegisterNewUser(ctx, in.GetEmail(), in.GetPassword())

	if err != nil {
		// Ошибку storage.ErrUserExists мы создадим ниже
		if errors.Is(err, storage.ErrUserExists) {

			return nil, status.Error(codes.AlreadyExists, "user already exists")

		}
		return nil, status.Error(codes.Internal, "failed to register user")

	}
	return &ssov1.RegisterResponse{UserId: uid}, nil

}
