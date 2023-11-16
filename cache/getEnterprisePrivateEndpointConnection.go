// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cache

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified private endpoint connection associated with the RedisEnterprise cluster.
// Azure REST API version: 2023-03-01-preview.
//
// Other available API versions: 2023-07-01, 2023-08-01-preview, 2023-10-01-preview, 2023-11-01.
func LookupEnterprisePrivateEndpointConnection(ctx *pulumi.Context, args *LookupEnterprisePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEnterprisePrivateEndpointConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEnterprisePrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:cache:getEnterprisePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnterprisePrivateEndpointConnectionArgs struct {
	// The name of the RedisEnterprise cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Private Endpoint Connection resource.
type LookupEnterprisePrivateEndpointConnectionResult struct {
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
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupEnterprisePrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupEnterprisePrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupEnterprisePrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnterprisePrivateEndpointConnectionResult, error) {
			args := v.(LookupEnterprisePrivateEndpointConnectionArgs)
			r, err := LookupEnterprisePrivateEndpointConnection(ctx, &args, opts...)
			var s LookupEnterprisePrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnterprisePrivateEndpointConnectionResultOutput)
}

type LookupEnterprisePrivateEndpointConnectionOutputArgs struct {
	// The name of the RedisEnterprise cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEnterprisePrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnterprisePrivateEndpointConnectionArgs)(nil)).Elem()
}

// The Private Endpoint Connection resource.
type LookupEnterprisePrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupEnterprisePrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnterprisePrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupEnterprisePrivateEndpointConnectionResultOutput) ToLookupEnterprisePrivateEndpointConnectionResultOutput() LookupEnterprisePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupEnterprisePrivateEndpointConnectionResultOutput) ToLookupEnterprisePrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupEnterprisePrivateEndpointConnectionResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEnterprisePrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterprisePrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEnterprisePrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterprisePrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o LookupEnterprisePrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupEnterprisePrivateEndpointConnectionResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o LookupEnterprisePrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupEnterprisePrivateEndpointConnectionResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o LookupEnterprisePrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterprisePrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEnterprisePrivateEndpointConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnterprisePrivateEndpointConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEnterprisePrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterprisePrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnterprisePrivateEndpointConnectionResultOutput{})
}
