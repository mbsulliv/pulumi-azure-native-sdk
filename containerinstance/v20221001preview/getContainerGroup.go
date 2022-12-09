// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A container group.
func LookupContainerGroup(ctx *pulumi.Context, args *LookupContainerGroupArgs, opts ...pulumi.InvokeOption) (*LookupContainerGroupResult, error) {
	var rv LookupContainerGroupResult
	err := ctx.Invoke("azure-native:containerinstance/v20221001preview:getContainerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupContainerGroupArgs struct {
	// The name of the container group.
	ContainerGroupName string `pulumi:"containerGroupName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A container group.
type LookupContainerGroupResult struct {
	// The properties for confidential container group
	ConfidentialComputeProperties *ConfidentialComputePropertiesResponse `pulumi:"confidentialComputeProperties"`
	// The containers within the container group.
	Containers []ContainerResponse `pulumi:"containers"`
	// The diagnostic information for a container group.
	Diagnostics *ContainerGroupDiagnosticsResponse `pulumi:"diagnostics"`
	// The DNS config information for a container group.
	DnsConfig *DnsConfigurationResponse `pulumi:"dnsConfig"`
	// The encryption properties for a container group.
	EncryptionProperties *EncryptionPropertiesResponse `pulumi:"encryptionProperties"`
	// extensions used by virtual kubelet
	Extensions []DeploymentExtensionSpecResponse `pulumi:"extensions"`
	// The resource id.
	Id string `pulumi:"id"`
	// The identity of the container group, if configured.
	Identity *ContainerGroupIdentityResponse `pulumi:"identity"`
	// The image registry credentials by which the container group is created from.
	ImageRegistryCredentials []ImageRegistryCredentialResponse `pulumi:"imageRegistryCredentials"`
	// The init containers for a container group.
	InitContainers []InitContainerDefinitionResponse `pulumi:"initContainers"`
	// The instance view of the container group. Only valid in response.
	InstanceView ContainerGroupPropertiesResponseInstanceView `pulumi:"instanceView"`
	// The IP address type of the container group.
	IpAddress *IpAddressResponse `pulumi:"ipAddress"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// The operating system type required by the containers in the container group.
	OsType string `pulumi:"osType"`
	// The priority of the container group.
	Priority *string `pulumi:"priority"`
	// The provisioning state of the container group. This only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// Restart policy for all containers within the container group.
	// - `Always` Always restart
	// - `OnFailure` Restart on failure
	// - `Never` Never restart
	RestartPolicy *string `pulumi:"restartPolicy"`
	// The SKU for a container group.
	Sku *string `pulumi:"sku"`
	// The subnet resource IDs for a container group.
	SubnetIds []ContainerGroupSubnetIdResponse `pulumi:"subnetIds"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
	// The list of volumes that can be mounted by containers in this container group.
	Volumes []VolumeResponse `pulumi:"volumes"`
	// The zones for the container group.
	Zones []string `pulumi:"zones"`
}

// Defaults sets the appropriate defaults for LookupContainerGroupResult
func (val *LookupContainerGroupResult) Defaults() *LookupContainerGroupResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.IpAddress = tmp.IpAddress.Defaults()

	return &tmp
}

func LookupContainerGroupOutput(ctx *pulumi.Context, args LookupContainerGroupOutputArgs, opts ...pulumi.InvokeOption) LookupContainerGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerGroupResult, error) {
			args := v.(LookupContainerGroupArgs)
			r, err := LookupContainerGroup(ctx, &args, opts...)
			var s LookupContainerGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerGroupResultOutput)
}

type LookupContainerGroupOutputArgs struct {
	// The name of the container group.
	ContainerGroupName pulumi.StringInput `pulumi:"containerGroupName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupContainerGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerGroupArgs)(nil)).Elem()
}

// A container group.
type LookupContainerGroupResultOutput struct{ *pulumi.OutputState }

func (LookupContainerGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerGroupResult)(nil)).Elem()
}

func (o LookupContainerGroupResultOutput) ToLookupContainerGroupResultOutput() LookupContainerGroupResultOutput {
	return o
}

func (o LookupContainerGroupResultOutput) ToLookupContainerGroupResultOutputWithContext(ctx context.Context) LookupContainerGroupResultOutput {
	return o
}

