// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents a connected cluster.
type ConnectedCluster struct {
	pulumi.CustomResourceState

	// Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
	AgentPublicKeyCertificate pulumi.StringOutput `pulumi:"agentPublicKeyCertificate"`
	// Version of the agent running on the connected cluster resource
	AgentVersion pulumi.StringOutput `pulumi:"agentVersion"`
	// Indicates whether Azure Hybrid Benefit is opted in
	AzureHybridBenefit pulumi.StringPtrOutput `pulumi:"azureHybridBenefit"`
	// Represents the connectivity status of the connected cluster.
	ConnectivityStatus pulumi.StringOutput `pulumi:"connectivityStatus"`
	// The Kubernetes distribution running on this connected cluster.
	Distribution pulumi.StringPtrOutput `pulumi:"distribution"`
	// The Kubernetes distribution version on this connected cluster.
	DistributionVersion pulumi.StringPtrOutput `pulumi:"distributionVersion"`
	// The identity of the connected cluster.
	Identity ConnectedClusterIdentityResponseOutput `pulumi:"identity"`
	// The infrastructure on which the Kubernetes cluster represented by this connected cluster is running on.
	Infrastructure pulumi.StringPtrOutput `pulumi:"infrastructure"`
	// The Kubernetes version of the connected cluster resource
	KubernetesVersion pulumi.StringOutput `pulumi:"kubernetesVersion"`
	// Time representing the last instance when heart beat was received from the cluster
	LastConnectivityTime pulumi.StringOutput `pulumi:"lastConnectivityTime"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Expiration time of the managed identity certificate
	ManagedIdentityCertificateExpirationTime pulumi.StringOutput `pulumi:"managedIdentityCertificateExpirationTime"`
	// More properties related to the Connected Cluster
	MiscellaneousProperties pulumi.StringMapOutput `pulumi:"miscellaneousProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Connected cluster offering
	Offering pulumi.StringOutput `pulumi:"offering"`
	// The resource id of the private link scope this connected cluster is assigned to, if any.
	PrivateLinkScopeResourceId pulumi.StringPtrOutput `pulumi:"privateLinkScopeResourceId"`
	// Property which describes the state of private link on a connected cluster resource.
	PrivateLinkState pulumi.StringPtrOutput `pulumi:"privateLinkState"`
	// Provisioning state of the connected cluster resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Number of CPU cores present in the connected cluster resource
	TotalCoreCount pulumi.IntOutput `pulumi:"totalCoreCount"`
	// Number of nodes present in the connected cluster resource
	TotalNodeCount pulumi.IntOutput `pulumi:"totalNodeCount"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnectedCluster registers a new resource with the given unique name, arguments, and options.
func NewConnectedCluster(ctx *pulumi.Context,
	name string, args *ConnectedClusterArgs, opts ...pulumi.ResourceOption) (*ConnectedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentPublicKeyCertificate == nil {
		return nil, errors.New("invalid value for required argument 'AgentPublicKeyCertificate'")
	}
	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.AzureHybridBenefit == nil {
		args.AzureHybridBenefit = pulumi.StringPtr("NotApplicable")
	}
	args.Identity = args.Identity.ToConnectedClusterIdentityOutput().ApplyT(func(v ConnectedClusterIdentity) ConnectedClusterIdentity { return *v.Defaults() }).(ConnectedClusterIdentityOutput)
	if args.PrivateLinkState == nil {
		args.PrivateLinkState = pulumi.StringPtr("Disabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kubernetes:ConnectedCluster"),
		},
		{
			Type: pulumi.String("azure-native:kubernetes/v20200101preview:ConnectedCluster"),
		},
		{
			Type: pulumi.String("azure-native:kubernetes/v20210301:ConnectedCluster"),
		},
		{
			Type: pulumi.String("azure-native:kubernetes/v20210401preview:ConnectedCluster"),
		},
		{
			Type: pulumi.String("azure-native:kubernetes/v20211001:ConnectedCluster"),
		},
		{
			Type: pulumi.String("azure-native:kubernetes/v20220501preview:ConnectedCluster"),
		},
		{
			Type: pulumi.String("azure-native:kubernetes/v20231101preview:ConnectedCluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConnectedCluster
	err := ctx.RegisterResource("azure-native:kubernetes/v20221001preview:ConnectedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectedCluster gets an existing ConnectedCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectedClusterState, opts ...pulumi.ResourceOption) (*ConnectedCluster, error) {
	var resource ConnectedCluster
	err := ctx.ReadResource("azure-native:kubernetes/v20221001preview:ConnectedCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectedCluster resources.
type connectedClusterState struct {
}

type ConnectedClusterState struct {
}

func (ConnectedClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedClusterState)(nil)).Elem()
}

type connectedClusterArgs struct {
	// Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
	AgentPublicKeyCertificate string `pulumi:"agentPublicKeyCertificate"`
	// Indicates whether Azure Hybrid Benefit is opted in
	AzureHybridBenefit *string `pulumi:"azureHybridBenefit"`
	// The name of the Kubernetes cluster on which get is called.
	ClusterName *string `pulumi:"clusterName"`
	// The Kubernetes distribution running on this connected cluster.
	Distribution *string `pulumi:"distribution"`
	// The Kubernetes distribution version on this connected cluster.
	DistributionVersion *string `pulumi:"distributionVersion"`
	// The identity of the connected cluster.
	Identity ConnectedClusterIdentity `pulumi:"identity"`
	// The infrastructure on which the Kubernetes cluster represented by this connected cluster is running on.
	Infrastructure *string `pulumi:"infrastructure"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The resource id of the private link scope this connected cluster is assigned to, if any.
	PrivateLinkScopeResourceId *string `pulumi:"privateLinkScopeResourceId"`
	// Property which describes the state of private link on a connected cluster resource.
	PrivateLinkState *string `pulumi:"privateLinkState"`
	// Provisioning state of the connected cluster resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ConnectedCluster resource.
type ConnectedClusterArgs struct {
	// Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
	AgentPublicKeyCertificate pulumi.StringInput
	// Indicates whether Azure Hybrid Benefit is opted in
	AzureHybridBenefit pulumi.StringPtrInput
	// The name of the Kubernetes cluster on which get is called.
	ClusterName pulumi.StringPtrInput
	// The Kubernetes distribution running on this connected cluster.
	Distribution pulumi.StringPtrInput
	// The Kubernetes distribution version on this connected cluster.
	DistributionVersion pulumi.StringPtrInput
	// The identity of the connected cluster.
	Identity ConnectedClusterIdentityInput
	// The infrastructure on which the Kubernetes cluster represented by this connected cluster is running on.
	Infrastructure pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The resource id of the private link scope this connected cluster is assigned to, if any.
	PrivateLinkScopeResourceId pulumi.StringPtrInput
	// Property which describes the state of private link on a connected cluster resource.
	PrivateLinkState pulumi.StringPtrInput
	// Provisioning state of the connected cluster resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ConnectedClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedClusterArgs)(nil)).Elem()
}

type ConnectedClusterInput interface {
	pulumi.Input

	ToConnectedClusterOutput() ConnectedClusterOutput
	ToConnectedClusterOutputWithContext(ctx context.Context) ConnectedClusterOutput
}

func (*ConnectedCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedCluster)(nil)).Elem()
}

