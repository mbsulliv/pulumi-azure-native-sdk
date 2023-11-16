// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An API collection as represented by Microsoft Defender for APIs.
// Azure REST API version: 2023-11-15.
type APICollectionByAzureApiManagementService struct {
	pulumi.CustomResourceState

	// The base URI for this API collection. All endpoints of this API collection extend this base URI.
	BaseUrl pulumi.StringOutput `pulumi:"baseUrl"`
	// The resource Id of the resource from where this API collection was discovered.
	DiscoveredVia pulumi.StringOutput `pulumi:"discoveredVia"`
	// The display name of the API collection.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of API endpoints discovered in this API collection.
	NumberOfApiEndpoints pulumi.Float64Output `pulumi:"numberOfApiEndpoints"`
	// The number of API endpoints in this API collection which are exposing sensitive data in their requests and/or responses.
	NumberOfApiEndpointsWithSensitiveDataExposed pulumi.Float64Output `pulumi:"numberOfApiEndpointsWithSensitiveDataExposed"`
	// The number of API endpoints in this API collection for which API traffic from the internet was observed.
	NumberOfExternalApiEndpoints pulumi.Float64Output `pulumi:"numberOfExternalApiEndpoints"`
	// The number of API endpoints in this API collection that have not received any API traffic in the last 30 days.
	NumberOfInactiveApiEndpoints pulumi.Float64Output `pulumi:"numberOfInactiveApiEndpoints"`
	// The number of API endpoints in this API collection that are unauthenticated.
	NumberOfUnauthenticatedApiEndpoints pulumi.Float64Output `pulumi:"numberOfUnauthenticatedApiEndpoints"`
	// Gets the provisioning state of the API collection.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The highest priority sensitivity label from Microsoft Purview in this API collection.
	SensitivityLabel pulumi.StringOutput `pulumi:"sensitivityLabel"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAPICollectionByAzureApiManagementService registers a new resource with the given unique name, arguments, and options.
func NewAPICollectionByAzureApiManagementService(ctx *pulumi.Context,
	name string, args *APICollectionByAzureApiManagementServiceArgs, opts ...pulumi.ResourceOption) (*APICollectionByAzureApiManagementService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security/v20221120preview:APICollectionByAzureApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:security/v20231115:APICollectionByAzureApiManagementService"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource APICollectionByAzureApiManagementService
	err := ctx.RegisterResource("azure-native:security:APICollectionByAzureApiManagementService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAPICollectionByAzureApiManagementService gets an existing APICollectionByAzureApiManagementService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAPICollectionByAzureApiManagementService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *APICollectionByAzureApiManagementServiceState, opts ...pulumi.ResourceOption) (*APICollectionByAzureApiManagementService, error) {
	var resource APICollectionByAzureApiManagementService
	err := ctx.ReadResource("azure-native:security:APICollectionByAzureApiManagementService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering APICollectionByAzureApiManagementService resources.
type apicollectionByAzureApiManagementServiceState struct {
}

type APICollectionByAzureApiManagementServiceState struct {
}

func (APICollectionByAzureApiManagementServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*apicollectionByAzureApiManagementServiceState)(nil)).Elem()
}

type apicollectionByAzureApiManagementServiceArgs struct {
	// API revision identifier. Must be unique in the API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId *string `pulumi:"apiId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a APICollectionByAzureApiManagementService resource.
type APICollectionByAzureApiManagementServiceArgs struct {
	// API revision identifier. Must be unique in the API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (APICollectionByAzureApiManagementServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apicollectionByAzureApiManagementServiceArgs)(nil)).Elem()
}

type APICollectionByAzureApiManagementServiceInput interface {
	pulumi.Input

	ToAPICollectionByAzureApiManagementServiceOutput() APICollectionByAzureApiManagementServiceOutput
	ToAPICollectionByAzureApiManagementServiceOutputWithContext(ctx context.Context) APICollectionByAzureApiManagementServiceOutput
}

func (*APICollectionByAzureApiManagementService) ElementType() reflect.Type {
	return reflect.TypeOf((**APICollectionByAzureApiManagementService)(nil)).Elem()
}

func (i *APICollectionByAzureApiManagementService) ToAPICollectionByAzureApiManagementServiceOutput() APICollectionByAzureApiManagementServiceOutput {
	return i.ToAPICollectionByAzureApiManagementServiceOutputWithContext(context.Background())
}

func (i *APICollectionByAzureApiManagementService) ToAPICollectionByAzureApiManagementServiceOutputWithContext(ctx context.Context) APICollectionByAzureApiManagementServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APICollectionByAzureApiManagementServiceOutput)
}

