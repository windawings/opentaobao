package baoxian

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/baoxian"
)

// AlipayBaoxianClaimUploadattachment 资料上传接口
// alipay.baoxian.claim.uploadattachment
//
// 给合作伙伴上传申请理赔材料
func AlipayBaoxianClaimUploadattachment(ctx context.Context, clt *core.SDKClient, req *baoxian.AlipayBaoxianClaimUploadattachmentAPIRequest, resp *baoxian.AlipayBaoxianClaimUploadattachmentAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
