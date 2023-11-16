// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// API details.
// Azure REST API version: 2022-08-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2016-07-07, 2016-10-10, 2017-03-01, 2018-06-01-preview, 2020-12-01, 2022-09-01-preview, 2023-03-01-preview.
type Api struct {
	pulumi.CustomResourceState

	// Describes the revision of the API. If no value is provided, default revision 1 is created
	ApiRevision pulumi.StringPtrOutput `pulumi:"apiRevision"`
	// Description of the API Revision.
	ApiRevisionDescription pulumi.StringPtrOutput `pulumi:"apiRevisionDescription"`
	// Type of API.
	ApiType pulumi.StringPtrOutput `pulumi:"apiType"`
	// Indicates the version identifier of the API if the API is versioned
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Description of the API Version.
	ApiVersionDescription pulumi.StringPtrOutput `pulumi:"apiVersionDescription"`
	// Version set details
	ApiVersionSet ApiVersionSetContractDetailsResponsePtrOutput `pulumi:"apiVersionSet"`
	// A resource identifier for the related ApiVersionSet.
	ApiVersionSetId pulumi.StringPtrOutput `pulumi:"apiVersionSetId"`
	// Collection of authentication settings included into this API.
	AuthenticationSettings AuthenticationSettingsContractResponsePtrOutput `pulumi:"authenticationSettings"`
	// Contact information for the API.
	Contact ApiContactInformationResponsePtrOutput `pulumi:"contact"`
	// Description of the API. May include HTML formatting tags.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// API name. Must be 1 to 300 characters long.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Indicates if API revision is current api revision.
	IsCurrent pulumi.BoolPtrOutput `pulumi:"isCurrent"`
	// Indicates if API revision is accessible via the gateway.
	IsOnline pulumi.BoolOutput `pulumi:"isOnline"`
	// License information for the API.
	License ApiLicenseInformationResponsePtrOutput `pulumi:"license"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
	Path pulumi.StringOutput `pulumi:"path"`
	// Describes on which protocols the operations in this API can be invoked.
	Protocols pulumi.StringArrayOutput `pulumi:"protocols"`
	// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
	ServiceUrl pulumi.StringPtrOutput `pulumi:"serviceUrl"`
	// API identifier of the source API.
	SourceApiId pulumi.StringPtrOutput `pulumi:"sourceApiId"`
	// Protocols over which API is made available.
	SubscriptionKeyParameterNames SubscriptionKeyParameterNamesContractResponsePtrOutput `pulumi:"subscriptionKeyParameterNames"`
	// Specifies whether an API or Product subscription is required for accessing the API.
	SubscriptionRequired pulumi.BoolPtrOutput `pulumi:"subscriptionRequired"`
	//  A URL to the Terms of Service for the API. MUST be in the format of a URL.
	TermsOfServiceUrl pulumi.StringPtrOutput `pulumi:"termsOfServiceUrl"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:Api"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Api
	err := ctx.RegisterResource("azure-native:apimanagement:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("azure-native:apimanagement:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type apiState struct {
}

type ApiState struct {
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId *string `pulumi:"apiId"`
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
	ApiVersionSet *ApiVersionSetContractDetails `pulumi:"apiVersionSet"`
	// A resource identifier for the related ApiVersionSet.
	ApiVersionSetId *string `pulumi:"apiVersionSetId"`
	// Collection of authentication settings included into this API.
	AuthenticationSettings *AuthenticationSettingsContract `pulumi:"authenticationSettings"`
	// Contact information for the API.
	Contact *ApiContactInformation `pulumi:"contact"`
	// Description of the API. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// API name. Must be 1 to 300 characters long.
	DisplayName *string `pulumi:"displayName"`
	// Format of the Content in which the API is getting imported.
	Format *string `pulumi:"format"`
	// Indicates if API revision is current api revision.
	IsCurrent *bool `pulumi:"isCurrent"`
	// License information for the API.
	License *ApiLicenseInformation `pulumi:"license"`
	// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
	Path string `pulumi:"path"`
	// Describes on which protocols the operations in this API can be invoked.
	Protocols []string `pulumi:"protocols"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
	ServiceUrl *string `pulumi:"serviceUrl"`
	// Type of API to create.
	//  * `http` creates a REST API
	//  * `soap` creates a SOAP pass-through API
	//  * `websocket` creates websocket API
	//  * `graphql` creates GraphQL API.
	SoapApiType *string `pulumi:"soapApiType"`
	// API identifier of the source API.
	SourceApiId *string `pulumi:"sourceApiId"`
	// Protocols over which API is made available.
	SubscriptionKeyParameterNames *SubscriptionKeyParameterNamesContract `pulumi:"subscriptionKeyParameterNames"`
	// Specifies whether an API or Product subscription is required for accessing the API.
	SubscriptionRequired *bool `pulumi:"subscriptionRequired"`
	//  A URL to the Terms of Service for the API. MUST be in the format of a URL.
	TermsOfServiceUrl *string `pulumi:"termsOfServiceUrl"`
	// Strategy of translating required query parameters to template ones. By default has value 'template'. Possible values: 'template', 'query'
	TranslateRequiredQueryParametersConduct *string `pulumi:"translateRequiredQueryParametersConduct"`
	// Content value when Importing an API.
	Value *string `pulumi:"value"`
	// Criteria to limit import of WSDL to a subset of the document.
	WsdlSelector *ApiCreateOrUpdatePropertiesWsdlSelector `pulumi:"wsdlSelector"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringPtrInput
	// Describes the revision of the API. If no value is provided, default revision 1 is created
	ApiRevision pulumi.StringPtrInput
	// Description of the API Revision.
	ApiRevisionDescription pulumi.StringPtrInput
	// Type of API.
	ApiType pulumi.StringPtrInput
	// Indicates the version identifier of the API if the API is versioned
	ApiVersion pulumi.StringPtrInput
	// Description of the API Version.
	ApiVersionDescription pulumi.StringPtrInput
	// Version set details
	ApiVersionSet ApiVersionSetContractDetailsPtrInput
	// A resource identifier for the related ApiVersionSet.
	ApiVersionSetId pulumi.StringPtrInput
	// Collection of authentication settings included into this API.
	AuthenticationSettings AuthenticationSettingsContractPtrInput
	// Contact information for the API.
	Contact ApiContactInformationPtrInput
	// Description of the API. May include HTML formatting tags.
	Description pulumi.StringPtrInput
	// API name. Must be 1 to 300 characters long.
	DisplayName pulumi.StringPtrInput
	// Format of the Content in which the API is getting imported.
	Format pulumi.StringPtrInput
	// Indicates if API revision is current api revision.
	IsCurrent pulumi.BoolPtrInput
	// License information for the API.
	License ApiLicenseInformationPtrInput
	// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
	Path pulumi.StringInput
	// Describes on which protocols the operations in this API can be invoked.
	Protocols pulumi.StringArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
	ServiceUrl pulumi.StringPtrInput
	// Type of API to create.
	//  * `http` creates a REST API
	//  * `soap` creates a SOAP pass-through API
	//  * `websocket` creates websocket API
	//  * `graphql` creates GraphQL API.
	SoapApiType pulumi.StringPtrInput
	// API identifier of the source API.
	SourceApiId pulumi.StringPtrInput
	// Protocols over which API is made available.
	SubscriptionKeyParameterNames SubscriptionKeyParameterNamesContractPtrInput
	// Specifies whether an API or Product subscription is required for accessing the API.
	SubscriptionRequired pulumi.BoolPtrInput
	//  A URL to the Terms of Service for the API. MUST be in the format of a URL.
	TermsOfServiceUrl pulumi.StringPtrInput
	// Strategy of translating required query parameters to template ones. By default has value 'template'. Possible values: 'template', 'query'
	TranslateRequiredQueryParametersConduct pulumi.StringPtrInput
	// Content value when Importing an API.
	Value pulumi.StringPtrInput
	// Criteria to limit import of WSDL to a subset of the document.
	WsdlSelector ApiCreateOrUpdatePropertiesWsdlSelectorPtrInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}

type ApiInput interface {
	pulumi.Input

	ToApiOutput() ApiOutput
	ToApiOutputWithContext(ctx context.Context) ApiOutput
}

func (*Api) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (i *Api) ToApiOutput() ApiOutput {
	return i.ToApiOutputWithContext(context.Background())
}

func (i *Api) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOutput)
}

