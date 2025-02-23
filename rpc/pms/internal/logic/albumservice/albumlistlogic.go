package albumservicelogic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlbumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumListLogic {
	return &AlbumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumListLogic) AlbumList(in *pmsclient.AlbumListReq) (*pmsclient.AlbumListResp, error) {
	all, err := l.svcCtx.PmsAlbumModel.FindAll(l.ctx, in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsAlbumModel.Count(l.ctx)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询相册列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pmsclient.AlbumListData
	for _, item := range *all {

		list = append(list, &pmsclient.AlbumListData{
			Id:          item.Id,
			Name:        item.Name,
			CoverPic:    item.CoverPic,
			PicCount:    item.PicCount,
			Sort:        item.Sort,
			Description: item.Description,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询相册列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pmsclient.AlbumListResp{
		Total: count,
		List:  list,
	}, nil
}
