// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200430

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// OpenShiftCluster represents an Azure Red Hat OpenShift cluster.
type OpenShiftCluster struct {
	pulumi.CustomResourceState

	// The cluster API server profile.
	ApiserverProfile APIServerProfileResponsePtrOutput `pulumi:"apiserverProfile"`
	// The cluster profile.
	ClusterProfile ClusterProfileResponsePtrOutput `pulumi:"clusterProfile"`
	// The console profile.
	ConsoleProfile ConsoleProfileResponsePtrOutput `pulumi:"consoleProfile"`
	// The cluster ingress profiles.
	IngressProfiles IngressProfileResponseArrayOutput `pulumi:"ingressProfiles"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The cluster master profile.
	MasterProfile MasterProfileResponsePtrOutput `pulumi:"masterProfile"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The cluster network profile.
	NetworkProfile NetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// The cluster provisioning state (immutable).
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The cluster service principal profile.
	ServicePrincipalProfile ServicePrincipalProfileResponsePtrOutput `pulumi:"servicePrincipalProfile"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The cluster worker profiles.
	WorkerProfiles WorkerProfileResponseArrayOutput `pulumi:"workerProfiles"`
}

// NewOpenShiftCluster registers a new resource with the given unique name, arguments, and options.
func NewOpenShiftCluster(ctx *pulumi.Context,
	name string, args *OpenShiftClusterArgs, opts ...pulumi.ResourceOption) (*OpenShiftCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:redhatopenshift:OpenShiftCluster"),
		},
		{
			Type: pulumi.String("azure-native:redhatopenshift/v20210901preview:OpenShiftCluster"),
		},
		{
			Type: pulumi.String("azure-native:redhatopenshift/v20220401:OpenShiftCluster"),
		},
		{
			Type: pulumi.String("azure-native:redhatopenshift/v20220904:OpenShiftCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource OpenShiftCluster
	err := ctx.RegisterResource("azure-native:redhatopenshift/v20200430:OpenShiftCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOpenShiftCluster gets an existing OpenShiftCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOpenShiftCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OpenShiftClusterState, opts ...pulumi.ResourceOption) (*OpenShiftCluster, error) {
	var resource OpenShiftCluster
	err := ctx.ReadResource("azure-native:redhatopenshift/v20200430:OpenShiftCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OpenShiftCluster resources.
type openShiftClusterState struct {
}

type OpenShiftClusterState struct {
}

func (OpenShiftClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*openShiftClusterState)(nil)).Elem()
}

type openShiftClusterArgs struct {
	// The cluster API server profile.
	ApiserverProfile *APIServerProfile `pulumi:"apiserverProfile"`
	// The cluster profile.
	ClusterProfile *ClusterProfile `pulumi:"clusterProfile"`
	// The console profile.
	ConsoleProfile *ConsoleProfile `pulumi:"consoleProfile"`
	// The cluster ingress profiles.
	IngressProfiles []IngressProfile `pulumi:"ingressProfiles"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The cluster master profile.
	MasterProfile *MasterProfile `pulumi:"masterProfile"`
	// The cluster network profile.
	NetworkProfile *NetworkProfile `pulumi:"networkProfile"`
	// The cluster provisioning state (immutable).
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the OpenShift cluster resource.
	ResourceName *string `pulumi:"resourceName"`
	// The cluster service principal profile.
	ServicePrincipalProfile *ServicePrincipalProfile `pulumi:"servicePrincipalProfile"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The cluster worker profiles.
	WorkerProfiles []WorkerProfile `pulumi:"workerProfiles"`
}

// The set of arguments for constructing a OpenShiftCluster resource.
type OpenShiftClusterArgs struct {
	// The cluster API server profile.
	ApiserverProfile APIServerProfilePtrInput
	// The cluster profile.
	ClusterProfile ClusterProfilePtrInput
	// The console profile.
	ConsoleProfile ConsoleProfilePtrInput
	// The cluster ingress profiles.
	IngressProfiles IngressProfileArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The cluster master profile.
	MasterProfile MasterProfilePtrInput
	// The cluster network profile.
	NetworkProfile NetworkProfilePtrInput
	// The cluster provisioning state (immutable).
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the OpenShift cluster resource.
	ResourceName pulumi.StringPtrInput
	// The cluster service principal profile.
	ServicePrincipalProfile ServicePrincipalProfilePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The cluster worker profiles.
	WorkerProfiles WorkerProfileArrayInput
}

func (OpenShiftClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*openShiftClusterArgs)(nil)).Elem()
}

type OpenShiftClusterInput interface {
	pulumi.Input

	ToOpenShiftClusterOutput() OpenShiftClusterOutput
	ToOpenShiftClusterOutputWithContext(ctx context.Context) OpenShiftClusterOutput
}

func (*OpenShiftCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftCluster)(nil)).Elem()
}

func (i *OpenShiftCluster) ToOpenShiftClusterOutput() OpenShiftClusterOutput {
	return i.ToOpenShiftClusterOutputWithContext(context.Background())
}

func (i *OpenShiftCluster) ToOpenShiftClusterOutputWithContext(ctx context.Context) OpenShiftClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftClusterOutput)
}

type OpenShiftClusterOutput struct{ *pulumi.OutputState }

func (OpenShiftClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftCluster)(nil)).Elem()
}

func (o OpenShiftClusterOutput) ToOpenShiftClusterOutput() OpenShiftClusterOutput {
	return o
}

func (o OpenShiftClusterOutput) ToOpenShiftClusterOutputWithContext(ctx context.Context) OpenShiftClusterOutput {
	return o
}

// The cluster API server profile.
func (o OpenShiftClusterOutput) ApiserverProfile() APIServerProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftCluster) APIServerProfileResponsePtrOutput { return v.ApiserverProfile }).(APIServerProfileResponsePtrOutput)
}

// The cluster profile.
func (o OpenShiftClusterOutput) ClusterProfile() ClusterProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftCluster) ClusterProfileResponsePtrOutput { return v.ClusterProfile }).(ClusterProfileResponsePtrOutput)
}

// The console profile.
func (o OpenShiftClusterOutput) ConsoleProfile() ConsoleProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftCluster) ConsoleProfileResponsePtrOutput { return v.ConsoleProfile }).(ConsoleProfileResponsePtrOutput)
}

// The cluster ingress profiles.
func (o OpenShiftClusterOutput) IngressProfiles() IngressProfileResponseArrayOutput {
	return o.ApplyT(func(v *OpenShiftCluster) IngressProfileResponseArrayOutput { return v.IngressProfiles }).(IngressProfileResponseArrayOutput)
}

// The geo-location where the resource lives
func (o OpenShiftClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The cluster master profile.
func (o OpenShiftClusterOutput) MasterProfile() MasterProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftCluster) MasterProfileResponsePtrOutput { return v.MasterProfile }).(MasterProfileResponsePtrOutput)
}

// The name of the resource
func (o OpenShiftClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The cluster network profile.
func (o OpenShiftClusterOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftCluster) NetworkProfileResponsePtrOutput { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

// The cluster provisioning state (immutable).
func (o OpenShiftClusterOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftCluster) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The cluster service principal profile.
func (o OpenShiftClusterOutput) ServicePrincipalProfile() ServicePrincipalProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftCluster) ServicePrincipalProfileResponsePtrOutput { return v.ServicePrincipalProfile }).(ServicePrincipalProfileResponsePtrOutput)
}

// Resource tags.
func (o OpenShiftClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OpenShiftCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o OpenShiftClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The cluster worker profiles.
func (o OpenShiftClusterOutput) WorkerProfiles() WorkerProfileResponseArrayOutput {
	return o.ApplyT(func(v *OpenShiftCluster) WorkerProfileResponseArrayOutput { return v.WorkerProfiles }).(WorkerProfileResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(OpenShiftClusterOutput{})
}
