// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Asset.
type Asset struct {
	pulumi.CustomResourceState

	// The alternate ID of the Asset.
	AlternateId pulumi.StringPtrOutput `pulumi:"alternateId"`
	// The Asset ID.
	AssetId pulumi.StringOutput `pulumi:"assetId"`
	// The name of the asset blob container.
	Container pulumi.StringPtrOutput `pulumi:"container"`
	// The creation date of the Asset.
	Created pulumi.StringOutput `pulumi:"created"`
	// The Asset description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Asset container encryption scope in the storage account.
	EncryptionScope pulumi.StringPtrOutput `pulumi:"encryptionScope"`
	// The last modified date of the Asset.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the storage account.
	StorageAccountName pulumi.StringPtrOutput `pulumi:"storageAccountName"`
	// The Asset encryption format. One of None or MediaStorageEncryption.
	StorageEncryptionFormat pulumi.StringOutput `pulumi:"storageEncryptionFormat"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAsset registers a new resource with the given unique name, arguments, and options.
func NewAsset(ctx *pulumi.Context,
	name string, args *AssetArgs, opts ...pulumi.ResourceOption) (*Asset, error) {
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
			Type: pulumi.String("azure-native:media:Asset"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:Asset"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:Asset"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:Asset"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:Asset"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:Asset"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:Asset"),
		},
		{
			Type: pulumi.String("azure-native:media/v20220801:Asset"),
		},
	})
	opts = append(opts, aliases)
	var resource Asset
	err := ctx.RegisterResource("azure-native:media/v20230101:Asset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAsset gets an existing Asset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAsset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssetState, opts ...pulumi.ResourceOption) (*Asset, error) {
	var resource Asset
	err := ctx.ReadResource("azure-native:media/v20230101:Asset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Asset resources.
type assetState struct {
}

type AssetState struct {
}

func (AssetState) ElementType() reflect.Type {
	return reflect.TypeOf((*assetState)(nil)).Elem()
}

type assetArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The alternate ID of the Asset.
	AlternateId *string `pulumi:"alternateId"`
	// The Asset name.
	AssetName *string `pulumi:"assetName"`
	// The name of the asset blob container.
	Container *string `pulumi:"container"`
	// The Asset description.
	Description *string `pulumi:"description"`
	// The Asset container encryption scope in the storage account.
	EncryptionScope *string `pulumi:"encryptionScope"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the storage account.
	StorageAccountName *string `pulumi:"storageAccountName"`
}

// The set of arguments for constructing a Asset resource.
type AssetArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The alternate ID of the Asset.
	AlternateId pulumi.StringPtrInput
	// The Asset name.
	AssetName pulumi.StringPtrInput
	// The name of the asset blob container.
	Container pulumi.StringPtrInput
	// The Asset description.
	Description pulumi.StringPtrInput
	// The Asset container encryption scope in the storage account.
	EncryptionScope pulumi.StringPtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The name of the storage account.
	StorageAccountName pulumi.StringPtrInput
}

func (AssetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assetArgs)(nil)).Elem()
}

type AssetInput interface {
	pulumi.Input

	ToAssetOutput() AssetOutput
	ToAssetOutputWithContext(ctx context.Context) AssetOutput
}

func (*Asset) ElementType() reflect.Type {
	return reflect.TypeOf((**Asset)(nil)).Elem()
}

func (i *Asset) ToAssetOutput() AssetOutput {
	return i.ToAssetOutputWithContext(context.Background())
}

func (i *Asset) ToAssetOutputWithContext(ctx context.Context) AssetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetOutput)
}

type AssetOutput struct{ *pulumi.OutputState }

func (AssetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Asset)(nil)).Elem()
}

func (o AssetOutput) ToAssetOutput() AssetOutput {
	return o
}

func (o AssetOutput) ToAssetOutputWithContext(ctx context.Context) AssetOutput {
	return o
}

// The alternate ID of the Asset.
func (o AssetOutput) AlternateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.AlternateId }).(pulumi.StringPtrOutput)
}

// The Asset ID.
func (o AssetOutput) AssetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.AssetId }).(pulumi.StringOutput)
}

// The name of the asset blob container.
func (o AssetOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.Container }).(pulumi.StringPtrOutput)
}

// The creation date of the Asset.
func (o AssetOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// The Asset description.
func (o AssetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Asset container encryption scope in the storage account.
func (o AssetOutput) EncryptionScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.EncryptionScope }).(pulumi.StringPtrOutput)
}

// The last modified date of the Asset.
func (o AssetOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

// The name of the resource
func (o AssetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the storage account.
func (o AssetOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

// The Asset encryption format. One of None or MediaStorageEncryption.
func (o AssetOutput) StorageEncryptionFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.StorageEncryptionFormat }).(pulumi.StringOutput)
}

// The system metadata relating to this resource.
func (o AssetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Asset) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AssetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssetOutput{})
}
