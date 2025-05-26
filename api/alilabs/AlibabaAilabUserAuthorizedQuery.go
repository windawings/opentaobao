package alilabs

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alilabs"
)

// AlibabaAilabUserAuthorizedQuery 查询授权状态接口
// alibaba.ailab.user.authorized.query
//
// 查询三方用户授权状态
func AlibabaAilabUserAuthorizedQuery(ctx context.Context, clt *core.SDKClient, req *alilabs.AlibabaAilabUserAuthorizedQueryAPIRequest, resp *alilabs.AlibabaAilabUserAuthorizedQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
