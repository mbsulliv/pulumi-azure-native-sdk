// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230615

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The NetworkTapRule resource definition.
type NetworkTapRule struct {
	pulumi.CustomResourceState

	// Administrative state of the resource.
	AdministrativeState pulumi.StringOutput `pulumi:"administrativeState"`
	// Switch configuration description.
	Annotation pulumi.StringPtrOutput `pulumi:"annotation"`
	// Configuration state of the resource.
	ConfigurationState pulumi.StringOutput `pulumi:"configurationState"`
	// Input method to configure Network Tap Rule.
	ConfigurationType pulumi.StringOutput `pulumi:"configurationType"`
	// List of dynamic match configurations.
	DynamicMatchConfigurations CommonDynamicMatchConfigurationResponseArrayOutput `pulumi:"dynamicMatchConfigurations"`
	// The last sync timestamp.
	LastSyncedTime pulumi.StringOutput `pulumi:"lastSyncedTime"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// List of match configurations.
	MatchConfigurations NetworkTapRuleMatchConfigurationResponseArrayOutput `pulumi:"matchConfigurations"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARM resource Id of the NetworkTap.
	NetworkTapId pulumi.StringOutput `pulumi:"networkTapId"`
	// Polling interval in seconds.
	PollingIntervalInSeconds pulumi.IntPtrOutput `pulumi:"pollingIntervalInSeconds"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Network Tap Rules file URL.
	TapRulesUrl pulumi.StringPtrOutput `pulumi:"tapRulesUrl"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkTapRule registers a new resource with the given unique name, arguments, and options.
func NewNetworkTapRule(ctx *pulumi.Context,
	name string, args *NetworkTapRuleArgs, opts ...pulumi.ResourceOption) (*NetworkTapRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationType == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.PollingIntervalInSeconds == nil {
		args.PollingIntervalInSeconds = pulumi.IntPtr(30)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetworkfabric:NetworkTapRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NetworkTapRule
	err := ctx.RegisterResource("azure-native:managednetworkfabric/v20230615:NetworkTapRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkTapRule gets an existing NetworkTapRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkTapRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkTapRuleState, opts ...pulumi.ResourceOption) (*NetworkTapRule, error) {
	var resource NetworkTapRule
	err := ctx.ReadResource("azure-native:managednetworkfabric/v20230615:NetworkTapRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkTapRule resources.
type networkTapRuleState struct {
}

type NetworkTapRuleState struct {
}

func (NetworkTapRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkTapRuleState)(nil)).Elem()
}

type networkTapRuleArgs struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Input method to configure Network Tap Rule.
	ConfigurationType string `pulumi:"configurationType"`
	// List of dynamic match configurations.
	DynamicMatchConfigurations []CommonDynamicMatchConfiguration `pulumi:"dynamicMatchConfigurations"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// List of match configurations.
	MatchConfigurations []NetworkTapRuleMatchConfiguration `pulumi:"matchConfigurations"`
	// Name of the Network Tap Rule.
	NetworkTapRuleName *string `pulumi:"networkTapRuleName"`
	// Polling interval in seconds.
	PollingIntervalInSeconds *int `pulumi:"pollingIntervalInSeconds"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Network Tap Rules file URL.
	TapRulesUrl *string `pulumi:"tapRulesUrl"`
}

// The set of arguments for constructing a NetworkTapRule resource.
type NetworkTapRuleArgs struct {
	// Switch configuration description.
	Annotation pulumi.StringPtrInput
	// Input method to configure Network Tap Rule.
	ConfigurationType pulumi.StringInput
	// List of dynamic match configurations.
	DynamicMatchConfigurations CommonDynamicMatchConfigurationArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// List of match configurations.
	MatchConfigurations NetworkTapRuleMatchConfigurationArrayInput
	// Name of the Network Tap Rule.
	NetworkTapRuleName pulumi.StringPtrInput
	// Polling interval in seconds.
	PollingIntervalInSeconds pulumi.IntPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Network Tap Rules file URL.
	TapRulesUrl pulumi.StringPtrInput
}

func (NetworkTapRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkTapRuleArgs)(nil)).Elem()
}

