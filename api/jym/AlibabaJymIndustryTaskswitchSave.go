package jym

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/jym"
)

// AlibabaJymIndustryTaskswitchSave 行业信息系统开关
// alibaba.jym.industry.taskswitch.save
//
// VMOS回调交易猫行业信息系统
func AlibabaJymIndustryTaskswitchSave(ctx context.Context, clt *core.SDKClient, req *jym.AlibabaJymIndustryTaskswitchSaveAPIRequest, resp *jym.AlibabaJymIndustryTaskswitchSaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
