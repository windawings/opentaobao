package mos

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/mos"
)

// AlibabaMjMosFundCreatebill 创建一个付款单
// alibaba.mj.mos.fund.createbill
//
// 创建一个付款单
func AlibabaMjMosFundCreatebill(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMjMosFundCreatebillAPIRequest, resp *mos.AlibabaMjMosFundCreatebillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
