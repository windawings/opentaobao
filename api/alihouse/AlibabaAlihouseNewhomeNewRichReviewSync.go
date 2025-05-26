package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeNewRichReviewSync 新评测改造接口同步
// alibaba.alihouse.newhome.new.rich.review.sync
//
// 新评测改造接口同步
func AlibabaAlihouseNewhomeNewRichReviewSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeNewRichReviewSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
