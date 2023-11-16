// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get details of the specified volume group
func LookupVolumeGroup(ctx *pulumi.Context, args *LookupVolumeGroupArgs, opts ...pulumi.InvokeOption) (*LookupVolumeGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVolumeGroupResult
	err := ctx.Invoke("azure-native:netapp/v20221101preview:getVolumeGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeGroupArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the volumeGroup
	VolumeGroupName string `pulumi:"volumeGroupName"`
}

// Volume group resource for create
type LookupVolumeGroupResult struct {
	// Volume group details
	GroupMetaData *VolumeGroupMetaDataResponse `pulumi:"groupMetaData"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource type
	Type string `pulumi:"type"`
	// List of volumes from group
	Volumes []VolumeGroupVolumePropertiesResponse `pulumi:"volumes"`
}

func LookupVolumeGroupOutput(ctx *pulumi.Context, args LookupVolumeGroupOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeGroupResult, error) {
			args := v.(LookupVolumeGroupArgs)
			r, err := LookupVolumeGroup(ctx, &args, opts...)
			var s LookupVolumeGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeGroupResultOutput)
}

type LookupVolumeGroupOutputArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the volumeGroup
	VolumeGroupName pulumi.StringInput `pulumi:"volumeGroupName"`
}

func (LookupVolumeGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeGroupArgs)(nil)).Elem()
}

// Volume group resource for create
type LookupVolumeGroupResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeGroupResult)(nil)).Elem()
}

func (o LookupVolumeGroupResultOutput) ToLookupVolumeGroupResultOutput() LookupVolumeGroupResultOutput {
	return o
}

func (o LookupVolumeGroupResultOutput) ToLookupVolumeGroupResultOutputWithContext(ctx context.Context) LookupVolumeGroupResultOutput {
	return o
}

// Volume group details
func (o LookupVolumeGroupResultOutput) GroupMetaData() VolumeGroupMetaDataResponsePtrOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) *VolumeGroupMetaDataResponse { return v.GroupMetaData }).(VolumeGroupMetaDataResponsePtrOutput)
}

// Resource Id
func (o LookupVolumeGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupVolumeGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o LookupVolumeGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o LookupVolumeGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type
func (o LookupVolumeGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// List of volumes from group
func (o LookupVolumeGroupResultOutput) Volumes() VolumeGroupVolumePropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) []VolumeGroupVolumePropertiesResponse { return v.Volumes }).(VolumeGroupVolumePropertiesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeGroupResultOutput{})
}
