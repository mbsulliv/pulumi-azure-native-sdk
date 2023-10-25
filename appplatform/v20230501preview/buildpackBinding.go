// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Buildpack Binding Resource object
type BuildpackBinding struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of a buildpack binding
	Properties BuildpackBindingPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBuildpackBinding registers a new resource with the given unique name, arguments, and options.
func NewBuildpackBinding(ctx *pulumi.Context,
	name string, args *BuildpackBindingArgs, opts ...pulumi.ResourceOption) (*BuildpackBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BuildServiceName == nil {
		return nil, errors.New("invalid value for required argument 'BuildServiceName'")
	}
	if args.BuilderName == nil {
		return nil, errors.New("invalid value for required argument 'BuilderName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230301preview:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230701preview:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230901preview:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20231101preview:BuildpackBinding"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BuildpackBinding
	err := ctx.RegisterResource("azure-native:appplatform/v20230501preview:BuildpackBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBuildpackBinding gets an existing BuildpackBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuildpackBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildpackBindingState, opts ...pulumi.ResourceOption) (*BuildpackBinding, error) {
	var resource BuildpackBinding
	err := ctx.ReadResource("azure-native:appplatform/v20230501preview:BuildpackBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BuildpackBinding resources.
type buildpackBindingState struct {
}

type BuildpackBindingState struct {
}

func (BuildpackBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildpackBindingState)(nil)).Elem()
}

type buildpackBindingArgs struct {
	// The name of the build service resource.
	BuildServiceName string `pulumi:"buildServiceName"`
	// The name of the builder resource.
	BuilderName string `pulumi:"builderName"`
	// The name of the Buildpack Binding Name
	BuildpackBindingName *string `pulumi:"buildpackBindingName"`
	// Properties of a buildpack binding
	Properties *BuildpackBindingProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a BuildpackBinding resource.
type BuildpackBindingArgs struct {
	// The name of the build service resource.
	BuildServiceName pulumi.StringInput
	// The name of the builder resource.
	BuilderName pulumi.StringInput
	// The name of the Buildpack Binding Name
	BuildpackBindingName pulumi.StringPtrInput
	// Properties of a buildpack binding
	Properties BuildpackBindingPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (BuildpackBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildpackBindingArgs)(nil)).Elem()
}

type BuildpackBindingInput interface {
	pulumi.Input

	ToBuildpackBindingOutput() BuildpackBindingOutput
	ToBuildpackBindingOutputWithContext(ctx context.Context) BuildpackBindingOutput
}

func (*BuildpackBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildpackBinding)(nil)).Elem()
}

func (i *BuildpackBinding) ToBuildpackBindingOutput() BuildpackBindingOutput {
	return i.ToBuildpackBindingOutputWithContext(context.Background())
}

func (i *BuildpackBinding) ToBuildpackBindingOutputWithContext(ctx context.Context) BuildpackBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildpackBindingOutput)
}

func (i *BuildpackBinding) ToOutput(ctx context.Context) pulumix.Output[*BuildpackBinding] {
	return pulumix.Output[*BuildpackBinding]{
		OutputState: i.ToBuildpackBindingOutputWithContext(ctx).OutputState,
	}
}

type BuildpackBindingOutput struct{ *pulumi.OutputState }

func (BuildpackBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildpackBinding)(nil)).Elem()
}

func (o BuildpackBindingOutput) ToBuildpackBindingOutput() BuildpackBindingOutput {
	return o
}

func (o BuildpackBindingOutput) ToBuildpackBindingOutputWithContext(ctx context.Context) BuildpackBindingOutput {
	return o
}

func (o BuildpackBindingOutput) ToOutput(ctx context.Context) pulumix.Output[*BuildpackBinding] {
	return pulumix.Output[*BuildpackBinding]{
		OutputState: o.OutputState,
	}
}

// The name of the resource.
func (o BuildpackBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildpackBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of a buildpack binding
func (o BuildpackBindingOutput) Properties() BuildpackBindingPropertiesResponseOutput {
	return o.ApplyT(func(v *BuildpackBinding) BuildpackBindingPropertiesResponseOutput { return v.Properties }).(BuildpackBindingPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o BuildpackBindingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BuildpackBinding) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o BuildpackBindingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildpackBinding) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BuildpackBindingOutput{})
}
