// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Data connector to pull Threat intelligence data from TAXII 2.0/2.1 server
type TiTaxiiDataConnector struct {
	pulumi.CustomResourceState

	// The collection id of the TAXII server.
	CollectionId pulumi.StringPtrOutput `pulumi:"collectionId"`
	// The available data types for Threat Intelligence TAXII data connector.
	DataTypes TiTaxiiDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The friendly name for the TAXII server.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// The kind of the data connector
	// Expected value is 'ThreatIntelligenceTaxii'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The password for the TAXII server.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The polling frequency for the TAXII server.
	PollingFrequency pulumi.StringOutput `pulumi:"pollingFrequency"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The lookback period for the TAXII server.
	TaxiiLookbackPeriod pulumi.StringPtrOutput `pulumi:"taxiiLookbackPeriod"`
	// The API root for the TAXII server.
	TaxiiServer pulumi.StringPtrOutput `pulumi:"taxiiServer"`
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The userName for the TAXII server.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
	// The workspace id.
	WorkspaceId pulumi.StringPtrOutput `pulumi:"workspaceId"`
}

// NewTiTaxiiDataConnector registers a new resource with the given unique name, arguments, and options.
func NewTiTaxiiDataConnector(ctx *pulumi.Context,
	name string, args *TiTaxiiDataConnectorArgs, opts ...pulumi.ResourceOption) (*TiTaxiiDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataTypes == nil {
		return nil, errors.New("invalid value for required argument 'DataTypes'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.PollingFrequency == nil {
		return nil, errors.New("invalid value for required argument 'PollingFrequency'")
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
	args.Kind = pulumi.String("ThreatIntelligenceTaxii")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:TiTaxiiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:TiTaxiiDataConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TiTaxiiDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20230701preview:TiTaxiiDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTiTaxiiDataConnector gets an existing TiTaxiiDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTiTaxiiDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TiTaxiiDataConnectorState, opts ...pulumi.ResourceOption) (*TiTaxiiDataConnector, error) {
	var resource TiTaxiiDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20230701preview:TiTaxiiDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TiTaxiiDataConnector resources.
type tiTaxiiDataConnectorState struct {
}

type TiTaxiiDataConnectorState struct {
}

func (TiTaxiiDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*tiTaxiiDataConnectorState)(nil)).Elem()
}

type tiTaxiiDataConnectorArgs struct {
	// The collection id of the TAXII server.
	CollectionId *string `pulumi:"collectionId"`
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for Threat Intelligence TAXII data connector.
	DataTypes TiTaxiiDataConnectorDataTypes `pulumi:"dataTypes"`
	// The friendly name for the TAXII server.
	FriendlyName *string `pulumi:"friendlyName"`
	// The kind of the data connector
	// Expected value is 'ThreatIntelligenceTaxii'.
	Kind string `pulumi:"kind"`
	// The password for the TAXII server.
	Password *string `pulumi:"password"`
	// The polling frequency for the TAXII server.
	PollingFrequency string `pulumi:"pollingFrequency"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The lookback period for the TAXII server.
	TaxiiLookbackPeriod *string `pulumi:"taxiiLookbackPeriod"`
	// The API root for the TAXII server.
	TaxiiServer *string `pulumi:"taxiiServer"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The userName for the TAXII server.
	UserName *string `pulumi:"userName"`
	// The workspace id.
	WorkspaceId *string `pulumi:"workspaceId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a TiTaxiiDataConnector resource.
type TiTaxiiDataConnectorArgs struct {
	// The collection id of the TAXII server.
	CollectionId pulumi.StringPtrInput
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for Threat Intelligence TAXII data connector.
	DataTypes TiTaxiiDataConnectorDataTypesInput
	// The friendly name for the TAXII server.
	FriendlyName pulumi.StringPtrInput
	// The kind of the data connector
	// Expected value is 'ThreatIntelligenceTaxii'.
	Kind pulumi.StringInput
	// The password for the TAXII server.
	Password pulumi.StringPtrInput
	// The polling frequency for the TAXII server.
	PollingFrequency pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The lookback period for the TAXII server.
	TaxiiLookbackPeriod pulumi.StringPtrInput
	// The API root for the TAXII server.
	TaxiiServer pulumi.StringPtrInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringInput
	// The userName for the TAXII server.
	UserName pulumi.StringPtrInput
	// The workspace id.
	WorkspaceId pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (TiTaxiiDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tiTaxiiDataConnectorArgs)(nil)).Elem()
}

