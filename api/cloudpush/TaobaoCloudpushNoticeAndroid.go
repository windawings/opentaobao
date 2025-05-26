package cloudpush

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/cloudpush"
)

// TaobaoCloudpushNoticeAndroid 百川云推送发送通知给android
// taobao.cloudpush.notice.android
//
// 百川云推送发送通知给android
func TaobaoCloudpushNoticeAndroid(ctx context.Context, clt *core.SDKClient, req *cloudpush.TaobaoCloudpushNoticeAndroidAPIRequest, resp *cloudpush.TaobaoCloudpushNoticeAndroidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
