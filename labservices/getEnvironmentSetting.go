// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package labservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get environment setting
// Azure REST API version: 2018-10-15.
func LookupEnvironmentSetting(ctx *pulumi.Context, args *LookupEnvironmentSettingArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentSettingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentSettingResult
	err := ctx.Invoke("azure-native:labservices:getEnvironmentSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentSettingArgs struct {
	// The name of the environment Setting.
	EnvironmentSettingName string `pulumi:"environmentSettingName"`
	// Specify the $expand query. Example: 'properties($select=publishingState)'
	Expand *string `pulumi:"expand"`
	// The name of the lab Account.
	LabAccountName string `pulumi:"labAccountName"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents settings of an environment, from which environment instances would be created
type LookupEnvironmentSettingResult struct {
	// Describes the user's progress in configuring their environment setting
	ConfigurationState *string `pulumi:"configurationState"`
	// Describes the environment and its resource settings
	Description *string `pulumi:"description"`
	// The identifier of the resource.
	Id string `pulumi:"id"`
	// Time when the template VM was last changed.
	LastChanged string `pulumi:"lastChanged"`
	// Time when the template VM was last sent for publishing.
	LastPublished string `pulumi:"lastPublished"`
	// The details of the latest operation. ex: status, error
	LatestOperationResult LatestOperationResultResponse `pulumi:"latestOperationResult"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Describes the readiness of this environment setting
	PublishingState string `pulumi:"publishingState"`
	// The resource specific settings
	ResourceSettings ResourceSettingsResponse `pulumi:"resourceSettings"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Brief title describing the environment and its resource settings
	Title *string `pulumi:"title"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
}

func LookupEnvironmentSettingOutput(ctx *pulumi.Context, args LookupEnvironmentSettingOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentSettingResult, error) {
			args := v.(LookupEnvironmentSettingArgs)
			r, err := LookupEnvironmentSetting(ctx, &args, opts...)
			var s LookupEnvironmentSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentSettingResultOutput)
}

type LookupEnvironmentSettingOutputArgs struct {
	// The name of the environment Setting.
	EnvironmentSettingName pulumi.StringInput `pulumi:"environmentSettingName"`
	// Specify the $expand query. Example: 'properties($select=publishingState)'
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the lab Account.
	LabAccountName pulumi.StringInput `pulumi:"labAccountName"`
	// The name of the lab.
	LabName pulumi.StringInput `pulumi:"labName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEnvironmentSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentSettingArgs)(nil)).Elem()
}

// Represents settings of an environment, from which environment instances would be created
type LookupEnvironmentSettingResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentSettingResult)(nil)).Elem()
}

func (o LookupEnvironmentSettingResultOutput) ToLookupEnvironmentSettingResultOutput() LookupEnvironmentSettingResultOutput {
	return o
}

func (o LookupEnvironmentSettingResultOutput) ToLookupEnvironmentSettingResultOutputWithContext(ctx context.Context) LookupEnvironmentSettingResultOutput {
	return o
}

// Describes the user's progress in configuring their environment setting
func (o LookupEnvironmentSettingResultOutput) ConfigurationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) *string { return v.ConfigurationState }).(pulumi.StringPtrOutput)
}

// Describes the environment and its resource settings
func (o LookupEnvironmentSettingResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The identifier of the resource.
func (o LookupEnvironmentSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Time when the template VM was last changed.
func (o LookupEnvironmentSettingResultOutput) LastChanged() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) string { return v.LastChanged }).(pulumi.StringOutput)
}

// Time when the template VM was last sent for publishing.
func (o LookupEnvironmentSettingResultOutput) LastPublished() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) string { return v.LastPublished }).(pulumi.StringOutput)
}

// The details of the latest operation. ex: status, error
func (o LookupEnvironmentSettingResultOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) LatestOperationResultResponse { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

// The location of the resource.
func (o LookupEnvironmentSettingResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupEnvironmentSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning status of the resource.
func (o LookupEnvironmentSettingResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Describes the readiness of this environment setting
func (o LookupEnvironmentSettingResultOutput) PublishingState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) string { return v.PublishingState }).(pulumi.StringOutput)
}

// The resource specific settings
func (o LookupEnvironmentSettingResultOutput) ResourceSettings() ResourceSettingsResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) ResourceSettingsResponse { return v.ResourceSettings }).(ResourceSettingsResponseOutput)
}

// The tags of the resource.
func (o LookupEnvironmentSettingResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Brief title describing the environment and its resource settings
func (o LookupEnvironmentSettingResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o LookupEnvironmentSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o LookupEnvironmentSettingResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentSettingResultOutput{})
}
