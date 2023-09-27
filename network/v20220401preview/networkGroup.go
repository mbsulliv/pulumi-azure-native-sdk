// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The network group resource
type NetworkGroup struct {
	pulumi.CustomResourceState

	// A description of the network group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Group member type.
	MemberType pulumi.StringOutput `pulumi:"memberType"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the scope assignment resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkGroup registers a new resource with the given unique name, arguments, and options.
func NewNetworkGroup(ctx *pulumi.Context,
	name string, args *NetworkGroupArgs, opts ...pulumi.ResourceOption) (*NetworkGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MemberType == nil {
		return nil, errors.New("invalid value for required argument 'MemberType'")
	}
	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:NetworkGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NetworkGroup
	err := ctx.RegisterResource("azure-native:network/v20220401preview:NetworkGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkGroup gets an existing NetworkGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkGroupState, opts ...pulumi.ResourceOption) (*NetworkGroup, error) {
	var resource NetworkGroup
	err := ctx.ReadResource("azure-native:network/v20220401preview:NetworkGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkGroup resources.
type networkGroupState struct {
}

type NetworkGroupState struct {
}

func (NetworkGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkGroupState)(nil)).Elem()
}

type networkGroupArgs struct {
	// A description of the network group.
	Description *string `pulumi:"description"`
	// Group member type.
	MemberType string `pulumi:"memberType"`
	// The name of the network group.
	NetworkGroupName *string `pulumi:"networkGroupName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a NetworkGroup resource.
type NetworkGroupArgs struct {
	// A description of the network group.
	Description pulumi.StringPtrInput
	// Group member type.
	MemberType pulumi.StringInput
	// The name of the network group.
	NetworkGroupName pulumi.StringPtrInput
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (NetworkGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkGroupArgs)(nil)).Elem()
}

type NetworkGroupInput interface {
	pulumi.Input

	ToNetworkGroupOutput() NetworkGroupOutput
	ToNetworkGroupOutputWithContext(ctx context.Context) NetworkGroupOutput
}

func (*NetworkGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkGroup)(nil)).Elem()
}

func (i *NetworkGroup) ToNetworkGroupOutput() NetworkGroupOutput {
	return i.ToNetworkGroupOutputWithContext(context.Background())
}

func (i *NetworkGroup) ToNetworkGroupOutputWithContext(ctx context.Context) NetworkGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkGroupOutput)
}

func (i *NetworkGroup) ToOutput(ctx context.Context) pulumix.Output[*NetworkGroup] {
	return pulumix.Output[*NetworkGroup]{
		OutputState: i.ToNetworkGroupOutputWithContext(ctx).OutputState,
	}
}

type NetworkGroupOutput struct{ *pulumi.OutputState }

func (NetworkGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkGroup)(nil)).Elem()
}

func (o NetworkGroupOutput) ToNetworkGroupOutput() NetworkGroupOutput {
	return o
}

func (o NetworkGroupOutput) ToNetworkGroupOutputWithContext(ctx context.Context) NetworkGroupOutput {
	return o
}

func (o NetworkGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*NetworkGroup] {
	return pulumix.Output[*NetworkGroup]{
		OutputState: o.OutputState,
	}
}

// A description of the network group.
func (o NetworkGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o NetworkGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Group member type.
func (o NetworkGroupOutput) MemberType() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkGroup) pulumi.StringOutput { return v.MemberType }).(pulumi.StringOutput)
}

// Resource name.
func (o NetworkGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the scope assignment resource.
func (o NetworkGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o NetworkGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NetworkGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o NetworkGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkGroupOutput{})
}
