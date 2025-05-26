package tmallservice

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardDesensitizationQuery 工单查询接口
// tmall.servicecenter.workcard.desensitization.query
//
// 工单查询接口
func TmallServicecenterWorkcardDesensitizationQuery(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardDesensitizationQueryAPIRequest, resp *tmallservice.TmallServicecenterWorkcardDesensitizationQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
