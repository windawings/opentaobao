package alicom

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alicom"
)

// AlibabaAliqinFcVoiceGetdetail 获取呼叫详情
// alibaba.aliqin.fc.voice.getdetail
//
// 通过呼叫id获取呼叫相关的数据
func AlibabaAliqinFcVoiceGetdetail(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAliqinFcVoiceGetdetailAPIRequest, resp *alicom.AlibabaAliqinFcVoiceGetdetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
