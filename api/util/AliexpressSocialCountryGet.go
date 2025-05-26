package util

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/util"
)

// AliexpressSocialCountryGet 获取国家列表
// aliexpress.social.country.get
//
// 获取目前AE支持的国家列表
func AliexpressSocialCountryGet(ctx context.Context, clt *core.SDKClient, req *util.AliexpressSocialCountryGetAPIRequest, resp *util.AliexpressSocialCountryGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
