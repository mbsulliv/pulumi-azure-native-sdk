// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191027preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// OpenShift Managed cluster.
type OpenShiftManagedCluster struct {
	pulumi.CustomResourceState

	// Configuration of OpenShift cluster VMs.
	AgentPoolProfiles OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput `pulumi:"agentPoolProfiles"`
	// Configures OpenShift authentication.
	AuthProfile OpenShiftManagedClusterAuthProfileResponsePtrOutput `pulumi:"authProfile"`
	// Version of OpenShift specified when creating the cluster.
	ClusterVersion pulumi.StringOutput `pulumi:"clusterVersion"`
	// Service generated FQDN for OpenShift API server loadbalancer internal hostname.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Configuration for OpenShift master VMs.
	MasterPoolProfile OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput `pulumi:"masterPoolProfile"`
	// Configures Log Analytics integration.
	MonitorProfile OpenShiftManagedClusterMonitorProfileResponsePtrOutput `pulumi:"monitorProfile"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration for OpenShift networking.
	NetworkProfile NetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// Version of OpenShift specified when creating the cluster.
	OpenShiftVersion pulumi.StringOutput `pulumi:"openShiftVersion"`
	// Define the resource plan as required by ARM for billing purposes
	Plan PurchasePlanResponsePtrOutput `pulumi:"plan"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Service generated FQDN or private IP for OpenShift API server.
	PublicHostname pulumi.StringOutput `pulumi:"publicHostname"`
	// Allows node rotation
	RefreshCluster pulumi.BoolPtrOutput `pulumi:"refreshCluster"`
	// Configuration for OpenShift router(s).
	RouterProfiles OpenShiftRouterProfileResponseArrayOutput `pulumi:"routerProfiles"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOpenShiftManagedCluster registers a new resource with the given unique name, arguments, and options.
func NewOpenShiftManagedCluster(ctx *pulumi.Context,
	name string, args *OpenShiftManagedClusterArgs, opts ...pulumi.ResourceOption) (*OpenShiftManagedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OpenShiftVersion == nil {
		return nil, errors.New("invalid value for required argument 'OpenShiftVersion'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.NetworkProfile != nil {
		args.NetworkProfile = args.NetworkProfile.ToNetworkProfilePtrOutput().ApplyT(func(v *NetworkProfile) *NetworkProfile { return v.Defaults() }).(NetworkProfilePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:OpenShiftManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20180930preview:OpenShiftManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190430:OpenShiftManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190930preview:OpenShiftManagedCluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource OpenShiftManagedCluster
	err := ctx.RegisterResource("azure-native:containerservice/v20191027preview:OpenShiftManagedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOpenShiftManagedCluster gets an existing OpenShiftManagedCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOpenShiftManagedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OpenShiftManagedClusterState, opts ...pulumi.ResourceOption) (*OpenShiftManagedCluster, error) {
	var resource OpenShiftManagedCluster
	err := ctx.ReadResource("azure-native:containerservice/v20191027preview:OpenShiftManagedCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OpenShiftManagedCluster resources.
type openShiftManagedClusterState struct {
}

type OpenShiftManagedClusterState struct {
}

func (OpenShiftManagedClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*openShiftManagedClusterState)(nil)).Elem()
}

type openShiftManagedClusterArgs struct {
	// Configuration of OpenShift cluster VMs.
	AgentPoolProfiles []OpenShiftManagedClusterAgentPoolProfile `pulumi:"agentPoolProfiles"`
	// Configures OpenShift authentication.
	AuthProfile *OpenShiftManagedClusterAuthProfile `pulumi:"authProfile"`
	// Resource location
	Location *string `pulumi:"location"`
	// Configuration for OpenShift master VMs.
	MasterPoolProfile *OpenShiftManagedClusterMasterPoolProfile `pulumi:"masterPoolProfile"`
	// Configures Log Analytics integration.
	MonitorProfile *OpenShiftManagedClusterMonitorProfile `pulumi:"monitorProfile"`
	// Configuration for OpenShift networking.
	NetworkProfile *NetworkProfile `pulumi:"networkProfile"`
	// Version of OpenShift specified when creating the cluster.
	OpenShiftVersion string `pulumi:"openShiftVersion"`
	// Define the resource plan as required by ARM for billing purposes
	Plan *PurchasePlan `pulumi:"plan"`
	// Allows node rotation
	RefreshCluster *bool `pulumi:"refreshCluster"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the OpenShift managed cluster resource.
	ResourceName *string `pulumi:"resourceName"`
	// Configuration for OpenShift router(s).
	RouterProfiles []OpenShiftRouterProfile `pulumi:"routerProfiles"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a OpenShiftManagedCluster resource.
type OpenShiftManagedClusterArgs struct {
	// Configuration of OpenShift cluster VMs.
	AgentPoolProfiles OpenShiftManagedClusterAgentPoolProfileArrayInput
	// Configures OpenShift authentication.
	AuthProfile OpenShiftManagedClusterAuthProfilePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Configuration for OpenShift master VMs.
	MasterPoolProfile OpenShiftManagedClusterMasterPoolProfilePtrInput
	// Configures Log Analytics integration.
	MonitorProfile OpenShiftManagedClusterMonitorProfilePtrInput
	// Configuration for OpenShift networking.
	NetworkProfile NetworkProfilePtrInput
	// Version of OpenShift specified when creating the cluster.
	OpenShiftVersion pulumi.StringInput
	// Define the resource plan as required by ARM for billing purposes
	Plan PurchasePlanPtrInput
	// Allows node rotation
	RefreshCluster pulumi.BoolPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the OpenShift managed cluster resource.
	ResourceName pulumi.StringPtrInput
	// Configuration for OpenShift router(s).
	RouterProfiles OpenShiftRouterProfileArrayInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (OpenShiftManagedClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*openShiftManagedClusterArgs)(nil)).Elem()
}

type OpenShiftManagedClusterInput interface {
	pulumi.Input

	ToOpenShiftManagedClusterOutput() OpenShiftManagedClusterOutput
	ToOpenShiftManagedClusterOutputWithContext(ctx context.Context) OpenShiftManagedClusterOutput
}

func (*OpenShiftManagedCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftManagedCluster)(nil)).Elem()
}

func (i *OpenShiftManagedCluster) ToOpenShiftManagedClusterOutput() OpenShiftManagedClusterOutput {
	return i.ToOpenShiftManagedClusterOutputWithContext(context.Background())
}

func (i *OpenShiftManagedCluster) ToOpenShiftManagedClusterOutputWithContext(ctx context.Context) OpenShiftManagedClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterOutput)
}

type OpenShiftManagedClusterOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftManagedCluster)(nil)).Elem()
}

func (o OpenShiftManagedClusterOutput) ToOpenShiftManagedClusterOutput() OpenShiftManagedClusterOutput {
	return o
}

func (o OpenShiftManagedClusterOutput) ToOpenShiftManagedClusterOutputWithContext(ctx context.Context) OpenShiftManagedClusterOutput {
	return o
}

// Configuration of OpenShift cluster VMs.
func (o OpenShiftManagedClusterOutput) AgentPoolProfiles() OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput {
		return v.AgentPoolProfiles
	}).(OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput)
}

// Configures OpenShift authentication.
func (o OpenShiftManagedClusterOutput) AuthProfile() OpenShiftManagedClusterAuthProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) OpenShiftManagedClusterAuthProfileResponsePtrOutput {
		return v.AuthProfile
	}).(OpenShiftManagedClusterAuthProfileResponsePtrOutput)
}

// Version of OpenShift specified when creating the cluster.
func (o OpenShiftManagedClusterOutput) ClusterVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.ClusterVersion }).(pulumi.StringOutput)
}

// Service generated FQDN for OpenShift API server loadbalancer internal hostname.
func (o OpenShiftManagedClusterOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// Resource location
func (o OpenShiftManagedClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Configuration for OpenShift master VMs.
func (o OpenShiftManagedClusterOutput) MasterPoolProfile() OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput {
		return v.MasterPoolProfile
	}).(OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput)
}

// Configures Log Analytics integration.
func (o OpenShiftManagedClusterOutput) MonitorProfile() OpenShiftManagedClusterMonitorProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) OpenShiftManagedClusterMonitorProfileResponsePtrOutput {
		return v.MonitorProfile
	}).(OpenShiftManagedClusterMonitorProfileResponsePtrOutput)
}

// Resource name
func (o OpenShiftManagedClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Configuration for OpenShift networking.
func (o OpenShiftManagedClusterOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) NetworkProfileResponsePtrOutput { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

// Version of OpenShift specified when creating the cluster.
func (o OpenShiftManagedClusterOutput) OpenShiftVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.OpenShiftVersion }).(pulumi.StringOutput)
}

// Define the resource plan as required by ARM for billing purposes
func (o OpenShiftManagedClusterOutput) Plan() PurchasePlanResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) PurchasePlanResponsePtrOutput { return v.Plan }).(PurchasePlanResponsePtrOutput)
}

// The current deployment or provisioning state, which only appears in the response.
func (o OpenShiftManagedClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Service generated FQDN or private IP for OpenShift API server.
func (o OpenShiftManagedClusterOutput) PublicHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.PublicHostname }).(pulumi.StringOutput)
}

// Allows node rotation
func (o OpenShiftManagedClusterOutput) RefreshCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.BoolPtrOutput { return v.RefreshCluster }).(pulumi.BoolPtrOutput)
}

// Configuration for OpenShift router(s).
func (o OpenShiftManagedClusterOutput) RouterProfiles() OpenShiftRouterProfileResponseArrayOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) OpenShiftRouterProfileResponseArrayOutput { return v.RouterProfiles }).(OpenShiftRouterProfileResponseArrayOutput)
}

// Resource tags
func (o OpenShiftManagedClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o OpenShiftManagedClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OpenShiftManagedClusterOutput{})
}
