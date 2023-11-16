// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230906

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified private endpoint connection associated with the workspace.
func LookupWorkspacePrivateEndpointConnection(ctx *pulumi.Context, args *LookupWorkspacePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWorkspacePrivateEndpointConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspacePrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:healthcareapis/v20230906:getWorkspacePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspacePrivateEndpointConnectionArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The Private Endpoint Connection resource.
type LookupWorkspacePrivateEndpointConnectionResult struct {
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
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkspacePrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupWorkspacePrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspacePrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspacePrivateEndpointConnectionResult, error) {
			args := v.(LookupWorkspacePrivateEndpointConnectionArgs)
			r, err := LookupWorkspacePrivateEndpointConnection(ctx, &args, opts...)
			var s LookupWorkspacePrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspacePrivateEndpointConnectionResultOutput)
}

type LookupWorkspacePrivateEndpointConnectionOutputArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspacePrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspacePrivateEndpointConnectionArgs)(nil)).Elem()
}

// The Private Endpoint Connection resource.
type LookupWorkspacePrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspacePrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspacePrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupWorkspacePrivateEndpointConnectionResultOutput) ToLookupWorkspacePrivateEndpointConnectionResultOutput() LookupWorkspacePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupWorkspacePrivateEndpointConnectionResultOutput) ToLookupWorkspacePrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupWorkspacePrivateEndpointConnectionResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspacePrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspacePrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspacePrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspacePrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o LookupWorkspacePrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspacePrivateEndpointConnectionResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o LookupWorkspacePrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupWorkspacePrivateEndpointConnectionResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o LookupWorkspacePrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspacePrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupWorkspacePrivateEndpointConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkspacePrivateEndpointConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspacePrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspacePrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspacePrivateEndpointConnectionResultOutput{})
}
