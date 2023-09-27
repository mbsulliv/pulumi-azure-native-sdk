// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Obtains the details of a suppression.
func LookupSuppression(ctx *pulumi.Context, args *LookupSuppressionArgs, opts ...pulumi.InvokeOption) (*LookupSuppressionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSuppressionResult
	err := ctx.Invoke("azure-native:advisor/v20230101:getSuppression", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSuppressionArgs struct {
	// The name of the suppression.
	Name string `pulumi:"name"`
	// The recommendation ID.
	RecommendationId string `pulumi:"recommendationId"`
	// The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
	ResourceUri string `pulumi:"resourceUri"`
}

// The details of the snoozed or dismissed rule; for example, the duration, name, and GUID associated with the rule.
type LookupSuppressionResult struct {
	// Gets or sets the expiration time stamp.
	ExpirationTimeStamp string `pulumi:"expirationTimeStamp"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The GUID of the suppression.
	SuppressionId *string `pulumi:"suppressionId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The duration for which the suppression is valid.
	Ttl *string `pulumi:"ttl"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSuppressionOutput(ctx *pulumi.Context, args LookupSuppressionOutputArgs, opts ...pulumi.InvokeOption) LookupSuppressionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSuppressionResult, error) {
			args := v.(LookupSuppressionArgs)
			r, err := LookupSuppression(ctx, &args, opts...)
			var s LookupSuppressionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSuppressionResultOutput)
}

type LookupSuppressionOutputArgs struct {
	// The name of the suppression.
	Name pulumi.StringInput `pulumi:"name"`
	// The recommendation ID.
	RecommendationId pulumi.StringInput `pulumi:"recommendationId"`
	// The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupSuppressionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSuppressionArgs)(nil)).Elem()
}

// The details of the snoozed or dismissed rule; for example, the duration, name, and GUID associated with the rule.
type LookupSuppressionResultOutput struct{ *pulumi.OutputState }

func (LookupSuppressionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSuppressionResult)(nil)).Elem()
}

func (o LookupSuppressionResultOutput) ToLookupSuppressionResultOutput() LookupSuppressionResultOutput {
	return o
}

func (o LookupSuppressionResultOutput) ToLookupSuppressionResultOutputWithContext(ctx context.Context) LookupSuppressionResultOutput {
	return o
}

func (o LookupSuppressionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSuppressionResult] {
	return pulumix.Output[LookupSuppressionResult]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the expiration time stamp.
func (o LookupSuppressionResultOutput) ExpirationTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionResult) string { return v.ExpirationTimeStamp }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupSuppressionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSuppressionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The GUID of the suppression.
func (o LookupSuppressionResultOutput) SuppressionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSuppressionResult) *string { return v.SuppressionId }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSuppressionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSuppressionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The duration for which the suppression is valid.
func (o LookupSuppressionResultOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSuppressionResult) *string { return v.Ttl }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSuppressionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSuppressionResultOutput{})
}
