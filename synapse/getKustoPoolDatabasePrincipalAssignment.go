// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a Kusto pool database principalAssignment.
// Azure REST API version: 2021-06-01-preview.
func LookupKustoPoolDatabasePrincipalAssignment(ctx *pulumi.Context, args *LookupKustoPoolDatabasePrincipalAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupKustoPoolDatabasePrincipalAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupKustoPoolDatabasePrincipalAssignmentResult
	err := ctx.Invoke("azure-native:synapse:getKustoPoolDatabasePrincipalAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoPoolDatabasePrincipalAssignmentArgs struct {
	// The name of the database in the Kusto pool.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the Kusto pool.
	KustoPoolName string `pulumi:"kustoPoolName"`
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName string `pulumi:"principalAssignmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Class representing a database principal assignment.
type LookupKustoPoolDatabasePrincipalAssignmentResult struct {
	// The service principal object id in AAD (Azure active directory)
	AadObjectId string `pulumi:"aadObjectId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
	PrincipalId string `pulumi:"principalId"`
	// The principal name
	PrincipalName string `pulumi:"principalName"`
	// Principal type.
	PrincipalType string `pulumi:"principalType"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Database principal role.
	Role string `pulumi:"role"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tenant id of the principal
	TenantId *string `pulumi:"tenantId"`
	// The tenant name of the principal
	TenantName string `pulumi:"tenantName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupKustoPoolDatabasePrincipalAssignmentOutput(ctx *pulumi.Context, args LookupKustoPoolDatabasePrincipalAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupKustoPoolDatabasePrincipalAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoPoolDatabasePrincipalAssignmentResult, error) {
			args := v.(LookupKustoPoolDatabasePrincipalAssignmentArgs)
			r, err := LookupKustoPoolDatabasePrincipalAssignment(ctx, &args, opts...)
			var s LookupKustoPoolDatabasePrincipalAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoPoolDatabasePrincipalAssignmentResultOutput)
}

type LookupKustoPoolDatabasePrincipalAssignmentOutputArgs struct {
	// The name of the database in the Kusto pool.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the Kusto pool.
	KustoPoolName pulumi.StringInput `pulumi:"kustoPoolName"`
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName pulumi.StringInput `pulumi:"principalAssignmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupKustoPoolDatabasePrincipalAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolDatabasePrincipalAssignmentArgs)(nil)).Elem()
}

// Class representing a database principal assignment.
type LookupKustoPoolDatabasePrincipalAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupKustoPoolDatabasePrincipalAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolDatabasePrincipalAssignmentResult)(nil)).Elem()
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) ToLookupKustoPoolDatabasePrincipalAssignmentResultOutput() LookupKustoPoolDatabasePrincipalAssignmentResultOutput {
	return o
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) ToLookupKustoPoolDatabasePrincipalAssignmentResultOutputWithContext(ctx context.Context) LookupKustoPoolDatabasePrincipalAssignmentResultOutput {
	return o
}

func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupKustoPoolDatabasePrincipalAssignmentResult] {
	return pulumix.Output[LookupKustoPoolDatabasePrincipalAssignmentResult]{
		OutputState: o.OutputState,
	}
}

// The service principal object id in AAD (Azure active directory)
func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) AadObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.AadObjectId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The principal name
func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.PrincipalName }).(pulumi.StringOutput)
}

// Principal type.
func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.PrincipalType }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Database principal role.
func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.Role }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id of the principal
func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The tenant name of the principal
func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.TenantName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupKustoPoolDatabasePrincipalAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabasePrincipalAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoPoolDatabasePrincipalAssignmentResultOutput{})
}
