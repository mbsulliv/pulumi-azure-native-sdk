// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents OfficeATP (Office 365 Advanced Threat Protection) data connector.
type OfficeATPDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'OfficeATP'.
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

// NewOfficeATPDataConnector registers a new resource with the given unique name, arguments, and options.
func NewOfficeATPDataConnector(ctx *pulumi.Context,
	name string, args *OfficeATPDataConnectorArgs, opts ...pulumi.ResourceOption) (*OfficeATPDataConnector, error) {
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
	args.Kind = pulumi.String("OfficeATP")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:OfficeATPDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource OfficeATPDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220401preview:OfficeATPDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOfficeATPDataConnector gets an existing OfficeATPDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOfficeATPDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OfficeATPDataConnectorState, opts ...pulumi.ResourceOption) (*OfficeATPDataConnector, error) {
	var resource OfficeATPDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220401preview:OfficeATPDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OfficeATPDataConnector resources.
type officeATPDataConnectorState struct {
}

type OfficeATPDataConnectorState struct {
}

func (OfficeATPDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*officeATPDataConnectorState)(nil)).Elem()
}

type officeATPDataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes *AlertsDataTypeOfDataConnector `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'OfficeATP'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a OfficeATPDataConnector resource.
type OfficeATPDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes AlertsDataTypeOfDataConnectorPtrInput
	// The kind of the data connector
	// Expected value is 'OfficeATP'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (OfficeATPDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*officeATPDataConnectorArgs)(nil)).Elem()
}

type OfficeATPDataConnectorInput interface {
	pulumi.Input

	ToOfficeATPDataConnectorOutput() OfficeATPDataConnectorOutput
	ToOfficeATPDataConnectorOutputWithContext(ctx context.Context) OfficeATPDataConnectorOutput
}

func (*OfficeATPDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeATPDataConnector)(nil)).Elem()
}

func (i *OfficeATPDataConnector) ToOfficeATPDataConnectorOutput() OfficeATPDataConnectorOutput {
	return i.ToOfficeATPDataConnectorOutputWithContext(context.Background())
}

func (i *OfficeATPDataConnector) ToOfficeATPDataConnectorOutputWithContext(ctx context.Context) OfficeATPDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeATPDataConnectorOutput)
}

type OfficeATPDataConnectorOutput struct{ *pulumi.OutputState }

func (OfficeATPDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeATPDataConnector)(nil)).Elem()
}

func (o OfficeATPDataConnectorOutput) ToOfficeATPDataConnectorOutput() OfficeATPDataConnectorOutput {
	return o
}

func (o OfficeATPDataConnectorOutput) ToOfficeATPDataConnectorOutputWithContext(ctx context.Context) OfficeATPDataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o OfficeATPDataConnectorOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v *OfficeATPDataConnector) AlertsDataTypeOfDataConnectorResponsePtrOutput { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

// Etag of the azure resource
func (o OfficeATPDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficeATPDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'OfficeATP'.
func (o OfficeATPDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeATPDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o OfficeATPDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeATPDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o OfficeATPDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OfficeATPDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o OfficeATPDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeATPDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o OfficeATPDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeATPDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OfficeATPDataConnectorOutput{})
}
