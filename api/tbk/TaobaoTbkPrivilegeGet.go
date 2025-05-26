package tbk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbk"
)

// TaobaoTbkPrivilegeGet 淘宝客-服务商-单品券高效转链
// taobao.tbk.privilege.get
//
// 单品券高效转链API
func TaobaoTbkPrivilegeGet(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkPrivilegeGetAPIRequest, resp *tbk.TaobaoTbkPrivilegeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
