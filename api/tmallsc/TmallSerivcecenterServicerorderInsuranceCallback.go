package tmallsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallsc"
)

// TmallSerivcecenterServicerorderInsuranceCallback 服务商回传保单信息
// tmall.serivcecenter.servicerorder.insurance.callback
//
// 服务商回传保单信息
func TmallSerivcecenterServicerorderInsuranceCallback(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallSerivcecenterServicerorderInsuranceCallbackAPIRequest, resp *tmallsc.TmallSerivcecenterServicerorderInsuranceCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
