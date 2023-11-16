// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetesconfiguration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a Azure Arc PrivateLinkScope.
// Azure REST API version: 2022-04-02-preview.
func LookupPrivateLinkScope(ctx *pulumi.Context, args *LookupPrivateLinkScopeArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkScopeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateLinkScopeResult
	err := ctx.Invoke("azure-native:kubernetesconfiguration:getPrivateLinkScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkScopeArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Azure Arc PrivateLinkScope resource.
	ScopeName string `pulumi:"scopeName"`
}

// An Azure Arc PrivateLinkScope definition.
type LookupPrivateLinkScopeResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties that define a Azure Arc PrivateLinkScope resource.
	Properties KubernetesConfigurationPrivateLinkScopePropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupPrivateLinkScopeOutput(ctx *pulumi.Context, args LookupPrivateLinkScopeOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateLinkScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateLinkScopeResult, error) {
			args := v.(LookupPrivateLinkScopeArgs)
			r, err := LookupPrivateLinkScope(ctx, &args, opts...)
			var s LookupPrivateLinkScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateLinkScopeResultOutput)
}

type LookupPrivateLinkScopeOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Azure Arc PrivateLinkScope resource.
	ScopeName pulumi.StringInput `pulumi:"scopeName"`
}

func (LookupPrivateLinkScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkScopeArgs)(nil)).Elem()
}

// An Azure Arc PrivateLinkScope definition.
type LookupPrivateLinkScopeResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateLinkScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkScopeResult)(nil)).Elem()
}

func (o LookupPrivateLinkScopeResultOutput) ToLookupPrivateLinkScopeResultOutput() LookupPrivateLinkScopeResultOutput {
	return o
}

func (o LookupPrivateLinkScopeResultOutput) ToLookupPrivateLinkScopeResultOutputWithContext(ctx context.Context) LookupPrivateLinkScopeResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupPrivateLinkScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupPrivateLinkScopeResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopeResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPrivateLinkScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties that define a Azure Arc PrivateLinkScope resource.
func (o LookupPrivateLinkScopeResultOutput) Properties() KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopeResult) KubernetesConfigurationPrivateLinkScopePropertiesResponse {
		return v.Properties
	}).(KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupPrivateLinkScopeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupPrivateLinkScopeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPrivateLinkScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateLinkScopeResultOutput{})
}
