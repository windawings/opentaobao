package vaccin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinNoticeTimebucketRemind 疫苗预约时间段提醒
// alibaba.health.vaccin.notice.timebucket.remind
//
// 疫苗预约时间段提醒
func AlibabaHealthVaccinNoticeTimebucketRemind(ctx context.Context, clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest, resp *vaccin.AlibabaHealthVaccinNoticeTimebucketRemindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
