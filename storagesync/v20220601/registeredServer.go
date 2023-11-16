// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Registered Server resource.
type RegisteredServer struct {
	pulumi.CustomResourceState

	// Registered Server Agent Version
	AgentVersion pulumi.StringPtrOutput `pulumi:"agentVersion"`
	// Registered Server Agent Version Expiration Date
	AgentVersionExpirationDate pulumi.StringOutput `pulumi:"agentVersionExpirationDate"`
	// Registered Server Agent Version Status
	AgentVersionStatus pulumi.StringOutput `pulumi:"agentVersionStatus"`
	// Registered Server clusterId
	ClusterId pulumi.StringPtrOutput `pulumi:"clusterId"`
	// Registered Server clusterName
	ClusterName pulumi.StringPtrOutput `pulumi:"clusterName"`
	// Resource discoveryEndpointUri
	DiscoveryEndpointUri pulumi.StringPtrOutput `pulumi:"discoveryEndpointUri"`
	// Friendly Name
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// Registered Server last heart beat
	LastHeartBeat pulumi.StringPtrOutput `pulumi:"lastHeartBeat"`
	// Resource Last Operation Name
	LastOperationName pulumi.StringPtrOutput `pulumi:"lastOperationName"`
	// Registered Server lastWorkflowId
	LastWorkflowId pulumi.StringPtrOutput `pulumi:"lastWorkflowId"`
	// Management Endpoint Uri
	ManagementEndpointUri pulumi.StringPtrOutput `pulumi:"managementEndpointUri"`
	// Monitoring Configuration
	MonitoringConfiguration pulumi.StringPtrOutput `pulumi:"monitoringConfiguration"`
	// Telemetry Endpoint Uri
	MonitoringEndpointUri pulumi.StringPtrOutput `pulumi:"monitoringEndpointUri"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Registered Server Provisioning State
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Resource Location
	ResourceLocation pulumi.StringPtrOutput `pulumi:"resourceLocation"`
	// Registered Server Certificate
	ServerCertificate pulumi.StringPtrOutput `pulumi:"serverCertificate"`
	// Registered Server serverId
	ServerId pulumi.StringPtrOutput `pulumi:"serverId"`
	// Registered Server Management Error Code
	ServerManagementErrorCode pulumi.IntPtrOutput `pulumi:"serverManagementErrorCode"`
	// Server name
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Registered Server OS Version
	ServerOSVersion pulumi.StringPtrOutput `pulumi:"serverOSVersion"`
	// Registered Server serverRole
	ServerRole pulumi.StringPtrOutput `pulumi:"serverRole"`
	// Service Location
	ServiceLocation pulumi.StringPtrOutput `pulumi:"serviceLocation"`
	// Registered Server storageSyncServiceUid
	StorageSyncServiceUid pulumi.StringPtrOutput `pulumi:"storageSyncServiceUid"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegisteredServer registers a new resource with the given unique name, arguments, and options.
func NewRegisteredServer(ctx *pulumi.Context,
	name string, args *RegisteredServerArgs, opts ...pulumi.ResourceOption) (*RegisteredServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageSyncServiceName == nil {
		return nil, errors.New("invalid value for required argument 'StorageSyncServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagesync:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20170605preview:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180402:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180701:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20181001:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190201:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190301:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190601:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20191001:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200301:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200901:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20220901:RegisteredServer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RegisteredServer
	err := ctx.RegisterResource("azure-native:storagesync/v20220601:RegisteredServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegisteredServer gets an existing RegisteredServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegisteredServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegisteredServerState, opts ...pulumi.ResourceOption) (*RegisteredServer, error) {
	var resource RegisteredServer
	err := ctx.ReadResource("azure-native:storagesync/v20220601:RegisteredServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegisteredServer resources.
type registeredServerState struct {
}

type RegisteredServerState struct {
}

func (RegisteredServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredServerState)(nil)).Elem()
}

type registeredServerArgs struct {
	// Registered Server Agent Version
	AgentVersion *string `pulumi:"agentVersion"`
	// Registered Server clusterId
	ClusterId *string `pulumi:"clusterId"`
	// Registered Server clusterName
	ClusterName *string `pulumi:"clusterName"`
	// Friendly Name
	FriendlyName *string `pulumi:"friendlyName"`
	// Registered Server last heart beat
	LastHeartBeat *string `pulumi:"lastHeartBeat"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Registered Server Certificate
	ServerCertificate *string `pulumi:"serverCertificate"`
	// Registered Server serverId
	ServerId *string `pulumi:"serverId"`
	// Registered Server OS Version
	ServerOSVersion *string `pulumi:"serverOSVersion"`
	// Registered Server serverRole
	ServerRole *string `pulumi:"serverRole"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
}

// The set of arguments for constructing a RegisteredServer resource.
type RegisteredServerArgs struct {
	// Registered Server Agent Version
	AgentVersion pulumi.StringPtrInput
	// Registered Server clusterId
	ClusterId pulumi.StringPtrInput
	// Registered Server clusterName
	ClusterName pulumi.StringPtrInput
	// Friendly Name
	FriendlyName pulumi.StringPtrInput
	// Registered Server last heart beat
	LastHeartBeat pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Registered Server Certificate
	ServerCertificate pulumi.StringPtrInput
	// Registered Server serverId
	ServerId pulumi.StringPtrInput
	// Registered Server OS Version
	ServerOSVersion pulumi.StringPtrInput
	// Registered Server serverRole
	ServerRole pulumi.StringPtrInput
	// Name of Storage Sync Service resource.
	StorageSyncServiceName pulumi.StringInput
}

func (RegisteredServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredServerArgs)(nil)).Elem()
}

type RegisteredServerInput interface {
	pulumi.Input

	ToRegisteredServerOutput() RegisteredServerOutput
	ToRegisteredServerOutputWithContext(ctx context.Context) RegisteredServerOutput
}

func (*RegisteredServer) ElementType() reflect.Type {
	return reflect.TypeOf((**RegisteredServer)(nil)).Elem()
}

func (i *RegisteredServer) ToRegisteredServerOutput() RegisteredServerOutput {
	return i.ToRegisteredServerOutputWithContext(context.Background())
}

func (i *RegisteredServer) ToRegisteredServerOutputWithContext(ctx context.Context) RegisteredServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegisteredServerOutput)
}

