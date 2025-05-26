package alihealthcrm

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthUicUserinfoHealthidGet 获取健康id
// alibaba.alihealth.uic.userinfo.healthid.get
//
// 根据支付宝用户ID获取用户健康ID
func AlibabaAlihealthUicUserinfoHealthidGet(ctx context.Context, clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthUicUserinfoHealthidGetAPIRequest, resp *alihealthcrm.AlibabaAlihealthUicUserinfoHealthidGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
