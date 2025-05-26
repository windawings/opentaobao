package nropen

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/nropen"
)

// AlibabaAscpIndustryDisivisonQuery 查询服务支持地区列表
// alibaba.ascp.industry.disivison.query
//
// 商家获取服务支持地区
func AlibabaAscpIndustryDisivisonQuery(ctx context.Context, clt *core.SDKClient, req *nropen.AlibabaAscpIndustryDisivisonQueryAPIRequest, resp *nropen.AlibabaAscpIndustryDisivisonQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
