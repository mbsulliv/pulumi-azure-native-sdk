// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The description of Fhir Service
func LookupFhirService(ctx *pulumi.Context, args *LookupFhirServiceArgs, opts ...pulumi.InvokeOption) (*LookupFhirServiceResult, error) {
	var rv LookupFhirServiceResult
	err := ctx.Invoke("azure-native:healthcareapis/v20221001preview:getFhirService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFhirServiceArgs struct {
	// The name of FHIR Service resource.
	FhirServiceName string `pulumi:"fhirServiceName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The description of Fhir Service
type LookupFhirServiceResult struct {
	// Fhir Service access policies.
	AccessPolicies []FhirServiceAccessPolicyEntryResponse `pulumi:"accessPolicies"`
	// Fhir Service Azure container registry configuration.
	AcrConfiguration *FhirServiceAcrConfigurationResponse `pulumi:"acrConfiguration"`
	// Fhir Service authentication configuration.
	AuthenticationConfiguration *FhirServiceAuthenticationConfigurationResponse `pulumi:"authenticationConfiguration"`
	// Fhir Service Cors configuration.
	CorsConfiguration *FhirServiceCorsConfigurationResponse `pulumi:"corsConfiguration"`
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string `pulumi:"etag"`
	// Fhir Service event support status.
	EventState string `pulumi:"eventState"`
	// Fhir Service export configuration.
	ExportConfiguration *FhirServiceExportConfigurationResponse `pulumi:"exportConfiguration"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityResponseIdentity `pulumi:"identity"`
	// Fhir Service import configuration.
	ImportConfiguration *FhirServiceImportConfigurationResponse `pulumi:"importConfiguration"`
	// The kind of the service.
	Kind *string `pulumi:"kind"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// The list of private endpoint connections that are set up for this resource.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// The provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess string `pulumi:"publicNetworkAccess"`
	// Determines tracking of history for resources.
	ResourceVersionPolicyConfiguration *ResourceVersionPolicyConfigurationResponse `pulumi:"resourceVersionPolicyConfiguration"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupFhirServiceOutput(ctx *pulumi.Context, args LookupFhirServiceOutputArgs, opts ...pulumi.InvokeOption) LookupFhirServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFhirServiceResult, error) {
			args := v.(LookupFhirServiceArgs)
			r, err := LookupFhirService(ctx, &args, opts...)
			var s LookupFhirServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFhirServiceResultOutput)
}

type LookupFhirServiceOutputArgs struct {
	// The name of FHIR Service resource.
	FhirServiceName pulumi.StringInput `pulumi:"fhirServiceName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupFhirServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFhirServiceArgs)(nil)).Elem()
}

// The description of Fhir Service
type LookupFhirServiceResultOutput struct{ *pulumi.OutputState }

func (LookupFhirServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFhirServiceResult)(nil)).Elem()
}

func (o LookupFhirServiceResultOutput) ToLookupFhirServiceResultOutput() LookupFhirServiceResultOutput {
	return o
}

func (o LookupFhirServiceResultOutput) ToLookupFhirServiceResultOutputWithContext(ctx context.Context) LookupFhirServiceResultOutput {
	return o
}

// Fhir Service access policies.
func (o LookupFhirServiceResultOutput) AccessPolicies() FhirServiceAccessPolicyEntryResponseArrayOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) []FhirServiceAccessPolicyEntryResponse { return v.AccessPolicies }).(FhirServiceAccessPolicyEntryResponseArrayOutput)
}

// Fhir Service Azure container registry configuration.
func (o LookupFhirServiceResultOutput) AcrConfiguration() FhirServiceAcrConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceAcrConfigurationResponse { return v.AcrConfiguration }).(FhirServiceAcrConfigurationResponsePtrOutput)
}

// Fhir Service authentication configuration.
func (o LookupFhirServiceResultOutput) AuthenticationConfiguration() FhirServiceAuthenticationConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceAuthenticationConfigurationResponse {
		return v.AuthenticationConfiguration
	}).(FhirServiceAuthenticationConfigurationResponsePtrOutput)
}

// Fhir Service Cors configuration.
func (o LookupFhirServiceResultOutput) CorsConfiguration() FhirServiceCorsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceCorsConfigurationResponse { return v.CorsConfiguration }).(FhirServiceCorsConfigurationResponsePtrOutput)
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o LookupFhirServiceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fhir Service event support status.
func (o LookupFhirServiceResultOutput) EventState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.EventState }).(pulumi.StringOutput)
}

// Fhir Service export configuration.
func (o LookupFhirServiceResultOutput) ExportConfiguration() FhirServiceExportConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceExportConfigurationResponse { return v.ExportConfiguration }).(FhirServiceExportConfigurationResponsePtrOutput)
}

// The resource identifier.
func (o LookupFhirServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Setting indicating whether the service has a managed identity associated with it.
func (o LookupFhirServiceResultOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *ServiceManagedIdentityResponseIdentity { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

// Fhir Service import configuration.
func (o LookupFhirServiceResultOutput) ImportConfiguration() FhirServiceImportConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceImportConfigurationResponse { return v.ImportConfiguration }).(FhirServiceImportConfigurationResponsePtrOutput)
}

// The kind of the service.
func (o LookupFhirServiceResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The resource location.
func (o LookupFhirServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o LookupFhirServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of private endpoint connections that are set up for this resource.
func (o LookupFhirServiceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// The provisioning state.
func (o LookupFhirServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
func (o LookupFhirServiceResultOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

// Determines tracking of history for resources.
func (o LookupFhirServiceResultOutput) ResourceVersionPolicyConfiguration() ResourceVersionPolicyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *ResourceVersionPolicyConfigurationResponse {
		return v.ResourceVersionPolicyConfiguration
	}).(ResourceVersionPolicyConfigurationResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupFhirServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupFhirServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupFhirServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFhirServiceResultOutput{})
}
