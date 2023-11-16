// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of a storage account’s Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.
func LookupBlobServiceProperties(ctx *pulumi.Context, args *LookupBlobServicePropertiesArgs, opts ...pulumi.InvokeOption) (*LookupBlobServicePropertiesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBlobServicePropertiesResult
	err := ctx.Invoke("azure-native:storage/v20220901:getBlobServiceProperties", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobServicePropertiesArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the blob Service within the specified storage account. Blob Service Name must be 'default'
	BlobServicesName string `pulumi:"blobServicesName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The properties of a storage account’s Blob service.
type LookupBlobServicePropertiesResult struct {
	// Deprecated in favor of isVersioningEnabled property.
	AutomaticSnapshotPolicyEnabled *bool `pulumi:"automaticSnapshotPolicyEnabled"`
	// The blob service properties for change feed events.
	ChangeFeed *ChangeFeedResponse `pulumi:"changeFeed"`
	// The blob service properties for container soft delete.
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicyResponse `pulumi:"containerDeleteRetentionPolicy"`
	// Specifies CORS rules for the Blob service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Blob service.
	Cors *CorsRulesResponse `pulumi:"cors"`
	// DefaultServiceVersion indicates the default version to use for requests to the Blob service if an incoming request’s version is not specified. Possible values include version 2008-10-27 and all more recent versions.
	DefaultServiceVersion *string `pulumi:"defaultServiceVersion"`
	// The blob service properties for blob soft delete.
	DeleteRetentionPolicy *DeleteRetentionPolicyResponse `pulumi:"deleteRetentionPolicy"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Versioning is enabled if set to true.
	IsVersioningEnabled *bool `pulumi:"isVersioningEnabled"`
	// The blob service property to configure last access time based tracking policy.
	LastAccessTimeTrackingPolicy *LastAccessTimeTrackingPolicyResponse `pulumi:"lastAccessTimeTrackingPolicy"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The blob service properties for blob restore policy.
	RestorePolicy *RestorePolicyPropertiesResponse `pulumi:"restorePolicy"`
	// Sku name and tier.
	Sku SkuResponse `pulumi:"sku"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupBlobServicePropertiesOutput(ctx *pulumi.Context, args LookupBlobServicePropertiesOutputArgs, opts ...pulumi.InvokeOption) LookupBlobServicePropertiesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobServicePropertiesResult, error) {
			args := v.(LookupBlobServicePropertiesArgs)
			r, err := LookupBlobServiceProperties(ctx, &args, opts...)
			var s LookupBlobServicePropertiesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobServicePropertiesResultOutput)
}

type LookupBlobServicePropertiesOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the blob Service within the specified storage account. Blob Service Name must be 'default'
	BlobServicesName pulumi.StringInput `pulumi:"blobServicesName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBlobServicePropertiesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobServicePropertiesArgs)(nil)).Elem()
}

// The properties of a storage account’s Blob service.
type LookupBlobServicePropertiesResultOutput struct{ *pulumi.OutputState }

func (LookupBlobServicePropertiesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobServicePropertiesResult)(nil)).Elem()
}

func (o LookupBlobServicePropertiesResultOutput) ToLookupBlobServicePropertiesResultOutput() LookupBlobServicePropertiesResultOutput {
	return o
}

func (o LookupBlobServicePropertiesResultOutput) ToLookupBlobServicePropertiesResultOutputWithContext(ctx context.Context) LookupBlobServicePropertiesResultOutput {
	return o
}

// Deprecated in favor of isVersioningEnabled property.
func (o LookupBlobServicePropertiesResultOutput) AutomaticSnapshotPolicyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *bool { return v.AutomaticSnapshotPolicyEnabled }).(pulumi.BoolPtrOutput)
}

// The blob service properties for change feed events.
func (o LookupBlobServicePropertiesResultOutput) ChangeFeed() ChangeFeedResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *ChangeFeedResponse { return v.ChangeFeed }).(ChangeFeedResponsePtrOutput)
}

// The blob service properties for container soft delete.
func (o LookupBlobServicePropertiesResultOutput) ContainerDeleteRetentionPolicy() DeleteRetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *DeleteRetentionPolicyResponse {
		return v.ContainerDeleteRetentionPolicy
	}).(DeleteRetentionPolicyResponsePtrOutput)
}

// Specifies CORS rules for the Blob service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Blob service.
func (o LookupBlobServicePropertiesResultOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *CorsRulesResponse { return v.Cors }).(CorsRulesResponsePtrOutput)
}

// DefaultServiceVersion indicates the default version to use for requests to the Blob service if an incoming request’s version is not specified. Possible values include version 2008-10-27 and all more recent versions.
func (o LookupBlobServicePropertiesResultOutput) DefaultServiceVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *string { return v.DefaultServiceVersion }).(pulumi.StringPtrOutput)
}

// The blob service properties for blob soft delete.
func (o LookupBlobServicePropertiesResultOutput) DeleteRetentionPolicy() DeleteRetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *DeleteRetentionPolicyResponse {
		return v.DeleteRetentionPolicy
	}).(DeleteRetentionPolicyResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupBlobServicePropertiesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Versioning is enabled if set to true.
func (o LookupBlobServicePropertiesResultOutput) IsVersioningEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *bool { return v.IsVersioningEnabled }).(pulumi.BoolPtrOutput)
}

// The blob service property to configure last access time based tracking policy.
func (o LookupBlobServicePropertiesResultOutput) LastAccessTimeTrackingPolicy() LastAccessTimeTrackingPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *LastAccessTimeTrackingPolicyResponse {
		return v.LastAccessTimeTrackingPolicy
	}).(LastAccessTimeTrackingPolicyResponsePtrOutput)
}

// The name of the resource
func (o LookupBlobServicePropertiesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) string { return v.Name }).(pulumi.StringOutput)
}

// The blob service properties for blob restore policy.
func (o LookupBlobServicePropertiesResultOutput) RestorePolicy() RestorePolicyPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *RestorePolicyPropertiesResponse { return v.RestorePolicy }).(RestorePolicyPropertiesResponsePtrOutput)
}

// Sku name and tier.
func (o LookupBlobServicePropertiesResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBlobServicePropertiesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobServicePropertiesResultOutput{})
}
