package mos

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/mos"
)

// AlibabaMjMoscarnivalReceiveencrypt 根据加密手机号领券
// alibaba.mj.moscarnival.receiveencrypt
//
// 根据加密手机号领券
func AlibabaMjMoscarnivalReceiveencrypt(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMjMoscarnivalReceiveencryptAPIRequest, resp *mos.AlibabaMjMoscarnivalReceiveencryptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
