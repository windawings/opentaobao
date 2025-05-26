package nrt

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/nrt"
)

// TmallNrtCoupontemplateQuery 券模板查询
// tmall.nrt.coupontemplate.query
//
// 新零售场景，商家拉取在新零售工作台设置的券数据
func TmallNrtCoupontemplateQuery(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtCoupontemplateQueryAPIRequest, resp *nrt.TmallNrtCoupontemplateQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
