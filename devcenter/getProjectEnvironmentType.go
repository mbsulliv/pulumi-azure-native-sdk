// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devcenter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a project environment type.
// Azure REST API version: 2023-04-01.
func LookupProjectEnvironmentType(ctx *pulumi.Context, args *LookupProjectEnvironmentTypeArgs, opts ...pulumi.InvokeOption) (*LookupProjectEnvironmentTypeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectEnvironmentTypeResult
	err := ctx.Invoke("azure-native:devcenter:getProjectEnvironmentType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectEnvironmentTypeArgs struct {
	// The name of the environment type.
	EnvironmentTypeName string `pulumi:"environmentTypeName"`
	// The name of the project.
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents an environment type.
type LookupProjectEnvironmentTypeResult struct {
	// The role definition assigned to the environment creator on backing resources.
	CreatorRoleAssignment *ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignment `pulumi:"creatorRoleAssignment"`
	// Id of a subscription that the environment type will be mapped to. The environment's resources will be deployed into this subscription.
	DeploymentTargetId *string `pulumi:"deploymentTargetId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Managed identity properties
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The geo-location for the environment type
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Defines whether this Environment Type can be used in this Project.
	Status *string `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Role Assignments created on environment backing resources. This is a mapping from a user object ID to an object of role definition IDs.
	UserRoleAssignments map[string]UserRoleAssignmentResponse `pulumi:"userRoleAssignments"`
}

func LookupProjectEnvironmentTypeOutput(ctx *pulumi.Context, args LookupProjectEnvironmentTypeOutputArgs, opts ...pulumi.InvokeOption) LookupProjectEnvironmentTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectEnvironmentTypeResult, error) {
			args := v.(LookupProjectEnvironmentTypeArgs)
			r, err := LookupProjectEnvironmentType(ctx, &args, opts...)
			var s LookupProjectEnvironmentTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectEnvironmentTypeResultOutput)
}

type LookupProjectEnvironmentTypeOutputArgs struct {
	// The name of the environment type.
	EnvironmentTypeName pulumi.StringInput `pulumi:"environmentTypeName"`
	// The name of the project.
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProjectEnvironmentTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectEnvironmentTypeArgs)(nil)).Elem()
}

// Represents an environment type.
type LookupProjectEnvironmentTypeResultOutput struct{ *pulumi.OutputState }

func (LookupProjectEnvironmentTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectEnvironmentTypeResult)(nil)).Elem()
}

func (o LookupProjectEnvironmentTypeResultOutput) ToLookupProjectEnvironmentTypeResultOutput() LookupProjectEnvironmentTypeResultOutput {
	return o
}

func (o LookupProjectEnvironmentTypeResultOutput) ToLookupProjectEnvironmentTypeResultOutputWithContext(ctx context.Context) LookupProjectEnvironmentTypeResultOutput {
	return o
}

func (o LookupProjectEnvironmentTypeResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupProjectEnvironmentTypeResult] {
	return pulumix.Output[LookupProjectEnvironmentTypeResult]{
		OutputState: o.OutputState,
	}
}

// The role definition assigned to the environment creator on backing resources.
func (o LookupProjectEnvironmentTypeResultOutput) CreatorRoleAssignment() ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) *ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignment {
		return v.CreatorRoleAssignment
	}).(ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput)
}

// Id of a subscription that the environment type will be mapped to. The environment's resources will be deployed into this subscription.
func (o LookupProjectEnvironmentTypeResultOutput) DeploymentTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) *string { return v.DeploymentTargetId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupProjectEnvironmentTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed identity properties
func (o LookupProjectEnvironmentTypeResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location for the environment type
func (o LookupProjectEnvironmentTypeResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupProjectEnvironmentTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupProjectEnvironmentTypeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Defines whether this Environment Type can be used in this Project.
func (o LookupProjectEnvironmentTypeResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupProjectEnvironmentTypeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupProjectEnvironmentTypeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupProjectEnvironmentTypeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) string { return v.Type }).(pulumi.StringOutput)
}

// Role Assignments created on environment backing resources. This is a mapping from a user object ID to an object of role definition IDs.
func (o LookupProjectEnvironmentTypeResultOutput) UserRoleAssignments() UserRoleAssignmentResponseMapOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) map[string]UserRoleAssignmentResponse {
		return v.UserRoleAssignments
	}).(UserRoleAssignmentResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectEnvironmentTypeResultOutput{})
}
