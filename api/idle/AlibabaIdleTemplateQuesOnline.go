package idle

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/idle"
)

// AlibabaIdleTemplateQuesOnline 预上线SPU问卷版本
// alibaba.idle.template.ques.online
//
// 获取SPU最新版本问卷
func AlibabaIdleTemplateQuesOnline(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleTemplateQuesOnlineAPIRequest, resp *idle.AlibabaIdleTemplateQuesOnlineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
