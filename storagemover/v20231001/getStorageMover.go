// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a Storage Mover resource.
func LookupStorageMover(ctx *pulumi.Context, args *LookupStorageMoverArgs, opts ...pulumi.InvokeOption) (*LookupStorageMoverResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStorageMoverResult
	err := ctx.Invoke("azure-native:storagemover/v20231001:getStorageMover", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageMoverArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Storage Mover resource.
	StorageMoverName string `pulumi:"storageMoverName"`
}

// The Storage Mover resource, which is a container for a group of Agents, Projects, and Endpoints.
type LookupStorageMoverResult struct {
	// A description for the Storage Mover.
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of this resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupStorageMoverOutput(ctx *pulumi.Context, args LookupStorageMoverOutputArgs, opts ...pulumi.InvokeOption) LookupStorageMoverResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageMoverResult, error) {
			args := v.(LookupStorageMoverArgs)
			r, err := LookupStorageMover(ctx, &args, opts...)
			var s LookupStorageMoverResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageMoverResultOutput)
}

type LookupStorageMoverOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Storage Mover resource.
	StorageMoverName pulumi.StringInput `pulumi:"storageMoverName"`
}

func (LookupStorageMoverOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageMoverArgs)(nil)).Elem()
}

// The Storage Mover resource, which is a container for a group of Agents, Projects, and Endpoints.
type LookupStorageMoverResultOutput struct{ *pulumi.OutputState }

func (LookupStorageMoverResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageMoverResult)(nil)).Elem()
}

func (o LookupStorageMoverResultOutput) ToLookupStorageMoverResultOutput() LookupStorageMoverResultOutput {
	return o
}

func (o LookupStorageMoverResultOutput) ToLookupStorageMoverResultOutputWithContext(ctx context.Context) LookupStorageMoverResultOutput {
	return o
}

func (o LookupStorageMoverResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupStorageMoverResult] {
	return pulumix.Output[LookupStorageMoverResult]{
		OutputState: o.OutputState,
	}
}

// A description for the Storage Mover.
func (o LookupStorageMoverResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupStorageMoverResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupStorageMoverResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupStorageMoverResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of this resource.
func (o LookupStorageMoverResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupStorageMoverResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupStorageMoverResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupStorageMoverResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageMoverResultOutput{})
}