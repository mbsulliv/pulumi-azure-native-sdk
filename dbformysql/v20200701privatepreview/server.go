// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200701privatepreview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a server.
type Server struct {
	pulumi.CustomResourceState

	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin pulumi.StringPtrOutput `pulumi:"administratorLogin"`
	// availability Zone information of the server.
	AvailabilityZone pulumi.StringPtrOutput `pulumi:"availabilityZone"`
	// Status showing whether the data encryption is enabled with customer-managed keys.
	ByokEnforcement pulumi.StringOutput `pulumi:"byokEnforcement"`
	// Delegated subnet arguments.
	DelegatedSubnetArguments DelegatedSubnetArgumentsResponsePtrOutput `pulumi:"delegatedSubnetArguments"`
	// Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate pulumi.StringOutput `pulumi:"earliestRestoreDate"`
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName pulumi.StringOutput `pulumi:"fullyQualifiedDomainName"`
	// Enable HA or not for a server.
	HaEnabled pulumi.StringPtrOutput `pulumi:"haEnabled"`
	// The state of a HA server.
	HaState pulumi.StringOutput `pulumi:"haState"`
	// The Azure Active Directory identity of the server.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Maintenance window of a server.
	MaintenanceWindow MaintenanceWindowResponsePtrOutput `pulumi:"maintenanceWindow"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess pulumi.StringOutput `pulumi:"publicNetworkAccess"`
	// The maximum number of replicas that a primary server can have.
	ReplicaCapacity pulumi.IntOutput `pulumi:"replicaCapacity"`
	// The replication role.
	ReplicationRole pulumi.StringPtrOutput `pulumi:"replicationRole"`
	// The SKU (pricing tier) of the server.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The source MySQL server id.
	SourceServerId pulumi.StringPtrOutput `pulumi:"sourceServerId"`
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement pulumi.StringPtrOutput `pulumi:"sslEnforcement"`
	// availability Zone information of the server.
	StandbyAvailabilityZone pulumi.StringOutput `pulumi:"standbyAvailabilityZone"`
	// The state of a server.
	State pulumi.StringOutput `pulumi:"state"`
	// Storage profile of a server.
	StorageProfile StorageProfileResponsePtrOutput `pulumi:"storageProfile"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Server version.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dbformysql:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20200701preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20210501:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20210501preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20211201preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20220101:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20220930preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20230601preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20230630:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20231001preview:Server"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Server
	err := ctx.RegisterResource("azure-native:dbformysql/v20200701privatepreview:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azure-native:dbformysql/v20200701privatepreview:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
}

type ServerState struct {
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The password of the administrator login (required for server creation).
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// availability Zone information of the server.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The mode to create a new MySQL server.
	CreateMode *string `pulumi:"createMode"`
	// Delegated subnet arguments.
	DelegatedSubnetArguments *DelegatedSubnetArguments `pulumi:"delegatedSubnetArguments"`
	// Enable HA or not for a server.
	HaEnabled *string `pulumi:"haEnabled"`
	// The Azure Active Directory identity of the server.
	Identity *Identity `pulumi:"identity"`
	// Status showing whether the server enabled infrastructure encryption.
	InfrastructureEncryption *string `pulumi:"infrastructureEncryption"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindow `pulumi:"maintenanceWindow"`
	// The replication role.
	ReplicationRole *string `pulumi:"replicationRole"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// The name of the server.
	ServerName *string `pulumi:"serverName"`
	// The SKU (pricing tier) of the server.
	Sku *Sku `pulumi:"sku"`
	// The source MySQL server id.
	SourceServerId *string `pulumi:"sourceServerId"`
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement *string `pulumi:"sslEnforcement"`
	// Storage profile of a server.
	StorageProfile *StorageProfile `pulumi:"storageProfile"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Server version.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin pulumi.StringPtrInput
	// The password of the administrator login (required for server creation).
	AdministratorLoginPassword pulumi.StringPtrInput
	// availability Zone information of the server.
	AvailabilityZone pulumi.StringPtrInput
	// The mode to create a new MySQL server.
	CreateMode pulumi.StringPtrInput
	// Delegated subnet arguments.
	DelegatedSubnetArguments DelegatedSubnetArgumentsPtrInput
	// Enable HA or not for a server.
	HaEnabled pulumi.StringPtrInput
	// The Azure Active Directory identity of the server.
	Identity IdentityPtrInput
	// Status showing whether the server enabled infrastructure encryption.
	InfrastructureEncryption pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Maintenance window of a server.
	MaintenanceWindow MaintenanceWindowPtrInput
	// The replication role.
	ReplicationRole pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime pulumi.StringPtrInput
	// The name of the server.
	ServerName pulumi.StringPtrInput
	// The SKU (pricing tier) of the server.
	Sku SkuPtrInput
	// The source MySQL server id.
	SourceServerId pulumi.StringPtrInput
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement pulumi.StringPtrInput
	// Storage profile of a server.
	StorageProfile StorageProfilePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Server version.
	Version pulumi.StringPtrInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}

type ServerInput interface {
	pulumi.Input

	ToServerOutput() ServerOutput
	ToServerOutputWithContext(ctx context.Context) ServerOutput
}

func (*Server) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (i *Server) ToServerOutput() ServerOutput {
	return i.ToServerOutputWithContext(context.Background())
}

func (i *Server) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerOutput)
}

type ServerOutput struct{ *pulumi.OutputState }

func (ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (o ServerOutput) ToServerOutput() ServerOutput {
	return o
}

func (o ServerOutput) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return o
}

// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
func (o ServerOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

// availability Zone information of the server.
func (o ServerOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// Status showing whether the data encryption is enabled with customer-managed keys.
func (o ServerOutput) ByokEnforcement() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.ByokEnforcement }).(pulumi.StringOutput)
}

// Delegated subnet arguments.
func (o ServerOutput) DelegatedSubnetArguments() DelegatedSubnetArgumentsResponsePtrOutput {
	return o.ApplyT(func(v *Server) DelegatedSubnetArgumentsResponsePtrOutput { return v.DelegatedSubnetArguments }).(DelegatedSubnetArgumentsResponsePtrOutput)
}

// Earliest restore point creation time (ISO8601 format)
func (o ServerOutput) EarliestRestoreDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.EarliestRestoreDate }).(pulumi.StringOutput)
}

// The fully qualified domain name of a server.
func (o ServerOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

// Enable HA or not for a server.
func (o ServerOutput) HaEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.HaEnabled }).(pulumi.StringPtrOutput)
}

// The state of a HA server.
func (o ServerOutput) HaState() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.HaState }).(pulumi.StringOutput)
}

// The Azure Active Directory identity of the server.
func (o ServerOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Server) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o ServerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Maintenance window of a server.
func (o ServerOutput) MaintenanceWindow() MaintenanceWindowResponsePtrOutput {
	return o.ApplyT(func(v *Server) MaintenanceWindowResponsePtrOutput { return v.MaintenanceWindow }).(MaintenanceWindowResponsePtrOutput)
}

// The name of the resource
func (o ServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
func (o ServerOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

// The maximum number of replicas that a primary server can have.
func (o ServerOutput) ReplicaCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *Server) pulumi.IntOutput { return v.ReplicaCapacity }).(pulumi.IntOutput)
}

// The replication role.
func (o ServerOutput) ReplicationRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.ReplicationRole }).(pulumi.StringPtrOutput)
}

// The SKU (pricing tier) of the server.
func (o ServerOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Server) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The source MySQL server id.
func (o ServerOutput) SourceServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SourceServerId }).(pulumi.StringPtrOutput)
}

// Enable ssl enforcement or not when connect to server.
func (o ServerOutput) SslEnforcement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SslEnforcement }).(pulumi.StringPtrOutput)
}

// availability Zone information of the server.
func (o ServerOutput) StandbyAvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.StandbyAvailabilityZone }).(pulumi.StringOutput)
}

// The state of a server.
func (o ServerOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Storage profile of a server.
func (o ServerOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *Server) StorageProfileResponsePtrOutput { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

// Resource tags.
func (o ServerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Server) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Server version.
func (o ServerOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerOutput{})
}