type ApiOutput struct{ *pulumi.OutputState }

func (ApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (o ApiOutput) ToApiOutput() ApiOutput {
	return o
}

func (o ApiOutput) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return o
}

// Describes the revision of the API. If no value is provided, default revision 1 is created
func (o ApiOutput) ApiRevision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiRevision }).(pulumi.StringPtrOutput)
}

// Description of the API Revision.
func (o ApiOutput) ApiRevisionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiRevisionDescription }).(pulumi.StringPtrOutput)
}

// Type of API.
func (o ApiOutput) ApiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiType }).(pulumi.StringPtrOutput)
}

// Indicates the version identifier of the API if the API is versioned
func (o ApiOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Description of the API Version.
func (o ApiOutput) ApiVersionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiVersionDescription }).(pulumi.StringPtrOutput)
}

// Version set details
func (o ApiOutput) ApiVersionSet() ApiVersionSetContractDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Api) ApiVersionSetContractDetailsResponsePtrOutput { return v.ApiVersionSet }).(ApiVersionSetContractDetailsResponsePtrOutput)
}

// A resource identifier for the related ApiVersionSet.
func (o ApiOutput) ApiVersionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiVersionSetId }).(pulumi.StringPtrOutput)
}

// Collection of authentication settings included into this API.
func (o ApiOutput) AuthenticationSettings() AuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyT(func(v *Api) AuthenticationSettingsContractResponsePtrOutput { return v.AuthenticationSettings }).(AuthenticationSettingsContractResponsePtrOutput)
}

