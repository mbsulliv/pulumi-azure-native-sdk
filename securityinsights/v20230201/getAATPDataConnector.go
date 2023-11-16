// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a data connector.
func LookupAATPDataConnector(ctx *pulumi.Context, args *LookupAATPDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAATPDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAATPDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20230201:getAATPDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAATPDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents AATP (Azure Advanced Threat Protection) data connector.
type LookupAATPDataConnectorResult struct {
	// The available data types for the connector.
	DataTypes *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'AzureAdvancedThreatProtection'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tenant id to connect to, and get the data from.
	TenantId *string `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAATPDataConnectorOutput(ctx *pulumi.Context, args LookupAATPDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAATPDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAATPDataConnectorResult, error) {
			args := v.(LookupAATPDataConnectorArgs)
			r, err := LookupAATPDataConnector(ctx, &args, opts...)
			var s LookupAATPDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAATPDataConnectorResultOutput)
}

type LookupAATPDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAATPDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAATPDataConnectorArgs)(nil)).Elem()
}

// Represents AATP (Azure Advanced Threat Protection) data connector.
type LookupAATPDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAATPDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAATPDataConnectorResult)(nil)).Elem()
}

func (o LookupAATPDataConnectorResultOutput) ToLookupAATPDataConnectorResultOutput() LookupAATPDataConnectorResultOutput {
	return o
}

func (o LookupAATPDataConnectorResultOutput) ToLookupAATPDataConnectorResultOutputWithContext(ctx context.Context) LookupAATPDataConnectorResultOutput {
	return o
}

// The available data types for the connector.
func (o LookupAATPDataConnectorResultOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) *AlertsDataTypeOfDataConnectorResponse { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

// Etag of the azure resource
func (o LookupAATPDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAATPDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'AzureAdvancedThreatProtection'.
func (o LookupAATPDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAATPDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAATPDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o LookupAATPDataConnectorResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAATPDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAATPDataConnectorResultOutput{})
}