type APICollectionByAzureApiManagementServiceOutput struct{ *pulumi.OutputState }

func (APICollectionByAzureApiManagementServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**APICollectionByAzureApiManagementService)(nil)).Elem()
}

func (o APICollectionByAzureApiManagementServiceOutput) ToAPICollectionByAzureApiManagementServiceOutput() APICollectionByAzureApiManagementServiceOutput {
	return o
}

func (o APICollectionByAzureApiManagementServiceOutput) ToAPICollectionByAzureApiManagementServiceOutputWithContext(ctx context.Context) APICollectionByAzureApiManagementServiceOutput {
	return o
}

// The base URI for this API collection. All endpoints of this API collection extend this base URI.
func (o APICollectionByAzureApiManagementServiceOutput) BaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *APICollectionByAzureApiManagementService) pulumi.StringOutput { return v.BaseUrl }).(pulumi.StringOutput)
}

// The resource Id of the resource from where this API collection was discovered.
func (o APICollectionByAzureApiManagementServiceOutput) DiscoveredVia() pulumi.StringOutput {
	return o.ApplyT(func(v *APICollectionByAzureApiManagementService) pulumi.StringOutput { return v.DiscoveredVia }).(pulumi.StringOutput)
}

// The display name of the API collection.
func (o APICollectionByAzureApiManagementServiceOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *APICollectionByAzureApiManagementService) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource name
func (o APICollectionByAzureApiManagementServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *APICollectionByAzureApiManagementService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of API endpoints discovered in this API collection.
func (o APICollectionByAzureApiManagementServiceOutput) NumberOfApiEndpoints() pulumi.Float64Output {
	return o.ApplyT(func(v *APICollectionByAzureApiManagementService) pulumi.Float64Output { return v.NumberOfApiEndpoints }).(pulumi.Float64Output)
}

// The number of API endpoints in this API collection which are exposing sensitive data in their requests and/or responses.
func (o APICollectionByAzureApiManagementServiceOutput) NumberOfApiEndpointsWithSensitiveDataExposed() pulumi.Float64Output {
	return o.ApplyT(func(v *APICollectionByAzureApiManagementService) pulumi.Float64Output {
		return v.NumberOfApiEndpointsWithSensitiveDataExposed
	}).(pulumi.Float64Output)
}

// The number of API endpoints in this API collection for which API traffic from the internet was observed.
func (o APICollectionByAzureApiManagementServiceOutput) NumberOfExternalApiEndpoints() pulumi.Float64Output {
	return o.ApplyT(func(v *APICollectionByAzureApiManagementService) pulumi.Float64Output {
		return v.NumberOfExternalApiEndpoints
	}).(pulumi.Float64Output)
}

// The number of API endpoints in this API collection that have not received any API traffic in the last 30 days.
func (o APICollectionByAzureApiManagementServiceOutput) NumberOfInactiveApiEndpoints() pulumi.Float64Output {
	return o.ApplyT(func(v *APICollectionByAzureApiManagementService) pulumi.Float64Output {
		return v.NumberOfInactiveApiEndpoints
	}).(pulumi.Float64Output)
}

// The number of API endpoints in this API collection that are unauthenticated.
func (o APICollectionByAzureApiManagementServiceOutput) NumberOfUnauthenticatedApiEndpoints() pulumi.Float64Output {
	return o.ApplyT(func(v *APICollectionByAzureApiManagementService) pulumi.Float64Output {
		return v.NumberOfUnauthenticatedApiEndpoints
	}).(pulumi.Float64Output)
}

// Gets the provisioning state of the API collection.
func (o APICollectionByAzureApiManagementServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *APICollectionByAzureApiManagementService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The highest priority sensitivity label from Microsoft Purview in this API collection.
func (o APICollectionByAzureApiManagementServiceOutput) SensitivityLabel() pulumi.StringOutput {
	return o.ApplyT(func(v *APICollectionByAzureApiManagementService) pulumi.StringOutput { return v.SensitivityLabel }).(pulumi.StringOutput)
}

// Resource type
func (o APICollectionByAzureApiManagementServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *APICollectionByAzureApiManagementService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(APICollectionByAzureApiManagementServiceOutput{})
}
