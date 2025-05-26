package eleenterpriseordernew

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/eleenterpriseordernew"
)

// AlibabaEleEnterpriseOrdernewCancel 订单取消
// alibaba.ele.enterprise.ordernew.cancel
//
// 订单取消
func AlibabaEleEnterpriseOrdernewCancel(ctx context.Context, clt *core.SDKClient, req *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewCancelAPIRequest, resp *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
