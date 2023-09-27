// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get Sentinel onboarding state
func LookupSentinelOnboardingState(ctx *pulumi.Context, args *LookupSentinelOnboardingStateArgs, opts ...pulumi.InvokeOption) (*LookupSentinelOnboardingStateResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSentinelOnboardingStateResult
	err := ctx.Invoke("azure-native:securityinsights/v20230201:getSentinelOnboardingState", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSentinelOnboardingStateArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Sentinel onboarding state name. Supports - default
	SentinelOnboardingStateName string `pulumi:"sentinelOnboardingStateName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Sentinel onboarding state
type LookupSentinelOnboardingStateResult struct {
	// Flag that indicates the status of the CMK setting
	CustomerManagedKey *bool `pulumi:"customerManagedKey"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSentinelOnboardingStateOutput(ctx *pulumi.Context, args LookupSentinelOnboardingStateOutputArgs, opts ...pulumi.InvokeOption) LookupSentinelOnboardingStateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSentinelOnboardingStateResult, error) {
			args := v.(LookupSentinelOnboardingStateArgs)
			r, err := LookupSentinelOnboardingState(ctx, &args, opts...)
			var s LookupSentinelOnboardingStateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSentinelOnboardingStateResultOutput)
}

type LookupSentinelOnboardingStateOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Sentinel onboarding state name. Supports - default
	SentinelOnboardingStateName pulumi.StringInput `pulumi:"sentinelOnboardingStateName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSentinelOnboardingStateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSentinelOnboardingStateArgs)(nil)).Elem()
}

// Sentinel onboarding state
type LookupSentinelOnboardingStateResultOutput struct{ *pulumi.OutputState }

func (LookupSentinelOnboardingStateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSentinelOnboardingStateResult)(nil)).Elem()
}

func (o LookupSentinelOnboardingStateResultOutput) ToLookupSentinelOnboardingStateResultOutput() LookupSentinelOnboardingStateResultOutput {
	return o
}

func (o LookupSentinelOnboardingStateResultOutput) ToLookupSentinelOnboardingStateResultOutputWithContext(ctx context.Context) LookupSentinelOnboardingStateResultOutput {
	return o
}

func (o LookupSentinelOnboardingStateResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSentinelOnboardingStateResult] {
	return pulumix.Output[LookupSentinelOnboardingStateResult]{
		OutputState: o.OutputState,
	}
}

// Flag that indicates the status of the CMK setting
func (o LookupSentinelOnboardingStateResultOutput) CustomerManagedKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSentinelOnboardingStateResult) *bool { return v.CustomerManagedKey }).(pulumi.BoolPtrOutput)
}

// Etag of the azure resource
func (o LookupSentinelOnboardingStateResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSentinelOnboardingStateResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSentinelOnboardingStateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentinelOnboardingStateResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSentinelOnboardingStateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentinelOnboardingStateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSentinelOnboardingStateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSentinelOnboardingStateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSentinelOnboardingStateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentinelOnboardingStateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSentinelOnboardingStateResultOutput{})
}
