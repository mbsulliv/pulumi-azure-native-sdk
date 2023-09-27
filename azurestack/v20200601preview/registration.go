// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Registration information.
type Registration struct {
	pulumi.CustomResourceState

	// Specifies the billing mode for the Azure Stack registration.
	BillingModel pulumi.StringPtrOutput `pulumi:"billingModel"`
	// The identifier of the registered Azure Stack.
	CloudId pulumi.StringPtrOutput `pulumi:"cloudId"`
	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the resource.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The object identifier associated with the Azure Stack connecting to Azure.
	ObjectId pulumi.StringPtrOutput `pulumi:"objectId"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Custom tags for the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of Resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistration registers a new resource with the given unique name, arguments, and options.
func NewRegistration(ctx *pulumi.Context,
	name string, args *RegistrationArgs, opts ...pulumi.ResourceOption) (*Registration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistrationToken == nil {
		return nil, errors.New("invalid value for required argument 'RegistrationToken'")
	}
	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestack:Registration"),
		},
		{
			Type: pulumi.String("azure-native:azurestack/v20160101:Registration"),
		},
		{
			Type: pulumi.String("azure-native:azurestack/v20170601:Registration"),
		},
		{
			Type: pulumi.String("azure-native:azurestack/v20220601:Registration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Registration
	err := ctx.RegisterResource("azure-native:azurestack/v20200601preview:Registration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistration gets an existing Registration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistrationState, opts ...pulumi.ResourceOption) (*Registration, error) {
	var resource Registration
	err := ctx.ReadResource("azure-native:azurestack/v20200601preview:Registration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Registration resources.
type registrationState struct {
}

type RegistrationState struct {
}

func (RegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationState)(nil)).Elem()
}

type registrationArgs struct {
	// Location of the resource.
	Location *string `pulumi:"location"`
	// Name of the Azure Stack registration.
	RegistrationName *string `pulumi:"registrationName"`
	// The token identifying registered Azure Stack
	RegistrationToken string `pulumi:"registrationToken"`
	// Name of the resource group.
	ResourceGroup string `pulumi:"resourceGroup"`
}

// The set of arguments for constructing a Registration resource.
type RegistrationArgs struct {
	// Location of the resource.
	Location pulumi.StringPtrInput
	// Name of the Azure Stack registration.
	RegistrationName pulumi.StringPtrInput
	// The token identifying registered Azure Stack
	RegistrationToken pulumi.StringInput
	// Name of the resource group.
	ResourceGroup pulumi.StringInput
}

func (RegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationArgs)(nil)).Elem()
}

type RegistrationInput interface {
	pulumi.Input

	ToRegistrationOutput() RegistrationOutput
	ToRegistrationOutputWithContext(ctx context.Context) RegistrationOutput
}

func (*Registration) ElementType() reflect.Type {
	return reflect.TypeOf((**Registration)(nil)).Elem()
}

func (i *Registration) ToRegistrationOutput() RegistrationOutput {
	return i.ToRegistrationOutputWithContext(context.Background())
}

func (i *Registration) ToRegistrationOutputWithContext(ctx context.Context) RegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationOutput)
}

func (i *Registration) ToOutput(ctx context.Context) pulumix.Output[*Registration] {
	return pulumix.Output[*Registration]{
		OutputState: i.ToRegistrationOutputWithContext(ctx).OutputState,
	}
}

type RegistrationOutput struct{ *pulumi.OutputState }

func (RegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Registration)(nil)).Elem()
}

func (o RegistrationOutput) ToRegistrationOutput() RegistrationOutput {
	return o
}

func (o RegistrationOutput) ToRegistrationOutputWithContext(ctx context.Context) RegistrationOutput {
	return o
}

func (o RegistrationOutput) ToOutput(ctx context.Context) pulumix.Output[*Registration] {
	return pulumix.Output[*Registration]{
		OutputState: o.OutputState,
	}
}

// Specifies the billing mode for the Azure Stack registration.
func (o RegistrationOutput) BillingModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringPtrOutput { return v.BillingModel }).(pulumi.StringPtrOutput)
}

// The identifier of the registered Azure Stack.
func (o RegistrationOutput) CloudId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringPtrOutput { return v.CloudId }).(pulumi.StringPtrOutput)
}

// The entity tag used for optimistic concurrency when modifying the resource.
func (o RegistrationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the resource.
func (o RegistrationOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Location of the resource.
func (o RegistrationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the resource.
func (o RegistrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The object identifier associated with the Azure Stack connecting to Azure.
func (o RegistrationOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringPtrOutput { return v.ObjectId }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o RegistrationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Registration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Custom tags for the resource.
func (o RegistrationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of Resource.
func (o RegistrationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistrationOutput{})
}
