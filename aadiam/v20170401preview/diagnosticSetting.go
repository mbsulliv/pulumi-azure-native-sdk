// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The diagnostic setting resource.
type DiagnosticSetting struct {
	pulumi.CustomResourceState

	// The resource Id for the event hub authorization rule.
	EventHubAuthorizationRuleId pulumi.StringPtrOutput `pulumi:"eventHubAuthorizationRuleId"`
	// The name of the event hub. If none is specified, the default event hub will be selected.
	EventHubName pulumi.StringPtrOutput `pulumi:"eventHubName"`
	// The list of logs settings.
	Logs LogSettingsResponseArrayOutput `pulumi:"logs"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
	ServiceBusRuleId pulumi.StringPtrOutput `pulumi:"serviceBusRuleId"`
	// The resource ID of the storage account to which you would like to send Diagnostic Logs.
	StorageAccountId pulumi.StringPtrOutput `pulumi:"storageAccountId"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The workspace ID (resource ID of a Log Analytics workspace) for a Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
	WorkspaceId pulumi.StringPtrOutput `pulumi:"workspaceId"`
}

// NewDiagnosticSetting registers a new resource with the given unique name, arguments, and options.
func NewDiagnosticSetting(ctx *pulumi.Context,
	name string, args *DiagnosticSettingArgs, opts ...pulumi.ResourceOption) (*DiagnosticSetting, error) {
	if args == nil {
		args = &DiagnosticSettingArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:aadiam:DiagnosticSetting"),
		},
		{
			Type: pulumi.String("azure-native:aadiam/v20170401:DiagnosticSetting"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DiagnosticSetting
	err := ctx.RegisterResource("azure-native:aadiam/v20170401preview:DiagnosticSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiagnosticSetting gets an existing DiagnosticSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiagnosticSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiagnosticSettingState, opts ...pulumi.ResourceOption) (*DiagnosticSetting, error) {
	var resource DiagnosticSetting
	err := ctx.ReadResource("azure-native:aadiam/v20170401preview:DiagnosticSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiagnosticSetting resources.
type diagnosticSettingState struct {
}

type DiagnosticSettingState struct {
}

func (DiagnosticSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticSettingState)(nil)).Elem()
}

type diagnosticSettingArgs struct {
	// The resource Id for the event hub authorization rule.
	EventHubAuthorizationRuleId *string `pulumi:"eventHubAuthorizationRuleId"`
	// The name of the event hub. If none is specified, the default event hub will be selected.
	EventHubName *string `pulumi:"eventHubName"`
	// The list of logs settings.
	Logs []LogSettings `pulumi:"logs"`
	// The name of the diagnostic setting.
	Name *string `pulumi:"name"`
	// The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
	ServiceBusRuleId *string `pulumi:"serviceBusRuleId"`
	// The resource ID of the storage account to which you would like to send Diagnostic Logs.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// The workspace ID (resource ID of a Log Analytics workspace) for a Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
	WorkspaceId *string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DiagnosticSetting resource.
type DiagnosticSettingArgs struct {
	// The resource Id for the event hub authorization rule.
	EventHubAuthorizationRuleId pulumi.StringPtrInput
	// The name of the event hub. If none is specified, the default event hub will be selected.
	EventHubName pulumi.StringPtrInput
	// The list of logs settings.
	Logs LogSettingsArrayInput
	// The name of the diagnostic setting.
	Name pulumi.StringPtrInput
	// The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
	ServiceBusRuleId pulumi.StringPtrInput
	// The resource ID of the storage account to which you would like to send Diagnostic Logs.
	StorageAccountId pulumi.StringPtrInput
	// The workspace ID (resource ID of a Log Analytics workspace) for a Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
	WorkspaceId pulumi.StringPtrInput
}

func (DiagnosticSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticSettingArgs)(nil)).Elem()
}

type DiagnosticSettingInput interface {
	pulumi.Input

	ToDiagnosticSettingOutput() DiagnosticSettingOutput
	ToDiagnosticSettingOutputWithContext(ctx context.Context) DiagnosticSettingOutput
}

func (*DiagnosticSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticSetting)(nil)).Elem()
}

func (i *DiagnosticSetting) ToDiagnosticSettingOutput() DiagnosticSettingOutput {
	return i.ToDiagnosticSettingOutputWithContext(context.Background())
}

func (i *DiagnosticSetting) ToDiagnosticSettingOutputWithContext(ctx context.Context) DiagnosticSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticSettingOutput)
}

type DiagnosticSettingOutput struct{ *pulumi.OutputState }

func (DiagnosticSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticSetting)(nil)).Elem()
}

func (o DiagnosticSettingOutput) ToDiagnosticSettingOutput() DiagnosticSettingOutput {
	return o
}

func (o DiagnosticSettingOutput) ToDiagnosticSettingOutputWithContext(ctx context.Context) DiagnosticSettingOutput {
	return o
}

// The resource Id for the event hub authorization rule.
func (o DiagnosticSettingOutput) EventHubAuthorizationRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticSetting) pulumi.StringPtrOutput { return v.EventHubAuthorizationRuleId }).(pulumi.StringPtrOutput)
}

// The name of the event hub. If none is specified, the default event hub will be selected.
func (o DiagnosticSettingOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticSetting) pulumi.StringPtrOutput { return v.EventHubName }).(pulumi.StringPtrOutput)
}

// The list of logs settings.
func (o DiagnosticSettingOutput) Logs() LogSettingsResponseArrayOutput {
	return o.ApplyT(func(v *DiagnosticSetting) LogSettingsResponseArrayOutput { return v.Logs }).(LogSettingsResponseArrayOutput)
}

// Azure resource name
func (o DiagnosticSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DiagnosticSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
func (o DiagnosticSettingOutput) ServiceBusRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticSetting) pulumi.StringPtrOutput { return v.ServiceBusRuleId }).(pulumi.StringPtrOutput)
}

// The resource ID of the storage account to which you would like to send Diagnostic Logs.
func (o DiagnosticSettingOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticSetting) pulumi.StringPtrOutput { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

// Azure resource type
func (o DiagnosticSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DiagnosticSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The workspace ID (resource ID of a Log Analytics workspace) for a Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
func (o DiagnosticSettingOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticSetting) pulumi.StringPtrOutput { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DiagnosticSettingOutput{})
}
