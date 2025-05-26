package dt

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/dt"
)

// AlibabaNrsItemPricetagRecognize 价签识别
// alibaba.nrs.item.pricetag.recognize
//
// 商品价签识别，用于识别RT上传的竞品分析照片，返回价签内容
func AlibabaNrsItemPricetagRecognize(ctx context.Context, clt *core.SDKClient, req *dt.AlibabaNrsItemPricetagRecognizeAPIRequest, resp *dt.AlibabaNrsItemPricetagRecognizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
