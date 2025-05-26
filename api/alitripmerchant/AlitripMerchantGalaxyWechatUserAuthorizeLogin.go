package alitripmerchant

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatUserAuthorizeLogin DFC-ID用户手机号授权登录
// alitrip.merchant.galaxy.wechat.user.authorize.login
//
// DFC-ID用户手机号授权登录
func AlitripMerchantGalaxyWechatUserAuthorizeLogin(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
