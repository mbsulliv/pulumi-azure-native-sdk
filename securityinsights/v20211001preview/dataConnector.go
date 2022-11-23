// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data connector
//
// Deprecated: Please use one of the variants: AADDataConnector, AATPDataConnector, ASCDataConnector, AwsCloudTrailDataConnector, AwsS3DataConnector, CodelessApiPollingDataConnector, CodelessUiDataConnector, Dynamics365DataConnector, MCASDataConnector, MDATPDataConnector, MSTIDataConnector, MTPDataConnector, Office365ProjectDataConnector, OfficeATPDataConnector, OfficeDataConnector, OfficeIRMDataConnector, OfficePowerBIDataConnector, TIDataConnector, TiTaxiiDataConnector.
type DataConnector struct {
	pulumi.CustomResourceState

	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The data connector kind
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataConnector registers a new resource with the given unique name, arguments, and options.
func NewDataConnector(ctx *pulumi.Context,
	name string, args *DataConnectorArgs, opts ...pulumi.ResourceOption) (*DataConnector, error) {
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:DataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource DataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20211001preview:DataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataConnector gets an existing DataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataConnectorState, opts ...pulumi.ResourceOption) (*DataConnector, error) {
	var resource DataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20211001preview:DataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataConnector resources.
type dataConnectorState struct {
}

type DataConnectorState struct {
}

func (DataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataConnectorState)(nil)).Elem()
}

type dataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The data connector kind
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a DataConnector resource.
type DataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The data connector kind
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (DataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataConnectorArgs)(nil)).Elem()
}

type DataConnectorInput interface {
	pulumi.Input

	ToDataConnectorOutput() DataConnectorOutput
	ToDataConnectorOutputWithContext(ctx context.Context) DataConnectorOutput
}

func (*DataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnector)(nil)).Elem()
}

func (i *DataConnector) ToDataConnectorOutput() DataConnectorOutput {
	return i.ToDataConnectorOutputWithContext(context.Background())
}

func (i *DataConnector) ToDataConnectorOutputWithContext(ctx context.Context) DataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorOutput)
}

type DataConnectorOutput struct{ *pulumi.OutputState }

func (DataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnector)(nil)).Elem()
}

func (o DataConnectorOutput) ToDataConnectorOutput() DataConnectorOutput {
	return o
}

func (o DataConnectorOutput) ToDataConnectorOutputWithContext(ctx context.Context) DataConnectorOutput {
	return o
}

// Etag of the azure resource
func (o DataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The data connector kind
func (o DataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o DataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataConnectorOutput{})
}
