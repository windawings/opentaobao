package alsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveSnsConverturl 防封接口
// alibaba.alsc.growth.interactive.sns.converturl
//
// 防封接口
func AlibabaAlscGrowthInteractiveSnsConverturl(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveSnsConverturlAPIRequest, resp *alsc.AlibabaAlscGrowthInteractiveSnsConverturlAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
