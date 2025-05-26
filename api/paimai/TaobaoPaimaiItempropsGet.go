package paimai

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/paimai"
)

// TaobaoPaimaiItempropsGet 拍卖相关类目属性
// taobao.paimai.itemprops.get
//
// 读取拍卖相关类目属性
func TaobaoPaimaiItempropsGet(ctx context.Context, clt *core.SDKClient, req *paimai.TaobaoPaimaiItempropsGetAPIRequest, resp *paimai.TaobaoPaimaiItempropsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
