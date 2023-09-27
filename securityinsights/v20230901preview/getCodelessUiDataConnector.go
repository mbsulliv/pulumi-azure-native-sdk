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
func LookupCodelessUiDataConnector(ctx *pulumi.Context, args *LookupCodelessUiDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupCodelessUiDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCodelessUiDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20230901preview:getCodelessUiDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCodelessUiDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents Codeless UI data connector.
type LookupCodelessUiDataConnectorResult struct {
	// Config to describe the instructions blade
	ConnectorUiConfig *CodelessUiConnectorConfigPropertiesResponse `pulumi:"connectorUiConfig"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'GenericUI'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupCodelessUiDataConnectorOutput(ctx *pulumi.Context, args LookupCodelessUiDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupCodelessUiDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCodelessUiDataConnectorResult, error) {
			args := v.(LookupCodelessUiDataConnectorArgs)
			r, err := LookupCodelessUiDataConnector(ctx, &args, opts...)
			var s LookupCodelessUiDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCodelessUiDataConnectorResultOutput)
}

type LookupCodelessUiDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupCodelessUiDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodelessUiDataConnectorArgs)(nil)).Elem()
}

// Represents Codeless UI data connector.
type LookupCodelessUiDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupCodelessUiDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodelessUiDataConnectorResult)(nil)).Elem()
}

func (o LookupCodelessUiDataConnectorResultOutput) ToLookupCodelessUiDataConnectorResultOutput() LookupCodelessUiDataConnectorResultOutput {
	return o
}

func (o LookupCodelessUiDataConnectorResultOutput) ToLookupCodelessUiDataConnectorResultOutputWithContext(ctx context.Context) LookupCodelessUiDataConnectorResultOutput {
	return o
}

func (o LookupCodelessUiDataConnectorResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCodelessUiDataConnectorResult] {
	return pulumix.Output[LookupCodelessUiDataConnectorResult]{
		OutputState: o.OutputState,
	}
}

// Config to describe the instructions blade
func (o LookupCodelessUiDataConnectorResultOutput) ConnectorUiConfig() CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupCodelessUiDataConnectorResult) *CodelessUiConnectorConfigPropertiesResponse {
		return v.ConnectorUiConfig
	}).(CodelessUiConnectorConfigPropertiesResponsePtrOutput)
}

// Etag of the azure resource
func (o LookupCodelessUiDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCodelessUiDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupCodelessUiDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodelessUiDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'GenericUI'.
func (o LookupCodelessUiDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodelessUiDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCodelessUiDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodelessUiDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupCodelessUiDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCodelessUiDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCodelessUiDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodelessUiDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCodelessUiDataConnectorResultOutput{})
}
