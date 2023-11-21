// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230515preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a single software update configuration by name.
func LookupSoftwareUpdateConfigurationByName(ctx *pulumi.Context, args *LookupSoftwareUpdateConfigurationByNameArgs, opts ...pulumi.InvokeOption) (*LookupSoftwareUpdateConfigurationByNameResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSoftwareUpdateConfigurationByNameResult
	err := ctx.Invoke("azure-native:automation/v20230515preview:getSoftwareUpdateConfigurationByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSoftwareUpdateConfigurationByNameArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the software update configuration to be created.
	SoftwareUpdateConfigurationName string `pulumi:"softwareUpdateConfigurationName"`
}

// Software update configuration properties.
type LookupSoftwareUpdateConfigurationByNameResult struct {
	// CreatedBy property, which only appears in the response.
	CreatedBy string `pulumi:"createdBy"`
	// Creation time of the resource, which only appears in the response.
	CreationTime string `pulumi:"creationTime"`
	// Details of provisioning error
	Error *ErrorResponseResponse `pulumi:"error"`
	// Resource Id.
	Id string `pulumi:"id"`
	// LastModifiedBy property, which only appears in the response.
	LastModifiedBy string `pulumi:"lastModifiedBy"`
	// Last time resource was modified, which only appears in the response.
	LastModifiedTime string `pulumi:"lastModifiedTime"`
	// Resource name.
	Name string `pulumi:"name"`
	// Provisioning state for the software update configuration, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// Schedule information for the Software update configuration
	ScheduleInfo SUCSchedulePropertiesResponse `pulumi:"scheduleInfo"`
	// Tasks information for the Software update configuration.
	Tasks *SoftwareUpdateConfigurationTasksResponse `pulumi:"tasks"`
	// Resource type
	Type string `pulumi:"type"`
	// update specific properties for the Software update configuration
	UpdateConfiguration UpdateConfigurationResponse `pulumi:"updateConfiguration"`
}

// Defaults sets the appropriate defaults for LookupSoftwareUpdateConfigurationByNameResult
func (val *LookupSoftwareUpdateConfigurationByNameResult) Defaults() *LookupSoftwareUpdateConfigurationByNameResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ScheduleInfo = *tmp.ScheduleInfo.Defaults()

	return &tmp
}

func LookupSoftwareUpdateConfigurationByNameOutput(ctx *pulumi.Context, args LookupSoftwareUpdateConfigurationByNameOutputArgs, opts ...pulumi.InvokeOption) LookupSoftwareUpdateConfigurationByNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSoftwareUpdateConfigurationByNameResult, error) {
			args := v.(LookupSoftwareUpdateConfigurationByNameArgs)
			r, err := LookupSoftwareUpdateConfigurationByName(ctx, &args, opts...)
			var s LookupSoftwareUpdateConfigurationByNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSoftwareUpdateConfigurationByNameResultOutput)
}

type LookupSoftwareUpdateConfigurationByNameOutputArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the software update configuration to be created.
	SoftwareUpdateConfigurationName pulumi.StringInput `pulumi:"softwareUpdateConfigurationName"`
}

func (LookupSoftwareUpdateConfigurationByNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSoftwareUpdateConfigurationByNameArgs)(nil)).Elem()
}

// Software update configuration properties.
type LookupSoftwareUpdateConfigurationByNameResultOutput struct{ *pulumi.OutputState }

func (LookupSoftwareUpdateConfigurationByNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSoftwareUpdateConfigurationByNameResult)(nil)).Elem()
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) ToLookupSoftwareUpdateConfigurationByNameResultOutput() LookupSoftwareUpdateConfigurationByNameResultOutput {
	return o
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) ToLookupSoftwareUpdateConfigurationByNameResultOutputWithContext(ctx context.Context) LookupSoftwareUpdateConfigurationByNameResultOutput {
	return o
}

// CreatedBy property, which only appears in the response.
func (o LookupSoftwareUpdateConfigurationByNameResultOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.CreatedBy }).(pulumi.StringOutput)
}

// Creation time of the resource, which only appears in the response.
func (o LookupSoftwareUpdateConfigurationByNameResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// Details of provisioning error
func (o LookupSoftwareUpdateConfigurationByNameResultOutput) Error() ErrorResponseResponsePtrOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) *ErrorResponseResponse { return v.Error }).(ErrorResponseResponsePtrOutput)
}

// Resource Id.
func (o LookupSoftwareUpdateConfigurationByNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.Id }).(pulumi.StringOutput)
}

// LastModifiedBy property, which only appears in the response.
func (o LookupSoftwareUpdateConfigurationByNameResultOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

// Last time resource was modified, which only appears in the response.
func (o LookupSoftwareUpdateConfigurationByNameResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupSoftwareUpdateConfigurationByNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state for the software update configuration, which only appears in the response.
func (o LookupSoftwareUpdateConfigurationByNameResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Schedule information for the Software update configuration
func (o LookupSoftwareUpdateConfigurationByNameResultOutput) ScheduleInfo() SUCSchedulePropertiesResponseOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) SUCSchedulePropertiesResponse {
		return v.ScheduleInfo
	}).(SUCSchedulePropertiesResponseOutput)
}

// Tasks information for the Software update configuration.
func (o LookupSoftwareUpdateConfigurationByNameResultOutput) Tasks() SoftwareUpdateConfigurationTasksResponsePtrOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) *SoftwareUpdateConfigurationTasksResponse {
		return v.Tasks
	}).(SoftwareUpdateConfigurationTasksResponsePtrOutput)
}

// Resource type
func (o LookupSoftwareUpdateConfigurationByNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.Type }).(pulumi.StringOutput)
}

// update specific properties for the Software update configuration
func (o LookupSoftwareUpdateConfigurationByNameResultOutput) UpdateConfiguration() UpdateConfigurationResponseOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) UpdateConfigurationResponse {
		return v.UpdateConfiguration
	}).(UpdateConfigurationResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSoftwareUpdateConfigurationByNameResultOutput{})
}