type TiTaxiiDataConnectorInput interface {
	pulumi.Input

	ToTiTaxiiDataConnectorOutput() TiTaxiiDataConnectorOutput
	ToTiTaxiiDataConnectorOutputWithContext(ctx context.Context) TiTaxiiDataConnectorOutput
}

func (*TiTaxiiDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**TiTaxiiDataConnector)(nil)).Elem()
}

func (i *TiTaxiiDataConnector) ToTiTaxiiDataConnectorOutput() TiTaxiiDataConnectorOutput {
	return i.ToTiTaxiiDataConnectorOutputWithContext(context.Background())
}

func (i *TiTaxiiDataConnector) ToTiTaxiiDataConnectorOutputWithContext(ctx context.Context) TiTaxiiDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiTaxiiDataConnectorOutput)
}

func (i *TiTaxiiDataConnector) ToOutput(ctx context.Context) pulumix.Output[*TiTaxiiDataConnector] {
	return pulumix.Output[*TiTaxiiDataConnector]{
		OutputState: i.ToTiTaxiiDataConnectorOutputWithContext(ctx).OutputState,
	}
}

type TiTaxiiDataConnectorOutput struct{ *pulumi.OutputState }

func (TiTaxiiDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TiTaxiiDataConnector)(nil)).Elem()
}

func (o TiTaxiiDataConnectorOutput) ToTiTaxiiDataConnectorOutput() TiTaxiiDataConnectorOutput {
	return o
}

func (o TiTaxiiDataConnectorOutput) ToTiTaxiiDataConnectorOutputWithContext(ctx context.Context) TiTaxiiDataConnectorOutput {
	return o
}

func (o TiTaxiiDataConnectorOutput) ToOutput(ctx context.Context) pulumix.Output[*TiTaxiiDataConnector] {
	return pulumix.Output[*TiTaxiiDataConnector]{
		OutputState: o.OutputState,
	}
}

// The collection id of the TAXII server.
func (o TiTaxiiDataConnectorOutput) CollectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) pulumi.StringPtrOutput { return v.CollectionId }).(pulumi.StringPtrOutput)
}

// The available data types for Threat Intelligence TAXII data connector.
func (o TiTaxiiDataConnectorOutput) DataTypes() TiTaxiiDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) TiTaxiiDataConnectorDataTypesResponseOutput { return v.DataTypes }).(TiTaxiiDataConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o TiTaxiiDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The friendly name for the TAXII server.
func (o TiTaxiiDataConnectorOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'ThreatIntelligenceTaxii'.
func (o TiTaxiiDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o TiTaxiiDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The password for the TAXII server.
func (o TiTaxiiDataConnectorOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The polling frequency for the TAXII server.
func (o TiTaxiiDataConnectorOutput) PollingFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) pulumi.StringOutput { return v.PollingFrequency }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o TiTaxiiDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The lookback period for the TAXII server.
func (o TiTaxiiDataConnectorOutput) TaxiiLookbackPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) pulumi.StringPtrOutput { return v.TaxiiLookbackPeriod }).(pulumi.StringPtrOutput)
}

// The API root for the TAXII server.
func (o TiTaxiiDataConnectorOutput) TaxiiServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) pulumi.StringPtrOutput { return v.TaxiiServer }).(pulumi.StringPtrOutput)
}

// The tenant id to connect to, and get the data from.
func (o TiTaxiiDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o TiTaxiiDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The userName for the TAXII server.
func (o TiTaxiiDataConnectorOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

// The workspace id.
func (o TiTaxiiDataConnectorOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnector) pulumi.StringPtrOutput { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(TiTaxiiDataConnectorOutput{})
}
