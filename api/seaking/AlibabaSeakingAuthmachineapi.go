package seaking

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/seaking"
)

// AlibabaSeakingAuthmachineapi 机翻Api授权
// alibaba.seaking.authmachineapi
//
// 机翻Api授权
func AlibabaSeakingAuthmachineapi(ctx context.Context, clt *core.SDKClient, req *seaking.AlibabaSeakingAuthmachineapiAPIRequest, resp *seaking.AlibabaSeakingAuthmachineapiAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
