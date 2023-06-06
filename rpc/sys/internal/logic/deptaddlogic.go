package logic

import (
	"context"
	"fmt"
	"strings"
	"zero-admin/rpc/model/sysmodel"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptAddLogic {
	return &DeptAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptAddLogic) DeptAdd(in *sys.DeptAddReq) (*sys.DeptAddResp, error) {
	_, err := l.svcCtx.DeptModel.Insert(l.ctx, &sysmodel.SysDept{
		Name:      in.Name,
		ParentId:  in.ParentId,
		OrderNum:  in.OrderNum,
		CreateBy:  in.CreateBy,
		DelFlag:   in.DelFlag,
		ParentIds: strings.Replace(strings.Trim(fmt.Sprint(in.ParentIds), "[]"), " ", ",", -1),
	})

	if err != nil {
		return nil, err
	}

	return &sys.DeptAddResp{}, nil
}
