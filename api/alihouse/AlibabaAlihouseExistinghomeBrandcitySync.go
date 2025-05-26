package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrandcitySync 二手房城市品牌店数据同步
// alibaba.alihouse.existinghome.brandcity.sync
//
// 二手房城市品牌店数据同步
func AlibabaAlihouseExistinghomeBrandcitySync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeBrandcitySyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
