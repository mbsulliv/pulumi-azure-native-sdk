// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a data connector.
func LookupTIDataConnector(ctx *pulumi.Context, args *LookupTIDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupTIDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTIDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20230801preview:getTIDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTIDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents threat intelligence data connector.
type LookupTIDataConnectorResult struct {
	// The available data types for the connector.
	DataTypes TIDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'ThreatIntelligence'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The lookback period for the feed to be imported.
	TipLookbackPeriod *string `pulumi:"tipLookbackPeriod"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupTIDataConnectorOutput(ctx *pulumi.Context, args LookupTIDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupTIDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTIDataConnectorResult, error) {
			args := v.(LookupTIDataConnectorArgs)
			r, err := LookupTIDataConnector(ctx, &args, opts...)
			var s LookupTIDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTIDataConnectorResultOutput)
}

type LookupTIDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupTIDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTIDataConnectorArgs)(nil)).Elem()
}

// Represents threat intelligence data connector.
type LookupTIDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupTIDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTIDataConnectorResult)(nil)).Elem()
}

func (o LookupTIDataConnectorResultOutput) ToLookupTIDataConnectorResultOutput() LookupTIDataConnectorResultOutput {
	return o
}

func (o LookupTIDataConnectorResultOutput) ToLookupTIDataConnectorResultOutputWithContext(ctx context.Context) LookupTIDataConnectorResultOutput {
	return o
}

// The available data types for the connector.
func (o LookupTIDataConnectorResultOutput) DataTypes() TIDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) TIDataConnectorDataTypesResponse { return v.DataTypes }).(TIDataConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o LookupTIDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupTIDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'ThreatIntelligence'.
func (o LookupTIDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupTIDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupTIDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o LookupTIDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The lookback period for the feed to be imported.
func (o LookupTIDataConnectorResultOutput) TipLookbackPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) *string { return v.TipLookbackPeriod }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupTIDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTIDataConnectorResultOutput{})
}
