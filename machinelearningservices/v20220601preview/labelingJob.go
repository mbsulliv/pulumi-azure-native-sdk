// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
type LabelingJob struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	LabelingJobProperties LabelingJobResponseOutput `pulumi:"labelingJobProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLabelingJob registers a new resource with the given unique name, arguments, and options.
func NewLabelingJob(ctx *pulumi.Context,
	name string, args *LabelingJobArgs, opts ...pulumi.ResourceOption) (*LabelingJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabelingJobProperties == nil {
		return nil, errors.New("invalid value for required argument 'LabelingJobProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.LabelingJobProperties = args.LabelingJobProperties.ToLabelingJobTypeOutput().ApplyT(func(v LabelingJobType) LabelingJobType { return *v.Defaults() }).(LabelingJobTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:LabelingJob"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:LabelingJob"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:LabelingJob"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:LabelingJob"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:LabelingJob"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:LabelingJob"),
		},
	})
	opts = append(opts, aliases)
	var resource LabelingJob
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220601preview:LabelingJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLabelingJob gets an existing LabelingJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLabelingJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabelingJobState, opts ...pulumi.ResourceOption) (*LabelingJob, error) {
	var resource LabelingJob
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220601preview:LabelingJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LabelingJob resources.
type labelingJobState struct {
}

type LabelingJobState struct {
}

func (LabelingJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*labelingJobState)(nil)).Elem()
}

type labelingJobArgs struct {
	// The name and identifier for the LabelingJob.
	Id *string `pulumi:"id"`
	// [Required] Additional attributes of the entity.
	LabelingJobProperties LabelingJobType `pulumi:"labelingJobProperties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a LabelingJob resource.
type LabelingJobArgs struct {
	// The name and identifier for the LabelingJob.
	Id pulumi.StringPtrInput
	// [Required] Additional attributes of the entity.
	LabelingJobProperties LabelingJobTypeInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (LabelingJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labelingJobArgs)(nil)).Elem()
}

type LabelingJobInput interface {
	pulumi.Input

	ToLabelingJobOutput() LabelingJobOutput
	ToLabelingJobOutputWithContext(ctx context.Context) LabelingJobOutput
}

func (*LabelingJob) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelingJob)(nil)).Elem()
}

func (i *LabelingJob) ToLabelingJobOutput() LabelingJobOutput {
	return i.ToLabelingJobOutputWithContext(context.Background())
}

func (i *LabelingJob) ToLabelingJobOutputWithContext(ctx context.Context) LabelingJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingJobOutput)
}

type LabelingJobOutput struct{ *pulumi.OutputState }

func (LabelingJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelingJob)(nil)).Elem()
}

func (o LabelingJobOutput) ToLabelingJobOutput() LabelingJobOutput {
	return o
}

func (o LabelingJobOutput) ToLabelingJobOutputWithContext(ctx context.Context) LabelingJobOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o LabelingJobOutput) LabelingJobProperties() LabelingJobResponseOutput {
	return o.ApplyT(func(v *LabelingJob) LabelingJobResponseOutput { return v.LabelingJobProperties }).(LabelingJobResponseOutput)
}

// The name of the resource
func (o LabelingJobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LabelingJob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LabelingJobOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LabelingJob) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LabelingJobOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LabelingJob) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LabelingJobOutput{})
}
