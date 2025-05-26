package openim

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/openim"
)

// TaobaoOpenimTribeGettribeinfo 获取群信息
// taobao.openim.tribe.gettribeinfo
//
// 获取群信息
func TaobaoOpenimTribeGettribeinfo(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimTribeGettribeinfoAPIRequest, resp *openim.TaobaoOpenimTribeGettribeinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
