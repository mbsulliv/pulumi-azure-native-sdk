// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkAssociation struct {
	pulumi.CustomResourceState

	// The pla name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The private link association properties.
	Properties PrivateLinkAssociationPropertiesExpandedResponseOutput `pulumi:"properties"`
	// The operation type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateLinkAssociation registers a new resource with the given unique name, arguments, and options.
func NewPrivateLinkAssociation(ctx *pulumi.Context,
	name string, args *PrivateLinkAssociationArgs, opts ...pulumi.ResourceOption) (*PrivateLinkAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:PrivateLinkAssociation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateLinkAssociation
	err := ctx.RegisterResource("azure-native:authorization/v20200501:PrivateLinkAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateLinkAssociation gets an existing PrivateLinkAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateLinkAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkAssociationState, opts ...pulumi.ResourceOption) (*PrivateLinkAssociation, error) {
	var resource PrivateLinkAssociation
	err := ctx.ReadResource("azure-native:authorization/v20200501:PrivateLinkAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateLinkAssociation resources.
type privateLinkAssociationState struct {
}

type PrivateLinkAssociationState struct {
}

func (PrivateLinkAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkAssociationState)(nil)).Elem()
}

type privateLinkAssociationArgs struct {
	// The management group ID.
	GroupId string `pulumi:"groupId"`
	// The ID of the PLA
	PlaId *string `pulumi:"plaId"`
	// The properties of the PrivateLinkAssociation.
	Properties *PrivateLinkAssociationProperties `pulumi:"properties"`
}

// The set of arguments for constructing a PrivateLinkAssociation resource.
type PrivateLinkAssociationArgs struct {
	// The management group ID.
	GroupId pulumi.StringInput
	// The ID of the PLA
	PlaId pulumi.StringPtrInput
	// The properties of the PrivateLinkAssociation.
	Properties PrivateLinkAssociationPropertiesPtrInput
}

func (PrivateLinkAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkAssociationArgs)(nil)).Elem()
}

type PrivateLinkAssociationInput interface {
	pulumi.Input

	ToPrivateLinkAssociationOutput() PrivateLinkAssociationOutput
	ToPrivateLinkAssociationOutputWithContext(ctx context.Context) PrivateLinkAssociationOutput
}

func (*PrivateLinkAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkAssociation)(nil)).Elem()
}

func (i *PrivateLinkAssociation) ToPrivateLinkAssociationOutput() PrivateLinkAssociationOutput {
	return i.ToPrivateLinkAssociationOutputWithContext(context.Background())
}

func (i *PrivateLinkAssociation) ToPrivateLinkAssociationOutputWithContext(ctx context.Context) PrivateLinkAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkAssociationOutput)
}

type PrivateLinkAssociationOutput struct{ *pulumi.OutputState }

func (PrivateLinkAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkAssociation)(nil)).Elem()
}

func (o PrivateLinkAssociationOutput) ToPrivateLinkAssociationOutput() PrivateLinkAssociationOutput {
	return o
}

func (o PrivateLinkAssociationOutput) ToPrivateLinkAssociationOutputWithContext(ctx context.Context) PrivateLinkAssociationOutput {
	return o
}

// The pla name.
func (o PrivateLinkAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The private link association properties.
func (o PrivateLinkAssociationOutput) Properties() PrivateLinkAssociationPropertiesExpandedResponseOutput {
	return o.ApplyT(func(v *PrivateLinkAssociation) PrivateLinkAssociationPropertiesExpandedResponseOutput {
		return v.Properties
	}).(PrivateLinkAssociationPropertiesExpandedResponseOutput)
}

// The operation type.
func (o PrivateLinkAssociationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkAssociation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkAssociationOutput{})
}
