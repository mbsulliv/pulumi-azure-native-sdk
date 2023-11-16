// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a data connector.
func LookupOfficeATPDataConnector(ctx *pulumi.Context, args *LookupOfficeATPDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupOfficeATPDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOfficeATPDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20230701preview:getOfficeATPDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOfficeATPDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents OfficeATP (Office 365 Advanced Threat Protection) data connector.
type LookupOfficeATPDataConnectorResult struct {
	// The available data types for the connector.
	DataTypes *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'OfficeATP'.
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

func LookupOfficeATPDataConnectorOutput(ctx *pulumi.Context, args LookupOfficeATPDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupOfficeATPDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOfficeATPDataConnectorResult, error) {
			args := v.(LookupOfficeATPDataConnectorArgs)
			r, err := LookupOfficeATPDataConnector(ctx, &args, opts...)
			var s LookupOfficeATPDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOfficeATPDataConnectorResultOutput)
}

type LookupOfficeATPDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupOfficeATPDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOfficeATPDataConnectorArgs)(nil)).Elem()
}

// Represents OfficeATP (Office 365 Advanced Threat Protection) data connector.
type LookupOfficeATPDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupOfficeATPDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOfficeATPDataConnectorResult)(nil)).Elem()
}

func (o LookupOfficeATPDataConnectorResultOutput) ToLookupOfficeATPDataConnectorResultOutput() LookupOfficeATPDataConnectorResultOutput {
	return o
}

func (o LookupOfficeATPDataConnectorResultOutput) ToLookupOfficeATPDataConnectorResultOutputWithContext(ctx context.Context) LookupOfficeATPDataConnectorResultOutput {
	return o
}

// The available data types for the connector.
func (o LookupOfficeATPDataConnectorResultOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) *AlertsDataTypeOfDataConnectorResponse { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

// Etag of the azure resource
func (o LookupOfficeATPDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupOfficeATPDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'OfficeATP'.
func (o LookupOfficeATPDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupOfficeATPDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupOfficeATPDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o LookupOfficeATPDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupOfficeATPDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOfficeATPDataConnectorResultOutput{})
}
