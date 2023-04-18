// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Codeless UI data connector.
type CodelessUiDataConnector struct {
	pulumi.CustomResourceState

	// Config to describe the instructions blade
	ConnectorUiConfig CodelessUiConnectorConfigPropertiesResponsePtrOutput `pulumi:"connectorUiConfig"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'GenericUI'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCodelessUiDataConnector registers a new resource with the given unique name, arguments, and options.
func NewCodelessUiDataConnector(ctx *pulumi.Context,
	name string, args *CodelessUiDataConnectorArgs, opts ...pulumi.ResourceOption) (*CodelessUiDataConnector, error) {
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
	args.Kind = pulumi.String("GenericUI")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:CodelessUiDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource CodelessUiDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220601preview:CodelessUiDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCodelessUiDataConnector gets an existing CodelessUiDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCodelessUiDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CodelessUiDataConnectorState, opts ...pulumi.ResourceOption) (*CodelessUiDataConnector, error) {
	var resource CodelessUiDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220601preview:CodelessUiDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CodelessUiDataConnector resources.
type codelessUiDataConnectorState struct {
}

type CodelessUiDataConnectorState struct {
}

func (CodelessUiDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*codelessUiDataConnectorState)(nil)).Elem()
}

type codelessUiDataConnectorArgs struct {
	// Config to describe the instructions blade
	ConnectorUiConfig *CodelessUiConnectorConfigProperties `pulumi:"connectorUiConfig"`
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The kind of the data connector
	// Expected value is 'GenericUI'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a CodelessUiDataConnector resource.
type CodelessUiDataConnectorArgs struct {
	// Config to describe the instructions blade
	ConnectorUiConfig CodelessUiConnectorConfigPropertiesPtrInput
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The kind of the data connector
	// Expected value is 'GenericUI'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (CodelessUiDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*codelessUiDataConnectorArgs)(nil)).Elem()
}

type CodelessUiDataConnectorInput interface {
	pulumi.Input

	ToCodelessUiDataConnectorOutput() CodelessUiDataConnectorOutput
	ToCodelessUiDataConnectorOutputWithContext(ctx context.Context) CodelessUiDataConnectorOutput
}

func (*CodelessUiDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessUiDataConnector)(nil)).Elem()
}

func (i *CodelessUiDataConnector) ToCodelessUiDataConnectorOutput() CodelessUiDataConnectorOutput {
	return i.ToCodelessUiDataConnectorOutputWithContext(context.Background())
}

func (i *CodelessUiDataConnector) ToCodelessUiDataConnectorOutputWithContext(ctx context.Context) CodelessUiDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiDataConnectorOutput)
}

type CodelessUiDataConnectorOutput struct{ *pulumi.OutputState }

func (CodelessUiDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessUiDataConnector)(nil)).Elem()
}

func (o CodelessUiDataConnectorOutput) ToCodelessUiDataConnectorOutput() CodelessUiDataConnectorOutput {
	return o
}

func (o CodelessUiDataConnectorOutput) ToCodelessUiDataConnectorOutputWithContext(ctx context.Context) CodelessUiDataConnectorOutput {
	return o
}

// Config to describe the instructions blade
func (o CodelessUiDataConnectorOutput) ConnectorUiConfig() CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CodelessUiDataConnector) CodelessUiConnectorConfigPropertiesResponsePtrOutput {
		return v.ConnectorUiConfig
	}).(CodelessUiConnectorConfigPropertiesResponsePtrOutput)
}

// Etag of the azure resource
func (o CodelessUiDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessUiDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'GenericUI'.
func (o CodelessUiDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CodelessUiDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o CodelessUiDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CodelessUiDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o CodelessUiDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CodelessUiDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CodelessUiDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CodelessUiDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CodelessUiDataConnectorOutput{})
}
