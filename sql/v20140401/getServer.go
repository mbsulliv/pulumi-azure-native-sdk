// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20140401

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
	err := ctx.Invoke("azure-native:sql/v20140401:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Represents a server.
type LookupServerResult struct {
	// Administrator username for the server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The display name of the Azure Active Directory object with admin permissions on this server. Legacy parameter, always null. To check for Active Directory admin, query .../servers/{serverName}/administrators
	ExternalAdministratorLogin string `pulumi:"externalAdministratorLogin"`
	// The ID of the Active Azure Directory object with admin permissions on this server. Legacy parameter, always null. To check for Active Directory admin, query .../servers/{serverName}/administrators.
	ExternalAdministratorSid string `pulumi:"externalAdministratorSid"`
	// The fully qualified domain name of the server.
	FullyQualifiedDomainName string `pulumi:"fullyQualifiedDomainName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Kind of sql server.  This is metadata used for the Azure portal experience.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The state of the server.
	State string `pulumi:"state"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The version of the server.
	Version *string `pulumi:"version"`
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
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerArgs)(nil)).Elem()
}

// Represents a server.
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

// Administrator username for the server. Can only be specified when the server is being created (and is required for creation).
func (o LookupServerResultOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

// The display name of the Azure Active Directory object with admin permissions on this server. Legacy parameter, always null. To check for Active Directory admin, query .../servers/{serverName}/administrators
func (o LookupServerResultOutput) ExternalAdministratorLogin() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.ExternalAdministratorLogin }).(pulumi.StringOutput)
}

// The ID of the Active Azure Directory object with admin permissions on this server. Legacy parameter, always null. To check for Active Directory admin, query .../servers/{serverName}/administrators.
func (o LookupServerResultOutput) ExternalAdministratorSid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.ExternalAdministratorSid }).(pulumi.StringOutput)
}

// The fully qualified domain name of the server.
func (o LookupServerResultOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of sql server.  This is metadata used for the Azure portal experience.
func (o LookupServerResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupServerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Name }).(pulumi.StringOutput)
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

func init() {
	pulumi.RegisterOutputType(LookupServerResultOutput{})
}
