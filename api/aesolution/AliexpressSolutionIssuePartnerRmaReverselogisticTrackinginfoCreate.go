package aesolution

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/aesolution"
)

// AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreate aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create
// aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create
//
// Receives information about reverse logistics tracking info
func AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreate(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIRequest, resp *aesolution.AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
