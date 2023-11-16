// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a managed instance administrator.
func LookupManagedInstanceAdministrator(ctx *pulumi.Context, args *LookupManagedInstanceAdministratorArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstanceAdministratorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedInstanceAdministratorResult
	err := ctx.Invoke("azure-native:sql/v20230501preview:getManagedInstanceAdministrator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstanceAdministratorArgs struct {
	AdministratorName string `pulumi:"administratorName"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure SQL managed instance administrator.
type LookupManagedInstanceAdministratorResult struct {
	// Type of the managed instance administrator.
	AdministratorType string `pulumi:"administratorType"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Login name of the managed instance administrator.
	Login string `pulumi:"login"`
	// Resource name.
	Name string `pulumi:"name"`
	// SID (object ID) of the managed instance administrator.
	Sid string `pulumi:"sid"`
	// Tenant ID of the managed instance administrator.
	TenantId *string `pulumi:"tenantId"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupManagedInstanceAdministratorOutput(ctx *pulumi.Context, args LookupManagedInstanceAdministratorOutputArgs, opts ...pulumi.InvokeOption) LookupManagedInstanceAdministratorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedInstanceAdministratorResult, error) {
			args := v.(LookupManagedInstanceAdministratorArgs)
			r, err := LookupManagedInstanceAdministrator(ctx, &args, opts...)
			var s LookupManagedInstanceAdministratorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedInstanceAdministratorResultOutput)
}

type LookupManagedInstanceAdministratorOutputArgs struct {
	AdministratorName pulumi.StringInput `pulumi:"administratorName"`
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedInstanceAdministratorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceAdministratorArgs)(nil)).Elem()
}

// An Azure SQL managed instance administrator.
type LookupManagedInstanceAdministratorResultOutput struct{ *pulumi.OutputState }

func (LookupManagedInstanceAdministratorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceAdministratorResult)(nil)).Elem()
}

func (o LookupManagedInstanceAdministratorResultOutput) ToLookupManagedInstanceAdministratorResultOutput() LookupManagedInstanceAdministratorResultOutput {
	return o
}

func (o LookupManagedInstanceAdministratorResultOutput) ToLookupManagedInstanceAdministratorResultOutputWithContext(ctx context.Context) LookupManagedInstanceAdministratorResultOutput {
	return o
}

// Type of the managed instance administrator.
func (o LookupManagedInstanceAdministratorResultOutput) AdministratorType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAdministratorResult) string { return v.AdministratorType }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupManagedInstanceAdministratorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAdministratorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Login name of the managed instance administrator.
func (o LookupManagedInstanceAdministratorResultOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAdministratorResult) string { return v.Login }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupManagedInstanceAdministratorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAdministratorResult) string { return v.Name }).(pulumi.StringOutput)
}

// SID (object ID) of the managed instance administrator.
func (o LookupManagedInstanceAdministratorResultOutput) Sid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAdministratorResult) string { return v.Sid }).(pulumi.StringOutput)
}

// Tenant ID of the managed instance administrator.
func (o LookupManagedInstanceAdministratorResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceAdministratorResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupManagedInstanceAdministratorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAdministratorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedInstanceAdministratorResultOutput{})
}
