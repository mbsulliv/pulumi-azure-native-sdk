// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221108

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a cluster.
type Cluster struct {
	pulumi.CustomResourceState

	// The administrator's login name of the servers in the cluster.
	AdministratorLogin pulumi.StringOutput `pulumi:"administratorLogin"`
	// The Citus extension version on all cluster servers.
	CitusVersion pulumi.StringPtrOutput `pulumi:"citusVersion"`
	// If public access is enabled on coordinator.
	CoordinatorEnablePublicIpAccess pulumi.BoolPtrOutput `pulumi:"coordinatorEnablePublicIpAccess"`
	// The edition of a coordinator server (default: GeneralPurpose). Required for creation.
	CoordinatorServerEdition pulumi.StringPtrOutput `pulumi:"coordinatorServerEdition"`
	// The storage of a server in MB. Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	CoordinatorStorageQuotaInMb pulumi.IntPtrOutput `pulumi:"coordinatorStorageQuotaInMb"`
	// The vCores count of a server (max: 96). Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	CoordinatorVCores pulumi.IntPtrOutput `pulumi:"coordinatorVCores"`
	// The earliest restore point time (ISO8601 format) for the cluster.
	EarliestRestoreTime pulumi.StringOutput `pulumi:"earliestRestoreTime"`
	// If high availability (HA) is enabled or not for the cluster.
	EnableHa pulumi.BoolPtrOutput `pulumi:"enableHa"`
	// If distributed tables are placed on coordinator or not. Should be set to 'true' on single node clusters. Requires shard rebalancing after value is changed.
	EnableShardsOnCoordinator pulumi.BoolPtrOutput `pulumi:"enableShardsOnCoordinator"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Maintenance window of a cluster.
	MaintenanceWindow MaintenanceWindowResponsePtrOutput `pulumi:"maintenanceWindow"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Worker node count of the cluster. When node count is 0, it represents a single node configuration with the ability to create distributed tables on that node. 2 or more worker nodes represent multi-node configuration. Node count value cannot be 1. Required for creation.
	NodeCount pulumi.IntPtrOutput `pulumi:"nodeCount"`
	// If public access is enabled on worker nodes.
	NodeEnablePublicIpAccess pulumi.BoolPtrOutput `pulumi:"nodeEnablePublicIpAccess"`
	// The edition of a node server (default: MemoryOptimized).
	NodeServerEdition pulumi.StringPtrOutput `pulumi:"nodeServerEdition"`
	// The storage in MB on each worker node. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	NodeStorageQuotaInMb pulumi.IntPtrOutput `pulumi:"nodeStorageQuotaInMb"`
	// The compute in vCores on each worker node (max: 104). See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	NodeVCores pulumi.IntPtrOutput `pulumi:"nodeVCores"`
	// Date and time in UTC (ISO8601 format) for cluster restore.
	PointInTimeUTC pulumi.StringPtrOutput `pulumi:"pointInTimeUTC"`
	// The major PostgreSQL version on all cluster servers.
	PostgresqlVersion pulumi.StringPtrOutput `pulumi:"postgresqlVersion"`
	// Preferred primary availability zone (AZ) for all cluster servers.
	PreferredPrimaryZone pulumi.StringPtrOutput `pulumi:"preferredPrimaryZone"`
	// The private endpoint connections for a cluster.
	PrivateEndpointConnections SimplePrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Provisioning state of the cluster
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The array of read replica clusters.
	ReadReplicas pulumi.StringArrayOutput `pulumi:"readReplicas"`
	// The list of server names in the cluster
	ServerNames ServerNameItemResponseArrayOutput `pulumi:"serverNames"`
	// The Azure region of source cluster for read replica clusters.
	SourceLocation pulumi.StringPtrOutput `pulumi:"sourceLocation"`
	// The resource id of source cluster for read replica clusters.
	SourceResourceId pulumi.StringPtrOutput `pulumi:"sourceResourceId"`
	// A state of a cluster/server that is visible to user.
	State pulumi.StringOutput `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
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
			Type: pulumi.String("azure-native:dbforpostgresql:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20201005privatepreview:Cluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:dbforpostgresql/v20221108:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:dbforpostgresql/v20221108:Cluster", name, id, state, &resource, opts...)
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
	// The password of the administrator login. Required for creation.
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// The Citus extension version on all cluster servers.
	CitusVersion *string `pulumi:"citusVersion"`
	// The name of the cluster.
	ClusterName *string `pulumi:"clusterName"`
	// If public access is enabled on coordinator.
	CoordinatorEnablePublicIpAccess *bool `pulumi:"coordinatorEnablePublicIpAccess"`
	// The edition of a coordinator server (default: GeneralPurpose). Required for creation.
	CoordinatorServerEdition *string `pulumi:"coordinatorServerEdition"`
	// The storage of a server in MB. Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	CoordinatorStorageQuotaInMb *int `pulumi:"coordinatorStorageQuotaInMb"`
	// The vCores count of a server (max: 96). Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	CoordinatorVCores *int `pulumi:"coordinatorVCores"`
	// If high availability (HA) is enabled or not for the cluster.
	EnableHa *bool `pulumi:"enableHa"`
	// If distributed tables are placed on coordinator or not. Should be set to 'true' on single node clusters. Requires shard rebalancing after value is changed.
	EnableShardsOnCoordinator *bool `pulumi:"enableShardsOnCoordinator"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Maintenance window of a cluster.
	MaintenanceWindow *MaintenanceWindow `pulumi:"maintenanceWindow"`
	// Worker node count of the cluster. When node count is 0, it represents a single node configuration with the ability to create distributed tables on that node. 2 or more worker nodes represent multi-node configuration. Node count value cannot be 1. Required for creation.
	NodeCount *int `pulumi:"nodeCount"`
	// If public access is enabled on worker nodes.
	NodeEnablePublicIpAccess *bool `pulumi:"nodeEnablePublicIpAccess"`
	// The edition of a node server (default: MemoryOptimized).
	NodeServerEdition *string `pulumi:"nodeServerEdition"`
	// The storage in MB on each worker node. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	NodeStorageQuotaInMb *int `pulumi:"nodeStorageQuotaInMb"`
	// The compute in vCores on each worker node (max: 104). See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	NodeVCores *int `pulumi:"nodeVCores"`
	// Date and time in UTC (ISO8601 format) for cluster restore.
	PointInTimeUTC *string `pulumi:"pointInTimeUTC"`
	// The major PostgreSQL version on all cluster servers.
	PostgresqlVersion *string `pulumi:"postgresqlVersion"`
	// Preferred primary availability zone (AZ) for all cluster servers.
	PreferredPrimaryZone *string `pulumi:"preferredPrimaryZone"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Azure region of source cluster for read replica clusters.
	SourceLocation *string `pulumi:"sourceLocation"`
	// The resource id of source cluster for read replica clusters.
	SourceResourceId *string `pulumi:"sourceResourceId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The password of the administrator login. Required for creation.
	AdministratorLoginPassword pulumi.StringPtrInput
	// The Citus extension version on all cluster servers.
	CitusVersion pulumi.StringPtrInput
	// The name of the cluster.
	ClusterName pulumi.StringPtrInput
	// If public access is enabled on coordinator.
	CoordinatorEnablePublicIpAccess pulumi.BoolPtrInput
	// The edition of a coordinator server (default: GeneralPurpose). Required for creation.
	CoordinatorServerEdition pulumi.StringPtrInput
	// The storage of a server in MB. Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	CoordinatorStorageQuotaInMb pulumi.IntPtrInput
	// The vCores count of a server (max: 96). Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	CoordinatorVCores pulumi.IntPtrInput
	// If high availability (HA) is enabled or not for the cluster.
	EnableHa pulumi.BoolPtrInput
	// If distributed tables are placed on coordinator or not. Should be set to 'true' on single node clusters. Requires shard rebalancing after value is changed.
	EnableShardsOnCoordinator pulumi.BoolPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Maintenance window of a cluster.
	MaintenanceWindow MaintenanceWindowPtrInput
	// Worker node count of the cluster. When node count is 0, it represents a single node configuration with the ability to create distributed tables on that node. 2 or more worker nodes represent multi-node configuration. Node count value cannot be 1. Required for creation.
	NodeCount pulumi.IntPtrInput
	// If public access is enabled on worker nodes.
	NodeEnablePublicIpAccess pulumi.BoolPtrInput
	// The edition of a node server (default: MemoryOptimized).
	NodeServerEdition pulumi.StringPtrInput
	// The storage in MB on each worker node. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	NodeStorageQuotaInMb pulumi.IntPtrInput
	// The compute in vCores on each worker node (max: 104). See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	NodeVCores pulumi.IntPtrInput
	// Date and time in UTC (ISO8601 format) for cluster restore.
	PointInTimeUTC pulumi.StringPtrInput
	// The major PostgreSQL version on all cluster servers.
	PostgresqlVersion pulumi.StringPtrInput
	// Preferred primary availability zone (AZ) for all cluster servers.
	PreferredPrimaryZone pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The Azure region of source cluster for read replica clusters.
	SourceLocation pulumi.StringPtrInput
	// The resource id of source cluster for read replica clusters.
	SourceResourceId pulumi.StringPtrInput
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

// The administrator's login name of the servers in the cluster.
func (o ClusterOutput) AdministratorLogin() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.AdministratorLogin }).(pulumi.StringOutput)
}

