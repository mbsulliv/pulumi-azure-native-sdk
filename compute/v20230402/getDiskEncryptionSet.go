// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230402

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a disk encryption set.
func LookupDiskEncryptionSet(ctx *pulumi.Context, args *LookupDiskEncryptionSetArgs, opts ...pulumi.InvokeOption) (*LookupDiskEncryptionSetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDiskEncryptionSetResult
	err := ctx.Invoke("azure-native:compute/v20230402:getDiskEncryptionSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskEncryptionSetArgs struct {
	// The name of the disk encryption set that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
	DiskEncryptionSetName string `pulumi:"diskEncryptionSetName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// disk encryption set resource.
type LookupDiskEncryptionSetResult struct {
	// The key vault key which is currently used by this disk encryption set.
	ActiveKey *KeyForDiskEncryptionSetResponse `pulumi:"activeKey"`
	// The error that was encountered during auto-key rotation. If an error is present, then auto-key rotation will not be attempted until the error on this disk encryption set is fixed.
	AutoKeyRotationError ApiErrorResponse `pulumi:"autoKeyRotationError"`
	// The type of key used to encrypt the data of the disk.
	EncryptionType *string `pulumi:"encryptionType"`
	// Multi-tenant application client id to access key vault in a different tenant. Setting the value to 'None' will clear the property.
	FederatedClientId *string `pulumi:"federatedClientId"`
	// Resource Id
	Id string `pulumi:"id"`
	// The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used to encrypt disks.
	Identity *EncryptionSetIdentityResponse `pulumi:"identity"`
	// The time when the active key of this disk encryption set was updated.
	LastKeyRotationTimestamp string `pulumi:"lastKeyRotationTimestamp"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// A readonly collection of key vault keys previously used by this disk encryption set while a key rotation is in progress. It will be empty if there is no ongoing key rotation.
	PreviousKeys []KeyForDiskEncryptionSetResponse `pulumi:"previousKeys"`
	// The disk encryption set provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Set this flag to true to enable auto-updating of this disk encryption set to the latest key version.
	RotationToLatestKeyVersionEnabled *bool `pulumi:"rotationToLatestKeyVersionEnabled"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupDiskEncryptionSetOutput(ctx *pulumi.Context, args LookupDiskEncryptionSetOutputArgs, opts ...pulumi.InvokeOption) LookupDiskEncryptionSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiskEncryptionSetResult, error) {
			args := v.(LookupDiskEncryptionSetArgs)
			r, err := LookupDiskEncryptionSet(ctx, &args, opts...)
			var s LookupDiskEncryptionSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiskEncryptionSetResultOutput)
}

type LookupDiskEncryptionSetOutputArgs struct {
	// The name of the disk encryption set that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
	DiskEncryptionSetName pulumi.StringInput `pulumi:"diskEncryptionSetName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDiskEncryptionSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskEncryptionSetArgs)(nil)).Elem()
}

// disk encryption set resource.
type LookupDiskEncryptionSetResultOutput struct{ *pulumi.OutputState }

func (LookupDiskEncryptionSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskEncryptionSetResult)(nil)).Elem()
}

func (o LookupDiskEncryptionSetResultOutput) ToLookupDiskEncryptionSetResultOutput() LookupDiskEncryptionSetResultOutput {
	return o
}

func (o LookupDiskEncryptionSetResultOutput) ToLookupDiskEncryptionSetResultOutputWithContext(ctx context.Context) LookupDiskEncryptionSetResultOutput {
	return o
}

// The key vault key which is currently used by this disk encryption set.
func (o LookupDiskEncryptionSetResultOutput) ActiveKey() KeyForDiskEncryptionSetResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) *KeyForDiskEncryptionSetResponse { return v.ActiveKey }).(KeyForDiskEncryptionSetResponsePtrOutput)
}

// The error that was encountered during auto-key rotation. If an error is present, then auto-key rotation will not be attempted until the error on this disk encryption set is fixed.
func (o LookupDiskEncryptionSetResultOutput) AutoKeyRotationError() ApiErrorResponseOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) ApiErrorResponse { return v.AutoKeyRotationError }).(ApiErrorResponseOutput)
}

// The type of key used to encrypt the data of the disk.
func (o LookupDiskEncryptionSetResultOutput) EncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) *string { return v.EncryptionType }).(pulumi.StringPtrOutput)
}

// Multi-tenant application client id to access key vault in a different tenant. Setting the value to 'None' will clear the property.
func (o LookupDiskEncryptionSetResultOutput) FederatedClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) *string { return v.FederatedClientId }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupDiskEncryptionSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) string { return v.Id }).(pulumi.StringOutput)
}

// The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used to encrypt disks.
func (o LookupDiskEncryptionSetResultOutput) Identity() EncryptionSetIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) *EncryptionSetIdentityResponse { return v.Identity }).(EncryptionSetIdentityResponsePtrOutput)
}

// The time when the active key of this disk encryption set was updated.
func (o LookupDiskEncryptionSetResultOutput) LastKeyRotationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) string { return v.LastKeyRotationTimestamp }).(pulumi.StringOutput)
}

// Resource location
func (o LookupDiskEncryptionSetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupDiskEncryptionSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) string { return v.Name }).(pulumi.StringOutput)
}

// A readonly collection of key vault keys previously used by this disk encryption set while a key rotation is in progress. It will be empty if there is no ongoing key rotation.
func (o LookupDiskEncryptionSetResultOutput) PreviousKeys() KeyForDiskEncryptionSetResponseArrayOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) []KeyForDiskEncryptionSetResponse { return v.PreviousKeys }).(KeyForDiskEncryptionSetResponseArrayOutput)
}

// The disk encryption set provisioning state.
func (o LookupDiskEncryptionSetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Set this flag to true to enable auto-updating of this disk encryption set to the latest key version.
func (o LookupDiskEncryptionSetResultOutput) RotationToLatestKeyVersionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) *bool { return v.RotationToLatestKeyVersionEnabled }).(pulumi.BoolPtrOutput)
}

// Resource tags
func (o LookupDiskEncryptionSetResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupDiskEncryptionSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiskEncryptionSetResultOutput{})
}
