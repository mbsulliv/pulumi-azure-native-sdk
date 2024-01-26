// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Arc Addon.
type ArcAddon struct {
	pulumi.CustomResourceState

	// Host OS supported by the Arc addon.
	HostPlatform pulumi.StringOutput `pulumi:"hostPlatform"`
	// Platform where the runtime is hosted.
	HostPlatformType pulumi.StringOutput `pulumi:"hostPlatformType"`
	// Addon type.
	// Expected value is 'ArcForKubernetes'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Addon Provisioning State
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Arc resource group name
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Arc resource location
	ResourceLocation pulumi.StringOutput `pulumi:"resourceLocation"`
	// Arc resource Name
	ResourceName pulumi.StringOutput `pulumi:"resourceName"`
	// Arc resource subscription Id
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// Metadata pertaining to creation and last modification of Addon
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
	// Arc resource version
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewArcAddon registers a new resource with the given unique name, arguments, and options.
func NewArcAddon(ctx *pulumi.Context,
	name string, args *ArcAddonArgs, opts ...pulumi.ResourceOption) (*ArcAddon, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceLocation == nil {
		return nil, errors.New("invalid value for required argument 'ResourceLocation'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	args.Kind = pulumi.String("ArcForKubernetes")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230101preview:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230701:ArcAddon"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ArcAddon
	err := ctx.RegisterResource("azure-native:databoxedge/v20231201:ArcAddon", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArcAddon gets an existing ArcAddon resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArcAddon(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArcAddonState, opts ...pulumi.ResourceOption) (*ArcAddon, error) {
	var resource ArcAddon
	err := ctx.ReadResource("azure-native:databoxedge/v20231201:ArcAddon", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ArcAddon resources.
type arcAddonState struct {
}

type ArcAddonState struct {
}

func (ArcAddonState) ElementType() reflect.Type {
	return reflect.TypeOf((*arcAddonState)(nil)).Elem()
}

type arcAddonArgs struct {
	// The addon name.
	AddonName *string `pulumi:"addonName"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// Addon type.
	// Expected value is 'ArcForKubernetes'.
	Kind string `pulumi:"kind"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Arc resource location
	ResourceLocation string `pulumi:"resourceLocation"`
	// Arc resource Name
	ResourceName string `pulumi:"resourceName"`
	// The role name.
	RoleName string `pulumi:"roleName"`
	// Arc resource subscription Id
	SubscriptionId string `pulumi:"subscriptionId"`
}

// The set of arguments for constructing a ArcAddon resource.
type ArcAddonArgs struct {
	// The addon name.
	AddonName pulumi.StringPtrInput
	// The device name.
	DeviceName pulumi.StringInput
	// Addon type.
	// Expected value is 'ArcForKubernetes'.
	Kind pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Arc resource location
	ResourceLocation pulumi.StringInput
	// Arc resource Name
	ResourceName pulumi.StringInput
	// The role name.
	RoleName pulumi.StringInput
	// Arc resource subscription Id
	SubscriptionId pulumi.StringInput
}

func (ArcAddonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*arcAddonArgs)(nil)).Elem()
}

type ArcAddonInput interface {
	pulumi.Input

	ToArcAddonOutput() ArcAddonOutput
	ToArcAddonOutputWithContext(ctx context.Context) ArcAddonOutput
}

func (*ArcAddon) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcAddon)(nil)).Elem()
}

func (i *ArcAddon) ToArcAddonOutput() ArcAddonOutput {
	return i.ToArcAddonOutputWithContext(context.Background())
}

func (i *ArcAddon) ToArcAddonOutputWithContext(ctx context.Context) ArcAddonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcAddonOutput)
}

type ArcAddonOutput struct{ *pulumi.OutputState }

func (ArcAddonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcAddon)(nil)).Elem()
}

func (o ArcAddonOutput) ToArcAddonOutput() ArcAddonOutput {
	return o
}

func (o ArcAddonOutput) ToArcAddonOutputWithContext(ctx context.Context) ArcAddonOutput {
	return o
}

// Host OS supported by the Arc addon.
func (o ArcAddonOutput) HostPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.HostPlatform }).(pulumi.StringOutput)
}

// Platform where the runtime is hosted.
func (o ArcAddonOutput) HostPlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.HostPlatformType }).(pulumi.StringOutput)
}

// Addon type.
// Expected value is 'ArcForKubernetes'.
func (o ArcAddonOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The object name.
func (o ArcAddonOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Addon Provisioning State
func (o ArcAddonOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Arc resource group name
func (o ArcAddonOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// Arc resource location
func (o ArcAddonOutput) ResourceLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.ResourceLocation }).(pulumi.StringOutput)
}

// Arc resource Name
func (o ArcAddonOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.ResourceName }).(pulumi.StringOutput)
}

// Arc resource subscription Id
func (o ArcAddonOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of Addon
func (o ArcAddonOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ArcAddon) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o ArcAddonOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Arc resource version
func (o ArcAddonOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ArcAddonOutput{})
}
