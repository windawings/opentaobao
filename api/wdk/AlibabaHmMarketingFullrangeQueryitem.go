package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaHmMarketingFullrangeQueryitem 全场活动查询换购品
// alibaba.hm.marketing.fullrange.queryitem
//
// 全场活动查询换购品
func AlibabaHmMarketingFullrangeQueryitem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingFullrangeQueryitemAPIRequest, resp *wdk.AlibabaHmMarketingFullrangeQueryitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
