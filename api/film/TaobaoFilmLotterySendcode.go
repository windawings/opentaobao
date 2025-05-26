package film

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/film"
)

// TaobaoFilmLotterySendcode 淘票票外部直发券
// taobao.film.lottery.sendcode
//
// 淘票票外部直发券
func TaobaoFilmLotterySendcode(ctx context.Context, clt *core.SDKClient, req *film.TaobaoFilmLotterySendcodeAPIRequest, resp *film.TaobaoFilmLotterySendcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
