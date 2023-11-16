// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents an instance of an Analysis Services resource.
type ServerDetails struct {
	pulumi.CustomResourceState

	// A collection of AS server administrators
	AsAdministrators ServerAdministratorsResponsePtrOutput `pulumi:"asAdministrators"`
	// The SAS container URI to the backup container.
	BackupBlobContainerUri pulumi.StringPtrOutput `pulumi:"backupBlobContainerUri"`
	// The gateway details configured for the AS server.
	GatewayDetails GatewayDetailsResponsePtrOutput `pulumi:"gatewayDetails"`
	// The firewall settings for the AS server.
	IpV4FirewallSettings IPv4FirewallSettingsResponsePtrOutput `pulumi:"ipV4FirewallSettings"`
	// Location of the Analysis Services resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The managed mode of the server (0 = not managed, 1 = managed).
	ManagedMode pulumi.IntPtrOutput `pulumi:"managedMode"`
	// The name of the Analysis Services resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current deployment state of Analysis Services resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// How the read-write server's participation in the query pool is controlled.<br/>It can have the following values: <ul><li>readOnly - indicates that the read-write server is intended not to participate in query operations</li><li>all - indicates that the read-write server can participate in query operations</li></ul>Specifying readOnly when capacity is 1 results in error.
	QuerypoolConnectionMode pulumi.StringPtrOutput `pulumi:"querypoolConnectionMode"`
	// The full name of the Analysis Services resource.
	ServerFullName pulumi.StringOutput `pulumi:"serverFullName"`
	// The server monitor mode for AS server
	ServerMonitorMode pulumi.IntPtrOutput `pulumi:"serverMonitorMode"`
	// The SKU of the Analysis Services resource.
	Sku ResourceSkuResponseOutput `pulumi:"sku"`
	// The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
	State pulumi.StringOutput `pulumi:"state"`
	// Key-value pairs of additional resource provisioning properties.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the Analysis Services resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServerDetails registers a new resource with the given unique name, arguments, and options.
func NewServerDetails(ctx *pulumi.Context,
	name string, args *ServerDetailsArgs, opts ...pulumi.ResourceOption) (*ServerDetails, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.ManagedMode == nil {
		args.ManagedMode = pulumi.IntPtr(1)
	}
	if args.QuerypoolConnectionMode == nil {
		args.QuerypoolConnectionMode = ConnectionMode("All")
	}
	if args.ServerMonitorMode == nil {
		args.ServerMonitorMode = pulumi.IntPtr(1)
	}
	args.Sku = args.Sku.ToResourceSkuOutput().ApplyT(func(v ResourceSku) ResourceSku { return *v.Defaults() }).(ResourceSkuOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:analysisservices:ServerDetails"),
		},
		{
			Type: pulumi.String("azure-native:analysisservices/v20160516:ServerDetails"),
		},
		{
			Type: pulumi.String("azure-native:analysisservices/v20170714:ServerDetails"),
		},
		{
			Type: pulumi.String("azure-native:analysisservices/v20170801beta:ServerDetails"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServerDetails
	err := ctx.RegisterResource("azure-native:analysisservices/v20170801:ServerDetails", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerDetails gets an existing ServerDetails resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerDetails(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerDetailsState, opts ...pulumi.ResourceOption) (*ServerDetails, error) {
	var resource ServerDetails
	err := ctx.ReadResource("azure-native:analysisservices/v20170801:ServerDetails", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerDetails resources.
type serverDetailsState struct {
}

type ServerDetailsState struct {
}

func (ServerDetailsState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverDetailsState)(nil)).Elem()
}

type serverDetailsArgs struct {
	// A collection of AS server administrators
	AsAdministrators *ServerAdministrators `pulumi:"asAdministrators"`
	// The SAS container URI to the backup container.
	BackupBlobContainerUri *string `pulumi:"backupBlobContainerUri"`
	// The gateway details configured for the AS server.
	GatewayDetails *GatewayDetails `pulumi:"gatewayDetails"`
	// The firewall settings for the AS server.
	IpV4FirewallSettings *IPv4FirewallSettings `pulumi:"ipV4FirewallSettings"`
	// Location of the Analysis Services resource.
	Location *string `pulumi:"location"`
	// The managed mode of the server (0 = not managed, 1 = managed).
	ManagedMode *int `pulumi:"managedMode"`
	// How the read-write server's participation in the query pool is controlled.<br/>It can have the following values: <ul><li>readOnly - indicates that the read-write server is intended not to participate in query operations</li><li>all - indicates that the read-write server can participate in query operations</li></ul>Specifying readOnly when capacity is 1 results in error.
	QuerypoolConnectionMode *ConnectionMode `pulumi:"querypoolConnectionMode"`
	// The name of the Azure Resource group of which a given Analysis Services server is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The server monitor mode for AS server
	ServerMonitorMode *int `pulumi:"serverMonitorMode"`
	// The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
	ServerName *string `pulumi:"serverName"`
	// The SKU of the Analysis Services resource.
	Sku ResourceSku `pulumi:"sku"`
	// Key-value pairs of additional resource provisioning properties.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ServerDetails resource.
type ServerDetailsArgs struct {
	// A collection of AS server administrators
	AsAdministrators ServerAdministratorsPtrInput
	// The SAS container URI to the backup container.
	BackupBlobContainerUri pulumi.StringPtrInput
	// The gateway details configured for the AS server.
	GatewayDetails GatewayDetailsPtrInput
	// The firewall settings for the AS server.
	IpV4FirewallSettings IPv4FirewallSettingsPtrInput
	// Location of the Analysis Services resource.
	Location pulumi.StringPtrInput
	// The managed mode of the server (0 = not managed, 1 = managed).
	ManagedMode pulumi.IntPtrInput
	// How the read-write server's participation in the query pool is controlled.<br/>It can have the following values: <ul><li>readOnly - indicates that the read-write server is intended not to participate in query operations</li><li>all - indicates that the read-write server can participate in query operations</li></ul>Specifying readOnly when capacity is 1 results in error.
	QuerypoolConnectionMode ConnectionModePtrInput
	// The name of the Azure Resource group of which a given Analysis Services server is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName pulumi.StringInput
	// The server monitor mode for AS server
	ServerMonitorMode pulumi.IntPtrInput
	// The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
	ServerName pulumi.StringPtrInput
	// The SKU of the Analysis Services resource.
	Sku ResourceSkuInput
	// Key-value pairs of additional resource provisioning properties.
	Tags pulumi.StringMapInput
}

func (ServerDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverDetailsArgs)(nil)).Elem()
}

type ServerDetailsInput interface {
	pulumi.Input

	ToServerDetailsOutput() ServerDetailsOutput
	ToServerDetailsOutputWithContext(ctx context.Context) ServerDetailsOutput
}

func (*ServerDetails) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerDetails)(nil)).Elem()
}

func (i *ServerDetails) ToServerDetailsOutput() ServerDetailsOutput {
	return i.ToServerDetailsOutputWithContext(context.Background())
}

func (i *ServerDetails) ToServerDetailsOutputWithContext(ctx context.Context) ServerDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerDetailsOutput)
}

