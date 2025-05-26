package mos

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/mos"
)

// AlibabaMosTmcSmsSend 银泰TMC发送短信
// alibaba.mos.tmc.sms.send
//
// 银泰TMC发送短信
func AlibabaMosTmcSmsSend(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMosTmcSmsSendAPIRequest, resp *mos.AlibabaMosTmcSmsSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
