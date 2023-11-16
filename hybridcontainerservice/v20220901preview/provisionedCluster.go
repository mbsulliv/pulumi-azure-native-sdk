// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provisionedClusters resource definition.
type ProvisionedCluster struct {
	pulumi.CustomResourceState

	ExtendedLocation ProvisionedClustersResponseResponseExtendedLocationPtrOutput `pulumi:"extendedLocation"`
	// Identity for the Provisioned cluster.
	Identity ProvisionedClusterIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name       pulumi.StringOutput                                 `pulumi:"name"`
	Properties ProvisionedClustersResponsePropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProvisionedCluster registers a new resource with the given unique name, arguments, and options.
func NewProvisionedCluster(ctx *pulumi.Context,
	name string, args *ProvisionedClusterArgs, opts ...pulumi.ResourceOption) (*ProvisionedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToProvisionedClustersAllPropertiesPtrOutput().ApplyT(func(v *ProvisionedClustersAllProperties) *ProvisionedClustersAllProperties { return v.Defaults() }).(ProvisionedClustersAllPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcontainerservice:ProvisionedCluster"),
		},
		{
			Type: pulumi.String("azure-native:hybridcontainerservice/v20220501preview:ProvisionedCluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProvisionedCluster
	err := ctx.RegisterResource("azure-native:hybridcontainerservice/v20220901preview:ProvisionedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProvisionedCluster gets an existing ProvisionedCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProvisionedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProvisionedClusterState, opts ...pulumi.ResourceOption) (*ProvisionedCluster, error) {
	var resource ProvisionedCluster
	err := ctx.ReadResource("azure-native:hybridcontainerservice/v20220901preview:ProvisionedCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProvisionedCluster resources.
type provisionedClusterState struct {
}

type ProvisionedClusterState struct {
}

func (ProvisionedClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*provisionedClusterState)(nil)).Elem()
}

type provisionedClusterArgs struct {
	ExtendedLocation *ProvisionedClustersExtendedLocation `pulumi:"extendedLocation"`
	// Identity for the Provisioned cluster.
	Identity *ProvisionedClusterIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// All properties of the provisioned cluster
	Properties *ProvisionedClustersAllProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Parameter for the name of the provisioned cluster
	ResourceName *string `pulumi:"resourceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ProvisionedCluster resource.
type ProvisionedClusterArgs struct {
	ExtendedLocation ProvisionedClustersExtendedLocationPtrInput
	// Identity for the Provisioned cluster.
	Identity ProvisionedClusterIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// All properties of the provisioned cluster
	Properties ProvisionedClustersAllPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Parameter for the name of the provisioned cluster
	ResourceName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ProvisionedClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*provisionedClusterArgs)(nil)).Elem()
}

type ProvisionedClusterInput interface {
	pulumi.Input

	ToProvisionedClusterOutput() ProvisionedClusterOutput
	ToProvisionedClusterOutputWithContext(ctx context.Context) ProvisionedClusterOutput
}

func (*ProvisionedCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedCluster)(nil)).Elem()
}

func (i *ProvisionedCluster) ToProvisionedClusterOutput() ProvisionedClusterOutput {
	return i.ToProvisionedClusterOutputWithContext(context.Background())
}

func (i *ProvisionedCluster) ToProvisionedClusterOutputWithContext(ctx context.Context) ProvisionedClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClusterOutput)
}

type ProvisionedClusterOutput struct{ *pulumi.OutputState }

func (ProvisionedClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedCluster)(nil)).Elem()
}

func (o ProvisionedClusterOutput) ToProvisionedClusterOutput() ProvisionedClusterOutput {
	return o
}

func (o ProvisionedClusterOutput) ToProvisionedClusterOutputWithContext(ctx context.Context) ProvisionedClusterOutput {
	return o
}

func (o ProvisionedClusterOutput) ExtendedLocation() ProvisionedClustersResponseResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v *ProvisionedCluster) ProvisionedClustersResponseResponseExtendedLocationPtrOutput {
		return v.ExtendedLocation
	}).(ProvisionedClustersResponseResponseExtendedLocationPtrOutput)
}

// Identity for the Provisioned cluster.
func (o ProvisionedClusterOutput) Identity() ProvisionedClusterIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ProvisionedCluster) ProvisionedClusterIdentityResponsePtrOutput { return v.Identity }).(ProvisionedClusterIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o ProvisionedClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ProvisionedClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProvisionedClusterOutput) Properties() ProvisionedClustersResponsePropertiesResponseOutput {
	return o.ApplyT(func(v *ProvisionedCluster) ProvisionedClustersResponsePropertiesResponseOutput { return v.Properties }).(ProvisionedClustersResponsePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ProvisionedClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ProvisionedCluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ProvisionedClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProvisionedCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ProvisionedClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProvisionedClusterOutput{})
}
