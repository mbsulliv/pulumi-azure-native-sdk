// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A managed cluster snapshot resource.
// Azure REST API version: 2023-05-02-preview.
//
// Other available API versions: 2023-06-02-preview, 2023-07-02-preview, 2023-08-02-preview.
func LookupManagedClusterSnapshot(ctx *pulumi.Context, args *LookupManagedClusterSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterSnapshotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedClusterSnapshotResult
	err := ctx.Invoke("azure-native:containerservice:getManagedClusterSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedClusterSnapshotArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
}

// A managed cluster snapshot resource.
type LookupManagedClusterSnapshotResult struct {
	// CreationData to be used to specify the source resource ID to create this snapshot.
	CreationData *CreationDataResponse `pulumi:"creationData"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// What the properties will be showed when getting managed cluster snapshot. Those properties are read-only.
	ManagedClusterPropertiesReadOnly ManagedClusterPropertiesForSnapshotResponse `pulumi:"managedClusterPropertiesReadOnly"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of a snapshot. The default is NodePool.
	SnapshotType *string `pulumi:"snapshotType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupManagedClusterSnapshotOutput(ctx *pulumi.Context, args LookupManagedClusterSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupManagedClusterSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedClusterSnapshotResult, error) {
			args := v.(LookupManagedClusterSnapshotArgs)
			r, err := LookupManagedClusterSnapshot(ctx, &args, opts...)
			var s LookupManagedClusterSnapshotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedClusterSnapshotResultOutput)
}

type LookupManagedClusterSnapshotOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupManagedClusterSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedClusterSnapshotArgs)(nil)).Elem()
}

// A managed cluster snapshot resource.
type LookupManagedClusterSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupManagedClusterSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedClusterSnapshotResult)(nil)).Elem()
}

func (o LookupManagedClusterSnapshotResultOutput) ToLookupManagedClusterSnapshotResultOutput() LookupManagedClusterSnapshotResultOutput {
	return o
}

func (o LookupManagedClusterSnapshotResultOutput) ToLookupManagedClusterSnapshotResultOutputWithContext(ctx context.Context) LookupManagedClusterSnapshotResultOutput {
	return o
}

func (o LookupManagedClusterSnapshotResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupManagedClusterSnapshotResult] {
	return pulumix.Output[LookupManagedClusterSnapshotResult]{
		OutputState: o.OutputState,
	}
}

// CreationData to be used to specify the source resource ID to create this snapshot.
func (o LookupManagedClusterSnapshotResultOutput) CreationData() CreationDataResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) *CreationDataResponse { return v.CreationData }).(CreationDataResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupManagedClusterSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupManagedClusterSnapshotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) string { return v.Location }).(pulumi.StringOutput)
}

// What the properties will be showed when getting managed cluster snapshot. Those properties are read-only.
func (o LookupManagedClusterSnapshotResultOutput) ManagedClusterPropertiesReadOnly() ManagedClusterPropertiesForSnapshotResponseOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) ManagedClusterPropertiesForSnapshotResponse {
		return v.ManagedClusterPropertiesReadOnly
	}).(ManagedClusterPropertiesForSnapshotResponseOutput)
}

// The name of the resource
func (o LookupManagedClusterSnapshotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of a snapshot. The default is NodePool.
func (o LookupManagedClusterSnapshotResultOutput) SnapshotType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) *string { return v.SnapshotType }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupManagedClusterSnapshotResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupManagedClusterSnapshotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupManagedClusterSnapshotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedClusterSnapshotResultOutput{})
}
