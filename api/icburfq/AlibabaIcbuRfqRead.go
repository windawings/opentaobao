package icburfq

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/icburfq"
)

// AlibabaIcbuRfqRead 是否已读RFQ
// alibaba.icbu.rfq.read
//
// 是否已读RFQ
func AlibabaIcbuRfqRead(ctx context.Context, clt *core.SDKClient, req *icburfq.AlibabaIcbuRfqReadAPIRequest, resp *icburfq.AlibabaIcbuRfqReadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
