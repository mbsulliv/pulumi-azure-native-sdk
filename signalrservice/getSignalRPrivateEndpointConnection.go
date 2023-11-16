// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalrservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the specified private endpoint connection
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview.
func LookupSignalRPrivateEndpointConnection(ctx *pulumi.Context, args *LookupSignalRPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupSignalRPrivateEndpointConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSignalRPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:signalrservice:getSignalRPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSignalRPrivateEndpointConnectionArgs struct {
	// The name of the private endpoint connection
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// A private endpoint connection to an azure resource
type LookupSignalRPrivateEndpointConnectionResult struct {
	// Group IDs
	GroupIds []string `pulumi:"groupIds"`
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Private endpoint
	PrivateEndpoint *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// Connection state of the private endpoint connection
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
	Type string `pulumi:"type"`
}

func LookupSignalRPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupSignalRPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupSignalRPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSignalRPrivateEndpointConnectionResult, error) {
			args := v.(LookupSignalRPrivateEndpointConnectionArgs)
			r, err := LookupSignalRPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupSignalRPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSignalRPrivateEndpointConnectionResultOutput)
}

type LookupSignalRPrivateEndpointConnectionOutputArgs struct {
	// The name of the private endpoint connection
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSignalRPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRPrivateEndpointConnectionArgs)(nil)).Elem()
}

// A private endpoint connection to an azure resource
type LookupSignalRPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupSignalRPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupSignalRPrivateEndpointConnectionResultOutput) ToLookupSignalRPrivateEndpointConnectionResultOutput() LookupSignalRPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupSignalRPrivateEndpointConnectionResultOutput) ToLookupSignalRPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupSignalRPrivateEndpointConnectionResultOutput {
	return o
}

// Group IDs
func (o LookupSignalRPrivateEndpointConnectionResultOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSignalRPrivateEndpointConnectionResult) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// Fully qualified resource Id for the resource.
func (o LookupSignalRPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupSignalRPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint
func (o LookupSignalRPrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRPrivateEndpointConnectionResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

// Connection state of the private endpoint connection
func (o LookupSignalRPrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRPrivateEndpointConnectionResult) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

// Provisioning state of the resource.
func (o LookupSignalRPrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRPrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupSignalRPrivateEndpointConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSignalRPrivateEndpointConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
func (o LookupSignalRPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSignalRPrivateEndpointConnectionResultOutput{})
}
