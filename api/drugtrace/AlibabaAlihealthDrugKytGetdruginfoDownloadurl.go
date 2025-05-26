package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytGetdruginfoDownloadurl 药品全量数据下载
// alibaba.alihealth.drug.kyt.getdruginfo.downloadurl
//
// 药品全量数据下载
func AlibabaAlihealthDrugKytGetdruginfoDownloadurl(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
