package alihealth2

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihealth2"
)

// AlibabaAlihealthDentalItemList ISV获取口腔标品列表
// alibaba.alihealth.dental.item.list
//
// ISV获取口腔标品列表
func AlibabaAlihealthDentalItemList(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalItemListAPIRequest, resp *alihealth2.AlibabaAlihealthDentalItemListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
