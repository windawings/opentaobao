package baichuan

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/baichuan"
)

// AlibabaBaichuanAsoQuery 查询app在设备上的安装信息
// alibaba.baichuan.aso.query
//
// 查询app在设备上的安装信息
func AlibabaBaichuanAsoQuery(ctx context.Context, clt *core.SDKClient, req *baichuan.AlibabaBaichuanAsoQueryAPIRequest, resp *baichuan.AlibabaBaichuanAsoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
