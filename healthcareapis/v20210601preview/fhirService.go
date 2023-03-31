// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The description of Fhir Service
type FhirService struct {
	pulumi.CustomResourceState

	// Fhir Service access policies.
	AccessPolicies FhirServiceAccessPolicyEntryResponseArrayOutput `pulumi:"accessPolicies"`
	// Fhir Service Azure container registry configuration.
	AcrConfiguration FhirServiceAcrConfigurationResponsePtrOutput `pulumi:"acrConfiguration"`
	// Fhir Service authentication configuration.
	AuthenticationConfiguration FhirServiceAuthenticationConfigurationResponsePtrOutput `pulumi:"authenticationConfiguration"`
	// Fhir Service Cors configuration.
	CorsConfiguration FhirServiceCorsConfigurationResponsePtrOutput `pulumi:"corsConfiguration"`
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Fhir Service export configuration.
	ExportConfiguration FhirServiceExportConfigurationResponsePtrOutput `pulumi:"exportConfiguration"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ServiceManagedIdentityResponseIdentityPtrOutput `pulumi:"identity"`
	// The kind of the service.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFhirService registers a new resource with the given unique name, arguments, and options.
func NewFhirService(ctx *pulumi.Context,
	name string, args *FhirServiceArgs, opts ...pulumi.ResourceOption) (*FhirService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:healthcareapis:FhirService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20211101:FhirService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220131preview:FhirService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220515:FhirService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220601:FhirService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20221001preview:FhirService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20221201:FhirService"),
		},
	})
	opts = append(opts, aliases)
	var resource FhirService
	err := ctx.RegisterResource("azure-native:healthcareapis/v20210601preview:FhirService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFhirService gets an existing FhirService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFhirService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FhirServiceState, opts ...pulumi.ResourceOption) (*FhirService, error) {
	var resource FhirService
	err := ctx.ReadResource("azure-native:healthcareapis/v20210601preview:FhirService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FhirService resources.
type fhirServiceState struct {
}

type FhirServiceState struct {
}

func (FhirServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirServiceState)(nil)).Elem()
}

type fhirServiceArgs struct {
	// Fhir Service access policies.
	AccessPolicies []FhirServiceAccessPolicyEntry `pulumi:"accessPolicies"`
	// Fhir Service Azure container registry configuration.
	AcrConfiguration *FhirServiceAcrConfiguration `pulumi:"acrConfiguration"`
	// Fhir Service authentication configuration.
	AuthenticationConfiguration *FhirServiceAuthenticationConfiguration `pulumi:"authenticationConfiguration"`
	// Fhir Service Cors configuration.
	CorsConfiguration *FhirServiceCorsConfiguration `pulumi:"corsConfiguration"`
	// Fhir Service export configuration.
	ExportConfiguration *FhirServiceExportConfiguration `pulumi:"exportConfiguration"`
	// The name of FHIR Service resource.
	FhirServiceName *string `pulumi:"fhirServiceName"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityIdentity `pulumi:"identity"`
	// The kind of the service.
	Kind *string `pulumi:"kind"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of workspace resource.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a FhirService resource.
type FhirServiceArgs struct {
	// Fhir Service access policies.
	AccessPolicies FhirServiceAccessPolicyEntryArrayInput
	// Fhir Service Azure container registry configuration.
	AcrConfiguration FhirServiceAcrConfigurationPtrInput
	// Fhir Service authentication configuration.
	AuthenticationConfiguration FhirServiceAuthenticationConfigurationPtrInput
	// Fhir Service Cors configuration.
	CorsConfiguration FhirServiceCorsConfigurationPtrInput
	// Fhir Service export configuration.
	ExportConfiguration FhirServiceExportConfigurationPtrInput
	// The name of FHIR Service resource.
	FhirServiceName pulumi.StringPtrInput
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ServiceManagedIdentityIdentityPtrInput
	// The kind of the service.
	Kind pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of workspace resource.
	WorkspaceName pulumi.StringInput
}

func (FhirServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirServiceArgs)(nil)).Elem()
}

type FhirServiceInput interface {
	pulumi.Input

	ToFhirServiceOutput() FhirServiceOutput
	ToFhirServiceOutputWithContext(ctx context.Context) FhirServiceOutput
}

func (*FhirService) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirService)(nil)).Elem()
}

func (i *FhirService) ToFhirServiceOutput() FhirServiceOutput {
	return i.ToFhirServiceOutputWithContext(context.Background())
}

func (i *FhirService) ToFhirServiceOutputWithContext(ctx context.Context) FhirServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceOutput)
}

type FhirServiceOutput struct{ *pulumi.OutputState }

func (FhirServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirService)(nil)).Elem()
}

func (o FhirServiceOutput) ToFhirServiceOutput() FhirServiceOutput {
	return o
}

func (o FhirServiceOutput) ToFhirServiceOutputWithContext(ctx context.Context) FhirServiceOutput {
	return o
}

// Fhir Service access policies.
func (o FhirServiceOutput) AccessPolicies() FhirServiceAccessPolicyEntryResponseArrayOutput {
	return o.ApplyT(func(v *FhirService) FhirServiceAccessPolicyEntryResponseArrayOutput { return v.AccessPolicies }).(FhirServiceAccessPolicyEntryResponseArrayOutput)
}

// Fhir Service Azure container registry configuration.
func (o FhirServiceOutput) AcrConfiguration() FhirServiceAcrConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *FhirService) FhirServiceAcrConfigurationResponsePtrOutput { return v.AcrConfiguration }).(FhirServiceAcrConfigurationResponsePtrOutput)
}

// Fhir Service authentication configuration.
func (o FhirServiceOutput) AuthenticationConfiguration() FhirServiceAuthenticationConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *FhirService) FhirServiceAuthenticationConfigurationResponsePtrOutput {
		return v.AuthenticationConfiguration
	}).(FhirServiceAuthenticationConfigurationResponsePtrOutput)
}

// Fhir Service Cors configuration.
func (o FhirServiceOutput) CorsConfiguration() FhirServiceCorsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *FhirService) FhirServiceCorsConfigurationResponsePtrOutput { return v.CorsConfiguration }).(FhirServiceCorsConfigurationResponsePtrOutput)
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o FhirServiceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fhir Service export configuration.
func (o FhirServiceOutput) ExportConfiguration() FhirServiceExportConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *FhirService) FhirServiceExportConfigurationResponsePtrOutput { return v.ExportConfiguration }).(FhirServiceExportConfigurationResponsePtrOutput)
}

// Setting indicating whether the service has a managed identity associated with it.
func (o FhirServiceOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v *FhirService) ServiceManagedIdentityResponseIdentityPtrOutput { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

// The kind of the service.
func (o FhirServiceOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The resource location.
func (o FhirServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o FhirServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state.
func (o FhirServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o FhirServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FhirService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o FhirServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o FhirServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FhirServiceOutput{})
}
