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

// Represents Dynamics365 data connector.
type Dynamics365DataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes Dynamics365DataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'Dynamics365'.
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

// NewDynamics365DataConnector registers a new resource with the given unique name, arguments, and options.
func NewDynamics365DataConnector(ctx *pulumi.Context,
	name string, args *Dynamics365DataConnectorArgs, opts ...pulumi.ResourceOption) (*Dynamics365DataConnector, error) {
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
	args.Kind = pulumi.String("Dynamics365")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:Dynamics365DataConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Dynamics365DataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20230601preview:Dynamics365DataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDynamics365DataConnector gets an existing Dynamics365DataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDynamics365DataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Dynamics365DataConnectorState, opts ...pulumi.ResourceOption) (*Dynamics365DataConnector, error) {
	var resource Dynamics365DataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20230601preview:Dynamics365DataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dynamics365DataConnector resources.
type dynamics365DataConnectorState struct {
}

type Dynamics365DataConnectorState struct {
}

func (Dynamics365DataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*dynamics365DataConnectorState)(nil)).Elem()
}

type dynamics365DataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes Dynamics365DataConnectorDataTypes `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'Dynamics365'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Dynamics365DataConnector resource.
type Dynamics365DataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes Dynamics365DataConnectorDataTypesInput
	// The kind of the data connector
	// Expected value is 'Dynamics365'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (Dynamics365DataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dynamics365DataConnectorArgs)(nil)).Elem()
}

type Dynamics365DataConnectorInput interface {
	pulumi.Input

	ToDynamics365DataConnectorOutput() Dynamics365DataConnectorOutput
	ToDynamics365DataConnectorOutputWithContext(ctx context.Context) Dynamics365DataConnectorOutput
}

func (*Dynamics365DataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**Dynamics365DataConnector)(nil)).Elem()
}

func (i *Dynamics365DataConnector) ToDynamics365DataConnectorOutput() Dynamics365DataConnectorOutput {
	return i.ToDynamics365DataConnectorOutputWithContext(context.Background())
}

func (i *Dynamics365DataConnector) ToDynamics365DataConnectorOutputWithContext(ctx context.Context) Dynamics365DataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Dynamics365DataConnectorOutput)
}

type Dynamics365DataConnectorOutput struct{ *pulumi.OutputState }

func (Dynamics365DataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dynamics365DataConnector)(nil)).Elem()
}

func (o Dynamics365DataConnectorOutput) ToDynamics365DataConnectorOutput() Dynamics365DataConnectorOutput {
	return o
}

func (o Dynamics365DataConnectorOutput) ToDynamics365DataConnectorOutputWithContext(ctx context.Context) Dynamics365DataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o Dynamics365DataConnectorOutput) DataTypes() Dynamics365DataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *Dynamics365DataConnector) Dynamics365DataConnectorDataTypesResponseOutput { return v.DataTypes }).(Dynamics365DataConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o Dynamics365DataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dynamics365DataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'Dynamics365'.
func (o Dynamics365DataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Dynamics365DataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o Dynamics365DataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dynamics365DataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o Dynamics365DataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Dynamics365DataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o Dynamics365DataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Dynamics365DataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o Dynamics365DataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Dynamics365DataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Dynamics365DataConnectorOutput{})
}
