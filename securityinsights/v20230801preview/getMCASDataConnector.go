// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a data connector.
func LookupMCASDataConnector(ctx *pulumi.Context, args *LookupMCASDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupMCASDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMCASDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20230801preview:getMCASDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMCASDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents MCAS (Microsoft Cloud App Security) data connector.
type LookupMCASDataConnectorResult struct {
	// The available data types for the connector.
	DataTypes MCASDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'MicrosoftCloudAppSecurity'.
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

func LookupMCASDataConnectorOutput(ctx *pulumi.Context, args LookupMCASDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupMCASDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMCASDataConnectorResult, error) {
			args := v.(LookupMCASDataConnectorArgs)
			r, err := LookupMCASDataConnector(ctx, &args, opts...)
			var s LookupMCASDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMCASDataConnectorResultOutput)
}

type LookupMCASDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMCASDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMCASDataConnectorArgs)(nil)).Elem()
}

// Represents MCAS (Microsoft Cloud App Security) data connector.
type LookupMCASDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupMCASDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMCASDataConnectorResult)(nil)).Elem()
}

func (o LookupMCASDataConnectorResultOutput) ToLookupMCASDataConnectorResultOutput() LookupMCASDataConnectorResultOutput {
	return o
}

func (o LookupMCASDataConnectorResultOutput) ToLookupMCASDataConnectorResultOutputWithContext(ctx context.Context) LookupMCASDataConnectorResultOutput {
	return o
}

func (o LookupMCASDataConnectorResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMCASDataConnectorResult] {
	return pulumix.Output[LookupMCASDataConnectorResult]{
		OutputState: o.OutputState,
	}
}

// The available data types for the connector.
func (o LookupMCASDataConnectorResultOutput) DataTypes() MCASDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) MCASDataConnectorDataTypesResponse { return v.DataTypes }).(MCASDataConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o LookupMCASDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupMCASDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'MicrosoftCloudAppSecurity'.
func (o LookupMCASDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupMCASDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMCASDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o LookupMCASDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMCASDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMCASDataConnectorResultOutput{})
}
