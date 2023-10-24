// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package virtualmachineimages

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get information about a virtual machine image template
// Azure REST API version: 2022-07-01.
//
// Other available API versions: 2018-02-01-preview, 2019-05-01-preview.
func LookupVirtualMachineImageTemplate(ctx *pulumi.Context, args *LookupVirtualMachineImageTemplateArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineImageTemplateResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualMachineImageTemplateResult
	err := ctx.Invoke("azure-native:virtualmachineimages:getVirtualMachineImageTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualMachineImageTemplateArgs struct {
	// The name of the image Template
	ImageTemplateName string `pulumi:"imageTemplateName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Image template is an ARM resource managed by Microsoft.VirtualMachineImages provider
type LookupVirtualMachineImageTemplateResult struct {
	// Maximum duration to wait while building the image template (includes all customizations, optimization, validations, and distributions). Omit or specify 0 to use the default (4 hours).
	BuildTimeoutInMinutes *int `pulumi:"buildTimeoutInMinutes"`
	// Specifies the properties used to describe the customization steps of the image, like Image source etc
	Customize []interface{} `pulumi:"customize"`
	// The distribution targets where the image output needs to go to.
	Distribute []interface{} `pulumi:"distribute"`
	// The staging resource group id in the same subscription as the image template that will be used to build the image. This read-only field differs from 'stagingResourceGroup' only if the value specified in the 'stagingResourceGroup' field is empty.
	ExactStagingResourceGroup string `pulumi:"exactStagingResourceGroup"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The identity of the image template, if configured.
	Identity ImageTemplateIdentityResponse `pulumi:"identity"`
	// State of 'run' that is currently executing or was last executed.
	LastRunStatus ImageTemplateLastRunStatusResponse `pulumi:"lastRunStatus"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Specifies optimization to be performed on image.
	Optimize *ImageTemplatePropertiesResponseOptimize `pulumi:"optimize"`
	// Provisioning error, if any
	ProvisioningError ProvisioningErrorResponse `pulumi:"provisioningError"`
	// Provisioning state of the resource
	ProvisioningState string `pulumi:"provisioningState"`
	// Specifies the properties used to describe the source image.
	Source interface{} `pulumi:"source"`
	// The staging resource group id in the same subscription as the image template that will be used to build the image. If this field is empty, a resource group with a random name will be created. If the resource group specified in this field doesn't exist, it will be created with the same name. If the resource group specified exists, it must be empty and in the same region as the image template. The resource group created will be deleted during template deletion if this field is empty or the resource group specified doesn't exist, but if the resource group specified exists the resources created in the resource group will be deleted during template deletion and the resource group itself will remain.
	StagingResourceGroup *string `pulumi:"stagingResourceGroup"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Configuration options and list of validations to be performed on the resulting image.
	Validate *ImageTemplatePropertiesResponseValidate `pulumi:"validate"`
	// Describes how virtual machine is set up to build images
	VmProfile *ImageTemplateVmProfileResponse `pulumi:"vmProfile"`
}

// Defaults sets the appropriate defaults for LookupVirtualMachineImageTemplateResult
func (val *LookupVirtualMachineImageTemplateResult) Defaults() *LookupVirtualMachineImageTemplateResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.BuildTimeoutInMinutes == nil {
		buildTimeoutInMinutes_ := 0
		tmp.BuildTimeoutInMinutes = &buildTimeoutInMinutes_
	}
	tmp.Validate = tmp.Validate.Defaults()

	tmp.VmProfile = tmp.VmProfile.Defaults()

	return &tmp
}

func LookupVirtualMachineImageTemplateOutput(ctx *pulumi.Context, args LookupVirtualMachineImageTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineImageTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineImageTemplateResult, error) {
			args := v.(LookupVirtualMachineImageTemplateArgs)
			r, err := LookupVirtualMachineImageTemplate(ctx, &args, opts...)
			var s LookupVirtualMachineImageTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineImageTemplateResultOutput)
}

