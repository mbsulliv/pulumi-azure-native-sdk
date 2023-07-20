// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230615

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The IP Extended Community resource definition.
type IpExtendedCommunity struct {
	pulumi.CustomResourceState

	// Administrative state of the resource.
	AdministrativeState pulumi.StringOutput `pulumi:"administrativeState"`
	// Switch configuration description.
	Annotation pulumi.StringPtrOutput `pulumi:"annotation"`
	// Configuration state of the resource.
	ConfigurationState pulumi.StringOutput `pulumi:"configurationState"`
	// List of IP Extended Community Rules.
	IpExtendedCommunityRules IpExtendedCommunityRuleResponseArrayOutput `pulumi:"ipExtendedCommunityRules"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIpExtendedCommunity registers a new resource with the given unique name, arguments, and options.
func NewIpExtendedCommunity(ctx *pulumi.Context,
	name string, args *IpExtendedCommunityArgs, opts ...pulumi.ResourceOption) (*IpExtendedCommunity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpExtendedCommunityRules == nil {
		return nil, errors.New("invalid value for required argument 'IpExtendedCommunityRules'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetworkfabric:IpExtendedCommunity"),
		},
		{
			Type: pulumi.String("azure-native:managednetworkfabric/v20230201preview:IpExtendedCommunity"),
		},
	})
	opts = append(opts, aliases)
	var resource IpExtendedCommunity
	err := ctx.RegisterResource("azure-native:managednetworkfabric/v20230615:IpExtendedCommunity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpExtendedCommunity gets an existing IpExtendedCommunity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpExtendedCommunity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpExtendedCommunityState, opts ...pulumi.ResourceOption) (*IpExtendedCommunity, error) {
	var resource IpExtendedCommunity
	err := ctx.ReadResource("azure-native:managednetworkfabric/v20230615:IpExtendedCommunity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpExtendedCommunity resources.
type ipExtendedCommunityState struct {
}

type IpExtendedCommunityState struct {
}

func (IpExtendedCommunityState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipExtendedCommunityState)(nil)).Elem()
}

type ipExtendedCommunityArgs struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Name of the IP Extended Community.
	IpExtendedCommunityName *string `pulumi:"ipExtendedCommunityName"`
	// List of IP Extended Community Rules.
	IpExtendedCommunityRules []IpExtendedCommunityRule `pulumi:"ipExtendedCommunityRules"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IpExtendedCommunity resource.
type IpExtendedCommunityArgs struct {
	// Switch configuration description.
	Annotation pulumi.StringPtrInput
	// Name of the IP Extended Community.
	IpExtendedCommunityName pulumi.StringPtrInput
	// List of IP Extended Community Rules.
	IpExtendedCommunityRules IpExtendedCommunityRuleArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (IpExtendedCommunityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipExtendedCommunityArgs)(nil)).Elem()
}

type IpExtendedCommunityInput interface {
	pulumi.Input

	ToIpExtendedCommunityOutput() IpExtendedCommunityOutput
	ToIpExtendedCommunityOutputWithContext(ctx context.Context) IpExtendedCommunityOutput
}

func (*IpExtendedCommunity) ElementType() reflect.Type {
	return reflect.TypeOf((**IpExtendedCommunity)(nil)).Elem()
}

func (i *IpExtendedCommunity) ToIpExtendedCommunityOutput() IpExtendedCommunityOutput {
	return i.ToIpExtendedCommunityOutputWithContext(context.Background())
}

func (i *IpExtendedCommunity) ToIpExtendedCommunityOutputWithContext(ctx context.Context) IpExtendedCommunityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpExtendedCommunityOutput)
}

type IpExtendedCommunityOutput struct{ *pulumi.OutputState }

func (IpExtendedCommunityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpExtendedCommunity)(nil)).Elem()
}

func (o IpExtendedCommunityOutput) ToIpExtendedCommunityOutput() IpExtendedCommunityOutput {
	return o
}

func (o IpExtendedCommunityOutput) ToIpExtendedCommunityOutputWithContext(ctx context.Context) IpExtendedCommunityOutput {
	return o
}

// Administrative state of the resource.
func (o IpExtendedCommunityOutput) AdministrativeState() pulumi.StringOutput {
	return o.ApplyT(func(v *IpExtendedCommunity) pulumi.StringOutput { return v.AdministrativeState }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o IpExtendedCommunityOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpExtendedCommunity) pulumi.StringPtrOutput { return v.Annotation }).(pulumi.StringPtrOutput)
}

// Configuration state of the resource.
func (o IpExtendedCommunityOutput) ConfigurationState() pulumi.StringOutput {
	return o.ApplyT(func(v *IpExtendedCommunity) pulumi.StringOutput { return v.ConfigurationState }).(pulumi.StringOutput)
}

// List of IP Extended Community Rules.
func (o IpExtendedCommunityOutput) IpExtendedCommunityRules() IpExtendedCommunityRuleResponseArrayOutput {
	return o.ApplyT(func(v *IpExtendedCommunity) IpExtendedCommunityRuleResponseArrayOutput {
		return v.IpExtendedCommunityRules
	}).(IpExtendedCommunityRuleResponseArrayOutput)
}

// The geo-location where the resource lives
func (o IpExtendedCommunityOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *IpExtendedCommunity) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o IpExtendedCommunityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpExtendedCommunity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o IpExtendedCommunityOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *IpExtendedCommunity) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o IpExtendedCommunityOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IpExtendedCommunity) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o IpExtendedCommunityOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IpExtendedCommunity) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o IpExtendedCommunityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IpExtendedCommunity) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IpExtendedCommunityOutput{})
}
