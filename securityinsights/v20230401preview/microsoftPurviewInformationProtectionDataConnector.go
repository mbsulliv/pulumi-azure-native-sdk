// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Microsoft Purview Information Protection data connector.
type MicrosoftPurviewInformationProtectionDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes MicrosoftPurviewInformationProtectionConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'MicrosoftPurviewInformationProtection'.
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

// NewMicrosoftPurviewInformationProtectionDataConnector registers a new resource with the given unique name, arguments, and options.
func NewMicrosoftPurviewInformationProtectionDataConnector(ctx *pulumi.Context,
	name string, args *MicrosoftPurviewInformationProtectionDataConnectorArgs, opts ...pulumi.ResourceOption) (*MicrosoftPurviewInformationProtectionDataConnector, error) {
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
	args.Kind = pulumi.String("MicrosoftPurviewInformationProtection")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:MicrosoftPurviewInformationProtectionDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:MicrosoftPurviewInformationProtectionDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource MicrosoftPurviewInformationProtectionDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20230401preview:MicrosoftPurviewInformationProtectionDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMicrosoftPurviewInformationProtectionDataConnector gets an existing MicrosoftPurviewInformationProtectionDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMicrosoftPurviewInformationProtectionDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MicrosoftPurviewInformationProtectionDataConnectorState, opts ...pulumi.ResourceOption) (*MicrosoftPurviewInformationProtectionDataConnector, error) {
	var resource MicrosoftPurviewInformationProtectionDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20230401preview:MicrosoftPurviewInformationProtectionDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MicrosoftPurviewInformationProtectionDataConnector resources.
type microsoftPurviewInformationProtectionDataConnectorState struct {
}

type MicrosoftPurviewInformationProtectionDataConnectorState struct {
}

func (MicrosoftPurviewInformationProtectionDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*microsoftPurviewInformationProtectionDataConnectorState)(nil)).Elem()
}

type microsoftPurviewInformationProtectionDataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes MicrosoftPurviewInformationProtectionConnectorDataTypes `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'MicrosoftPurviewInformationProtection'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a MicrosoftPurviewInformationProtectionDataConnector resource.
type MicrosoftPurviewInformationProtectionDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes MicrosoftPurviewInformationProtectionConnectorDataTypesInput
	// The kind of the data connector
	// Expected value is 'MicrosoftPurviewInformationProtection'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (MicrosoftPurviewInformationProtectionDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*microsoftPurviewInformationProtectionDataConnectorArgs)(nil)).Elem()
}

type MicrosoftPurviewInformationProtectionDataConnectorInput interface {
	pulumi.Input

	ToMicrosoftPurviewInformationProtectionDataConnectorOutput() MicrosoftPurviewInformationProtectionDataConnectorOutput
	ToMicrosoftPurviewInformationProtectionDataConnectorOutputWithContext(ctx context.Context) MicrosoftPurviewInformationProtectionDataConnectorOutput
}

func (*MicrosoftPurviewInformationProtectionDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**MicrosoftPurviewInformationProtectionDataConnector)(nil)).Elem()
}

func (i *MicrosoftPurviewInformationProtectionDataConnector) ToMicrosoftPurviewInformationProtectionDataConnectorOutput() MicrosoftPurviewInformationProtectionDataConnectorOutput {
	return i.ToMicrosoftPurviewInformationProtectionDataConnectorOutputWithContext(context.Background())
}

func (i *MicrosoftPurviewInformationProtectionDataConnector) ToMicrosoftPurviewInformationProtectionDataConnectorOutputWithContext(ctx context.Context) MicrosoftPurviewInformationProtectionDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MicrosoftPurviewInformationProtectionDataConnectorOutput)
}

type MicrosoftPurviewInformationProtectionDataConnectorOutput struct{ *pulumi.OutputState }

func (MicrosoftPurviewInformationProtectionDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MicrosoftPurviewInformationProtectionDataConnector)(nil)).Elem()
}

func (o MicrosoftPurviewInformationProtectionDataConnectorOutput) ToMicrosoftPurviewInformationProtectionDataConnectorOutput() MicrosoftPurviewInformationProtectionDataConnectorOutput {
	return o
}

func (o MicrosoftPurviewInformationProtectionDataConnectorOutput) ToMicrosoftPurviewInformationProtectionDataConnectorOutputWithContext(ctx context.Context) MicrosoftPurviewInformationProtectionDataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o MicrosoftPurviewInformationProtectionDataConnectorOutput) DataTypes() MicrosoftPurviewInformationProtectionConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *MicrosoftPurviewInformationProtectionDataConnector) MicrosoftPurviewInformationProtectionConnectorDataTypesResponseOutput {
		return v.DataTypes
	}).(MicrosoftPurviewInformationProtectionConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o MicrosoftPurviewInformationProtectionDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MicrosoftPurviewInformationProtectionDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'MicrosoftPurviewInformationProtection'.
func (o MicrosoftPurviewInformationProtectionDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftPurviewInformationProtectionDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o MicrosoftPurviewInformationProtectionDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftPurviewInformationProtectionDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MicrosoftPurviewInformationProtectionDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MicrosoftPurviewInformationProtectionDataConnector) SystemDataResponseOutput {
		return v.SystemData
	}).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o MicrosoftPurviewInformationProtectionDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftPurviewInformationProtectionDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MicrosoftPurviewInformationProtectionDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftPurviewInformationProtectionDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MicrosoftPurviewInformationProtectionDataConnectorOutput{})
}