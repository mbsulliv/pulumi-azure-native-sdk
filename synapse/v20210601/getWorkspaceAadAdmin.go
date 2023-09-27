// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a workspace active directory admin
func LookupWorkspaceAadAdmin(ctx *pulumi.Context, args *LookupWorkspaceAadAdminArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceAadAdminResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceAadAdminResult
	err := ctx.Invoke("azure-native:synapse/v20210601:getWorkspaceAadAdmin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceAadAdminArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Workspace active directory administrator
type LookupWorkspaceAadAdminResult struct {
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

func LookupWorkspaceAadAdminOutput(ctx *pulumi.Context, args LookupWorkspaceAadAdminOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceAadAdminResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceAadAdminResult, error) {
			args := v.(LookupWorkspaceAadAdminArgs)
			r, err := LookupWorkspaceAadAdmin(ctx, &args, opts...)
			var s LookupWorkspaceAadAdminResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceAadAdminResultOutput)
}

type LookupWorkspaceAadAdminOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceAadAdminOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceAadAdminArgs)(nil)).Elem()
}

// Workspace active directory administrator
type LookupWorkspaceAadAdminResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceAadAdminResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceAadAdminResult)(nil)).Elem()
}

func (o LookupWorkspaceAadAdminResultOutput) ToLookupWorkspaceAadAdminResultOutput() LookupWorkspaceAadAdminResultOutput {
	return o
}

func (o LookupWorkspaceAadAdminResultOutput) ToLookupWorkspaceAadAdminResultOutputWithContext(ctx context.Context) LookupWorkspaceAadAdminResultOutput {
	return o
}

func (o LookupWorkspaceAadAdminResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkspaceAadAdminResult] {
	return pulumix.Output[LookupWorkspaceAadAdminResult]{
		OutputState: o.OutputState,
	}
}

// Workspace active directory administrator type
func (o LookupWorkspaceAadAdminResultOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceAadAdminResult) *string { return v.AdministratorType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceAadAdminResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceAadAdminResult) string { return v.Id }).(pulumi.StringOutput)
}

// Login of the workspace active directory administrator
func (o LookupWorkspaceAadAdminResultOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceAadAdminResult) *string { return v.Login }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupWorkspaceAadAdminResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceAadAdminResult) string { return v.Name }).(pulumi.StringOutput)
}

// Object ID of the workspace active directory administrator
func (o LookupWorkspaceAadAdminResultOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceAadAdminResult) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

// Tenant ID of the workspace active directory administrator
func (o LookupWorkspaceAadAdminResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceAadAdminResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceAadAdminResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceAadAdminResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceAadAdminResultOutput{})
}
