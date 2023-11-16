// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents IoT data connector.
type IoTDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'IOT'.
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

// NewIoTDataConnector registers a new resource with the given unique name, arguments, and options.
func NewIoTDataConnector(ctx *pulumi.Context,
	name string, args *IoTDataConnectorArgs, opts ...pulumi.ResourceOption) (*IoTDataConnector, error) {
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
	args.Kind = pulumi.String("IOT")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:IoTDataConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IoTDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20231001preview:IoTDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIoTDataConnector gets an existing IoTDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIoTDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IoTDataConnectorState, opts ...pulumi.ResourceOption) (*IoTDataConnector, error) {
	var resource IoTDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20231001preview:IoTDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IoTDataConnector resources.
type ioTDataConnectorState struct {
}

type IoTDataConnectorState struct {
}

func (IoTDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTDataConnectorState)(nil)).Elem()
}

type ioTDataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes *AlertsDataTypeOfDataConnector `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'IOT'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The subscription id to connect to, and get the data from.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a IoTDataConnector resource.
type IoTDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes AlertsDataTypeOfDataConnectorPtrInput
	// The kind of the data connector
	// Expected value is 'IOT'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The subscription id to connect to, and get the data from.
	SubscriptionId pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (IoTDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTDataConnectorArgs)(nil)).Elem()
}

type IoTDataConnectorInput interface {
	pulumi.Input

	ToIoTDataConnectorOutput() IoTDataConnectorOutput
	ToIoTDataConnectorOutputWithContext(ctx context.Context) IoTDataConnectorOutput
}

func (*IoTDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTDataConnector)(nil)).Elem()
}

func (i *IoTDataConnector) ToIoTDataConnectorOutput() IoTDataConnectorOutput {
	return i.ToIoTDataConnectorOutputWithContext(context.Background())
}

func (i *IoTDataConnector) ToIoTDataConnectorOutputWithContext(ctx context.Context) IoTDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTDataConnectorOutput)
}

type IoTDataConnectorOutput struct{ *pulumi.OutputState }

func (IoTDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTDataConnector)(nil)).Elem()
}

func (o IoTDataConnectorOutput) ToIoTDataConnectorOutput() IoTDataConnectorOutput {
	return o
}

func (o IoTDataConnectorOutput) ToIoTDataConnectorOutputWithContext(ctx context.Context) IoTDataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o IoTDataConnectorOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v *IoTDataConnector) AlertsDataTypeOfDataConnectorResponsePtrOutput { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

// Etag of the azure resource
func (o IoTDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'IOT'.
func (o IoTDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o IoTDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The subscription id to connect to, and get the data from.
func (o IoTDataConnectorOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTDataConnector) pulumi.StringPtrOutput { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o IoTDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IoTDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o IoTDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IoTDataConnectorOutput{})
}
