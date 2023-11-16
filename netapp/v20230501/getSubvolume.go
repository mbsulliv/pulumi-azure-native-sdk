// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the path associated with the subvolumeName provided
func LookupSubvolume(ctx *pulumi.Context, args *LookupSubvolumeArgs, opts ...pulumi.InvokeOption) (*LookupSubvolumeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSubvolumeResult
	err := ctx.Invoke("azure-native:netapp/v20230501:getSubvolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubvolumeArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the subvolume.
	SubvolumeName string `pulumi:"subvolumeName"`
	// The name of the volume
	VolumeName string `pulumi:"volumeName"`
}

// Subvolume Information properties
type LookupSubvolumeResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// parent path to the subvolume
	ParentPath *string `pulumi:"parentPath"`
	// Path to the subvolume
	Path *string `pulumi:"path"`
	// Azure lifecycle management
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSubvolumeOutput(ctx *pulumi.Context, args LookupSubvolumeOutputArgs, opts ...pulumi.InvokeOption) LookupSubvolumeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubvolumeResult, error) {
			args := v.(LookupSubvolumeArgs)
			r, err := LookupSubvolume(ctx, &args, opts...)
			var s LookupSubvolumeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubvolumeResultOutput)
}

type LookupSubvolumeOutputArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName pulumi.StringInput `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the subvolume.
	SubvolumeName pulumi.StringInput `pulumi:"subvolumeName"`
	// The name of the volume
	VolumeName pulumi.StringInput `pulumi:"volumeName"`
}

func (LookupSubvolumeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubvolumeArgs)(nil)).Elem()
}

// Subvolume Information properties
type LookupSubvolumeResultOutput struct{ *pulumi.OutputState }

func (LookupSubvolumeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubvolumeResult)(nil)).Elem()
}

func (o LookupSubvolumeResultOutput) ToLookupSubvolumeResultOutput() LookupSubvolumeResultOutput {
	return o
}

func (o LookupSubvolumeResultOutput) ToLookupSubvolumeResultOutputWithContext(ctx context.Context) LookupSubvolumeResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSubvolumeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubvolumeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSubvolumeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubvolumeResult) string { return v.Name }).(pulumi.StringOutput)
}

// parent path to the subvolume
func (o LookupSubvolumeResultOutput) ParentPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubvolumeResult) *string { return v.ParentPath }).(pulumi.StringPtrOutput)
}

// Path to the subvolume
func (o LookupSubvolumeResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubvolumeResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// Azure lifecycle management
func (o LookupSubvolumeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubvolumeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSubvolumeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSubvolumeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSubvolumeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubvolumeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubvolumeResultOutput{})
}
