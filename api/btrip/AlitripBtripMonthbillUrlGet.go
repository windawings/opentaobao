package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripMonthbillUrlGet 月账单数据查询
// alitrip.btrip.monthbill.url.get
//
// 月账单数据查询
func AlitripBtripMonthbillUrlGet(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripMonthbillUrlGetAPIRequest, resp *btrip.AlitripBtripMonthbillUrlGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
