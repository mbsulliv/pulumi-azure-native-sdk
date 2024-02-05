// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents MCAS (Microsoft Cloud App Security) data connector.
type MCASDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes MCASDataConnectorDataTypesResponsePtrOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'MicrosoftCloudAppSecurity'.
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

// NewMCASDataConnector registers a new resource with the given unique name, arguments, and options.
func NewMCASDataConnector(ctx *pulumi.Context,
	name string, args *MCASDataConnectorArgs, opts ...pulumi.ResourceOption) (*MCASDataConnector, error) {
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
	args.Kind = pulumi.String("MicrosoftCloudAppSecurity")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:MCASDataConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MCASDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20230201:MCASDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMCASDataConnector gets an existing MCASDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMCASDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MCASDataConnectorState, opts ...pulumi.ResourceOption) (*MCASDataConnector, error) {
	var resource MCASDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20230201:MCASDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MCASDataConnector resources.
type mcasdataConnectorState struct {
}

type MCASDataConnectorState struct {
}

func (MCASDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*mcasdataConnectorState)(nil)).Elem()
}

type mcasdataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes *MCASDataConnectorDataTypes `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'MicrosoftCloudAppSecurity'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId *string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a MCASDataConnector resource.
type MCASDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes MCASDataConnectorDataTypesPtrInput
	// The kind of the data connector
	// Expected value is 'MicrosoftCloudAppSecurity'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (MCASDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mcasdataConnectorArgs)(nil)).Elem()
}

type MCASDataConnectorInput interface {
	pulumi.Input

	ToMCASDataConnectorOutput() MCASDataConnectorOutput
	ToMCASDataConnectorOutputWithContext(ctx context.Context) MCASDataConnectorOutput
}

func (*MCASDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**MCASDataConnector)(nil)).Elem()
}

func (i *MCASDataConnector) ToMCASDataConnectorOutput() MCASDataConnectorOutput {
	return i.ToMCASDataConnectorOutputWithContext(context.Background())
}

func (i *MCASDataConnector) ToMCASDataConnectorOutputWithContext(ctx context.Context) MCASDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorOutput)
}

type MCASDataConnectorOutput struct{ *pulumi.OutputState }

func (MCASDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MCASDataConnector)(nil)).Elem()
}

func (o MCASDataConnectorOutput) ToMCASDataConnectorOutput() MCASDataConnectorOutput {
	return o
}

func (o MCASDataConnectorOutput) ToMCASDataConnectorOutputWithContext(ctx context.Context) MCASDataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o MCASDataConnectorOutput) DataTypes() MCASDataConnectorDataTypesResponsePtrOutput {
	return o.ApplyT(func(v *MCASDataConnector) MCASDataConnectorDataTypesResponsePtrOutput { return v.DataTypes }).(MCASDataConnectorDataTypesResponsePtrOutput)
}

// Etag of the azure resource
func (o MCASDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MCASDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'MicrosoftCloudAppSecurity'.
func (o MCASDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MCASDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o MCASDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MCASDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MCASDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MCASDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o MCASDataConnectorOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MCASDataConnector) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MCASDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MCASDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MCASDataConnectorOutput{})
}
