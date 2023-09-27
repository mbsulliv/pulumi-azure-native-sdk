// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Settings with single toggle.
type IPSyncer struct {
	pulumi.CustomResourceState

	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Determines whether the setting is enable or disabled.
	IsEnabled pulumi.BoolOutput `pulumi:"isEnabled"`
	// Expected value is 'IPSyncer'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIPSyncer registers a new resource with the given unique name, arguments, and options.
func NewIPSyncer(ctx *pulumi.Context,
	name string, args *IPSyncerArgs, opts ...pulumi.ResourceOption) (*IPSyncer, error) {
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
	args.Kind = pulumi.String("IPSyncer")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:IPSyncer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IPSyncer
	err := ctx.RegisterResource("azure-native:securityinsights/v20190101preview:IPSyncer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIPSyncer gets an existing IPSyncer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIPSyncer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPSyncerState, opts ...pulumi.ResourceOption) (*IPSyncer, error) {
	var resource IPSyncer
	err := ctx.ReadResource("azure-native:securityinsights/v20190101preview:IPSyncer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IPSyncer resources.
type ipsyncerState struct {
}

type IPSyncerState struct {
}

func (IPSyncerState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsyncerState)(nil)).Elem()
}

type ipsyncerArgs struct {
	// Expected value is 'IPSyncer'.
	Kind string `pulumi:"kind"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName *string `pulumi:"settingsName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a IPSyncer resource.
type IPSyncerArgs struct {
	// Expected value is 'IPSyncer'.
	Kind pulumi.StringInput
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (IPSyncerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsyncerArgs)(nil)).Elem()
}

type IPSyncerInput interface {
	pulumi.Input

	ToIPSyncerOutput() IPSyncerOutput
	ToIPSyncerOutputWithContext(ctx context.Context) IPSyncerOutput
}

func (*IPSyncer) ElementType() reflect.Type {
	return reflect.TypeOf((**IPSyncer)(nil)).Elem()
}

func (i *IPSyncer) ToIPSyncerOutput() IPSyncerOutput {
	return i.ToIPSyncerOutputWithContext(context.Background())
}

func (i *IPSyncer) ToIPSyncerOutputWithContext(ctx context.Context) IPSyncerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPSyncerOutput)
}

func (i *IPSyncer) ToOutput(ctx context.Context) pulumix.Output[*IPSyncer] {
	return pulumix.Output[*IPSyncer]{
		OutputState: i.ToIPSyncerOutputWithContext(ctx).OutputState,
	}
}

type IPSyncerOutput struct{ *pulumi.OutputState }

func (IPSyncerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPSyncer)(nil)).Elem()
}

func (o IPSyncerOutput) ToIPSyncerOutput() IPSyncerOutput {
	return o
}

func (o IPSyncerOutput) ToIPSyncerOutputWithContext(ctx context.Context) IPSyncerOutput {
	return o
}

func (o IPSyncerOutput) ToOutput(ctx context.Context) pulumix.Output[*IPSyncer] {
	return pulumix.Output[*IPSyncer]{
		OutputState: o.OutputState,
	}
}

// Etag of the azure resource
func (o IPSyncerOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPSyncer) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Determines whether the setting is enable or disabled.
func (o IPSyncerOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *IPSyncer) pulumi.BoolOutput { return v.IsEnabled }).(pulumi.BoolOutput)
}

// Expected value is 'IPSyncer'.
func (o IPSyncerOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *IPSyncer) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Azure resource name
func (o IPSyncerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IPSyncer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure resource type
func (o IPSyncerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IPSyncer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IPSyncerOutput{})
}
