// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A datastore resource
func LookupDatastore(ctx *pulumi.Context, args *LookupDatastoreArgs, opts ...pulumi.InvokeOption) (*LookupDatastoreResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDatastoreResult
	err := ctx.Invoke("azure-native:avs/v20220501:getDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDatastoreArgs struct {
	// Name of the cluster in the private cloud
	ClusterName string `pulumi:"clusterName"`
	// Name of the datastore in the private cloud cluster
	DatastoreName string `pulumi:"datastoreName"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A datastore resource
type LookupDatastoreResult struct {
	// An iSCSI volume
	DiskPoolVolume *DiskPoolVolumeResponse `pulumi:"diskPoolVolume"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// An Azure NetApp Files volume
	NetAppVolume *NetAppVolumeResponse `pulumi:"netAppVolume"`
	// The state of the datastore provisioning
	ProvisioningState string `pulumi:"provisioningState"`
	// The operational status of the datastore
	Status string `pulumi:"status"`
	// Resource type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupDatastoreResult
func (val *LookupDatastoreResult) Defaults() *LookupDatastoreResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DiskPoolVolume = tmp.DiskPoolVolume.Defaults()

	return &tmp
}

func LookupDatastoreOutput(ctx *pulumi.Context, args LookupDatastoreOutputArgs, opts ...pulumi.InvokeOption) LookupDatastoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatastoreResult, error) {
			args := v.(LookupDatastoreArgs)
			r, err := LookupDatastore(ctx, &args, opts...)
			var s LookupDatastoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatastoreResultOutput)
}

type LookupDatastoreOutputArgs struct {
	// Name of the cluster in the private cloud
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// Name of the datastore in the private cloud cluster
	DatastoreName pulumi.StringInput `pulumi:"datastoreName"`
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatastoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreArgs)(nil)).Elem()
}

// A datastore resource
type LookupDatastoreResultOutput struct{ *pulumi.OutputState }

func (LookupDatastoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreResult)(nil)).Elem()
}

func (o LookupDatastoreResultOutput) ToLookupDatastoreResultOutput() LookupDatastoreResultOutput {
	return o
}

func (o LookupDatastoreResultOutput) ToLookupDatastoreResultOutputWithContext(ctx context.Context) LookupDatastoreResultOutput {
	return o
}

// An iSCSI volume
func (o LookupDatastoreResultOutput) DiskPoolVolume() DiskPoolVolumeResponsePtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *DiskPoolVolumeResponse { return v.DiskPoolVolume }).(DiskPoolVolumeResponsePtrOutput)
}

// Resource ID.
func (o LookupDatastoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupDatastoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Name }).(pulumi.StringOutput)
}

// An Azure NetApp Files volume
func (o LookupDatastoreResultOutput) NetAppVolume() NetAppVolumeResponsePtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *NetAppVolumeResponse { return v.NetAppVolume }).(NetAppVolumeResponsePtrOutput)
}

// The state of the datastore provisioning
func (o LookupDatastoreResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The operational status of the datastore
func (o LookupDatastoreResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Status }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupDatastoreResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatastoreResultOutput{})
}
