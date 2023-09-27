// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a data connector.
func LookupOfficeIRMDataConnector(ctx *pulumi.Context, args *LookupOfficeIRMDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupOfficeIRMDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOfficeIRMDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20230901preview:getOfficeIRMDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOfficeIRMDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents OfficeIRM (Microsoft Insider Risk Management) data connector.
type LookupOfficeIRMDataConnectorResult struct {
	// The available data types for the connector.
	DataTypes *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'OfficeIRM'.
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

func LookupOfficeIRMDataConnectorOutput(ctx *pulumi.Context, args LookupOfficeIRMDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupOfficeIRMDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOfficeIRMDataConnectorResult, error) {
			args := v.(LookupOfficeIRMDataConnectorArgs)
			r, err := LookupOfficeIRMDataConnector(ctx, &args, opts...)
			var s LookupOfficeIRMDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOfficeIRMDataConnectorResultOutput)
}

type LookupOfficeIRMDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupOfficeIRMDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOfficeIRMDataConnectorArgs)(nil)).Elem()
}

// Represents OfficeIRM (Microsoft Insider Risk Management) data connector.
type LookupOfficeIRMDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupOfficeIRMDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOfficeIRMDataConnectorResult)(nil)).Elem()
}

func (o LookupOfficeIRMDataConnectorResultOutput) ToLookupOfficeIRMDataConnectorResultOutput() LookupOfficeIRMDataConnectorResultOutput {
	return o
}

func (o LookupOfficeIRMDataConnectorResultOutput) ToLookupOfficeIRMDataConnectorResultOutputWithContext(ctx context.Context) LookupOfficeIRMDataConnectorResultOutput {
	return o
}

func (o LookupOfficeIRMDataConnectorResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOfficeIRMDataConnectorResult] {
	return pulumix.Output[LookupOfficeIRMDataConnectorResult]{
		OutputState: o.OutputState,
	}
}

// The available data types for the connector.
func (o LookupOfficeIRMDataConnectorResultOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) *AlertsDataTypeOfDataConnectorResponse { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

// Etag of the azure resource
func (o LookupOfficeIRMDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupOfficeIRMDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'OfficeIRM'.
func (o LookupOfficeIRMDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupOfficeIRMDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupOfficeIRMDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o LookupOfficeIRMDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupOfficeIRMDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOfficeIRMDataConnectorResultOutput{})
}
