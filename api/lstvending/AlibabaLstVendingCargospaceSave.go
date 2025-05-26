package lstvending

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/lstvending"
)

// AlibabaLstVendingCargospaceSave 自动售卖机货道数据回流
// alibaba.lst.vending.cargospace.save
//
// 自动售卖机货道数据回流接口，ISV通过调用此接口上传售卖机货道信息。
func AlibabaLstVendingCargospaceSave(ctx context.Context, clt *core.SDKClient, req *lstvending.AlibabaLstVendingCargospaceSaveAPIRequest, resp *lstvending.AlibabaLstVendingCargospaceSaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
