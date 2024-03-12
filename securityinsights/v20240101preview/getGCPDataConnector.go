// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a data connector.
func LookupGCPDataConnector(ctx *pulumi.Context, args *LookupGCPDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupGCPDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGCPDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20240101preview:getGCPDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGCPDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents Google Cloud Platform data connector.
type LookupGCPDataConnectorResult struct {
	// The auth section of the connector.
	Auth GCPAuthPropertiesResponse `pulumi:"auth"`
	// The name of the connector definition that represents the UI config.
	ConnectorDefinitionName string `pulumi:"connectorDefinitionName"`
	// The configuration of the destination of the data.
	DcrConfig *DCRConfigurationResponse `pulumi:"dcrConfig"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'GCP'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The request section of the connector.
	Request GCPRequestPropertiesResponse `pulumi:"request"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupGCPDataConnectorOutput(ctx *pulumi.Context, args LookupGCPDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupGCPDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGCPDataConnectorResult, error) {
			args := v.(LookupGCPDataConnectorArgs)
			r, err := LookupGCPDataConnector(ctx, &args, opts...)
			var s LookupGCPDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGCPDataConnectorResultOutput)
}

type LookupGCPDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupGCPDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGCPDataConnectorArgs)(nil)).Elem()
}

// Represents Google Cloud Platform data connector.
type LookupGCPDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupGCPDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGCPDataConnectorResult)(nil)).Elem()
}

func (o LookupGCPDataConnectorResultOutput) ToLookupGCPDataConnectorResultOutput() LookupGCPDataConnectorResultOutput {
	return o
}

func (o LookupGCPDataConnectorResultOutput) ToLookupGCPDataConnectorResultOutputWithContext(ctx context.Context) LookupGCPDataConnectorResultOutput {
	return o
}

// The auth section of the connector.
func (o LookupGCPDataConnectorResultOutput) Auth() GCPAuthPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGCPDataConnectorResult) GCPAuthPropertiesResponse { return v.Auth }).(GCPAuthPropertiesResponseOutput)
}

// The name of the connector definition that represents the UI config.
func (o LookupGCPDataConnectorResultOutput) ConnectorDefinitionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGCPDataConnectorResult) string { return v.ConnectorDefinitionName }).(pulumi.StringOutput)
}

// The configuration of the destination of the data.
func (o LookupGCPDataConnectorResultOutput) DcrConfig() DCRConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupGCPDataConnectorResult) *DCRConfigurationResponse { return v.DcrConfig }).(DCRConfigurationResponsePtrOutput)
}

// Etag of the azure resource
func (o LookupGCPDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGCPDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupGCPDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGCPDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'GCP'.
func (o LookupGCPDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGCPDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupGCPDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGCPDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// The request section of the connector.
func (o LookupGCPDataConnectorResultOutput) Request() GCPRequestPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGCPDataConnectorResult) GCPRequestPropertiesResponse { return v.Request }).(GCPRequestPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupGCPDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGCPDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupGCPDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGCPDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGCPDataConnectorResultOutput{})
}
