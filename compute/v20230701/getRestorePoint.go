// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The operation to get the restore point.
func LookupRestorePoint(ctx *pulumi.Context, args *LookupRestorePointArgs, opts ...pulumi.InvokeOption) (*LookupRestorePointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRestorePointResult
	err := ctx.Invoke("azure-native:compute/v20230701:getRestorePoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRestorePointArgs struct {
	// The expand expression to apply on the operation. 'InstanceView' retrieves information about the run-time state of a restore point.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the restore point collection.
	RestorePointCollectionName string `pulumi:"restorePointCollectionName"`
	// The name of the restore point.
	RestorePointName string `pulumi:"restorePointName"`
}

// Restore Point details.
type LookupRestorePointResult struct {
	// ConsistencyMode of the RestorePoint. Can be specified in the input while creating a restore point. For now, only CrashConsistent is accepted as a valid input. Please refer to https://aka.ms/RestorePoints for more details.
	ConsistencyMode *string `pulumi:"consistencyMode"`
	// List of disk resource ids that the customer wishes to exclude from the restore point. If no disks are specified, all disks will be included.
	ExcludeDisks []ApiEntityReferenceResponse `pulumi:"excludeDisks"`
	// Resource Id
	Id string `pulumi:"id"`
	// The restore point instance view.
	InstanceView RestorePointInstanceViewResponse `pulumi:"instanceView"`
	// Resource name
	Name string `pulumi:"name"`
	// Gets the provisioning state of the restore point.
	ProvisioningState string `pulumi:"provisioningState"`
	// Gets the details of the VM captured at the time of the restore point creation.
	SourceMetadata *RestorePointSourceMetadataResponse `pulumi:"sourceMetadata"`
	// Resource Id of the source restore point from which a copy needs to be created.
	SourceRestorePoint *ApiEntityReferenceResponse `pulumi:"sourceRestorePoint"`
	// Gets the creation time of the restore point.
	TimeCreated *string `pulumi:"timeCreated"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupRestorePointOutput(ctx *pulumi.Context, args LookupRestorePointOutputArgs, opts ...pulumi.InvokeOption) LookupRestorePointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRestorePointResult, error) {
			args := v.(LookupRestorePointArgs)
			r, err := LookupRestorePoint(ctx, &args, opts...)
			var s LookupRestorePointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRestorePointResultOutput)
}

type LookupRestorePointOutputArgs struct {
	// The expand expression to apply on the operation. 'InstanceView' retrieves information about the run-time state of a restore point.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the restore point collection.
	RestorePointCollectionName pulumi.StringInput `pulumi:"restorePointCollectionName"`
	// The name of the restore point.
	RestorePointName pulumi.StringInput `pulumi:"restorePointName"`
}

func (LookupRestorePointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestorePointArgs)(nil)).Elem()
}

// Restore Point details.
type LookupRestorePointResultOutput struct{ *pulumi.OutputState }

func (LookupRestorePointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestorePointResult)(nil)).Elem()
}

func (o LookupRestorePointResultOutput) ToLookupRestorePointResultOutput() LookupRestorePointResultOutput {
	return o
}

func (o LookupRestorePointResultOutput) ToLookupRestorePointResultOutputWithContext(ctx context.Context) LookupRestorePointResultOutput {
	return o
}

func (o LookupRestorePointResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRestorePointResult] {
	return pulumix.Output[LookupRestorePointResult]{
		OutputState: o.OutputState,
	}
}

// ConsistencyMode of the RestorePoint. Can be specified in the input while creating a restore point. For now, only CrashConsistent is accepted as a valid input. Please refer to https://aka.ms/RestorePoints for more details.
func (o LookupRestorePointResultOutput) ConsistencyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRestorePointResult) *string { return v.ConsistencyMode }).(pulumi.StringPtrOutput)
}

// List of disk resource ids that the customer wishes to exclude from the restore point. If no disks are specified, all disks will be included.
func (o LookupRestorePointResultOutput) ExcludeDisks() ApiEntityReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupRestorePointResult) []ApiEntityReferenceResponse { return v.ExcludeDisks }).(ApiEntityReferenceResponseArrayOutput)
}

// Resource Id
func (o LookupRestorePointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointResult) string { return v.Id }).(pulumi.StringOutput)
}

// The restore point instance view.
func (o LookupRestorePointResultOutput) InstanceView() RestorePointInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupRestorePointResult) RestorePointInstanceViewResponse { return v.InstanceView }).(RestorePointInstanceViewResponseOutput)
}

// Resource name
func (o LookupRestorePointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state of the restore point.
func (o LookupRestorePointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets the details of the VM captured at the time of the restore point creation.
func (o LookupRestorePointResultOutput) SourceMetadata() RestorePointSourceMetadataResponsePtrOutput {
	return o.ApplyT(func(v LookupRestorePointResult) *RestorePointSourceMetadataResponse { return v.SourceMetadata }).(RestorePointSourceMetadataResponsePtrOutput)
}

// Resource Id of the source restore point from which a copy needs to be created.
func (o LookupRestorePointResultOutput) SourceRestorePoint() ApiEntityReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupRestorePointResult) *ApiEntityReferenceResponse { return v.SourceRestorePoint }).(ApiEntityReferenceResponsePtrOutput)
}

// Gets the creation time of the restore point.
func (o LookupRestorePointResultOutput) TimeCreated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRestorePointResult) *string { return v.TimeCreated }).(pulumi.StringPtrOutput)
}

// Resource type
func (o LookupRestorePointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRestorePointResultOutput{})
}
