package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusVisitorImageUpload 访客大厅图片上传及查看
// alibaba.campus.visitor.image.upload
//
// 访客大厅图片上传及查看
func AlibabaCampusVisitorImageUpload(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusVisitorImageUploadAPIRequest, resp *campus.AlibabaCampusVisitorImageUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
