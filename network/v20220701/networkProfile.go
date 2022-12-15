// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Network profile resource.
type NetworkProfile struct {
	pulumi.CustomResourceState

	// List of chid container network interface configurations.
	ContainerNetworkInterfaceConfigurations ContainerNetworkInterfaceConfigurationResponseArrayOutput `pulumi:"containerNetworkInterfaceConfigurations"`
	// List of child container network interfaces.
	ContainerNetworkInterfaces ContainerNetworkInterfaceResponseArrayOutput `pulumi:"containerNetworkInterfaces"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the network profile resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource GUID property of the network profile resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkProfile registers a new resource with the given unique name, arguments, and options.
func NewNetworkProfile(ctx *pulumi.Context,
	name string, args *NetworkProfileArgs, opts ...pulumi.ResourceOption) (*NetworkProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:NetworkProfile"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkProfile
	err := ctx.RegisterResource("azure-native:network/v20220701:NetworkProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkProfile gets an existing NetworkProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkProfileState, opts ...pulumi.ResourceOption) (*NetworkProfile, error) {
	var resource NetworkProfile
	err := ctx.ReadResource("azure-native:network/v20220701:NetworkProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkProfile resources.
type networkProfileState struct {
}

type NetworkProfileState struct {
}

func (NetworkProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkProfileState)(nil)).Elem()
}

type networkProfileArgs struct {
	// List of chid container network interface configurations.
	ContainerNetworkInterfaceConfigurations []ContainerNetworkInterfaceConfiguration `pulumi:"containerNetworkInterfaceConfigurations"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the network profile.
	NetworkProfileName *string `pulumi:"networkProfileName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkProfile resource.
type NetworkProfileArgs struct {
	// List of chid container network interface configurations.
	ContainerNetworkInterfaceConfigurations ContainerNetworkInterfaceConfigurationArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the network profile.
	NetworkProfileName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkProfileArgs)(nil)).Elem()
}

type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput
}

func (*NetworkProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *NetworkProfile) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i *NetworkProfile) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

// List of chid container network interface configurations.
func (o NetworkProfileOutput) ContainerNetworkInterfaceConfigurations() ContainerNetworkInterfaceConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) ContainerNetworkInterfaceConfigurationResponseArrayOutput {
		return v.ContainerNetworkInterfaceConfigurations
	}).(ContainerNetworkInterfaceConfigurationResponseArrayOutput)
}

// List of child container network interfaces.
func (o NetworkProfileOutput) ContainerNetworkInterfaces() ContainerNetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) ContainerNetworkInterfaceResponseArrayOutput {
		return v.ContainerNetworkInterfaces
	}).(ContainerNetworkInterfaceResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o NetworkProfileOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource location.
func (o NetworkProfileOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o NetworkProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the network profile resource.
func (o NetworkProfileOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the network profile resource.
func (o NetworkProfileOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Resource tags.
func (o NetworkProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o NetworkProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkProfileOutput{})
}