package vaccin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/vaccin"
)

// AlibabaAlihealthVaccineRegisterSubmit cdc回传疫苗登记数据
// alibaba.alihealth.vaccine.register.submit
//
// cdc回传疫苗登记信息
func AlibabaAlihealthVaccineRegisterSubmit(ctx context.Context, clt *core.SDKClient, req *vaccin.AlibabaAlihealthVaccineRegisterSubmitAPIRequest, resp *vaccin.AlibabaAlihealthVaccineRegisterSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
