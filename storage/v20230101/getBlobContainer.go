// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets properties of a specified container.
func LookupBlobContainer(ctx *pulumi.Context, args *LookupBlobContainerArgs, opts ...pulumi.InvokeOption) (*LookupBlobContainerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBlobContainerResult
	err := ctx.Invoke("azure-native:storage/v20230101:getBlobContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobContainerArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ContainerName string `pulumi:"containerName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Properties of the blob container, including Id, resource name, resource type, Etag.
type LookupBlobContainerResult struct {
	// Default the container to use specified encryption scope for all writes.
	DefaultEncryptionScope *string `pulumi:"defaultEncryptionScope"`
	// Indicates whether the blob container was deleted.
	Deleted bool `pulumi:"deleted"`
	// Blob container deletion time.
	DeletedTime string `pulumi:"deletedTime"`
	// Block override of encryption scope from the container default.
	DenyEncryptionScopeOverride *bool `pulumi:"denyEncryptionScopeOverride"`
	// Enable NFSv3 all squash on blob container.
	EnableNfsV3AllSquash *bool `pulumi:"enableNfsV3AllSquash"`
	// Enable NFSv3 root squash on blob container.
	EnableNfsV3RootSquash *bool `pulumi:"enableNfsV3RootSquash"`
	// Resource Etag.
	Etag string `pulumi:"etag"`
	// The hasImmutabilityPolicy public property is set to true by SRP if ImmutabilityPolicy has been created for this container. The hasImmutabilityPolicy public property is set to false by SRP if ImmutabilityPolicy has not been created for this container.
	HasImmutabilityPolicy bool `pulumi:"hasImmutabilityPolicy"`
	// The hasLegalHold public property is set to true by SRP if there are at least one existing tag. The hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out. There can be a maximum of 1000 blob containers with hasLegalHold=true for a given account.
	HasLegalHold bool `pulumi:"hasLegalHold"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The ImmutabilityPolicy property of the container.
	ImmutabilityPolicy ImmutabilityPolicyPropertiesResponse `pulumi:"immutabilityPolicy"`
	// The object level immutability property of the container. The property is immutable and can only be set to true at the container creation time. Existing containers must undergo a migration process.
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioningResponse `pulumi:"immutableStorageWithVersioning"`
	// Returns the date and time the container was last modified.
	LastModifiedTime string `pulumi:"lastModifiedTime"`
	// Specifies whether the lease on a container is of infinite or fixed duration, only when the container is leased.
	LeaseDuration string `pulumi:"leaseDuration"`
	// Lease state of the container.
	LeaseState string `pulumi:"leaseState"`
	// The lease status of the container.
	LeaseStatus string `pulumi:"leaseStatus"`
	// The LegalHold property of the container.
	LegalHold LegalHoldPropertiesResponse `pulumi:"legalHold"`
	// A name-value pair to associate with the container as metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Specifies whether data in the container may be accessed publicly and the level of access.
	PublicAccess *string `pulumi:"publicAccess"`
	// Remaining retention days for soft deleted blob container.
	RemainingRetentionDays int `pulumi:"remainingRetentionDays"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The version of the deleted blob container.
	Version string `pulumi:"version"`
}

func LookupBlobContainerOutput(ctx *pulumi.Context, args LookupBlobContainerOutputArgs, opts ...pulumi.InvokeOption) LookupBlobContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobContainerResult, error) {
			args := v.(LookupBlobContainerArgs)
			r, err := LookupBlobContainer(ctx, &args, opts...)
			var s LookupBlobContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobContainerResultOutput)
}

type LookupBlobContainerOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ContainerName pulumi.StringInput `pulumi:"containerName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBlobContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobContainerArgs)(nil)).Elem()
}

// Properties of the blob container, including Id, resource name, resource type, Etag.
type LookupBlobContainerResultOutput struct{ *pulumi.OutputState }

func (LookupBlobContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobContainerResult)(nil)).Elem()
}

func (o LookupBlobContainerResultOutput) ToLookupBlobContainerResultOutput() LookupBlobContainerResultOutput {
	return o
}

func (o LookupBlobContainerResultOutput) ToLookupBlobContainerResultOutputWithContext(ctx context.Context) LookupBlobContainerResultOutput {
	return o
}

// Default the container to use specified encryption scope for all writes.
func (o LookupBlobContainerResultOutput) DefaultEncryptionScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) *string { return v.DefaultEncryptionScope }).(pulumi.StringPtrOutput)
}

// Indicates whether the blob container was deleted.
func (o LookupBlobContainerResultOutput) Deleted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) bool { return v.Deleted }).(pulumi.BoolOutput)
}

// Blob container deletion time.
func (o LookupBlobContainerResultOutput) DeletedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.DeletedTime }).(pulumi.StringOutput)
}

// Block override of encryption scope from the container default.
func (o LookupBlobContainerResultOutput) DenyEncryptionScopeOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) *bool { return v.DenyEncryptionScopeOverride }).(pulumi.BoolPtrOutput)
}

// Enable NFSv3 all squash on blob container.
func (o LookupBlobContainerResultOutput) EnableNfsV3AllSquash() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) *bool { return v.EnableNfsV3AllSquash }).(pulumi.BoolPtrOutput)
}

// Enable NFSv3 root squash on blob container.
func (o LookupBlobContainerResultOutput) EnableNfsV3RootSquash() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) *bool { return v.EnableNfsV3RootSquash }).(pulumi.BoolPtrOutput)
}

// Resource Etag.
func (o LookupBlobContainerResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The hasImmutabilityPolicy public property is set to true by SRP if ImmutabilityPolicy has been created for this container. The hasImmutabilityPolicy public property is set to false by SRP if ImmutabilityPolicy has not been created for this container.
func (o LookupBlobContainerResultOutput) HasImmutabilityPolicy() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) bool { return v.HasImmutabilityPolicy }).(pulumi.BoolOutput)
}

// The hasLegalHold public property is set to true by SRP if there are at least one existing tag. The hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out. There can be a maximum of 1000 blob containers with hasLegalHold=true for a given account.
func (o LookupBlobContainerResultOutput) HasLegalHold() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) bool { return v.HasLegalHold }).(pulumi.BoolOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupBlobContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ImmutabilityPolicy property of the container.
func (o LookupBlobContainerResultOutput) ImmutabilityPolicy() ImmutabilityPolicyPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) ImmutabilityPolicyPropertiesResponse { return v.ImmutabilityPolicy }).(ImmutabilityPolicyPropertiesResponseOutput)
}

// The object level immutability property of the container. The property is immutable and can only be set to true at the container creation time. Existing containers must undergo a migration process.
func (o LookupBlobContainerResultOutput) ImmutableStorageWithVersioning() ImmutableStorageWithVersioningResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) *ImmutableStorageWithVersioningResponse {
		return v.ImmutableStorageWithVersioning
	}).(ImmutableStorageWithVersioningResponsePtrOutput)
}

// Returns the date and time the container was last modified.
func (o LookupBlobContainerResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// Specifies whether the lease on a container is of infinite or fixed duration, only when the container is leased.
func (o LookupBlobContainerResultOutput) LeaseDuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.LeaseDuration }).(pulumi.StringOutput)
}

// Lease state of the container.
func (o LookupBlobContainerResultOutput) LeaseState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.LeaseState }).(pulumi.StringOutput)
}

// The lease status of the container.
func (o LookupBlobContainerResultOutput) LeaseStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.LeaseStatus }).(pulumi.StringOutput)
}

// The LegalHold property of the container.
func (o LookupBlobContainerResultOutput) LegalHold() LegalHoldPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) LegalHoldPropertiesResponse { return v.LegalHold }).(LegalHoldPropertiesResponseOutput)
}

// A name-value pair to associate with the container as metadata.
func (o LookupBlobContainerResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// The name of the resource
func (o LookupBlobContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies whether data in the container may be accessed publicly and the level of access.
func (o LookupBlobContainerResultOutput) PublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) *string { return v.PublicAccess }).(pulumi.StringPtrOutput)
}

// Remaining retention days for soft deleted blob container.
func (o LookupBlobContainerResultOutput) RemainingRetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) int { return v.RemainingRetentionDays }).(pulumi.IntOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBlobContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

// The version of the deleted blob container.
func (o LookupBlobContainerResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobContainerResultOutput{})
}
