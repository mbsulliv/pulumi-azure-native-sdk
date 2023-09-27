// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a data connector definition.
// Azure REST API version: 2023-07-01-preview.
func LookupCustomizableConnectorDefinition(ctx *pulumi.Context, args *LookupCustomizableConnectorDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupCustomizableConnectorDefinitionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCustomizableConnectorDefinitionResult
	err := ctx.Invoke("azure-native:securityinsights:getCustomizableConnectorDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomizableConnectorDefinitionArgs struct {
	// The data connector definition name.
	DataConnectorDefinitionName string `pulumi:"dataConnectorDefinitionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Connector definition for kind 'Customizable'.
type LookupCustomizableConnectorDefinitionResult struct {
	// The UiConfig for 'Customizable' connector definition kind.
	ConnectionsConfig *CustomizableConnectionsConfigResponse `pulumi:"connectionsConfig"`
	// The UiConfig for 'Customizable' connector definition kind.
	ConnectorUiConfig CustomizableConnectorUiConfigResponse `pulumi:"connectorUiConfig"`
	// Gets or sets the connector definition created date in UTC format.
	CreatedTimeUtc *string `pulumi:"createdTimeUtc"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the data connector definitions
	// Expected value is 'Customizable'.
	Kind string `pulumi:"kind"`
	// Gets or sets the connector definition last modified date in UTC format.
	LastModifiedUtc *string `pulumi:"lastModifiedUtc"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupCustomizableConnectorDefinitionOutput(ctx *pulumi.Context, args LookupCustomizableConnectorDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupCustomizableConnectorDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomizableConnectorDefinitionResult, error) {
			args := v.(LookupCustomizableConnectorDefinitionArgs)
			r, err := LookupCustomizableConnectorDefinition(ctx, &args, opts...)
			var s LookupCustomizableConnectorDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomizableConnectorDefinitionResultOutput)
}

type LookupCustomizableConnectorDefinitionOutputArgs struct {
	// The data connector definition name.
	DataConnectorDefinitionName pulumi.StringInput `pulumi:"dataConnectorDefinitionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupCustomizableConnectorDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomizableConnectorDefinitionArgs)(nil)).Elem()
}

// Connector definition for kind 'Customizable'.
type LookupCustomizableConnectorDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupCustomizableConnectorDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomizableConnectorDefinitionResult)(nil)).Elem()
}

func (o LookupCustomizableConnectorDefinitionResultOutput) ToLookupCustomizableConnectorDefinitionResultOutput() LookupCustomizableConnectorDefinitionResultOutput {
	return o
}

func (o LookupCustomizableConnectorDefinitionResultOutput) ToLookupCustomizableConnectorDefinitionResultOutputWithContext(ctx context.Context) LookupCustomizableConnectorDefinitionResultOutput {
	return o
}

func (o LookupCustomizableConnectorDefinitionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCustomizableConnectorDefinitionResult] {
	return pulumix.Output[LookupCustomizableConnectorDefinitionResult]{
		OutputState: o.OutputState,
	}
}

// The UiConfig for 'Customizable' connector definition kind.
func (o LookupCustomizableConnectorDefinitionResultOutput) ConnectionsConfig() CustomizableConnectionsConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomizableConnectorDefinitionResult) *CustomizableConnectionsConfigResponse {
		return v.ConnectionsConfig
	}).(CustomizableConnectionsConfigResponsePtrOutput)
}

// The UiConfig for 'Customizable' connector definition kind.
func (o LookupCustomizableConnectorDefinitionResultOutput) ConnectorUiConfig() CustomizableConnectorUiConfigResponseOutput {
	return o.ApplyT(func(v LookupCustomizableConnectorDefinitionResult) CustomizableConnectorUiConfigResponse {
		return v.ConnectorUiConfig
	}).(CustomizableConnectorUiConfigResponseOutput)
}

// Gets or sets the connector definition created date in UTC format.
func (o LookupCustomizableConnectorDefinitionResultOutput) CreatedTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomizableConnectorDefinitionResult) *string { return v.CreatedTimeUtc }).(pulumi.StringPtrOutput)
}

// Etag of the azure resource
func (o LookupCustomizableConnectorDefinitionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomizableConnectorDefinitionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupCustomizableConnectorDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomizableConnectorDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector definitions
// Expected value is 'Customizable'.
func (o LookupCustomizableConnectorDefinitionResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomizableConnectorDefinitionResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Gets or sets the connector definition last modified date in UTC format.
func (o LookupCustomizableConnectorDefinitionResultOutput) LastModifiedUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomizableConnectorDefinitionResult) *string { return v.LastModifiedUtc }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupCustomizableConnectorDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomizableConnectorDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupCustomizableConnectorDefinitionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCustomizableConnectorDefinitionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCustomizableConnectorDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomizableConnectorDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomizableConnectorDefinitionResultOutput{})
}
