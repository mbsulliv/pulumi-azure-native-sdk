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
// Azure REST API version: 2022-09-01-preview.
//
// Other available API versions: 2023-03-01-preview.
type WorkspaceApi struct {
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

// NewWorkspaceApi registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceApi(ctx *pulumi.Context,
	name string, args *WorkspaceApiArgs, opts ...pulumi.ResourceOption) (*WorkspaceApi, error) {
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
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:WorkspaceApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:WorkspaceApi"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceApi
	err := ctx.RegisterResource("azure-native:apimanagement:WorkspaceApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceApi gets an existing WorkspaceApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceApiState, opts ...pulumi.ResourceOption) (*WorkspaceApi, error) {
	var resource WorkspaceApi
	err := ctx.ReadResource("azure-native:apimanagement:WorkspaceApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceApi resources.
type workspaceApiState struct {
}

type WorkspaceApiState struct {
}

func (WorkspaceApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceApiState)(nil)).Elem()
}

type workspaceApiArgs struct {
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
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
	// Criteria to limit import of WSDL to a subset of the document.
	WsdlSelector *ApiCreateOrUpdatePropertiesWsdlSelector `pulumi:"wsdlSelector"`
}

// The set of arguments for constructing a WorkspaceApi resource.
type WorkspaceApiArgs struct {
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
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput
	// Criteria to limit import of WSDL to a subset of the document.
	WsdlSelector ApiCreateOrUpdatePropertiesWsdlSelectorPtrInput
}

func (WorkspaceApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceApiArgs)(nil)).Elem()
}

type WorkspaceApiInput interface {
	pulumi.Input

	ToWorkspaceApiOutput() WorkspaceApiOutput
	ToWorkspaceApiOutputWithContext(ctx context.Context) WorkspaceApiOutput
}

func (*WorkspaceApi) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceApi)(nil)).Elem()
}

func (i *WorkspaceApi) ToWorkspaceApiOutput() WorkspaceApiOutput {
	return i.ToWorkspaceApiOutputWithContext(context.Background())
}

func (i *WorkspaceApi) ToWorkspaceApiOutputWithContext(ctx context.Context) WorkspaceApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceApiOutput)
}

type WorkspaceApiOutput struct{ *pulumi.OutputState }

func (WorkspaceApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceApi)(nil)).Elem()
}

func (o WorkspaceApiOutput) ToWorkspaceApiOutput() WorkspaceApiOutput {
	return o
}

func (o WorkspaceApiOutput) ToWorkspaceApiOutputWithContext(ctx context.Context) WorkspaceApiOutput {
	return o
}

// Describes the revision of the API. If no value is provided, default revision 1 is created
func (o WorkspaceApiOutput) ApiRevision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringPtrOutput { return v.ApiRevision }).(pulumi.StringPtrOutput)
}

// Description of the API Revision.
func (o WorkspaceApiOutput) ApiRevisionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringPtrOutput { return v.ApiRevisionDescription }).(pulumi.StringPtrOutput)
}

// Type of API.
func (o WorkspaceApiOutput) ApiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringPtrOutput { return v.ApiType }).(pulumi.StringPtrOutput)
}

// Indicates the version identifier of the API if the API is versioned
func (o WorkspaceApiOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Description of the API Version.
func (o WorkspaceApiOutput) ApiVersionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringPtrOutput { return v.ApiVersionDescription }).(pulumi.StringPtrOutput)
}

// Version set details
func (o WorkspaceApiOutput) ApiVersionSet() ApiVersionSetContractDetailsResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) ApiVersionSetContractDetailsResponsePtrOutput { return v.ApiVersionSet }).(ApiVersionSetContractDetailsResponsePtrOutput)
}

// A resource identifier for the related ApiVersionSet.
func (o WorkspaceApiOutput) ApiVersionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringPtrOutput { return v.ApiVersionSetId }).(pulumi.StringPtrOutput)
}

// Collection of authentication settings included into this API.
func (o WorkspaceApiOutput) AuthenticationSettings() AuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) AuthenticationSettingsContractResponsePtrOutput { return v.AuthenticationSettings }).(AuthenticationSettingsContractResponsePtrOutput)
}

// Contact information for the API.
func (o WorkspaceApiOutput) Contact() ApiContactInformationResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) ApiContactInformationResponsePtrOutput { return v.Contact }).(ApiContactInformationResponsePtrOutput)
}

// Description of the API. May include HTML formatting tags.
func (o WorkspaceApiOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// API name. Must be 1 to 300 characters long.
func (o WorkspaceApiOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Indicates if API revision is current api revision.
func (o WorkspaceApiOutput) IsCurrent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.BoolPtrOutput { return v.IsCurrent }).(pulumi.BoolPtrOutput)
}

// Indicates if API revision is accessible via the gateway.
func (o WorkspaceApiOutput) IsOnline() pulumi.BoolOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.BoolOutput { return v.IsOnline }).(pulumi.BoolOutput)
}

// License information for the API.
func (o WorkspaceApiOutput) License() ApiLicenseInformationResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) ApiLicenseInformationResponsePtrOutput { return v.License }).(ApiLicenseInformationResponsePtrOutput)
}

// The name of the resource
func (o WorkspaceApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
func (o WorkspaceApiOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Describes on which protocols the operations in this API can be invoked.
func (o WorkspaceApiOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
func (o WorkspaceApiOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringPtrOutput { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

// API identifier of the source API.
func (o WorkspaceApiOutput) SourceApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringPtrOutput { return v.SourceApiId }).(pulumi.StringPtrOutput)
}

// Protocols over which API is made available.
func (o WorkspaceApiOutput) SubscriptionKeyParameterNames() SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) SubscriptionKeyParameterNamesContractResponsePtrOutput {
		return v.SubscriptionKeyParameterNames
	}).(SubscriptionKeyParameterNamesContractResponsePtrOutput)
}

// Specifies whether an API or Product subscription is required for accessing the API.
func (o WorkspaceApiOutput) SubscriptionRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.BoolPtrOutput { return v.SubscriptionRequired }).(pulumi.BoolPtrOutput)
}

// A URL to the Terms of Service for the API. MUST be in the format of a URL.
func (o WorkspaceApiOutput) TermsOfServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringPtrOutput { return v.TermsOfServiceUrl }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceApiOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceApi) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceApiOutput{})
}
