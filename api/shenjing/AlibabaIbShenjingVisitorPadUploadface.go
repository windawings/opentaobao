package shenjing

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/shenjing"
)

// AlibabaIbShenjingVisitorPadUploadface 访客PAD上传人脸
// alibaba.ib.shenjing.visitor.pad.uploadface
//
// 访客PAD端上传人脸。
func AlibabaIbShenjingVisitorPadUploadface(ctx context.Context, clt *core.SDKClient, req *shenjing.AlibabaIbShenjingVisitorPadUploadfaceAPIRequest, resp *shenjing.AlibabaIbShenjingVisitorPadUploadfaceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
