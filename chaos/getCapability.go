// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package chaos

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a Capability resource that extends a Target resource.
// Azure REST API version: 2023-04-15-preview.
//
// Other available API versions: 2023-09-01-preview, 2023-10-27-preview.
func LookupCapability(ctx *pulumi.Context, args *LookupCapabilityArgs, opts ...pulumi.InvokeOption) (*LookupCapabilityResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCapabilityResult
	err := ctx.Invoke("azure-native:chaos:getCapability", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCapabilityArgs struct {
	// String that represents a Capability resource name.
	CapabilityName string `pulumi:"capabilityName"`
	// String that represents a resource provider namespace.
	ParentProviderNamespace string `pulumi:"parentProviderNamespace"`
	// String that represents a resource name.
	ParentResourceName string `pulumi:"parentResourceName"`
	// String that represents a resource type.
	ParentResourceType string `pulumi:"parentResourceType"`
	// String that represents an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// String that represents a Target resource name.
	TargetName string `pulumi:"targetName"`
}

// Model that represents a Capability resource.
type LookupCapabilityResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The properties of a capability resource.
	Properties CapabilityPropertiesResponse `pulumi:"properties"`
	// The standard system metadata of a resource type.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupCapabilityOutput(ctx *pulumi.Context, args LookupCapabilityOutputArgs, opts ...pulumi.InvokeOption) LookupCapabilityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCapabilityResult, error) {
			args := v.(LookupCapabilityArgs)
			r, err := LookupCapability(ctx, &args, opts...)
			var s LookupCapabilityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCapabilityResultOutput)
}

type LookupCapabilityOutputArgs struct {
	// String that represents a Capability resource name.
	CapabilityName pulumi.StringInput `pulumi:"capabilityName"`
	// String that represents a resource provider namespace.
	ParentProviderNamespace pulumi.StringInput `pulumi:"parentProviderNamespace"`
	// String that represents a resource name.
	ParentResourceName pulumi.StringInput `pulumi:"parentResourceName"`
	// String that represents a resource type.
	ParentResourceType pulumi.StringInput `pulumi:"parentResourceType"`
	// String that represents an Azure resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// String that represents a Target resource name.
	TargetName pulumi.StringInput `pulumi:"targetName"`
}

func (LookupCapabilityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapabilityArgs)(nil)).Elem()
}

// Model that represents a Capability resource.
type LookupCapabilityResultOutput struct{ *pulumi.OutputState }

func (LookupCapabilityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapabilityResult)(nil)).Elem()
}

func (o LookupCapabilityResultOutput) ToLookupCapabilityResultOutput() LookupCapabilityResultOutput {
	return o
}

func (o LookupCapabilityResultOutput) ToLookupCapabilityResultOutputWithContext(ctx context.Context) LookupCapabilityResultOutput {
	return o
}

func (o LookupCapabilityResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCapabilityResult] {
	return pulumix.Output[LookupCapabilityResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupCapabilityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapabilityResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCapabilityResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapabilityResult) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of a capability resource.
func (o LookupCapabilityResultOutput) Properties() CapabilityPropertiesResponseOutput {
	return o.ApplyT(func(v LookupCapabilityResult) CapabilityPropertiesResponse { return v.Properties }).(CapabilityPropertiesResponseOutput)
}

// The standard system metadata of a resource type.
func (o LookupCapabilityResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCapabilityResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCapabilityResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapabilityResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCapabilityResultOutput{})
}
