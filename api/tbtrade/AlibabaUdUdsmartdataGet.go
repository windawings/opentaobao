package tbtrade

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbtrade"
)

// AlibabaUdUdsmartdataGet udsmart获取数据回传接口
// alibaba.ud.udsmartdata.get
//
// udsmart获取数据回传接口
func AlibabaUdUdsmartdataGet(ctx context.Context, clt *core.SDKClient, req *tbtrade.AlibabaUdUdsmartdataGetAPIRequest, resp *tbtrade.AlibabaUdUdsmartdataGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
