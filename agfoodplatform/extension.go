// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package agfoodplatform

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Extension resource.
// Azure REST API version: 2023-06-01-preview. Prior API version in Azure Native 1.x: 2020-05-12-preview.
//
// Other available API versions: 2021-09-01-preview.
type Extension struct {
	pulumi.CustomResourceState

	// Additional Api Properties.
	AdditionalApiProperties ApiPropertiesResponseMapOutput `pulumi:"additionalApiProperties"`
	// The ETag value to implement optimistic concurrency.
	ETag pulumi.StringOutput `pulumi:"eTag"`
	// Extension api docs link.
	ExtensionApiDocsLink pulumi.StringOutput `pulumi:"extensionApiDocsLink"`
	// Extension auth link.
	ExtensionAuthLink pulumi.StringOutput `pulumi:"extensionAuthLink"`
	// Extension category. e.g. weather/sensor/satellite.
	ExtensionCategory pulumi.StringOutput `pulumi:"extensionCategory"`
	// Extension Id.
	ExtensionId pulumi.StringOutput `pulumi:"extensionId"`
	// Installed extension version.
	InstalledExtensionVersion pulumi.StringOutput `pulumi:"installedExtensionVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewExtension registers a new resource with the given unique name, arguments, and options.
func NewExtension(ctx *pulumi.Context,
	name string, args *ExtensionArgs, opts ...pulumi.ResourceOption) (*Extension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataManagerForAgricultureResourceName == nil {
		return nil, errors.New("invalid value for required argument 'DataManagerForAgricultureResourceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:agfoodplatform/v20200512preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:agfoodplatform/v20210901preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:agfoodplatform/v20230601preview:Extension"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Extension
	err := ctx.RegisterResource("azure-native:agfoodplatform:Extension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtension gets an existing Extension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionState, opts ...pulumi.ResourceOption) (*Extension, error) {
	var resource Extension
	err := ctx.ReadResource("azure-native:agfoodplatform:Extension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Extension resources.
type extensionState struct {
}

type ExtensionState struct {
}

func (ExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionState)(nil)).Elem()
}

type extensionArgs struct {
	// Additional Api Properties.
	AdditionalApiProperties map[string]ApiProperties `pulumi:"additionalApiProperties"`
	// DataManagerForAgriculture resource name.
	DataManagerForAgricultureResourceName string `pulumi:"dataManagerForAgricultureResourceName"`
	// Id of extension resource.
	ExtensionId *string `pulumi:"extensionId"`
	// Extension Version.
	ExtensionVersion *string `pulumi:"extensionVersion"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Extension resource.
type ExtensionArgs struct {
	// Additional Api Properties.
	AdditionalApiProperties ApiPropertiesMapInput
	// DataManagerForAgriculture resource name.
	DataManagerForAgricultureResourceName pulumi.StringInput
	// Id of extension resource.
	ExtensionId pulumi.StringPtrInput
	// Extension Version.
	ExtensionVersion pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionArgs)(nil)).Elem()
}

type ExtensionInput interface {
	pulumi.Input

	ToExtensionOutput() ExtensionOutput
	ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput
}

func (*Extension) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (i *Extension) ToExtensionOutput() ExtensionOutput {
	return i.ToExtensionOutputWithContext(context.Background())
}

func (i *Extension) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionOutput)
}

func (i *Extension) ToOutput(ctx context.Context) pulumix.Output[*Extension] {
	return pulumix.Output[*Extension]{
		OutputState: i.ToExtensionOutputWithContext(ctx).OutputState,
	}
}

type ExtensionOutput struct{ *pulumi.OutputState }

func (ExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (o ExtensionOutput) ToExtensionOutput() ExtensionOutput {
	return o
}

func (o ExtensionOutput) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return o
}

func (o ExtensionOutput) ToOutput(ctx context.Context) pulumix.Output[*Extension] {
	return pulumix.Output[*Extension]{
		OutputState: o.OutputState,
	}
}

// Additional Api Properties.
func (o ExtensionOutput) AdditionalApiProperties() ApiPropertiesResponseMapOutput {
	return o.ApplyT(func(v *Extension) ApiPropertiesResponseMapOutput { return v.AdditionalApiProperties }).(ApiPropertiesResponseMapOutput)
}

// The ETag value to implement optimistic concurrency.
func (o ExtensionOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

// Extension api docs link.
func (o ExtensionOutput) ExtensionApiDocsLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.ExtensionApiDocsLink }).(pulumi.StringOutput)
}

// Extension auth link.
func (o ExtensionOutput) ExtensionAuthLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.ExtensionAuthLink }).(pulumi.StringOutput)
}

// Extension category. e.g. weather/sensor/satellite.
func (o ExtensionOutput) ExtensionCategory() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.ExtensionCategory }).(pulumi.StringOutput)
}

// Extension Id.
func (o ExtensionOutput) ExtensionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.ExtensionId }).(pulumi.StringOutput)
}

// Installed extension version.
func (o ExtensionOutput) InstalledExtensionVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.InstalledExtensionVersion }).(pulumi.StringOutput)
}

// The name of the resource
func (o ExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ExtensionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Extension) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExtensionOutput{})
}
