// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a privateLinkHub
func LookupPrivateLinkHub(ctx *pulumi.Context, args *LookupPrivateLinkHubArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkHubResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateLinkHubResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getPrivateLinkHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkHubArgs struct {
	// Name of the privateLinkHub
	PrivateLinkHubName string `pulumi:"privateLinkHubName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A privateLinkHub
type LookupPrivateLinkHubResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of private endpoint connections
	PrivateEndpointConnections []PrivateEndpointConnectionForPrivateLinkHubBasicResponse `pulumi:"privateEndpointConnections"`
	// PrivateLinkHub provisioning state
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupPrivateLinkHubOutput(ctx *pulumi.Context, args LookupPrivateLinkHubOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateLinkHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateLinkHubResult, error) {
			args := v.(LookupPrivateLinkHubArgs)
			r, err := LookupPrivateLinkHub(ctx, &args, opts...)
			var s LookupPrivateLinkHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateLinkHubResultOutput)
}

type LookupPrivateLinkHubOutputArgs struct {
	// Name of the privateLinkHub
	PrivateLinkHubName pulumi.StringInput `pulumi:"privateLinkHubName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateLinkHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkHubArgs)(nil)).Elem()
}

// A privateLinkHub
type LookupPrivateLinkHubResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateLinkHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkHubResult)(nil)).Elem()
}

func (o LookupPrivateLinkHubResultOutput) ToLookupPrivateLinkHubResultOutput() LookupPrivateLinkHubResultOutput {
	return o
}

func (o LookupPrivateLinkHubResultOutput) ToLookupPrivateLinkHubResultOutputWithContext(ctx context.Context) LookupPrivateLinkHubResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupPrivateLinkHubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkHubResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupPrivateLinkHubResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkHubResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPrivateLinkHubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkHubResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of private endpoint connections
func (o LookupPrivateLinkHubResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateLinkHubResult) []PrivateEndpointConnectionForPrivateLinkHubBasicResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput)
}

// PrivateLinkHub provisioning state
func (o LookupPrivateLinkHubResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkHubResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupPrivateLinkHubResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateLinkHubResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPrivateLinkHubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkHubResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateLinkHubResultOutput{})
}
