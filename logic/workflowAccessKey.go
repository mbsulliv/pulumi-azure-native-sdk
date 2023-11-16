// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure REST API version: 2015-02-01-preview. Prior API version in Azure Native 1.x: 2015-02-01-preview.
type WorkflowAccessKey struct {
	pulumi.CustomResourceState

	// Gets the workflow access key name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the not-after time.
	NotAfter pulumi.StringPtrOutput `pulumi:"notAfter"`
	// Gets or sets the not-before time.
	NotBefore pulumi.StringPtrOutput `pulumi:"notBefore"`
	// Gets the workflow access key type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkflowAccessKey registers a new resource with the given unique name, arguments, and options.
func NewWorkflowAccessKey(ctx *pulumi.Context,
	name string, args *WorkflowAccessKeyArgs, opts ...pulumi.ResourceOption) (*WorkflowAccessKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkflowName == nil {
		return nil, errors.New("invalid value for required argument 'WorkflowName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic/v20150201preview:WorkflowAccessKey"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkflowAccessKey
	err := ctx.RegisterResource("azure-native:logic:WorkflowAccessKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowAccessKey gets an existing WorkflowAccessKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowAccessKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowAccessKeyState, opts ...pulumi.ResourceOption) (*WorkflowAccessKey, error) {
	var resource WorkflowAccessKey
	err := ctx.ReadResource("azure-native:logic:WorkflowAccessKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowAccessKey resources.
type workflowAccessKeyState struct {
}

type WorkflowAccessKeyState struct {
}

func (WorkflowAccessKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowAccessKeyState)(nil)).Elem()
}

type workflowAccessKeyArgs struct {
	// The workflow access key name.
	AccessKeyName *string `pulumi:"accessKeyName"`
	// Gets or sets the resource id.
	Id *string `pulumi:"id"`
	// Gets or sets the not-after time.
	NotAfter *string `pulumi:"notAfter"`
	// Gets or sets the not-before time.
	NotBefore *string `pulumi:"notBefore"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workflow name.
	WorkflowName string `pulumi:"workflowName"`
}

// The set of arguments for constructing a WorkflowAccessKey resource.
type WorkflowAccessKeyArgs struct {
	// The workflow access key name.
	AccessKeyName pulumi.StringPtrInput
	// Gets or sets the resource id.
	Id pulumi.StringPtrInput
	// Gets or sets the not-after time.
	NotAfter pulumi.StringPtrInput
	// Gets or sets the not-before time.
	NotBefore pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The workflow name.
	WorkflowName pulumi.StringInput
}

func (WorkflowAccessKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowAccessKeyArgs)(nil)).Elem()
}

type WorkflowAccessKeyInput interface {
	pulumi.Input

	ToWorkflowAccessKeyOutput() WorkflowAccessKeyOutput
	ToWorkflowAccessKeyOutputWithContext(ctx context.Context) WorkflowAccessKeyOutput
}

func (*WorkflowAccessKey) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowAccessKey)(nil)).Elem()
}

func (i *WorkflowAccessKey) ToWorkflowAccessKeyOutput() WorkflowAccessKeyOutput {
	return i.ToWorkflowAccessKeyOutputWithContext(context.Background())
}

func (i *WorkflowAccessKey) ToWorkflowAccessKeyOutputWithContext(ctx context.Context) WorkflowAccessKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowAccessKeyOutput)
}

type WorkflowAccessKeyOutput struct{ *pulumi.OutputState }

func (WorkflowAccessKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowAccessKey)(nil)).Elem()
}

func (o WorkflowAccessKeyOutput) ToWorkflowAccessKeyOutput() WorkflowAccessKeyOutput {
	return o
}

func (o WorkflowAccessKeyOutput) ToWorkflowAccessKeyOutputWithContext(ctx context.Context) WorkflowAccessKeyOutput {
	return o
}

// Gets the workflow access key name.
func (o WorkflowAccessKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowAccessKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the not-after time.
func (o WorkflowAccessKeyOutput) NotAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowAccessKey) pulumi.StringPtrOutput { return v.NotAfter }).(pulumi.StringPtrOutput)
}

// Gets or sets the not-before time.
func (o WorkflowAccessKeyOutput) NotBefore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowAccessKey) pulumi.StringPtrOutput { return v.NotBefore }).(pulumi.StringPtrOutput)
}

// Gets the workflow access key type.
func (o WorkflowAccessKeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowAccessKey) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkflowAccessKeyOutput{})
}
