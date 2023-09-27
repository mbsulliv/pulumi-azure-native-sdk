// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a private endpoint connection.
func LookupManagedInstancePrivateEndpointConnection(ctx *pulumi.Context, args *LookupManagedInstancePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstancePrivateEndpointConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedInstancePrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:sql/v20221101preview:getManagedInstancePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstancePrivateEndpointConnectionArgs struct {
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The name of the private endpoint connection.
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A private endpoint connection
type LookupManagedInstancePrivateEndpointConnectionResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Private endpoint which the connection belongs to.
	PrivateEndpoint *ManagedInstancePrivateEndpointPropertyResponse `pulumi:"privateEndpoint"`
	// Connection State of the Private Endpoint Connection.
	PrivateLinkServiceConnectionState *ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	// State of the Private Endpoint Connection.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupManagedInstancePrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupManagedInstancePrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupManagedInstancePrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedInstancePrivateEndpointConnectionResult, error) {
			args := v.(LookupManagedInstancePrivateEndpointConnectionArgs)
			r, err := LookupManagedInstancePrivateEndpointConnection(ctx, &args, opts...)
			var s LookupManagedInstancePrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedInstancePrivateEndpointConnectionResultOutput)
}

type LookupManagedInstancePrivateEndpointConnectionOutputArgs struct {
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	// The name of the private endpoint connection.
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedInstancePrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstancePrivateEndpointConnectionArgs)(nil)).Elem()
}

// A private endpoint connection
type LookupManagedInstancePrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupManagedInstancePrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstancePrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) ToLookupManagedInstancePrivateEndpointConnectionResultOutput() LookupManagedInstancePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) ToLookupManagedInstancePrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupManagedInstancePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupManagedInstancePrivateEndpointConnectionResult] {
	return pulumix.Output[LookupManagedInstancePrivateEndpointConnectionResult]{
		OutputState: o.OutputState,
	}
}

// Resource ID.
func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstancePrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstancePrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint which the connection belongs to.
func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) PrivateEndpoint() ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedInstancePrivateEndpointConnectionResult) *ManagedInstancePrivateEndpointPropertyResponse {
		return v.PrivateEndpoint
	}).(ManagedInstancePrivateEndpointPropertyResponsePtrOutput)
}

// Connection State of the Private Endpoint Connection.
func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedInstancePrivateEndpointConnectionResult) *ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

// State of the Private Endpoint Connection.
func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstancePrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstancePrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedInstancePrivateEndpointConnectionResultOutput{})
}
