package idle

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/idle"
)

// AlibabaIdleConsignmentSpuStatistics 闲鱼帮卖同步服务商交易统计信息
// alibaba.idle.consignment.spu.statistics
//
// 闲鱼帮卖同步服务商交易统计信息
func AlibabaIdleConsignmentSpuStatistics(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleConsignmentSpuStatisticsAPIRequest, resp *idle.AlibabaIdleConsignmentSpuStatisticsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
