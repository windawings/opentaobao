package wangwang

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wangwang"
)

// TaobaoQianniuKefuevalGet 客服评价详情接口
// taobao.qianniu.kefueval.get
//
// 获取买家对客服的服务评价
func TaobaoQianniuKefuevalGet(ctx context.Context, clt *core.SDKClient, req *wangwang.TaobaoQianniuKefuevalGetAPIRequest, resp *wangwang.TaobaoQianniuKefuevalGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
