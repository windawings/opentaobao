package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryNotItemPromotion 其他主体报表查询
// taobao.universalbp.report.query.not.item.promotion
//
// 其他主体报表查询
func TaobaoUniversalbpReportQueryNotItemPromotion(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest, resp *simba.TaobaoUniversalbpReportQueryNotItemPromotionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
