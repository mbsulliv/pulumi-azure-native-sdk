// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents the DNSSEC configuration.
type DnssecConfig struct {
	pulumi.CustomResourceState

	// The etag of the DNSSEC configuration.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The name of the DNSSEC configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning State of the DNSSEC configuration.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The list of signing keys.
	SigningKeys SigningKeyResponseArrayOutput `pulumi:"signingKeys"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the DNSSEC configuration.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDnssecConfig registers a new resource with the given unique name, arguments, and options.
func NewDnssecConfig(ctx *pulumi.Context,
	name string, args *DnssecConfigArgs, opts ...pulumi.ResourceOption) (*DnssecConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ZoneName == nil {
		return nil, errors.New("invalid value for required argument 'ZoneName'")
	}
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DnssecConfig
	err := ctx.RegisterResource("azure-native:network/v20230701preview:DnssecConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnssecConfig gets an existing DnssecConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnssecConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnssecConfigState, opts ...pulumi.ResourceOption) (*DnssecConfig, error) {
	var resource DnssecConfig
	err := ctx.ReadResource("azure-native:network/v20230701preview:DnssecConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnssecConfig resources.
type dnssecConfigState struct {
}

type DnssecConfigState struct {
}

func (DnssecConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnssecConfigState)(nil)).Elem()
}

type dnssecConfigArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the DNS zone (without a terminating dot).
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a DnssecConfig resource.
type DnssecConfigArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the DNS zone (without a terminating dot).
	ZoneName pulumi.StringInput
}

func (DnssecConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnssecConfigArgs)(nil)).Elem()
}

type DnssecConfigInput interface {
	pulumi.Input

	ToDnssecConfigOutput() DnssecConfigOutput
	ToDnssecConfigOutputWithContext(ctx context.Context) DnssecConfigOutput
}

func (*DnssecConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**DnssecConfig)(nil)).Elem()
}

func (i *DnssecConfig) ToDnssecConfigOutput() DnssecConfigOutput {
	return i.ToDnssecConfigOutputWithContext(context.Background())
}

func (i *DnssecConfig) ToDnssecConfigOutputWithContext(ctx context.Context) DnssecConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnssecConfigOutput)
}

type DnssecConfigOutput struct{ *pulumi.OutputState }

func (DnssecConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnssecConfig)(nil)).Elem()
}

func (o DnssecConfigOutput) ToDnssecConfigOutput() DnssecConfigOutput {
	return o
}

func (o DnssecConfigOutput) ToDnssecConfigOutputWithContext(ctx context.Context) DnssecConfigOutput {
	return o
}

// The etag of the DNSSEC configuration.
func (o DnssecConfigOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnssecConfig) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The name of the DNSSEC configuration.
func (o DnssecConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DnssecConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning State of the DNSSEC configuration.
func (o DnssecConfigOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DnssecConfig) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The list of signing keys.
func (o DnssecConfigOutput) SigningKeys() SigningKeyResponseArrayOutput {
	return o.ApplyT(func(v *DnssecConfig) SigningKeyResponseArrayOutput { return v.SigningKeys }).(SigningKeyResponseArrayOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o DnssecConfigOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DnssecConfig) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the DNSSEC configuration.
func (o DnssecConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DnssecConfig) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DnssecConfigOutput{})
}
