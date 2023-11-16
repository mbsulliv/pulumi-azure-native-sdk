// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datamigration

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A task resource
// Azure REST API version: 2021-06-30. Prior API version in Azure Native 1.x: 2018-04-19.
//
// Other available API versions: 2022-03-30-preview.
type Task struct {
	pulumi.CustomResourceState

	// HTTP strong entity tag value. This is ignored if submitted.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Custom task properties
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTask registers a new resource with the given unique name, arguments, and options.
func NewTask(ctx *pulumi.Context,
	name string, args *TaskArgs, opts ...pulumi.ResourceOption) (*Task, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datamigration/v20171115preview:Task"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180315preview:Task"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180331preview:Task"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180419:Task"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180715preview:Task"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20210630:Task"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20211030preview:Task"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220130preview:Task"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220330preview:Task"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Task
	err := ctx.RegisterResource("azure-native:datamigration:Task", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTask gets an existing Task resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskState, opts ...pulumi.ResourceOption) (*Task, error) {
	var resource Task
	err := ctx.ReadResource("azure-native:datamigration:Task", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Task resources.
type taskState struct {
}

type TaskState struct {
}

func (TaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*taskState)(nil)).Elem()
}

type taskArgs struct {
	// Name of the resource group
	GroupName string `pulumi:"groupName"`
	// Name of the project
	ProjectName string `pulumi:"projectName"`
	// Custom task properties
	Properties interface{} `pulumi:"properties"`
	// Name of the service
	ServiceName string `pulumi:"serviceName"`
	// Name of the Task
	TaskName *string `pulumi:"taskName"`
}

// The set of arguments for constructing a Task resource.
type TaskArgs struct {
	// Name of the resource group
	GroupName pulumi.StringInput
	// Name of the project
	ProjectName pulumi.StringInput
	// Custom task properties
	Properties pulumi.Input
	// Name of the service
	ServiceName pulumi.StringInput
	// Name of the Task
	TaskName pulumi.StringPtrInput
}

func (TaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taskArgs)(nil)).Elem()
}

type TaskInput interface {
	pulumi.Input

	ToTaskOutput() TaskOutput
	ToTaskOutputWithContext(ctx context.Context) TaskOutput
}

func (*Task) ElementType() reflect.Type {
	return reflect.TypeOf((**Task)(nil)).Elem()
}

func (i *Task) ToTaskOutput() TaskOutput {
	return i.ToTaskOutputWithContext(context.Background())
}

func (i *Task) ToTaskOutputWithContext(ctx context.Context) TaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskOutput)
}

type TaskOutput struct{ *pulumi.OutputState }

func (TaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Task)(nil)).Elem()
}

func (o TaskOutput) ToTaskOutput() TaskOutput {
	return o
}

func (o TaskOutput) ToTaskOutputWithContext(ctx context.Context) TaskOutput {
	return o
}

// HTTP strong entity tag value. This is ignored if submitted.
func (o TaskOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Task) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o TaskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Custom task properties
func (o TaskOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Task) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o TaskOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Task) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o TaskOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TaskOutput{})
}
