// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a data connector.
func LookupAADDataConnector(ctx *pulumi.Context, args *LookupAADDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAADDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAADDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20240301:getAADDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAADDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents AAD (Azure Active Directory) data connector.
type LookupAADDataConnectorResult struct {
	// The available data types for the connector.
	DataTypes *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'AzureActiveDirectory'.
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

func LookupAADDataConnectorOutput(ctx *pulumi.Context, args LookupAADDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAADDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAADDataConnectorResult, error) {
			args := v.(LookupAADDataConnectorArgs)
			r, err := LookupAADDataConnector(ctx, &args, opts...)
			var s LookupAADDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAADDataConnectorResultOutput)
}

type LookupAADDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAADDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAADDataConnectorArgs)(nil)).Elem()
}

// Represents AAD (Azure Active Directory) data connector.
type LookupAADDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAADDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAADDataConnectorResult)(nil)).Elem()
}

func (o LookupAADDataConnectorResultOutput) ToLookupAADDataConnectorResultOutput() LookupAADDataConnectorResultOutput {
	return o
}

func (o LookupAADDataConnectorResultOutput) ToLookupAADDataConnectorResultOutputWithContext(ctx context.Context) LookupAADDataConnectorResultOutput {
	return o
}

// The available data types for the connector.
func (o LookupAADDataConnectorResultOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) *AlertsDataTypeOfDataConnectorResponse { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

// Etag of the azure resource
func (o LookupAADDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAADDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'AzureActiveDirectory'.
func (o LookupAADDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAADDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAADDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o LookupAADDataConnectorResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAADDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAADDataConnectorResultOutput{})
}
