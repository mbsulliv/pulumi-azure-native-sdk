// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified private endpoint connection associated with the service.
// Azure REST API version: 2021-03-08.
func LookupPrivateEndpointConnectionsForSCCPowershell(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionsForSCCPowershellArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionsForSCCPowershellResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateEndpointConnectionsForSCCPowershellResult
	err := ctx.Invoke("azure-native:securityandcompliance:getPrivateEndpointConnectionsForSCCPowershell", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionsForSCCPowershellArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName string `pulumi:"resourceName"`
}

// The Private Endpoint Connection resource.
type LookupPrivateEndpointConnectionsForSCCPowershellResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Required property for system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupPrivateEndpointConnectionsForSCCPowershellOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionsForSCCPowershellOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionsForSCCPowershellResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionsForSCCPowershellResult, error) {
			args := v.(LookupPrivateEndpointConnectionsForSCCPowershellArgs)
			r, err := LookupPrivateEndpointConnectionsForSCCPowershell(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionsForSCCPowershellResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionsForSCCPowershellResultOutput)
}

type LookupPrivateEndpointConnectionsForSCCPowershellOutputArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupPrivateEndpointConnectionsForSCCPowershellOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsForSCCPowershellArgs)(nil)).Elem()
}

// The Private Endpoint Connection resource.
type LookupPrivateEndpointConnectionsForSCCPowershellResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsForSCCPowershellResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) ToLookupPrivateEndpointConnectionsForSCCPowershellResultOutput() LookupPrivateEndpointConnectionsForSCCPowershellResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) ToLookupPrivateEndpointConnectionsForSCCPowershellResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionsForSCCPowershellResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPrivateEndpointConnectionsForSCCPowershellResult] {
	return pulumix.Output[LookupPrivateEndpointConnectionsForSCCPowershellResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForSCCPowershellResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForSCCPowershellResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForSCCPowershellResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForSCCPowershellResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForSCCPowershellResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Required property for system data
func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForSCCPowershellResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForSCCPowershellResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionsForSCCPowershellResultOutput{})
}
