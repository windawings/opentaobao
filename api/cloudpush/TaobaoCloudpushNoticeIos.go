package cloudpush

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/cloudpush"
)

// TaobaoCloudpushNoticeIos 推送通知给ios设备
// taobao.cloudpush.notice.ios
//
// 推送通知给ios设备
func TaobaoCloudpushNoticeIos(ctx context.Context, clt *core.SDKClient, req *cloudpush.TaobaoCloudpushNoticeIosAPIRequest, resp *cloudpush.TaobaoCloudpushNoticeIosAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
