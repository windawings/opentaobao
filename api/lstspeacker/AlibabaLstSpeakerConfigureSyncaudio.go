package lstspeacker

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/lstspeacker"
)

// AlibabaLstSpeakerConfigureSyncaudio 音频同步
// alibaba.lst.speaker.configure.syncaudio
//
// 音频同步
func AlibabaLstSpeakerConfigureSyncaudio(ctx context.Context, clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerConfigureSyncaudioAPIRequest, resp *lstspeacker.AlibabaLstSpeakerConfigureSyncaudioAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
