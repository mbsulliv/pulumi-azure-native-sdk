// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Amazon Web Services S3 data connector.
type AwsS3DataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes AwsS3DataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	// The logs destination table name in LogAnalytics.
	DestinationTable pulumi.StringOutput `pulumi:"destinationTable"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'AmazonWebServicesS3'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Aws Role Arn that is used to access the Aws account.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The AWS sqs urls for the connector.
	SqsUrls pulumi.StringArrayOutput `pulumi:"sqsUrls"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAwsS3DataConnector registers a new resource with the given unique name, arguments, and options.
func NewAwsS3DataConnector(ctx *pulumi.Context,
	name string, args *AwsS3DataConnectorArgs, opts ...pulumi.ResourceOption) (*AwsS3DataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataTypes == nil {
		return nil, errors.New("invalid value for required argument 'DataTypes'")
	}
	if args.DestinationTable == nil {
		return nil, errors.New("invalid value for required argument 'DestinationTable'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.SqsUrls == nil {
		return nil, errors.New("invalid value for required argument 'SqsUrls'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("AmazonWebServicesS3")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:AwsS3DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:AwsS3DataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource AwsS3DataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20221001preview:AwsS3DataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAwsS3DataConnector gets an existing AwsS3DataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAwsS3DataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsS3DataConnectorState, opts ...pulumi.ResourceOption) (*AwsS3DataConnector, error) {
	var resource AwsS3DataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20221001preview:AwsS3DataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AwsS3DataConnector resources.
type awsS3DataConnectorState struct {
}

type AwsS3DataConnectorState struct {
}

func (AwsS3DataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsS3DataConnectorState)(nil)).Elem()
}

type awsS3DataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes AwsS3DataConnectorDataTypes `pulumi:"dataTypes"`
	// The logs destination table name in LogAnalytics.
	DestinationTable string `pulumi:"destinationTable"`
	// The kind of the data connector
	// Expected value is 'AmazonWebServicesS3'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Aws Role Arn that is used to access the Aws account.
	RoleArn string `pulumi:"roleArn"`
	// The AWS sqs urls for the connector.
	SqsUrls []string `pulumi:"sqsUrls"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a AwsS3DataConnector resource.
type AwsS3DataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes AwsS3DataConnectorDataTypesInput
	// The logs destination table name in LogAnalytics.
	DestinationTable pulumi.StringInput
	// The kind of the data connector
	// Expected value is 'AmazonWebServicesS3'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The Aws Role Arn that is used to access the Aws account.
	RoleArn pulumi.StringInput
	// The AWS sqs urls for the connector.
	SqsUrls pulumi.StringArrayInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (AwsS3DataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsS3DataConnectorArgs)(nil)).Elem()
}

type AwsS3DataConnectorInput interface {
	pulumi.Input

	ToAwsS3DataConnectorOutput() AwsS3DataConnectorOutput
	ToAwsS3DataConnectorOutputWithContext(ctx context.Context) AwsS3DataConnectorOutput
}

func (*AwsS3DataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsS3DataConnector)(nil)).Elem()
}

func (i *AwsS3DataConnector) ToAwsS3DataConnectorOutput() AwsS3DataConnectorOutput {
	return i.ToAwsS3DataConnectorOutputWithContext(context.Background())
}

func (i *AwsS3DataConnector) ToAwsS3DataConnectorOutputWithContext(ctx context.Context) AwsS3DataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsS3DataConnectorOutput)
}

type AwsS3DataConnectorOutput struct{ *pulumi.OutputState }

func (AwsS3DataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsS3DataConnector)(nil)).Elem()
}

func (o AwsS3DataConnectorOutput) ToAwsS3DataConnectorOutput() AwsS3DataConnectorOutput {
	return o
}

func (o AwsS3DataConnectorOutput) ToAwsS3DataConnectorOutputWithContext(ctx context.Context) AwsS3DataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o AwsS3DataConnectorOutput) DataTypes() AwsS3DataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) AwsS3DataConnectorDataTypesResponseOutput { return v.DataTypes }).(AwsS3DataConnectorDataTypesResponseOutput)
}

// The logs destination table name in LogAnalytics.
func (o AwsS3DataConnectorOutput) DestinationTable() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) pulumi.StringOutput { return v.DestinationTable }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o AwsS3DataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'AmazonWebServicesS3'.
func (o AwsS3DataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o AwsS3DataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Aws Role Arn that is used to access the Aws account.
func (o AwsS3DataConnectorOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// The AWS sqs urls for the connector.
func (o AwsS3DataConnectorOutput) SqsUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) pulumi.StringArrayOutput { return v.SqsUrls }).(pulumi.StringArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AwsS3DataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AwsS3DataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsS3DataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AwsS3DataConnectorOutput{})
}
