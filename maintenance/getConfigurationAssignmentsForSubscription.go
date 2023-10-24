// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package maintenance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get configuration assignment for resource..
// Azure REST API version: 2023-04-01.
//
// Other available API versions: 2023-09-01-preview.
func LookupConfigurationAssignmentsForSubscription(ctx *pulumi.Context, args *LookupConfigurationAssignmentsForSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationAssignmentsForSubscriptionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConfigurationAssignmentsForSubscriptionResult
	err := ctx.Invoke("azure-native:maintenance:getConfigurationAssignmentsForSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationAssignmentsForSubscriptionArgs struct {
	// Configuration assignment name
	ConfigurationAssignmentName string `pulumi:"configurationAssignmentName"`
}

// Configuration Assignment
type LookupConfigurationAssignmentsForSubscriptionResult struct {
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

func LookupConfigurationAssignmentsForSubscriptionOutput(ctx *pulumi.Context, args LookupConfigurationAssignmentsForSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationAssignmentsForSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationAssignmentsForSubscriptionResult, error) {
			args := v.(LookupConfigurationAssignmentsForSubscriptionArgs)
			r, err := LookupConfigurationAssignmentsForSubscription(ctx, &args, opts...)
			var s LookupConfigurationAssignmentsForSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationAssignmentsForSubscriptionResultOutput)
}

type LookupConfigurationAssignmentsForSubscriptionOutputArgs struct {
	// Configuration assignment name
	ConfigurationAssignmentName pulumi.StringInput `pulumi:"configurationAssignmentName"`
}

func (LookupConfigurationAssignmentsForSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationAssignmentsForSubscriptionArgs)(nil)).Elem()
}

// Configuration Assignment
type LookupConfigurationAssignmentsForSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationAssignmentsForSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationAssignmentsForSubscriptionResult)(nil)).Elem()
}

func (o LookupConfigurationAssignmentsForSubscriptionResultOutput) ToLookupConfigurationAssignmentsForSubscriptionResultOutput() LookupConfigurationAssignmentsForSubscriptionResultOutput {
	return o
}

func (o LookupConfigurationAssignmentsForSubscriptionResultOutput) ToLookupConfigurationAssignmentsForSubscriptionResultOutputWithContext(ctx context.Context) LookupConfigurationAssignmentsForSubscriptionResultOutput {
	return o
}

func (o LookupConfigurationAssignmentsForSubscriptionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupConfigurationAssignmentsForSubscriptionResult] {
	return pulumix.Output[LookupConfigurationAssignmentsForSubscriptionResult]{
		OutputState: o.OutputState,
	}
}

// Properties of the configuration assignment
func (o LookupConfigurationAssignmentsForSubscriptionResultOutput) Filter() ConfigurationAssignmentFilterPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForSubscriptionResult) *ConfigurationAssignmentFilterPropertiesResponse {
		return v.Filter
	}).(ConfigurationAssignmentFilterPropertiesResponsePtrOutput)
}

// Fully qualified identifier of the resource
func (o LookupConfigurationAssignmentsForSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Location of the resource
func (o LookupConfigurationAssignmentsForSubscriptionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForSubscriptionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The maintenance configuration Id
func (o LookupConfigurationAssignmentsForSubscriptionResultOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForSubscriptionResult) *string {
		return v.MaintenanceConfigurationId
	}).(pulumi.StringPtrOutput)
}

// Name of the resource
func (o LookupConfigurationAssignmentsForSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The unique resourceId
func (o LookupConfigurationAssignmentsForSubscriptionResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForSubscriptionResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupConfigurationAssignmentsForSubscriptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForSubscriptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the resource
func (o LookupConfigurationAssignmentsForSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentsForSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationAssignmentsForSubscriptionResultOutput{})
}
