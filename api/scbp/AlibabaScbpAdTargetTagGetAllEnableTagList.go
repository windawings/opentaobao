package scbp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/scbp"
)

// AlibabaScbpAdTargetTagGetAllEnableTagList 查询所有可添加标签信息
// alibaba.scbp.ad.target.tag.get.all.enable.tag.list
//
// 查询标签数据
func AlibabaScbpAdTargetTagGetAllEnableTagList(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest, resp *scbp.AlibabaScbpAdTargetTagGetAllEnableTagListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
