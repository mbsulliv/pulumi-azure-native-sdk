// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a Kusto pool principalAssignment.
// Azure REST API version: 2021-06-01-preview.
func LookupKustoPoolPrincipalAssignment(ctx *pulumi.Context, args *LookupKustoPoolPrincipalAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupKustoPoolPrincipalAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupKustoPoolPrincipalAssignmentResult
	err := ctx.Invoke("azure-native:synapse:getKustoPoolPrincipalAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoPoolPrincipalAssignmentArgs struct {
	// The name of the Kusto pool.
	KustoPoolName string `pulumi:"kustoPoolName"`
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName string `pulumi:"principalAssignmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Class representing a cluster principal assignment.
type LookupKustoPoolPrincipalAssignmentResult struct {
	// The service principal object id in AAD (Azure active directory)
	AadObjectId string `pulumi:"aadObjectId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
	PrincipalId string `pulumi:"principalId"`
	// The principal name
	PrincipalName string `pulumi:"principalName"`
	// Principal type.
	PrincipalType string `pulumi:"principalType"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Cluster principal role.
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

func LookupKustoPoolPrincipalAssignmentOutput(ctx *pulumi.Context, args LookupKustoPoolPrincipalAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupKustoPoolPrincipalAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoPoolPrincipalAssignmentResult, error) {
			args := v.(LookupKustoPoolPrincipalAssignmentArgs)
			r, err := LookupKustoPoolPrincipalAssignment(ctx, &args, opts...)
			var s LookupKustoPoolPrincipalAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoPoolPrincipalAssignmentResultOutput)
}

type LookupKustoPoolPrincipalAssignmentOutputArgs struct {
	// The name of the Kusto pool.
	KustoPoolName pulumi.StringInput `pulumi:"kustoPoolName"`
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName pulumi.StringInput `pulumi:"principalAssignmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupKustoPoolPrincipalAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolPrincipalAssignmentArgs)(nil)).Elem()
}

// Class representing a cluster principal assignment.
type LookupKustoPoolPrincipalAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupKustoPoolPrincipalAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolPrincipalAssignmentResult)(nil)).Elem()
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) ToLookupKustoPoolPrincipalAssignmentResultOutput() LookupKustoPoolPrincipalAssignmentResultOutput {
	return o
}

func (o LookupKustoPoolPrincipalAssignmentResultOutput) ToLookupKustoPoolPrincipalAssignmentResultOutputWithContext(ctx context.Context) LookupKustoPoolPrincipalAssignmentResultOutput {
	return o
}

// The service principal object id in AAD (Azure active directory)
func (o LookupKustoPoolPrincipalAssignmentResultOutput) AadObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.AadObjectId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupKustoPoolPrincipalAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupKustoPoolPrincipalAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
func (o LookupKustoPoolPrincipalAssignmentResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The principal name
func (o LookupKustoPoolPrincipalAssignmentResultOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.PrincipalName }).(pulumi.StringOutput)
}

// Principal type.
func (o LookupKustoPoolPrincipalAssignmentResultOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.PrincipalType }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o LookupKustoPoolPrincipalAssignmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Cluster principal role.
func (o LookupKustoPoolPrincipalAssignmentResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.Role }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupKustoPoolPrincipalAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id of the principal
func (o LookupKustoPoolPrincipalAssignmentResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The tenant name of the principal
func (o LookupKustoPoolPrincipalAssignmentResultOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.TenantName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupKustoPoolPrincipalAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolPrincipalAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoPoolPrincipalAssignmentResultOutput{})
}
