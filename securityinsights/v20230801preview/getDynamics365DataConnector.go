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
func LookupDynamics365DataConnector(ctx *pulumi.Context, args *LookupDynamics365DataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupDynamics365DataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDynamics365DataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20230801preview:getDynamics365DataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDynamics365DataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents Dynamics365 data connector.
type LookupDynamics365DataConnectorResult struct {
	// The available data types for the connector.
	DataTypes Dynamics365DataConnectorDataTypesResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'Dynamics365'.
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

func LookupDynamics365DataConnectorOutput(ctx *pulumi.Context, args LookupDynamics365DataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupDynamics365DataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDynamics365DataConnectorResult, error) {
			args := v.(LookupDynamics365DataConnectorArgs)
			r, err := LookupDynamics365DataConnector(ctx, &args, opts...)
			var s LookupDynamics365DataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDynamics365DataConnectorResultOutput)
}

type LookupDynamics365DataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupDynamics365DataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDynamics365DataConnectorArgs)(nil)).Elem()
}

// Represents Dynamics365 data connector.
type LookupDynamics365DataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupDynamics365DataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDynamics365DataConnectorResult)(nil)).Elem()
}

func (o LookupDynamics365DataConnectorResultOutput) ToLookupDynamics365DataConnectorResultOutput() LookupDynamics365DataConnectorResultOutput {
	return o
}

func (o LookupDynamics365DataConnectorResultOutput) ToLookupDynamics365DataConnectorResultOutputWithContext(ctx context.Context) LookupDynamics365DataConnectorResultOutput {
	return o
}

func (o LookupDynamics365DataConnectorResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDynamics365DataConnectorResult] {
	return pulumix.Output[LookupDynamics365DataConnectorResult]{
		OutputState: o.OutputState,
	}
}

// The available data types for the connector.
func (o LookupDynamics365DataConnectorResultOutput) DataTypes() Dynamics365DataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) Dynamics365DataConnectorDataTypesResponse {
		return v.DataTypes
	}).(Dynamics365DataConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o LookupDynamics365DataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDynamics365DataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'Dynamics365'.
func (o LookupDynamics365DataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDynamics365DataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDynamics365DataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o LookupDynamics365DataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDynamics365DataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDynamics365DataConnectorResultOutput{})
}
