package scbp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/scbp"
)

// AlibabaScbpEffectAccountList 账户-报表
// alibaba.scbp.effect.account.list
//
// 账户-报表,支持最近7天，最近30天，以及180天内时间区间。
func AlibabaScbpEffectAccountList(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpEffectAccountListAPIRequest, resp *scbp.AlibabaScbpEffectAccountListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