// Contact information for the API.
func (o ApiOutput) Contact() ApiContactInformationResponsePtrOutput {
	return o.ApplyT(func(v *Api) ApiContactInformationResponsePtrOutput { return v.Contact }).(ApiContactInformationResponsePtrOutput)
}

// Description of the API. May include HTML formatting tags.
func (o ApiOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// API name. Must be 1 to 300 characters long.
func (o ApiOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Indicates if API revision is current api revision.
func (o ApiOutput) IsCurrent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.BoolPtrOutput { return v.IsCurrent }).(pulumi.BoolPtrOutput)
}

// Indicates if API revision is accessible via the gateway.
func (o ApiOutput) IsOnline() pulumi.BoolOutput {
	return o.ApplyT(func(v *Api) pulumi.BoolOutput { return v.IsOnline }).(pulumi.BoolOutput)
}

// License information for the API.
func (o ApiOutput) License() ApiLicenseInformationResponsePtrOutput {
	return o.ApplyT(func(v *Api) ApiLicenseInformationResponsePtrOutput { return v.License }).(ApiLicenseInformationResponsePtrOutput)
}

// The name of the resource
func (o ApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
func (o ApiOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Describes on which protocols the operations in this API can be invoked.
func (o ApiOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Api) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
func (o ApiOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

// API identifier of the source API.
func (o ApiOutput) SourceApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.SourceApiId }).(pulumi.StringPtrOutput)
}

// Protocols over which API is made available.
func (o ApiOutput) SubscriptionKeyParameterNames() SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return o.ApplyT(func(v *Api) SubscriptionKeyParameterNamesContractResponsePtrOutput {
		return v.SubscriptionKeyParameterNames
	}).(SubscriptionKeyParameterNamesContractResponsePtrOutput)
}

// Specifies whether an API or Product subscription is required for accessing the API.
func (o ApiOutput) SubscriptionRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.BoolPtrOutput { return v.SubscriptionRequired }).(pulumi.BoolPtrOutput)
}

// A URL to the Terms of Service for the API. MUST be in the format of a URL.
func (o ApiOutput) TermsOfServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.TermsOfServiceUrl }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ApiOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiOutput{})
}
