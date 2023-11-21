// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Single Namespace item in List or Get Operation
type Namespace struct {
	pulumi.CustomResourceState

	// Alternate name specified when alias and namespace names are same.
	AlternateName pulumi.StringPtrOutput `pulumi:"alternateName"`
	// Cluster ARM ID of the Namespace.
	ClusterArmId pulumi.StringPtrOutput `pulumi:"clusterArmId"`
	// The time the Namespace was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// This property disables SAS authentication for the Event Hubs namespace.
	DisableLocalAuth pulumi.BoolPtrOutput `pulumi:"disableLocalAuth"`
	// Properties of BYOK Encryption description
	Encryption EncryptionResponsePtrOutput `pulumi:"encryption"`
	// Properties of BYOK Identity description
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Value that indicates whether AutoInflate is enabled for eventhub namespace.
	IsAutoInflateEnabled pulumi.BoolPtrOutput `pulumi:"isAutoInflateEnabled"`
	// Value that indicates whether Kafka is enabled for eventhub namespace.
	KafkaEnabled pulumi.BoolPtrOutput `pulumi:"kafkaEnabled"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units. ( '0' if AutoInflateEnabled = true)
	MaximumThroughputUnits pulumi.IntPtrOutput `pulumi:"maximumThroughputUnits"`
	// Identifier for Azure Insights metrics.
	MetricId pulumi.StringOutput `pulumi:"metricId"`
	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion pulumi.StringPtrOutput `pulumi:"minimumTlsVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of private endpoint connections.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Provisioning state of the Namespace.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint pulumi.StringOutput `pulumi:"serviceBusEndpoint"`
	// Properties of sku resource
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Status of the Namespace.
	Status pulumi.StringOutput `pulumi:"status"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the Namespace was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.
	ZoneRedundant pulumi.BoolPtrOutput `pulumi:"zoneRedundant"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Encryption != nil {
		args.Encryption = args.Encryption.ToEncryptionPtrOutput().ApplyT(func(v *Encryption) *Encryption { return v.Defaults() }).(EncryptionPtrOutput)
	}
	if args.PublicNetworkAccess == nil {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventhub:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20140901:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20150801:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20170401:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20180101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210601preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20211101:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20220101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20221001preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20230101preview:Namespace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Namespace
	err := ctx.RegisterResource("azure-native:eventhub/v20240101:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure-native:eventhub/v20240101:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
}

type NamespaceState struct {
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// Alternate name specified when alias and namespace names are same.
	AlternateName *string `pulumi:"alternateName"`
	// Cluster ARM ID of the Namespace.
	ClusterArmId *string `pulumi:"clusterArmId"`
	// This property disables SAS authentication for the Event Hubs namespace.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// Properties of BYOK Encryption description
	Encryption *Encryption `pulumi:"encryption"`
	// Properties of BYOK Identity description
	Identity *Identity `pulumi:"identity"`
	// Value that indicates whether AutoInflate is enabled for eventhub namespace.
	IsAutoInflateEnabled *bool `pulumi:"isAutoInflateEnabled"`
	// Value that indicates whether Kafka is enabled for eventhub namespace.
	KafkaEnabled *bool `pulumi:"kafkaEnabled"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units. ( '0' if AutoInflateEnabled = true)
	MaximumThroughputUnits *int `pulumi:"maximumThroughputUnits"`
	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// The Namespace name
	NamespaceName *string `pulumi:"namespaceName"`
	// List of private endpoint connections.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	PrivateEndpointConnections []PrivateEndpointConnectionType `pulumi:"privateEndpointConnections"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Properties of sku resource
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Alternate name specified when alias and namespace names are same.
	AlternateName pulumi.StringPtrInput
	// Cluster ARM ID of the Namespace.
	ClusterArmId pulumi.StringPtrInput
	// This property disables SAS authentication for the Event Hubs namespace.
	DisableLocalAuth pulumi.BoolPtrInput
	// Properties of BYOK Encryption description
	Encryption EncryptionPtrInput
	// Properties of BYOK Identity description
	Identity IdentityPtrInput
	// Value that indicates whether AutoInflate is enabled for eventhub namespace.
	IsAutoInflateEnabled pulumi.BoolPtrInput
	// Value that indicates whether Kafka is enabled for eventhub namespace.
	KafkaEnabled pulumi.BoolPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units. ( '0' if AutoInflateEnabled = true)
	MaximumThroughputUnits pulumi.IntPtrInput
	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion pulumi.StringPtrInput
	// The Namespace name
	NamespaceName pulumi.StringPtrInput
	// List of private endpoint connections.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	PrivateEndpointConnections PrivateEndpointConnectionTypeArrayInput
	// This determines if traffic is allowed over public network. By default it is enabled.
	PublicNetworkAccess pulumi.StringPtrInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput
	// Properties of sku resource
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (*Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

type NamespaceOutput struct{ *pulumi.OutputState }

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

// Alternate name specified when alias and namespace names are same.
func (o NamespaceOutput) AlternateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.AlternateName }).(pulumi.StringPtrOutput)
}

// Cluster ARM ID of the Namespace.
func (o NamespaceOutput) ClusterArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.ClusterArmId }).(pulumi.StringPtrOutput)
}

// The time the Namespace was created.
func (o NamespaceOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// This property disables SAS authentication for the Event Hubs namespace.
func (o NamespaceOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.BoolPtrOutput { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// Properties of BYOK Encryption description
func (o NamespaceOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v *Namespace) EncryptionResponsePtrOutput { return v.Encryption }).(EncryptionResponsePtrOutput)
}

// Properties of BYOK Identity description
func (o NamespaceOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Namespace) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// Value that indicates whether AutoInflate is enabled for eventhub namespace.
func (o NamespaceOutput) IsAutoInflateEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.BoolPtrOutput { return v.IsAutoInflateEnabled }).(pulumi.BoolPtrOutput)
}

// Value that indicates whether Kafka is enabled for eventhub namespace.
func (o NamespaceOutput) KafkaEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.BoolPtrOutput { return v.KafkaEnabled }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o NamespaceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units. ( '0' if AutoInflateEnabled = true)
func (o NamespaceOutput) MaximumThroughputUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.IntPtrOutput { return v.MaximumThroughputUnits }).(pulumi.IntPtrOutput)
}

// Identifier for Azure Insights metrics.
func (o NamespaceOutput) MetricId() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.MetricId }).(pulumi.StringOutput)
}

// The minimum TLS version for the cluster to support, e.g. '1.2'
func (o NamespaceOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o NamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of private endpoint connections.
func (o NamespaceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Namespace) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the Namespace.
func (o NamespaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// This determines if traffic is allowed over public network. By default it is enabled.
func (o NamespaceOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Endpoint you can use to perform Service Bus operations.
func (o NamespaceOutput) ServiceBusEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.ServiceBusEndpoint }).(pulumi.StringOutput)
}

// Properties of sku resource
func (o NamespaceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Namespace) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Status of the Namespace.
func (o NamespaceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The system meta data relating to this resource.
func (o NamespaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Namespace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o NamespaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o NamespaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The time the Namespace was updated.
func (o NamespaceOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.
func (o NamespaceOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.BoolPtrOutput { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceOutput{})
}
