package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripProjectDelete 删除项目
// alitrip.btrip.project.delete
//
// 删除项目
func AlitripBtripProjectDelete(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripProjectDeleteAPIRequest, resp *btrip.AlitripBtripProjectDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
