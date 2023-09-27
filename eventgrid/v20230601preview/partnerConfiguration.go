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

// Partner configuration information
type PartnerConfiguration struct {
	pulumi.CustomResourceState

	// Location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The details of authorized partners.
	PartnerAuthorization PartnerAuthorizationResponsePtrOutput `pulumi:"partnerAuthorization"`
	// Provisioning state of the partner configuration.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The system metadata relating to partner configuration resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPartnerConfiguration registers a new resource with the given unique name, arguments, and options.
func NewPartnerConfiguration(ctx *pulumi.Context,
	name string, args *PartnerConfigurationArgs, opts ...pulumi.ResourceOption) (*PartnerConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:PartnerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:PartnerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:PartnerConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PartnerConfiguration
	err := ctx.RegisterResource("azure-native:eventgrid/v20230601preview:PartnerConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPartnerConfiguration gets an existing PartnerConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPartnerConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerConfigurationState, opts ...pulumi.ResourceOption) (*PartnerConfiguration, error) {
	var resource PartnerConfiguration
	err := ctx.ReadResource("azure-native:eventgrid/v20230601preview:PartnerConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PartnerConfiguration resources.
type partnerConfigurationState struct {
}

type PartnerConfigurationState struct {
}

func (PartnerConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerConfigurationState)(nil)).Elem()
}

type partnerConfigurationArgs struct {
	// Location of the resource.
	Location *string `pulumi:"location"`
	// The details of authorized partners.
	PartnerAuthorization *PartnerAuthorization `pulumi:"partnerAuthorization"`
	// Provisioning state of the partner configuration.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PartnerConfiguration resource.
type PartnerConfigurationArgs struct {
	// Location of the resource.
	Location pulumi.StringPtrInput
	// The details of authorized partners.
	PartnerAuthorization PartnerAuthorizationPtrInput
	// Provisioning state of the partner configuration.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
}

func (PartnerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerConfigurationArgs)(nil)).Elem()
}

type PartnerConfigurationInput interface {
	pulumi.Input

	ToPartnerConfigurationOutput() PartnerConfigurationOutput
	ToPartnerConfigurationOutputWithContext(ctx context.Context) PartnerConfigurationOutput
}

func (*PartnerConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerConfiguration)(nil)).Elem()
}

func (i *PartnerConfiguration) ToPartnerConfigurationOutput() PartnerConfigurationOutput {
	return i.ToPartnerConfigurationOutputWithContext(context.Background())
}

func (i *PartnerConfiguration) ToPartnerConfigurationOutputWithContext(ctx context.Context) PartnerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerConfigurationOutput)
}

func (i *PartnerConfiguration) ToOutput(ctx context.Context) pulumix.Output[*PartnerConfiguration] {
	return pulumix.Output[*PartnerConfiguration]{
		OutputState: i.ToPartnerConfigurationOutputWithContext(ctx).OutputState,
	}
}

type PartnerConfigurationOutput struct{ *pulumi.OutputState }

func (PartnerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerConfiguration)(nil)).Elem()
}

func (o PartnerConfigurationOutput) ToPartnerConfigurationOutput() PartnerConfigurationOutput {
	return o
}

func (o PartnerConfigurationOutput) ToPartnerConfigurationOutputWithContext(ctx context.Context) PartnerConfigurationOutput {
	return o
}

func (o PartnerConfigurationOutput) ToOutput(ctx context.Context) pulumix.Output[*PartnerConfiguration] {
	return pulumix.Output[*PartnerConfiguration]{
		OutputState: o.OutputState,
	}
}

// Location of the resource.
func (o PartnerConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerConfiguration) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o PartnerConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The details of authorized partners.
func (o PartnerConfigurationOutput) PartnerAuthorization() PartnerAuthorizationResponsePtrOutput {
	return o.ApplyT(func(v *PartnerConfiguration) PartnerAuthorizationResponsePtrOutput { return v.PartnerAuthorization }).(PartnerAuthorizationResponsePtrOutput)
}

// Provisioning state of the partner configuration.
func (o PartnerConfigurationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerConfiguration) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The system metadata relating to partner configuration resource.
func (o PartnerConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PartnerConfiguration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags of the resource.
func (o PartnerConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PartnerConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the resource.
func (o PartnerConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PartnerConfigurationOutput{})
}
