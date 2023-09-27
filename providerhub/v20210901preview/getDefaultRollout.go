// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the default rollout details.
func LookupDefaultRollout(ctx *pulumi.Context, args *LookupDefaultRolloutArgs, opts ...pulumi.InvokeOption) (*LookupDefaultRolloutResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDefaultRolloutResult
	err := ctx.Invoke("azure-native:providerhub/v20210901preview:getDefaultRollout", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDefaultRolloutArgs struct {
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace string `pulumi:"providerNamespace"`
	// The rollout name.
	RolloutName string `pulumi:"rolloutName"`
}

// Default rollout definition.
type LookupDefaultRolloutResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties of the rollout.
	Properties DefaultRolloutResponseProperties `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDefaultRolloutOutput(ctx *pulumi.Context, args LookupDefaultRolloutOutputArgs, opts ...pulumi.InvokeOption) LookupDefaultRolloutResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDefaultRolloutResult, error) {
			args := v.(LookupDefaultRolloutArgs)
			r, err := LookupDefaultRollout(ctx, &args, opts...)
			var s LookupDefaultRolloutResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDefaultRolloutResultOutput)
}

type LookupDefaultRolloutOutputArgs struct {
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace pulumi.StringInput `pulumi:"providerNamespace"`
	// The rollout name.
	RolloutName pulumi.StringInput `pulumi:"rolloutName"`
}

func (LookupDefaultRolloutOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultRolloutArgs)(nil)).Elem()
}

// Default rollout definition.
type LookupDefaultRolloutResultOutput struct{ *pulumi.OutputState }

func (LookupDefaultRolloutResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultRolloutResult)(nil)).Elem()
}

func (o LookupDefaultRolloutResultOutput) ToLookupDefaultRolloutResultOutput() LookupDefaultRolloutResultOutput {
	return o
}

func (o LookupDefaultRolloutResultOutput) ToLookupDefaultRolloutResultOutputWithContext(ctx context.Context) LookupDefaultRolloutResultOutput {
	return o
}

func (o LookupDefaultRolloutResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDefaultRolloutResult] {
	return pulumix.Output[LookupDefaultRolloutResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDefaultRolloutResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDefaultRolloutResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the rollout.
func (o LookupDefaultRolloutResultOutput) Properties() DefaultRolloutResponsePropertiesOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) DefaultRolloutResponseProperties { return v.Properties }).(DefaultRolloutResponsePropertiesOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupDefaultRolloutResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDefaultRolloutResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDefaultRolloutResultOutput{})
}
