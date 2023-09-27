// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The build service agent pool resource
type BuildServiceAgentPool struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// build service agent pool properties
	Properties BuildServiceAgentPoolPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBuildServiceAgentPool registers a new resource with the given unique name, arguments, and options.
func NewBuildServiceAgentPool(ctx *pulumi.Context,
	name string, args *BuildServiceAgentPoolArgs, opts ...pulumi.ResourceOption) (*BuildServiceAgentPool, error) {
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
			Type: pulumi.String("azure-native:appplatform:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230301preview:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230501preview:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230701preview:BuildServiceAgentPool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BuildServiceAgentPool
	err := ctx.RegisterResource("azure-native:appplatform/v20230901preview:BuildServiceAgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBuildServiceAgentPool gets an existing BuildServiceAgentPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuildServiceAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildServiceAgentPoolState, opts ...pulumi.ResourceOption) (*BuildServiceAgentPool, error) {
	var resource BuildServiceAgentPool
	err := ctx.ReadResource("azure-native:appplatform/v20230901preview:BuildServiceAgentPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BuildServiceAgentPool resources.
type buildServiceAgentPoolState struct {
}

type BuildServiceAgentPoolState struct {
}

func (BuildServiceAgentPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildServiceAgentPoolState)(nil)).Elem()
}

type buildServiceAgentPoolArgs struct {
	// The name of the build service agent pool resource.
	AgentPoolName *string `pulumi:"agentPoolName"`
	// The name of the build service resource.
	BuildServiceName string `pulumi:"buildServiceName"`
	// build service agent pool properties
	Properties *BuildServiceAgentPoolProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a BuildServiceAgentPool resource.
type BuildServiceAgentPoolArgs struct {
	// The name of the build service agent pool resource.
	AgentPoolName pulumi.StringPtrInput
	// The name of the build service resource.
	BuildServiceName pulumi.StringInput
	// build service agent pool properties
	Properties BuildServiceAgentPoolPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (BuildServiceAgentPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildServiceAgentPoolArgs)(nil)).Elem()
}

type BuildServiceAgentPoolInput interface {
	pulumi.Input

	ToBuildServiceAgentPoolOutput() BuildServiceAgentPoolOutput
	ToBuildServiceAgentPoolOutputWithContext(ctx context.Context) BuildServiceAgentPoolOutput
}

func (*BuildServiceAgentPool) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildServiceAgentPool)(nil)).Elem()
}

func (i *BuildServiceAgentPool) ToBuildServiceAgentPoolOutput() BuildServiceAgentPoolOutput {
	return i.ToBuildServiceAgentPoolOutputWithContext(context.Background())
}

func (i *BuildServiceAgentPool) ToBuildServiceAgentPoolOutputWithContext(ctx context.Context) BuildServiceAgentPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildServiceAgentPoolOutput)
}

func (i *BuildServiceAgentPool) ToOutput(ctx context.Context) pulumix.Output[*BuildServiceAgentPool] {
	return pulumix.Output[*BuildServiceAgentPool]{
		OutputState: i.ToBuildServiceAgentPoolOutputWithContext(ctx).OutputState,
	}
}

type BuildServiceAgentPoolOutput struct{ *pulumi.OutputState }

func (BuildServiceAgentPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildServiceAgentPool)(nil)).Elem()
}

func (o BuildServiceAgentPoolOutput) ToBuildServiceAgentPoolOutput() BuildServiceAgentPoolOutput {
	return o
}

func (o BuildServiceAgentPoolOutput) ToBuildServiceAgentPoolOutputWithContext(ctx context.Context) BuildServiceAgentPoolOutput {
	return o
}

func (o BuildServiceAgentPoolOutput) ToOutput(ctx context.Context) pulumix.Output[*BuildServiceAgentPool] {
	return pulumix.Output[*BuildServiceAgentPool]{
		OutputState: o.OutputState,
	}
}

// The name of the resource.
func (o BuildServiceAgentPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildServiceAgentPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// build service agent pool properties
func (o BuildServiceAgentPoolOutput) Properties() BuildServiceAgentPoolPropertiesResponseOutput {
	return o.ApplyT(func(v *BuildServiceAgentPool) BuildServiceAgentPoolPropertiesResponseOutput { return v.Properties }).(BuildServiceAgentPoolPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o BuildServiceAgentPoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BuildServiceAgentPool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o BuildServiceAgentPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildServiceAgentPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BuildServiceAgentPoolOutput{})
}
