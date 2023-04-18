// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Codeless API Polling data connector.
type CodelessApiPollingDataConnector struct {
	pulumi.CustomResourceState

	// Config to describe the instructions blade
	ConnectorUiConfig CodelessUiConnectorConfigPropertiesResponsePtrOutput `pulumi:"connectorUiConfig"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'APIPolling'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Config to describe the polling instructions
	PollingConfig CodelessConnectorPollingConfigPropertiesResponsePtrOutput `pulumi:"pollingConfig"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCodelessApiPollingDataConnector registers a new resource with the given unique name, arguments, and options.
func NewCodelessApiPollingDataConnector(ctx *pulumi.Context,
	name string, args *CodelessApiPollingDataConnectorArgs, opts ...pulumi.ResourceOption) (*CodelessApiPollingDataConnector, error) {
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
	args.Kind = pulumi.String("APIPolling")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:CodelessApiPollingDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource CodelessApiPollingDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220901preview:CodelessApiPollingDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCodelessApiPollingDataConnector gets an existing CodelessApiPollingDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCodelessApiPollingDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CodelessApiPollingDataConnectorState, opts ...pulumi.ResourceOption) (*CodelessApiPollingDataConnector, error) {
	var resource CodelessApiPollingDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220901preview:CodelessApiPollingDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CodelessApiPollingDataConnector resources.
type codelessApiPollingDataConnectorState struct {
}

type CodelessApiPollingDataConnectorState struct {
}

func (CodelessApiPollingDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*codelessApiPollingDataConnectorState)(nil)).Elem()
}

type codelessApiPollingDataConnectorArgs struct {
	// Config to describe the instructions blade
	ConnectorUiConfig *CodelessUiConnectorConfigProperties `pulumi:"connectorUiConfig"`
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The kind of the data connector
	// Expected value is 'APIPolling'.
	Kind string `pulumi:"kind"`
	// Config to describe the polling instructions
	PollingConfig *CodelessConnectorPollingConfigProperties `pulumi:"pollingConfig"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a CodelessApiPollingDataConnector resource.
type CodelessApiPollingDataConnectorArgs struct {
	// Config to describe the instructions blade
	ConnectorUiConfig CodelessUiConnectorConfigPropertiesPtrInput
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The kind of the data connector
	// Expected value is 'APIPolling'.
	Kind pulumi.StringInput
	// Config to describe the polling instructions
	PollingConfig CodelessConnectorPollingConfigPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (CodelessApiPollingDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*codelessApiPollingDataConnectorArgs)(nil)).Elem()
}

type CodelessApiPollingDataConnectorInput interface {
	pulumi.Input

	ToCodelessApiPollingDataConnectorOutput() CodelessApiPollingDataConnectorOutput
	ToCodelessApiPollingDataConnectorOutputWithContext(ctx context.Context) CodelessApiPollingDataConnectorOutput
}

func (*CodelessApiPollingDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessApiPollingDataConnector)(nil)).Elem()
}

func (i *CodelessApiPollingDataConnector) ToCodelessApiPollingDataConnectorOutput() CodelessApiPollingDataConnectorOutput {
	return i.ToCodelessApiPollingDataConnectorOutputWithContext(context.Background())
}

func (i *CodelessApiPollingDataConnector) ToCodelessApiPollingDataConnectorOutputWithContext(ctx context.Context) CodelessApiPollingDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessApiPollingDataConnectorOutput)
}

type CodelessApiPollingDataConnectorOutput struct{ *pulumi.OutputState }

func (CodelessApiPollingDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessApiPollingDataConnector)(nil)).Elem()
}

func (o CodelessApiPollingDataConnectorOutput) ToCodelessApiPollingDataConnectorOutput() CodelessApiPollingDataConnectorOutput {
	return o
}

func (o CodelessApiPollingDataConnectorOutput) ToCodelessApiPollingDataConnectorOutputWithContext(ctx context.Context) CodelessApiPollingDataConnectorOutput {
	return o
}

// Config to describe the instructions blade
func (o CodelessApiPollingDataConnectorOutput) ConnectorUiConfig() CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CodelessApiPollingDataConnector) CodelessUiConnectorConfigPropertiesResponsePtrOutput {
		return v.ConnectorUiConfig
	}).(CodelessUiConnectorConfigPropertiesResponsePtrOutput)
}

// Etag of the azure resource
func (o CodelessApiPollingDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessApiPollingDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'APIPolling'.
func (o CodelessApiPollingDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CodelessApiPollingDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o CodelessApiPollingDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CodelessApiPollingDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Config to describe the polling instructions
func (o CodelessApiPollingDataConnectorOutput) PollingConfig() CodelessConnectorPollingConfigPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CodelessApiPollingDataConnector) CodelessConnectorPollingConfigPropertiesResponsePtrOutput {
		return v.PollingConfig
	}).(CodelessConnectorPollingConfigPropertiesResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o CodelessApiPollingDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CodelessApiPollingDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CodelessApiPollingDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CodelessApiPollingDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CodelessApiPollingDataConnectorOutput{})
}
