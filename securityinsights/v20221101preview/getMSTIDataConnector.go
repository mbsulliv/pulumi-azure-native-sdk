// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Microsoft Threat Intelligence data connector.
func LookupMSTIDataConnector(ctx *pulumi.Context, args *LookupMSTIDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupMSTIDataConnectorResult, error) {
	var rv LookupMSTIDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20221101preview:getMSTIDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMSTIDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents Microsoft Threat Intelligence data connector.
type LookupMSTIDataConnectorResult struct {
	// The available data types for the connector.
	DataTypes MSTIDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'MicrosoftThreatIntelligence'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupMSTIDataConnectorOutput(ctx *pulumi.Context, args LookupMSTIDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupMSTIDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMSTIDataConnectorResult, error) {
			args := v.(LookupMSTIDataConnectorArgs)
			r, err := LookupMSTIDataConnector(ctx, &args, opts...)
			var s LookupMSTIDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMSTIDataConnectorResultOutput)
}

type LookupMSTIDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMSTIDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMSTIDataConnectorArgs)(nil)).Elem()
}

// Represents Microsoft Threat Intelligence data connector.
type LookupMSTIDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupMSTIDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMSTIDataConnectorResult)(nil)).Elem()
}

func (o LookupMSTIDataConnectorResultOutput) ToLookupMSTIDataConnectorResultOutput() LookupMSTIDataConnectorResultOutput {
	return o
}

func (o LookupMSTIDataConnectorResultOutput) ToLookupMSTIDataConnectorResultOutputWithContext(ctx context.Context) LookupMSTIDataConnectorResultOutput {
	return o
}

// The available data types for the connector.
func (o LookupMSTIDataConnectorResultOutput) DataTypes() MSTIDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) MSTIDataConnectorDataTypesResponse { return v.DataTypes }).(MSTIDataConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o LookupMSTIDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupMSTIDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'MicrosoftThreatIntelligence'.
func (o LookupMSTIDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupMSTIDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMSTIDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o LookupMSTIDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMSTIDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMSTIDataConnectorResultOutput{})
}