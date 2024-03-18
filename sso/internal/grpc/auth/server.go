import (
	"context"
	"google.golang.org."

	ssov1 "github.com/MedetGaidarov/go-protos"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer // Хитрая штука, о ней ниже
	auth                          Auth
}


// Тот самый интерфейс, котрый мы передавали в grpcApp
type Auth interface  {
	Login (
		ctx context.Context, 
		email string, 
		password string,
		appId int
	) (token string, err error)
	RegiserNewUser (
		ctx context.Context, 
        email string, 
        password string,
       
	) (userId int64, err error)
}

func Register(gRPCServer *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPCServer, &serverApi{auth: auth})
}


func (s *serverApi) Login (
	ctx context.context,
	in *ssov1.LoginRequest,
) (*ssov1.LoginRespone, error) {
	// TODO:
}

func (s *serverApi ) Register (
	ctx context.context,
    in *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	// TODO:
}
