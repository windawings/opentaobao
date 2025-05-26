package aesolution

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/aesolution"
)

// AliexpressSolutionIssuePartnerRmaScreeningCreate aliexpress.solution.issue.partner.rma.screening.create
// aliexpress.solution.issue.partner.rma.screening.create
//
// Receives information about screening results from after sales partners
func AliexpressSolutionIssuePartnerRmaScreeningCreate(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionIssuePartnerRmaScreeningCreateAPIRequest, resp *aesolution.AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
