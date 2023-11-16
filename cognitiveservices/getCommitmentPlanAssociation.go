// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognitiveservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the association of the Cognitive Services commitment plan.
// Azure REST API version: 2023-05-01.
//
// Other available API versions: 2023-10-01-preview.
func LookupCommitmentPlanAssociation(ctx *pulumi.Context, args *LookupCommitmentPlanAssociationArgs, opts ...pulumi.InvokeOption) (*LookupCommitmentPlanAssociationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCommitmentPlanAssociationResult
	err := ctx.Invoke("azure-native:cognitiveservices:getCommitmentPlanAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCommitmentPlanAssociationArgs struct {
	// The name of the commitment plan association with the Cognitive Services Account
	CommitmentPlanAssociationName string `pulumi:"commitmentPlanAssociationName"`
	// The name of the commitmentPlan associated with the Cognitive Services Account
	CommitmentPlanName string `pulumi:"commitmentPlanName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The commitment plan association.
type LookupCommitmentPlanAssociationResult struct {
	// The Azure resource id of the account.
	AccountId *string `pulumi:"accountId"`
	// Resource Etag.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupCommitmentPlanAssociationOutput(ctx *pulumi.Context, args LookupCommitmentPlanAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupCommitmentPlanAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCommitmentPlanAssociationResult, error) {
			args := v.(LookupCommitmentPlanAssociationArgs)
			r, err := LookupCommitmentPlanAssociation(ctx, &args, opts...)
			var s LookupCommitmentPlanAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCommitmentPlanAssociationResultOutput)
}

type LookupCommitmentPlanAssociationOutputArgs struct {
	// The name of the commitment plan association with the Cognitive Services Account
	CommitmentPlanAssociationName pulumi.StringInput `pulumi:"commitmentPlanAssociationName"`
	// The name of the commitmentPlan associated with the Cognitive Services Account
	CommitmentPlanName pulumi.StringInput `pulumi:"commitmentPlanName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCommitmentPlanAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommitmentPlanAssociationArgs)(nil)).Elem()
}

// The commitment plan association.
type LookupCommitmentPlanAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupCommitmentPlanAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommitmentPlanAssociationResult)(nil)).Elem()
}

func (o LookupCommitmentPlanAssociationResultOutput) ToLookupCommitmentPlanAssociationResultOutput() LookupCommitmentPlanAssociationResultOutput {
	return o
}

func (o LookupCommitmentPlanAssociationResultOutput) ToLookupCommitmentPlanAssociationResultOutputWithContext(ctx context.Context) LookupCommitmentPlanAssociationResultOutput {
	return o
}

// The Azure resource id of the account.
func (o LookupCommitmentPlanAssociationResultOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCommitmentPlanAssociationResult) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

// Resource Etag.
func (o LookupCommitmentPlanAssociationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanAssociationResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupCommitmentPlanAssociationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanAssociationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCommitmentPlanAssociationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanAssociationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupCommitmentPlanAssociationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCommitmentPlanAssociationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCommitmentPlanAssociationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanAssociationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCommitmentPlanAssociationResultOutput{})
}
