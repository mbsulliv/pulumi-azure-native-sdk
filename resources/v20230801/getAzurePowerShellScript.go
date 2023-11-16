// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a deployment script with a given name.
func LookupAzurePowerShellScript(ctx *pulumi.Context, args *LookupAzurePowerShellScriptArgs, opts ...pulumi.InvokeOption) (*LookupAzurePowerShellScriptResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAzurePowerShellScriptResult
	err := ctx.Invoke("azure-native:resources/v20230801:getAzurePowerShellScript", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAzurePowerShellScriptArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment script.
	ScriptName string `pulumi:"scriptName"`
}

// Object model for the Azure PowerShell script.
type LookupAzurePowerShellScriptResult struct {
	// Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2'
	Arguments *string `pulumi:"arguments"`
	// Azure PowerShell module version to be used.
	AzPowerShellVersion string `pulumi:"azPowerShellVersion"`
	// The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
	CleanupPreference *string `pulumi:"cleanupPreference"`
	// Container settings.
	ContainerSettings *ContainerConfigurationResponse `pulumi:"containerSettings"`
	// The environment variables to pass over to the script.
	EnvironmentVariables []EnvironmentVariableResponse `pulumi:"environmentVariables"`
	// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// String Id used to locate any resource on Azure.
	Id string `pulumi:"id"`
	// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// Type of the script.
	// Expected value is 'AzurePowerShell'.
	Kind string `pulumi:"kind"`
	// The location of the ACI and the storage account for the deployment script.
	Location string `pulumi:"location"`
	// Name of this resource.
	Name string `pulumi:"name"`
	// List of script outputs.
	Outputs map[string]interface{} `pulumi:"outputs"`
	// Uri for the script. This is the entry point for the external script.
	PrimaryScriptUri *string `pulumi:"primaryScriptUri"`
	// State of the script execution. This only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
	RetentionInterval string `pulumi:"retentionInterval"`
	// Script body.
	ScriptContent *string `pulumi:"scriptContent"`
	// Contains the results of script execution.
	Status ScriptStatusResponse `pulumi:"status"`
	// Storage Account settings.
	StorageAccountSettings *StorageAccountConfigurationResponse `pulumi:"storageAccountSettings"`
	// Supporting files for the external script.
	SupportingScriptUris []string `pulumi:"supportingScriptUris"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
	Timeout *string `pulumi:"timeout"`
	// Type of this resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupAzurePowerShellScriptResult
func (val *LookupAzurePowerShellScriptResult) Defaults() *LookupAzurePowerShellScriptResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.CleanupPreference == nil {
		cleanupPreference_ := "Always"
		tmp.CleanupPreference = &cleanupPreference_
	}
	if tmp.Timeout == nil {
		timeout_ := "P1D"
		tmp.Timeout = &timeout_
	}
	return &tmp
}

func LookupAzurePowerShellScriptOutput(ctx *pulumi.Context, args LookupAzurePowerShellScriptOutputArgs, opts ...pulumi.InvokeOption) LookupAzurePowerShellScriptResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAzurePowerShellScriptResult, error) {
			args := v.(LookupAzurePowerShellScriptArgs)
			r, err := LookupAzurePowerShellScript(ctx, &args, opts...)
			var s LookupAzurePowerShellScriptResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAzurePowerShellScriptResultOutput)
}

type LookupAzurePowerShellScriptOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment script.
	ScriptName pulumi.StringInput `pulumi:"scriptName"`
}

func (LookupAzurePowerShellScriptOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzurePowerShellScriptArgs)(nil)).Elem()
}

// Object model for the Azure PowerShell script.
type LookupAzurePowerShellScriptResultOutput struct{ *pulumi.OutputState }

func (LookupAzurePowerShellScriptResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzurePowerShellScriptResult)(nil)).Elem()
}

func (o LookupAzurePowerShellScriptResultOutput) ToLookupAzurePowerShellScriptResultOutput() LookupAzurePowerShellScriptResultOutput {
	return o
}

func (o LookupAzurePowerShellScriptResultOutput) ToLookupAzurePowerShellScriptResultOutputWithContext(ctx context.Context) LookupAzurePowerShellScriptResultOutput {
	return o
}

// Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2'
func (o LookupAzurePowerShellScriptResultOutput) Arguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *string { return v.Arguments }).(pulumi.StringPtrOutput)
}

// Azure PowerShell module version to be used.
func (o LookupAzurePowerShellScriptResultOutput) AzPowerShellVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.AzPowerShellVersion }).(pulumi.StringOutput)
}

// The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
func (o LookupAzurePowerShellScriptResultOutput) CleanupPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *string { return v.CleanupPreference }).(pulumi.StringPtrOutput)
}

// Container settings.
func (o LookupAzurePowerShellScriptResultOutput) ContainerSettings() ContainerConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *ContainerConfigurationResponse { return v.ContainerSettings }).(ContainerConfigurationResponsePtrOutput)
}

// The environment variables to pass over to the script.
func (o LookupAzurePowerShellScriptResultOutput) EnvironmentVariables() EnvironmentVariableResponseArrayOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) []EnvironmentVariableResponse { return v.EnvironmentVariables }).(EnvironmentVariableResponseArrayOutput)
}

// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
func (o LookupAzurePowerShellScriptResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

// String Id used to locate any resource on Azure.
func (o LookupAzurePowerShellScriptResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.Id }).(pulumi.StringOutput)
}

// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
func (o LookupAzurePowerShellScriptResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Type of the script.
// Expected value is 'AzurePowerShell'.
func (o LookupAzurePowerShellScriptResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The location of the ACI and the storage account for the deployment script.
func (o LookupAzurePowerShellScriptResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.Location }).(pulumi.StringOutput)
}

// Name of this resource.
func (o LookupAzurePowerShellScriptResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of script outputs.
func (o LookupAzurePowerShellScriptResultOutput) Outputs() pulumi.MapOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) map[string]interface{} { return v.Outputs }).(pulumi.MapOutput)
}

// Uri for the script. This is the entry point for the external script.
func (o LookupAzurePowerShellScriptResultOutput) PrimaryScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *string { return v.PrimaryScriptUri }).(pulumi.StringPtrOutput)
}

// State of the script execution. This only appears in the response.
func (o LookupAzurePowerShellScriptResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
func (o LookupAzurePowerShellScriptResultOutput) RetentionInterval() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.RetentionInterval }).(pulumi.StringOutput)
}

// Script body.
func (o LookupAzurePowerShellScriptResultOutput) ScriptContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *string { return v.ScriptContent }).(pulumi.StringPtrOutput)
}

// Contains the results of script execution.
func (o LookupAzurePowerShellScriptResultOutput) Status() ScriptStatusResponseOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) ScriptStatusResponse { return v.Status }).(ScriptStatusResponseOutput)
}

// Storage Account settings.
func (o LookupAzurePowerShellScriptResultOutput) StorageAccountSettings() StorageAccountConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *StorageAccountConfigurationResponse {
		return v.StorageAccountSettings
	}).(StorageAccountConfigurationResponsePtrOutput)
}

// Supporting files for the external script.
func (o LookupAzurePowerShellScriptResultOutput) SupportingScriptUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) []string { return v.SupportingScriptUris }).(pulumi.StringArrayOutput)
}

// The system metadata related to this resource.
func (o LookupAzurePowerShellScriptResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAzurePowerShellScriptResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
func (o LookupAzurePowerShellScriptResultOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *string { return v.Timeout }).(pulumi.StringPtrOutput)
}

// Type of this resource.
func (o LookupAzurePowerShellScriptResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzurePowerShellScriptResultOutput{})
}
