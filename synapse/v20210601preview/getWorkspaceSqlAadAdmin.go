// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a workspace SQL active directory admin
func LookupWorkspaceSqlAadAdmin(ctx *pulumi.Context, args *LookupWorkspaceSqlAadAdminArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceSqlAadAdminResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceSqlAadAdminResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getWorkspaceSqlAadAdmin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceSqlAadAdminArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Workspace active directory administrator
type LookupWorkspaceSqlAadAdminResult struct {
	// Workspace active directory administrator type
	AdministratorType *string `pulumi:"administratorType"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Login of the workspace active directory administrator
	Login *string `pulumi:"login"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Object ID of the workspace active directory administrator
	Sid *string `pulumi:"sid"`
	// Tenant ID of the workspace active directory administrator
	TenantId *string `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkspaceSqlAadAdminOutput(ctx *pulumi.Context, args LookupWorkspaceSqlAadAdminOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceSqlAadAdminResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceSqlAadAdminResult, error) {
			args := v.(LookupWorkspaceSqlAadAdminArgs)
			r, err := LookupWorkspaceSqlAadAdmin(ctx, &args, opts...)
			var s LookupWorkspaceSqlAadAdminResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceSqlAadAdminResultOutput)
}

type LookupWorkspaceSqlAadAdminOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceSqlAadAdminOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceSqlAadAdminArgs)(nil)).Elem()
}

// Workspace active directory administrator
type LookupWorkspaceSqlAadAdminResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceSqlAadAdminResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceSqlAadAdminResult)(nil)).Elem()
}

func (o LookupWorkspaceSqlAadAdminResultOutput) ToLookupWorkspaceSqlAadAdminResultOutput() LookupWorkspaceSqlAadAdminResultOutput {
	return o
}

func (o LookupWorkspaceSqlAadAdminResultOutput) ToLookupWorkspaceSqlAadAdminResultOutputWithContext(ctx context.Context) LookupWorkspaceSqlAadAdminResultOutput {
	return o
}

func (o LookupWorkspaceSqlAadAdminResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkspaceSqlAadAdminResult] {
	return pulumix.Output[LookupWorkspaceSqlAadAdminResult]{
		OutputState: o.OutputState,
	}
}

// Workspace active directory administrator type
func (o LookupWorkspaceSqlAadAdminResultOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSqlAadAdminResult) *string { return v.AdministratorType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceSqlAadAdminResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSqlAadAdminResult) string { return v.Id }).(pulumi.StringOutput)
}

// Login of the workspace active directory administrator
func (o LookupWorkspaceSqlAadAdminResultOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSqlAadAdminResult) *string { return v.Login }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupWorkspaceSqlAadAdminResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSqlAadAdminResult) string { return v.Name }).(pulumi.StringOutput)
}

// Object ID of the workspace active directory administrator
func (o LookupWorkspaceSqlAadAdminResultOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSqlAadAdminResult) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

// Tenant ID of the workspace active directory administrator
func (o LookupWorkspaceSqlAadAdminResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSqlAadAdminResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceSqlAadAdminResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSqlAadAdminResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceSqlAadAdminResultOutput{})
}
