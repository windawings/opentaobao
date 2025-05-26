package tmallgenie

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMemoMeetingDelete 天猫精灵会议删除
// taobao.ailab.aicloud.top.memo.meeting.delete
//
// 天猫精灵会议删除
func TaobaoAilabAicloudTopMemoMeetingDelete(ctx context.Context, clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopMemoMeetingDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
