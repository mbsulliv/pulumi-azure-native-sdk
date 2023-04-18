// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents MTP (Microsoft Threat Protection) data connector.
type MTPDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes MTPDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'MicrosoftThreatProtection'.
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

// NewMTPDataConnector registers a new resource with the given unique name, arguments, and options.
func NewMTPDataConnector(ctx *pulumi.Context,
	name string, args *MTPDataConnectorArgs, opts ...pulumi.ResourceOption) (*MTPDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataTypes == nil {
		return nil, errors.New("invalid value for required argument 'DataTypes'")
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
	args.Kind = pulumi.String("MicrosoftThreatProtection")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:MTPDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource MTPDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220501preview:MTPDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMTPDataConnector gets an existing MTPDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMTPDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MTPDataConnectorState, opts ...pulumi.ResourceOption) (*MTPDataConnector, error) {
	var resource MTPDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220501preview:MTPDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MTPDataConnector resources.
type mtpdataConnectorState struct {
}

type MTPDataConnectorState struct {
}

func (MTPDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*mtpdataConnectorState)(nil)).Elem()
}

type mtpdataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes MTPDataConnectorDataTypes `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'MicrosoftThreatProtection'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a MTPDataConnector resource.
type MTPDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes MTPDataConnectorDataTypesInput
	// The kind of the data connector
	// Expected value is 'MicrosoftThreatProtection'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (MTPDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mtpdataConnectorArgs)(nil)).Elem()
}

type MTPDataConnectorInput interface {
	pulumi.Input

	ToMTPDataConnectorOutput() MTPDataConnectorOutput
	ToMTPDataConnectorOutputWithContext(ctx context.Context) MTPDataConnectorOutput
}

func (*MTPDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**MTPDataConnector)(nil)).Elem()
}

func (i *MTPDataConnector) ToMTPDataConnectorOutput() MTPDataConnectorOutput {
	return i.ToMTPDataConnectorOutputWithContext(context.Background())
}

func (i *MTPDataConnector) ToMTPDataConnectorOutputWithContext(ctx context.Context) MTPDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorOutput)
}

type MTPDataConnectorOutput struct{ *pulumi.OutputState }

func (MTPDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MTPDataConnector)(nil)).Elem()
}

func (o MTPDataConnectorOutput) ToMTPDataConnectorOutput() MTPDataConnectorOutput {
	return o
}

func (o MTPDataConnectorOutput) ToMTPDataConnectorOutputWithContext(ctx context.Context) MTPDataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o MTPDataConnectorOutput) DataTypes() MTPDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *MTPDataConnector) MTPDataConnectorDataTypesResponseOutput { return v.DataTypes }).(MTPDataConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o MTPDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MTPDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'MicrosoftThreatProtection'.
func (o MTPDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MTPDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o MTPDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MTPDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MTPDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MTPDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o MTPDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *MTPDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MTPDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MTPDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MTPDataConnectorOutput{})
}
