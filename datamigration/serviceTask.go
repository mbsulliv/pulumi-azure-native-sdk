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
// Azure REST API version: 2021-06-30.
//
// Other available API versions: 2022-03-30-preview.
type ServiceTask struct {
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

// NewServiceTask registers a new resource with the given unique name, arguments, and options.
func NewServiceTask(ctx *pulumi.Context,
	name string, args *ServiceTaskArgs, opts ...pulumi.ResourceOption) (*ServiceTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datamigration/v20180715preview:ServiceTask"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20210630:ServiceTask"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20211030preview:ServiceTask"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220130preview:ServiceTask"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220330preview:ServiceTask"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServiceTask
	err := ctx.RegisterResource("azure-native:datamigration:ServiceTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceTask gets an existing ServiceTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceTaskState, opts ...pulumi.ResourceOption) (*ServiceTask, error) {
	var resource ServiceTask
	err := ctx.ReadResource("azure-native:datamigration:ServiceTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceTask resources.
type serviceTaskState struct {
}

type ServiceTaskState struct {
}

func (ServiceTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTaskState)(nil)).Elem()
}

type serviceTaskArgs struct {
	// Name of the resource group
	GroupName string `pulumi:"groupName"`
	// Custom task properties
	Properties interface{} `pulumi:"properties"`
	// Name of the service
	ServiceName string `pulumi:"serviceName"`
	// Name of the Task
	TaskName *string `pulumi:"taskName"`
}

// The set of arguments for constructing a ServiceTask resource.
type ServiceTaskArgs struct {
	// Name of the resource group
	GroupName pulumi.StringInput
	// Custom task properties
	Properties pulumi.Input
	// Name of the service
	ServiceName pulumi.StringInput
	// Name of the Task
	TaskName pulumi.StringPtrInput
}

func (ServiceTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTaskArgs)(nil)).Elem()
}

type ServiceTaskInput interface {
	pulumi.Input

	ToServiceTaskOutput() ServiceTaskOutput
	ToServiceTaskOutputWithContext(ctx context.Context) ServiceTaskOutput
}

func (*ServiceTask) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTask)(nil)).Elem()
}

func (i *ServiceTask) ToServiceTaskOutput() ServiceTaskOutput {
	return i.ToServiceTaskOutputWithContext(context.Background())
}

func (i *ServiceTask) ToServiceTaskOutputWithContext(ctx context.Context) ServiceTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTaskOutput)
}

type ServiceTaskOutput struct{ *pulumi.OutputState }

func (ServiceTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTask)(nil)).Elem()
}

func (o ServiceTaskOutput) ToServiceTaskOutput() ServiceTaskOutput {
	return o
}

func (o ServiceTaskOutput) ToServiceTaskOutputWithContext(ctx context.Context) ServiceTaskOutput {
	return o
}

// HTTP strong entity tag value. This is ignored if submitted.
func (o ServiceTaskOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceTask) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ServiceTaskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceTask) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Custom task properties
func (o ServiceTaskOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ServiceTask) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ServiceTaskOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ServiceTask) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o ServiceTaskOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceTask) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceTaskOutput{})
}
