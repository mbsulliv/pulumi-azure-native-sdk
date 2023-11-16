// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbforpostgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a server.
// Azure REST API version: 2022-12-01.
//
// Other available API versions: 2023-03-01-preview, 2023-06-01-preview.
func LookupAdministrator(ctx *pulumi.Context, args *LookupAdministratorArgs, opts ...pulumi.InvokeOption) (*LookupAdministratorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAdministratorResult
	err := ctx.Invoke("azure-native:dbforpostgresql:getAdministrator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAdministratorArgs struct {
	// Guid of the objectId for the administrator.
	ObjectId string `pulumi:"objectId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Represents an Active Directory administrator.
type LookupAdministratorResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The objectId of the Active Directory administrator.
	ObjectId *string `pulumi:"objectId"`
	// Active Directory administrator principal name.
	PrincipalName *string `pulumi:"principalName"`
	// The principal type used to represent the type of Active Directory Administrator.
	PrincipalType *string `pulumi:"principalType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tenantId of the Active Directory administrator.
	TenantId *string `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAdministratorOutput(ctx *pulumi.Context, args LookupAdministratorOutputArgs, opts ...pulumi.InvokeOption) LookupAdministratorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAdministratorResult, error) {
			args := v.(LookupAdministratorArgs)
			r, err := LookupAdministrator(ctx, &args, opts...)
			var s LookupAdministratorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAdministratorResultOutput)
}

type LookupAdministratorOutputArgs struct {
	// Guid of the objectId for the administrator.
	ObjectId pulumi.StringInput `pulumi:"objectId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupAdministratorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdministratorArgs)(nil)).Elem()
}

// Represents an Active Directory administrator.
type LookupAdministratorResultOutput struct{ *pulumi.OutputState }

func (LookupAdministratorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdministratorResult)(nil)).Elem()
}

func (o LookupAdministratorResultOutput) ToLookupAdministratorResultOutput() LookupAdministratorResultOutput {
	return o
}

func (o LookupAdministratorResultOutput) ToLookupAdministratorResultOutputWithContext(ctx context.Context) LookupAdministratorResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAdministratorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdministratorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAdministratorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdministratorResult) string { return v.Name }).(pulumi.StringOutput)
}

// The objectId of the Active Directory administrator.
func (o LookupAdministratorResultOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAdministratorResult) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

// Active Directory administrator principal name.
func (o LookupAdministratorResultOutput) PrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAdministratorResult) *string { return v.PrincipalName }).(pulumi.StringPtrOutput)
}

// The principal type used to represent the type of Active Directory Administrator.
func (o LookupAdministratorResultOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAdministratorResult) *string { return v.PrincipalType }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAdministratorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAdministratorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenantId of the Active Directory administrator.
func (o LookupAdministratorResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAdministratorResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAdministratorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdministratorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAdministratorResultOutput{})
}
