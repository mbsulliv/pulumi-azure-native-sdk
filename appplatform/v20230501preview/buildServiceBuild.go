// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Build resource payload
type BuildServiceBuild struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the build resource
	Properties BuildPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBuildServiceBuild registers a new resource with the given unique name, arguments, and options.
func NewBuildServiceBuild(ctx *pulumi.Context,
	name string, args *BuildServiceBuildArgs, opts ...pulumi.ResourceOption) (*BuildServiceBuild, error) {
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
	if args.Properties != nil {
		args.Properties = args.Properties.ToBuildPropertiesPtrOutput().ApplyT(func(v *BuildProperties) *BuildProperties { return v.Defaults() }).(BuildPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:BuildServiceBuild"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230301preview:BuildServiceBuild"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230701preview:BuildServiceBuild"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230901preview:BuildServiceBuild"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20231101preview:BuildServiceBuild"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BuildServiceBuild
	err := ctx.RegisterResource("azure-native:appplatform/v20230501preview:BuildServiceBuild", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBuildServiceBuild gets an existing BuildServiceBuild resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuildServiceBuild(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildServiceBuildState, opts ...pulumi.ResourceOption) (*BuildServiceBuild, error) {
	var resource BuildServiceBuild
	err := ctx.ReadResource("azure-native:appplatform/v20230501preview:BuildServiceBuild", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BuildServiceBuild resources.
type buildServiceBuildState struct {
}

type BuildServiceBuildState struct {
}

func (BuildServiceBuildState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildServiceBuildState)(nil)).Elem()
}

type buildServiceBuildArgs struct {
	// The name of the build resource.
	BuildName *string `pulumi:"buildName"`
	// The name of the build service resource.
	BuildServiceName string `pulumi:"buildServiceName"`
	// Properties of the build resource
	Properties *BuildProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a BuildServiceBuild resource.
type BuildServiceBuildArgs struct {
	// The name of the build resource.
	BuildName pulumi.StringPtrInput
	// The name of the build service resource.
	BuildServiceName pulumi.StringInput
	// Properties of the build resource
	Properties BuildPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (BuildServiceBuildArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildServiceBuildArgs)(nil)).Elem()
}

type BuildServiceBuildInput interface {
	pulumi.Input

	ToBuildServiceBuildOutput() BuildServiceBuildOutput
	ToBuildServiceBuildOutputWithContext(ctx context.Context) BuildServiceBuildOutput
}

func (*BuildServiceBuild) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildServiceBuild)(nil)).Elem()
}

func (i *BuildServiceBuild) ToBuildServiceBuildOutput() BuildServiceBuildOutput {
	return i.ToBuildServiceBuildOutputWithContext(context.Background())
}

func (i *BuildServiceBuild) ToBuildServiceBuildOutputWithContext(ctx context.Context) BuildServiceBuildOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildServiceBuildOutput)
}

type BuildServiceBuildOutput struct{ *pulumi.OutputState }

func (BuildServiceBuildOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildServiceBuild)(nil)).Elem()
}

func (o BuildServiceBuildOutput) ToBuildServiceBuildOutput() BuildServiceBuildOutput {
	return o
}

func (o BuildServiceBuildOutput) ToBuildServiceBuildOutputWithContext(ctx context.Context) BuildServiceBuildOutput {
	return o
}

// The name of the resource.
func (o BuildServiceBuildOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildServiceBuild) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the build resource
func (o BuildServiceBuildOutput) Properties() BuildPropertiesResponseOutput {
	return o.ApplyT(func(v *BuildServiceBuild) BuildPropertiesResponseOutput { return v.Properties }).(BuildPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o BuildServiceBuildOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BuildServiceBuild) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o BuildServiceBuildOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildServiceBuild) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BuildServiceBuildOutput{})
}
