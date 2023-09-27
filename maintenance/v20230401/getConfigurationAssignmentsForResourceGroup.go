// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get configuration assignment for resource..
func LookupConfigurationAssignmentsForResourceGroup(ctx *pulumi.Context, args *LookupConfigurationAssignmentsForResourceGroupArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationAssignmentsForResourceGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConfigurationAssignmentsForResourceGroupResult
	err := ctx.Invoke("azure-native:maintenance/v20230401:getConfigurationAssignmentsForResourceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationAssignmentsForResourceGroupArgs struct {
	// Configuration assignment name
	ConfigurationAssignmentName string `pulumi:"configurationAssignmentName"`
	// Resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Configuration Assignment
type LookupConfigurationAssignmentsForResourceGroupResult struct {
	// Properties of the configuration assignment
	Filter *ConfigurationAssignmentFilterPropertiesResponse `pulumi:"filter"`
	// Fully qualified identifier of the resource
	Id string `pulumi:"id"`
	// Location of the resource
	Location *string `pulumi:"location"`
	// The maintenance configuration Id
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// Name of the resource
	Name string `pulumi:"name"`
	// The unique resourceId
	ResourceId *string `pulumi:"resourceId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the resource
	Type string `pulumi:"type"`
}

func LookupConfigurationAssignmentsForResourceGroupOutput(ctx *pulumi.Context, args LookupConfigurationAssignmentsForResourceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationAssignmentsForResourceGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationAssignmentsForResourceGroupResult, error) {
			args := v.(LookupConfigurationAssignmentsForResourceGroupArgs)
			r, err := LookupConfigurationAssignmentsForResourceGroup(ctx, &args, opts...)
			var s LookupConfigurationAssignmentsForResourceGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationAssignmentsForResourceGroupResultOutput)
}

type LookupConfigurationAssignmentsForResourceGroupOutputArgs struct {
	// Configuration assignment name
	ConfigurationAssignmentName pulumi.StringInput `pulumi:"configurationAssignmentName"`
	// Resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConfigurationAssignmentsForResourceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationAssignmentsForResourceGroupArgs)(nil)).Elem()
}

// Configuration Assignment
type LookupConfigurationAssignmentsForResourceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationAssignmentsForResourceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationAssignmentsForResourceGroupResult)(nil)).Elem()
}

func (o LookupConfigurationAssignmentsForResourceGroupResultOutput) ToLookupConfigurationAssignmentsForResourceGroupResultOutput() LookupConfigurationAssignmentsForResourceGroupResultOutput {
	return o
}

func (o LookupConfigurationAssignmentsForResourceGroupResultOutput) ToLookupConfigurationAssignmentsForResourceGroupResultOutputWithContext(ctx context.Context) LookupConfigurationAssignmentsForResourceGroupResultOutput {
	return o
}

func (o LookupConfigurationAssignmentsForResourceGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupConfigurationAssignmentsForResourceGroupResult] {
	return pulumix.Output[LookupConfigurationAssignmentsForResourceGroupResult]{
		OutputState: o.OutputState,
	}
}

// Properties of the configuration assignment
func (o LookupConfigurationAssignmentsForResourceGroupResultOutput) Filter() ConfigurationAssignmentFilterPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForResourceGroupResult) *ConfigurationAssignmentFilterPropertiesResponse {
		return v.Filter
	}).(ConfigurationAssignmentFilterPropertiesResponsePtrOutput)
}

// Fully qualified identifier of the resource
func (o LookupConfigurationAssignmentsForResourceGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForResourceGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Location of the resource
func (o LookupConfigurationAssignmentsForResourceGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForResourceGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The maintenance configuration Id
func (o LookupConfigurationAssignmentsForResourceGroupResultOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForResourceGroupResult) *string {
		return v.MaintenanceConfigurationId
	}).(pulumi.StringPtrOutput)
}

// Name of the resource
func (o LookupConfigurationAssignmentsForResourceGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForResourceGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The unique resourceId
func (o LookupConfigurationAssignmentsForResourceGroupResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForResourceGroupResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupConfigurationAssignmentsForResourceGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForResourceGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the resource
func (o LookupConfigurationAssignmentsForResourceGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForResourceGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationAssignmentsForResourceGroupResultOutput{})
}
