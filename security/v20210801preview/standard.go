// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Security Standard on a resource
type Standard struct {
	pulumi.CustomResourceState

	// category of the standard provided
	Category pulumi.StringPtrOutput `pulumi:"category"`
	// List of component objects containing component unique keys (such as assessment keys) to apply to standard scope.  Currently only supports assessment keys.
	Components StandardComponentPropertiesResponseArrayOutput `pulumi:"components"`
	// description of the standard
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// display name of the standard, equivalent to the standardId
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Entity tag is used for comparing two or more entities from the same requested resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Kind of the resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Location where the resource is stored
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// standard type (Custom or BuiltIn only currently)
	StandardType pulumi.StringOutput `pulumi:"standardType"`
	// List of all standard supported clouds.
	SupportedClouds pulumi.StringArrayOutput `pulumi:"supportedClouds"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// A list of key value pairs that describe the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStandard registers a new resource with the given unique name, arguments, and options.
func NewStandard(ctx *pulumi.Context,
	name string, args *StandardArgs, opts ...pulumi.ResourceOption) (*Standard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security:Standard"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Standard
	err := ctx.RegisterResource("azure-native:security/v20210801preview:Standard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStandard gets an existing Standard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStandard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StandardState, opts ...pulumi.ResourceOption) (*Standard, error) {
	var resource Standard
	err := ctx.ReadResource("azure-native:security/v20210801preview:Standard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Standard resources.
type standardState struct {
}

type StandardState struct {
}

func (StandardState) ElementType() reflect.Type {
	return reflect.TypeOf((*standardState)(nil)).Elem()
}

type standardArgs struct {
	// category of the standard provided
	Category *string `pulumi:"category"`
	// List of component objects containing component unique keys (such as assessment keys) to apply to standard scope.  Currently only supports assessment keys.
	Components []StandardComponentProperties `pulumi:"components"`
	// description of the standard
	Description *string `pulumi:"description"`
	// display name of the standard, equivalent to the standardId
	DisplayName *string `pulumi:"displayName"`
	// Kind of the resource
	Kind *string `pulumi:"kind"`
	// Location where the resource is stored
	Location *string `pulumi:"location"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Security Standard key - unique key for the standard type
	StandardId *string `pulumi:"standardId"`
	// List of all standard supported clouds.
	SupportedClouds []StandardSupportedClouds `pulumi:"supportedClouds"`
	// A list of key value pairs that describe the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Standard resource.
type StandardArgs struct {
	// category of the standard provided
	Category pulumi.StringPtrInput
	// List of component objects containing component unique keys (such as assessment keys) to apply to standard scope.  Currently only supports assessment keys.
	Components StandardComponentPropertiesArrayInput
	// description of the standard
	Description pulumi.StringPtrInput
	// display name of the standard, equivalent to the standardId
	DisplayName pulumi.StringPtrInput
	// Kind of the resource
	Kind pulumi.StringPtrInput
	// Location where the resource is stored
	Location pulumi.StringPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The Security Standard key - unique key for the standard type
	StandardId pulumi.StringPtrInput
	// List of all standard supported clouds.
	SupportedClouds StandardSupportedCloudsArrayInput
	// A list of key value pairs that describe the resource.
	Tags pulumi.StringMapInput
}

func (StandardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*standardArgs)(nil)).Elem()
}

type StandardInput interface {
	pulumi.Input

	ToStandardOutput() StandardOutput
	ToStandardOutputWithContext(ctx context.Context) StandardOutput
}

func (*Standard) ElementType() reflect.Type {
	return reflect.TypeOf((**Standard)(nil)).Elem()
}

func (i *Standard) ToStandardOutput() StandardOutput {
	return i.ToStandardOutputWithContext(context.Background())
}

func (i *Standard) ToStandardOutputWithContext(ctx context.Context) StandardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardOutput)
}

func (i *Standard) ToOutput(ctx context.Context) pulumix.Output[*Standard] {
	return pulumix.Output[*Standard]{
		OutputState: i.ToStandardOutputWithContext(ctx).OutputState,
	}
}

type StandardOutput struct{ *pulumi.OutputState }

func (StandardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Standard)(nil)).Elem()
}

func (o StandardOutput) ToStandardOutput() StandardOutput {
	return o
}

func (o StandardOutput) ToStandardOutputWithContext(ctx context.Context) StandardOutput {
	return o
}

func (o StandardOutput) ToOutput(ctx context.Context) pulumix.Output[*Standard] {
	return pulumix.Output[*Standard]{
		OutputState: o.OutputState,
	}
}

// category of the standard provided
func (o StandardOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Standard) pulumi.StringPtrOutput { return v.Category }).(pulumi.StringPtrOutput)
}

// List of component objects containing component unique keys (such as assessment keys) to apply to standard scope.  Currently only supports assessment keys.
func (o StandardOutput) Components() StandardComponentPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *Standard) StandardComponentPropertiesResponseArrayOutput { return v.Components }).(StandardComponentPropertiesResponseArrayOutput)
}

// description of the standard
func (o StandardOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Standard) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// display name of the standard, equivalent to the standardId
func (o StandardOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Standard) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Entity tag is used for comparing two or more entities from the same requested resource.
func (o StandardOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Standard) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Kind of the resource
func (o StandardOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Standard) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Location where the resource is stored
func (o StandardOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Standard) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o StandardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Standard) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// standard type (Custom or BuiltIn only currently)
func (o StandardOutput) StandardType() pulumi.StringOutput {
	return o.ApplyT(func(v *Standard) pulumi.StringOutput { return v.StandardType }).(pulumi.StringOutput)
}

// List of all standard supported clouds.
func (o StandardOutput) SupportedClouds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Standard) pulumi.StringArrayOutput { return v.SupportedClouds }).(pulumi.StringArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o StandardOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Standard) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// A list of key value pairs that describe the resource.
func (o StandardOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Standard) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o StandardOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Standard) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StandardOutput{})
}