func (i *ConnectedCluster) ToConnectedClusterOutput() ConnectedClusterOutput {
	return i.ToConnectedClusterOutputWithContext(context.Background())
}

func (i *ConnectedCluster) ToConnectedClusterOutputWithContext(ctx context.Context) ConnectedClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterOutput)
}

func (i *ConnectedCluster) ToOutput(ctx context.Context) pulumix.Output[*ConnectedCluster] {
	return pulumix.Output[*ConnectedCluster]{
		OutputState: i.ToConnectedClusterOutputWithContext(ctx).OutputState,
	}
}

type ConnectedClusterOutput struct{ *pulumi.OutputState }

func (ConnectedClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedCluster)(nil)).Elem()
}

func (o ConnectedClusterOutput) ToConnectedClusterOutput() ConnectedClusterOutput {
	return o
}

func (o ConnectedClusterOutput) ToConnectedClusterOutputWithContext(ctx context.Context) ConnectedClusterOutput {
	return o
}

func (o ConnectedClusterOutput) ToOutput(ctx context.Context) pulumix.Output[*ConnectedCluster] {
	return pulumix.Output[*ConnectedCluster]{
		OutputState: o.OutputState,
	}
}

// Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
func (o ConnectedClusterOutput) AgentPublicKeyCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.AgentPublicKeyCertificate }).(pulumi.StringOutput)
}

