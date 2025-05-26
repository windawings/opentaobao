package alimember

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alimember"
)

// AlibabaMemberIsvPageQuery isv离线会员数据分页查询
// alibaba.member.isv.page.query
//
// isv离线会员数据分页查询
func AlibabaMemberIsvPageQuery(ctx context.Context, clt *core.SDKClient, req *alimember.AlibabaMemberIsvPageQueryAPIRequest, resp *alimember.AlibabaMemberIsvPageQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
