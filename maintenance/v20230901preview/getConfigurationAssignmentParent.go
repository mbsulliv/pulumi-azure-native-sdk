// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get configuration assignment for resource..
func LookupConfigurationAssignmentParent(ctx *pulumi.Context, args *LookupConfigurationAssignmentParentArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationAssignmentParentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConfigurationAssignmentParentResult
	err := ctx.Invoke("azure-native:maintenance/v20230901preview:getConfigurationAssignmentParent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationAssignmentParentArgs struct {
	// Configuration assignment name
	ConfigurationAssignmentName string `pulumi:"configurationAssignmentName"`
	// Resource provider name
	ProviderName string `pulumi:"providerName"`
	// Resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource identifier
	ResourceName string `pulumi:"resourceName"`
	// Resource parent identifier
	ResourceParentName string `pulumi:"resourceParentName"`
	// Resource parent type
	ResourceParentType string `pulumi:"resourceParentType"`
	// Resource type
	ResourceType string `pulumi:"resourceType"`
}

// Configuration Assignment
type LookupConfigurationAssignmentParentResult struct {
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

func LookupConfigurationAssignmentParentOutput(ctx *pulumi.Context, args LookupConfigurationAssignmentParentOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationAssignmentParentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationAssignmentParentResult, error) {
			args := v.(LookupConfigurationAssignmentParentArgs)
			r, err := LookupConfigurationAssignmentParent(ctx, &args, opts...)
			var s LookupConfigurationAssignmentParentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationAssignmentParentResultOutput)
}

type LookupConfigurationAssignmentParentOutputArgs struct {
	// Configuration assignment name
	ConfigurationAssignmentName pulumi.StringInput `pulumi:"configurationAssignmentName"`
	// Resource provider name
	ProviderName pulumi.StringInput `pulumi:"providerName"`
	// Resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Resource identifier
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// Resource parent identifier
	ResourceParentName pulumi.StringInput `pulumi:"resourceParentName"`
	// Resource parent type
	ResourceParentType pulumi.StringInput `pulumi:"resourceParentType"`
	// Resource type
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
}

func (LookupConfigurationAssignmentParentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationAssignmentParentArgs)(nil)).Elem()
}

// Configuration Assignment
type LookupConfigurationAssignmentParentResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationAssignmentParentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationAssignmentParentResult)(nil)).Elem()
}

func (o LookupConfigurationAssignmentParentResultOutput) ToLookupConfigurationAssignmentParentResultOutput() LookupConfigurationAssignmentParentResultOutput {
	return o
}

func (o LookupConfigurationAssignmentParentResultOutput) ToLookupConfigurationAssignmentParentResultOutputWithContext(ctx context.Context) LookupConfigurationAssignmentParentResultOutput {
	return o
}

func (o LookupConfigurationAssignmentParentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupConfigurationAssignmentParentResult] {
	return pulumix.Output[LookupConfigurationAssignmentParentResult]{
		OutputState: o.OutputState,
	}
}

// Properties of the configuration assignment
func (o LookupConfigurationAssignmentParentResultOutput) Filter() ConfigurationAssignmentFilterPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) *ConfigurationAssignmentFilterPropertiesResponse {
		return v.Filter
	}).(ConfigurationAssignmentFilterPropertiesResponsePtrOutput)
}

// Fully qualified identifier of the resource
func (o LookupConfigurationAssignmentParentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Location of the resource
func (o LookupConfigurationAssignmentParentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The maintenance configuration Id
func (o LookupConfigurationAssignmentParentResultOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) *string { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

// Name of the resource
func (o LookupConfigurationAssignmentParentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The unique resourceId
func (o LookupConfigurationAssignmentParentResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupConfigurationAssignmentParentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the resource
func (o LookupConfigurationAssignmentParentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationAssignmentParentResultOutput{})
}
