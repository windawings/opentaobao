package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusAdminmapPoiinfoGetlistbygroup 根据分组条件查询分组下的空间单元不包涵业务属性信息
// alibaba.campus.adminmap.poiinfo.getlistbygroup
//
// 根据分组条件查询分组下的空间单元不包涵业务属性信息
func AlibabaCampusAdminmapPoiinfoGetlistbygroup(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIRequest, resp *campus.AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
