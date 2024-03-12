// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Anomaly Security ML Analytics Settings
type AnomalySecurityMLAnalyticsSettings struct {
	pulumi.CustomResourceState

	// The anomaly settings version of the Anomaly security ml analytics settings that dictates whether job version gets updated or not.
	AnomalySettingsVersion pulumi.IntPtrOutput `pulumi:"anomalySettingsVersion"`
	// The anomaly version of the AnomalySecurityMLAnalyticsSettings.
	AnomalyVersion pulumi.StringOutput `pulumi:"anomalyVersion"`
	// The customizable observations of the AnomalySecurityMLAnalyticsSettings.
	CustomizableObservations pulumi.AnyOutput `pulumi:"customizableObservations"`
	// The description of the SecurityMLAnalyticsSettings.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name for settings created by this SecurityMLAnalyticsSettings.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Determines whether this settings is enabled or disabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The frequency that this SecurityMLAnalyticsSettings will be run.
	Frequency pulumi.StringOutput `pulumi:"frequency"`
	// Determines whether this anomaly security ml analytics settings is a default settings
	IsDefaultSettings pulumi.BoolOutput `pulumi:"isDefaultSettings"`
	// The kind of security ML analytics settings
	// Expected value is 'Anomaly'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The last time that this SecurityMLAnalyticsSettings has been modified.
	LastModifiedUtc pulumi.StringOutput `pulumi:"lastModifiedUtc"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The required data sources for this SecurityMLAnalyticsSettings
	RequiredDataConnectors SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput `pulumi:"requiredDataConnectors"`
	// The anomaly settings definition Id
	SettingsDefinitionId pulumi.StringPtrOutput `pulumi:"settingsDefinitionId"`
	// The anomaly SecurityMLAnalyticsSettings status
	SettingsStatus pulumi.StringOutput `pulumi:"settingsStatus"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tactics of the SecurityMLAnalyticsSettings
	Tactics pulumi.StringArrayOutput `pulumi:"tactics"`
	// The techniques of the SecurityMLAnalyticsSettings
	Techniques pulumi.StringArrayOutput `pulumi:"techniques"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAnomalySecurityMLAnalyticsSettings registers a new resource with the given unique name, arguments, and options.
func NewAnomalySecurityMLAnalyticsSettings(ctx *pulumi.Context,
	name string, args *AnomalySecurityMLAnalyticsSettingsArgs, opts ...pulumi.ResourceOption) (*AnomalySecurityMLAnalyticsSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AnomalyVersion == nil {
		return nil, errors.New("invalid value for required argument 'AnomalyVersion'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.IsDefaultSettings == nil {
		return nil, errors.New("invalid value for required argument 'IsDefaultSettings'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SettingsStatus == nil {
		return nil, errors.New("invalid value for required argument 'SettingsStatus'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("Anomaly")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:AnomalySecurityMLAnalyticsSettings"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AnomalySecurityMLAnalyticsSettings
	err := ctx.RegisterResource("azure-native:securityinsights/v20230801preview:AnomalySecurityMLAnalyticsSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnomalySecurityMLAnalyticsSettings gets an existing AnomalySecurityMLAnalyticsSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnomalySecurityMLAnalyticsSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnomalySecurityMLAnalyticsSettingsState, opts ...pulumi.ResourceOption) (*AnomalySecurityMLAnalyticsSettings, error) {
	var resource AnomalySecurityMLAnalyticsSettings
	err := ctx.ReadResource("azure-native:securityinsights/v20230801preview:AnomalySecurityMLAnalyticsSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnomalySecurityMLAnalyticsSettings resources.
type anomalySecurityMLAnalyticsSettingsState struct {
}

type AnomalySecurityMLAnalyticsSettingsState struct {
}

func (AnomalySecurityMLAnalyticsSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*anomalySecurityMLAnalyticsSettingsState)(nil)).Elem()
}

type anomalySecurityMLAnalyticsSettingsArgs struct {
	// The anomaly settings version of the Anomaly security ml analytics settings that dictates whether job version gets updated or not.
	AnomalySettingsVersion *int `pulumi:"anomalySettingsVersion"`
	// The anomaly version of the AnomalySecurityMLAnalyticsSettings.
	AnomalyVersion string `pulumi:"anomalyVersion"`
	// The customizable observations of the AnomalySecurityMLAnalyticsSettings.
	CustomizableObservations interface{} `pulumi:"customizableObservations"`
	// The description of the SecurityMLAnalyticsSettings.
	Description *string `pulumi:"description"`
	// The display name for settings created by this SecurityMLAnalyticsSettings.
	DisplayName string `pulumi:"displayName"`
	// Determines whether this settings is enabled or disabled.
	Enabled bool `pulumi:"enabled"`
	// The frequency that this SecurityMLAnalyticsSettings will be run.
	Frequency string `pulumi:"frequency"`
	// Determines whether this anomaly security ml analytics settings is a default settings
	IsDefaultSettings bool `pulumi:"isDefaultSettings"`
	// The kind of security ML analytics settings
	// Expected value is 'Anomaly'.
	Kind string `pulumi:"kind"`
	// The required data sources for this SecurityMLAnalyticsSettings
	RequiredDataConnectors []SecurityMLAnalyticsSettingsDataSource `pulumi:"requiredDataConnectors"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The anomaly settings definition Id
	SettingsDefinitionId *string `pulumi:"settingsDefinitionId"`
	// Security ML Analytics Settings resource name
	SettingsResourceName *string `pulumi:"settingsResourceName"`
	// The anomaly SecurityMLAnalyticsSettings status
	SettingsStatus string `pulumi:"settingsStatus"`
	// The tactics of the SecurityMLAnalyticsSettings
	Tactics []string `pulumi:"tactics"`
	// The techniques of the SecurityMLAnalyticsSettings
	Techniques []string `pulumi:"techniques"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a AnomalySecurityMLAnalyticsSettings resource.
type AnomalySecurityMLAnalyticsSettingsArgs struct {
	// The anomaly settings version of the Anomaly security ml analytics settings that dictates whether job version gets updated or not.
	AnomalySettingsVersion pulumi.IntPtrInput
	// The anomaly version of the AnomalySecurityMLAnalyticsSettings.
	AnomalyVersion pulumi.StringInput
	// The customizable observations of the AnomalySecurityMLAnalyticsSettings.
	CustomizableObservations pulumi.Input
	// The description of the SecurityMLAnalyticsSettings.
	Description pulumi.StringPtrInput
	// The display name for settings created by this SecurityMLAnalyticsSettings.
	DisplayName pulumi.StringInput
	// Determines whether this settings is enabled or disabled.
	Enabled pulumi.BoolInput
	// The frequency that this SecurityMLAnalyticsSettings will be run.
	Frequency pulumi.StringInput
	// Determines whether this anomaly security ml analytics settings is a default settings
	IsDefaultSettings pulumi.BoolInput
	// The kind of security ML analytics settings
	// Expected value is 'Anomaly'.
	Kind pulumi.StringInput
	// The required data sources for this SecurityMLAnalyticsSettings
	RequiredDataConnectors SecurityMLAnalyticsSettingsDataSourceArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The anomaly settings definition Id
	SettingsDefinitionId pulumi.StringPtrInput
	// Security ML Analytics Settings resource name
	SettingsResourceName pulumi.StringPtrInput
	// The anomaly SecurityMLAnalyticsSettings status
	SettingsStatus pulumi.StringInput
	// The tactics of the SecurityMLAnalyticsSettings
	Tactics pulumi.StringArrayInput
	// The techniques of the SecurityMLAnalyticsSettings
	Techniques pulumi.StringArrayInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (AnomalySecurityMLAnalyticsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*anomalySecurityMLAnalyticsSettingsArgs)(nil)).Elem()
}

type AnomalySecurityMLAnalyticsSettingsInput interface {
	pulumi.Input

	ToAnomalySecurityMLAnalyticsSettingsOutput() AnomalySecurityMLAnalyticsSettingsOutput
	ToAnomalySecurityMLAnalyticsSettingsOutputWithContext(ctx context.Context) AnomalySecurityMLAnalyticsSettingsOutput
}

func (*AnomalySecurityMLAnalyticsSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**AnomalySecurityMLAnalyticsSettings)(nil)).Elem()
}

func (i *AnomalySecurityMLAnalyticsSettings) ToAnomalySecurityMLAnalyticsSettingsOutput() AnomalySecurityMLAnalyticsSettingsOutput {
	return i.ToAnomalySecurityMLAnalyticsSettingsOutputWithContext(context.Background())
}

func (i *AnomalySecurityMLAnalyticsSettings) ToAnomalySecurityMLAnalyticsSettingsOutputWithContext(ctx context.Context) AnomalySecurityMLAnalyticsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnomalySecurityMLAnalyticsSettingsOutput)
}

type AnomalySecurityMLAnalyticsSettingsOutput struct{ *pulumi.OutputState }

func (AnomalySecurityMLAnalyticsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnomalySecurityMLAnalyticsSettings)(nil)).Elem()
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) ToAnomalySecurityMLAnalyticsSettingsOutput() AnomalySecurityMLAnalyticsSettingsOutput {
	return o
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) ToAnomalySecurityMLAnalyticsSettingsOutputWithContext(ctx context.Context) AnomalySecurityMLAnalyticsSettingsOutput {
	return o
}

// The anomaly settings version of the Anomaly security ml analytics settings that dictates whether job version gets updated or not.
func (o AnomalySecurityMLAnalyticsSettingsOutput) AnomalySettingsVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.IntPtrOutput { return v.AnomalySettingsVersion }).(pulumi.IntPtrOutput)
}

// The anomaly version of the AnomalySecurityMLAnalyticsSettings.
func (o AnomalySecurityMLAnalyticsSettingsOutput) AnomalyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.AnomalyVersion }).(pulumi.StringOutput)
}

// The customizable observations of the AnomalySecurityMLAnalyticsSettings.
func (o AnomalySecurityMLAnalyticsSettingsOutput) CustomizableObservations() pulumi.AnyOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.AnyOutput { return v.CustomizableObservations }).(pulumi.AnyOutput)
}

// The description of the SecurityMLAnalyticsSettings.
func (o AnomalySecurityMLAnalyticsSettingsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name for settings created by this SecurityMLAnalyticsSettings.
func (o AnomalySecurityMLAnalyticsSettingsOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Determines whether this settings is enabled or disabled.
func (o AnomalySecurityMLAnalyticsSettingsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Etag of the azure resource
func (o AnomalySecurityMLAnalyticsSettingsOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The frequency that this SecurityMLAnalyticsSettings will be run.
func (o AnomalySecurityMLAnalyticsSettingsOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.Frequency }).(pulumi.StringOutput)
}

// Determines whether this anomaly security ml analytics settings is a default settings
func (o AnomalySecurityMLAnalyticsSettingsOutput) IsDefaultSettings() pulumi.BoolOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.BoolOutput { return v.IsDefaultSettings }).(pulumi.BoolOutput)
}

// The kind of security ML analytics settings
// Expected value is 'Anomaly'.
func (o AnomalySecurityMLAnalyticsSettingsOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The last time that this SecurityMLAnalyticsSettings has been modified.
func (o AnomalySecurityMLAnalyticsSettingsOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o AnomalySecurityMLAnalyticsSettingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The required data sources for this SecurityMLAnalyticsSettings
func (o AnomalySecurityMLAnalyticsSettingsOutput) RequiredDataConnectors() SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput {
		return v.RequiredDataConnectors
	}).(SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput)
}

// The anomaly settings definition Id
func (o AnomalySecurityMLAnalyticsSettingsOutput) SettingsDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringPtrOutput { return v.SettingsDefinitionId }).(pulumi.StringPtrOutput)
}

// The anomaly SecurityMLAnalyticsSettings status
func (o AnomalySecurityMLAnalyticsSettingsOutput) SettingsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.SettingsStatus }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AnomalySecurityMLAnalyticsSettingsOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tactics of the SecurityMLAnalyticsSettings
func (o AnomalySecurityMLAnalyticsSettingsOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

// The techniques of the SecurityMLAnalyticsSettings
func (o AnomalySecurityMLAnalyticsSettingsOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringArrayOutput { return v.Techniques }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AnomalySecurityMLAnalyticsSettingsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AnomalySecurityMLAnalyticsSettingsOutput{})
}
