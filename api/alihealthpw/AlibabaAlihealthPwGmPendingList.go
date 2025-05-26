package alihealthpw

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwGmPendingList 同情用药待审核工单查询接口
// alibaba.alihealth.pw.gm.pending.list
//
// 同情用药待审核工单查询接口，提供给合作方用来查询待处理工单列表
func AlibabaAlihealthPwGmPendingList(ctx context.Context, clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwGmPendingListAPIRequest, resp *alihealthpw.AlibabaAlihealthPwGmPendingListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
