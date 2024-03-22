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
	// availability Zone information of the server.
	AvailabilityZone pulumi.StringPtrOutput `pulumi:"availabilityZone"`
	// Backup related properties of a server.
	Backup BackupResponsePtrOutput `pulumi:"backup"`
	// The Data Encryption for CMK.
	DataEncryption DataEncryptionResponsePtrOutput `pulumi:"dataEncryption"`
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName pulumi.StringOutput `pulumi:"fullyQualifiedDomainName"`
	// High availability related properties of a server.
	HighAvailability HighAvailabilityResponsePtrOutput `pulumi:"highAvailability"`
	// The cmk identity for the server.
	Identity MySQLServerIdentityResponsePtrOutput `pulumi:"identity"`
	// Source properties for import from storage.
	ImportSourceProperties ImportSourcePropertiesResponsePtrOutput `pulumi:"importSourceProperties"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Maintenance window of a server.
	MaintenanceWindow MaintenanceWindowResponsePtrOutput `pulumi:"maintenanceWindow"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Network related properties of a server.
	Network NetworkResponsePtrOutput `pulumi:"network"`
	// PrivateEndpointConnections related properties of a server.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The maximum number of replicas that a primary server can have.
	ReplicaCapacity pulumi.IntOutput `pulumi:"replicaCapacity"`
	// The replication role.
	ReplicationRole pulumi.StringPtrOutput `pulumi:"replicationRole"`
	// The SKU (pricing tier) of the server.
	Sku MySQLServerSkuResponsePtrOutput `pulumi:"sku"`
	// The source MySQL server id.
	SourceServerResourceId pulumi.StringPtrOutput `pulumi:"sourceServerResourceId"`
	// The state of a server.
	State pulumi.StringOutput `pulumi:"state"`
	// Storage related properties of a server.
	Storage StorageResponsePtrOutput `pulumi:"storage"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
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
	if args.Backup != nil {
		args.Backup = args.Backup.ToBackupPtrOutput().ApplyT(func(v *Backup) *Backup { return v.Defaults() }).(BackupPtrOutput)
	}
	if args.Storage != nil {
		args.Storage = args.Storage.ToStoragePtrOutput().ApplyT(func(v *Storage) *Storage { return v.Defaults() }).(StoragePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dbformysql:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20200701preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20200701privatepreview:Server"),
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
		{
			Type: pulumi.String("azure-native:dbformysql/v20231230:Server"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Server
	err := ctx.RegisterResource("azure-native:dbformysql/v20231201preview:Server", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:dbformysql/v20231201preview:Server", name, id, state, &resource, opts...)
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
	// Backup related properties of a server.
	Backup *Backup `pulumi:"backup"`
	// The mode to create a new MySQL server.
	CreateMode *string `pulumi:"createMode"`
	// The Data Encryption for CMK.
	DataEncryption *DataEncryption `pulumi:"dataEncryption"`
	// High availability related properties of a server.
	HighAvailability *HighAvailability `pulumi:"highAvailability"`
	// The cmk identity for the server.
	Identity *MySQLServerIdentity `pulumi:"identity"`
	// Source properties for import from storage.
	ImportSourceProperties *ImportSourceProperties `pulumi:"importSourceProperties"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindow `pulumi:"maintenanceWindow"`
	// Network related properties of a server.
	Network *Network `pulumi:"network"`
	// The replication role.
	ReplicationRole *string `pulumi:"replicationRole"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// The name of the server.
	ServerName *string `pulumi:"serverName"`
	// The SKU (pricing tier) of the server.
	Sku *MySQLServerSku `pulumi:"sku"`
	// The source MySQL server id.
	SourceServerResourceId *string `pulumi:"sourceServerResourceId"`
	// Storage related properties of a server.
	Storage *Storage `pulumi:"storage"`
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
	// Backup related properties of a server.
	Backup BackupPtrInput
	// The mode to create a new MySQL server.
	CreateMode pulumi.StringPtrInput
	// The Data Encryption for CMK.
	DataEncryption DataEncryptionPtrInput
	// High availability related properties of a server.
	HighAvailability HighAvailabilityPtrInput
	// The cmk identity for the server.
	Identity MySQLServerIdentityPtrInput
	// Source properties for import from storage.
	ImportSourceProperties ImportSourcePropertiesPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Maintenance window of a server.
	MaintenanceWindow MaintenanceWindowPtrInput
	// Network related properties of a server.
	Network NetworkPtrInput
	// The replication role.
	ReplicationRole pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime pulumi.StringPtrInput
	// The name of the server.
	ServerName pulumi.StringPtrInput
	// The SKU (pricing tier) of the server.
	Sku MySQLServerSkuPtrInput
	// The source MySQL server id.
	SourceServerResourceId pulumi.StringPtrInput
	// Storage related properties of a server.
	Storage StoragePtrInput
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

// Backup related properties of a server.
func (o ServerOutput) Backup() BackupResponsePtrOutput {
	return o.ApplyT(func(v *Server) BackupResponsePtrOutput { return v.Backup }).(BackupResponsePtrOutput)
}

// The Data Encryption for CMK.
func (o ServerOutput) DataEncryption() DataEncryptionResponsePtrOutput {
	return o.ApplyT(func(v *Server) DataEncryptionResponsePtrOutput { return v.DataEncryption }).(DataEncryptionResponsePtrOutput)
}

// The fully qualified domain name of a server.
func (o ServerOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

// High availability related properties of a server.
func (o ServerOutput) HighAvailability() HighAvailabilityResponsePtrOutput {
	return o.ApplyT(func(v *Server) HighAvailabilityResponsePtrOutput { return v.HighAvailability }).(HighAvailabilityResponsePtrOutput)
}

// The cmk identity for the server.
func (o ServerOutput) Identity() MySQLServerIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Server) MySQLServerIdentityResponsePtrOutput { return v.Identity }).(MySQLServerIdentityResponsePtrOutput)
}

// Source properties for import from storage.
func (o ServerOutput) ImportSourceProperties() ImportSourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Server) ImportSourcePropertiesResponsePtrOutput { return v.ImportSourceProperties }).(ImportSourcePropertiesResponsePtrOutput)
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

// Network related properties of a server.
func (o ServerOutput) Network() NetworkResponsePtrOutput {
	return o.ApplyT(func(v *Server) NetworkResponsePtrOutput { return v.Network }).(NetworkResponsePtrOutput)
}

// PrivateEndpointConnections related properties of a server.
func (o ServerOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Server) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
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
func (o ServerOutput) Sku() MySQLServerSkuResponsePtrOutput {
	return o.ApplyT(func(v *Server) MySQLServerSkuResponsePtrOutput { return v.Sku }).(MySQLServerSkuResponsePtrOutput)
}

// The source MySQL server id.
func (o ServerOutput) SourceServerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SourceServerResourceId }).(pulumi.StringPtrOutput)
}

// The state of a server.
func (o ServerOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Storage related properties of a server.
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

// Server version.
func (o ServerOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerOutput{})
}
