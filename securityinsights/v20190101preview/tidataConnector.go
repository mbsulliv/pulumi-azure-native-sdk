// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data connector to pull threat intelligence data from TIP products.
type TIDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes TIDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Expected value is 'ThreatIntelligence'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The lookback period for the feed to be imported.
	TipLookbackPeriod pulumi.StringPtrOutput `pulumi:"tipLookbackPeriod"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTIDataConnector registers a new resource with the given unique name, arguments, and options.
func NewTIDataConnector(ctx *pulumi.Context,
	name string, args *TIDataConnectorArgs, opts ...pulumi.ResourceOption) (*TIDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataTypes == nil {
		return nil, errors.New("invalid value for required argument 'DataTypes'")
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
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("ThreatIntelligence")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:TIDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource TIDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20190101preview:TIDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTIDataConnector gets an existing TIDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTIDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TIDataConnectorState, opts ...pulumi.ResourceOption) (*TIDataConnector, error) {
	var resource TIDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20190101preview:TIDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TIDataConnector resources.
type tidataConnectorState struct {
}

type TIDataConnectorState struct {
}

func (TIDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*tidataConnectorState)(nil)).Elem()
}

type tidataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes TIDataConnectorDataTypes `pulumi:"dataTypes"`
	// Expected value is 'ThreatIntelligence'.
	Kind string `pulumi:"kind"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The lookback period for the feed to be imported.
	TipLookbackPeriod *string `pulumi:"tipLookbackPeriod"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a TIDataConnector resource.
type TIDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes TIDataConnectorDataTypesInput
	// Expected value is 'ThreatIntelligence'.
	Kind pulumi.StringInput
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringInput
	// The lookback period for the feed to be imported.
	TipLookbackPeriod pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (TIDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tidataConnectorArgs)(nil)).Elem()
}

type TIDataConnectorInput interface {
	pulumi.Input

	ToTIDataConnectorOutput() TIDataConnectorOutput
	ToTIDataConnectorOutputWithContext(ctx context.Context) TIDataConnectorOutput
}

func (*TIDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**TIDataConnector)(nil)).Elem()
}

func (i *TIDataConnector) ToTIDataConnectorOutput() TIDataConnectorOutput {
	return i.ToTIDataConnectorOutputWithContext(context.Background())
}

func (i *TIDataConnector) ToTIDataConnectorOutputWithContext(ctx context.Context) TIDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TIDataConnectorOutput)
}

type TIDataConnectorOutput struct{ *pulumi.OutputState }

func (TIDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TIDataConnector)(nil)).Elem()
}

func (o TIDataConnectorOutput) ToTIDataConnectorOutput() TIDataConnectorOutput {
	return o
}

func (o TIDataConnectorOutput) ToTIDataConnectorOutputWithContext(ctx context.Context) TIDataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o TIDataConnectorOutput) DataTypes() TIDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *TIDataConnector) TIDataConnectorDataTypesResponseOutput { return v.DataTypes }).(TIDataConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o TIDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TIDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Expected value is 'ThreatIntelligence'.
func (o TIDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *TIDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Azure resource name
func (o TIDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TIDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The tenant id to connect to, and get the data from.
func (o TIDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *TIDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The lookback period for the feed to be imported.
func (o TIDataConnectorOutput) TipLookbackPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TIDataConnector) pulumi.StringPtrOutput { return v.TipLookbackPeriod }).(pulumi.StringPtrOutput)
}

// Azure resource type
func (o TIDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TIDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TIDataConnectorOutput{})
}
