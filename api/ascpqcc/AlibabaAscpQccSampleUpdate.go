package ascpqcc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/ascpqcc"
)

// AlibabaAscpQccSampleUpdate 品控中心更新样品信息
// alibaba.ascp.qcc.sample.update
//
// 品控中心更新样品信息
func AlibabaAscpQccSampleUpdate(ctx context.Context, clt *core.SDKClient, req *ascpqcc.AlibabaAscpQccSampleUpdateAPIRequest, resp *ascpqcc.AlibabaAscpQccSampleUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
