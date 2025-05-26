package miniappcloud

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/miniappcloud"
)

// TaobaoMiniappCloudMongoUpdate 更新MongoDB中的数据
// taobao.miniapp.cloud.mongo.update
//
// 更新MongoDB中的数据
func TaobaoMiniappCloudMongoUpdate(ctx context.Context, clt *core.SDKClient, req *miniappcloud.TaobaoMiniappCloudMongoUpdateAPIRequest, resp *miniappcloud.TaobaoMiniappCloudMongoUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