// Version of the agent running on the connected cluster resource
func (o ConnectedClusterOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.AgentVersion }).(pulumi.StringOutput)
}

// Indicates whether Azure Hybrid Benefit is opted in
func (o ConnectedClusterOutput) AzureHybridBenefit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringPtrOutput { return v.AzureHybridBenefit }).(pulumi.StringPtrOutput)
}

// Represents the connectivity status of the connected cluster.
func (o ConnectedClusterOutput) ConnectivityStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.ConnectivityStatus }).(pulumi.StringOutput)
}

// The Kubernetes distribution running on this connected cluster.
func (o ConnectedClusterOutput) Distribution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringPtrOutput { return v.Distribution }).(pulumi.StringPtrOutput)
}

// The Kubernetes distribution version on this connected cluster.
func (o ConnectedClusterOutput) DistributionVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringPtrOutput { return v.DistributionVersion }).(pulumi.StringPtrOutput)
}

// The identity of the connected cluster.
func (o ConnectedClusterOutput) Identity() ConnectedClusterIdentityResponseOutput {
	return o.ApplyT(func(v *ConnectedCluster) ConnectedClusterIdentityResponseOutput { return v.Identity }).(ConnectedClusterIdentityResponseOutput)
}

// The infrastructure on which the Kubernetes cluster represented by this connected cluster is running on.
func (o ConnectedClusterOutput) Infrastructure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringPtrOutput { return v.Infrastructure }).(pulumi.StringPtrOutput)
}

// The Kubernetes version of the connected cluster resource
func (o ConnectedClusterOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.KubernetesVersion }).(pulumi.StringOutput)
}

// Time representing the last instance when heart beat was received from the cluster
func (o ConnectedClusterOutput) LastConnectivityTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.LastConnectivityTime }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o ConnectedClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Expiration time of the managed identity certificate
func (o ConnectedClusterOutput) ManagedIdentityCertificateExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.ManagedIdentityCertificateExpirationTime }).(pulumi.StringOutput)
}

// More properties related to the Connected Cluster
func (o ConnectedClusterOutput) MiscellaneousProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringMapOutput { return v.MiscellaneousProperties }).(pulumi.StringMapOutput)
}

// The name of the resource
func (o ConnectedClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Connected cluster offering
func (o ConnectedClusterOutput) Offering() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.Offering }).(pulumi.StringOutput)
}

// The resource id of the private link scope this connected cluster is assigned to, if any.
func (o ConnectedClusterOutput) PrivateLinkScopeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringPtrOutput { return v.PrivateLinkScopeResourceId }).(pulumi.StringPtrOutput)
}

// Property which describes the state of private link on a connected cluster resource.
func (o ConnectedClusterOutput) PrivateLinkState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringPtrOutput { return v.PrivateLinkState }).(pulumi.StringPtrOutput)
}

// Provisioning state of the connected cluster resource.
func (o ConnectedClusterOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource
func (o ConnectedClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectedCluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ConnectedClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Number of CPU cores present in the connected cluster resource
func (o ConnectedClusterOutput) TotalCoreCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.IntOutput { return v.TotalCoreCount }).(pulumi.IntOutput)
}

// Number of nodes present in the connected cluster resource
func (o ConnectedClusterOutput) TotalNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.IntOutput { return v.TotalNodeCount }).(pulumi.IntOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConnectedClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectedClusterOutput{})
}
