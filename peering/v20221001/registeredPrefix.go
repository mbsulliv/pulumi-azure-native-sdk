// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The customer's prefix that is registered by the peering service provider.
type RegisteredPrefix struct {
	pulumi.CustomResourceState

	// The error message associated with the validation state, if any.
	ErrorMessage pulumi.StringOutput `pulumi:"errorMessage"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The peering service prefix key that is to be shared with the customer.
	PeeringServicePrefixKey pulumi.StringOutput `pulumi:"peeringServicePrefixKey"`
	// The customer's prefix from which traffic originates.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
	// The prefix validation state.
	PrefixValidationState pulumi.StringOutput `pulumi:"prefixValidationState"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegisteredPrefix registers a new resource with the given unique name, arguments, and options.
func NewRegisteredPrefix(ctx *pulumi.Context,
	name string, args *RegisteredPrefixArgs, opts ...pulumi.ResourceOption) (*RegisteredPrefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeeringName == nil {
		return nil, errors.New("invalid value for required argument 'PeeringName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:peering:RegisteredPrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200101preview:RegisteredPrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200401:RegisteredPrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20201001:RegisteredPrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210101:RegisteredPrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210601:RegisteredPrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220101:RegisteredPrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220601:RegisteredPrefix"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RegisteredPrefix
	err := ctx.RegisterResource("azure-native:peering/v20221001:RegisteredPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegisteredPrefix gets an existing RegisteredPrefix resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegisteredPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegisteredPrefixState, opts ...pulumi.ResourceOption) (*RegisteredPrefix, error) {
	var resource RegisteredPrefix
	err := ctx.ReadResource("azure-native:peering/v20221001:RegisteredPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegisteredPrefix resources.
type registeredPrefixState struct {
}

type RegisteredPrefixState struct {
}

func (RegisteredPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredPrefixState)(nil)).Elem()
}

type registeredPrefixArgs struct {
	// The name of the peering.
	PeeringName string `pulumi:"peeringName"`
	// The customer's prefix from which traffic originates.
	Prefix *string `pulumi:"prefix"`
	// The name of the registered prefix.
	RegisteredPrefixName *string `pulumi:"registeredPrefixName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a RegisteredPrefix resource.
type RegisteredPrefixArgs struct {
	// The name of the peering.
	PeeringName pulumi.StringInput
	// The customer's prefix from which traffic originates.
	Prefix pulumi.StringPtrInput
	// The name of the registered prefix.
	RegisteredPrefixName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (RegisteredPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredPrefixArgs)(nil)).Elem()
}

type RegisteredPrefixInput interface {
	pulumi.Input

	ToRegisteredPrefixOutput() RegisteredPrefixOutput
	ToRegisteredPrefixOutputWithContext(ctx context.Context) RegisteredPrefixOutput
}

func (*RegisteredPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((**RegisteredPrefix)(nil)).Elem()
}

func (i *RegisteredPrefix) ToRegisteredPrefixOutput() RegisteredPrefixOutput {
	return i.ToRegisteredPrefixOutputWithContext(context.Background())
}

func (i *RegisteredPrefix) ToRegisteredPrefixOutputWithContext(ctx context.Context) RegisteredPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegisteredPrefixOutput)
}

func (i *RegisteredPrefix) ToOutput(ctx context.Context) pulumix.Output[*RegisteredPrefix] {
	return pulumix.Output[*RegisteredPrefix]{
		OutputState: i.ToRegisteredPrefixOutputWithContext(ctx).OutputState,
	}
}

type RegisteredPrefixOutput struct{ *pulumi.OutputState }

func (RegisteredPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegisteredPrefix)(nil)).Elem()
}

func (o RegisteredPrefixOutput) ToRegisteredPrefixOutput() RegisteredPrefixOutput {
	return o
}

func (o RegisteredPrefixOutput) ToRegisteredPrefixOutputWithContext(ctx context.Context) RegisteredPrefixOutput {
	return o
}

func (o RegisteredPrefixOutput) ToOutput(ctx context.Context) pulumix.Output[*RegisteredPrefix] {
	return pulumix.Output[*RegisteredPrefix]{
		OutputState: o.OutputState,
	}
}

// The error message associated with the validation state, if any.
func (o RegisteredPrefixOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredPrefix) pulumi.StringOutput { return v.ErrorMessage }).(pulumi.StringOutput)
}

// The name of the resource.
func (o RegisteredPrefixOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredPrefix) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The peering service prefix key that is to be shared with the customer.
func (o RegisteredPrefixOutput) PeeringServicePrefixKey() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredPrefix) pulumi.StringOutput { return v.PeeringServicePrefixKey }).(pulumi.StringOutput)
}

// The customer's prefix from which traffic originates.
func (o RegisteredPrefixOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredPrefix) pulumi.StringPtrOutput { return v.Prefix }).(pulumi.StringPtrOutput)
}

// The prefix validation state.
func (o RegisteredPrefixOutput) PrefixValidationState() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredPrefix) pulumi.StringOutput { return v.PrefixValidationState }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o RegisteredPrefixOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredPrefix) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the resource.
func (o RegisteredPrefixOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredPrefix) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegisteredPrefixOutput{})
}
