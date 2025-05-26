package ju

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/ju"
)

// AlibabaJhsCommunityIdentityStore 用户信息存储
// alibaba.jhs.community.identity.store
//
// 用户信息存储
func AlibabaJhsCommunityIdentityStore(ctx context.Context, clt *core.SDKClient, req *ju.AlibabaJhsCommunityIdentityStoreAPIRequest, resp *ju.AlibabaJhsCommunityIdentityStoreAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
