package alitripmerchant

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberCompleteSwitch 会员切换模式接口
// alitrip.merchant.galaxy.member.complete.switch
//
// 小程序老用户调用德比接口进行会员切换
func AlitripMerchantGalaxyMemberCompleteSwitch(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMemberCompleteSwitchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
