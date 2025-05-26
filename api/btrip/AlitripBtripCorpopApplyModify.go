package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripCorpopApplyModify 【商旅】修改出差审批单（行程）
// alitrip.btrip.corpop.apply.modify
//
// 【商旅】修改出差审批单（行程）
func AlitripBtripCorpopApplyModify(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCorpopApplyModifyAPIRequest, resp *btrip.AlitripBtripCorpopApplyModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
