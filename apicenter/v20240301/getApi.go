// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns details of the API.
func LookupApi(ctx *pulumi.Context, args *LookupApiArgs, opts ...pulumi.InvokeOption) (*LookupApiResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApiResult
	err := ctx.Invoke("azure-native:apicenter/v20240301:getApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiArgs struct {
	// The name of the API.
	ApiName string `pulumi:"apiName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of Azure API Center service.
	ServiceName string `pulumi:"serviceName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// API entity.
type LookupApiResult struct {
	Contacts []ContactResponse `pulumi:"contacts"`
	// The custom metadata defined for API catalog entities.
	CustomProperties interface{} `pulumi:"customProperties"`
	// Description of the API.
	Description           *string                         `pulumi:"description"`
	ExternalDocumentation []ExternalDocumentationResponse `pulumi:"externalDocumentation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Kind of API. For example, REST or GraphQL.
	Kind string `pulumi:"kind"`
	// The license information for the API.
	License *LicenseResponse `pulumi:"license"`
	// Current lifecycle stage of the API.
	LifecycleStage string `pulumi:"lifecycleStage"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Short description of the API.
	Summary *string `pulumi:"summary"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Terms of service for the API.
	TermsOfService *TermsOfServiceResponse `pulumi:"termsOfService"`
	// API title.
	Title string `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupApiOutput(ctx *pulumi.Context, args LookupApiOutputArgs, opts ...pulumi.InvokeOption) LookupApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiResult, error) {
			args := v.(LookupApiArgs)
			r, err := LookupApi(ctx, &args, opts...)
			var s LookupApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiResultOutput)
}

type LookupApiOutputArgs struct {
	// The name of the API.
	ApiName pulumi.StringInput `pulumi:"apiName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of Azure API Center service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiArgs)(nil)).Elem()
}

// API entity.
type LookupApiResultOutput struct{ *pulumi.OutputState }

func (LookupApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiResult)(nil)).Elem()
}

func (o LookupApiResultOutput) ToLookupApiResultOutput() LookupApiResultOutput {
	return o
}

func (o LookupApiResultOutput) ToLookupApiResultOutputWithContext(ctx context.Context) LookupApiResultOutput {
	return o
}

func (o LookupApiResultOutput) Contacts() ContactResponseArrayOutput {
	return o.ApplyT(func(v LookupApiResult) []ContactResponse { return v.Contacts }).(ContactResponseArrayOutput)
}

// The custom metadata defined for API catalog entities.
func (o LookupApiResultOutput) CustomProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupApiResult) interface{} { return v.CustomProperties }).(pulumi.AnyOutput)
}

// Description of the API.
func (o LookupApiResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApiResultOutput) ExternalDocumentation() ExternalDocumentationResponseArrayOutput {
	return o.ApplyT(func(v LookupApiResult) []ExternalDocumentationResponse { return v.ExternalDocumentation }).(ExternalDocumentationResponseArrayOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of API. For example, REST or GraphQL.
func (o LookupApiResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The license information for the API.
func (o LookupApiResultOutput) License() LicenseResponsePtrOutput {
	return o.ApplyT(func(v LookupApiResult) *LicenseResponse { return v.License }).(LicenseResponsePtrOutput)
}

// Current lifecycle stage of the API.
func (o LookupApiResultOutput) LifecycleStage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.LifecycleStage }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Name }).(pulumi.StringOutput)
}

// Short description of the API.
func (o LookupApiResultOutput) Summary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.Summary }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupApiResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApiResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Terms of service for the API.
func (o LookupApiResultOutput) TermsOfService() TermsOfServiceResponsePtrOutput {
	return o.ApplyT(func(v LookupApiResult) *TermsOfServiceResponse { return v.TermsOfService }).(TermsOfServiceResponsePtrOutput)
}

// API title.
func (o LookupApiResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Title }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupApiResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiResultOutput{})
}
