package drug

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drug"
)

// AlibabaAlihealthNrTradeMedicalInsuranceGet 阿里健康医保支付信息获取
// alibaba.alihealth.nr.trade.medical.insurance.get
//
// 阿里健康医保支付信息获取
func AlibabaAlihealthNrTradeMedicalInsuranceGet(ctx context.Context, clt *core.SDKClient, req *drug.AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest, resp *drug.AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
