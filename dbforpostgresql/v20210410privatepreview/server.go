// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210410privatepreview

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
	ByokEnforcement          pulumi.StringOutput                                       `pulumi:"byokEnforcement"`
	DelegatedSubnetArguments ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput `pulumi:"delegatedSubnetArguments"`
	// The display name of a server.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName pulumi.StringOutput `pulumi:"fullyQualifiedDomainName"`
	// stand by count value can be either enabled or disabled
	HaEnabled pulumi.StringPtrOutput `pulumi:"haEnabled"`
	// A state of a HA server that is visible to user.
	HaState pulumi.StringOutput `pulumi:"haState"`
	// The Azure Active Directory identity of the server.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Maintenance window of a server.
	MaintenanceWindow MaintenanceWindowResponsePtrOutput `pulumi:"maintenanceWindow"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	PointInTimeUTC          pulumi.StringPtrOutput                                   `pulumi:"pointInTimeUTC"`
	PrivateDnsZoneArguments ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput `pulumi:"privateDnsZoneArguments"`
	// public network access is enabled or not
	PublicNetworkAccess pulumi.StringOutput `pulumi:"publicNetworkAccess"`
	// The SKU (pricing tier) of the server.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The resource group name of source PostgreSQL server name to restore from.
	SourceResourceGroupName pulumi.StringPtrOutput `pulumi:"sourceResourceGroupName"`
	// The source PostgreSQL server name to restore from.
	SourceServerName pulumi.StringPtrOutput `pulumi:"sourceServerName"`
	// The subscription id of source PostgreSQL server name to restore from.
	SourceSubscriptionId pulumi.StringPtrOutput `pulumi:"sourceSubscriptionId"`
	// availability Zone information of the server.
	StandbyAvailabilityZone pulumi.StringOutput `pulumi:"standbyAvailabilityZone"`
	// A state of a server that is visible to user.
	State pulumi.StringOutput `pulumi:"state"`
	// Storage profile of a server.
	StorageProfile StorageProfileResponsePtrOutput `pulumi:"storageProfile"`
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
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20231201preview:Server"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Server
	err := ctx.RegisterResource("azure-native:dbforpostgresql/v20210410privatepreview:Server", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:dbforpostgresql/v20210410privatepreview:Server", name, id, state, &resource, opts...)
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
	// availability Zone information of the server.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The mode to create a new PostgreSQL server.
	CreateMode               *string                                   `pulumi:"createMode"`
	DelegatedSubnetArguments *ServerPropertiesDelegatedSubnetArguments `pulumi:"delegatedSubnetArguments"`
	// The display name of a server.
	DisplayName *string `pulumi:"displayName"`
	// stand by count value can be either enabled or disabled
	HaEnabled *HAEnabledEnum `pulumi:"haEnabled"`
	// The Azure Active Directory identity of the server.
	Identity *Identity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindow `pulumi:"maintenanceWindow"`
	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	PointInTimeUTC          *string                                  `pulumi:"pointInTimeUTC"`
	PrivateDnsZoneArguments *ServerPropertiesPrivateDnsZoneArguments `pulumi:"privateDnsZoneArguments"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName *string `pulumi:"serverName"`
	// The SKU (pricing tier) of the server.
	Sku *Sku `pulumi:"sku"`
	// The resource group name of source PostgreSQL server name to restore from.
	SourceResourceGroupName *string `pulumi:"sourceResourceGroupName"`
	// The source PostgreSQL server name to restore from.
	SourceServerName *string `pulumi:"sourceServerName"`
	// The subscription id of source PostgreSQL server name to restore from.
	SourceSubscriptionId *string `pulumi:"sourceSubscriptionId"`
	// Storage profile of a server.
	StorageProfile *StorageProfile `pulumi:"storageProfile"`
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
	// availability Zone information of the server.
	AvailabilityZone pulumi.StringPtrInput
	// The mode to create a new PostgreSQL server.
	CreateMode               pulumi.StringPtrInput
	DelegatedSubnetArguments ServerPropertiesDelegatedSubnetArgumentsPtrInput
	// The display name of a server.
	DisplayName pulumi.StringPtrInput
	// stand by count value can be either enabled or disabled
	HaEnabled HAEnabledEnumPtrInput
	// The Azure Active Directory identity of the server.
	Identity IdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Maintenance window of a server.
	MaintenanceWindow MaintenanceWindowPtrInput
	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	PointInTimeUTC          pulumi.StringPtrInput
	PrivateDnsZoneArguments ServerPropertiesPrivateDnsZoneArgumentsPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringPtrInput
	// The SKU (pricing tier) of the server.
	Sku SkuPtrInput
	// The resource group name of source PostgreSQL server name to restore from.
	SourceResourceGroupName pulumi.StringPtrInput
	// The source PostgreSQL server name to restore from.
	SourceServerName pulumi.StringPtrInput
	// The subscription id of source PostgreSQL server name to restore from.
	SourceSubscriptionId pulumi.StringPtrInput
	// Storage profile of a server.
	StorageProfile StorageProfilePtrInput
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

// availability Zone information of the server.
func (o ServerOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// Status showing whether the data encryption is enabled with customer-managed keys.
func (o ServerOutput) ByokEnforcement() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.ByokEnforcement }).(pulumi.StringOutput)
}

func (o ServerOutput) DelegatedSubnetArguments() ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput {
	return o.ApplyT(func(v *Server) ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput {
		return v.DelegatedSubnetArguments
	}).(ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput)
}

// The display name of a server.
func (o ServerOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The fully qualified domain name of a server.
func (o ServerOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

// stand by count value can be either enabled or disabled
func (o ServerOutput) HaEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.HaEnabled }).(pulumi.StringPtrOutput)
}

// A state of a HA server that is visible to user.
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

// Restore point creation time (ISO8601 format), specifying the time to restore from.
func (o ServerOutput) PointInTimeUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.PointInTimeUTC }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) PrivateDnsZoneArguments() ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput {
	return o.ApplyT(func(v *Server) ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput {
		return v.PrivateDnsZoneArguments
	}).(ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput)
}

// public network access is enabled or not
func (o ServerOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

// The SKU (pricing tier) of the server.
func (o ServerOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Server) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The resource group name of source PostgreSQL server name to restore from.
func (o ServerOutput) SourceResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SourceResourceGroupName }).(pulumi.StringPtrOutput)
}

// The source PostgreSQL server name to restore from.
func (o ServerOutput) SourceServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SourceServerName }).(pulumi.StringPtrOutput)
}

// The subscription id of source PostgreSQL server name to restore from.
func (o ServerOutput) SourceSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SourceSubscriptionId }).(pulumi.StringPtrOutput)
}

// availability Zone information of the server.
func (o ServerOutput) StandbyAvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.StandbyAvailabilityZone }).(pulumi.StringOutput)
}

// A state of a server that is visible to user.
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

// PostgreSQL Server version.
func (o ServerOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerOutput{})
}
