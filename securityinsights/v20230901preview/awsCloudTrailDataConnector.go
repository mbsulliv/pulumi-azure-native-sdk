// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Amazon Web Services CloudTrail data connector.
type AwsCloudTrailDataConnector struct {
	pulumi.CustomResourceState

	// The Aws Role Arn (with CloudTrailReadOnly policy) that is used to access the Aws account.
	AwsRoleArn pulumi.StringPtrOutput `pulumi:"awsRoleArn"`
	// The available data types for the connector.
	DataTypes AwsCloudTrailDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'AmazonWebServicesCloudTrail'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAwsCloudTrailDataConnector registers a new resource with the given unique name, arguments, and options.
func NewAwsCloudTrailDataConnector(ctx *pulumi.Context,
	name string, args *AwsCloudTrailDataConnectorArgs, opts ...pulumi.ResourceOption) (*AwsCloudTrailDataConnector, error) {
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
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("AmazonWebServicesCloudTrail")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:AwsCloudTrailDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:AwsCloudTrailDataConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AwsCloudTrailDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20230901preview:AwsCloudTrailDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAwsCloudTrailDataConnector gets an existing AwsCloudTrailDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAwsCloudTrailDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsCloudTrailDataConnectorState, opts ...pulumi.ResourceOption) (*AwsCloudTrailDataConnector, error) {
	var resource AwsCloudTrailDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20230901preview:AwsCloudTrailDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AwsCloudTrailDataConnector resources.
type awsCloudTrailDataConnectorState struct {
}

type AwsCloudTrailDataConnectorState struct {
}

func (AwsCloudTrailDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsCloudTrailDataConnectorState)(nil)).Elem()
}

type awsCloudTrailDataConnectorArgs struct {
	// The Aws Role Arn (with CloudTrailReadOnly policy) that is used to access the Aws account.
	AwsRoleArn *string `pulumi:"awsRoleArn"`
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes AwsCloudTrailDataConnectorDataTypes `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'AmazonWebServicesCloudTrail'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a AwsCloudTrailDataConnector resource.
type AwsCloudTrailDataConnectorArgs struct {
	// The Aws Role Arn (with CloudTrailReadOnly policy) that is used to access the Aws account.
	AwsRoleArn pulumi.StringPtrInput
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes AwsCloudTrailDataConnectorDataTypesInput
	// The kind of the data connector
	// Expected value is 'AmazonWebServicesCloudTrail'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (AwsCloudTrailDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsCloudTrailDataConnectorArgs)(nil)).Elem()
}

type AwsCloudTrailDataConnectorInput interface {
	pulumi.Input

	ToAwsCloudTrailDataConnectorOutput() AwsCloudTrailDataConnectorOutput
	ToAwsCloudTrailDataConnectorOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorOutput
}

func (*AwsCloudTrailDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCloudTrailDataConnector)(nil)).Elem()
}

func (i *AwsCloudTrailDataConnector) ToAwsCloudTrailDataConnectorOutput() AwsCloudTrailDataConnectorOutput {
	return i.ToAwsCloudTrailDataConnectorOutputWithContext(context.Background())
}

func (i *AwsCloudTrailDataConnector) ToAwsCloudTrailDataConnectorOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCloudTrailDataConnectorOutput)
}

type AwsCloudTrailDataConnectorOutput struct{ *pulumi.OutputState }

func (AwsCloudTrailDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCloudTrailDataConnector)(nil)).Elem()
}

func (o AwsCloudTrailDataConnectorOutput) ToAwsCloudTrailDataConnectorOutput() AwsCloudTrailDataConnectorOutput {
	return o
}

func (o AwsCloudTrailDataConnectorOutput) ToAwsCloudTrailDataConnectorOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorOutput {
	return o
}

// The Aws Role Arn (with CloudTrailReadOnly policy) that is used to access the Aws account.
func (o AwsCloudTrailDataConnectorOutput) AwsRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnector) pulumi.StringPtrOutput { return v.AwsRoleArn }).(pulumi.StringPtrOutput)
}

// The available data types for the connector.
func (o AwsCloudTrailDataConnectorOutput) DataTypes() AwsCloudTrailDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnector) AwsCloudTrailDataConnectorDataTypesResponseOutput {
		return v.DataTypes
	}).(AwsCloudTrailDataConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o AwsCloudTrailDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'AmazonWebServicesCloudTrail'.
func (o AwsCloudTrailDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o AwsCloudTrailDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AwsCloudTrailDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AwsCloudTrailDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorOutput{})
}
