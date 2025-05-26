package normalvisa

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/normalvisa"
)

// TaobaoAlitripTravelNormalvisaUpdatepersonstauts 更新签证办理进度
// taobao.alitrip.travel.normalvisa.updatepersonstauts
//
// 更新签证办理进度
func TaobaoAlitripTravelNormalvisaUpdatepersonstauts(ctx context.Context, clt *core.SDKClient, req *normalvisa.TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest, resp *normalvisa.TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
