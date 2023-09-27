// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Information about a partner registration.
type PartnerRegistration struct {
	pulumi.CustomResourceState

	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The immutableId of the corresponding partner registration.
	// Note: This property is marked for deprecation and is not supported in any future GA API version
	PartnerRegistrationImmutableId pulumi.StringPtrOutput `pulumi:"partnerRegistrationImmutableId"`
	// Provisioning state of the partner registration.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system metadata relating to Partner Registration resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPartnerRegistration registers a new resource with the given unique name, arguments, and options.
func NewPartnerRegistration(ctx *pulumi.Context,
	name string, args *PartnerRegistrationArgs, opts ...pulumi.ResourceOption) (*PartnerRegistration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:PartnerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:PartnerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:PartnerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:PartnerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:PartnerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:PartnerRegistration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PartnerRegistration
	err := ctx.RegisterResource("azure-native:eventgrid/v20230601preview:PartnerRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPartnerRegistration gets an existing PartnerRegistration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPartnerRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerRegistrationState, opts ...pulumi.ResourceOption) (*PartnerRegistration, error) {
	var resource PartnerRegistration
	err := ctx.ReadResource("azure-native:eventgrid/v20230601preview:PartnerRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PartnerRegistration resources.
type partnerRegistrationState struct {
}

type PartnerRegistrationState struct {
}

func (PartnerRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerRegistrationState)(nil)).Elem()
}

type partnerRegistrationArgs struct {
	// Location of the resource.
	Location *string `pulumi:"location"`
	// The immutableId of the corresponding partner registration.
	// Note: This property is marked for deprecation and is not supported in any future GA API version
	PartnerRegistrationImmutableId *string `pulumi:"partnerRegistrationImmutableId"`
	// Name of the partner registration.
	PartnerRegistrationName *string `pulumi:"partnerRegistrationName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PartnerRegistration resource.
type PartnerRegistrationArgs struct {
	// Location of the resource.
	Location pulumi.StringPtrInput
	// The immutableId of the corresponding partner registration.
	// Note: This property is marked for deprecation and is not supported in any future GA API version
	PartnerRegistrationImmutableId pulumi.StringPtrInput
	// Name of the partner registration.
	PartnerRegistrationName pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
}

func (PartnerRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerRegistrationArgs)(nil)).Elem()
}

type PartnerRegistrationInput interface {
	pulumi.Input

	ToPartnerRegistrationOutput() PartnerRegistrationOutput
	ToPartnerRegistrationOutputWithContext(ctx context.Context) PartnerRegistrationOutput
}

func (*PartnerRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerRegistration)(nil)).Elem()
}

func (i *PartnerRegistration) ToPartnerRegistrationOutput() PartnerRegistrationOutput {
	return i.ToPartnerRegistrationOutputWithContext(context.Background())
}

func (i *PartnerRegistration) ToPartnerRegistrationOutputWithContext(ctx context.Context) PartnerRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerRegistrationOutput)
}

func (i *PartnerRegistration) ToOutput(ctx context.Context) pulumix.Output[*PartnerRegistration] {
	return pulumix.Output[*PartnerRegistration]{
		OutputState: i.ToPartnerRegistrationOutputWithContext(ctx).OutputState,
	}
}

type PartnerRegistrationOutput struct{ *pulumi.OutputState }

func (PartnerRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerRegistration)(nil)).Elem()
}

func (o PartnerRegistrationOutput) ToPartnerRegistrationOutput() PartnerRegistrationOutput {
	return o
}

func (o PartnerRegistrationOutput) ToPartnerRegistrationOutputWithContext(ctx context.Context) PartnerRegistrationOutput {
	return o
}

func (o PartnerRegistrationOutput) ToOutput(ctx context.Context) pulumix.Output[*PartnerRegistration] {
	return pulumix.Output[*PartnerRegistration]{
		OutputState: o.OutputState,
	}
}

// Location of the resource.
func (o PartnerRegistrationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the resource.
func (o PartnerRegistrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The immutableId of the corresponding partner registration.
// Note: This property is marked for deprecation and is not supported in any future GA API version
func (o PartnerRegistrationOutput) PartnerRegistrationImmutableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringPtrOutput { return v.PartnerRegistrationImmutableId }).(pulumi.StringPtrOutput)
}

// Provisioning state of the partner registration.
func (o PartnerRegistrationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system metadata relating to Partner Registration resource.
func (o PartnerRegistrationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PartnerRegistration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags of the resource.
func (o PartnerRegistrationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the resource.
func (o PartnerRegistrationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PartnerRegistrationOutput{})
}
