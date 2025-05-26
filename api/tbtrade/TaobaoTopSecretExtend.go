package tbtrade

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbtrade"
)

// TaobaoTopSecretExtend 虚拟号延期
// taobao.top.secret.extend
//
// 虚拟号延期
func TaobaoTopSecretExtend(ctx context.Context, clt *core.SDKClient, req *tbtrade.TaobaoTopSecretExtendAPIRequest, resp *tbtrade.TaobaoTopSecretExtendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
