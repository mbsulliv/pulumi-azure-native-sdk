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

// VpnServerConfigurationPolicyGroup Resource.
type ConfigurationPolicyGroup struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Shows if this is a Default VpnServerConfigurationPolicyGroup or not.
	IsDefault pulumi.BoolPtrOutput `pulumi:"isDefault"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// List of references to P2SConnectionConfigurations.
	P2SConnectionConfigurations SubResourceResponseArrayOutput `pulumi:"p2SConnectionConfigurations"`
	// Multiple PolicyMembers for VpnServerConfigurationPolicyGroup.
	PolicyMembers VpnServerConfigurationPolicyGroupMemberResponseArrayOutput `pulumi:"policyMembers"`
	// Priority for VpnServerConfigurationPolicyGroup.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The provisioning state of the VpnServerConfigurationPolicyGroup resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationPolicyGroup registers a new resource with the given unique name, arguments, and options.
func NewConfigurationPolicyGroup(ctx *pulumi.Context,
	name string, args *ConfigurationPolicyGroupArgs, opts ...pulumi.ResourceOption) (*ConfigurationPolicyGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VpnServerConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'VpnServerConfigurationName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ConfigurationPolicyGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ConfigurationPolicyGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ConfigurationPolicyGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ConfigurationPolicyGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ConfigurationPolicyGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:ConfigurationPolicyGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:ConfigurationPolicyGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:ConfigurationPolicyGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:ConfigurationPolicyGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigurationPolicyGroup
	err := ctx.RegisterResource("azure-native:network/v20230401:ConfigurationPolicyGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationPolicyGroup gets an existing ConfigurationPolicyGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationPolicyGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationPolicyGroupState, opts ...pulumi.ResourceOption) (*ConfigurationPolicyGroup, error) {
	var resource ConfigurationPolicyGroup
	err := ctx.ReadResource("azure-native:network/v20230401:ConfigurationPolicyGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationPolicyGroup resources.
type configurationPolicyGroupState struct {
}

type ConfigurationPolicyGroupState struct {
}

func (ConfigurationPolicyGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationPolicyGroupState)(nil)).Elem()
}

type configurationPolicyGroupArgs struct {
	// The name of the ConfigurationPolicyGroup.
	ConfigurationPolicyGroupName *string `pulumi:"configurationPolicyGroupName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Shows if this is a Default VpnServerConfigurationPolicyGroup or not.
	IsDefault *bool `pulumi:"isDefault"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Multiple PolicyMembers for VpnServerConfigurationPolicyGroup.
	PolicyMembers []VpnServerConfigurationPolicyGroupMember `pulumi:"policyMembers"`
	// Priority for VpnServerConfigurationPolicyGroup.
	Priority *int `pulumi:"priority"`
	// The resource group name of the ConfigurationPolicyGroup.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VpnServerConfiguration.
	VpnServerConfigurationName string `pulumi:"vpnServerConfigurationName"`
}

// The set of arguments for constructing a ConfigurationPolicyGroup resource.
type ConfigurationPolicyGroupArgs struct {
	// The name of the ConfigurationPolicyGroup.
	ConfigurationPolicyGroupName pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Shows if this is a Default VpnServerConfigurationPolicyGroup or not.
	IsDefault pulumi.BoolPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Multiple PolicyMembers for VpnServerConfigurationPolicyGroup.
	PolicyMembers VpnServerConfigurationPolicyGroupMemberArrayInput
	// Priority for VpnServerConfigurationPolicyGroup.
	Priority pulumi.IntPtrInput
	// The resource group name of the ConfigurationPolicyGroup.
	ResourceGroupName pulumi.StringInput
	// The name of the VpnServerConfiguration.
	VpnServerConfigurationName pulumi.StringInput
}

func (ConfigurationPolicyGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationPolicyGroupArgs)(nil)).Elem()
}

type ConfigurationPolicyGroupInput interface {
	pulumi.Input

	ToConfigurationPolicyGroupOutput() ConfigurationPolicyGroupOutput
	ToConfigurationPolicyGroupOutputWithContext(ctx context.Context) ConfigurationPolicyGroupOutput
}

func (*ConfigurationPolicyGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationPolicyGroup)(nil)).Elem()
}

func (i *ConfigurationPolicyGroup) ToConfigurationPolicyGroupOutput() ConfigurationPolicyGroupOutput {
	return i.ToConfigurationPolicyGroupOutputWithContext(context.Background())
}

func (i *ConfigurationPolicyGroup) ToConfigurationPolicyGroupOutputWithContext(ctx context.Context) ConfigurationPolicyGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationPolicyGroupOutput)
}

func (i *ConfigurationPolicyGroup) ToOutput(ctx context.Context) pulumix.Output[*ConfigurationPolicyGroup] {
	return pulumix.Output[*ConfigurationPolicyGroup]{
		OutputState: i.ToConfigurationPolicyGroupOutputWithContext(ctx).OutputState,
	}
}

type ConfigurationPolicyGroupOutput struct{ *pulumi.OutputState }

func (ConfigurationPolicyGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationPolicyGroup)(nil)).Elem()
}

func (o ConfigurationPolicyGroupOutput) ToConfigurationPolicyGroupOutput() ConfigurationPolicyGroupOutput {
	return o
}

func (o ConfigurationPolicyGroupOutput) ToConfigurationPolicyGroupOutputWithContext(ctx context.Context) ConfigurationPolicyGroupOutput {
	return o
}

func (o ConfigurationPolicyGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*ConfigurationPolicyGroup] {
	return pulumix.Output[*ConfigurationPolicyGroup]{
		OutputState: o.OutputState,
	}
}

// A unique read-only string that changes whenever the resource is updated.
func (o ConfigurationPolicyGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Shows if this is a Default VpnServerConfigurationPolicyGroup or not.
func (o ConfigurationPolicyGroupOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) pulumi.BoolPtrOutput { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o ConfigurationPolicyGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// List of references to P2SConnectionConfigurations.
func (o ConfigurationPolicyGroupOutput) P2SConnectionConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) SubResourceResponseArrayOutput { return v.P2SConnectionConfigurations }).(SubResourceResponseArrayOutput)
}

// Multiple PolicyMembers for VpnServerConfigurationPolicyGroup.
func (o ConfigurationPolicyGroupOutput) PolicyMembers() VpnServerConfigurationPolicyGroupMemberResponseArrayOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) VpnServerConfigurationPolicyGroupMemberResponseArrayOutput {
		return v.PolicyMembers
	}).(VpnServerConfigurationPolicyGroupMemberResponseArrayOutput)
}

// Priority for VpnServerConfigurationPolicyGroup.
func (o ConfigurationPolicyGroupOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// The provisioning state of the VpnServerConfigurationPolicyGroup resource.
func (o ConfigurationPolicyGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o ConfigurationPolicyGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationPolicyGroupOutput{})
}
