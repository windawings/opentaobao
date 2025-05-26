package normalvisa

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/normalvisa"
)

// TaobaoAlitripTravelNormalvisaUploadfile 上传电子签证
// taobao.alitrip.travel.normalvisa.uploadfile
//
// 上传电子签证
func TaobaoAlitripTravelNormalvisaUploadfile(ctx context.Context, clt *core.SDKClient, req *normalvisa.TaobaoAlitripTravelNormalvisaUploadfileAPIRequest, resp *normalvisa.TaobaoAlitripTravelNormalvisaUploadfileAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