type ServerDetailsOutput struct{ *pulumi.OutputState }

func (ServerDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerDetails)(nil)).Elem()
}

func (o ServerDetailsOutput) ToServerDetailsOutput() ServerDetailsOutput {
	return o
}

func (o ServerDetailsOutput) ToServerDetailsOutputWithContext(ctx context.Context) ServerDetailsOutput {
	return o
}

// A collection of AS server administrators
func (o ServerDetailsOutput) AsAdministrators() ServerAdministratorsResponsePtrOutput {
	return o.ApplyT(func(v *ServerDetails) ServerAdministratorsResponsePtrOutput { return v.AsAdministrators }).(ServerAdministratorsResponsePtrOutput)
}

// The SAS container URI to the backup container.
func (o ServerDetailsOutput) BackupBlobContainerUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringPtrOutput { return v.BackupBlobContainerUri }).(pulumi.StringPtrOutput)
}

// The gateway details configured for the AS server.
func (o ServerDetailsOutput) GatewayDetails() GatewayDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ServerDetails) GatewayDetailsResponsePtrOutput { return v.GatewayDetails }).(GatewayDetailsResponsePtrOutput)
}

// The firewall settings for the AS server.
func (o ServerDetailsOutput) IpV4FirewallSettings() IPv4FirewallSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ServerDetails) IPv4FirewallSettingsResponsePtrOutput { return v.IpV4FirewallSettings }).(IPv4FirewallSettingsResponsePtrOutput)
}

// Location of the Analysis Services resource.
func (o ServerDetailsOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The managed mode of the server (0 = not managed, 1 = managed).
func (o ServerDetailsOutput) ManagedMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.IntPtrOutput { return v.ManagedMode }).(pulumi.IntPtrOutput)
}

// The name of the Analysis Services resource.
func (o ServerDetailsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current deployment state of Analysis Services resource. The provisioningState is to indicate states for resource provisioning.
func (o ServerDetailsOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// How the read-write server's participation in the query pool is controlled.<br/>It can have the following values: <ul><li>readOnly - indicates that the read-write server is intended not to participate in query operations</li><li>all - indicates that the read-write server can participate in query operations</li></ul>Specifying readOnly when capacity is 1 results in error.
func (o ServerDetailsOutput) QuerypoolConnectionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringPtrOutput { return v.QuerypoolConnectionMode }).(pulumi.StringPtrOutput)
}

// The full name of the Analysis Services resource.
func (o ServerDetailsOutput) ServerFullName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringOutput { return v.ServerFullName }).(pulumi.StringOutput)
}

// The server monitor mode for AS server
func (o ServerDetailsOutput) ServerMonitorMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.IntPtrOutput { return v.ServerMonitorMode }).(pulumi.IntPtrOutput)
}

// The SKU of the Analysis Services resource.
func (o ServerDetailsOutput) Sku() ResourceSkuResponseOutput {
	return o.ApplyT(func(v *ServerDetails) ResourceSkuResponseOutput { return v.Sku }).(ResourceSkuResponseOutput)
}

// The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
func (o ServerDetailsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Key-value pairs of additional resource provisioning properties.
func (o ServerDetailsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the Analysis Services resource.
func (o ServerDetailsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerDetailsOutput{})
}
