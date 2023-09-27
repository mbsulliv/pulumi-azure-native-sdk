// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scom

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A monitored resource.
// Azure REST API version: 2023-07-07-preview.
type MonitoredResource struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of a monitored resource.
	Properties MonitoredResourcePropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMonitoredResource registers a new resource with the given unique name, arguments, and options.
func NewMonitoredResource(ctx *pulumi.Context,
	name string, args *MonitoredResourceArgs, opts ...pulumi.ResourceOption) (*MonitoredResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scom/v20230707preview:MonitoredResource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MonitoredResource
	err := ctx.RegisterResource("azure-native:scom:MonitoredResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitoredResource gets an existing MonitoredResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitoredResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitoredResourceState, opts ...pulumi.ResourceOption) (*MonitoredResource, error) {
	var resource MonitoredResource
	err := ctx.ReadResource("azure-native:scom:MonitoredResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MonitoredResource resources.
type monitoredResourceState struct {
}

type MonitoredResourceState struct {
}

func (MonitoredResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoredResourceState)(nil)).Elem()
}

type monitoredResourceArgs struct {
	// Name of the SCOM managed instance.
	InstanceName string `pulumi:"instanceName"`
	// The monitored resource name.
	MonitoredResourceName *string `pulumi:"monitoredResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a MonitoredResource resource.
type MonitoredResourceArgs struct {
	// Name of the SCOM managed instance.
	InstanceName pulumi.StringInput
	// The monitored resource name.
	MonitoredResourceName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (MonitoredResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoredResourceArgs)(nil)).Elem()
}

type MonitoredResourceInput interface {
	pulumi.Input

	ToMonitoredResourceOutput() MonitoredResourceOutput
	ToMonitoredResourceOutputWithContext(ctx context.Context) MonitoredResourceOutput
}

func (*MonitoredResource) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoredResource)(nil)).Elem()
}

func (i *MonitoredResource) ToMonitoredResourceOutput() MonitoredResourceOutput {
	return i.ToMonitoredResourceOutputWithContext(context.Background())
}

func (i *MonitoredResource) ToMonitoredResourceOutputWithContext(ctx context.Context) MonitoredResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoredResourceOutput)
}

func (i *MonitoredResource) ToOutput(ctx context.Context) pulumix.Output[*MonitoredResource] {
	return pulumix.Output[*MonitoredResource]{
		OutputState: i.ToMonitoredResourceOutputWithContext(ctx).OutputState,
	}
}

type MonitoredResourceOutput struct{ *pulumi.OutputState }

func (MonitoredResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoredResource)(nil)).Elem()
}

func (o MonitoredResourceOutput) ToMonitoredResourceOutput() MonitoredResourceOutput {
	return o
}

func (o MonitoredResourceOutput) ToMonitoredResourceOutputWithContext(ctx context.Context) MonitoredResourceOutput {
	return o
}

func (o MonitoredResourceOutput) ToOutput(ctx context.Context) pulumix.Output[*MonitoredResource] {
	return pulumix.Output[*MonitoredResource]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o MonitoredResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoredResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The properties of a monitored resource.
func (o MonitoredResourceOutput) Properties() MonitoredResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *MonitoredResource) MonitoredResourcePropertiesResponseOutput { return v.Properties }).(MonitoredResourcePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MonitoredResourceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MonitoredResource) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MonitoredResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoredResource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MonitoredResourceOutput{})
}
