// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The Get ManagedNetworks operation gets a Managed Network Resource, specified by the resource group and Managed Network name
func LookupManagedNetwork(ctx *pulumi.Context, args *LookupManagedNetworkArgs, opts ...pulumi.InvokeOption) (*LookupManagedNetworkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedNetworkResult
	err := ctx.Invoke("azure-native:managednetwork/v20190601preview:getManagedNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedNetworkArgs struct {
	// The name of the Managed Network.
	ManagedNetworkName string `pulumi:"managedNetworkName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Managed Network resource
type LookupManagedNetworkResult struct {
	// The collection of groups and policies concerned with connectivity
	Connectivity ConnectivityCollectionResponse `pulumi:"connectivity"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the ManagedNetwork resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The collection of management groups, subscriptions, virtual networks, and subnets by the Managed Network. This is a read-only property that is reflective of all ScopeAssignments for this Managed Network
	Scope *ScopeResponse `pulumi:"scope"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}

func LookupManagedNetworkOutput(ctx *pulumi.Context, args LookupManagedNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupManagedNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedNetworkResult, error) {
			args := v.(LookupManagedNetworkArgs)
			r, err := LookupManagedNetwork(ctx, &args, opts...)
			var s LookupManagedNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedNetworkResultOutput)
}

type LookupManagedNetworkOutputArgs struct {
	// The name of the Managed Network.
	ManagedNetworkName pulumi.StringInput `pulumi:"managedNetworkName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedNetworkArgs)(nil)).Elem()
}

// The Managed Network resource
type LookupManagedNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupManagedNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedNetworkResult)(nil)).Elem()
}

func (o LookupManagedNetworkResultOutput) ToLookupManagedNetworkResultOutput() LookupManagedNetworkResultOutput {
	return o
}

func (o LookupManagedNetworkResultOutput) ToLookupManagedNetworkResultOutputWithContext(ctx context.Context) LookupManagedNetworkResultOutput {
	return o
}

func (o LookupManagedNetworkResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupManagedNetworkResult] {
	return pulumix.Output[LookupManagedNetworkResult]{
		OutputState: o.OutputState,
	}
}

// The collection of groups and policies concerned with connectivity
func (o LookupManagedNetworkResultOutput) Connectivity() ConnectivityCollectionResponseOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) ConnectivityCollectionResponse { return v.Connectivity }).(ConnectivityCollectionResponseOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupManagedNetworkResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupManagedNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupManagedNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupManagedNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the ManagedNetwork resource.
func (o LookupManagedNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The collection of management groups, subscriptions, virtual networks, and subnets by the Managed Network. This is a read-only property that is reflective of all ScopeAssignments for this Managed Network
func (o LookupManagedNetworkResultOutput) Scope() ScopeResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) *ScopeResponse { return v.Scope }).(ScopeResponsePtrOutput)
}

// Resource tags
func (o LookupManagedNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o LookupManagedNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedNetworkResultOutput{})
}
