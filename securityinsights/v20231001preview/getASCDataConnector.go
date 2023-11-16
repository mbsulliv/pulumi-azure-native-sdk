// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a data connector.
func LookupASCDataConnector(ctx *pulumi.Context, args *LookupASCDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupASCDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupASCDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20231001preview:getASCDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupASCDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents ASC (Azure Security Center) data connector.
type LookupASCDataConnectorResult struct {
	// The available data types for the connector.
	DataTypes *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'AzureSecurityCenter'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The subscription id to connect to, and get the data from.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupASCDataConnectorOutput(ctx *pulumi.Context, args LookupASCDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupASCDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupASCDataConnectorResult, error) {
			args := v.(LookupASCDataConnectorArgs)
			r, err := LookupASCDataConnector(ctx, &args, opts...)
			var s LookupASCDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupASCDataConnectorResultOutput)
}

type LookupASCDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupASCDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupASCDataConnectorArgs)(nil)).Elem()
}

// Represents ASC (Azure Security Center) data connector.
type LookupASCDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupASCDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupASCDataConnectorResult)(nil)).Elem()
}

func (o LookupASCDataConnectorResultOutput) ToLookupASCDataConnectorResultOutput() LookupASCDataConnectorResultOutput {
	return o
}

func (o LookupASCDataConnectorResultOutput) ToLookupASCDataConnectorResultOutputWithContext(ctx context.Context) LookupASCDataConnectorResultOutput {
	return o
}

// The available data types for the connector.
func (o LookupASCDataConnectorResultOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) *AlertsDataTypeOfDataConnectorResponse { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

// Etag of the azure resource
func (o LookupASCDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupASCDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'AzureSecurityCenter'.
func (o LookupASCDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupASCDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// The subscription id to connect to, and get the data from.
func (o LookupASCDataConnectorResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupASCDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupASCDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupASCDataConnectorResultOutput{})
}
