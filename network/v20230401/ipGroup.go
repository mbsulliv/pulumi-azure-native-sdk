// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The IpGroups resource information.
type IpGroup struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// List of references to Firewall Policies resources that this IpGroups is associated with.
	FirewallPolicies SubResourceResponseArrayOutput `pulumi:"firewallPolicies"`
	// List of references to Firewall resources that this IpGroups is associated with.
	Firewalls SubResourceResponseArrayOutput `pulumi:"firewalls"`
	// IpAddresses/IpAddressPrefixes in the IpGroups resource.
	IpAddresses pulumi.StringArrayOutput `pulumi:"ipAddresses"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the IpGroups resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIpGroup registers a new resource with the given unique name, arguments, and options.
func NewIpGroup(ctx *pulumi.Context,
	name string, args *IpGroupArgs, opts ...pulumi.ResourceOption) (*IpGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:IpGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IpGroup
	err := ctx.RegisterResource("azure-native:network/v20230401:IpGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpGroup gets an existing IpGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpGroupState, opts ...pulumi.ResourceOption) (*IpGroup, error) {
	var resource IpGroup
	err := ctx.ReadResource("azure-native:network/v20230401:IpGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpGroup resources.
type ipGroupState struct {
}

type IpGroupState struct {
}

func (IpGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipGroupState)(nil)).Elem()
}

type ipGroupArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// IpAddresses/IpAddressPrefixes in the IpGroups resource.
	IpAddresses []string `pulumi:"ipAddresses"`
	// The name of the ipGroups.
	IpGroupsName *string `pulumi:"ipGroupsName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IpGroup resource.
type IpGroupArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// IpAddresses/IpAddressPrefixes in the IpGroups resource.
	IpAddresses pulumi.StringArrayInput
	// The name of the ipGroups.
	IpGroupsName pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (IpGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipGroupArgs)(nil)).Elem()
}

type IpGroupInput interface {
	pulumi.Input

	ToIpGroupOutput() IpGroupOutput
	ToIpGroupOutputWithContext(ctx context.Context) IpGroupOutput
}

func (*IpGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**IpGroup)(nil)).Elem()
}

func (i *IpGroup) ToIpGroupOutput() IpGroupOutput {
	return i.ToIpGroupOutputWithContext(context.Background())
}

func (i *IpGroup) ToIpGroupOutputWithContext(ctx context.Context) IpGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpGroupOutput)
}

func (i *IpGroup) ToOutput(ctx context.Context) pulumix.Output[*IpGroup] {
	return pulumix.Output[*IpGroup]{
		OutputState: i.ToIpGroupOutputWithContext(ctx).OutputState,
	}
}

type IpGroupOutput struct{ *pulumi.OutputState }

func (IpGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpGroup)(nil)).Elem()
}

func (o IpGroupOutput) ToIpGroupOutput() IpGroupOutput {
	return o
}

func (o IpGroupOutput) ToIpGroupOutputWithContext(ctx context.Context) IpGroupOutput {
	return o
}

func (o IpGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*IpGroup] {
	return pulumix.Output[*IpGroup]{
		OutputState: o.OutputState,
	}
}

// A unique read-only string that changes whenever the resource is updated.
func (o IpGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IpGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// List of references to Firewall Policies resources that this IpGroups is associated with.
func (o IpGroupOutput) FirewallPolicies() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *IpGroup) SubResourceResponseArrayOutput { return v.FirewallPolicies }).(SubResourceResponseArrayOutput)
}

// List of references to Firewall resources that this IpGroups is associated with.
func (o IpGroupOutput) Firewalls() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *IpGroup) SubResourceResponseArrayOutput { return v.Firewalls }).(SubResourceResponseArrayOutput)
}

// IpAddresses/IpAddressPrefixes in the IpGroups resource.
func (o IpGroupOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IpGroup) pulumi.StringArrayOutput { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// Resource location.
func (o IpGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o IpGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the IpGroups resource.
func (o IpGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *IpGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o IpGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IpGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o IpGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IpGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IpGroupOutput{})
}
