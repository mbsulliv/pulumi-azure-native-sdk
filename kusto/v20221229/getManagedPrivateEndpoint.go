// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221229

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a managed private endpoint.
func LookupManagedPrivateEndpoint(ctx *pulumi.Context, args *LookupManagedPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupManagedPrivateEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedPrivateEndpointResult
	err := ctx.Invoke("azure-native:kusto/v20221229:getManagedPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedPrivateEndpointArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the managed private endpoint.
	ManagedPrivateEndpointName string `pulumi:"managedPrivateEndpointName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing a managed private endpoint.
type LookupManagedPrivateEndpointResult struct {
	// The groupId in which the managed private endpoint is created.
	GroupId string `pulumi:"groupId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The ARM resource ID of the resource for which the managed private endpoint is created.
	PrivateLinkResourceId string `pulumi:"privateLinkResourceId"`
	// The region of the resource to which the managed private endpoint is created.
	PrivateLinkResourceRegion *string `pulumi:"privateLinkResourceRegion"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The user request message.
	RequestMessage *string `pulumi:"requestMessage"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupManagedPrivateEndpointOutput(ctx *pulumi.Context, args LookupManagedPrivateEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupManagedPrivateEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedPrivateEndpointResult, error) {
			args := v.(LookupManagedPrivateEndpointArgs)
			r, err := LookupManagedPrivateEndpoint(ctx, &args, opts...)
			var s LookupManagedPrivateEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedPrivateEndpointResultOutput)
}

type LookupManagedPrivateEndpointOutputArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the managed private endpoint.
	ManagedPrivateEndpointName pulumi.StringInput `pulumi:"managedPrivateEndpointName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedPrivateEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedPrivateEndpointArgs)(nil)).Elem()
}

// Class representing a managed private endpoint.
type LookupManagedPrivateEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupManagedPrivateEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedPrivateEndpointResult)(nil)).Elem()
}

func (o LookupManagedPrivateEndpointResultOutput) ToLookupManagedPrivateEndpointResultOutput() LookupManagedPrivateEndpointResultOutput {
	return o
}

func (o LookupManagedPrivateEndpointResultOutput) ToLookupManagedPrivateEndpointResultOutputWithContext(ctx context.Context) LookupManagedPrivateEndpointResultOutput {
	return o
}

// The groupId in which the managed private endpoint is created.
func (o LookupManagedPrivateEndpointResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupManagedPrivateEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupManagedPrivateEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ARM resource ID of the resource for which the managed private endpoint is created.
func (o LookupManagedPrivateEndpointResultOutput) PrivateLinkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.PrivateLinkResourceId }).(pulumi.StringOutput)
}

// The region of the resource to which the managed private endpoint is created.
func (o LookupManagedPrivateEndpointResultOutput) PrivateLinkResourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) *string { return v.PrivateLinkResourceRegion }).(pulumi.StringPtrOutput)
}

// The provisioned state of the resource.
func (o LookupManagedPrivateEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The user request message.
func (o LookupManagedPrivateEndpointResultOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupManagedPrivateEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupManagedPrivateEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedPrivateEndpointResultOutput{})
}
