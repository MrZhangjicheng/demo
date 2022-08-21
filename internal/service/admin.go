package service

import (
	"context"
	pb "demo/api/auth/service/v1"
	"errors"
)

func (a *Auth) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("参数为空")
	}
	resp := &pb.LoginReply{}
	resp.AccessTocken = "sssss"
	resp.TockenType = "test"
	resp.ExpiresIn = 9000
	a.log.Debug("zap 日志测试")
	return resp, nil
}

func (a *Auth) Logout(context.Context, *pb.LogoutReq) (*pb.LogoutReply, error) {
	return nil, nil

}
func (a *Auth) Refresh(context.Context, *pb.RefreshReq) (*pb.RefreshReply, error) {
	return nil, nil
}
func (a *Auth) UpdatePassword(context.Context, *pb.UpdatePasswordReq) (*pb.UpdatePasswordReply, error) {
	return nil, nil
}
func (a *Auth) UpdateUser(context.Context, *pb.UpdateUserReq) (*pb.UpdateUserReply, error) {
	return nil, nil
}
func (a *Auth) PasswordForgotten(context.Context, *pb.PasswordForgottenReq) (*pb.PasswordForgottenReply, error) {
	return nil, nil
}
func (a *Auth) PasswordForgottenTocken(context.Context, *pb.PasswordForgottenTockenReq) (*pb.PasswordForgottenTockenReply, error) {
	return nil, nil
}
