// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets all legal agreements that user needs to accept before purchasing a domain.
func ListTopLevelDomainAgreements(ctx *pulumi.Context, args *ListTopLevelDomainAgreementsArgs, opts ...pulumi.InvokeOption) (*ListTopLevelDomainAgreementsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListTopLevelDomainAgreementsResult
	err := ctx.Invoke("azure-native:domainregistration/v20210201:listTopLevelDomainAgreements", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTopLevelDomainAgreementsArgs struct {
	// If <code>true</code>, then the list of agreements will include agreements for domain transfer as well; otherwise, <code>false</code>.
	ForTransfer *bool `pulumi:"forTransfer"`
	// If <code>true</code>, then the list of agreements will include agreements for domain privacy as well; otherwise, <code>false</code>.
	IncludePrivacy *bool `pulumi:"includePrivacy"`
	// Name of the top-level domain.
	Name string `pulumi:"name"`
}

// Collection of top-level domain legal agreements.
type ListTopLevelDomainAgreementsResult struct {
	// Link to next page of resources.
	NextLink string `pulumi:"nextLink"`
	// Collection of resources.
	Value []TldLegalAgreementResponse `pulumi:"value"`
}

func ListTopLevelDomainAgreementsOutput(ctx *pulumi.Context, args ListTopLevelDomainAgreementsOutputArgs, opts ...pulumi.InvokeOption) ListTopLevelDomainAgreementsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListTopLevelDomainAgreementsResult, error) {
			args := v.(ListTopLevelDomainAgreementsArgs)
			r, err := ListTopLevelDomainAgreements(ctx, &args, opts...)
			var s ListTopLevelDomainAgreementsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListTopLevelDomainAgreementsResultOutput)
}

type ListTopLevelDomainAgreementsOutputArgs struct {
	// If <code>true</code>, then the list of agreements will include agreements for domain transfer as well; otherwise, <code>false</code>.
	ForTransfer pulumi.BoolPtrInput `pulumi:"forTransfer"`
	// If <code>true</code>, then the list of agreements will include agreements for domain privacy as well; otherwise, <code>false</code>.
	IncludePrivacy pulumi.BoolPtrInput `pulumi:"includePrivacy"`
	// Name of the top-level domain.
	Name pulumi.StringInput `pulumi:"name"`
}

func (ListTopLevelDomainAgreementsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTopLevelDomainAgreementsArgs)(nil)).Elem()
}

// Collection of top-level domain legal agreements.
type ListTopLevelDomainAgreementsResultOutput struct{ *pulumi.OutputState }

func (ListTopLevelDomainAgreementsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTopLevelDomainAgreementsResult)(nil)).Elem()
}

func (o ListTopLevelDomainAgreementsResultOutput) ToListTopLevelDomainAgreementsResultOutput() ListTopLevelDomainAgreementsResultOutput {
	return o
}

func (o ListTopLevelDomainAgreementsResultOutput) ToListTopLevelDomainAgreementsResultOutputWithContext(ctx context.Context) ListTopLevelDomainAgreementsResultOutput {
	return o
}

// Link to next page of resources.
func (o ListTopLevelDomainAgreementsResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListTopLevelDomainAgreementsResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// Collection of resources.
func (o ListTopLevelDomainAgreementsResultOutput) Value() TldLegalAgreementResponseArrayOutput {
	return o.ApplyT(func(v ListTopLevelDomainAgreementsResult) []TldLegalAgreementResponse { return v.Value }).(TldLegalAgreementResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListTopLevelDomainAgreementsResultOutput{})
}
