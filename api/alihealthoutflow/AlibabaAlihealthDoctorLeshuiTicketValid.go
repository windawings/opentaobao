package alihealthoutflow

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthDoctorLeshuiTicketValid 乐税token验证
// alibaba.alihealth.doctor.leshui.ticket.valid
//
// 乐税token验证
func AlibabaAlihealthDoctorLeshuiTicketValid(ctx context.Context, clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthDoctorLeshuiTicketValidAPIRequest, resp *alihealthoutflow.AlibabaAlihealthDoctorLeshuiTicketValidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
