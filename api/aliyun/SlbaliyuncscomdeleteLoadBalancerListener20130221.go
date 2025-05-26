package aliyun

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/aliyun"
)

// SlbAliyuncsComDeleteLoadBalancerListener20130221 删除 slb listener
// slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21
//
// delete_vip
func SlbAliyuncsComDeleteLoadBalancerListener20130221(ctx context.Context, clt *core.SDKClient, req *aliyun.SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest, resp *aliyun.SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
