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
func LookupOffice365ProjectDataConnector(ctx *pulumi.Context, args *LookupOffice365ProjectDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupOffice365ProjectDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOffice365ProjectDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20231001preview:getOffice365ProjectDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOffice365ProjectDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents Office Microsoft Project data connector.
type LookupOffice365ProjectDataConnectorResult struct {
	// The available data types for the connector.
	DataTypes Office365ProjectConnectorDataTypesResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'Office365Project'.
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

func LookupOffice365ProjectDataConnectorOutput(ctx *pulumi.Context, args LookupOffice365ProjectDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupOffice365ProjectDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOffice365ProjectDataConnectorResult, error) {
			args := v.(LookupOffice365ProjectDataConnectorArgs)
			r, err := LookupOffice365ProjectDataConnector(ctx, &args, opts...)
			var s LookupOffice365ProjectDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOffice365ProjectDataConnectorResultOutput)
}

type LookupOffice365ProjectDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupOffice365ProjectDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOffice365ProjectDataConnectorArgs)(nil)).Elem()
}

// Represents Office Microsoft Project data connector.
type LookupOffice365ProjectDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupOffice365ProjectDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOffice365ProjectDataConnectorResult)(nil)).Elem()
}

func (o LookupOffice365ProjectDataConnectorResultOutput) ToLookupOffice365ProjectDataConnectorResultOutput() LookupOffice365ProjectDataConnectorResultOutput {
	return o
}

func (o LookupOffice365ProjectDataConnectorResultOutput) ToLookupOffice365ProjectDataConnectorResultOutputWithContext(ctx context.Context) LookupOffice365ProjectDataConnectorResultOutput {
	return o
}

// The available data types for the connector.
func (o LookupOffice365ProjectDataConnectorResultOutput) DataTypes() Office365ProjectConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) Office365ProjectConnectorDataTypesResponse {
		return v.DataTypes
	}).(Office365ProjectConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o LookupOffice365ProjectDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupOffice365ProjectDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'Office365Project'.
func (o LookupOffice365ProjectDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupOffice365ProjectDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupOffice365ProjectDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o LookupOffice365ProjectDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupOffice365ProjectDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOffice365ProjectDataConnectorResultOutput{})
}
