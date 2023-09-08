// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a AAD server administrator.
func LookupServerAdministrator(ctx *pulumi.Context, args *LookupServerAdministratorArgs, opts ...pulumi.InvokeOption) (*LookupServerAdministratorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServerAdministratorResult
	err := ctx.Invoke("azure-native:dbformysql/v20180601privatepreview:getServerAdministrator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerAdministratorArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Represents a and external administrator to be created.
type LookupServerAdministratorResult struct {
	// The type of administrator.
	AdministratorType string `pulumi:"administratorType"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The server administrator login account name.
	Login string `pulumi:"login"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The server administrator Sid (Secure ID).
	Sid string `pulumi:"sid"`
	// The server Active Directory Administrator tenant id.
	TenantId string `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupServerAdministratorOutput(ctx *pulumi.Context, args LookupServerAdministratorOutputArgs, opts ...pulumi.InvokeOption) LookupServerAdministratorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerAdministratorResult, error) {
			args := v.(LookupServerAdministratorArgs)
			r, err := LookupServerAdministrator(ctx, &args, opts...)
			var s LookupServerAdministratorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerAdministratorResultOutput)
}

type LookupServerAdministratorOutputArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerAdministratorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerAdministratorArgs)(nil)).Elem()
}

// Represents a and external administrator to be created.
type LookupServerAdministratorResultOutput struct{ *pulumi.OutputState }

func (LookupServerAdministratorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerAdministratorResult)(nil)).Elem()
}

func (o LookupServerAdministratorResultOutput) ToLookupServerAdministratorResultOutput() LookupServerAdministratorResultOutput {
	return o
}

func (o LookupServerAdministratorResultOutput) ToLookupServerAdministratorResultOutputWithContext(ctx context.Context) LookupServerAdministratorResultOutput {
	return o
}

// The type of administrator.
func (o LookupServerAdministratorResultOutput) AdministratorType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdministratorResult) string { return v.AdministratorType }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupServerAdministratorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdministratorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The server administrator login account name.
func (o LookupServerAdministratorResultOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdministratorResult) string { return v.Login }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupServerAdministratorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdministratorResult) string { return v.Name }).(pulumi.StringOutput)
}

// The server administrator Sid (Secure ID).
func (o LookupServerAdministratorResultOutput) Sid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdministratorResult) string { return v.Sid }).(pulumi.StringOutput)
}

// The server Active Directory Administrator tenant id.
func (o LookupServerAdministratorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdministratorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupServerAdministratorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdministratorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerAdministratorResultOutput{})
}