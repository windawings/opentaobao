package alihealth2

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalbaseHosSyncnew 直连医院上传接口
// alibaba.alihealth.medicalbase.hos.syncnew
//
// 直连医院上传接口
func AlibabaAlihealthMedicalbaseHosSyncnew(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest, resp *alihealth2.AlibabaAlihealthMedicalbaseHosSyncnewAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
