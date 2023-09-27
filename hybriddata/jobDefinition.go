// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybriddata

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Job Definition.
// Azure REST API version: 2019-06-01. Prior API version in Azure Native 1.x: 2019-06-01
type JobDefinition struct {
	pulumi.CustomResourceState

	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
	CustomerSecrets CustomerSecretResponseArrayOutput `pulumi:"customerSecrets"`
	// A generic json used differently by each data service type.
	DataServiceInput pulumi.AnyOutput `pulumi:"dataServiceInput"`
	// Data Sink Id associated to the job definition.
	DataSinkId pulumi.StringOutput `pulumi:"dataSinkId"`
	// Data Source Id associated to the job definition.
	DataSourceId pulumi.StringOutput `pulumi:"dataSourceId"`
	// Last modified time of the job definition.
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	// Name of the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// This is the preferred geo location for the job to run.
	RunLocation pulumi.StringPtrOutput `pulumi:"runLocation"`
	// Schedule for running the job definition
	Schedules ScheduleResponseArrayOutput `pulumi:"schedules"`
	// State of the job definition.
	State pulumi.StringOutput `pulumi:"state"`
	// Type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
	// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
	UserConfirmation pulumi.StringPtrOutput `pulumi:"userConfirmation"`
}

// NewJobDefinition registers a new resource with the given unique name, arguments, and options.
func NewJobDefinition(ctx *pulumi.Context,
	name string, args *JobDefinitionArgs, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataManagerName == nil {
		return nil, errors.New("invalid value for required argument 'DataManagerName'")
	}
	if args.DataServiceName == nil {
		return nil, errors.New("invalid value for required argument 'DataServiceName'")
	}
	if args.DataSinkId == nil {
		return nil, errors.New("invalid value for required argument 'DataSinkId'")
	}
	if args.DataSourceId == nil {
		return nil, errors.New("invalid value for required argument 'DataSourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	if args.UserConfirmation == nil {
		args.UserConfirmation = UserConfirmation("NotRequired")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybriddata/v20160601:JobDefinition"),
		},
		{
			Type: pulumi.String("azure-native:hybriddata/v20190601:JobDefinition"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource JobDefinition
	err := ctx.RegisterResource("azure-native:hybriddata:JobDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobDefinition gets an existing JobDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobDefinitionState, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	var resource JobDefinition
	err := ctx.ReadResource("azure-native:hybriddata:JobDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobDefinition resources.
type jobDefinitionState struct {
}

type JobDefinitionState struct {
}

func (JobDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionState)(nil)).Elem()
}

type jobDefinitionArgs struct {
	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
	CustomerSecrets []CustomerSecret `pulumi:"customerSecrets"`
	// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	DataManagerName string `pulumi:"dataManagerName"`
	// A generic json used differently by each data service type.
	DataServiceInput interface{} `pulumi:"dataServiceInput"`
	// The data service type of the job definition.
	DataServiceName string `pulumi:"dataServiceName"`
	// Data Sink Id associated to the job definition.
	DataSinkId string `pulumi:"dataSinkId"`
	// Data Source Id associated to the job definition.
	DataSourceId string `pulumi:"dataSourceId"`
	// The job definition name to be created or updated.
	JobDefinitionName *string `pulumi:"jobDefinitionName"`
	// Last modified time of the job definition.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// This is the preferred geo location for the job to run.
	RunLocation *RunLocation `pulumi:"runLocation"`
	// Schedule for running the job definition
	Schedules []Schedule `pulumi:"schedules"`
	// State of the job definition.
	State State `pulumi:"state"`
	// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
	UserConfirmation *UserConfirmation `pulumi:"userConfirmation"`
}

// The set of arguments for constructing a JobDefinition resource.
type JobDefinitionArgs struct {
	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
	CustomerSecrets CustomerSecretArrayInput
	// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	DataManagerName pulumi.StringInput
	// A generic json used differently by each data service type.
	DataServiceInput pulumi.Input
	// The data service type of the job definition.
	DataServiceName pulumi.StringInput
	// Data Sink Id associated to the job definition.
	DataSinkId pulumi.StringInput
	// Data Source Id associated to the job definition.
	DataSourceId pulumi.StringInput
	// The job definition name to be created or updated.
	JobDefinitionName pulumi.StringPtrInput
	// Last modified time of the job definition.
	LastModifiedTime pulumi.StringPtrInput
	// The Resource Group Name
	ResourceGroupName pulumi.StringInput
	// This is the preferred geo location for the job to run.
	RunLocation RunLocationPtrInput
	// Schedule for running the job definition
	Schedules ScheduleArrayInput
	// State of the job definition.
	State StateInput
	// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
	UserConfirmation UserConfirmationPtrInput
}

func (JobDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionArgs)(nil)).Elem()
}

type JobDefinitionInput interface {
	pulumi.Input

	ToJobDefinitionOutput() JobDefinitionOutput
	ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput
}

func (*JobDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinition)(nil)).Elem()
}

func (i *JobDefinition) ToJobDefinitionOutput() JobDefinitionOutput {
	return i.ToJobDefinitionOutputWithContext(context.Background())
}

func (i *JobDefinition) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionOutput)
}

func (i *JobDefinition) ToOutput(ctx context.Context) pulumix.Output[*JobDefinition] {
	return pulumix.Output[*JobDefinition]{
		OutputState: i.ToJobDefinitionOutputWithContext(ctx).OutputState,
	}
}

type JobDefinitionOutput struct{ *pulumi.OutputState }

func (JobDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinition)(nil)).Elem()
}

func (o JobDefinitionOutput) ToJobDefinitionOutput() JobDefinitionOutput {
	return o
}

func (o JobDefinitionOutput) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return o
}

func (o JobDefinitionOutput) ToOutput(ctx context.Context) pulumix.Output[*JobDefinition] {
	return pulumix.Output[*JobDefinition]{
		OutputState: o.OutputState,
	}
}

// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
func (o JobDefinitionOutput) CustomerSecrets() CustomerSecretResponseArrayOutput {
	return o.ApplyT(func(v *JobDefinition) CustomerSecretResponseArrayOutput { return v.CustomerSecrets }).(CustomerSecretResponseArrayOutput)
}

// A generic json used differently by each data service type.
func (o JobDefinitionOutput) DataServiceInput() pulumi.AnyOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.AnyOutput { return v.DataServiceInput }).(pulumi.AnyOutput)
}

// Data Sink Id associated to the job definition.
func (o JobDefinitionOutput) DataSinkId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.DataSinkId }).(pulumi.StringOutput)
}

// Data Source Id associated to the job definition.
func (o JobDefinitionOutput) DataSourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.DataSourceId }).(pulumi.StringOutput)
}

// Last modified time of the job definition.
func (o JobDefinitionOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// Name of the object.
func (o JobDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// This is the preferred geo location for the job to run.
func (o JobDefinitionOutput) RunLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.RunLocation }).(pulumi.StringPtrOutput)
}

// Schedule for running the job definition
func (o JobDefinitionOutput) Schedules() ScheduleResponseArrayOutput {
	return o.ApplyT(func(v *JobDefinition) ScheduleResponseArrayOutput { return v.Schedules }).(ScheduleResponseArrayOutput)
}

// State of the job definition.
func (o JobDefinitionOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Type of the object.
func (o JobDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
func (o JobDefinitionOutput) UserConfirmation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.UserConfirmation }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(JobDefinitionOutput{})
}
