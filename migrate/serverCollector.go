// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure REST API version: 2019-10-01. Prior API version in Azure Native 1.x: 2019-10-01.
type ServerCollector struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput            `pulumi:"eTag"`
	Name       pulumi.StringOutput               `pulumi:"name"`
	Properties CollectorPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput               `pulumi:"type"`
}

// NewServerCollector registers a new resource with the given unique name, arguments, and options.
func NewServerCollector(ctx *pulumi.Context,
	name string, args *ServerCollectorArgs, opts ...pulumi.ResourceOption) (*ServerCollector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20191001:ServerCollector"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230315:ServerCollector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServerCollector
	err := ctx.RegisterResource("azure-native:migrate:ServerCollector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerCollector gets an existing ServerCollector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerCollector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerCollectorState, opts ...pulumi.ResourceOption) (*ServerCollector, error) {
	var resource ServerCollector
	err := ctx.ReadResource("azure-native:migrate:ServerCollector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerCollector resources.
type serverCollectorState struct {
}

type ServerCollectorState struct {
}

func (ServerCollectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverCollectorState)(nil)).Elem()
}

type serverCollectorArgs struct {
	ETag *string `pulumi:"eTag"`
	// Name of the Azure Migrate project.
	ProjectName string               `pulumi:"projectName"`
	Properties  *CollectorProperties `pulumi:"properties"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Unique name of a Server collector within a project.
	ServerCollectorName *string `pulumi:"serverCollectorName"`
}

// The set of arguments for constructing a ServerCollector resource.
type ServerCollectorArgs struct {
	ETag pulumi.StringPtrInput
	// Name of the Azure Migrate project.
	ProjectName pulumi.StringInput
	Properties  CollectorPropertiesPtrInput
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput
	// Unique name of a Server collector within a project.
	ServerCollectorName pulumi.StringPtrInput
}

func (ServerCollectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverCollectorArgs)(nil)).Elem()
}

type ServerCollectorInput interface {
	pulumi.Input

	ToServerCollectorOutput() ServerCollectorOutput
	ToServerCollectorOutputWithContext(ctx context.Context) ServerCollectorOutput
}

func (*ServerCollector) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerCollector)(nil)).Elem()
}

func (i *ServerCollector) ToServerCollectorOutput() ServerCollectorOutput {
	return i.ToServerCollectorOutputWithContext(context.Background())
}

func (i *ServerCollector) ToServerCollectorOutputWithContext(ctx context.Context) ServerCollectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCollectorOutput)
}

func (i *ServerCollector) ToOutput(ctx context.Context) pulumix.Output[*ServerCollector] {
	return pulumix.Output[*ServerCollector]{
		OutputState: i.ToServerCollectorOutputWithContext(ctx).OutputState,
	}
}

type ServerCollectorOutput struct{ *pulumi.OutputState }

func (ServerCollectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerCollector)(nil)).Elem()
}

func (o ServerCollectorOutput) ToServerCollectorOutput() ServerCollectorOutput {
	return o
}

func (o ServerCollectorOutput) ToServerCollectorOutputWithContext(ctx context.Context) ServerCollectorOutput {
	return o
}

func (o ServerCollectorOutput) ToOutput(ctx context.Context) pulumix.Output[*ServerCollector] {
	return pulumix.Output[*ServerCollector]{
		OutputState: o.OutputState,
	}
}

func (o ServerCollectorOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCollector) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o ServerCollectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerCollector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerCollectorOutput) Properties() CollectorPropertiesResponseOutput {
	return o.ApplyT(func(v *ServerCollector) CollectorPropertiesResponseOutput { return v.Properties }).(CollectorPropertiesResponseOutput)
}

func (o ServerCollectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerCollector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerCollectorOutput{})
}
