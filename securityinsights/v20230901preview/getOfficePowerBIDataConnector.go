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
func LookupOfficePowerBIDataConnector(ctx *pulumi.Context, args *LookupOfficePowerBIDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupOfficePowerBIDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOfficePowerBIDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20230901preview:getOfficePowerBIDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOfficePowerBIDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents Office Microsoft PowerBI data connector.
type LookupOfficePowerBIDataConnectorResult struct {
	// The available data types for the connector.
	DataTypes OfficePowerBIConnectorDataTypesResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'OfficePowerBI'.
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

func LookupOfficePowerBIDataConnectorOutput(ctx *pulumi.Context, args LookupOfficePowerBIDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupOfficePowerBIDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOfficePowerBIDataConnectorResult, error) {
			args := v.(LookupOfficePowerBIDataConnectorArgs)
			r, err := LookupOfficePowerBIDataConnector(ctx, &args, opts...)
			var s LookupOfficePowerBIDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOfficePowerBIDataConnectorResultOutput)
}

type LookupOfficePowerBIDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupOfficePowerBIDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOfficePowerBIDataConnectorArgs)(nil)).Elem()
}

// Represents Office Microsoft PowerBI data connector.
type LookupOfficePowerBIDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupOfficePowerBIDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOfficePowerBIDataConnectorResult)(nil)).Elem()
}

func (o LookupOfficePowerBIDataConnectorResultOutput) ToLookupOfficePowerBIDataConnectorResultOutput() LookupOfficePowerBIDataConnectorResultOutput {
	return o
}

func (o LookupOfficePowerBIDataConnectorResultOutput) ToLookupOfficePowerBIDataConnectorResultOutputWithContext(ctx context.Context) LookupOfficePowerBIDataConnectorResultOutput {
	return o
}

func (o LookupOfficePowerBIDataConnectorResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOfficePowerBIDataConnectorResult] {
	return pulumix.Output[LookupOfficePowerBIDataConnectorResult]{
		OutputState: o.OutputState,
	}
}

// The available data types for the connector.
func (o LookupOfficePowerBIDataConnectorResultOutput) DataTypes() OfficePowerBIConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) OfficePowerBIConnectorDataTypesResponse {
		return v.DataTypes
	}).(OfficePowerBIConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o LookupOfficePowerBIDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupOfficePowerBIDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'OfficePowerBI'.
func (o LookupOfficePowerBIDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupOfficePowerBIDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupOfficePowerBIDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o LookupOfficePowerBIDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupOfficePowerBIDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOfficePowerBIDataConnectorResultOutput{})
}
