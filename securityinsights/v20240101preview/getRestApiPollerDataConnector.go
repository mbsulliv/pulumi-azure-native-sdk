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
func LookupRestApiPollerDataConnector(ctx *pulumi.Context, args *LookupRestApiPollerDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupRestApiPollerDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRestApiPollerDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20240101preview:getRestApiPollerDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRestApiPollerDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents Rest Api Poller data connector.
type LookupRestApiPollerDataConnectorResult struct {
	// The add on attributes. The key name will become attribute name (a column) and the value will become the attribute value in the payload.
	AddOnAttributes map[string]string `pulumi:"addOnAttributes"`
	// The a authentication model.
	Auth interface{} `pulumi:"auth"`
	// The connector definition name (the dataConnectorDefinition resource id).
	ConnectorDefinitionName string `pulumi:"connectorDefinitionName"`
	// The Log Analytics table destination.
	DataType *string `pulumi:"dataType"`
	// The DCR related properties.
	DcrConfig *DCRConfigurationResponse `pulumi:"dcrConfig"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Indicates whether the connector is active or not.
	IsActive *bool `pulumi:"isActive"`
	// The kind of the data connector
	// Expected value is 'RestApiPoller'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The paging configuration.
	Paging *RestApiPollerRequestPagingConfigResponse `pulumi:"paging"`
	// The request configuration.
	Request RestApiPollerRequestConfigResponse `pulumi:"request"`
	// The response configuration.
	Response *CcpResponseConfigResponse `pulumi:"response"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupRestApiPollerDataConnectorResult
func (val *LookupRestApiPollerDataConnectorResult) Defaults() *LookupRestApiPollerDataConnectorResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Response = tmp.Response.Defaults()

	return &tmp
}

func LookupRestApiPollerDataConnectorOutput(ctx *pulumi.Context, args LookupRestApiPollerDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupRestApiPollerDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRestApiPollerDataConnectorResult, error) {
			args := v.(LookupRestApiPollerDataConnectorArgs)
			r, err := LookupRestApiPollerDataConnector(ctx, &args, opts...)
			var s LookupRestApiPollerDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRestApiPollerDataConnectorResultOutput)
}

type LookupRestApiPollerDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupRestApiPollerDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestApiPollerDataConnectorArgs)(nil)).Elem()
}

// Represents Rest Api Poller data connector.
type LookupRestApiPollerDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupRestApiPollerDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestApiPollerDataConnectorResult)(nil)).Elem()
}

func (o LookupRestApiPollerDataConnectorResultOutput) ToLookupRestApiPollerDataConnectorResultOutput() LookupRestApiPollerDataConnectorResultOutput {
	return o
}

func (o LookupRestApiPollerDataConnectorResultOutput) ToLookupRestApiPollerDataConnectorResultOutputWithContext(ctx context.Context) LookupRestApiPollerDataConnectorResultOutput {
	return o
}

// The add on attributes. The key name will become attribute name (a column) and the value will become the attribute value in the payload.
func (o LookupRestApiPollerDataConnectorResultOutput) AddOnAttributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) map[string]string { return v.AddOnAttributes }).(pulumi.StringMapOutput)
}

// The a authentication model.
func (o LookupRestApiPollerDataConnectorResultOutput) Auth() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) interface{} { return v.Auth }).(pulumi.AnyOutput)
}

// The connector definition name (the dataConnectorDefinition resource id).
func (o LookupRestApiPollerDataConnectorResultOutput) ConnectorDefinitionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) string { return v.ConnectorDefinitionName }).(pulumi.StringOutput)
}

// The Log Analytics table destination.
func (o LookupRestApiPollerDataConnectorResultOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

// The DCR related properties.
func (o LookupRestApiPollerDataConnectorResultOutput) DcrConfig() DCRConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) *DCRConfigurationResponse { return v.DcrConfig }).(DCRConfigurationResponsePtrOutput)
}

// Etag of the azure resource
func (o LookupRestApiPollerDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRestApiPollerDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates whether the connector is active or not.
func (o LookupRestApiPollerDataConnectorResultOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) *bool { return v.IsActive }).(pulumi.BoolPtrOutput)
}

// The kind of the data connector
// Expected value is 'RestApiPoller'.
func (o LookupRestApiPollerDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupRestApiPollerDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// The paging configuration.
func (o LookupRestApiPollerDataConnectorResultOutput) Paging() RestApiPollerRequestPagingConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) *RestApiPollerRequestPagingConfigResponse {
		return v.Paging
	}).(RestApiPollerRequestPagingConfigResponsePtrOutput)
}

// The request configuration.
func (o LookupRestApiPollerDataConnectorResultOutput) Request() RestApiPollerRequestConfigResponseOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) RestApiPollerRequestConfigResponse { return v.Request }).(RestApiPollerRequestConfigResponseOutput)
}

// The response configuration.
func (o LookupRestApiPollerDataConnectorResultOutput) Response() CcpResponseConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) *CcpResponseConfigResponse { return v.Response }).(CcpResponseConfigResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRestApiPollerDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRestApiPollerDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestApiPollerDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRestApiPollerDataConnectorResultOutput{})
}
