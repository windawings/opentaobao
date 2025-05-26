package larkiot

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/larkiot"
)

// TaobaoLarkPosBasedataGetworkstation 根据影城id工作站和macId获取工作站
// taobao.lark.pos.basedata.getworkstation
//
// 获取单独工作站
func TaobaoLarkPosBasedataGetworkstation(ctx context.Context, clt *core.SDKClient, req *larkiot.TaobaoLarkPosBasedataGetworkstationAPIRequest, resp *larkiot.TaobaoLarkPosBasedataGetworkstationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
