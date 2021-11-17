package rpc

import (
	pb "common_service/internal/proto"
	"common_service/internal/service"
	"common_service/pkg/app"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServer struct{}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func Ui32PToUiP(ui *uint32) *uint {
	if ui == nil {
		return nil
	}
	i := uint(*ui)
	return &i
}

func (s *UserServer) GetUserList(ctx context.Context, r *pb.GetUserListRequest) (*pb.GetUserListReply, error) {
	svc := service.New(nil)
	page := app.Pager{
		Page:     int(r.Page),
		PageSize: int(r.PageSize),
	}
	param := service.ListUserRequest{
		Nickname: r.Nickname,
		Status:   Ui32PToUiP(r.State),
		RoleId:   r.RoleId,
	}
	users, totalCount, err := svc.ListUser(&param, &page)
	if err != nil {
		return nil, err
	}

	var dataList []*pb.User
	for _, u := range users {
		data := &pb.User{
			Id:        u.ID,
			Username:  u.Username,
			Nickname:  u.Nickname,
			State:     uint32(u.Status),
			RoleId:    u.RoleId,
			RoleName:  u.RoleName,
			CreatedAt: timestamppb.New(u.CreatedAt),
			UpdatedAt: timestamppb.New(u.UpdatedAt),
		}
		dataList = append(dataList, data)
	}
	pager := &pb.Pager{
		Page:      r.Page,
		PageSize:  r.PageSize,
		TotalRows: int64(totalCount),
	}
	replay := &pb.GetUserListReply{
		Data:  dataList,
		Pager: pager,
	}

	return replay, nil
}
