package westcrm

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/westcrm"
)

// AlibabaWestcrmCustomerInfoGet 会员信息查询接口
// alibaba.westcrm.customer.info.get
//
// 会员信息查询接口
func AlibabaWestcrmCustomerInfoGet(ctx context.Context, clt *core.SDKClient, req *westcrm.AlibabaWestcrmCustomerInfoGetAPIRequest, resp *westcrm.AlibabaWestcrmCustomerInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
