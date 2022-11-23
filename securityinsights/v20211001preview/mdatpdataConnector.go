// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents MDATP (Microsoft Defender Advanced Threat Protection) data connector.
type MDATPDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'MicrosoftDefenderAdvancedThreatProtection'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMDATPDataConnector registers a new resource with the given unique name, arguments, and options.
func NewMDATPDataConnector(ctx *pulumi.Context,
	name string, args *MDATPDataConnectorArgs, opts ...pulumi.ResourceOption) (*MDATPDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
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
	args.Kind = pulumi.String("MicrosoftDefenderAdvancedThreatProtection")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:MDATPDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource MDATPDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20211001preview:MDATPDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMDATPDataConnector gets an existing MDATPDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMDATPDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MDATPDataConnectorState, opts ...pulumi.ResourceOption) (*MDATPDataConnector, error) {
	var resource MDATPDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20211001preview:MDATPDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MDATPDataConnector resources.
type mdatpdataConnectorState struct {
}

type MDATPDataConnectorState struct {
}

func (MDATPDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*mdatpdataConnectorState)(nil)).Elem()
}

type mdatpdataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes *AlertsDataTypeOfDataConnector `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'MicrosoftDefenderAdvancedThreatProtection'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a MDATPDataConnector resource.
type MDATPDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes AlertsDataTypeOfDataConnectorPtrInput
	// The kind of the data connector
	// Expected value is 'MicrosoftDefenderAdvancedThreatProtection'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (MDATPDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mdatpdataConnectorArgs)(nil)).Elem()
}

type MDATPDataConnectorInput interface {
	pulumi.Input

	ToMDATPDataConnectorOutput() MDATPDataConnectorOutput
	ToMDATPDataConnectorOutputWithContext(ctx context.Context) MDATPDataConnectorOutput
}

func (*MDATPDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**MDATPDataConnector)(nil)).Elem()
}

func (i *MDATPDataConnector) ToMDATPDataConnectorOutput() MDATPDataConnectorOutput {
	return i.ToMDATPDataConnectorOutputWithContext(context.Background())
}

func (i *MDATPDataConnector) ToMDATPDataConnectorOutputWithContext(ctx context.Context) MDATPDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MDATPDataConnectorOutput)
}

type MDATPDataConnectorOutput struct{ *pulumi.OutputState }

func (MDATPDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MDATPDataConnector)(nil)).Elem()
}

func (o MDATPDataConnectorOutput) ToMDATPDataConnectorOutput() MDATPDataConnectorOutput {
	return o
}

func (o MDATPDataConnectorOutput) ToMDATPDataConnectorOutputWithContext(ctx context.Context) MDATPDataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o MDATPDataConnectorOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v *MDATPDataConnector) AlertsDataTypeOfDataConnectorResponsePtrOutput { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

// Etag of the azure resource
func (o MDATPDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MDATPDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'MicrosoftDefenderAdvancedThreatProtection'.
func (o MDATPDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MDATPDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o MDATPDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MDATPDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MDATPDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MDATPDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o MDATPDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *MDATPDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MDATPDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MDATPDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MDATPDataConnectorOutput{})
}
