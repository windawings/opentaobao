package scbp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/scbp"
)

// AlibabaScbpAdAccountLevelGet 查询推广账户等级
// alibaba.scbp.ad.account.level.get
//
// 查询推广账户等级
func AlibabaScbpAdAccountLevelGet(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdAccountLevelGetAPIRequest, resp *scbp.AlibabaScbpAdAccountLevelGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
