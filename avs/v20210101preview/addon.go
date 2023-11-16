// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An addon resource
type Addon struct {
	pulumi.CustomResourceState

	// The type of private cloud addon
	AddonType pulumi.StringPtrOutput `pulumi:"addonType"`
	// The SRM license
	LicenseKey pulumi.StringPtrOutput `pulumi:"licenseKey"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The state of the addon provisioning
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAddon registers a new resource with the given unique name, arguments, and options.
func NewAddon(ctx *pulumi.Context,
	name string, args *AddonArgs, opts ...pulumi.ResourceOption) (*Addon, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:Addon"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:Addon"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:Addon"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:Addon"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:Addon"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230301:Addon"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Addon
	err := ctx.RegisterResource("azure-native:avs/v20210101preview:Addon", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddon gets an existing Addon resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddon(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddonState, opts ...pulumi.ResourceOption) (*Addon, error) {
	var resource Addon
	err := ctx.ReadResource("azure-native:avs/v20210101preview:Addon", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Addon resources.
type addonState struct {
}

type AddonState struct {
}

func (AddonState) ElementType() reflect.Type {
	return reflect.TypeOf((*addonState)(nil)).Elem()
}

type addonArgs struct {
	// Name of the addon for the private cloud
	AddonName *string `pulumi:"addonName"`
	// The type of private cloud addon
	AddonType *string `pulumi:"addonType"`
	// The SRM license
	LicenseKey *string `pulumi:"licenseKey"`
	// The name of the private cloud.
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Addon resource.
type AddonArgs struct {
	// Name of the addon for the private cloud
	AddonName pulumi.StringPtrInput
	// The type of private cloud addon
	AddonType pulumi.StringPtrInput
	// The SRM license
	LicenseKey pulumi.StringPtrInput
	// The name of the private cloud.
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (AddonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addonArgs)(nil)).Elem()
}

type AddonInput interface {
	pulumi.Input

	ToAddonOutput() AddonOutput
	ToAddonOutputWithContext(ctx context.Context) AddonOutput
}

func (*Addon) ElementType() reflect.Type {
	return reflect.TypeOf((**Addon)(nil)).Elem()
}

func (i *Addon) ToAddonOutput() AddonOutput {
	return i.ToAddonOutputWithContext(context.Background())
}

func (i *Addon) ToAddonOutputWithContext(ctx context.Context) AddonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonOutput)
}

type AddonOutput struct{ *pulumi.OutputState }

func (AddonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Addon)(nil)).Elem()
}

func (o AddonOutput) ToAddonOutput() AddonOutput {
	return o
}

func (o AddonOutput) ToAddonOutputWithContext(ctx context.Context) AddonOutput {
	return o
}

// The type of private cloud addon
func (o AddonOutput) AddonType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringPtrOutput { return v.AddonType }).(pulumi.StringPtrOutput)
}

// The SRM license
func (o AddonOutput) LicenseKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringPtrOutput { return v.LicenseKey }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o AddonOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The state of the addon provisioning
func (o AddonOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o AddonOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AddonOutput{})
}
