package sungari

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/sungari"
)

// TaobaoCloudbridgeCaseinvestExecute 红盾云桥案件协查服务
// taobao.cloudbridge.caseinvest.execute
//
// 通过API接口直接提供政府部门录入及查询函件服务
func TaobaoCloudbridgeCaseinvestExecute(ctx context.Context, clt *core.SDKClient, req *sungari.TaobaoCloudbridgeCaseinvestExecuteAPIRequest, resp *sungari.TaobaoCloudbridgeCaseinvestExecuteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
