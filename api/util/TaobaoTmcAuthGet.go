package util

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/util"
)

// TaobaoTmcAuthGet TMC授权token
// taobao.tmc.auth.get
//
// TMC连接授权Token
func TaobaoTmcAuthGet(ctx context.Context, clt *core.SDKClient, req *util.TaobaoTmcAuthGetAPIRequest, resp *util.TaobaoTmcAuthGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
