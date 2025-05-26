package zqs

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/zqs"
)

// AlibabaZqsFulfillComplete 周期购履约完成接口
// alibaba.zqs.fulfill.complete
//
// 周期购履约完成接口
func AlibabaZqsFulfillComplete(ctx context.Context, clt *core.SDKClient, req *zqs.AlibabaZqsFulfillCompleteAPIRequest, resp *zqs.AlibabaZqsFulfillCompleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
