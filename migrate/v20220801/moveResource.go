// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines the move resource.
type MoveResource struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Defines the move resource properties.
	Properties MoveResourcePropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMoveResource registers a new resource with the given unique name, arguments, and options.
func NewMoveResource(ctx *pulumi.Context,
	name string, args *MoveResourceArgs, opts ...pulumi.ResourceOption) (*MoveResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MoveCollectionName == nil {
		return nil, errors.New("invalid value for required argument 'MoveCollectionName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate:MoveResource"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20191001preview:MoveResource"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20210101:MoveResource"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20210801:MoveResource"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230801:MoveResource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MoveResource
	err := ctx.RegisterResource("azure-native:migrate/v20220801:MoveResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMoveResource gets an existing MoveResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMoveResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MoveResourceState, opts ...pulumi.ResourceOption) (*MoveResource, error) {
	var resource MoveResource
	err := ctx.ReadResource("azure-native:migrate/v20220801:MoveResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MoveResource resources.
type moveResourceState struct {
}

type MoveResourceState struct {
}

func (MoveResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*moveResourceState)(nil)).Elem()
}

type moveResourceArgs struct {
	// The Move Collection Name.
	MoveCollectionName string `pulumi:"moveCollectionName"`
	// The Move Resource Name.
	MoveResourceName *string `pulumi:"moveResourceName"`
	// Defines the move resource properties.
	Properties *MoveResourceProperties `pulumi:"properties"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a MoveResource resource.
type MoveResourceArgs struct {
	// The Move Collection Name.
	MoveCollectionName pulumi.StringInput
	// The Move Resource Name.
	MoveResourceName pulumi.StringPtrInput
	// Defines the move resource properties.
	Properties MoveResourcePropertiesPtrInput
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput
}

func (MoveResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*moveResourceArgs)(nil)).Elem()
}

type MoveResourceInput interface {
	pulumi.Input

	ToMoveResourceOutput() MoveResourceOutput
	ToMoveResourceOutputWithContext(ctx context.Context) MoveResourceOutput
}

func (*MoveResource) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResource)(nil)).Elem()
}

func (i *MoveResource) ToMoveResourceOutput() MoveResourceOutput {
	return i.ToMoveResourceOutputWithContext(context.Background())
}

func (i *MoveResource) ToMoveResourceOutputWithContext(ctx context.Context) MoveResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceOutput)
}

type MoveResourceOutput struct{ *pulumi.OutputState }

func (MoveResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResource)(nil)).Elem()
}

func (o MoveResourceOutput) ToMoveResourceOutput() MoveResourceOutput {
	return o
}

func (o MoveResourceOutput) ToMoveResourceOutputWithContext(ctx context.Context) MoveResourceOutput {
	return o
}

// The name of the resource
func (o MoveResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MoveResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defines the move resource properties.
func (o MoveResourceOutput) Properties() MoveResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *MoveResource) MoveResourcePropertiesResponseOutput { return v.Properties }).(MoveResourcePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o MoveResourceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MoveResource) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o MoveResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MoveResource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MoveResourceOutput{})
}
