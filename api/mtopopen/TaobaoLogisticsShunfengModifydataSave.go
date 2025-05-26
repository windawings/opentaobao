package mtopopen

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/mtopopen"
)

// TaobaoLogisticsShunfengModifydataSave 顺丰小程序修改配送信息回传接口
// taobao.logistics.shunfeng.modifydata.save
//
// 顺丰小程序修改配送信息回传接口
func TaobaoLogisticsShunfengModifydataSave(ctx context.Context, clt *core.SDKClient, req *mtopopen.TaobaoLogisticsShunfengModifydataSaveAPIRequest, resp *mtopopen.TaobaoLogisticsShunfengModifydataSaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
