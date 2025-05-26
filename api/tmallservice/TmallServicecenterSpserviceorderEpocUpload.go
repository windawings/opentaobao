package tmallservice

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallservice"
)

// TmallServicecenterSpserviceorderEpocUpload 电子保单文件上传接口
// tmall.servicecenter.spserviceorder.epoc.upload
//
// 电子保单文件上传接口
func TmallServicecenterSpserviceorderEpocUpload(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterSpserviceorderEpocUploadAPIRequest, resp *tmallservice.TmallServicecenterSpserviceorderEpocUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
