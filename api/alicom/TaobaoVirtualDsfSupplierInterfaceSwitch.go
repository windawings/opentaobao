package alicom

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alicom"
)

// TaobaoVirtualDsfSupplierInterfaceSwitch 虚拟供应商履约接口切换
// taobao.virtual.dsf.supplier.interface.switch
//
// 虚拟供应商履约接口切换
func TaobaoVirtualDsfSupplierInterfaceSwitch(ctx context.Context, clt *core.SDKClient, req *alicom.TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest, resp *alicom.TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
