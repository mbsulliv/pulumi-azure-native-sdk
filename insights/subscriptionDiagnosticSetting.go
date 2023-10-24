// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The subscription diagnostic setting resource.
// Azure REST API version: 2021-05-01-preview. Prior API version in Azure Native 1.x: 2017-05-01-preview.
//
// Other available API versions: 2017-05-01-preview.
type SubscriptionDiagnosticSetting struct {
	pulumi.CustomResourceState

	// The resource Id for the event hub authorization rule.
	EventHubAuthorizationRuleId pulumi.StringPtrOutput `pulumi:"eventHubAuthorizationRuleId"`
	// The name of the event hub. If none is specified, the default event hub will be selected.
	EventHubName pulumi.StringPtrOutput `pulumi:"eventHubName"`
	// The list of logs settings.
	Logs SubscriptionLogSettingsResponseArrayOutput `pulumi:"logs"`
	// The full ARM resource ID of the Marketplace resource to which you would like to send Diagnostic Logs.
	MarketplacePartnerId pulumi.StringPtrOutput `pulumi:"marketplacePartnerId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
	ServiceBusRuleId pulumi.StringPtrOutput `pulumi:"serviceBusRuleId"`
	// The resource ID of the storage account to which you would like to send Diagnostic Logs.
	StorageAccountId pulumi.StringPtrOutput `pulumi:"storageAccountId"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
	WorkspaceId pulumi.StringPtrOutput `pulumi:"workspaceId"`
}

// NewSubscriptionDiagnosticSetting registers a new resource with the given unique name, arguments, and options.
func NewSubscriptionDiagnosticSetting(ctx *pulumi.Context,
	name string, args *SubscriptionDiagnosticSettingArgs, opts ...pulumi.ResourceOption) (*SubscriptionDiagnosticSetting, error) {
	if args == nil {
		args = &SubscriptionDiagnosticSettingArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights/v20170501preview:SubscriptionDiagnosticSetting"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210501preview:SubscriptionDiagnosticSetting"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SubscriptionDiagnosticSetting
	err := ctx.RegisterResource("azure-native:insights:SubscriptionDiagnosticSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscriptionDiagnosticSetting gets an existing SubscriptionDiagnosticSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscriptionDiagnosticSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionDiagnosticSettingState, opts ...pulumi.ResourceOption) (*SubscriptionDiagnosticSetting, error) {
	var resource SubscriptionDiagnosticSetting
	err := ctx.ReadResource("azure-native:insights:SubscriptionDiagnosticSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscriptionDiagnosticSetting resources.
type subscriptionDiagnosticSettingState struct {
}

type SubscriptionDiagnosticSettingState struct {
}

func (SubscriptionDiagnosticSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionDiagnosticSettingState)(nil)).Elem()
}

type subscriptionDiagnosticSettingArgs struct {
	// The resource Id for the event hub authorization rule.
	EventHubAuthorizationRuleId *string `pulumi:"eventHubAuthorizationRuleId"`
	// The name of the event hub. If none is specified, the default event hub will be selected.
	EventHubName *string `pulumi:"eventHubName"`
	// The list of logs settings.
	Logs []SubscriptionLogSettings `pulumi:"logs"`
	// The full ARM resource ID of the Marketplace resource to which you would like to send Diagnostic Logs.
	MarketplacePartnerId *string `pulumi:"marketplacePartnerId"`
	// The name of the diagnostic setting.
	Name *string `pulumi:"name"`
	// The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
	ServiceBusRuleId *string `pulumi:"serviceBusRuleId"`
	// The resource ID of the storage account to which you would like to send Diagnostic Logs.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
	WorkspaceId *string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SubscriptionDiagnosticSetting resource.
type SubscriptionDiagnosticSettingArgs struct {
	// The resource Id for the event hub authorization rule.
	EventHubAuthorizationRuleId pulumi.StringPtrInput
	// The name of the event hub. If none is specified, the default event hub will be selected.
	EventHubName pulumi.StringPtrInput
	// The list of logs settings.
	Logs SubscriptionLogSettingsArrayInput
	// The full ARM resource ID of the Marketplace resource to which you would like to send Diagnostic Logs.
	MarketplacePartnerId pulumi.StringPtrInput
	// The name of the diagnostic setting.
	Name pulumi.StringPtrInput
	// The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
	ServiceBusRuleId pulumi.StringPtrInput
	// The resource ID of the storage account to which you would like to send Diagnostic Logs.
	StorageAccountId pulumi.StringPtrInput
	// The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
	WorkspaceId pulumi.StringPtrInput
}

func (SubscriptionDiagnosticSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionDiagnosticSettingArgs)(nil)).Elem()
}

type SubscriptionDiagnosticSettingInput interface {
	pulumi.Input

	ToSubscriptionDiagnosticSettingOutput() SubscriptionDiagnosticSettingOutput
	ToSubscriptionDiagnosticSettingOutputWithContext(ctx context.Context) SubscriptionDiagnosticSettingOutput
}

func (*SubscriptionDiagnosticSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionDiagnosticSetting)(nil)).Elem()
}

