// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Software update configuration properties.
type SoftwareUpdateConfigurationByName struct {
	pulumi.CustomResourceState

	// CreatedBy property, which only appears in the response.
	CreatedBy pulumi.StringOutput `pulumi:"createdBy"`
	// Creation time of the resource, which only appears in the response.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Details of provisioning error
	Error ErrorResponseResponsePtrOutput `pulumi:"error"`
	// LastModifiedBy property, which only appears in the response.
	LastModifiedBy pulumi.StringOutput `pulumi:"lastModifiedBy"`
	// Last time resource was modified, which only appears in the response.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state for the software update configuration, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Schedule information for the Software update configuration
	ScheduleInfo SUCSchedulePropertiesResponseOutput `pulumi:"scheduleInfo"`
	// Tasks information for the Software update configuration.
	Tasks SoftwareUpdateConfigurationTasksResponsePtrOutput `pulumi:"tasks"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// update specific properties for the Software update configuration
	UpdateConfiguration UpdateConfigurationResponseOutput `pulumi:"updateConfiguration"`
}

// NewSoftwareUpdateConfigurationByName registers a new resource with the given unique name, arguments, and options.
func NewSoftwareUpdateConfigurationByName(ctx *pulumi.Context,
	name string, args *SoftwareUpdateConfigurationByNameArgs, opts ...pulumi.ResourceOption) (*SoftwareUpdateConfigurationByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScheduleInfo == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleInfo'")
	}
	if args.UpdateConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'UpdateConfiguration'")
	}
	args.ScheduleInfo = args.ScheduleInfo.ToSUCSchedulePropertiesOutput().ApplyT(func(v SUCScheduleProperties) SUCScheduleProperties { return *v.Defaults() }).(SUCSchedulePropertiesOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:SoftwareUpdateConfigurationByName"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20170515preview:SoftwareUpdateConfigurationByName"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SoftwareUpdateConfigurationByName
	err := ctx.RegisterResource("azure-native:automation/v20190601:SoftwareUpdateConfigurationByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSoftwareUpdateConfigurationByName gets an existing SoftwareUpdateConfigurationByName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSoftwareUpdateConfigurationByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SoftwareUpdateConfigurationByNameState, opts ...pulumi.ResourceOption) (*SoftwareUpdateConfigurationByName, error) {
	var resource SoftwareUpdateConfigurationByName
	err := ctx.ReadResource("azure-native:automation/v20190601:SoftwareUpdateConfigurationByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SoftwareUpdateConfigurationByName resources.
type softwareUpdateConfigurationByNameState struct {
}

type SoftwareUpdateConfigurationByNameState struct {
}

func (SoftwareUpdateConfigurationByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*softwareUpdateConfigurationByNameState)(nil)).Elem()
}

type softwareUpdateConfigurationByNameArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Details of provisioning error
	Error *ErrorResponse `pulumi:"error"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Schedule information for the Software update configuration
	ScheduleInfo SUCScheduleProperties `pulumi:"scheduleInfo"`
	// The name of the software update configuration to be created.
	SoftwareUpdateConfigurationName *string `pulumi:"softwareUpdateConfigurationName"`
	// Tasks information for the Software update configuration.
	Tasks *SoftwareUpdateConfigurationTasks `pulumi:"tasks"`
	// update specific properties for the Software update configuration
	UpdateConfiguration UpdateConfiguration `pulumi:"updateConfiguration"`
}

// The set of arguments for constructing a SoftwareUpdateConfigurationByName resource.
type SoftwareUpdateConfigurationByNameArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Details of provisioning error
	Error ErrorResponsePtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Schedule information for the Software update configuration
	ScheduleInfo SUCSchedulePropertiesInput
	// The name of the software update configuration to be created.
	SoftwareUpdateConfigurationName pulumi.StringPtrInput
	// Tasks information for the Software update configuration.
	Tasks SoftwareUpdateConfigurationTasksPtrInput
	// update specific properties for the Software update configuration
	UpdateConfiguration UpdateConfigurationInput
}

func (SoftwareUpdateConfigurationByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*softwareUpdateConfigurationByNameArgs)(nil)).Elem()
}

type SoftwareUpdateConfigurationByNameInput interface {
	pulumi.Input

	ToSoftwareUpdateConfigurationByNameOutput() SoftwareUpdateConfigurationByNameOutput
	ToSoftwareUpdateConfigurationByNameOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationByNameOutput
}

func (*SoftwareUpdateConfigurationByName) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwareUpdateConfigurationByName)(nil)).Elem()
}

func (i *SoftwareUpdateConfigurationByName) ToSoftwareUpdateConfigurationByNameOutput() SoftwareUpdateConfigurationByNameOutput {
	return i.ToSoftwareUpdateConfigurationByNameOutputWithContext(context.Background())
}

func (i *SoftwareUpdateConfigurationByName) ToSoftwareUpdateConfigurationByNameOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftwareUpdateConfigurationByNameOutput)
}

func (i *SoftwareUpdateConfigurationByName) ToOutput(ctx context.Context) pulumix.Output[*SoftwareUpdateConfigurationByName] {
	return pulumix.Output[*SoftwareUpdateConfigurationByName]{
		OutputState: i.ToSoftwareUpdateConfigurationByNameOutputWithContext(ctx).OutputState,
	}
}

type SoftwareUpdateConfigurationByNameOutput struct{ *pulumi.OutputState }

func (SoftwareUpdateConfigurationByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwareUpdateConfigurationByName)(nil)).Elem()
}

func (o SoftwareUpdateConfigurationByNameOutput) ToSoftwareUpdateConfigurationByNameOutput() SoftwareUpdateConfigurationByNameOutput {
	return o
}

func (o SoftwareUpdateConfigurationByNameOutput) ToSoftwareUpdateConfigurationByNameOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationByNameOutput {
	return o
}

func (o SoftwareUpdateConfigurationByNameOutput) ToOutput(ctx context.Context) pulumix.Output[*SoftwareUpdateConfigurationByName] {
	return pulumix.Output[*SoftwareUpdateConfigurationByName]{
		OutputState: o.OutputState,
	}
}

// CreatedBy property, which only appears in the response.
func (o SoftwareUpdateConfigurationByNameOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

// Creation time of the resource, which only appears in the response.
func (o SoftwareUpdateConfigurationByNameOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Details of provisioning error
func (o SoftwareUpdateConfigurationByNameOutput) Error() ErrorResponseResponsePtrOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) ErrorResponseResponsePtrOutput { return v.Error }).(ErrorResponseResponsePtrOutput)
}

// LastModifiedBy property, which only appears in the response.
func (o SoftwareUpdateConfigurationByNameOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) pulumi.StringOutput { return v.LastModifiedBy }).(pulumi.StringOutput)
}

// Last time resource was modified, which only appears in the response.
func (o SoftwareUpdateConfigurationByNameOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// Resource name.
func (o SoftwareUpdateConfigurationByNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state for the software update configuration, which only appears in the response.
func (o SoftwareUpdateConfigurationByNameOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Schedule information for the Software update configuration
func (o SoftwareUpdateConfigurationByNameOutput) ScheduleInfo() SUCSchedulePropertiesResponseOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) SUCSchedulePropertiesResponseOutput { return v.ScheduleInfo }).(SUCSchedulePropertiesResponseOutput)
}

// Tasks information for the Software update configuration.
func (o SoftwareUpdateConfigurationByNameOutput) Tasks() SoftwareUpdateConfigurationTasksResponsePtrOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) SoftwareUpdateConfigurationTasksResponsePtrOutput {
		return v.Tasks
	}).(SoftwareUpdateConfigurationTasksResponsePtrOutput)
}

// Resource type
func (o SoftwareUpdateConfigurationByNameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// update specific properties for the Software update configuration
func (o SoftwareUpdateConfigurationByNameOutput) UpdateConfiguration() UpdateConfigurationResponseOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) UpdateConfigurationResponseOutput {
		return v.UpdateConfiguration
	}).(UpdateConfigurationResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(SoftwareUpdateConfigurationByNameOutput{})
}
