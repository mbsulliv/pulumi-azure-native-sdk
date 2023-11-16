// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A group of job targets.
// Azure REST API version: 2021-11-01. Prior API version in Azure Native 1.x: 2020-11-01-preview.
//
// Other available API versions: 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview.
type JobTargetGroup struct {
	pulumi.CustomResourceState

	// Members of the target group.
	Members JobTargetResponseArrayOutput `pulumi:"members"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJobTargetGroup registers a new resource with the given unique name, arguments, and options.
func NewJobTargetGroup(ctx *pulumi.Context,
	name string, args *JobTargetGroupArgs, opts ...pulumi.ResourceOption) (*JobTargetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobAgentName == nil {
		return nil, errors.New("invalid value for required argument 'JobAgentName'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:JobTargetGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource JobTargetGroup
	err := ctx.RegisterResource("azure-native:sql:JobTargetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobTargetGroup gets an existing JobTargetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobTargetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobTargetGroupState, opts ...pulumi.ResourceOption) (*JobTargetGroup, error) {
	var resource JobTargetGroup
	err := ctx.ReadResource("azure-native:sql:JobTargetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobTargetGroup resources.
type jobTargetGroupState struct {
}

type JobTargetGroupState struct {
}

func (JobTargetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobTargetGroupState)(nil)).Elem()
}

type jobTargetGroupArgs struct {
	// The name of the job agent.
	JobAgentName string `pulumi:"jobAgentName"`
	// Members of the target group.
	Members []JobTarget `pulumi:"members"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the target group.
	TargetGroupName *string `pulumi:"targetGroupName"`
}

// The set of arguments for constructing a JobTargetGroup resource.
type JobTargetGroupArgs struct {
	// The name of the job agent.
	JobAgentName pulumi.StringInput
	// Members of the target group.
	Members JobTargetArrayInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The name of the target group.
	TargetGroupName pulumi.StringPtrInput
}

func (JobTargetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobTargetGroupArgs)(nil)).Elem()
}

type JobTargetGroupInput interface {
	pulumi.Input

	ToJobTargetGroupOutput() JobTargetGroupOutput
	ToJobTargetGroupOutputWithContext(ctx context.Context) JobTargetGroupOutput
}

func (*JobTargetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**JobTargetGroup)(nil)).Elem()
}

func (i *JobTargetGroup) ToJobTargetGroupOutput() JobTargetGroupOutput {
	return i.ToJobTargetGroupOutputWithContext(context.Background())
}

func (i *JobTargetGroup) ToJobTargetGroupOutputWithContext(ctx context.Context) JobTargetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTargetGroupOutput)
}

type JobTargetGroupOutput struct{ *pulumi.OutputState }

func (JobTargetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobTargetGroup)(nil)).Elem()
}

func (o JobTargetGroupOutput) ToJobTargetGroupOutput() JobTargetGroupOutput {
	return o
}

func (o JobTargetGroupOutput) ToJobTargetGroupOutputWithContext(ctx context.Context) JobTargetGroupOutput {
	return o
}

// Members of the target group.
func (o JobTargetGroupOutput) Members() JobTargetResponseArrayOutput {
	return o.ApplyT(func(v *JobTargetGroup) JobTargetResponseArrayOutput { return v.Members }).(JobTargetResponseArrayOutput)
}

// Resource name.
func (o JobTargetGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTargetGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o JobTargetGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTargetGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobTargetGroupOutput{})
}
