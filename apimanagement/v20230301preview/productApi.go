// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// API details.
type ProductApi struct {
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

// NewProductApi registers a new resource with the given unique name, arguments, and options.
func NewProductApi(ctx *pulumi.Context,
	name string, args *ProductApiArgs, opts ...pulumi.ResourceOption) (*ProductApi, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:ProductApi"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProductApi
	err := ctx.RegisterResource("azure-native:apimanagement/v20230301preview:ProductApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProductApi gets an existing ProductApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProductApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductApiState, opts ...pulumi.ResourceOption) (*ProductApi, error) {
	var resource ProductApi
	err := ctx.ReadResource("azure-native:apimanagement/v20230301preview:ProductApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProductApi resources.
type productApiState struct {
}

type ProductApiState struct {
}

func (ProductApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*productApiState)(nil)).Elem()
}

type productApiArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId *string `pulumi:"apiId"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ProductApi resource.
type ProductApiArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringPtrInput
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (ProductApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productApiArgs)(nil)).Elem()
}

type ProductApiInput interface {
	pulumi.Input

	ToProductApiOutput() ProductApiOutput
	ToProductApiOutputWithContext(ctx context.Context) ProductApiOutput
}

func (*ProductApi) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductApi)(nil)).Elem()
}

func (i *ProductApi) ToProductApiOutput() ProductApiOutput {
	return i.ToProductApiOutputWithContext(context.Background())
}

func (i *ProductApi) ToProductApiOutputWithContext(ctx context.Context) ProductApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductApiOutput)
}

func (i *ProductApi) ToOutput(ctx context.Context) pulumix.Output[*ProductApi] {
	return pulumix.Output[*ProductApi]{
		OutputState: i.ToProductApiOutputWithContext(ctx).OutputState,
	}
}

type ProductApiOutput struct{ *pulumi.OutputState }

func (ProductApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductApi)(nil)).Elem()
}

func (o ProductApiOutput) ToProductApiOutput() ProductApiOutput {
	return o
}

func (o ProductApiOutput) ToProductApiOutputWithContext(ctx context.Context) ProductApiOutput {
	return o
}

func (o ProductApiOutput) ToOutput(ctx context.Context) pulumix.Output[*ProductApi] {
	return pulumix.Output[*ProductApi]{
		OutputState: o.OutputState,
	}
}

// Describes the revision of the API. If no value is provided, default revision 1 is created
func (o ProductApiOutput) ApiRevision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.ApiRevision }).(pulumi.StringPtrOutput)
}

// Description of the API Revision.
func (o ProductApiOutput) ApiRevisionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.ApiRevisionDescription }).(pulumi.StringPtrOutput)
}

// Type of API.
func (o ProductApiOutput) ApiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.ApiType }).(pulumi.StringPtrOutput)
}

// Indicates the version identifier of the API if the API is versioned
func (o ProductApiOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Description of the API Version.
func (o ProductApiOutput) ApiVersionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.ApiVersionDescription }).(pulumi.StringPtrOutput)
}

// Version set details
func (o ProductApiOutput) ApiVersionSet() ApiVersionSetContractDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ProductApi) ApiVersionSetContractDetailsResponsePtrOutput { return v.ApiVersionSet }).(ApiVersionSetContractDetailsResponsePtrOutput)
}

// A resource identifier for the related ApiVersionSet.
func (o ProductApiOutput) ApiVersionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.ApiVersionSetId }).(pulumi.StringPtrOutput)
}

// Collection of authentication settings included into this API.
func (o ProductApiOutput) AuthenticationSettings() AuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyT(func(v *ProductApi) AuthenticationSettingsContractResponsePtrOutput { return v.AuthenticationSettings }).(AuthenticationSettingsContractResponsePtrOutput)
}

// Contact information for the API.
func (o ProductApiOutput) Contact() ApiContactInformationResponsePtrOutput {
	return o.ApplyT(func(v *ProductApi) ApiContactInformationResponsePtrOutput { return v.Contact }).(ApiContactInformationResponsePtrOutput)
}

// Description of the API. May include HTML formatting tags.
func (o ProductApiOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// API name. Must be 1 to 300 characters long.
func (o ProductApiOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Indicates if API revision is current api revision.
func (o ProductApiOutput) IsCurrent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.BoolPtrOutput { return v.IsCurrent }).(pulumi.BoolPtrOutput)
}

// Indicates if API revision is accessible via the gateway.
func (o ProductApiOutput) IsOnline() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.BoolOutput { return v.IsOnline }).(pulumi.BoolOutput)
}

// License information for the API.
func (o ProductApiOutput) License() ApiLicenseInformationResponsePtrOutput {
	return o.ApplyT(func(v *ProductApi) ApiLicenseInformationResponsePtrOutput { return v.License }).(ApiLicenseInformationResponsePtrOutput)
}

// The name of the resource
func (o ProductApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
func (o ProductApiOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Describes on which protocols the operations in this API can be invoked.
func (o ProductApiOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
func (o ProductApiOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

// API identifier of the source API.
func (o ProductApiOutput) SourceApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.SourceApiId }).(pulumi.StringPtrOutput)
}

// Protocols over which API is made available.
func (o ProductApiOutput) SubscriptionKeyParameterNames() SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return o.ApplyT(func(v *ProductApi) SubscriptionKeyParameterNamesContractResponsePtrOutput {
		return v.SubscriptionKeyParameterNames
	}).(SubscriptionKeyParameterNamesContractResponsePtrOutput)
}

// Specifies whether an API or Product subscription is required for accessing the API.
func (o ProductApiOutput) SubscriptionRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.BoolPtrOutput { return v.SubscriptionRequired }).(pulumi.BoolPtrOutput)
}

// A URL to the Terms of Service for the API. MUST be in the format of a URL.
func (o ProductApiOutput) TermsOfServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.TermsOfServiceUrl }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ProductApiOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProductApiOutput{})
}
