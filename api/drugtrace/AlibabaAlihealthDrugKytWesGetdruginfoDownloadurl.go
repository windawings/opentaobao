package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesGetdruginfoDownloadurl 药品全量数据下载
// alibaba.alihealth.drug.kyt.wes.getdruginfo.downloadurl
//
// 药品全量数据下载
func AlibabaAlihealthDrugKytWesGetdruginfoDownloadurl(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
