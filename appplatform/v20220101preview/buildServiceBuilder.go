// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// KPack Builder resource
type BuildServiceBuilder struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Property of the Builder resource.
	Properties BuilderPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBuildServiceBuilder registers a new resource with the given unique name, arguments, and options.
func NewBuildServiceBuilder(ctx *pulumi.Context,
	name string, args *BuildServiceBuilderArgs, opts ...pulumi.ResourceOption) (*BuildServiceBuilder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BuildServiceName == nil {
		return nil, errors.New("invalid value for required argument 'BuildServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:BuildServiceBuilder"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:BuildServiceBuilder"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:BuildServiceBuilder"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:BuildServiceBuilder"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:BuildServiceBuilder"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:BuildServiceBuilder"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:BuildServiceBuilder"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:BuildServiceBuilder"),
		},
	})
	opts = append(opts, aliases)
	var resource BuildServiceBuilder
	err := ctx.RegisterResource("azure-native:appplatform/v20220101preview:BuildServiceBuilder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBuildServiceBuilder gets an existing BuildServiceBuilder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuildServiceBuilder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildServiceBuilderState, opts ...pulumi.ResourceOption) (*BuildServiceBuilder, error) {
	var resource BuildServiceBuilder
	err := ctx.ReadResource("azure-native:appplatform/v20220101preview:BuildServiceBuilder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BuildServiceBuilder resources.
type buildServiceBuilderState struct {
}

type BuildServiceBuilderState struct {
}

func (BuildServiceBuilderState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildServiceBuilderState)(nil)).Elem()
}

type buildServiceBuilderArgs struct {
	// The name of the build service resource.
	BuildServiceName string `pulumi:"buildServiceName"`
	// The name of the builder resource.
	BuilderName *string `pulumi:"builderName"`
	// Property of the Builder resource.
	Properties *BuilderProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a BuildServiceBuilder resource.
type BuildServiceBuilderArgs struct {
	// The name of the build service resource.
	BuildServiceName pulumi.StringInput
	// The name of the builder resource.
	BuilderName pulumi.StringPtrInput
	// Property of the Builder resource.
	Properties BuilderPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (BuildServiceBuilderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildServiceBuilderArgs)(nil)).Elem()
}

type BuildServiceBuilderInput interface {
	pulumi.Input

	ToBuildServiceBuilderOutput() BuildServiceBuilderOutput
	ToBuildServiceBuilderOutputWithContext(ctx context.Context) BuildServiceBuilderOutput
}

func (*BuildServiceBuilder) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildServiceBuilder)(nil)).Elem()
}

func (i *BuildServiceBuilder) ToBuildServiceBuilderOutput() BuildServiceBuilderOutput {
	return i.ToBuildServiceBuilderOutputWithContext(context.Background())
}

func (i *BuildServiceBuilder) ToBuildServiceBuilderOutputWithContext(ctx context.Context) BuildServiceBuilderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildServiceBuilderOutput)
}

type BuildServiceBuilderOutput struct{ *pulumi.OutputState }

func (BuildServiceBuilderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildServiceBuilder)(nil)).Elem()
}

func (o BuildServiceBuilderOutput) ToBuildServiceBuilderOutput() BuildServiceBuilderOutput {
	return o
}

func (o BuildServiceBuilderOutput) ToBuildServiceBuilderOutputWithContext(ctx context.Context) BuildServiceBuilderOutput {
	return o
}

// The name of the resource.
func (o BuildServiceBuilderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildServiceBuilder) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Property of the Builder resource.
func (o BuildServiceBuilderOutput) Properties() BuilderPropertiesResponseOutput {
	return o.ApplyT(func(v *BuildServiceBuilder) BuilderPropertiesResponseOutput { return v.Properties }).(BuilderPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o BuildServiceBuilderOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BuildServiceBuilder) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o BuildServiceBuilderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildServiceBuilder) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BuildServiceBuilderOutput{})
}
