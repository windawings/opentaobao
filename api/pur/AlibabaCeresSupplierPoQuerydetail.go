package pur

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/pur"
)

// AlibabaCeresSupplierPoQuerydetail 采购供应商订单明细查询接口
// alibaba.ceres.supplier.po.querydetail
//
// 采购供应商订单明细查询接口
func AlibabaCeresSupplierPoQuerydetail(ctx context.Context, clt *core.SDKClient, req *pur.AlibabaCeresSupplierPoQuerydetailAPIRequest, resp *pur.AlibabaCeresSupplierPoQuerydetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