// The properties for confidential container group
func (o LookupContainerGroupResultOutput) ConfidentialComputeProperties() ConfidentialComputePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *ConfidentialComputePropertiesResponse {
		return v.ConfidentialComputeProperties
	}).(ConfidentialComputePropertiesResponsePtrOutput)
}

// The containers within the container group.
func (o LookupContainerGroupResultOutput) Containers() ContainerResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []ContainerResponse { return v.Containers }).(ContainerResponseArrayOutput)
}

// The diagnostic information for a container group.
func (o LookupContainerGroupResultOutput) Diagnostics() ContainerGroupDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *ContainerGroupDiagnosticsResponse { return v.Diagnostics }).(ContainerGroupDiagnosticsResponsePtrOutput)
}

// The DNS config information for a container group.
func (o LookupContainerGroupResultOutput) DnsConfig() DnsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *DnsConfigurationResponse { return v.DnsConfig }).(DnsConfigurationResponsePtrOutput)
}

// The encryption properties for a container group.
func (o LookupContainerGroupResultOutput) EncryptionProperties() EncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *EncryptionPropertiesResponse { return v.EncryptionProperties }).(EncryptionPropertiesResponsePtrOutput)
}

// extensions used by virtual kubelet
func (o LookupContainerGroupResultOutput) Extensions() DeploymentExtensionSpecResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []DeploymentExtensionSpecResponse { return v.Extensions }).(DeploymentExtensionSpecResponseArrayOutput)
}

// The resource id.
func (o LookupContainerGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the container group, if configured.
func (o LookupContainerGroupResultOutput) Identity() ContainerGroupIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *ContainerGroupIdentityResponse { return v.Identity }).(ContainerGroupIdentityResponsePtrOutput)
}

// The image registry credentials by which the container group is created from.
func (o LookupContainerGroupResultOutput) ImageRegistryCredentials() ImageRegistryCredentialResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []ImageRegistryCredentialResponse {
		return v.ImageRegistryCredentials
	}).(ImageRegistryCredentialResponseArrayOutput)
}

// The init containers for a container group.
func (o LookupContainerGroupResultOutput) InitContainers() InitContainerDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []InitContainerDefinitionResponse { return v.InitContainers }).(InitContainerDefinitionResponseArrayOutput)
}

// The instance view of the container group. Only valid in response.
func (o LookupContainerGroupResultOutput) InstanceView() ContainerGroupPropertiesResponseInstanceViewOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) ContainerGroupPropertiesResponseInstanceView { return v.InstanceView }).(ContainerGroupPropertiesResponseInstanceViewOutput)
}

// The IP address type of the container group.
func (o LookupContainerGroupResultOutput) IpAddress() IpAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *IpAddressResponse { return v.IpAddress }).(IpAddressResponsePtrOutput)
}

// The resource location.
func (o LookupContainerGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o LookupContainerGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The operating system type required by the containers in the container group.
func (o LookupContainerGroupResultOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.OsType }).(pulumi.StringOutput)
}

// The priority of the container group.
func (o LookupContainerGroupResultOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.Priority }).(pulumi.StringPtrOutput)
}

// The provisioning state of the container group. This only appears in the response.
func (o LookupContainerGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Restart policy for all containers within the container group.
// - `Always` Always restart
// - `OnFailure` Restart on failure
// - `Never` Never restart
func (o LookupContainerGroupResultOutput) RestartPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.RestartPolicy }).(pulumi.StringPtrOutput)
}

// The SKU for a container group.
func (o LookupContainerGroupResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

// The subnet resource IDs for a container group.
func (o LookupContainerGroupResultOutput) SubnetIds() ContainerGroupSubnetIdResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []ContainerGroupSubnetIdResponse { return v.SubnetIds }).(ContainerGroupSubnetIdResponseArrayOutput)
}

// The resource tags.
func (o LookupContainerGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupContainerGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// The list of volumes that can be mounted by containers in this container group.
func (o LookupContainerGroupResultOutput) Volumes() VolumeResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []VolumeResponse { return v.Volumes }).(VolumeResponseArrayOutput)
}

// The zones for the container group.
func (o LookupContainerGroupResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerGroupResultOutput{})
}
