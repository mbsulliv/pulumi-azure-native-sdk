// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a data connector.
func LookupAwsS3DataConnector(ctx *pulumi.Context, args *LookupAwsS3DataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAwsS3DataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAwsS3DataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20230601preview:getAwsS3DataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAwsS3DataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents Amazon Web Services S3 data connector.
type LookupAwsS3DataConnectorResult struct {
	// The available data types for the connector.
	DataTypes AwsS3DataConnectorDataTypesResponse `pulumi:"dataTypes"`
	// The logs destination table name in LogAnalytics.
	DestinationTable string `pulumi:"destinationTable"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'AmazonWebServicesS3'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The Aws Role Arn that is used to access the Aws account.
	RoleArn string `pulumi:"roleArn"`
	// The AWS sqs urls for the connector.
	SqsUrls []string `pulumi:"sqsUrls"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAwsS3DataConnectorOutput(ctx *pulumi.Context, args LookupAwsS3DataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAwsS3DataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAwsS3DataConnectorResult, error) {
			args := v.(LookupAwsS3DataConnectorArgs)
			r, err := LookupAwsS3DataConnector(ctx, &args, opts...)
			var s LookupAwsS3DataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAwsS3DataConnectorResultOutput)
}

type LookupAwsS3DataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAwsS3DataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAwsS3DataConnectorArgs)(nil)).Elem()
}

// Represents Amazon Web Services S3 data connector.
type LookupAwsS3DataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAwsS3DataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAwsS3DataConnectorResult)(nil)).Elem()
}

func (o LookupAwsS3DataConnectorResultOutput) ToLookupAwsS3DataConnectorResultOutput() LookupAwsS3DataConnectorResultOutput {
	return o
}

func (o LookupAwsS3DataConnectorResultOutput) ToLookupAwsS3DataConnectorResultOutputWithContext(ctx context.Context) LookupAwsS3DataConnectorResultOutput {
	return o
}

func (o LookupAwsS3DataConnectorResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAwsS3DataConnectorResult] {
	return pulumix.Output[LookupAwsS3DataConnectorResult]{
		OutputState: o.OutputState,
	}
}

// The available data types for the connector.
func (o LookupAwsS3DataConnectorResultOutput) DataTypes() AwsS3DataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) AwsS3DataConnectorDataTypesResponse { return v.DataTypes }).(AwsS3DataConnectorDataTypesResponseOutput)
}

// The logs destination table name in LogAnalytics.
func (o LookupAwsS3DataConnectorResultOutput) DestinationTable() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) string { return v.DestinationTable }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o LookupAwsS3DataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAwsS3DataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'AmazonWebServicesS3'.
func (o LookupAwsS3DataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAwsS3DataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Aws Role Arn that is used to access the Aws account.
func (o LookupAwsS3DataConnectorResultOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) string { return v.RoleArn }).(pulumi.StringOutput)
}

// The AWS sqs urls for the connector.
func (o LookupAwsS3DataConnectorResultOutput) SqsUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) []string { return v.SqsUrls }).(pulumi.StringArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAwsS3DataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAwsS3DataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAwsS3DataConnectorResultOutput{})
}
