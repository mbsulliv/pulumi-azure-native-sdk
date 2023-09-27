// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents office data connector.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-01-01
type OfficeDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes OfficeDataConnectorDataTypesResponsePtrOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'Office365'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOfficeDataConnector registers a new resource with the given unique name, arguments, and options.
func NewOfficeDataConnector(ctx *pulumi.Context,
	name string, args *OfficeDataConnectorArgs, opts ...pulumi.ResourceOption) (*OfficeDataConnector, error) {
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
	args.Kind = pulumi.String("Office365")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:OfficeDataConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource OfficeDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights:OfficeDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOfficeDataConnector gets an existing OfficeDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOfficeDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OfficeDataConnectorState, opts ...pulumi.ResourceOption) (*OfficeDataConnector, error) {
	var resource OfficeDataConnector
	err := ctx.ReadResource("azure-native:securityinsights:OfficeDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OfficeDataConnector resources.
type officeDataConnectorState struct {
}

type OfficeDataConnectorState struct {
}

func (OfficeDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*officeDataConnectorState)(nil)).Elem()
}

type officeDataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes *OfficeDataConnectorDataTypes `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'Office365'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId *string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a OfficeDataConnector resource.
type OfficeDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes OfficeDataConnectorDataTypesPtrInput
	// The kind of the data connector
	// Expected value is 'Office365'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (OfficeDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*officeDataConnectorArgs)(nil)).Elem()
}

type OfficeDataConnectorInput interface {
	pulumi.Input

	ToOfficeDataConnectorOutput() OfficeDataConnectorOutput
	ToOfficeDataConnectorOutputWithContext(ctx context.Context) OfficeDataConnectorOutput
}

func (*OfficeDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnector)(nil)).Elem()
}

func (i *OfficeDataConnector) ToOfficeDataConnectorOutput() OfficeDataConnectorOutput {
	return i.ToOfficeDataConnectorOutputWithContext(context.Background())
}

func (i *OfficeDataConnector) ToOfficeDataConnectorOutputWithContext(ctx context.Context) OfficeDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorOutput)
}

func (i *OfficeDataConnector) ToOutput(ctx context.Context) pulumix.Output[*OfficeDataConnector] {
	return pulumix.Output[*OfficeDataConnector]{
		OutputState: i.ToOfficeDataConnectorOutputWithContext(ctx).OutputState,
	}
}

type OfficeDataConnectorOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnector)(nil)).Elem()
}

func (o OfficeDataConnectorOutput) ToOfficeDataConnectorOutput() OfficeDataConnectorOutput {
	return o
}

func (o OfficeDataConnectorOutput) ToOfficeDataConnectorOutputWithContext(ctx context.Context) OfficeDataConnectorOutput {
	return o
}

func (o OfficeDataConnectorOutput) ToOutput(ctx context.Context) pulumix.Output[*OfficeDataConnector] {
	return pulumix.Output[*OfficeDataConnector]{
		OutputState: o.OutputState,
	}
}

// The available data types for the connector.
func (o OfficeDataConnectorOutput) DataTypes() OfficeDataConnectorDataTypesResponsePtrOutput {
	return o.ApplyT(func(v *OfficeDataConnector) OfficeDataConnectorDataTypesResponsePtrOutput { return v.DataTypes }).(OfficeDataConnectorDataTypesResponsePtrOutput)
}

// Etag of the azure resource
func (o OfficeDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'Office365'.
func (o OfficeDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o OfficeDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o OfficeDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OfficeDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o OfficeDataConnectorOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnector) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o OfficeDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OfficeDataConnectorOutput{})
}