type NetworkTapRuleInput interface {
	pulumi.Input

	ToNetworkTapRuleOutput() NetworkTapRuleOutput
	ToNetworkTapRuleOutputWithContext(ctx context.Context) NetworkTapRuleOutput
}

func (*NetworkTapRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkTapRule)(nil)).Elem()
}

func (i *NetworkTapRule) ToNetworkTapRuleOutput() NetworkTapRuleOutput {
	return i.ToNetworkTapRuleOutputWithContext(context.Background())
}

func (i *NetworkTapRule) ToNetworkTapRuleOutputWithContext(ctx context.Context) NetworkTapRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkTapRuleOutput)
}

func (i *NetworkTapRule) ToOutput(ctx context.Context) pulumix.Output[*NetworkTapRule] {
	return pulumix.Output[*NetworkTapRule]{
		OutputState: i.ToNetworkTapRuleOutputWithContext(ctx).OutputState,
	}
}

type NetworkTapRuleOutput struct{ *pulumi.OutputState }

func (NetworkTapRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkTapRule)(nil)).Elem()
}

func (o NetworkTapRuleOutput) ToNetworkTapRuleOutput() NetworkTapRuleOutput {
	return o
}

func (o NetworkTapRuleOutput) ToNetworkTapRuleOutputWithContext(ctx context.Context) NetworkTapRuleOutput {
	return o
}

func (o NetworkTapRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*NetworkTapRule] {
	return pulumix.Output[*NetworkTapRule]{
		OutputState: o.OutputState,
	}
}

// Administrative state of the resource.
func (o NetworkTapRuleOutput) AdministrativeState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkTapRule) pulumi.StringOutput { return v.AdministrativeState }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o NetworkTapRuleOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkTapRule) pulumi.StringPtrOutput { return v.Annotation }).(pulumi.StringPtrOutput)
}

// Configuration state of the resource.
func (o NetworkTapRuleOutput) ConfigurationState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkTapRule) pulumi.StringOutput { return v.ConfigurationState }).(pulumi.StringOutput)
}

// Input method to configure Network Tap Rule.
func (o NetworkTapRuleOutput) ConfigurationType() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkTapRule) pulumi.StringOutput { return v.ConfigurationType }).(pulumi.StringOutput)
}

// List of dynamic match configurations.
func (o NetworkTapRuleOutput) DynamicMatchConfigurations() CommonDynamicMatchConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkTapRule) CommonDynamicMatchConfigurationResponseArrayOutput {
		return v.DynamicMatchConfigurations
	}).(CommonDynamicMatchConfigurationResponseArrayOutput)
}

// The last sync timestamp.
func (o NetworkTapRuleOutput) LastSyncedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkTapRule) pulumi.StringOutput { return v.LastSyncedTime }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o NetworkTapRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkTapRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// List of match configurations.
func (o NetworkTapRuleOutput) MatchConfigurations() NetworkTapRuleMatchConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkTapRule) NetworkTapRuleMatchConfigurationResponseArrayOutput {
		return v.MatchConfigurations
	}).(NetworkTapRuleMatchConfigurationResponseArrayOutput)
}

// The name of the resource
func (o NetworkTapRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkTapRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ARM resource Id of the NetworkTap.
func (o NetworkTapRuleOutput) NetworkTapId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkTapRule) pulumi.StringOutput { return v.NetworkTapId }).(pulumi.StringOutput)
}

// Polling interval in seconds.
func (o NetworkTapRuleOutput) PollingIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetworkTapRule) pulumi.IntPtrOutput { return v.PollingIntervalInSeconds }).(pulumi.IntPtrOutput)
}

// Provisioning state of the resource.
func (o NetworkTapRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkTapRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o NetworkTapRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NetworkTapRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o NetworkTapRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkTapRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Network Tap Rules file URL.
func (o NetworkTapRuleOutput) TapRulesUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkTapRule) pulumi.StringPtrOutput { return v.TapRulesUrl }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o NetworkTapRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkTapRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkTapRuleOutput{})
}
