package lstspeacker

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/lstspeacker"
)

// AlibabaLstSpeakerStatusGet 音箱设备在线状态
// alibaba.lst.speaker.status.get
//
// 音箱设备在线状态查询
func AlibabaLstSpeakerStatusGet(ctx context.Context, clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerStatusGetAPIRequest, resp *lstspeacker.AlibabaLstSpeakerStatusGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