// The Citus extension version on all cluster servers.
func (o ClusterOutput) CitusVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.CitusVersion }).(pulumi.StringPtrOutput)
}

// If public access is enabled on coordinator.
func (o ClusterOutput) CoordinatorEnablePublicIpAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.CoordinatorEnablePublicIpAccess }).(pulumi.BoolPtrOutput)
}

// The edition of a coordinator server (default: GeneralPurpose). Required for creation.
func (o ClusterOutput) CoordinatorServerEdition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.CoordinatorServerEdition }).(pulumi.StringPtrOutput)
}

// The storage of a server in MB. Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
func (o ClusterOutput) CoordinatorStorageQuotaInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntPtrOutput { return v.CoordinatorStorageQuotaInMb }).(pulumi.IntPtrOutput)
}

// The vCores count of a server (max: 96). Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
func (o ClusterOutput) CoordinatorVCores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntPtrOutput { return v.CoordinatorVCores }).(pulumi.IntPtrOutput)
}

// The earliest restore point time (ISO8601 format) for the cluster.
func (o ClusterOutput) EarliestRestoreTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.EarliestRestoreTime }).(pulumi.StringOutput)
}

// If high availability (HA) is enabled or not for the cluster.
func (o ClusterOutput) EnableHa() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.EnableHa }).(pulumi.BoolPtrOutput)
}

