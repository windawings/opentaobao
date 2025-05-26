package charity

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/charity"
)

// AlibabaCharityCharitytimeQuerytime 外部查询公益时
// alibaba.charity.charitytime.querytime
//
// 外部查询公益时
func AlibabaCharityCharitytimeQuerytime(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCharityCharitytimeQuerytimeAPIRequest, resp *charity.AlibabaCharityCharitytimeQuerytimeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
