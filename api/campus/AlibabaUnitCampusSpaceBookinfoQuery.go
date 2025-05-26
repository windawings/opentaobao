package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaUnitCampusSpaceBookinfoQuery 环路资源信息查询单元环境
// alibaba.unit.campus.space.bookinfo.query
//
// 环路资源信息查询单元环境
func AlibabaUnitCampusSpaceBookinfoQuery(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaUnitCampusSpaceBookinfoQueryAPIRequest, resp *campus.AlibabaUnitCampusSpaceBookinfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
