package ju

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/ju"
)

// AlibabaJhsCommunityWechatLogin 聚划算用增淘外社群登录
// alibaba.jhs.community.wechat.login
//
// 聚划算用增淘外社群登录
func AlibabaJhsCommunityWechatLogin(ctx context.Context, clt *core.SDKClient, req *ju.AlibabaJhsCommunityWechatLoginAPIRequest, resp *ju.AlibabaJhsCommunityWechatLoginAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
