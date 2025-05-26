package qt

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/qt"
)

// TaobaoTsPropertyGet 淘宝服务属性查询
// taobao.ts.property.get
//
// 淘宝服务属性查询
func TaobaoTsPropertyGet(ctx context.Context, clt *core.SDKClient, req *qt.TaobaoTsPropertyGetAPIRequest, resp *qt.TaobaoTsPropertyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
