// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes an Azure AI Search service and its current state.
type Service struct {
	pulumi.CustomResourceState

	// Defines the options for how the data plane API of a search service authenticates requests. This cannot be set if 'disableLocalAuth' is set to true.
	AuthOptions DataPlaneAuthOptionsResponsePtrOutput `pulumi:"authOptions"`
	// When set to true, calls to the search service will not be permitted to utilize API keys for authentication. This cannot be set to true if 'dataPlaneAuthOptions' are defined.
	DisableLocalAuth pulumi.BoolPtrOutput `pulumi:"disableLocalAuth"`
	// A list of data exfiltration scenarios that are explicitly disallowed for the search service. Currently, the only supported value is 'All' to disable all possible data export scenarios with more fine grained controls planned for the future.
	DisabledDataExfiltrationOptions pulumi.StringArrayOutput `pulumi:"disabledDataExfiltrationOptions"`
	// A system generated property representing the service's etag that can be for optimistic concurrency control during updates.
	ETag pulumi.StringOutput `pulumi:"eTag"`
	// Specifies any policy regarding encryption of resources (such as indexes) using customer manager keys within a search service.
	EncryptionWithCmk EncryptionWithCmkResponsePtrOutput `pulumi:"encryptionWithCmk"`
	// Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'.
	HostingMode pulumi.StringPtrOutput `pulumi:"hostingMode"`
	// The identity of the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Network specific rules that determine how the Azure AI Search service may be reached.
	NetworkRuleSet NetworkRuleSetResponsePtrOutput `pulumi:"networkRuleSet"`
	// The number of partitions in the search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For 'standard3' services with hostingMode set to 'highDensity', the allowed values are between 1 and 3.
	PartitionCount pulumi.IntPtrOutput `pulumi:"partitionCount"`
	// The list of private endpoint connections to the Azure AI Search service.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The state of the last provisioning operation performed on the search service. Provisioning is an intermediate state that occurs while service capacity is being established. After capacity is set up, provisioningState changes to either 'succeeded' or 'failed'. Client applications can poll provisioning status (the recommended polling interval is from 30 seconds to one minute) by using the Get Search Service operation to see when an operation is completed. If you are using the free service, this value tends to come back as 'succeeded' directly in the call to Create search service. This is because the free service uses capacity that is already set up.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The number of replicas in the search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.
	ReplicaCount pulumi.IntPtrOutput `pulumi:"replicaCount"`
	// Sets options that control the availability of semantic search. This configuration is only possible for certain Azure AI Search SKUs in certain locations.
	SemanticSearch pulumi.StringPtrOutput `pulumi:"semanticSearch"`
	// The list of shared private link resources managed by the Azure AI Search service.
	SharedPrivateLinkResources SharedPrivateLinkResourceResponseArrayOutput `pulumi:"sharedPrivateLinkResources"`
	// The SKU of the search service, which determines price tier and capacity limits. This property is required when creating a new search service.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The status of the search service. Possible values include: 'running': The search service is running and no provisioning operations are underway. 'provisioning': The search service is being provisioned or scaled up or down. 'deleting': The search service is being deleted. 'degraded': The search service is degraded. This can occur when the underlying search units are not healthy. The search service is most likely operational, but performance might be slow and some requests might be dropped. 'disabled': The search service is disabled. In this state, the service will reject all API requests. 'error': The search service is in an error state. 'stopped': The search service is in a subscription that's disabled. If your service is in the degraded, disabled, or error states, it means the Azure AI Search team is actively investigating the underlying issue. Dedicated services in these states are still chargeable based on the number of search units provisioned.
	Status pulumi.StringOutput `pulumi:"status"`
	// The details of the search service status.
	StatusDetails pulumi.StringOutput `pulumi:"statusDetails"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.HostingMode == nil {
		args.HostingMode = HostingMode("default")
	}
	if args.PartitionCount == nil {
		args.PartitionCount = pulumi.IntPtr(1)
	}
	if args.PublicNetworkAccess == nil {
		args.PublicNetworkAccess = pulumi.StringPtr("enabled")
	}
	if args.ReplicaCount == nil {
		args.ReplicaCount = pulumi.IntPtr(1)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:search:Service"),
		},
		{
			Type: pulumi.String("azure-native:search/v20150819:Service"),
		},
		{
			Type: pulumi.String("azure-native:search/v20191001preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:search/v20200313:Service"),
		},
		{
			Type: pulumi.String("azure-native:search/v20200801:Service"),
		},
		{
			Type: pulumi.String("azure-native:search/v20200801preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:search/v20210401preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:search/v20220901:Service"),
		},
		{
			Type: pulumi.String("azure-native:search/v20231101:Service"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Service
	err := ctx.RegisterResource("azure-native:search/v20240301preview:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:search/v20240301preview:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
}

type ServiceState struct {
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Defines the options for how the data plane API of a search service authenticates requests. This cannot be set if 'disableLocalAuth' is set to true.
	AuthOptions *DataPlaneAuthOptions `pulumi:"authOptions"`
	// When set to true, calls to the search service will not be permitted to utilize API keys for authentication. This cannot be set to true if 'dataPlaneAuthOptions' are defined.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// A list of data exfiltration scenarios that are explicitly disallowed for the search service. Currently, the only supported value is 'All' to disable all possible data export scenarios with more fine grained controls planned for the future.
	DisabledDataExfiltrationOptions []string `pulumi:"disabledDataExfiltrationOptions"`
	// Specifies any policy regarding encryption of resources (such as indexes) using customer manager keys within a search service.
	EncryptionWithCmk *EncryptionWithCmk `pulumi:"encryptionWithCmk"`
	// Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'.
	HostingMode *HostingMode `pulumi:"hostingMode"`
	// The identity of the resource.
	Identity *Identity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Network specific rules that determine how the Azure AI Search service may be reached.
	NetworkRuleSet *NetworkRuleSet `pulumi:"networkRuleSet"`
	// The number of partitions in the search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For 'standard3' services with hostingMode set to 'highDensity', the allowed values are between 1 and 3.
	PartitionCount *int `pulumi:"partitionCount"`
	// This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The number of replicas in the search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.
	ReplicaCount *int `pulumi:"replicaCount"`
	// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Azure AI Search service to create or update. Search service names must only contain lowercase letters, digits or dashes, cannot use dash as the first two or last one characters, cannot contain consecutive dashes, and must be between 2 and 60 characters in length. Search service names must be globally unique since they are part of the service URI (https://<name>.search.windows.net). You cannot change the service name after the service is created.
	SearchServiceName *string `pulumi:"searchServiceName"`
	// Sets options that control the availability of semantic search. This configuration is only possible for certain Azure AI Search SKUs in certain locations.
	SemanticSearch *string `pulumi:"semanticSearch"`
	// The SKU of the search service, which determines price tier and capacity limits. This property is required when creating a new search service.
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Defines the options for how the data plane API of a search service authenticates requests. This cannot be set if 'disableLocalAuth' is set to true.
	AuthOptions DataPlaneAuthOptionsPtrInput
	// When set to true, calls to the search service will not be permitted to utilize API keys for authentication. This cannot be set to true if 'dataPlaneAuthOptions' are defined.
	DisableLocalAuth pulumi.BoolPtrInput
	// A list of data exfiltration scenarios that are explicitly disallowed for the search service. Currently, the only supported value is 'All' to disable all possible data export scenarios with more fine grained controls planned for the future.
	DisabledDataExfiltrationOptions pulumi.StringArrayInput
	// Specifies any policy regarding encryption of resources (such as indexes) using customer manager keys within a search service.
	EncryptionWithCmk EncryptionWithCmkPtrInput
	// Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'.
	HostingMode HostingModePtrInput
	// The identity of the resource.
	Identity IdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Network specific rules that determine how the Azure AI Search service may be reached.
	NetworkRuleSet NetworkRuleSetPtrInput
	// The number of partitions in the search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For 'standard3' services with hostingMode set to 'highDensity', the allowed values are between 1 and 3.
	PartitionCount pulumi.IntPtrInput
	// This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
	PublicNetworkAccess pulumi.StringPtrInput
	// The number of replicas in the search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.
	ReplicaCount pulumi.IntPtrInput
	// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Azure AI Search service to create or update. Search service names must only contain lowercase letters, digits or dashes, cannot use dash as the first two or last one characters, cannot contain consecutive dashes, and must be between 2 and 60 characters in length. Search service names must be globally unique since they are part of the service URI (https://<name>.search.windows.net). You cannot change the service name after the service is created.
	SearchServiceName pulumi.StringPtrInput
	// Sets options that control the availability of semantic search. This configuration is only possible for certain Azure AI Search SKUs in certain locations.
	SemanticSearch pulumi.StringPtrInput
	// The SKU of the search service, which determines price tier and capacity limits. This property is required when creating a new search service.
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

// Defines the options for how the data plane API of a search service authenticates requests. This cannot be set if 'disableLocalAuth' is set to true.
func (o ServiceOutput) AuthOptions() DataPlaneAuthOptionsResponsePtrOutput {
	return o.ApplyT(func(v *Service) DataPlaneAuthOptionsResponsePtrOutput { return v.AuthOptions }).(DataPlaneAuthOptionsResponsePtrOutput)
}

// When set to true, calls to the search service will not be permitted to utilize API keys for authentication. This cannot be set to true if 'dataPlaneAuthOptions' are defined.
func (o ServiceOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.BoolPtrOutput { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// A list of data exfiltration scenarios that are explicitly disallowed for the search service. Currently, the only supported value is 'All' to disable all possible data export scenarios with more fine grained controls planned for the future.
func (o ServiceOutput) DisabledDataExfiltrationOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Service) pulumi.StringArrayOutput { return v.DisabledDataExfiltrationOptions }).(pulumi.StringArrayOutput)
}

// A system generated property representing the service's etag that can be for optimistic concurrency control during updates.
func (o ServiceOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

// Specifies any policy regarding encryption of resources (such as indexes) using customer manager keys within a search service.
func (o ServiceOutput) EncryptionWithCmk() EncryptionWithCmkResponsePtrOutput {
	return o.ApplyT(func(v *Service) EncryptionWithCmkResponsePtrOutput { return v.EncryptionWithCmk }).(EncryptionWithCmkResponsePtrOutput)
}

// Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'.
func (o ServiceOutput) HostingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.HostingMode }).(pulumi.StringPtrOutput)
}

// The identity of the resource.
func (o ServiceOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Service) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o ServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network specific rules that determine how the Azure AI Search service may be reached.
func (o ServiceOutput) NetworkRuleSet() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v *Service) NetworkRuleSetResponsePtrOutput { return v.NetworkRuleSet }).(NetworkRuleSetResponsePtrOutput)
}

// The number of partitions in the search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For 'standard3' services with hostingMode set to 'highDensity', the allowed values are between 1 and 3.
func (o ServiceOutput) PartitionCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.IntPtrOutput { return v.PartitionCount }).(pulumi.IntPtrOutput)
}

// The list of private endpoint connections to the Azure AI Search service.
func (o ServiceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Service) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// The state of the last provisioning operation performed on the search service. Provisioning is an intermediate state that occurs while service capacity is being established. After capacity is set up, provisioningState changes to either 'succeeded' or 'failed'. Client applications can poll provisioning status (the recommended polling interval is from 30 seconds to one minute) by using the Get Search Service operation to see when an operation is completed. If you are using the free service, this value tends to come back as 'succeeded' directly in the call to Create search service. This is because the free service uses capacity that is already set up.
func (o ServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
func (o ServiceOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The number of replicas in the search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.
func (o ServiceOutput) ReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.IntPtrOutput { return v.ReplicaCount }).(pulumi.IntPtrOutput)
}

// Sets options that control the availability of semantic search. This configuration is only possible for certain Azure AI Search SKUs in certain locations.
func (o ServiceOutput) SemanticSearch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.SemanticSearch }).(pulumi.StringPtrOutput)
}

// The list of shared private link resources managed by the Azure AI Search service.
func (o ServiceOutput) SharedPrivateLinkResources() SharedPrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v *Service) SharedPrivateLinkResourceResponseArrayOutput { return v.SharedPrivateLinkResources }).(SharedPrivateLinkResourceResponseArrayOutput)
}

// The SKU of the search service, which determines price tier and capacity limits. This property is required when creating a new search service.
func (o ServiceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Service) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The status of the search service. Possible values include: 'running': The search service is running and no provisioning operations are underway. 'provisioning': The search service is being provisioned or scaled up or down. 'deleting': The search service is being deleted. 'degraded': The search service is degraded. This can occur when the underlying search units are not healthy. The search service is most likely operational, but performance might be slow and some requests might be dropped. 'disabled': The search service is disabled. In this state, the service will reject all API requests. 'error': The search service is in an error state. 'stopped': The search service is in a subscription that's disabled. If your service is in the degraded, disabled, or error states, it means the Azure AI Search team is actively investigating the underlying issue. Dedicated services in these states are still chargeable based on the number of search units provisioned.
func (o ServiceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The details of the search service status.
func (o ServiceOutput) StatusDetails() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.StatusDetails }).(pulumi.StringOutput)
}

// Resource tags.
func (o ServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
