// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Virtual machine guest diagnostic settings resource.
// Azure REST API version: 2018-06-01-preview. Prior API version in Azure Native 1.x: 2018-06-01-preview.
type GuestDiagnosticsSettingsAssociation struct {
	pulumi.CustomResourceState

	// The guest diagnostic settings name.
	GuestDiagnosticSettingsName pulumi.StringOutput `pulumi:"guestDiagnosticSettingsName"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGuestDiagnosticsSettingsAssociation registers a new resource with the given unique name, arguments, and options.
func NewGuestDiagnosticsSettingsAssociation(ctx *pulumi.Context,
	name string, args *GuestDiagnosticsSettingsAssociationArgs, opts ...pulumi.ResourceOption) (*GuestDiagnosticsSettingsAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GuestDiagnosticSettingsName == nil {
		return nil, errors.New("invalid value for required argument 'GuestDiagnosticSettingsName'")
	}
	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights/v20180601preview:GuestDiagnosticsSettingsAssociation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GuestDiagnosticsSettingsAssociation
	err := ctx.RegisterResource("azure-native:insights:GuestDiagnosticsSettingsAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGuestDiagnosticsSettingsAssociation gets an existing GuestDiagnosticsSettingsAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGuestDiagnosticsSettingsAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestDiagnosticsSettingsAssociationState, opts ...pulumi.ResourceOption) (*GuestDiagnosticsSettingsAssociation, error) {
	var resource GuestDiagnosticsSettingsAssociation
	err := ctx.ReadResource("azure-native:insights:GuestDiagnosticsSettingsAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GuestDiagnosticsSettingsAssociation resources.
type guestDiagnosticsSettingsAssociationState struct {
}

type GuestDiagnosticsSettingsAssociationState struct {
}

func (GuestDiagnosticsSettingsAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestDiagnosticsSettingsAssociationState)(nil)).Elem()
}

type guestDiagnosticsSettingsAssociationArgs struct {
	// The name of the diagnostic settings association.
	AssociationName *string `pulumi:"associationName"`
	// The guest diagnostic settings name.
	GuestDiagnosticSettingsName string `pulumi:"guestDiagnosticSettingsName"`
	// Resource location
	Location *string `pulumi:"location"`
	// The fully qualified ID of the resource, including the resource name and resource type.
	ResourceUri string `pulumi:"resourceUri"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GuestDiagnosticsSettingsAssociation resource.
type GuestDiagnosticsSettingsAssociationArgs struct {
	// The name of the diagnostic settings association.
	AssociationName pulumi.StringPtrInput
	// The guest diagnostic settings name.
	GuestDiagnosticSettingsName pulumi.StringInput
	// Resource location
	Location pulumi.StringPtrInput
	// The fully qualified ID of the resource, including the resource name and resource type.
	ResourceUri pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (GuestDiagnosticsSettingsAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestDiagnosticsSettingsAssociationArgs)(nil)).Elem()
}

type GuestDiagnosticsSettingsAssociationInput interface {
	pulumi.Input

	ToGuestDiagnosticsSettingsAssociationOutput() GuestDiagnosticsSettingsAssociationOutput
	ToGuestDiagnosticsSettingsAssociationOutputWithContext(ctx context.Context) GuestDiagnosticsSettingsAssociationOutput
}

func (*GuestDiagnosticsSettingsAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestDiagnosticsSettingsAssociation)(nil)).Elem()
}

func (i *GuestDiagnosticsSettingsAssociation) ToGuestDiagnosticsSettingsAssociationOutput() GuestDiagnosticsSettingsAssociationOutput {
	return i.ToGuestDiagnosticsSettingsAssociationOutputWithContext(context.Background())
}

func (i *GuestDiagnosticsSettingsAssociation) ToGuestDiagnosticsSettingsAssociationOutputWithContext(ctx context.Context) GuestDiagnosticsSettingsAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestDiagnosticsSettingsAssociationOutput)
}

type GuestDiagnosticsSettingsAssociationOutput struct{ *pulumi.OutputState }

func (GuestDiagnosticsSettingsAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestDiagnosticsSettingsAssociation)(nil)).Elem()
}

func (o GuestDiagnosticsSettingsAssociationOutput) ToGuestDiagnosticsSettingsAssociationOutput() GuestDiagnosticsSettingsAssociationOutput {
	return o
}

func (o GuestDiagnosticsSettingsAssociationOutput) ToGuestDiagnosticsSettingsAssociationOutputWithContext(ctx context.Context) GuestDiagnosticsSettingsAssociationOutput {
	return o
}

// The guest diagnostic settings name.
func (o GuestDiagnosticsSettingsAssociationOutput) GuestDiagnosticSettingsName() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestDiagnosticsSettingsAssociation) pulumi.StringOutput { return v.GuestDiagnosticSettingsName }).(pulumi.StringOutput)
}

// Resource location
func (o GuestDiagnosticsSettingsAssociationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestDiagnosticsSettingsAssociation) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name
func (o GuestDiagnosticsSettingsAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestDiagnosticsSettingsAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource tags
func (o GuestDiagnosticsSettingsAssociationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GuestDiagnosticsSettingsAssociation) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type
func (o GuestDiagnosticsSettingsAssociationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestDiagnosticsSettingsAssociation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GuestDiagnosticsSettingsAssociationOutput{})
}
