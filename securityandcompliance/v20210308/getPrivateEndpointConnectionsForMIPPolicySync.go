// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210308

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified private endpoint connection associated with the service.
func LookupPrivateEndpointConnectionsForMIPPolicySync(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionsForMIPPolicySyncArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionsForMIPPolicySyncResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateEndpointConnectionsForMIPPolicySyncResult
	err := ctx.Invoke("azure-native:securityandcompliance/v20210308:getPrivateEndpointConnectionsForMIPPolicySync", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionsForMIPPolicySyncArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName string `pulumi:"resourceName"`
}

// The Private Endpoint Connection resource.
type LookupPrivateEndpointConnectionsForMIPPolicySyncResult struct {
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

func LookupPrivateEndpointConnectionsForMIPPolicySyncOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionsForMIPPolicySyncOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionsForMIPPolicySyncResult, error) {
			args := v.(LookupPrivateEndpointConnectionsForMIPPolicySyncArgs)
			r, err := LookupPrivateEndpointConnectionsForMIPPolicySync(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionsForMIPPolicySyncResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput)
}

type LookupPrivateEndpointConnectionsForMIPPolicySyncOutputArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupPrivateEndpointConnectionsForMIPPolicySyncOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsForMIPPolicySyncArgs)(nil)).Elem()
}

// The Private Endpoint Connection resource.
type LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsForMIPPolicySyncResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) ToLookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput() LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) ToLookupPrivateEndpointConnectionsForMIPPolicySyncResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPrivateEndpointConnectionsForMIPPolicySyncResult] {
	return pulumix.Output[LookupPrivateEndpointConnectionsForMIPPolicySyncResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForMIPPolicySyncResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForMIPPolicySyncResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForMIPPolicySyncResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForMIPPolicySyncResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForMIPPolicySyncResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Required property for system data
func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForMIPPolicySyncResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForMIPPolicySyncResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput{})
}
