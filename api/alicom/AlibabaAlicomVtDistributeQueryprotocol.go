package alicom

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alicom"
)

// AlibabaAlicomVtDistributeQueryprotocol 通信业务外放协议查询
// alibaba.alicom.vt.distribute.queryprotocol
//
// 通信业务外放协议查询
func AlibabaAlicomVtDistributeQueryprotocol(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAlicomVtDistributeQueryprotocolAPIRequest, resp *alicom.AlibabaAlicomVtDistributeQueryprotocolAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
