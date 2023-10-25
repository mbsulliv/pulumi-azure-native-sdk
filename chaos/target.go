// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package chaos

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Model that represents a Target resource.
// Azure REST API version: 2023-04-15-preview. Prior API version in Azure Native 1.x: 2021-09-15-preview.
//
// Other available API versions: 2023-09-01-preview, 2023-10-27-preview.
type Target struct {
	pulumi.CustomResourceState

	// Location of the target resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the target resource.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The system metadata of the target resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTarget registers a new resource with the given unique name, arguments, and options.
func NewTarget(ctx *pulumi.Context,
	name string, args *TargetArgs, opts ...pulumi.ResourceOption) (*Target, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ParentProviderNamespace'")
	}
	if args.ParentResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourceName'")
	}
	if args.ParentResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourceType'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:chaos/v20210915preview:Target"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20220701preview:Target"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20221001preview:Target"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20230401preview:Target"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20230415preview:Target"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20230901preview:Target"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20231027preview:Target"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Target
	err := ctx.RegisterResource("azure-native:chaos:Target", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTarget gets an existing Target resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetState, opts ...pulumi.ResourceOption) (*Target, error) {
	var resource Target
	err := ctx.ReadResource("azure-native:chaos:Target", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Target resources.
type targetState struct {
}

type TargetState struct {
}

func (TargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetState)(nil)).Elem()
}

type targetArgs struct {
	// Location of the target resource.
	Location *string `pulumi:"location"`
	// String that represents a resource provider namespace.
	ParentProviderNamespace string `pulumi:"parentProviderNamespace"`
	// String that represents a resource name.
	ParentResourceName string `pulumi:"parentResourceName"`
	// String that represents a resource type.
	ParentResourceType string `pulumi:"parentResourceType"`
	// The properties of the target resource.
	Properties interface{} `pulumi:"properties"`
	// String that represents an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// String that represents a Target resource name.
	TargetName *string `pulumi:"targetName"`
}

// The set of arguments for constructing a Target resource.
type TargetArgs struct {
	// Location of the target resource.
	Location pulumi.StringPtrInput
	// String that represents a resource provider namespace.
	ParentProviderNamespace pulumi.StringInput
	// String that represents a resource name.
	ParentResourceName pulumi.StringInput
	// String that represents a resource type.
	ParentResourceType pulumi.StringInput
	// The properties of the target resource.
	Properties pulumi.Input
	// String that represents an Azure resource group.
	ResourceGroupName pulumi.StringInput
	// String that represents a Target resource name.
	TargetName pulumi.StringPtrInput
}

func (TargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetArgs)(nil)).Elem()
}

type TargetInput interface {
	pulumi.Input

	ToTargetOutput() TargetOutput
	ToTargetOutputWithContext(ctx context.Context) TargetOutput
}

func (*Target) ElementType() reflect.Type {
	return reflect.TypeOf((**Target)(nil)).Elem()
}

func (i *Target) ToTargetOutput() TargetOutput {
	return i.ToTargetOutputWithContext(context.Background())
}

func (i *Target) ToTargetOutputWithContext(ctx context.Context) TargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetOutput)
}

func (i *Target) ToOutput(ctx context.Context) pulumix.Output[*Target] {
	return pulumix.Output[*Target]{
		OutputState: i.ToTargetOutputWithContext(ctx).OutputState,
	}
}

type TargetOutput struct{ *pulumi.OutputState }

func (TargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Target)(nil)).Elem()
}

func (o TargetOutput) ToTargetOutput() TargetOutput {
	return o
}

func (o TargetOutput) ToTargetOutputWithContext(ctx context.Context) TargetOutput {
	return o
}

func (o TargetOutput) ToOutput(ctx context.Context) pulumix.Output[*Target] {
	return pulumix.Output[*Target]{
		OutputState: o.OutputState,
	}
}

// Location of the target resource.
func (o TargetOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Target) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o TargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Target) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The properties of the target resource.
func (o TargetOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Target) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// The system metadata of the target resource.
func (o TargetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Target) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o TargetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Target) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TargetOutput{})
}
