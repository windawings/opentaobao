package aliyun

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/aliyun"
)

// AccountAliyuncsComListAppkeyByOwnerAndBid20130701 根据渠道id和状态列出appkey
// account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01
//
// 根据渠道id和状态列出appkey
func AccountAliyuncsComListAppkeyByOwnerAndBid20130701(ctx context.Context, clt *core.SDKClient, req *aliyun.AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIRequest, resp *aliyun.AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
