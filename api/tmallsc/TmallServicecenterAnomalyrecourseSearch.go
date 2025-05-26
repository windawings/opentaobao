package tmallsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseSearch 天猫服务平台服务商一键求助单查询
// tmall.servicecenter.anomalyrecourse.search
//
// 天猫服务平台服务商一键求助单查询
func TmallServicecenterAnomalyrecourseSearch(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseSearchAPIRequest, resp *tmallsc.TmallServicecenterAnomalyrecourseSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
