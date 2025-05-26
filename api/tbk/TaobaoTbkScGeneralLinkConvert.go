package tbk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbk"
)

// TaobaoTbkScGeneralLinkConvert 淘宝客-服务商-万能转链
// taobao.tbk.sc.general.link.convert
//
// 淘宝客-服务商-万能转链
func TaobaoTbkScGeneralLinkConvert(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkScGeneralLinkConvertAPIRequest, resp *tbk.TaobaoTbkScGeneralLinkConvertAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
