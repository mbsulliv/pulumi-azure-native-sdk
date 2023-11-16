// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents AADIP (Azure Active Directory Identity Protection) data connector.
type AADDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'AzureActiveDirectory'.
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

// NewAADDataConnector registers a new resource with the given unique name, arguments, and options.
func NewAADDataConnector(ctx *pulumi.Context,
	name string, args *AADDataConnectorArgs, opts ...pulumi.ResourceOption) (*AADDataConnector, error) {
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
	args.Kind = pulumi.String("AzureActiveDirectory")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:AADDataConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AADDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20230601preview:AADDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAADDataConnector gets an existing AADDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAADDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AADDataConnectorState, opts ...pulumi.ResourceOption) (*AADDataConnector, error) {
	var resource AADDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20230601preview:AADDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AADDataConnector resources.
type aaddataConnectorState struct {
}

type AADDataConnectorState struct {
}

func (AADDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*aaddataConnectorState)(nil)).Elem()
}

type aaddataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes *AlertsDataTypeOfDataConnector `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'AzureActiveDirectory'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a AADDataConnector resource.
type AADDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes AlertsDataTypeOfDataConnectorPtrInput
	// The kind of the data connector
	// Expected value is 'AzureActiveDirectory'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (AADDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aaddataConnectorArgs)(nil)).Elem()
}

type AADDataConnectorInput interface {
	pulumi.Input

	ToAADDataConnectorOutput() AADDataConnectorOutput
	ToAADDataConnectorOutputWithContext(ctx context.Context) AADDataConnectorOutput
}

func (*AADDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**AADDataConnector)(nil)).Elem()
}

func (i *AADDataConnector) ToAADDataConnectorOutput() AADDataConnectorOutput {
	return i.ToAADDataConnectorOutputWithContext(context.Background())
}

func (i *AADDataConnector) ToAADDataConnectorOutputWithContext(ctx context.Context) AADDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADDataConnectorOutput)
}

type AADDataConnectorOutput struct{ *pulumi.OutputState }

func (AADDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AADDataConnector)(nil)).Elem()
}

func (o AADDataConnectorOutput) ToAADDataConnectorOutput() AADDataConnectorOutput {
	return o
}

func (o AADDataConnectorOutput) ToAADDataConnectorOutputWithContext(ctx context.Context) AADDataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o AADDataConnectorOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v *AADDataConnector) AlertsDataTypeOfDataConnectorResponsePtrOutput { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

// Etag of the azure resource
func (o AADDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'AzureActiveDirectory'.
func (o AADDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AADDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o AADDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AADDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AADDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AADDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o AADDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *AADDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AADDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AADDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AADDataConnectorOutput{})
}
