package tmallnr

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmActivityStoreGetstorelist ISV查询门店
// alibaba.lsy.crm.activity.store.getstorelist
//
// ISV查询门店
func AlibabaLsyCrmActivityStoreGetstorelist(ctx context.Context, clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivityStoreGetstorelistAPIRequest, resp *tmallnr.AlibabaLsyCrmActivityStoreGetstorelistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
