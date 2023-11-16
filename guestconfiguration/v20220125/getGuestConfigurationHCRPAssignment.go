// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220125

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a guest configuration assignment
func LookupGuestConfigurationHCRPAssignment(ctx *pulumi.Context, args *LookupGuestConfigurationHCRPAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupGuestConfigurationHCRPAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGuestConfigurationHCRPAssignmentResult
	err := ctx.Invoke("azure-native:guestconfiguration/v20220125:getGuestConfigurationHCRPAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupGuestConfigurationHCRPAssignmentArgs struct {
	// The guest configuration assignment name.
	GuestConfigurationAssignmentName string `pulumi:"guestConfigurationAssignmentName"`
	// The name of the ARC machine.
	MachineName string `pulumi:"machineName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Guest configuration assignment is an association between a machine and guest configuration.
type LookupGuestConfigurationHCRPAssignmentResult struct {
	// ARM resource id of the guest configuration assignment.
	Id string `pulumi:"id"`
	// Region where the VM is located.
	Location *string `pulumi:"location"`
	// Name of the guest configuration assignment.
	Name *string `pulumi:"name"`
	// Properties of the Guest configuration assignment.
	Properties GuestConfigurationAssignmentPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupGuestConfigurationHCRPAssignmentResult
func (val *LookupGuestConfigurationHCRPAssignmentResult) Defaults() *LookupGuestConfigurationHCRPAssignmentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupGuestConfigurationHCRPAssignmentOutput(ctx *pulumi.Context, args LookupGuestConfigurationHCRPAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupGuestConfigurationHCRPAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGuestConfigurationHCRPAssignmentResult, error) {
			args := v.(LookupGuestConfigurationHCRPAssignmentArgs)
			r, err := LookupGuestConfigurationHCRPAssignment(ctx, &args, opts...)
			var s LookupGuestConfigurationHCRPAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGuestConfigurationHCRPAssignmentResultOutput)
}

type LookupGuestConfigurationHCRPAssignmentOutputArgs struct {
	// The guest configuration assignment name.
	GuestConfigurationAssignmentName pulumi.StringInput `pulumi:"guestConfigurationAssignmentName"`
	// The name of the ARC machine.
	MachineName pulumi.StringInput `pulumi:"machineName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGuestConfigurationHCRPAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestConfigurationHCRPAssignmentArgs)(nil)).Elem()
}

// Guest configuration assignment is an association between a machine and guest configuration.
type LookupGuestConfigurationHCRPAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupGuestConfigurationHCRPAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestConfigurationHCRPAssignmentResult)(nil)).Elem()
}

func (o LookupGuestConfigurationHCRPAssignmentResultOutput) ToLookupGuestConfigurationHCRPAssignmentResultOutput() LookupGuestConfigurationHCRPAssignmentResultOutput {
	return o
}

func (o LookupGuestConfigurationHCRPAssignmentResultOutput) ToLookupGuestConfigurationHCRPAssignmentResultOutputWithContext(ctx context.Context) LookupGuestConfigurationHCRPAssignmentResultOutput {
	return o
}

// ARM resource id of the guest configuration assignment.
func (o LookupGuestConfigurationHCRPAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestConfigurationHCRPAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Region where the VM is located.
func (o LookupGuestConfigurationHCRPAssignmentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGuestConfigurationHCRPAssignmentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the guest configuration assignment.
func (o LookupGuestConfigurationHCRPAssignmentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGuestConfigurationHCRPAssignmentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Properties of the Guest configuration assignment.
func (o LookupGuestConfigurationHCRPAssignmentResultOutput) Properties() GuestConfigurationAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGuestConfigurationHCRPAssignmentResult) GuestConfigurationAssignmentPropertiesResponse {
		return v.Properties
	}).(GuestConfigurationAssignmentPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupGuestConfigurationHCRPAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGuestConfigurationHCRPAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupGuestConfigurationHCRPAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestConfigurationHCRPAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGuestConfigurationHCRPAssignmentResultOutput{})
}
