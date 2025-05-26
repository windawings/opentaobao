package alicom

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alicom"
)

// AlibabaWttUserRegioninfoByipGet 根据ip获取省市信息
// alibaba.wtt.user.regioninfo.byip.get
//
// 通过ip获取省市信息
func AlibabaWttUserRegioninfoByipGet(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaWttUserRegioninfoByipGetAPIRequest, resp *alicom.AlibabaWttUserRegioninfoByipGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
