package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusSpaceUnitGetlistbygroupid 根据分组ID查询相应的空间单元
// alibaba.campus.space.unit.getlistbygroupid
//
// 根据分组ID查询相应的空间单元
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getListByGroupId
func AlibabaCampusSpaceUnitGetlistbygroupid(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest, resp *campus.AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
