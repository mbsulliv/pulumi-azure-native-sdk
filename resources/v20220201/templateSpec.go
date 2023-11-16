// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Template Spec object.
type TemplateSpec struct {
	pulumi.CustomResourceState

	// Template Spec description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Template Spec display name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The location of the Template Spec. It cannot be changed after Template Spec creation. It must be one of the supported Azure locations.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Template Spec metadata. Metadata is an open-ended object and is typically a collection of key-value pairs.
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// High-level information about the versions within this Template Spec. The keys are the version names. Only populated if the $expand query parameter is set to 'versions'.
	Versions TemplateSpecVersionInfoResponseMapOutput `pulumi:"versions"`
}

// NewTemplateSpec registers a new resource with the given unique name, arguments, and options.
func NewTemplateSpec(ctx *pulumi.Context,
	name string, args *TemplateSpecArgs, opts ...pulumi.ResourceOption) (*TemplateSpec, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:TemplateSpec"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190601preview:TemplateSpec"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210301preview:TemplateSpec"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210501:TemplateSpec"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TemplateSpec
	err := ctx.RegisterResource("azure-native:resources/v20220201:TemplateSpec", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplateSpec gets an existing TemplateSpec resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplateSpec(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateSpecState, opts ...pulumi.ResourceOption) (*TemplateSpec, error) {
	var resource TemplateSpec
	err := ctx.ReadResource("azure-native:resources/v20220201:TemplateSpec", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TemplateSpec resources.
type templateSpecState struct {
}

type TemplateSpecState struct {
}

func (TemplateSpecState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateSpecState)(nil)).Elem()
}

type templateSpecArgs struct {
	// Template Spec description.
	Description *string `pulumi:"description"`
	// Template Spec display name.
	DisplayName *string `pulumi:"displayName"`
	// The location of the Template Spec. It cannot be changed after Template Spec creation. It must be one of the supported Azure locations.
	Location *string `pulumi:"location"`
	// The Template Spec metadata. Metadata is an open-ended object and is typically a collection of key-value pairs.
	Metadata interface{} `pulumi:"metadata"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Name of the Template Spec.
	TemplateSpecName *string `pulumi:"templateSpecName"`
}

// The set of arguments for constructing a TemplateSpec resource.
type TemplateSpecArgs struct {
	// Template Spec description.
	Description pulumi.StringPtrInput
	// Template Spec display name.
	DisplayName pulumi.StringPtrInput
	// The location of the Template Spec. It cannot be changed after Template Spec creation. It must be one of the supported Azure locations.
	Location pulumi.StringPtrInput
	// The Template Spec metadata. Metadata is an open-ended object and is typically a collection of key-value pairs.
	Metadata pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Name of the Template Spec.
	TemplateSpecName pulumi.StringPtrInput
}

func (TemplateSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateSpecArgs)(nil)).Elem()
}

type TemplateSpecInput interface {
	pulumi.Input

	ToTemplateSpecOutput() TemplateSpecOutput
	ToTemplateSpecOutputWithContext(ctx context.Context) TemplateSpecOutput
}

func (*TemplateSpec) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateSpec)(nil)).Elem()
}

func (i *TemplateSpec) ToTemplateSpecOutput() TemplateSpecOutput {
	return i.ToTemplateSpecOutputWithContext(context.Background())
}

func (i *TemplateSpec) ToTemplateSpecOutputWithContext(ctx context.Context) TemplateSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateSpecOutput)
}

type TemplateSpecOutput struct{ *pulumi.OutputState }

func (TemplateSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateSpec)(nil)).Elem()
}

func (o TemplateSpecOutput) ToTemplateSpecOutput() TemplateSpecOutput {
	return o
}

func (o TemplateSpecOutput) ToTemplateSpecOutputWithContext(ctx context.Context) TemplateSpecOutput {
	return o
}

// Template Spec description.
func (o TemplateSpecOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateSpec) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Template Spec display name.
func (o TemplateSpecOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateSpec) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The location of the Template Spec. It cannot be changed after Template Spec creation. It must be one of the supported Azure locations.
func (o TemplateSpecOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateSpec) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The Template Spec metadata. Metadata is an open-ended object and is typically a collection of key-value pairs.
func (o TemplateSpecOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *TemplateSpec) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

// Name of this resource.
func (o TemplateSpecOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateSpec) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o TemplateSpecOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TemplateSpec) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o TemplateSpecOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TemplateSpec) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of this resource.
func (o TemplateSpecOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateSpec) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// High-level information about the versions within this Template Spec. The keys are the version names. Only populated if the $expand query parameter is set to 'versions'.
func (o TemplateSpecOutput) Versions() TemplateSpecVersionInfoResponseMapOutput {
	return o.ApplyT(func(v *TemplateSpec) TemplateSpecVersionInfoResponseMapOutput { return v.Versions }).(TemplateSpecVersionInfoResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(TemplateSpecOutput{})
}
