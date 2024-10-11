package ctrl

import (
	"context"
	"fmt"
	"studyGrpc/internal/service"
	userPorto "studyGrpc/proto"
)

type AuthController struct {
	// 必须嵌入进入
	userPorto.UnimplementedAuthServiceServer
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// 验证参数
func validateRequest(request *userPorto.LoginRequest) error {
	return nil
}

func (a AuthController) Login(ctx context.Context, request *userPorto.LoginRequest) (*userPorto.LoginResponse, error) {

	//1. 验证
	if err := validateRequest(request); err != nil {
		return nil, err
	}

	fmt.Println("user login --> ", request.Username, request.Password)
	// 2. 业务 biz
	user, err := service.NewAuthService().Login(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	// 3. 组装相应
	resp := &userPorto.LoginResponse{
		Token: "123",
		User: &userPorto.User{
			Id: user.Id,
		},
	}
	// 4. 返回
	return resp, nil
}

func (a AuthController) Register(ctx context.Context, request *userPorto.RegisterRequest) (*userPorto.RegisterResponse, error) {
	return nil, nil
}
