// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get installed extension details by extension id.
func LookupExtension(ctx *pulumi.Context, args *LookupExtensionArgs, opts ...pulumi.InvokeOption) (*LookupExtensionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExtensionResult
	err := ctx.Invoke("azure-native:agfoodplatform/v20230601preview:getExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExtensionArgs struct {
	// DataManagerForAgriculture resource name.
	DataManagerForAgricultureResourceName string `pulumi:"dataManagerForAgricultureResourceName"`
	// Id of extension resource.
	ExtensionId string `pulumi:"extensionId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Extension resource.
type LookupExtensionResult struct {
	// Additional Api Properties.
	AdditionalApiProperties map[string]ApiPropertiesResponse `pulumi:"additionalApiProperties"`
	// The ETag value to implement optimistic concurrency.
	ETag string `pulumi:"eTag"`
	// Extension api docs link.
	ExtensionApiDocsLink string `pulumi:"extensionApiDocsLink"`
	// Extension auth link.
	ExtensionAuthLink string `pulumi:"extensionAuthLink"`
	// Extension category. e.g. weather/sensor/satellite.
	ExtensionCategory string `pulumi:"extensionCategory"`
	// Extension Id.
	ExtensionId string `pulumi:"extensionId"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Installed extension version.
	InstalledExtensionVersion string `pulumi:"installedExtensionVersion"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupExtensionOutput(ctx *pulumi.Context, args LookupExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupExtensionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExtensionResult, error) {
			args := v.(LookupExtensionArgs)
			r, err := LookupExtension(ctx, &args, opts...)
			var s LookupExtensionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExtensionResultOutput)
}

type LookupExtensionOutputArgs struct {
	// DataManagerForAgriculture resource name.
	DataManagerForAgricultureResourceName pulumi.StringInput `pulumi:"dataManagerForAgricultureResourceName"`
	// Id of extension resource.
	ExtensionId pulumi.StringInput `pulumi:"extensionId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtensionArgs)(nil)).Elem()
}

// Extension resource.
type LookupExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtensionResult)(nil)).Elem()
}

func (o LookupExtensionResultOutput) ToLookupExtensionResultOutput() LookupExtensionResultOutput {
	return o
}

func (o LookupExtensionResultOutput) ToLookupExtensionResultOutputWithContext(ctx context.Context) LookupExtensionResultOutput {
	return o
}

// Additional Api Properties.
func (o LookupExtensionResultOutput) AdditionalApiProperties() ApiPropertiesResponseMapOutput {
	return o.ApplyT(func(v LookupExtensionResult) map[string]ApiPropertiesResponse { return v.AdditionalApiProperties }).(ApiPropertiesResponseMapOutput)
}

// The ETag value to implement optimistic concurrency.
func (o LookupExtensionResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.ETag }).(pulumi.StringOutput)
}

// Extension api docs link.
func (o LookupExtensionResultOutput) ExtensionApiDocsLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.ExtensionApiDocsLink }).(pulumi.StringOutput)
}

// Extension auth link.
func (o LookupExtensionResultOutput) ExtensionAuthLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.ExtensionAuthLink }).(pulumi.StringOutput)
}

// Extension category. e.g. weather/sensor/satellite.
func (o LookupExtensionResultOutput) ExtensionCategory() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.ExtensionCategory }).(pulumi.StringOutput)
}

// Extension Id.
func (o LookupExtensionResultOutput) ExtensionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.ExtensionId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Installed extension version.
func (o LookupExtensionResultOutput) InstalledExtensionVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.InstalledExtensionVersion }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupExtensionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupExtensionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupExtensionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExtensionResultOutput{})
}