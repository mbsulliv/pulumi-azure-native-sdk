// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the API specified by its identifier.
func LookupApi(ctx *pulumi.Context, args *LookupApiArgs, opts ...pulumi.InvokeOption) (*LookupApiResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApiResult
	err := ctx.Invoke("azure-native:apimanagement/v20180601preview:getApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// API details.
type LookupApiResult struct {
	// Describes the Revision of the Api. If no value is provided, default revision 1 is created
	ApiRevision *string `pulumi:"apiRevision"`
	// Description of the Api Revision.
	ApiRevisionDescription *string `pulumi:"apiRevisionDescription"`
	// Type of API.
	ApiType *string `pulumi:"apiType"`
	// Indicates the Version identifier of the API if the API is versioned
	ApiVersion *string `pulumi:"apiVersion"`
	// Description of the Api Version.
	ApiVersionDescription *string `pulumi:"apiVersionDescription"`
	// An API Version Set contains the common configuration for a set of API Versions relating
	ApiVersionSet *ApiVersionSetContractDetailsResponse `pulumi:"apiVersionSet"`
	// A resource identifier for the related ApiVersionSet.
	ApiVersionSetId *string `pulumi:"apiVersionSetId"`
	// Collection of authentication settings included into this API.
	AuthenticationSettings *AuthenticationSettingsContractResponse `pulumi:"authenticationSettings"`
	// Description of the API. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// API name.
	DisplayName *string `pulumi:"displayName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Indicates if API revision is current api revision.
	IsCurrent bool `pulumi:"isCurrent"`
	// Indicates if API revision is accessible via the gateway.
	IsOnline bool `pulumi:"isOnline"`
	// Resource name.
	Name string `pulumi:"name"`
	// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
	Path string `pulumi:"path"`
	// Describes on which protocols the operations in this API can be invoked.
	Protocols []string `pulumi:"protocols"`
	// Absolute URL of the backend service implementing this API.
	ServiceUrl *string `pulumi:"serviceUrl"`
	// Protocols over which API is made available.
	SubscriptionKeyParameterNames *SubscriptionKeyParameterNamesContractResponse `pulumi:"subscriptionKeyParameterNames"`
	// Specifies whether an API or Product subscription is required for accessing the API.
	SubscriptionRequired *bool `pulumi:"subscriptionRequired"`
	// Resource type for API Management resource.
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
	// The name of the resource group.
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

func (o LookupApiResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupApiResult] {
	return pulumix.Output[LookupApiResult]{
		OutputState: o.OutputState,
	}
}

// Describes the Revision of the Api. If no value is provided, default revision 1 is created
func (o LookupApiResultOutput) ApiRevision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiRevision }).(pulumi.StringPtrOutput)
}

// Description of the Api Revision.
func (o LookupApiResultOutput) ApiRevisionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiRevisionDescription }).(pulumi.StringPtrOutput)
}

// Type of API.
func (o LookupApiResultOutput) ApiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiType }).(pulumi.StringPtrOutput)
}

// Indicates the Version identifier of the API if the API is versioned
func (o LookupApiResultOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Description of the Api Version.
func (o LookupApiResultOutput) ApiVersionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiVersionDescription }).(pulumi.StringPtrOutput)
}

// An API Version Set contains the common configuration for a set of API Versions relating
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

// Description of the API. May include HTML formatting tags.
func (o LookupApiResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// API name.
func (o LookupApiResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates if API revision is current api revision.
func (o LookupApiResultOutput) IsCurrent() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApiResult) bool { return v.IsCurrent }).(pulumi.BoolOutput)
}

// Indicates if API revision is accessible via the gateway.
func (o LookupApiResultOutput) IsOnline() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApiResult) bool { return v.IsOnline }).(pulumi.BoolOutput)
}

// Resource name.
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

// Absolute URL of the backend service implementing this API.
func (o LookupApiResultOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ServiceUrl }).(pulumi.StringPtrOutput)
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

// Resource type for API Management resource.
func (o LookupApiResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiResultOutput{})
}
