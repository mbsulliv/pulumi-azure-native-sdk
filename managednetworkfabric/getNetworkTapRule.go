// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managednetworkfabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get Network Tap Rule resource details.
// Azure REST API version: 2023-06-15.
func LookupNetworkTapRule(ctx *pulumi.Context, args *LookupNetworkTapRuleArgs, opts ...pulumi.InvokeOption) (*LookupNetworkTapRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkTapRuleResult
	err := ctx.Invoke("azure-native:managednetworkfabric:getNetworkTapRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNetworkTapRuleArgs struct {
	// Name of the Network Tap Rule.
	NetworkTapRuleName string `pulumi:"networkTapRuleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The NetworkTapRule resource definition.
type LookupNetworkTapRuleResult struct {
	// Administrative state of the resource.
	AdministrativeState string `pulumi:"administrativeState"`
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Configuration state of the resource.
	ConfigurationState string `pulumi:"configurationState"`
	// Input method to configure Network Tap Rule.
	ConfigurationType string `pulumi:"configurationType"`
	// List of dynamic match configurations.
	DynamicMatchConfigurations []CommonDynamicMatchConfigurationResponse `pulumi:"dynamicMatchConfigurations"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The last sync timestamp.
	LastSyncedTime string `pulumi:"lastSyncedTime"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// List of match configurations.
	MatchConfigurations []NetworkTapRuleMatchConfigurationResponse `pulumi:"matchConfigurations"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The ARM resource Id of the NetworkTap.
	NetworkTapId string `pulumi:"networkTapId"`
	// Polling interval in seconds.
	PollingIntervalInSeconds *int `pulumi:"pollingIntervalInSeconds"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Network Tap Rules file URL.
	TapRulesUrl *string `pulumi:"tapRulesUrl"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupNetworkTapRuleResult
func (val *LookupNetworkTapRuleResult) Defaults() *LookupNetworkTapRuleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.PollingIntervalInSeconds == nil {
		pollingIntervalInSeconds_ := 30
		tmp.PollingIntervalInSeconds = &pollingIntervalInSeconds_
	}
	return &tmp
}

func LookupNetworkTapRuleOutput(ctx *pulumi.Context, args LookupNetworkTapRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkTapRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkTapRuleResult, error) {
			args := v.(LookupNetworkTapRuleArgs)
			r, err := LookupNetworkTapRule(ctx, &args, opts...)
			var s LookupNetworkTapRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkTapRuleResultOutput)
}

type LookupNetworkTapRuleOutputArgs struct {
	// Name of the Network Tap Rule.
	NetworkTapRuleName pulumi.StringInput `pulumi:"networkTapRuleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkTapRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkTapRuleArgs)(nil)).Elem()
}

// The NetworkTapRule resource definition.
type LookupNetworkTapRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkTapRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkTapRuleResult)(nil)).Elem()
}

func (o LookupNetworkTapRuleResultOutput) ToLookupNetworkTapRuleResultOutput() LookupNetworkTapRuleResultOutput {
	return o
}

func (o LookupNetworkTapRuleResultOutput) ToLookupNetworkTapRuleResultOutputWithContext(ctx context.Context) LookupNetworkTapRuleResultOutput {
	return o
}

// Administrative state of the resource.
func (o LookupNetworkTapRuleResultOutput) AdministrativeState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) string { return v.AdministrativeState }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o LookupNetworkTapRuleResultOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) *string { return v.Annotation }).(pulumi.StringPtrOutput)
}

// Configuration state of the resource.
func (o LookupNetworkTapRuleResultOutput) ConfigurationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) string { return v.ConfigurationState }).(pulumi.StringOutput)
}

// Input method to configure Network Tap Rule.
func (o LookupNetworkTapRuleResultOutput) ConfigurationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) string { return v.ConfigurationType }).(pulumi.StringOutput)
}

// List of dynamic match configurations.
func (o LookupNetworkTapRuleResultOutput) DynamicMatchConfigurations() CommonDynamicMatchConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) []CommonDynamicMatchConfigurationResponse {
		return v.DynamicMatchConfigurations
	}).(CommonDynamicMatchConfigurationResponseArrayOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupNetworkTapRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The last sync timestamp.
func (o LookupNetworkTapRuleResultOutput) LastSyncedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) string { return v.LastSyncedTime }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupNetworkTapRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

// List of match configurations.
func (o LookupNetworkTapRuleResultOutput) MatchConfigurations() NetworkTapRuleMatchConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) []NetworkTapRuleMatchConfigurationResponse {
		return v.MatchConfigurations
	}).(NetworkTapRuleMatchConfigurationResponseArrayOutput)
}

// The name of the resource
func (o LookupNetworkTapRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ARM resource Id of the NetworkTap.
func (o LookupNetworkTapRuleResultOutput) NetworkTapId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) string { return v.NetworkTapId }).(pulumi.StringOutput)
}

// Polling interval in seconds.
func (o LookupNetworkTapRuleResultOutput) PollingIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) *int { return v.PollingIntervalInSeconds }).(pulumi.IntPtrOutput)
}

// Provisioning state of the resource.
func (o LookupNetworkTapRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupNetworkTapRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupNetworkTapRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Network Tap Rules file URL.
func (o LookupNetworkTapRuleResultOutput) TapRulesUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) *string { return v.TapRulesUrl }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNetworkTapRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkTapRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkTapRuleResultOutput{})
}