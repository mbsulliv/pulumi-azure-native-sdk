// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Asset definition.
type Asset struct {
	pulumi.CustomResourceState

	// The extended location.
	ExtendedLocation AssetResponseExtendedLocationOutput `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Asset resource properties.
	Properties AssetResponsePropertiesOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAsset registers a new resource with the given unique name, arguments, and options.
func NewAsset(ctx *pulumi.Context,
	name string, args *AssetArgs, opts ...pulumi.ResourceOption) (*Asset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:deviceregistry:Asset"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Asset
	err := ctx.RegisterResource("azure-native:deviceregistry/v20231101preview:Asset", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:deviceregistry/v20231101preview:Asset", name, id, state, &resource, opts...)
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
	// Asset name parameter.
	AssetName *string `pulumi:"assetName"`
	// The extended location.
	ExtendedLocation AssetExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Asset resource properties.
	Properties *AssetProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Asset resource.
type AssetArgs struct {
	// Asset name parameter.
	AssetName pulumi.StringPtrInput
	// The extended location.
	ExtendedLocation AssetExtendedLocationInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Asset resource properties.
	Properties AssetPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
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

func (i *Asset) ToOutput(ctx context.Context) pulumix.Output[*Asset] {
	return pulumix.Output[*Asset]{
		OutputState: i.ToAssetOutputWithContext(ctx).OutputState,
	}
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

func (o AssetOutput) ToOutput(ctx context.Context) pulumix.Output[*Asset] {
	return pulumix.Output[*Asset]{
		OutputState: o.OutputState,
	}
}

// The extended location.
func (o AssetOutput) ExtendedLocation() AssetResponseExtendedLocationOutput {
	return o.ApplyT(func(v *Asset) AssetResponseExtendedLocationOutput { return v.ExtendedLocation }).(AssetResponseExtendedLocationOutput)
}

// The geo-location where the resource lives
func (o AssetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o AssetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Asset resource properties.
func (o AssetOutput) Properties() AssetResponsePropertiesOutput {
	return o.ApplyT(func(v *Asset) AssetResponsePropertiesOutput { return v.Properties }).(AssetResponsePropertiesOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AssetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Asset) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AssetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AssetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssetOutput{})
}