type LookupVirtualMachineImageTemplateOutputArgs struct {
	// The name of the image Template
	ImageTemplateName pulumi.StringInput `pulumi:"imageTemplateName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupVirtualMachineImageTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineImageTemplateArgs)(nil)).Elem()
}

// Image template is an ARM resource managed by Microsoft.VirtualMachineImages provider
type LookupVirtualMachineImageTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineImageTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineImageTemplateResult)(nil)).Elem()
}

func (o LookupVirtualMachineImageTemplateResultOutput) ToLookupVirtualMachineImageTemplateResultOutput() LookupVirtualMachineImageTemplateResultOutput {
	return o
}

func (o LookupVirtualMachineImageTemplateResultOutput) ToLookupVirtualMachineImageTemplateResultOutputWithContext(ctx context.Context) LookupVirtualMachineImageTemplateResultOutput {
	return o
}

func (o LookupVirtualMachineImageTemplateResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualMachineImageTemplateResult] {
	return pulumix.Output[LookupVirtualMachineImageTemplateResult]{
		OutputState: o.OutputState,
	}
}

// Maximum duration to wait while building the image template (includes all customizations, optimization, validations, and distributions). Omit or specify 0 to use the default (4 hours).
func (o LookupVirtualMachineImageTemplateResultOutput) BuildTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) *int { return v.BuildTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

// Specifies the properties used to describe the customization steps of the image, like Image source etc
func (o LookupVirtualMachineImageTemplateResultOutput) Customize() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) []interface{} { return v.Customize }).(pulumi.ArrayOutput)
}

// The distribution targets where the image output needs to go to.
func (o LookupVirtualMachineImageTemplateResultOutput) Distribute() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) []interface{} { return v.Distribute }).(pulumi.ArrayOutput)
}

// The staging resource group id in the same subscription as the image template that will be used to build the image. This read-only field differs from 'stagingResourceGroup' only if the value specified in the 'stagingResourceGroup' field is empty.
func (o LookupVirtualMachineImageTemplateResultOutput) ExactStagingResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) string { return v.ExactStagingResourceGroup }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupVirtualMachineImageTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the image template, if configured.
func (o LookupVirtualMachineImageTemplateResultOutput) Identity() ImageTemplateIdentityResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) ImageTemplateIdentityResponse { return v.Identity }).(ImageTemplateIdentityResponseOutput)
}

// State of 'run' that is currently executing or was last executed.
func (o LookupVirtualMachineImageTemplateResultOutput) LastRunStatus() ImageTemplateLastRunStatusResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) ImageTemplateLastRunStatusResponse {
		return v.LastRunStatus
	}).(ImageTemplateLastRunStatusResponseOutput)
}

// The geo-location where the resource lives
func (o LookupVirtualMachineImageTemplateResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupVirtualMachineImageTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies optimization to be performed on image.
func (o LookupVirtualMachineImageTemplateResultOutput) Optimize() ImageTemplatePropertiesResponseOptimizePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) *ImageTemplatePropertiesResponseOptimize {
		return v.Optimize
	}).(ImageTemplatePropertiesResponseOptimizePtrOutput)
}

// Provisioning error, if any
func (o LookupVirtualMachineImageTemplateResultOutput) ProvisioningError() ProvisioningErrorResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) ProvisioningErrorResponse { return v.ProvisioningError }).(ProvisioningErrorResponseOutput)
}

// Provisioning state of the resource
func (o LookupVirtualMachineImageTemplateResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Specifies the properties used to describe the source image.
func (o LookupVirtualMachineImageTemplateResultOutput) Source() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) interface{} { return v.Source }).(pulumi.AnyOutput)
}

// The staging resource group id in the same subscription as the image template that will be used to build the image. If this field is empty, a resource group with a random name will be created. If the resource group specified in this field doesn't exist, it will be created with the same name. If the resource group specified exists, it must be empty and in the same region as the image template. The resource group created will be deleted during template deletion if this field is empty or the resource group specified doesn't exist, but if the resource group specified exists the resources created in the resource group will be deleted during template deletion and the resource group itself will remain.
func (o LookupVirtualMachineImageTemplateResultOutput) StagingResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) *string { return v.StagingResourceGroup }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupVirtualMachineImageTemplateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupVirtualMachineImageTemplateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVirtualMachineImageTemplateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) string { return v.Type }).(pulumi.StringOutput)
}

// Configuration options and list of validations to be performed on the resulting image.
func (o LookupVirtualMachineImageTemplateResultOutput) Validate() ImageTemplatePropertiesResponseValidatePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) *ImageTemplatePropertiesResponseValidate {
		return v.Validate
	}).(ImageTemplatePropertiesResponseValidatePtrOutput)
}

// Describes how virtual machine is set up to build images
func (o LookupVirtualMachineImageTemplateResultOutput) VmProfile() ImageTemplateVmProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) *ImageTemplateVmProfileResponse { return v.VmProfile }).(ImageTemplateVmProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineImageTemplateResultOutput{})
}
