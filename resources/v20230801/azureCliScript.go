// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Object model for the Azure CLI script.
type AzureCliScript struct {
	pulumi.CustomResourceState

	// Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2'
	Arguments pulumi.StringPtrOutput `pulumi:"arguments"`
	// Azure CLI module version to be used.
	AzCliVersion pulumi.StringOutput `pulumi:"azCliVersion"`
	// The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
	CleanupPreference pulumi.StringPtrOutput `pulumi:"cleanupPreference"`
	// Container settings.
	ContainerSettings ContainerConfigurationResponsePtrOutput `pulumi:"containerSettings"`
	// The environment variables to pass over to the script.
	EnvironmentVariables EnvironmentVariableResponseArrayOutput `pulumi:"environmentVariables"`
	// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
	ForceUpdateTag pulumi.StringPtrOutput `pulumi:"forceUpdateTag"`
	// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Type of the script.
	// Expected value is 'AzureCLI'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The location of the ACI and the storage account for the deployment script.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of script outputs.
	Outputs pulumi.MapOutput `pulumi:"outputs"`
	// Uri for the script. This is the entry point for the external script.
	PrimaryScriptUri pulumi.StringPtrOutput `pulumi:"primaryScriptUri"`
	// State of the script execution. This only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
	RetentionInterval pulumi.StringOutput `pulumi:"retentionInterval"`
	// Script body.
	ScriptContent pulumi.StringPtrOutput `pulumi:"scriptContent"`
	// Contains the results of script execution.
	Status ScriptStatusResponseOutput `pulumi:"status"`
	// Storage Account settings.
	StorageAccountSettings StorageAccountConfigurationResponsePtrOutput `pulumi:"storageAccountSettings"`
	// Supporting files for the external script.
	SupportingScriptUris pulumi.StringArrayOutput `pulumi:"supportingScriptUris"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
	Timeout pulumi.StringPtrOutput `pulumi:"timeout"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAzureCliScript registers a new resource with the given unique name, arguments, and options.
func NewAzureCliScript(ctx *pulumi.Context,
	name string, args *AzureCliScriptArgs, opts ...pulumi.ResourceOption) (*AzureCliScript, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzCliVersion == nil {
		return nil, errors.New("invalid value for required argument 'AzCliVersion'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RetentionInterval == nil {
		return nil, errors.New("invalid value for required argument 'RetentionInterval'")
	}
	if args.CleanupPreference == nil {
		args.CleanupPreference = pulumi.StringPtr("Always")
	}
	args.Kind = pulumi.String("AzureCLI")
	if args.Timeout == nil {
		args.Timeout = pulumi.StringPtr("P1D")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:AzureCliScript"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001preview:AzureCliScript"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:AzureCliScript"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AzureCliScript
	err := ctx.RegisterResource("azure-native:resources/v20230801:AzureCliScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureCliScript gets an existing AzureCliScript resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureCliScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureCliScriptState, opts ...pulumi.ResourceOption) (*AzureCliScript, error) {
	var resource AzureCliScript
	err := ctx.ReadResource("azure-native:resources/v20230801:AzureCliScript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureCliScript resources.
type azureCliScriptState struct {
}

type AzureCliScriptState struct {
}

func (AzureCliScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureCliScriptState)(nil)).Elem()
}

type azureCliScriptArgs struct {
	// Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2'
	Arguments *string `pulumi:"arguments"`
	// Azure CLI module version to be used.
	AzCliVersion string `pulumi:"azCliVersion"`
	// The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
	CleanupPreference *string `pulumi:"cleanupPreference"`
	// Container settings.
	ContainerSettings *ContainerConfiguration `pulumi:"containerSettings"`
	// The environment variables to pass over to the script.
	EnvironmentVariables []EnvironmentVariable `pulumi:"environmentVariables"`
	// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Type of the script.
	// Expected value is 'AzureCLI'.
	Kind string `pulumi:"kind"`
	// The location of the ACI and the storage account for the deployment script.
	Location *string `pulumi:"location"`
	// Uri for the script. This is the entry point for the external script.
	PrimaryScriptUri *string `pulumi:"primaryScriptUri"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
	RetentionInterval string `pulumi:"retentionInterval"`
	// Script body.
	ScriptContent *string `pulumi:"scriptContent"`
	// Name of the deployment script.
	ScriptName *string `pulumi:"scriptName"`
	// Storage Account settings.
	StorageAccountSettings *StorageAccountConfiguration `pulumi:"storageAccountSettings"`
	// Supporting files for the external script.
	SupportingScriptUris []string `pulumi:"supportingScriptUris"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
	Timeout *string `pulumi:"timeout"`
}

// The set of arguments for constructing a AzureCliScript resource.
type AzureCliScriptArgs struct {
	// Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2'
	Arguments pulumi.StringPtrInput
	// Azure CLI module version to be used.
	AzCliVersion pulumi.StringInput
	// The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
	CleanupPreference pulumi.StringPtrInput
	// Container settings.
	ContainerSettings ContainerConfigurationPtrInput
	// The environment variables to pass over to the script.
	EnvironmentVariables EnvironmentVariableArrayInput
	// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
	ForceUpdateTag pulumi.StringPtrInput
	// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
	Identity ManagedServiceIdentityPtrInput
	// Type of the script.
	// Expected value is 'AzureCLI'.
	Kind pulumi.StringInput
	// The location of the ACI and the storage account for the deployment script.
	Location pulumi.StringPtrInput
	// Uri for the script. This is the entry point for the external script.
	PrimaryScriptUri pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
	RetentionInterval pulumi.StringInput
	// Script body.
	ScriptContent pulumi.StringPtrInput
	// Name of the deployment script.
	ScriptName pulumi.StringPtrInput
	// Storage Account settings.
	StorageAccountSettings StorageAccountConfigurationPtrInput
	// Supporting files for the external script.
	SupportingScriptUris pulumi.StringArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
	Timeout pulumi.StringPtrInput
}

func (AzureCliScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureCliScriptArgs)(nil)).Elem()
}

type AzureCliScriptInput interface {
	pulumi.Input

	ToAzureCliScriptOutput() AzureCliScriptOutput
	ToAzureCliScriptOutputWithContext(ctx context.Context) AzureCliScriptOutput
}

func (*AzureCliScript) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureCliScript)(nil)).Elem()
}