// If distributed tables are placed on coordinator or not. Should be set to 'true' on single node clusters. Requires shard rebalancing after value is changed.
func (o ClusterOutput) EnableShardsOnCoordinator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.EnableShardsOnCoordinator }).(pulumi.BoolPtrOutput)
}

// The geo-location where the resource lives
func (o ClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Maintenance window of a cluster.
func (o ClusterOutput) MaintenanceWindow() MaintenanceWindowResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) MaintenanceWindowResponsePtrOutput { return v.MaintenanceWindow }).(MaintenanceWindowResponsePtrOutput)
}

// The name of the resource
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Worker node count of the cluster. When node count is 0, it represents a single node configuration with the ability to create distributed tables on that node. 2 or more worker nodes represent multi-node configuration. Node count value cannot be 1. Required for creation.
func (o ClusterOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntPtrOutput { return v.NodeCount }).(pulumi.IntPtrOutput)
}

// If public access is enabled on worker nodes.
func (o ClusterOutput) NodeEnablePublicIpAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.NodeEnablePublicIpAccess }).(pulumi.BoolPtrOutput)
}

// The edition of a node server (default: MemoryOptimized).
func (o ClusterOutput) NodeServerEdition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.NodeServerEdition }).(pulumi.StringPtrOutput)
}

// The storage in MB on each worker node. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
func (o ClusterOutput) NodeStorageQuotaInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntPtrOutput { return v.NodeStorageQuotaInMb }).(pulumi.IntPtrOutput)
}

// The compute in vCores on each worker node (max: 104). See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
func (o ClusterOutput) NodeVCores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntPtrOutput { return v.NodeVCores }).(pulumi.IntPtrOutput)
}

// Date and time in UTC (ISO8601 format) for cluster restore.
func (o ClusterOutput) PointInTimeUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.PointInTimeUTC }).(pulumi.StringPtrOutput)
}

// The major PostgreSQL version on all cluster servers.
func (o ClusterOutput) PostgresqlVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.PostgresqlVersion }).(pulumi.StringPtrOutput)
}

// Preferred primary availability zone (AZ) for all cluster servers.
func (o ClusterOutput) PreferredPrimaryZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.PreferredPrimaryZone }).(pulumi.StringPtrOutput)
}

// The private endpoint connections for a cluster.
func (o ClusterOutput) PrivateEndpointConnections() SimplePrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) SimplePrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(SimplePrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the cluster
func (o ClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The array of read replica clusters.
func (o ClusterOutput) ReadReplicas() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.ReadReplicas }).(pulumi.StringArrayOutput)
}

// The list of server names in the cluster
func (o ClusterOutput) ServerNames() ServerNameItemResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) ServerNameItemResponseArrayOutput { return v.ServerNames }).(ServerNameItemResponseArrayOutput)
}

// The Azure region of source cluster for read replica clusters.
func (o ClusterOutput) SourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.SourceLocation }).(pulumi.StringPtrOutput)
}

// The resource id of source cluster for read replica clusters.
func (o ClusterOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

// A state of a cluster/server that is visible to user.
func (o ClusterOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Cluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
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
