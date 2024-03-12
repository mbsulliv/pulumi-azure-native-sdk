// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a storage container
func LookupStorageContainer(ctx *pulumi.Context, args *LookupStorageContainerArgs, opts ...pulumi.InvokeOption) (*LookupStorageContainerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStorageContainerResult
	err := ctx.Invoke("azure-native:azurestackhci/v20240101:getStorageContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageContainerArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the storage container
	StorageContainerName string `pulumi:"storageContainerName"`
}

// The storage container resource definition.
type LookupStorageContainerResult struct {
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Path of the storage container on the disk
	Path string `pulumi:"path"`
	// Provisioning state of the storage container.
	ProvisioningState string `pulumi:"provisioningState"`
	// The observed state of storage containers
	Status StorageContainerStatusResponse `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupStorageContainerOutput(ctx *pulumi.Context, args LookupStorageContainerOutputArgs, opts ...pulumi.InvokeOption) LookupStorageContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageContainerResult, error) {
			args := v.(LookupStorageContainerArgs)
			r, err := LookupStorageContainer(ctx, &args, opts...)
			var s LookupStorageContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageContainerResultOutput)
}

type LookupStorageContainerOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the storage container
	StorageContainerName pulumi.StringInput `pulumi:"storageContainerName"`
}

func (LookupStorageContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageContainerArgs)(nil)).Elem()
}

// The storage container resource definition.
type LookupStorageContainerResultOutput struct{ *pulumi.OutputState }

func (LookupStorageContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageContainerResult)(nil)).Elem()
}

func (o LookupStorageContainerResultOutput) ToLookupStorageContainerResultOutput() LookupStorageContainerResultOutput {
	return o
}

func (o LookupStorageContainerResultOutput) ToLookupStorageContainerResultOutputWithContext(ctx context.Context) LookupStorageContainerResultOutput {
	return o
}

// The extendedLocation of the resource.
func (o LookupStorageContainerResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageContainerResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupStorageContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupStorageContainerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageContainerResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupStorageContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Path of the storage container on the disk
func (o LookupStorageContainerResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageContainerResult) string { return v.Path }).(pulumi.StringOutput)
}

// Provisioning state of the storage container.
func (o LookupStorageContainerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageContainerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The observed state of storage containers
func (o LookupStorageContainerResultOutput) Status() StorageContainerStatusResponseOutput {
	return o.ApplyT(func(v LookupStorageContainerResult) StorageContainerStatusResponse { return v.Status }).(StorageContainerStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupStorageContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStorageContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupStorageContainerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStorageContainerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupStorageContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageContainerResultOutput{})
}
