// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Integration runtime resource type.
type IntegrationRuntime struct {
	pulumi.CustomResourceState

	// Etag identifies change in the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Integration runtime properties.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegrationRuntime registers a new resource with the given unique name, arguments, and options.
func NewIntegrationRuntime(ctx *pulumi.Context,
	name string, args *IntegrationRuntimeArgs, opts ...pulumi.ResourceOption) (*IntegrationRuntime, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FactoryName == nil {
		return nil, errors.New("invalid value for required argument 'FactoryName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datafactory:IntegrationRuntime"),
		},
		{
			Type: pulumi.String("azure-native:datafactory/v20170901preview:IntegrationRuntime"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IntegrationRuntime
	err := ctx.RegisterResource("azure-native:datafactory/v20180601:IntegrationRuntime", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationRuntime gets an existing IntegrationRuntime resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationRuntime(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationRuntimeState, opts ...pulumi.ResourceOption) (*IntegrationRuntime, error) {
	var resource IntegrationRuntime
	err := ctx.ReadResource("azure-native:datafactory/v20180601:IntegrationRuntime", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationRuntime resources.
type integrationRuntimeState struct {
}

type IntegrationRuntimeState struct {
}

func (IntegrationRuntimeState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRuntimeState)(nil)).Elem()
}

type integrationRuntimeArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The integration runtime name.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Integration runtime properties.
	Properties interface{} `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a IntegrationRuntime resource.
type IntegrationRuntimeArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput
	// The integration runtime name.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Integration runtime properties.
	Properties pulumi.Input
	// The resource group name.
	ResourceGroupName pulumi.StringInput
}

func (IntegrationRuntimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRuntimeArgs)(nil)).Elem()
}

type IntegrationRuntimeInput interface {
	pulumi.Input

	ToIntegrationRuntimeOutput() IntegrationRuntimeOutput
	ToIntegrationRuntimeOutputWithContext(ctx context.Context) IntegrationRuntimeOutput
}

func (*IntegrationRuntime) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRuntime)(nil)).Elem()
}

func (i *IntegrationRuntime) ToIntegrationRuntimeOutput() IntegrationRuntimeOutput {
	return i.ToIntegrationRuntimeOutputWithContext(context.Background())
}

func (i *IntegrationRuntime) ToIntegrationRuntimeOutputWithContext(ctx context.Context) IntegrationRuntimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeOutput)
}

type IntegrationRuntimeOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRuntime)(nil)).Elem()
}

func (o IntegrationRuntimeOutput) ToIntegrationRuntimeOutput() IntegrationRuntimeOutput {
	return o
}

func (o IntegrationRuntimeOutput) ToIntegrationRuntimeOutputWithContext(ctx context.Context) IntegrationRuntimeOutput {
	return o
}

// Etag identifies change in the resource.
func (o IntegrationRuntimeOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRuntime) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The resource name.
func (o IntegrationRuntimeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRuntime) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Integration runtime properties.
func (o IntegrationRuntimeOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *IntegrationRuntime) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// The resource type.
func (o IntegrationRuntimeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRuntime) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationRuntimeOutput{})
}
