// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231201preview

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
	// AuthConfig properties of a server.
	AuthConfig AuthConfigResponsePtrOutput `pulumi:"authConfig"`
	// availability zone information of the server.
	AvailabilityZone pulumi.StringPtrOutput `pulumi:"availabilityZone"`
	// Backup properties of a server.
	Backup BackupResponsePtrOutput `pulumi:"backup"`
	// Data encryption properties of a server.
	DataEncryption DataEncryptionResponsePtrOutput `pulumi:"dataEncryption"`
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName pulumi.StringOutput `pulumi:"fullyQualifiedDomainName"`
	// High availability properties of a server.
	HighAvailability HighAvailabilityResponsePtrOutput `pulumi:"highAvailability"`
	// Describes the identity of the application.
	Identity UserAssignedIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Maintenance window properties of a server.
	MaintenanceWindow MaintenanceWindowResponsePtrOutput `pulumi:"maintenanceWindow"`
	// The minor version of the server.
	MinorVersion pulumi.StringOutput `pulumi:"minorVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Network properties of a server. This Network property is required to be passed only in case you want the server to be Private access server.
	Network NetworkResponsePtrOutput `pulumi:"network"`
	// List of private endpoint connections associated with the specified resource.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Replica properties of a server. These Replica properties are required to be passed only in case you want to Promote a server.
	Replica ReplicaResponsePtrOutput `pulumi:"replica"`
	// Replicas allowed for a server.
	ReplicaCapacity pulumi.IntOutput `pulumi:"replicaCapacity"`
	// Replication role of the server
	ReplicationRole pulumi.StringPtrOutput `pulumi:"replicationRole"`
	// The SKU (pricing tier) of the server.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The source server resource ID to restore from. It's required when 'createMode' is 'PointInTimeRestore' or 'GeoRestore' or 'Replica' or 'ReviveDropped'. This property is returned only for Replica server
	SourceServerResourceId pulumi.StringPtrOutput `pulumi:"sourceServerResourceId"`
	// A state of a server that is visible to user.
	State pulumi.StringOutput `pulumi:"state"`
	// Storage properties of a server.
	Storage StorageResponsePtrOutput `pulumi:"storage"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// PostgreSQL Server version.
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
	if args.AuthConfig != nil {
		args.AuthConfig = args.AuthConfig.ToAuthConfigPtrOutput().ApplyT(func(v *AuthConfig) *AuthConfig { return v.Defaults() }).(AuthConfigPtrOutput)
	}
	if args.AvailabilityZone == nil {
		args.AvailabilityZone = pulumi.StringPtr("")
	}
	if args.Backup != nil {
		args.Backup = args.Backup.ToBackupPtrOutput().ApplyT(func(v *Backup) *Backup { return v.Defaults() }).(BackupPtrOutput)
	}
	if args.HighAvailability != nil {
		args.HighAvailability = args.HighAvailability.ToHighAvailabilityPtrOutput().ApplyT(func(v *HighAvailability) *HighAvailability { return v.Defaults() }).(HighAvailabilityPtrOutput)
	}
	if args.MaintenanceWindow != nil {
		args.MaintenanceWindow = args.MaintenanceWindow.ToMaintenanceWindowPtrOutput().ApplyT(func(v *MaintenanceWindow) *MaintenanceWindow { return v.Defaults() }).(MaintenanceWindowPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dbforpostgresql:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20200214preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20200214privatepreview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210410privatepreview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210601:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210601preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210615privatepreview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20220120preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20220308preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20221201:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20230301preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20230601preview:Server"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Server
	err := ctx.RegisterResource("azure-native:dbforpostgresql/v20231201preview:Server", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:dbforpostgresql/v20231201preview:Server", name, id, state, &resource, opts...)
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
	// The administrator login password (required for server creation).
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// AuthConfig properties of a server.
	AuthConfig *AuthConfig `pulumi:"authConfig"`
	// availability zone information of the server.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Backup properties of a server.
	Backup *Backup `pulumi:"backup"`
	// The mode to create a new PostgreSQL server.
	CreateMode *string `pulumi:"createMode"`
	// Data encryption properties of a server.
	DataEncryption *DataEncryption `pulumi:"dataEncryption"`
	// High availability properties of a server.
	HighAvailability *HighAvailability `pulumi:"highAvailability"`
	// Describes the identity of the application.
	Identity *UserAssignedIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Maintenance window properties of a server.
	MaintenanceWindow *MaintenanceWindow `pulumi:"maintenanceWindow"`
	// Network properties of a server. This Network property is required to be passed only in case you want the server to be Private access server.
	Network *Network `pulumi:"network"`
	// Restore point creation time (ISO8601 format), specifying the time to restore from. It's required when 'createMode' is 'PointInTimeRestore' or 'GeoRestore' or 'ReviveDropped'.
	PointInTimeUTC *string `pulumi:"pointInTimeUTC"`
	// Replica properties of a server. These Replica properties are required to be passed only in case you want to Promote a server.
	Replica *Replica `pulumi:"replica"`
	// Replication role of the server
	ReplicationRole *string `pulumi:"replicationRole"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName *string `pulumi:"serverName"`
	// The SKU (pricing tier) of the server.
	Sku *Sku `pulumi:"sku"`
	// The source server resource ID to restore from. It's required when 'createMode' is 'PointInTimeRestore' or 'GeoRestore' or 'Replica' or 'ReviveDropped'. This property is returned only for Replica server
	SourceServerResourceId *string `pulumi:"sourceServerResourceId"`
	// Storage properties of a server.
	Storage *Storage `pulumi:"storage"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// PostgreSQL Server version.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin pulumi.StringPtrInput
	// The administrator login password (required for server creation).
	AdministratorLoginPassword pulumi.StringPtrInput
	// AuthConfig properties of a server.
	AuthConfig AuthConfigPtrInput
	// availability zone information of the server.
	AvailabilityZone pulumi.StringPtrInput
	// Backup properties of a server.
	Backup BackupPtrInput
	// The mode to create a new PostgreSQL server.
	CreateMode pulumi.StringPtrInput
	// Data encryption properties of a server.
	DataEncryption DataEncryptionPtrInput
	// High availability properties of a server.
	HighAvailability HighAvailabilityPtrInput
	// Describes the identity of the application.
	Identity UserAssignedIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Maintenance window properties of a server.
	MaintenanceWindow MaintenanceWindowPtrInput
	// Network properties of a server. This Network property is required to be passed only in case you want the server to be Private access server.
	Network NetworkPtrInput
	// Restore point creation time (ISO8601 format), specifying the time to restore from. It's required when 'createMode' is 'PointInTimeRestore' or 'GeoRestore' or 'ReviveDropped'.
	PointInTimeUTC pulumi.StringPtrInput
	// Replica properties of a server. These Replica properties are required to be passed only in case you want to Promote a server.
	Replica ReplicaPtrInput
	// Replication role of the server
	ReplicationRole pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringPtrInput
	// The SKU (pricing tier) of the server.
	Sku SkuPtrInput
	// The source server resource ID to restore from. It's required when 'createMode' is 'PointInTimeRestore' or 'GeoRestore' or 'Replica' or 'ReviveDropped'. This property is returned only for Replica server
	SourceServerResourceId pulumi.StringPtrInput
	// Storage properties of a server.
	Storage StoragePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// PostgreSQL Server version.
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

// AuthConfig properties of a server.
func (o ServerOutput) AuthConfig() AuthConfigResponsePtrOutput {
	return o.ApplyT(func(v *Server) AuthConfigResponsePtrOutput { return v.AuthConfig }).(AuthConfigResponsePtrOutput)
}

// availability zone information of the server.
func (o ServerOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// Backup properties of a server.
func (o ServerOutput) Backup() BackupResponsePtrOutput {
	return o.ApplyT(func(v *Server) BackupResponsePtrOutput { return v.Backup }).(BackupResponsePtrOutput)
}

// Data encryption properties of a server.
func (o ServerOutput) DataEncryption() DataEncryptionResponsePtrOutput {
	return o.ApplyT(func(v *Server) DataEncryptionResponsePtrOutput { return v.DataEncryption }).(DataEncryptionResponsePtrOutput)
}

// The fully qualified domain name of a server.
func (o ServerOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

// High availability properties of a server.
func (o ServerOutput) HighAvailability() HighAvailabilityResponsePtrOutput {
	return o.ApplyT(func(v *Server) HighAvailabilityResponsePtrOutput { return v.HighAvailability }).(HighAvailabilityResponsePtrOutput)
}

// Describes the identity of the application.
func (o ServerOutput) Identity() UserAssignedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Server) UserAssignedIdentityResponsePtrOutput { return v.Identity }).(UserAssignedIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o ServerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Maintenance window properties of a server.
func (o ServerOutput) MaintenanceWindow() MaintenanceWindowResponsePtrOutput {
	return o.ApplyT(func(v *Server) MaintenanceWindowResponsePtrOutput { return v.MaintenanceWindow }).(MaintenanceWindowResponsePtrOutput)
}

// The minor version of the server.
func (o ServerOutput) MinorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.MinorVersion }).(pulumi.StringOutput)
}

// The name of the resource
func (o ServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network properties of a server. This Network property is required to be passed only in case you want the server to be Private access server.
func (o ServerOutput) Network() NetworkResponsePtrOutput {
	return o.ApplyT(func(v *Server) NetworkResponsePtrOutput { return v.Network }).(NetworkResponsePtrOutput)
}

// List of private endpoint connections associated with the specified resource.
func (o ServerOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Server) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Replica properties of a server. These Replica properties are required to be passed only in case you want to Promote a server.
func (o ServerOutput) Replica() ReplicaResponsePtrOutput {
	return o.ApplyT(func(v *Server) ReplicaResponsePtrOutput { return v.Replica }).(ReplicaResponsePtrOutput)
}

// Replicas allowed for a server.
func (o ServerOutput) ReplicaCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *Server) pulumi.IntOutput { return v.ReplicaCapacity }).(pulumi.IntOutput)
}

// Replication role of the server
func (o ServerOutput) ReplicationRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.ReplicationRole }).(pulumi.StringPtrOutput)
}

// The SKU (pricing tier) of the server.
func (o ServerOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Server) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The source server resource ID to restore from. It's required when 'createMode' is 'PointInTimeRestore' or 'GeoRestore' or 'Replica' or 'ReviveDropped'. This property is returned only for Replica server
func (o ServerOutput) SourceServerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SourceServerResourceId }).(pulumi.StringPtrOutput)
}

// A state of a server that is visible to user.
func (o ServerOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Storage properties of a server.
func (o ServerOutput) Storage() StorageResponsePtrOutput {
	return o.ApplyT(func(v *Server) StorageResponsePtrOutput { return v.Storage }).(StorageResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ServerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Server) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ServerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Server) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// PostgreSQL Server version.
func (o ServerOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerOutput{})
}
