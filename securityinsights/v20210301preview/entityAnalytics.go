// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Settings with single toggle.
type EntityAnalytics struct {
	pulumi.CustomResourceState

	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Determines whether the setting is enable or disabled.
	IsEnabled pulumi.BoolOutput `pulumi:"isEnabled"`
	// The kind of the setting
	// Expected value is 'EntityAnalytics'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEntityAnalytics registers a new resource with the given unique name, arguments, and options.
func NewEntityAnalytics(ctx *pulumi.Context,
	name string, args *EntityAnalyticsArgs, opts ...pulumi.ResourceOption) (*EntityAnalytics, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("EntityAnalytics")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:EntityAnalytics"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EntityAnalytics
	err := ctx.RegisterResource("azure-native:securityinsights/v20210301preview:EntityAnalytics", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntityAnalytics gets an existing EntityAnalytics resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntityAnalytics(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntityAnalyticsState, opts ...pulumi.ResourceOption) (*EntityAnalytics, error) {
	var resource EntityAnalytics
	err := ctx.ReadResource("azure-native:securityinsights/v20210301preview:EntityAnalytics", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntityAnalytics resources.
type entityAnalyticsState struct {
}

type EntityAnalyticsState struct {
}

func (EntityAnalyticsState) ElementType() reflect.Type {
	return reflect.TypeOf((*entityAnalyticsState)(nil)).Elem()
}

type entityAnalyticsArgs struct {
	// The kind of the setting
	// Expected value is 'EntityAnalytics'.
	Kind string `pulumi:"kind"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName *string `pulumi:"settingsName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a EntityAnalytics resource.
type EntityAnalyticsArgs struct {
	// The kind of the setting
	// Expected value is 'EntityAnalytics'.
	Kind pulumi.StringInput
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (EntityAnalyticsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entityAnalyticsArgs)(nil)).Elem()
}

type EntityAnalyticsInput interface {
	pulumi.Input

	ToEntityAnalyticsOutput() EntityAnalyticsOutput
	ToEntityAnalyticsOutputWithContext(ctx context.Context) EntityAnalyticsOutput
}

func (*EntityAnalytics) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityAnalytics)(nil)).Elem()
}

func (i *EntityAnalytics) ToEntityAnalyticsOutput() EntityAnalyticsOutput {
	return i.ToEntityAnalyticsOutputWithContext(context.Background())
}

func (i *EntityAnalytics) ToEntityAnalyticsOutputWithContext(ctx context.Context) EntityAnalyticsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityAnalyticsOutput)
}

type EntityAnalyticsOutput struct{ *pulumi.OutputState }

func (EntityAnalyticsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityAnalytics)(nil)).Elem()
}

func (o EntityAnalyticsOutput) ToEntityAnalyticsOutput() EntityAnalyticsOutput {
	return o
}

func (o EntityAnalyticsOutput) ToEntityAnalyticsOutputWithContext(ctx context.Context) EntityAnalyticsOutput {
	return o
}

// Etag of the azure resource
func (o EntityAnalyticsOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EntityAnalytics) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Determines whether the setting is enable or disabled.
func (o EntityAnalyticsOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *EntityAnalytics) pulumi.BoolOutput { return v.IsEnabled }).(pulumi.BoolOutput)
}

// The kind of the setting
// Expected value is 'EntityAnalytics'.
func (o EntityAnalyticsOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityAnalytics) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Azure resource name
func (o EntityAnalyticsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityAnalytics) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o EntityAnalyticsOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EntityAnalytics) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource type
func (o EntityAnalyticsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityAnalytics) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EntityAnalyticsOutput{})
}