type RegisteredServerOutput struct{ *pulumi.OutputState }

func (RegisteredServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegisteredServer)(nil)).Elem()
}

func (o RegisteredServerOutput) ToRegisteredServerOutput() RegisteredServerOutput {
	return o
}

func (o RegisteredServerOutput) ToRegisteredServerOutputWithContext(ctx context.Context) RegisteredServerOutput {
	return o
}

// Registered Server Agent Version
func (o RegisteredServerOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

// Registered Server Agent Version Expiration Date
func (o RegisteredServerOutput) AgentVersionExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringOutput { return v.AgentVersionExpirationDate }).(pulumi.StringOutput)
}

// Registered Server Agent Version Status
func (o RegisteredServerOutput) AgentVersionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringOutput { return v.AgentVersionStatus }).(pulumi.StringOutput)
}

// Registered Server clusterId
func (o RegisteredServerOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// Registered Server clusterName
func (o RegisteredServerOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.ClusterName }).(pulumi.StringPtrOutput)
}

// Resource discoveryEndpointUri
func (o RegisteredServerOutput) DiscoveryEndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.DiscoveryEndpointUri }).(pulumi.StringPtrOutput)
}

// Friendly Name
func (o RegisteredServerOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// Registered Server last heart beat
func (o RegisteredServerOutput) LastHeartBeat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.LastHeartBeat }).(pulumi.StringPtrOutput)
}

// Resource Last Operation Name
func (o RegisteredServerOutput) LastOperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.LastOperationName }).(pulumi.StringPtrOutput)
}

// Registered Server lastWorkflowId
func (o RegisteredServerOutput) LastWorkflowId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.LastWorkflowId }).(pulumi.StringPtrOutput)
}

// Management Endpoint Uri
func (o RegisteredServerOutput) ManagementEndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.ManagementEndpointUri }).(pulumi.StringPtrOutput)
}

// Monitoring Configuration
func (o RegisteredServerOutput) MonitoringConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.MonitoringConfiguration }).(pulumi.StringPtrOutput)
}

// Telemetry Endpoint Uri
func (o RegisteredServerOutput) MonitoringEndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.MonitoringEndpointUri }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o RegisteredServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Registered Server Provisioning State
func (o RegisteredServerOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Resource Location
func (o RegisteredServerOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

// Registered Server Certificate
func (o RegisteredServerOutput) ServerCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.ServerCertificate }).(pulumi.StringPtrOutput)
}

// Registered Server serverId
func (o RegisteredServerOutput) ServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.ServerId }).(pulumi.StringPtrOutput)
}

// Registered Server Management Error Code
func (o RegisteredServerOutput) ServerManagementErrorCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.IntPtrOutput { return v.ServerManagementErrorCode }).(pulumi.IntPtrOutput)
}

// Server name
func (o RegisteredServerOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringOutput { return v.ServerName }).(pulumi.StringOutput)
}

// Registered Server OS Version
func (o RegisteredServerOutput) ServerOSVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.ServerOSVersion }).(pulumi.StringPtrOutput)
}

// Registered Server serverRole
func (o RegisteredServerOutput) ServerRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.ServerRole }).(pulumi.StringPtrOutput)
}

// Service Location
func (o RegisteredServerOutput) ServiceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.ServiceLocation }).(pulumi.StringPtrOutput)
}

// Registered Server storageSyncServiceUid
func (o RegisteredServerOutput) StorageSyncServiceUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringPtrOutput { return v.StorageSyncServiceUid }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RegisteredServerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegisteredServer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RegisteredServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegisteredServerOutput{})
}