func (i *SubscriptionDiagnosticSetting) ToSubscriptionDiagnosticSettingOutput() SubscriptionDiagnosticSettingOutput {
	return i.ToSubscriptionDiagnosticSettingOutputWithContext(context.Background())
}

func (i *SubscriptionDiagnosticSetting) ToSubscriptionDiagnosticSettingOutputWithContext(ctx context.Context) SubscriptionDiagnosticSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionDiagnosticSettingOutput)
}

func (i *SubscriptionDiagnosticSetting) ToOutput(ctx context.Context) pulumix.Output[*SubscriptionDiagnosticSetting] {
	return pulumix.Output[*SubscriptionDiagnosticSetting]{
		OutputState: i.ToSubscriptionDiagnosticSettingOutputWithContext(ctx).OutputState,
	}
}

type SubscriptionDiagnosticSettingOutput struct{ *pulumi.OutputState }

func (SubscriptionDiagnosticSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionDiagnosticSetting)(nil)).Elem()
}

func (o SubscriptionDiagnosticSettingOutput) ToSubscriptionDiagnosticSettingOutput() SubscriptionDiagnosticSettingOutput {
	return o
}

func (o SubscriptionDiagnosticSettingOutput) ToSubscriptionDiagnosticSettingOutputWithContext(ctx context.Context) SubscriptionDiagnosticSettingOutput {
	return o
}

func (o SubscriptionDiagnosticSettingOutput) ToOutput(ctx context.Context) pulumix.Output[*SubscriptionDiagnosticSetting] {
	return pulumix.Output[*SubscriptionDiagnosticSetting]{
		OutputState: o.OutputState,
	}
}

// The resource Id for the event hub authorization rule.
func (o SubscriptionDiagnosticSettingOutput) EventHubAuthorizationRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionDiagnosticSetting) pulumi.StringPtrOutput { return v.EventHubAuthorizationRuleId }).(pulumi.StringPtrOutput)
}

// The name of the event hub. If none is specified, the default event hub will be selected.
func (o SubscriptionDiagnosticSettingOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionDiagnosticSetting) pulumi.StringPtrOutput { return v.EventHubName }).(pulumi.StringPtrOutput)
}

// The list of logs settings.
func (o SubscriptionDiagnosticSettingOutput) Logs() SubscriptionLogSettingsResponseArrayOutput {
	return o.ApplyT(func(v *SubscriptionDiagnosticSetting) SubscriptionLogSettingsResponseArrayOutput { return v.Logs }).(SubscriptionLogSettingsResponseArrayOutput)
}

// The full ARM resource ID of the Marketplace resource to which you would like to send Diagnostic Logs.
func (o SubscriptionDiagnosticSettingOutput) MarketplacePartnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionDiagnosticSetting) pulumi.StringPtrOutput { return v.MarketplacePartnerId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o SubscriptionDiagnosticSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscriptionDiagnosticSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
func (o SubscriptionDiagnosticSettingOutput) ServiceBusRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionDiagnosticSetting) pulumi.StringPtrOutput { return v.ServiceBusRuleId }).(pulumi.StringPtrOutput)
}

// The resource ID of the storage account to which you would like to send Diagnostic Logs.
func (o SubscriptionDiagnosticSettingOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionDiagnosticSetting) pulumi.StringPtrOutput { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

// The system metadata related to this resource.
func (o SubscriptionDiagnosticSettingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SubscriptionDiagnosticSetting) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SubscriptionDiagnosticSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscriptionDiagnosticSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
func (o SubscriptionDiagnosticSettingOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionDiagnosticSetting) pulumi.StringPtrOutput { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SubscriptionDiagnosticSettingOutput{})
}
