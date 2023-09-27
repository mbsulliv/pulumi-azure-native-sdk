// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An Azure SQL Database server.
type Server struct {
	pulumi.CustomResourceState

	// Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin pulumi.StringPtrOutput `pulumi:"administratorLogin"`
	// The Azure Active Directory administrator of the server. This can only be used at server create time. If used for server update, it will be ignored or it will result in an error. For updates individual APIs will need to be used.
	Administrators ServerExternalAdministratorResponsePtrOutput `pulumi:"administrators"`
	// Status of external governance.
	ExternalGovernanceStatus pulumi.StringOutput `pulumi:"externalGovernanceStatus"`
	// The Client id used for cross tenant CMK scenario
	FederatedClientId pulumi.StringPtrOutput `pulumi:"federatedClientId"`
	// The fully qualified domain name of the server.
	FullyQualifiedDomainName pulumi.StringOutput `pulumi:"fullyQualifiedDomainName"`
	// The Azure Active Directory identity of the server.
	Identity ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	// Whether or not to enable IPv6 support for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	IsIPv6Enabled pulumi.StringPtrOutput `pulumi:"isIPv6Enabled"`
	// A CMK URI of the key to use for encryption.
	KeyId pulumi.StringPtrOutput `pulumi:"keyId"`
	// Kind of sql server. This is metadata used for the Azure portal experience.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Minimal TLS version. Allowed values: 'None', '1.0', '1.1', '1.2'
	MinimalTlsVersion pulumi.StringPtrOutput `pulumi:"minimalTlsVersion"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityId pulumi.StringPtrOutput `pulumi:"primaryUserAssignedIdentityId"`
	// List of private endpoint connections on a server
	PrivateEndpointConnections ServerPrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled' or 'SecuredByPerimeter'
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Whether or not to restrict outbound network access for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	RestrictOutboundNetworkAccess pulumi.StringPtrOutput `pulumi:"restrictOutboundNetworkAccess"`
	// The state of the server.
	State pulumi.StringOutput `pulumi:"state"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The version of the server.
	Version pulumi.StringPtrOutput `pulumi:"version"`
	// Whether or not existing server has a workspace created and if it allows connection from workspace
	WorkspaceFeature pulumi.StringOutput `pulumi:"workspaceFeature"`
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
			Type: pulumi.String("azure-native:sql:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:Server"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Server
	err := ctx.RegisterResource("azure-native:sql/v20230201preview:Server", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:sql/v20230201preview:Server", name, id, state, &resource, opts...)
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
	// Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The administrator login password (required for server creation).
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// The Azure Active Directory administrator of the server. This can only be used at server create time. If used for server update, it will be ignored or it will result in an error. For updates individual APIs will need to be used.
	Administrators *ServerExternalAdministrator `pulumi:"administrators"`
	// The Client id used for cross tenant CMK scenario
	FederatedClientId *string `pulumi:"federatedClientId"`
	// The Azure Active Directory identity of the server.
	Identity *ResourceIdentity `pulumi:"identity"`
	// Whether or not to enable IPv6 support for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	IsIPv6Enabled *string `pulumi:"isIPv6Enabled"`
	// A CMK URI of the key to use for encryption.
	KeyId *string `pulumi:"keyId"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Minimal TLS version. Allowed values: 'None', '1.0', '1.1', '1.2'
	MinimalTlsVersion *string `pulumi:"minimalTlsVersion"`
	// The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityId *string `pulumi:"primaryUserAssignedIdentityId"`
	// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled' or 'SecuredByPerimeter'
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Whether or not to restrict outbound network access for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	RestrictOutboundNetworkAccess *string `pulumi:"restrictOutboundNetworkAccess"`
	// The name of the server.
	ServerName *string `pulumi:"serverName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The version of the server.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin pulumi.StringPtrInput
	// The administrator login password (required for server creation).
	AdministratorLoginPassword pulumi.StringPtrInput
	// The Azure Active Directory administrator of the server. This can only be used at server create time. If used for server update, it will be ignored or it will result in an error. For updates individual APIs will need to be used.
	Administrators ServerExternalAdministratorPtrInput
	// The Client id used for cross tenant CMK scenario
	FederatedClientId pulumi.StringPtrInput
	// The Azure Active Directory identity of the server.
	Identity ResourceIdentityPtrInput
	// Whether or not to enable IPv6 support for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	IsIPv6Enabled pulumi.StringPtrInput
	// A CMK URI of the key to use for encryption.
	KeyId pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Minimal TLS version. Allowed values: 'None', '1.0', '1.1', '1.2'
	MinimalTlsVersion pulumi.StringPtrInput
	// The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityId pulumi.StringPtrInput
	// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled' or 'SecuredByPerimeter'
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Whether or not to restrict outbound network access for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	RestrictOutboundNetworkAccess pulumi.StringPtrInput
	// The name of the server.
	ServerName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The version of the server.
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

func (i *Server) ToOutput(ctx context.Context) pulumix.Output[*Server] {
	return pulumix.Output[*Server]{
		OutputState: i.ToServerOutputWithContext(ctx).OutputState,
	}
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

func (o ServerOutput) ToOutput(ctx context.Context) pulumix.Output[*Server] {
	return pulumix.Output[*Server]{
		OutputState: o.OutputState,
	}
}

// Administrator username for the server. Once created it cannot be changed.
func (o ServerOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

// The Azure Active Directory administrator of the server. This can only be used at server create time. If used for server update, it will be ignored or it will result in an error. For updates individual APIs will need to be used.
func (o ServerOutput) Administrators() ServerExternalAdministratorResponsePtrOutput {
	return o.ApplyT(func(v *Server) ServerExternalAdministratorResponsePtrOutput { return v.Administrators }).(ServerExternalAdministratorResponsePtrOutput)
}

// Status of external governance.
func (o ServerOutput) ExternalGovernanceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.ExternalGovernanceStatus }).(pulumi.StringOutput)
}

// The Client id used for cross tenant CMK scenario
func (o ServerOutput) FederatedClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.FederatedClientId }).(pulumi.StringPtrOutput)
}

// The fully qualified domain name of the server.
func (o ServerOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

// The Azure Active Directory identity of the server.
func (o ServerOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Server) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// Whether or not to enable IPv6 support for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
func (o ServerOutput) IsIPv6Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.IsIPv6Enabled }).(pulumi.StringPtrOutput)
}

// A CMK URI of the key to use for encryption.
func (o ServerOutput) KeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.KeyId }).(pulumi.StringPtrOutput)
}

// Kind of sql server. This is metadata used for the Azure portal experience.
func (o ServerOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o ServerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Minimal TLS version. Allowed values: 'None', '1.0', '1.1', '1.2'
func (o ServerOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource id of a user assigned identity to be used by default.
func (o ServerOutput) PrimaryUserAssignedIdentityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.PrimaryUserAssignedIdentityId }).(pulumi.StringPtrOutput)
}

// List of private endpoint connections on a server
func (o ServerOutput) PrivateEndpointConnections() ServerPrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Server) ServerPrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(ServerPrivateEndpointConnectionResponseArrayOutput)
}

// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled' or 'SecuredByPerimeter'
func (o ServerOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Whether or not to restrict outbound network access for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
func (o ServerOutput) RestrictOutboundNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.RestrictOutboundNetworkAccess }).(pulumi.StringPtrOutput)
}

// The state of the server.
func (o ServerOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Resource tags.
func (o ServerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Server) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The version of the server.
func (o ServerOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

// Whether or not existing server has a workspace created and if it allows connection from workspace
func (o ServerOutput) WorkspaceFeature() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.WorkspaceFeature }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerOutput{})
}
