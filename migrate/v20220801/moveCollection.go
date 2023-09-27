// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Define the move collection.
type MoveCollection struct {
	pulumi.CustomResourceState

	// The etag of the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Defines the MSI properties of the Move Collection.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Defines the move collection properties.
	Properties MoveCollectionPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMoveCollection registers a new resource with the given unique name, arguments, and options.
func NewMoveCollection(ctx *pulumi.Context,
	name string, args *MoveCollectionArgs, opts ...pulumi.ResourceOption) (*MoveCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate:MoveCollection"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20191001preview:MoveCollection"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20210101:MoveCollection"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20210801:MoveCollection"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230801:MoveCollection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MoveCollection
	err := ctx.RegisterResource("azure-native:migrate/v20220801:MoveCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMoveCollection gets an existing MoveCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMoveCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MoveCollectionState, opts ...pulumi.ResourceOption) (*MoveCollection, error) {
	var resource MoveCollection
	err := ctx.ReadResource("azure-native:migrate/v20220801:MoveCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MoveCollection resources.
type moveCollectionState struct {
}

type MoveCollectionState struct {
}

func (MoveCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*moveCollectionState)(nil)).Elem()
}

type moveCollectionArgs struct {
	// Defines the MSI properties of the Move Collection.
	Identity *Identity `pulumi:"identity"`
	// The geo-location where the resource lives.
	Location *string `pulumi:"location"`
	// The Move Collection Name.
	MoveCollectionName *string `pulumi:"moveCollectionName"`
	// Defines the move collection properties.
	Properties *MoveCollectionProperties `pulumi:"properties"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MoveCollection resource.
type MoveCollectionArgs struct {
	// Defines the MSI properties of the Move Collection.
	Identity IdentityPtrInput
	// The geo-location where the resource lives.
	Location pulumi.StringPtrInput
	// The Move Collection Name.
	MoveCollectionName pulumi.StringPtrInput
	// Defines the move collection properties.
	Properties MoveCollectionPropertiesPtrInput
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (MoveCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*moveCollectionArgs)(nil)).Elem()
}

type MoveCollectionInput interface {
	pulumi.Input

	ToMoveCollectionOutput() MoveCollectionOutput
	ToMoveCollectionOutputWithContext(ctx context.Context) MoveCollectionOutput
}

func (*MoveCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollection)(nil)).Elem()
}

func (i *MoveCollection) ToMoveCollectionOutput() MoveCollectionOutput {
	return i.ToMoveCollectionOutputWithContext(context.Background())
}

func (i *MoveCollection) ToMoveCollectionOutputWithContext(ctx context.Context) MoveCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionOutput)
}

func (i *MoveCollection) ToOutput(ctx context.Context) pulumix.Output[*MoveCollection] {
	return pulumix.Output[*MoveCollection]{
		OutputState: i.ToMoveCollectionOutputWithContext(ctx).OutputState,
	}
}

type MoveCollectionOutput struct{ *pulumi.OutputState }

func (MoveCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollection)(nil)).Elem()
}

func (o MoveCollectionOutput) ToMoveCollectionOutput() MoveCollectionOutput {
	return o
}

func (o MoveCollectionOutput) ToMoveCollectionOutputWithContext(ctx context.Context) MoveCollectionOutput {
	return o
}

func (o MoveCollectionOutput) ToOutput(ctx context.Context) pulumix.Output[*MoveCollection] {
	return pulumix.Output[*MoveCollection]{
		OutputState: o.OutputState,
	}
}

// The etag of the resource.
func (o MoveCollectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *MoveCollection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Defines the MSI properties of the Move Collection.
func (o MoveCollectionOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *MoveCollection) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The geo-location where the resource lives.
func (o MoveCollectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveCollection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o MoveCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MoveCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defines the move collection properties.
func (o MoveCollectionOutput) Properties() MoveCollectionPropertiesResponseOutput {
	return o.ApplyT(func(v *MoveCollection) MoveCollectionPropertiesResponseOutput { return v.Properties }).(MoveCollectionPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o MoveCollectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MoveCollection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o MoveCollectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MoveCollection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o MoveCollectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MoveCollection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MoveCollectionOutput{})
}
