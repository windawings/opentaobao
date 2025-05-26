package tttm

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tttm"
)

// AliyunIndustryTttmProduceSync 天天特卖生产进度同步
// aliyun.industry.tttm.produce.sync
//
// 天天特卖生产进度同步
func AliyunIndustryTttmProduceSync(ctx context.Context, clt *core.SDKClient, req *tttm.AliyunIndustryTttmProduceSyncAPIRequest, resp *tttm.AliyunIndustryTttmProduceSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
