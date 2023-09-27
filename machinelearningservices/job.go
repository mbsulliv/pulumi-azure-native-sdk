// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
// Azure REST API version: 2023-04-01. Prior API version in Azure Native 1.x: 2021-03-01-preview
type Job struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	JobBaseProperties pulumi.AnyOutput `pulumi:"jobBaseProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobBaseProperties == nil {
		return nil, errors.New("invalid value for required argument 'JobBaseProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:Job"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:Job"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:Job"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:Job"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Job
	err := ctx.RegisterResource("azure-native:machinelearningservices:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("azure-native:machinelearningservices:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
}

type JobState struct {
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// The name and identifier for the Job. This is case-sensitive.
	Id *string `pulumi:"id"`
	// [Required] Additional attributes of the entity.
	JobBaseProperties interface{} `pulumi:"jobBaseProperties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// The name and identifier for the Job. This is case-sensitive.
	Id pulumi.StringPtrInput
	// [Required] Additional attributes of the entity.
	JobBaseProperties pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

func (i *Job) ToOutput(ctx context.Context) pulumix.Output[*Job] {
	return pulumix.Output[*Job]{
		OutputState: i.ToJobOutputWithContext(ctx).OutputState,
	}
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

func (o JobOutput) ToOutput(ctx context.Context) pulumix.Output[*Job] {
	return pulumix.Output[*Job]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o JobOutput) JobBaseProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Job) pulumi.AnyOutput { return v.JobBaseProperties }).(pulumi.AnyOutput)
}

// The name of the resource
func (o JobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o JobOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Job) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o JobOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobOutput{})
}