func (i *AzureCliScript) ToAzureCliScriptOutput() AzureCliScriptOutput {
	return i.ToAzureCliScriptOutputWithContext(context.Background())
}

func (i *AzureCliScript) ToAzureCliScriptOutputWithContext(ctx context.Context) AzureCliScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureCliScriptOutput)
}

type AzureCliScriptOutput struct{ *pulumi.OutputState }

func (AzureCliScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureCliScript)(nil)).Elem()
}

func (o AzureCliScriptOutput) ToAzureCliScriptOutput() AzureCliScriptOutput {
	return o
}

func (o AzureCliScriptOutput) ToAzureCliScriptOutputWithContext(ctx context.Context) AzureCliScriptOutput {
	return o
}

// Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2'
func (o AzureCliScriptOutput) Arguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringPtrOutput { return v.Arguments }).(pulumi.StringPtrOutput)
}

// Azure CLI module version to be used.
func (o AzureCliScriptOutput) AzCliVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringOutput { return v.AzCliVersion }).(pulumi.StringOutput)
}

// The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
func (o AzureCliScriptOutput) CleanupPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringPtrOutput { return v.CleanupPreference }).(pulumi.StringPtrOutput)
}

// Container settings.
func (o AzureCliScriptOutput) ContainerSettings() ContainerConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AzureCliScript) ContainerConfigurationResponsePtrOutput { return v.ContainerSettings }).(ContainerConfigurationResponsePtrOutput)
}

// The environment variables to pass over to the script.
func (o AzureCliScriptOutput) EnvironmentVariables() EnvironmentVariableResponseArrayOutput {
	return o.ApplyT(func(v *AzureCliScript) EnvironmentVariableResponseArrayOutput { return v.EnvironmentVariables }).(EnvironmentVariableResponseArrayOutput)
}

// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
func (o AzureCliScriptOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
func (o AzureCliScriptOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *AzureCliScript) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Type of the script.
// Expected value is 'AzureCLI'.
func (o AzureCliScriptOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The location of the ACI and the storage account for the deployment script.
func (o AzureCliScriptOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of this resource.
func (o AzureCliScriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of script outputs.
func (o AzureCliScriptOutput) Outputs() pulumi.MapOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.MapOutput { return v.Outputs }).(pulumi.MapOutput)
}

// Uri for the script. This is the entry point for the external script.
func (o AzureCliScriptOutput) PrimaryScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringPtrOutput { return v.PrimaryScriptUri }).(pulumi.StringPtrOutput)
}

// State of the script execution. This only appears in the response.
func (o AzureCliScriptOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
func (o AzureCliScriptOutput) RetentionInterval() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringOutput { return v.RetentionInterval }).(pulumi.StringOutput)
}

// Script body.
func (o AzureCliScriptOutput) ScriptContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringPtrOutput { return v.ScriptContent }).(pulumi.StringPtrOutput)
}

// Contains the results of script execution.
func (o AzureCliScriptOutput) Status() ScriptStatusResponseOutput {
	return o.ApplyT(func(v *AzureCliScript) ScriptStatusResponseOutput { return v.Status }).(ScriptStatusResponseOutput)
}

// Storage Account settings.
func (o AzureCliScriptOutput) StorageAccountSettings() StorageAccountConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AzureCliScript) StorageAccountConfigurationResponsePtrOutput { return v.StorageAccountSettings }).(StorageAccountConfigurationResponsePtrOutput)
}

// Supporting files for the external script.
func (o AzureCliScriptOutput) SupportingScriptUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringArrayOutput { return v.SupportingScriptUris }).(pulumi.StringArrayOutput)
}

// The system metadata related to this resource.
func (o AzureCliScriptOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AzureCliScript) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AzureCliScriptOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
func (o AzureCliScriptOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringPtrOutput { return v.Timeout }).(pulumi.StringPtrOutput)
}

// Type of this resource.
func (o AzureCliScriptOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureCliScriptOutput{})
}
