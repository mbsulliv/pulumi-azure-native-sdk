// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The properties of a storage account’s Blob service.
type BlobServiceProperties struct {
	pulumi.CustomResourceState

	// Deprecated in favor of isVersioningEnabled property.
	AutomaticSnapshotPolicyEnabled pulumi.BoolPtrOutput `pulumi:"automaticSnapshotPolicyEnabled"`
	// The blob service properties for change feed events.
	ChangeFeed ChangeFeedResponsePtrOutput `pulumi:"changeFeed"`
	// The blob service properties for container soft delete.
	ContainerDeleteRetentionPolicy DeleteRetentionPolicyResponsePtrOutput `pulumi:"containerDeleteRetentionPolicy"`
	// Specifies CORS rules for the Blob service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Blob service.
	Cors CorsRulesResponsePtrOutput `pulumi:"cors"`
	// DefaultServiceVersion indicates the default version to use for requests to the Blob service if an incoming request’s version is not specified. Possible values include version 2008-10-27 and all more recent versions.
	DefaultServiceVersion pulumi.StringPtrOutput `pulumi:"defaultServiceVersion"`
	// The blob service properties for blob soft delete.
	DeleteRetentionPolicy DeleteRetentionPolicyResponsePtrOutput `pulumi:"deleteRetentionPolicy"`
	// Versioning is enabled if set to true.
	IsVersioningEnabled pulumi.BoolPtrOutput `pulumi:"isVersioningEnabled"`
	// The blob service property to configure last access time based tracking policy.
	LastAccessTimeTrackingPolicy LastAccessTimeTrackingPolicyResponsePtrOutput `pulumi:"lastAccessTimeTrackingPolicy"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The blob service properties for blob restore policy.
	RestorePolicy RestorePolicyPropertiesResponsePtrOutput `pulumi:"restorePolicy"`
	// Sku name and tier.
	Sku SkuResponseOutput `pulumi:"sku"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBlobServiceProperties registers a new resource with the given unique name, arguments, and options.
func NewBlobServiceProperties(ctx *pulumi.Context,
	name string, args *BlobServicePropertiesArgs, opts ...pulumi.ResourceOption) (*BlobServiceProperties, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180701:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20181101:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190401:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220901:BlobServiceProperties"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BlobServiceProperties
	err := ctx.RegisterResource("azure-native:storage/v20230101:BlobServiceProperties", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlobServiceProperties gets an existing BlobServiceProperties resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlobServiceProperties(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobServicePropertiesState, opts ...pulumi.ResourceOption) (*BlobServiceProperties, error) {
	var resource BlobServiceProperties
	err := ctx.ReadResource("azure-native:storage/v20230101:BlobServiceProperties", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlobServiceProperties resources.
type blobServicePropertiesState struct {
}

type BlobServicePropertiesState struct {
}

func (BlobServicePropertiesState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobServicePropertiesState)(nil)).Elem()
}

type blobServicePropertiesArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// Deprecated in favor of isVersioningEnabled property.
	AutomaticSnapshotPolicyEnabled *bool `pulumi:"automaticSnapshotPolicyEnabled"`
	// The name of the blob Service within the specified storage account. Blob Service Name must be 'default'
	BlobServicesName *string `pulumi:"blobServicesName"`
	// The blob service properties for change feed events.
	ChangeFeed *ChangeFeed `pulumi:"changeFeed"`
	// The blob service properties for container soft delete.
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicy `pulumi:"containerDeleteRetentionPolicy"`
	// Specifies CORS rules for the Blob service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Blob service.
	Cors *CorsRules `pulumi:"cors"`
	// DefaultServiceVersion indicates the default version to use for requests to the Blob service if an incoming request’s version is not specified. Possible values include version 2008-10-27 and all more recent versions.
	DefaultServiceVersion *string `pulumi:"defaultServiceVersion"`
	// The blob service properties for blob soft delete.
	DeleteRetentionPolicy *DeleteRetentionPolicy `pulumi:"deleteRetentionPolicy"`
	// Versioning is enabled if set to true.
	IsVersioningEnabled *bool `pulumi:"isVersioningEnabled"`
	// The blob service property to configure last access time based tracking policy.
	LastAccessTimeTrackingPolicy *LastAccessTimeTrackingPolicy `pulumi:"lastAccessTimeTrackingPolicy"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The blob service properties for blob restore policy.
	RestorePolicy *RestorePolicyProperties `pulumi:"restorePolicy"`
}

// The set of arguments for constructing a BlobServiceProperties resource.
type BlobServicePropertiesArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// Deprecated in favor of isVersioningEnabled property.
	AutomaticSnapshotPolicyEnabled pulumi.BoolPtrInput
	// The name of the blob Service within the specified storage account. Blob Service Name must be 'default'
	BlobServicesName pulumi.StringPtrInput
	// The blob service properties for change feed events.
	ChangeFeed ChangeFeedPtrInput
	// The blob service properties for container soft delete.
	ContainerDeleteRetentionPolicy DeleteRetentionPolicyPtrInput
	// Specifies CORS rules for the Blob service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Blob service.
	Cors CorsRulesPtrInput
	// DefaultServiceVersion indicates the default version to use for requests to the Blob service if an incoming request’s version is not specified. Possible values include version 2008-10-27 and all more recent versions.
	DefaultServiceVersion pulumi.StringPtrInput
	// The blob service properties for blob soft delete.
	DeleteRetentionPolicy DeleteRetentionPolicyPtrInput
	// Versioning is enabled if set to true.
	IsVersioningEnabled pulumi.BoolPtrInput
	// The blob service property to configure last access time based tracking policy.
	LastAccessTimeTrackingPolicy LastAccessTimeTrackingPolicyPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The blob service properties for blob restore policy.
	RestorePolicy RestorePolicyPropertiesPtrInput
}

func (BlobServicePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobServicePropertiesArgs)(nil)).Elem()
}

type BlobServicePropertiesInput interface {
	pulumi.Input

	ToBlobServicePropertiesOutput() BlobServicePropertiesOutput
	ToBlobServicePropertiesOutputWithContext(ctx context.Context) BlobServicePropertiesOutput
}

func (*BlobServiceProperties) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobServiceProperties)(nil)).Elem()
}

func (i *BlobServiceProperties) ToBlobServicePropertiesOutput() BlobServicePropertiesOutput {
	return i.ToBlobServicePropertiesOutputWithContext(context.Background())
}

func (i *BlobServiceProperties) ToBlobServicePropertiesOutputWithContext(ctx context.Context) BlobServicePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobServicePropertiesOutput)
}

type BlobServicePropertiesOutput struct{ *pulumi.OutputState }

func (BlobServicePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobServiceProperties)(nil)).Elem()
}

func (o BlobServicePropertiesOutput) ToBlobServicePropertiesOutput() BlobServicePropertiesOutput {
	return o
}

func (o BlobServicePropertiesOutput) ToBlobServicePropertiesOutputWithContext(ctx context.Context) BlobServicePropertiesOutput {
	return o
}

// Deprecated in favor of isVersioningEnabled property.
func (o BlobServicePropertiesOutput) AutomaticSnapshotPolicyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlobServiceProperties) pulumi.BoolPtrOutput { return v.AutomaticSnapshotPolicyEnabled }).(pulumi.BoolPtrOutput)
}

// The blob service properties for change feed events.
func (o BlobServicePropertiesOutput) ChangeFeed() ChangeFeedResponsePtrOutput {
	return o.ApplyT(func(v *BlobServiceProperties) ChangeFeedResponsePtrOutput { return v.ChangeFeed }).(ChangeFeedResponsePtrOutput)
}

// The blob service properties for container soft delete.
func (o BlobServicePropertiesOutput) ContainerDeleteRetentionPolicy() DeleteRetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v *BlobServiceProperties) DeleteRetentionPolicyResponsePtrOutput {
		return v.ContainerDeleteRetentionPolicy
	}).(DeleteRetentionPolicyResponsePtrOutput)
}

// Specifies CORS rules for the Blob service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Blob service.
func (o BlobServicePropertiesOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v *BlobServiceProperties) CorsRulesResponsePtrOutput { return v.Cors }).(CorsRulesResponsePtrOutput)
}

// DefaultServiceVersion indicates the default version to use for requests to the Blob service if an incoming request’s version is not specified. Possible values include version 2008-10-27 and all more recent versions.
func (o BlobServicePropertiesOutput) DefaultServiceVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobServiceProperties) pulumi.StringPtrOutput { return v.DefaultServiceVersion }).(pulumi.StringPtrOutput)
}

// The blob service properties for blob soft delete.
func (o BlobServicePropertiesOutput) DeleteRetentionPolicy() DeleteRetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v *BlobServiceProperties) DeleteRetentionPolicyResponsePtrOutput { return v.DeleteRetentionPolicy }).(DeleteRetentionPolicyResponsePtrOutput)
}

// Versioning is enabled if set to true.
func (o BlobServicePropertiesOutput) IsVersioningEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlobServiceProperties) pulumi.BoolPtrOutput { return v.IsVersioningEnabled }).(pulumi.BoolPtrOutput)
}

// The blob service property to configure last access time based tracking policy.
func (o BlobServicePropertiesOutput) LastAccessTimeTrackingPolicy() LastAccessTimeTrackingPolicyResponsePtrOutput {
	return o.ApplyT(func(v *BlobServiceProperties) LastAccessTimeTrackingPolicyResponsePtrOutput {
		return v.LastAccessTimeTrackingPolicy
	}).(LastAccessTimeTrackingPolicyResponsePtrOutput)
}

// The name of the resource
func (o BlobServicePropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobServiceProperties) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The blob service properties for blob restore policy.
func (o BlobServicePropertiesOutput) RestorePolicy() RestorePolicyPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *BlobServiceProperties) RestorePolicyPropertiesResponsePtrOutput { return v.RestorePolicy }).(RestorePolicyPropertiesResponsePtrOutput)
}

// Sku name and tier.
func (o BlobServicePropertiesOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *BlobServiceProperties) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o BlobServicePropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobServiceProperties) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BlobServicePropertiesOutput{})
}
