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

// Represents ASC (Azure Security Center) data connector.
type ASCDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'AzureSecurityCenter'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The subscription id to connect to, and get the data from.
	SubscriptionId pulumi.StringPtrOutput `pulumi:"subscriptionId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewASCDataConnector registers a new resource with the given unique name, arguments, and options.
func NewASCDataConnector(ctx *pulumi.Context,
	name string, args *ASCDataConnectorArgs, opts ...pulumi.ResourceOption) (*ASCDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("AzureSecurityCenter")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:ASCDataConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ASCDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20230601preview:ASCDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetASCDataConnector gets an existing ASCDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetASCDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ASCDataConnectorState, opts ...pulumi.ResourceOption) (*ASCDataConnector, error) {
	var resource ASCDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20230601preview:ASCDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ASCDataConnector resources.
type ascdataConnectorState struct {
}

type ASCDataConnectorState struct {
}

func (ASCDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*ascdataConnectorState)(nil)).Elem()
}

type ascdataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes *AlertsDataTypeOfDataConnector `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'AzureSecurityCenter'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The subscription id to connect to, and get the data from.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ASCDataConnector resource.
type ASCDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes AlertsDataTypeOfDataConnectorPtrInput
	// The kind of the data connector
	// Expected value is 'AzureSecurityCenter'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The subscription id to connect to, and get the data from.
	SubscriptionId pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (ASCDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ascdataConnectorArgs)(nil)).Elem()
}

type ASCDataConnectorInput interface {
	pulumi.Input

	ToASCDataConnectorOutput() ASCDataConnectorOutput
	ToASCDataConnectorOutputWithContext(ctx context.Context) ASCDataConnectorOutput
}

func (*ASCDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**ASCDataConnector)(nil)).Elem()
}

func (i *ASCDataConnector) ToASCDataConnectorOutput() ASCDataConnectorOutput {
	return i.ToASCDataConnectorOutputWithContext(context.Background())
}

func (i *ASCDataConnector) ToASCDataConnectorOutputWithContext(ctx context.Context) ASCDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ASCDataConnectorOutput)
}

type ASCDataConnectorOutput struct{ *pulumi.OutputState }

func (ASCDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ASCDataConnector)(nil)).Elem()
}

func (o ASCDataConnectorOutput) ToASCDataConnectorOutput() ASCDataConnectorOutput {
	return o
}

func (o ASCDataConnectorOutput) ToASCDataConnectorOutputWithContext(ctx context.Context) ASCDataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o ASCDataConnectorOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v *ASCDataConnector) AlertsDataTypeOfDataConnectorResponsePtrOutput { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

// Etag of the azure resource
func (o ASCDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ASCDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'AzureSecurityCenter'.
func (o ASCDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ASCDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o ASCDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ASCDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The subscription id to connect to, and get the data from.
func (o ASCDataConnectorOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ASCDataConnector) pulumi.StringPtrOutput { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ASCDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ASCDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ASCDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ASCDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ASCDataConnectorOutput{})
}
