// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databoxedge

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a specific role by name.
// Azure REST API version: 2022-03-01.
func LookupCloudEdgeManagementRole(ctx *pulumi.Context, args *LookupCloudEdgeManagementRoleArgs, opts ...pulumi.InvokeOption) (*LookupCloudEdgeManagementRoleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCloudEdgeManagementRoleResult
	err := ctx.Invoke("azure-native:databoxedge:getCloudEdgeManagementRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCloudEdgeManagementRoleArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The role name.
	Name string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The preview of Virtual Machine Cloud Management from the Azure supports deploying and managing VMs on your Azure Stack Edge device from Azure Portal.
// For more information, refer to: https://docs.microsoft.com/en-us/azure/databox-online/azure-stack-edge-gpu-virtual-machine-overview
// By using this feature, you agree to the preview legal terms. See the https://azure.microsoft.com/en-us/support/legal/preview-supplemental-terms/ for additional details.
type LookupCloudEdgeManagementRoleResult struct {
	// Edge Profile of the resource
	EdgeProfile EdgeProfileResponse `pulumi:"edgeProfile"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// Role type.
	// Expected value is 'CloudEdgeManagement'.
	Kind string `pulumi:"kind"`
	// Local Edge Management Status
	LocalManagementStatus string `pulumi:"localManagementStatus"`
	// The object name.
	Name string `pulumi:"name"`
	// Role status.
	RoleStatus string `pulumi:"roleStatus"`
	// Metadata pertaining to creation and last modification of Role
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
}

func LookupCloudEdgeManagementRoleOutput(ctx *pulumi.Context, args LookupCloudEdgeManagementRoleOutputArgs, opts ...pulumi.InvokeOption) LookupCloudEdgeManagementRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudEdgeManagementRoleResult, error) {
			args := v.(LookupCloudEdgeManagementRoleArgs)
			r, err := LookupCloudEdgeManagementRole(ctx, &args, opts...)
			var s LookupCloudEdgeManagementRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudEdgeManagementRoleResultOutput)
}

type LookupCloudEdgeManagementRoleOutputArgs struct {
	// The device name.
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The role name.
	Name pulumi.StringInput `pulumi:"name"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCloudEdgeManagementRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudEdgeManagementRoleArgs)(nil)).Elem()
}

// The preview of Virtual Machine Cloud Management from the Azure supports deploying and managing VMs on your Azure Stack Edge device from Azure Portal.
// For more information, refer to: https://docs.microsoft.com/en-us/azure/databox-online/azure-stack-edge-gpu-virtual-machine-overview
// By using this feature, you agree to the preview legal terms. See the https://azure.microsoft.com/en-us/support/legal/preview-supplemental-terms/ for additional details.
type LookupCloudEdgeManagementRoleResultOutput struct{ *pulumi.OutputState }

func (LookupCloudEdgeManagementRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudEdgeManagementRoleResult)(nil)).Elem()
}

func (o LookupCloudEdgeManagementRoleResultOutput) ToLookupCloudEdgeManagementRoleResultOutput() LookupCloudEdgeManagementRoleResultOutput {
	return o
}

func (o LookupCloudEdgeManagementRoleResultOutput) ToLookupCloudEdgeManagementRoleResultOutputWithContext(ctx context.Context) LookupCloudEdgeManagementRoleResultOutput {
	return o
}

// Edge Profile of the resource
func (o LookupCloudEdgeManagementRoleResultOutput) EdgeProfile() EdgeProfileResponseOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) EdgeProfileResponse { return v.EdgeProfile }).(EdgeProfileResponseOutput)
}

// The path ID that uniquely identifies the object.
func (o LookupCloudEdgeManagementRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Role type.
// Expected value is 'CloudEdgeManagement'.
func (o LookupCloudEdgeManagementRoleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Local Edge Management Status
func (o LookupCloudEdgeManagementRoleResultOutput) LocalManagementStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) string { return v.LocalManagementStatus }).(pulumi.StringOutput)
}

// The object name.
func (o LookupCloudEdgeManagementRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Role status.
func (o LookupCloudEdgeManagementRoleResultOutput) RoleStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) string { return v.RoleStatus }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of Role
func (o LookupCloudEdgeManagementRoleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o LookupCloudEdgeManagementRoleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudEdgeManagementRoleResultOutput{})
}
