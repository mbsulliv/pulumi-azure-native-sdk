// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a data connector.
func LookupMicrosoftPurviewInformationProtectionDataConnector(ctx *pulumi.Context, args *LookupMicrosoftPurviewInformationProtectionDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupMicrosoftPurviewInformationProtectionDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMicrosoftPurviewInformationProtectionDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20240101preview:getMicrosoftPurviewInformationProtectionDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMicrosoftPurviewInformationProtectionDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents Microsoft Purview Information Protection data connector.
type LookupMicrosoftPurviewInformationProtectionDataConnectorResult struct {
	// The available data types for the connector.
	DataTypes MicrosoftPurviewInformationProtectionConnectorDataTypesResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'MicrosoftPurviewInformationProtection'.
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

func LookupMicrosoftPurviewInformationProtectionDataConnectorOutput(ctx *pulumi.Context, args LookupMicrosoftPurviewInformationProtectionDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMicrosoftPurviewInformationProtectionDataConnectorResult, error) {
			args := v.(LookupMicrosoftPurviewInformationProtectionDataConnectorArgs)
			r, err := LookupMicrosoftPurviewInformationProtectionDataConnector(ctx, &args, opts...)
			var s LookupMicrosoftPurviewInformationProtectionDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput)
}

type LookupMicrosoftPurviewInformationProtectionDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMicrosoftPurviewInformationProtectionDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMicrosoftPurviewInformationProtectionDataConnectorArgs)(nil)).Elem()
}

// Represents Microsoft Purview Information Protection data connector.
type LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMicrosoftPurviewInformationProtectionDataConnectorResult)(nil)).Elem()
}

func (o LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput) ToLookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput() LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput {
	return o
}

func (o LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput) ToLookupMicrosoftPurviewInformationProtectionDataConnectorResultOutputWithContext(ctx context.Context) LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput {
	return o
}

// The available data types for the connector.
func (o LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput) DataTypes() MicrosoftPurviewInformationProtectionConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupMicrosoftPurviewInformationProtectionDataConnectorResult) MicrosoftPurviewInformationProtectionConnectorDataTypesResponse {
		return v.DataTypes
	}).(MicrosoftPurviewInformationProtectionConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMicrosoftPurviewInformationProtectionDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftPurviewInformationProtectionDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'MicrosoftPurviewInformationProtection'.
func (o LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftPurviewInformationProtectionDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftPurviewInformationProtectionDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMicrosoftPurviewInformationProtectionDataConnectorResult) SystemDataResponse {
		return v.SystemData
	}).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftPurviewInformationProtectionDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftPurviewInformationProtectionDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMicrosoftPurviewInformationProtectionDataConnectorResultOutput{})
}
