// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a server.
func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:sql/v20230201preview:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerArgs struct {
	// The child resources to include in the response.
	Expand *string `pulumi:"expand"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// An Azure SQL Database server.
type LookupServerResult struct {
	// Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The Azure Active Directory administrator of the server. This can only be used at server create time. If used for server update, it will be ignored or it will result in an error. For updates individual APIs will need to be used.
	Administrators *ServerExternalAdministratorResponse `pulumi:"administrators"`
	// Status of external governance.
	ExternalGovernanceStatus string `pulumi:"externalGovernanceStatus"`
	// The Client id used for cross tenant CMK scenario
	FederatedClientId *string `pulumi:"federatedClientId"`
	// The fully qualified domain name of the server.
	FullyQualifiedDomainName string `pulumi:"fullyQualifiedDomainName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The Azure Active Directory identity of the server.
	Identity *ResourceIdentityResponse `pulumi:"identity"`
	// Whether or not to enable IPv6 support for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	IsIPv6Enabled *string `pulumi:"isIPv6Enabled"`
	// A CMK URI of the key to use for encryption.
	KeyId *string `pulumi:"keyId"`
	// Kind of sql server. This is metadata used for the Azure portal experience.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location string `pulumi:"location"`
	// Minimal TLS version. Allowed values: 'None', '1.0', '1.1', '1.2'
	MinimalTlsVersion *string `pulumi:"minimalTlsVersion"`
	// Resource name.
	Name string `pulumi:"name"`
	// The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityId *string `pulumi:"primaryUserAssignedIdentityId"`
	// List of private endpoint connections on a server
	PrivateEndpointConnections []ServerPrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled' or 'SecuredByPerimeter'
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Whether or not to restrict outbound network access for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	RestrictOutboundNetworkAccess *string `pulumi:"restrictOutboundNetworkAccess"`
	// The state of the server.
	State string `pulumi:"state"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The version of the server.
	Version *string `pulumi:"version"`
	// Whether or not existing server has a workspace created and if it allows connection from workspace
	WorkspaceFeature string `pulumi:"workspaceFeature"`
}

func LookupServerOutput(ctx *pulumi.Context, args LookupServerOutputArgs, opts ...pulumi.InvokeOption) LookupServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerResult, error) {
			args := v.(LookupServerArgs)
			r, err := LookupServer(ctx, &args, opts...)
			var s LookupServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerResultOutput)
}

type LookupServerOutputArgs struct {
	// The child resources to include in the response.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerArgs)(nil)).Elem()
}

// An Azure SQL Database server.
type LookupServerResultOutput struct{ *pulumi.OutputState }

func (LookupServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerResult)(nil)).Elem()
}

func (o LookupServerResultOutput) ToLookupServerResultOutput() LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) ToLookupServerResultOutputWithContext(ctx context.Context) LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupServerResult] {
	return pulumix.Output[LookupServerResult]{
		OutputState: o.OutputState,
	}
}

// Administrator username for the server. Once created it cannot be changed.
func (o LookupServerResultOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

// The Azure Active Directory administrator of the server. This can only be used at server create time. If used for server update, it will be ignored or it will result in an error. For updates individual APIs will need to be used.
func (o LookupServerResultOutput) Administrators() ServerExternalAdministratorResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *ServerExternalAdministratorResponse { return v.Administrators }).(ServerExternalAdministratorResponsePtrOutput)
}

// Status of external governance.
func (o LookupServerResultOutput) ExternalGovernanceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.ExternalGovernanceStatus }).(pulumi.StringOutput)
}

// The Client id used for cross tenant CMK scenario
func (o LookupServerResultOutput) FederatedClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.FederatedClientId }).(pulumi.StringPtrOutput)
}

// The fully qualified domain name of the server.
func (o LookupServerResultOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Azure Active Directory identity of the server.
func (o LookupServerResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// Whether or not to enable IPv6 support for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
func (o LookupServerResultOutput) IsIPv6Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.IsIPv6Enabled }).(pulumi.StringPtrOutput)
}

// A CMK URI of the key to use for encryption.
func (o LookupServerResultOutput) KeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.KeyId }).(pulumi.StringPtrOutput)
}

// Kind of sql server. This is metadata used for the Azure portal experience.
func (o LookupServerResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupServerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Location }).(pulumi.StringOutput)
}

// Minimal TLS version. Allowed values: 'None', '1.0', '1.1', '1.2'
func (o LookupServerResultOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource id of a user assigned identity to be used by default.
func (o LookupServerResultOutput) PrimaryUserAssignedIdentityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.PrimaryUserAssignedIdentityId }).(pulumi.StringPtrOutput)
}

// List of private endpoint connections on a server
func (o LookupServerResultOutput) PrivateEndpointConnections() ServerPrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupServerResult) []ServerPrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(ServerPrivateEndpointConnectionResponseArrayOutput)
}

// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled' or 'SecuredByPerimeter'
func (o LookupServerResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Whether or not to restrict outbound network access for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
func (o LookupServerResultOutput) RestrictOutboundNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.RestrictOutboundNetworkAccess }).(pulumi.StringPtrOutput)
}

// The state of the server.
func (o LookupServerResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.State }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupServerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Type }).(pulumi.StringOutput)
}

// The version of the server.
func (o LookupServerResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

// Whether or not existing server has a workspace created and if it allows connection from workspace
func (o LookupServerResultOutput) WorkspaceFeature() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.WorkspaceFeature }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerResultOutput{})
}
