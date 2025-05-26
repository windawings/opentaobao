package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripProjectAdd 添加项目
// alitrip.btrip.project.add
//
// 添加项目
func AlitripBtripProjectAdd(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripProjectAddAPIRequest, resp *btrip.AlitripBtripProjectAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
