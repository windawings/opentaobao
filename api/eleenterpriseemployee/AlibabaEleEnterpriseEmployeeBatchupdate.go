package eleenterpriseemployee

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/eleenterpriseemployee"
)

// AlibabaEleEnterpriseEmployeeBatchupdate 批量新增更新员工
// alibaba.ele.enterprise.employee.batchupdate
//
// 批量新增更新员工
func AlibabaEleEnterpriseEmployeeBatchupdate(ctx context.Context, clt *core.SDKClient, req *eleenterpriseemployee.AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest, resp *eleenterpriseemployee.AlibabaEleEnterpriseEmployeeBatchupdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
