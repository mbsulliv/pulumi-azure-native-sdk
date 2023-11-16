// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managednetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the specified scope assignment.
// Azure REST API version: 2019-06-01-preview.
func LookupScopeAssignment(ctx *pulumi.Context, args *LookupScopeAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupScopeAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupScopeAssignmentResult
	err := ctx.Invoke("azure-native:managednetwork:getScopeAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScopeAssignmentArgs struct {
	// The base resource of the scope assignment.
	Scope string `pulumi:"scope"`
	// The name of the scope assignment to get.
	ScopeAssignmentName string `pulumi:"scopeAssignmentName"`
}

// The Managed Network resource
type LookupScopeAssignmentResult struct {
	// The managed network ID with scope will be assigned to.
	AssignedManagedNetwork *string `pulumi:"assignedManagedNetwork"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the ManagedNetwork resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}

func LookupScopeAssignmentOutput(ctx *pulumi.Context, args LookupScopeAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupScopeAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScopeAssignmentResult, error) {
			args := v.(LookupScopeAssignmentArgs)
			r, err := LookupScopeAssignment(ctx, &args, opts...)
			var s LookupScopeAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScopeAssignmentResultOutput)
}

type LookupScopeAssignmentOutputArgs struct {
	// The base resource of the scope assignment.
	Scope pulumi.StringInput `pulumi:"scope"`
	// The name of the scope assignment to get.
	ScopeAssignmentName pulumi.StringInput `pulumi:"scopeAssignmentName"`
}

func (LookupScopeAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeAssignmentArgs)(nil)).Elem()
}

// The Managed Network resource
type LookupScopeAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupScopeAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeAssignmentResult)(nil)).Elem()
}

func (o LookupScopeAssignmentResultOutput) ToLookupScopeAssignmentResultOutput() LookupScopeAssignmentResultOutput {
	return o
}

func (o LookupScopeAssignmentResultOutput) ToLookupScopeAssignmentResultOutputWithContext(ctx context.Context) LookupScopeAssignmentResultOutput {
	return o
}

// The managed network ID with scope will be assigned to.
func (o LookupScopeAssignmentResultOutput) AssignedManagedNetwork() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAssignmentResult) *string { return v.AssignedManagedNetwork }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupScopeAssignmentResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAssignmentResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupScopeAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupScopeAssignmentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAssignmentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupScopeAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the ManagedNetwork resource.
func (o LookupScopeAssignmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAssignmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o LookupScopeAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScopeAssignmentResultOutput{})
}
