// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package healthcareapis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of the specified DICOM Service.
// Azure REST API version: 2023-02-28.
//
// Other available API versions: 2023-09-06.
func LookupDicomService(ctx *pulumi.Context, args *LookupDicomServiceArgs, opts ...pulumi.InvokeOption) (*LookupDicomServiceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDicomServiceResult
	err := ctx.Invoke("azure-native:healthcareapis:getDicomService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDicomServiceArgs struct {
	// The name of DICOM Service resource.
	DicomServiceName string `pulumi:"dicomServiceName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The description of Dicom Service
type LookupDicomServiceResult struct {
	// Dicom Service authentication configuration.
	AuthenticationConfiguration *DicomServiceAuthenticationConfigurationResponse `pulumi:"authenticationConfiguration"`
	// Dicom Service Cors configuration.
	CorsConfiguration *CorsConfigurationResponse `pulumi:"corsConfiguration"`
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string `pulumi:"etag"`
	// DICOM Service event support status.
	EventState string `pulumi:"eventState"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityResponseIdentity `pulumi:"identity"`
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
	// The url of the Dicom Services.
	ServiceUrl string `pulumi:"serviceUrl"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupDicomServiceOutput(ctx *pulumi.Context, args LookupDicomServiceOutputArgs, opts ...pulumi.InvokeOption) LookupDicomServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDicomServiceResult, error) {
			args := v.(LookupDicomServiceArgs)
			r, err := LookupDicomService(ctx, &args, opts...)
			var s LookupDicomServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDicomServiceResultOutput)
}

type LookupDicomServiceOutputArgs struct {
	// The name of DICOM Service resource.
	DicomServiceName pulumi.StringInput `pulumi:"dicomServiceName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupDicomServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDicomServiceArgs)(nil)).Elem()
}

// The description of Dicom Service
type LookupDicomServiceResultOutput struct{ *pulumi.OutputState }

func (LookupDicomServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDicomServiceResult)(nil)).Elem()
}

func (o LookupDicomServiceResultOutput) ToLookupDicomServiceResultOutput() LookupDicomServiceResultOutput {
	return o
}

func (o LookupDicomServiceResultOutput) ToLookupDicomServiceResultOutputWithContext(ctx context.Context) LookupDicomServiceResultOutput {
	return o
}

// Dicom Service authentication configuration.
func (o LookupDicomServiceResultOutput) AuthenticationConfiguration() DicomServiceAuthenticationConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) *DicomServiceAuthenticationConfigurationResponse {
		return v.AuthenticationConfiguration
	}).(DicomServiceAuthenticationConfigurationResponsePtrOutput)
}

// Dicom Service Cors configuration.
func (o LookupDicomServiceResultOutput) CorsConfiguration() CorsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) *CorsConfigurationResponse { return v.CorsConfiguration }).(CorsConfigurationResponsePtrOutput)
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o LookupDicomServiceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// DICOM Service event support status.
func (o LookupDicomServiceResultOutput) EventState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) string { return v.EventState }).(pulumi.StringOutput)
}

// The resource identifier.
func (o LookupDicomServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Setting indicating whether the service has a managed identity associated with it.
func (o LookupDicomServiceResultOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) *ServiceManagedIdentityResponseIdentity { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

// The resource location.
func (o LookupDicomServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o LookupDicomServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of private endpoint connections that are set up for this resource.
func (o LookupDicomServiceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// The provisioning state.
func (o LookupDicomServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
func (o LookupDicomServiceResultOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) string { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

// The url of the Dicom Services.
func (o LookupDicomServiceResultOutput) ServiceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) string { return v.ServiceUrl }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupDicomServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupDicomServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupDicomServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDicomServiceResultOutput{})
}
