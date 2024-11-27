package services

import (
	"context"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"qinglong-envs/internal/db/queries"
	"qinglong-envs/pkg/api"
)

type UserService struct {
	db           *sql.DB
	queries      *queries.Queries
	tokenService *TokenService
}

func NewUserService(conn *sql.DB, tokenService *TokenService) *UserService {
	return &UserService{
		db:           conn,
		queries:      queries.New(conn),
		tokenService: tokenService,
	}
}

type User struct {
	Id          int    `json:"id,string"`
	Username    string `json:"username"`
	Password    string `json:"-"`
	WxPusherUid string `json:"wxPusherUID"`
}

func (s *UserService) GetUserById(ctx context.Context, id int) (*User, error) {
	user, err := s.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &User{
		Id:          user.ID,
		Username:    user.Username,
		Password:    user.Password.String,
		WxPusherUid: user.WxPusherUID.String,
	}, nil
}

func (s *UserService) FindUserByUsername(ctx context.Context, username string) (*User, error) {
	user, err := s.queries.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return &User{
		Id:          user.ID,
		Username:    user.Username,
		Password:    user.Password.String,
		WxPusherUid: user.WxPusherUID.String,
	}, nil
}

type UsernamePasswordInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginSuccess struct {
	Token string `json:"token"`
}

func (s *UserService) LoginByLocal(ctx context.Context, input UsernamePasswordInput) (*LoginSuccess, error) {
	user, err := s.queries.GetUserByUsername(ctx, input.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, api.NotPermittedError("用户名或密码错误")
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(input.Password))
	if err != nil {
		return nil, api.NotPermittedError("用户名或密码错误")
	}
	token, err := s.tokenService.CreateToken(ctx, user.ID)
	if err != nil {
		return nil, api.InternalError("登录失败: 创建令牌失败")
	}
	return &LoginSuccess{Token: token}, nil
}

func (s *UserService) Register(ctx context.Context, input UsernamePasswordInput) (*User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user, err := s.queries.CreateUser(ctx, queries.CreateUserParams{
		Username:    input.Username,
		Password:    sql.NullString{String: string(password), Valid: true},
		WxPusherUID: sql.NullString{},
	})
	if err != nil {
		return nil, err
	}
	return &User{
		Id:       user.ID,
		Username: user.Username,
		Password: user.Password.String,
	}, nil
}
