package iot

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/iot"
)

// AlibabaAilabsIotBusinessRecipestepInsertorupdate 插入或更新食谱步骤
// alibaba.ailabs.iot.business.recipestep.insertorupdate
//
// 插入或更新食谱步骤
func AlibabaAilabsIotBusinessRecipestepInsertorupdate(ctx context.Context, clt *core.SDKClient, req *iot.AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest, resp *iot.AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
