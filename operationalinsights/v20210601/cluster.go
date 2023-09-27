// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The top level Log Analytics cluster resource container.
type Cluster struct {
	pulumi.CustomResourceState

	// The list of Log Analytics workspaces associated with the cluster
	AssociatedWorkspaces AssociatedWorkspaceResponseArrayOutput `pulumi:"associatedWorkspaces"`
	// The cluster's billing type.
	BillingType pulumi.StringPtrOutput `pulumi:"billingType"`
	// Additional properties for capacity reservation
	CapacityReservationProperties CapacityReservationPropertiesResponsePtrOutput `pulumi:"capacityReservationProperties"`
	// The ID associated with the cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The cluster creation time
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The identity of the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Sets whether the cluster will support availability zones. This can be set as true only in regions where Azure Data Explorer support Availability Zones. This Property can not be modified after cluster creation. Default value is 'true' if region supports Availability Zones.
	IsAvailabilityZonesEnabled pulumi.BoolPtrOutput `pulumi:"isAvailabilityZonesEnabled"`
	// The associated key properties.
	KeyVaultProperties KeyVaultPropertiesResponsePtrOutput `pulumi:"keyVaultProperties"`
	// The last time the cluster was updated.
	LastModifiedDate pulumi.StringOutput `pulumi:"lastModifiedDate"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the cluster.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The sku properties.
	Sku ClusterSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190801preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20201001:Cluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:operationalinsights/v20210601:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:operationalinsights/v20210601:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The cluster's billing type.
	BillingType *string `pulumi:"billingType"`
	// The name of the Log Analytics cluster.
	ClusterName *string `pulumi:"clusterName"`
	// The identity of the resource.
	Identity *Identity `pulumi:"identity"`
	// Sets whether the cluster will support availability zones. This can be set as true only in regions where Azure Data Explorer support Availability Zones. This Property can not be modified after cluster creation. Default value is 'true' if region supports Availability Zones.
	IsAvailabilityZonesEnabled *bool `pulumi:"isAvailabilityZonesEnabled"`
	// Configures whether cluster will use double encryption. This Property can not be modified after cluster creation. Default value is 'true'
	IsDoubleEncryptionEnabled *bool `pulumi:"isDoubleEncryptionEnabled"`
	// The associated key properties.
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku properties.
	Sku *ClusterSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The cluster's billing type.
	BillingType pulumi.StringPtrInput
	// The name of the Log Analytics cluster.
	ClusterName pulumi.StringPtrInput
	// The identity of the resource.
	Identity IdentityPtrInput
	// Sets whether the cluster will support availability zones. This can be set as true only in regions where Azure Data Explorer support Availability Zones. This Property can not be modified after cluster creation. Default value is 'true' if region supports Availability Zones.
	IsAvailabilityZonesEnabled pulumi.BoolPtrInput
	// Configures whether cluster will use double encryption. This Property can not be modified after cluster creation. Default value is 'true'
	IsDoubleEncryptionEnabled pulumi.BoolPtrInput
	// The associated key properties.
	KeyVaultProperties KeyVaultPropertiesPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The sku properties.
	Sku ClusterSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

func (i *Cluster) ToOutput(ctx context.Context) pulumix.Output[*Cluster] {
	return pulumix.Output[*Cluster]{
		OutputState: i.ToClusterOutputWithContext(ctx).OutputState,
	}
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) ToOutput(ctx context.Context) pulumix.Output[*Cluster] {
	return pulumix.Output[*Cluster]{
		OutputState: o.OutputState,
	}
}

// The list of Log Analytics workspaces associated with the cluster
func (o ClusterOutput) AssociatedWorkspaces() AssociatedWorkspaceResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) AssociatedWorkspaceResponseArrayOutput { return v.AssociatedWorkspaces }).(AssociatedWorkspaceResponseArrayOutput)
}

// The cluster's billing type.
func (o ClusterOutput) BillingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.BillingType }).(pulumi.StringPtrOutput)
}

// Additional properties for capacity reservation
func (o ClusterOutput) CapacityReservationProperties() CapacityReservationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) CapacityReservationPropertiesResponsePtrOutput {
		return v.CapacityReservationProperties
	}).(CapacityReservationPropertiesResponsePtrOutput)
}

// The ID associated with the cluster.
func (o ClusterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The cluster creation time
func (o ClusterOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// The identity of the resource.
func (o ClusterOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// Sets whether the cluster will support availability zones. This can be set as true only in regions where Azure Data Explorer support Availability Zones. This Property can not be modified after cluster creation. Default value is 'true' if region supports Availability Zones.
func (o ClusterOutput) IsAvailabilityZonesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.IsAvailabilityZonesEnabled }).(pulumi.BoolPtrOutput)
}

// The associated key properties.
func (o ClusterOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) KeyVaultPropertiesResponsePtrOutput { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

// The last time the cluster was updated.
func (o ClusterOutput) LastModifiedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.LastModifiedDate }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o ClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the cluster.
func (o ClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The sku properties.
func (o ClusterOutput) Sku() ClusterSkuResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterSkuResponsePtrOutput { return v.Sku }).(ClusterSkuResponsePtrOutput)
}

// Resource tags.
func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
