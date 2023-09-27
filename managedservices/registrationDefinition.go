// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managedservices

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The registration definition.
// Azure REST API version: 2022-10-01. Prior API version in Azure Native 1.x: 2019-09-01
type RegistrationDefinition struct {
	pulumi.CustomResourceState

	// The name of the registration definition.
	Name pulumi.StringOutput `pulumi:"name"`
	// The details for the Managed Services offer’s plan in Azure Marketplace.
	Plan PlanResponsePtrOutput `pulumi:"plan"`
	// The properties of a registration definition.
	Properties RegistrationDefinitionPropertiesResponseOutput `pulumi:"properties"`
	// The metadata for the registration assignment resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the Azure resource (Microsoft.ManagedServices/registrationDefinitions).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistrationDefinition registers a new resource with the given unique name, arguments, and options.
func NewRegistrationDefinition(ctx *pulumi.Context,
	name string, args *RegistrationDefinitionArgs, opts ...pulumi.ResourceOption) (*RegistrationDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managedservices/v20180601preview:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20190401preview:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20190601:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20190901:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20200201preview:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20220101preview:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20221001:RegistrationDefinition"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RegistrationDefinition
	err := ctx.RegisterResource("azure-native:managedservices:RegistrationDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistrationDefinition gets an existing RegistrationDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistrationDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistrationDefinitionState, opts ...pulumi.ResourceOption) (*RegistrationDefinition, error) {
	var resource RegistrationDefinition
	err := ctx.ReadResource("azure-native:managedservices:RegistrationDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistrationDefinition resources.
type registrationDefinitionState struct {
}

type RegistrationDefinitionState struct {
}

func (RegistrationDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationDefinitionState)(nil)).Elem()
}

type registrationDefinitionArgs struct {
	// The details for the Managed Services offer’s plan in Azure Marketplace.
	Plan *Plan `pulumi:"plan"`
	// The properties of a registration definition.
	Properties *RegistrationDefinitionProperties `pulumi:"properties"`
	// The GUID of the registration definition.
	RegistrationDefinitionId *string `pulumi:"registrationDefinitionId"`
	// The scope of the resource.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a RegistrationDefinition resource.
type RegistrationDefinitionArgs struct {
	// The details for the Managed Services offer’s plan in Azure Marketplace.
	Plan PlanPtrInput
	// The properties of a registration definition.
	Properties RegistrationDefinitionPropertiesPtrInput
	// The GUID of the registration definition.
	RegistrationDefinitionId pulumi.StringPtrInput
	// The scope of the resource.
	Scope pulumi.StringInput
}

func (RegistrationDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationDefinitionArgs)(nil)).Elem()
}

type RegistrationDefinitionInput interface {
	pulumi.Input

	ToRegistrationDefinitionOutput() RegistrationDefinitionOutput
	ToRegistrationDefinitionOutputWithContext(ctx context.Context) RegistrationDefinitionOutput
}

func (*RegistrationDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationDefinition)(nil)).Elem()
}

func (i *RegistrationDefinition) ToRegistrationDefinitionOutput() RegistrationDefinitionOutput {
	return i.ToRegistrationDefinitionOutputWithContext(context.Background())
}

func (i *RegistrationDefinition) ToRegistrationDefinitionOutputWithContext(ctx context.Context) RegistrationDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationDefinitionOutput)
}

func (i *RegistrationDefinition) ToOutput(ctx context.Context) pulumix.Output[*RegistrationDefinition] {
	return pulumix.Output[*RegistrationDefinition]{
		OutputState: i.ToRegistrationDefinitionOutputWithContext(ctx).OutputState,
	}
}

type RegistrationDefinitionOutput struct{ *pulumi.OutputState }

func (RegistrationDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationDefinition)(nil)).Elem()
}

func (o RegistrationDefinitionOutput) ToRegistrationDefinitionOutput() RegistrationDefinitionOutput {
	return o
}

func (o RegistrationDefinitionOutput) ToRegistrationDefinitionOutputWithContext(ctx context.Context) RegistrationDefinitionOutput {
	return o
}

func (o RegistrationDefinitionOutput) ToOutput(ctx context.Context) pulumix.Output[*RegistrationDefinition] {
	return pulumix.Output[*RegistrationDefinition]{
		OutputState: o.OutputState,
	}
}

// The name of the registration definition.
func (o RegistrationDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistrationDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The details for the Managed Services offer’s plan in Azure Marketplace.
func (o RegistrationDefinitionOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v *RegistrationDefinition) PlanResponsePtrOutput { return v.Plan }).(PlanResponsePtrOutput)
}

// The properties of a registration definition.
func (o RegistrationDefinitionOutput) Properties() RegistrationDefinitionPropertiesResponseOutput {
	return o.ApplyT(func(v *RegistrationDefinition) RegistrationDefinitionPropertiesResponseOutput { return v.Properties }).(RegistrationDefinitionPropertiesResponseOutput)
}

// The metadata for the registration assignment resource.
func (o RegistrationDefinitionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistrationDefinition) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the Azure resource (Microsoft.ManagedServices/registrationDefinitions).
func (o RegistrationDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistrationDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistrationDefinitionOutput{})
}
