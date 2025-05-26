package film

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/film"
)

// TaobaoFilmDataThirdPartyRefundOrder 退票接口
// taobao.film.data.third.party.refund.order
//
// 淘票票第三方退票接口
func TaobaoFilmDataThirdPartyRefundOrder(ctx context.Context, clt *core.SDKClient, req *film.TaobaoFilmDataThirdPartyRefundOrderAPIRequest, resp *film.TaobaoFilmDataThirdPartyRefundOrderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
