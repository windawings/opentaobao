package alihealthcrm

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthAlipaypfmDietRecord 用户每日摄入卡路里总量回传接口
// alibaba.alihealth.alipaypfm.diet.record
//
// 用户每日摄入卡路里总量回传接口
func AlibabaAlihealthAlipaypfmDietRecord(ctx context.Context, clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthAlipaypfmDietRecordAPIRequest, resp *alihealthcrm.AlibabaAlihealthAlipaypfmDietRecordAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
