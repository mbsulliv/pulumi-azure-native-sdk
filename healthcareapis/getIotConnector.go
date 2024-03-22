// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package healthcareapis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of the specified IoT Connector.
// Azure REST API version: 2023-02-28.
//
// Other available API versions: 2023-09-06, 2023-11-01, 2023-12-01, 2024-03-01, 2024-03-31.
func LookupIotConnector(ctx *pulumi.Context, args *LookupIotConnectorArgs, opts ...pulumi.InvokeOption) (*LookupIotConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIotConnectorResult
	err := ctx.Invoke("azure-native:healthcareapis:getIotConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotConnectorArgs struct {
	// The name of IoT Connector resource.
	IotConnectorName string `pulumi:"iotConnectorName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName string `pulumi:"workspaceName"`
}

// IoT Connector definition.
type LookupIotConnectorResult struct {
	// Device Mappings.
	DeviceMapping *IotMappingPropertiesResponse `pulumi:"deviceMapping"`
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string `pulumi:"etag"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityResponseIdentity `pulumi:"identity"`
	// Source configuration.
	IngestionEndpointConfiguration *IotEventHubIngestionEndpointConfigurationResponse `pulumi:"ingestionEndpointConfiguration"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// The provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupIotConnectorOutput(ctx *pulumi.Context, args LookupIotConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupIotConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotConnectorResult, error) {
			args := v.(LookupIotConnectorArgs)
			r, err := LookupIotConnector(ctx, &args, opts...)
			var s LookupIotConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotConnectorResultOutput)
}

type LookupIotConnectorOutputArgs struct {
	// The name of IoT Connector resource.
	IotConnectorName pulumi.StringInput `pulumi:"iotConnectorName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIotConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotConnectorArgs)(nil)).Elem()
}

// IoT Connector definition.
type LookupIotConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupIotConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotConnectorResult)(nil)).Elem()
}

func (o LookupIotConnectorResultOutput) ToLookupIotConnectorResultOutput() LookupIotConnectorResultOutput {
	return o
}

func (o LookupIotConnectorResultOutput) ToLookupIotConnectorResultOutputWithContext(ctx context.Context) LookupIotConnectorResultOutput {
	return o
}

// Device Mappings.
func (o LookupIotConnectorResultOutput) DeviceMapping() IotMappingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) *IotMappingPropertiesResponse { return v.DeviceMapping }).(IotMappingPropertiesResponsePtrOutput)
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o LookupIotConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The resource identifier.
func (o LookupIotConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Setting indicating whether the service has a managed identity associated with it.
func (o LookupIotConnectorResultOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) *ServiceManagedIdentityResponseIdentity { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

// Source configuration.
func (o LookupIotConnectorResultOutput) IngestionEndpointConfiguration() IotEventHubIngestionEndpointConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) *IotEventHubIngestionEndpointConfigurationResponse {
		return v.IngestionEndpointConfiguration
	}).(IotEventHubIngestionEndpointConfigurationResponsePtrOutput)
}

// The resource location.
func (o LookupIotConnectorResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o LookupIotConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state.
func (o LookupIotConnectorResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupIotConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupIotConnectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupIotConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotConnectorResultOutput{})
}
