package alisports

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportOauthAlipaygrant 阿里体育会员系统-支付宝授权接口
// alibaba.alisports.passport.oauth.alipaygrant
//
// 开放给乐心运动使用的支付宝授权接口
func AlibabaAlisportsPassportOauthAlipaygrant(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportOauthAlipaygrantAPIRequest, resp *alisports.AlibabaAlisportsPassportOauthAlipaygrantAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
