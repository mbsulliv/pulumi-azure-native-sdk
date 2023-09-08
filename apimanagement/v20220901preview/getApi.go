// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the API specified by its identifier.
func LookupApi(ctx *pulumi.Context, args *LookupApiArgs, opts ...pulumi.InvokeOption) (*LookupApiResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApiResult
	err := ctx.Invoke("azure-native:apimanagement/v20220901preview:getApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// API details.
type LookupApiResult struct {
	// Describes the revision of the API. If no value is provided, default revision 1 is created
	ApiRevision *string `pulumi:"apiRevision"`
	// Description of the API Revision.
	ApiRevisionDescription *string `pulumi:"apiRevisionDescription"`
	// Type of API.
	ApiType *string `pulumi:"apiType"`
	// Indicates the version identifier of the API if the API is versioned
	ApiVersion *string `pulumi:"apiVersion"`
	// Description of the API Version.
	ApiVersionDescription *string `pulumi:"apiVersionDescription"`
	// Version set details
	ApiVersionSet *ApiVersionSetContractDetailsResponse `pulumi:"apiVersionSet"`
	// A resource identifier for the related ApiVersionSet.
	ApiVersionSetId *string `pulumi:"apiVersionSetId"`
	// Collection of authentication settings included into this API.
	AuthenticationSettings *AuthenticationSettingsContractResponse `pulumi:"authenticationSettings"`
	// Contact information for the API.
	Contact *ApiContactInformationResponse `pulumi:"contact"`
	// Description of the API. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// API name. Must be 1 to 300 characters long.
	DisplayName *string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Indicates if API revision is current api revision.
	IsCurrent *bool `pulumi:"isCurrent"`
	// Indicates if API revision is accessible via the gateway.
	IsOnline bool `pulumi:"isOnline"`
	// License information for the API.
	License *ApiLicenseInformationResponse `pulumi:"license"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
	Path string `pulumi:"path"`
	// Describes on which protocols the operations in this API can be invoked.
	Protocols []string `pulumi:"protocols"`
	// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
	ServiceUrl *string `pulumi:"serviceUrl"`
	// API identifier of the source API.
	SourceApiId *string `pulumi:"sourceApiId"`
	// Protocols over which API is made available.
	SubscriptionKeyParameterNames *SubscriptionKeyParameterNamesContractResponse `pulumi:"subscriptionKeyParameterNames"`
	// Specifies whether an API or Product subscription is required for accessing the API.
	SubscriptionRequired *bool `pulumi:"subscriptionRequired"`
	//  A URL to the Terms of Service for the API. MUST be in the format of a URL.
	TermsOfServiceUrl *string `pulumi:"termsOfServiceUrl"`
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
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiArgs)(nil)).Elem()
}

// API details.
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

// Describes the revision of the API. If no value is provided, default revision 1 is created
func (o LookupApiResultOutput) ApiRevision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiRevision }).(pulumi.StringPtrOutput)
}

// Description of the API Revision.
func (o LookupApiResultOutput) ApiRevisionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiRevisionDescription }).(pulumi.StringPtrOutput)
}

// Type of API.
func (o LookupApiResultOutput) ApiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiType }).(pulumi.StringPtrOutput)
}

// Indicates the version identifier of the API if the API is versioned
func (o LookupApiResultOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Description of the API Version.
func (o LookupApiResultOutput) ApiVersionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiVersionDescription }).(pulumi.StringPtrOutput)
}

// Version set details
func (o LookupApiResultOutput) ApiVersionSet() ApiVersionSetContractDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupApiResult) *ApiVersionSetContractDetailsResponse { return v.ApiVersionSet }).(ApiVersionSetContractDetailsResponsePtrOutput)
}

// A resource identifier for the related ApiVersionSet.
func (o LookupApiResultOutput) ApiVersionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiVersionSetId }).(pulumi.StringPtrOutput)
}

// Collection of authentication settings included into this API.
func (o LookupApiResultOutput) AuthenticationSettings() AuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyT(func(v LookupApiResult) *AuthenticationSettingsContractResponse { return v.AuthenticationSettings }).(AuthenticationSettingsContractResponsePtrOutput)
}

// Contact information for the API.
func (o LookupApiResultOutput) Contact() ApiContactInformationResponsePtrOutput {
	return o.ApplyT(func(v LookupApiResult) *ApiContactInformationResponse { return v.Contact }).(ApiContactInformationResponsePtrOutput)
}

// Description of the API. May include HTML formatting tags.
func (o LookupApiResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// API name. Must be 1 to 300 characters long.
func (o LookupApiResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates if API revision is current api revision.
func (o LookupApiResultOutput) IsCurrent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *bool { return v.IsCurrent }).(pulumi.BoolPtrOutput)
}

// Indicates if API revision is accessible via the gateway.
func (o LookupApiResultOutput) IsOnline() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApiResult) bool { return v.IsOnline }).(pulumi.BoolOutput)
}

// License information for the API.
func (o LookupApiResultOutput) License() ApiLicenseInformationResponsePtrOutput {
	return o.ApplyT(func(v LookupApiResult) *ApiLicenseInformationResponse { return v.License }).(ApiLicenseInformationResponsePtrOutput)
}

// The name of the resource
func (o LookupApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Name }).(pulumi.StringOutput)
}

// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
func (o LookupApiResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Path }).(pulumi.StringOutput)
}

// Describes on which protocols the operations in this API can be invoked.
func (o LookupApiResultOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApiResult) []string { return v.Protocols }).(pulumi.StringArrayOutput)
}

// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
func (o LookupApiResultOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

// API identifier of the source API.
func (o LookupApiResultOutput) SourceApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.SourceApiId }).(pulumi.StringPtrOutput)
}

// Protocols over which API is made available.
func (o LookupApiResultOutput) SubscriptionKeyParameterNames() SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return o.ApplyT(func(v LookupApiResult) *SubscriptionKeyParameterNamesContractResponse {
		return v.SubscriptionKeyParameterNames
	}).(SubscriptionKeyParameterNamesContractResponsePtrOutput)
}

// Specifies whether an API or Product subscription is required for accessing the API.
func (o LookupApiResultOutput) SubscriptionRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *bool { return v.SubscriptionRequired }).(pulumi.BoolPtrOutput)
}

// A URL to the Terms of Service for the API. MUST be in the format of a URL.
func (o LookupApiResultOutput) TermsOfServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.TermsOfServiceUrl }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupApiResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiResultOutput{})
}