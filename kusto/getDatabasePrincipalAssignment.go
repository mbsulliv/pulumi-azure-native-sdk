// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a Kusto cluster database principalAssignment.
// Azure REST API version: 2022-12-29.
func LookupDatabasePrincipalAssignment(ctx *pulumi.Context, args *LookupDatabasePrincipalAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupDatabasePrincipalAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabasePrincipalAssignmentResult
	err := ctx.Invoke("azure-native:kusto:getDatabasePrincipalAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabasePrincipalAssignmentArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName string `pulumi:"principalAssignmentName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing a database principal assignment.
type LookupDatabasePrincipalAssignmentResult struct {
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
	// The tenant id of the principal
	TenantId *string `pulumi:"tenantId"`
	// The tenant name of the principal
	TenantName string `pulumi:"tenantName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDatabasePrincipalAssignmentOutput(ctx *pulumi.Context, args LookupDatabasePrincipalAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupDatabasePrincipalAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabasePrincipalAssignmentResult, error) {
			args := v.(LookupDatabasePrincipalAssignmentArgs)
			r, err := LookupDatabasePrincipalAssignment(ctx, &args, opts...)
			var s LookupDatabasePrincipalAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabasePrincipalAssignmentResultOutput)
}

type LookupDatabasePrincipalAssignmentOutputArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName pulumi.StringInput `pulumi:"principalAssignmentName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabasePrincipalAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasePrincipalAssignmentArgs)(nil)).Elem()
}

// Class representing a database principal assignment.
type LookupDatabasePrincipalAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupDatabasePrincipalAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasePrincipalAssignmentResult)(nil)).Elem()
}

func (o LookupDatabasePrincipalAssignmentResultOutput) ToLookupDatabasePrincipalAssignmentResultOutput() LookupDatabasePrincipalAssignmentResultOutput {
	return o
}

func (o LookupDatabasePrincipalAssignmentResultOutput) ToLookupDatabasePrincipalAssignmentResultOutputWithContext(ctx context.Context) LookupDatabasePrincipalAssignmentResultOutput {
	return o
}

func (o LookupDatabasePrincipalAssignmentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDatabasePrincipalAssignmentResult] {
	return pulumix.Output[LookupDatabasePrincipalAssignmentResult]{
		OutputState: o.OutputState,
	}
}

// The service principal object id in AAD (Azure active directory)
func (o LookupDatabasePrincipalAssignmentResultOutput) AadObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.AadObjectId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDatabasePrincipalAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDatabasePrincipalAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
func (o LookupDatabasePrincipalAssignmentResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The principal name
func (o LookupDatabasePrincipalAssignmentResultOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.PrincipalName }).(pulumi.StringOutput)
}

// Principal type.
func (o LookupDatabasePrincipalAssignmentResultOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.PrincipalType }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o LookupDatabasePrincipalAssignmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Database principal role.
func (o LookupDatabasePrincipalAssignmentResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.Role }).(pulumi.StringOutput)
}

// The tenant id of the principal
func (o LookupDatabasePrincipalAssignmentResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The tenant name of the principal
func (o LookupDatabasePrincipalAssignmentResultOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.TenantName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDatabasePrincipalAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrincipalAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabasePrincipalAssignmentResultOutput{})
}
