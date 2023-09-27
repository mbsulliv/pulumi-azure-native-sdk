// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a data connector.
// Azure REST API version: 2023-02-01.
func LookupAwsCloudTrailDataConnector(ctx *pulumi.Context, args *LookupAwsCloudTrailDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAwsCloudTrailDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAwsCloudTrailDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights:getAwsCloudTrailDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAwsCloudTrailDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents Amazon Web Services CloudTrail data connector.
type LookupAwsCloudTrailDataConnectorResult struct {
	// The Aws Role Arn (with CloudTrailReadOnly policy) that is used to access the Aws account.
	AwsRoleArn *string `pulumi:"awsRoleArn"`
	// The available data types for the connector.
	DataTypes *AwsCloudTrailDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'AmazonWebServicesCloudTrail'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAwsCloudTrailDataConnectorOutput(ctx *pulumi.Context, args LookupAwsCloudTrailDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAwsCloudTrailDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAwsCloudTrailDataConnectorResult, error) {
			args := v.(LookupAwsCloudTrailDataConnectorArgs)
			r, err := LookupAwsCloudTrailDataConnector(ctx, &args, opts...)
			var s LookupAwsCloudTrailDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAwsCloudTrailDataConnectorResultOutput)
}

type LookupAwsCloudTrailDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAwsCloudTrailDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAwsCloudTrailDataConnectorArgs)(nil)).Elem()
}

// Represents Amazon Web Services CloudTrail data connector.
type LookupAwsCloudTrailDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAwsCloudTrailDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAwsCloudTrailDataConnectorResult)(nil)).Elem()
}

func (o LookupAwsCloudTrailDataConnectorResultOutput) ToLookupAwsCloudTrailDataConnectorResultOutput() LookupAwsCloudTrailDataConnectorResultOutput {
	return o
}

func (o LookupAwsCloudTrailDataConnectorResultOutput) ToLookupAwsCloudTrailDataConnectorResultOutputWithContext(ctx context.Context) LookupAwsCloudTrailDataConnectorResultOutput {
	return o
}

func (o LookupAwsCloudTrailDataConnectorResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAwsCloudTrailDataConnectorResult] {
	return pulumix.Output[LookupAwsCloudTrailDataConnectorResult]{
		OutputState: o.OutputState,
	}
}

// The Aws Role Arn (with CloudTrailReadOnly policy) that is used to access the Aws account.
func (o LookupAwsCloudTrailDataConnectorResultOutput) AwsRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) *string { return v.AwsRoleArn }).(pulumi.StringPtrOutput)
}

// The available data types for the connector.
func (o LookupAwsCloudTrailDataConnectorResultOutput) DataTypes() AwsCloudTrailDataConnectorDataTypesResponsePtrOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) *AwsCloudTrailDataConnectorDataTypesResponse {
		return v.DataTypes
	}).(AwsCloudTrailDataConnectorDataTypesResponsePtrOutput)
}

// Etag of the azure resource
func (o LookupAwsCloudTrailDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAwsCloudTrailDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'AmazonWebServicesCloudTrail'.
func (o LookupAwsCloudTrailDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAwsCloudTrailDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAwsCloudTrailDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAwsCloudTrailDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAwsCloudTrailDataConnectorResultOutput{})
}
