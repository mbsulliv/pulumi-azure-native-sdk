// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200515preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Linked workspace.
type LinkedWorkspace struct {
	pulumi.CustomResourceState

	// Friendly name of the linked workspace.
	Name pulumi.StringOutput `pulumi:"name"`
	// LinkedWorkspace specific properties.
	Properties LinkedWorkspacePropsResponseOutput `pulumi:"properties"`
	// Resource type of linked workspace.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLinkedWorkspace registers a new resource with the given unique name, arguments, and options.
func NewLinkedWorkspace(ctx *pulumi.Context,
	name string, args *LinkedWorkspaceArgs, opts ...pulumi.ResourceOption) (*LinkedWorkspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:LinkedWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:LinkedWorkspace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LinkedWorkspace
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20200515preview:LinkedWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedWorkspace gets an existing LinkedWorkspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedWorkspaceState, opts ...pulumi.ResourceOption) (*LinkedWorkspace, error) {
	var resource LinkedWorkspace
	err := ctx.ReadResource("azure-native:machinelearningservices/v20200515preview:LinkedWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedWorkspace resources.
type linkedWorkspaceState struct {
}

type LinkedWorkspaceState struct {
}

func (LinkedWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedWorkspaceState)(nil)).Elem()
}

type linkedWorkspaceArgs struct {
	// Friendly name of the linked workspace
	LinkName *string `pulumi:"linkName"`
	// Friendly name of the linked workspace
	Name *string `pulumi:"name"`
	// LinkedWorkspace specific properties.
	Properties *LinkedWorkspaceProps `pulumi:"properties"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a LinkedWorkspace resource.
type LinkedWorkspaceArgs struct {
	// Friendly name of the linked workspace
	LinkName pulumi.StringPtrInput
	// Friendly name of the linked workspace
	Name pulumi.StringPtrInput
	// LinkedWorkspace specific properties.
	Properties LinkedWorkspacePropsPtrInput
	// Name of the resource group in which workspace is located.
	ResourceGroupName pulumi.StringInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (LinkedWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedWorkspaceArgs)(nil)).Elem()
}

type LinkedWorkspaceInput interface {
	pulumi.Input

	ToLinkedWorkspaceOutput() LinkedWorkspaceOutput
	ToLinkedWorkspaceOutputWithContext(ctx context.Context) LinkedWorkspaceOutput
}

func (*LinkedWorkspace) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedWorkspace)(nil)).Elem()
}

func (i *LinkedWorkspace) ToLinkedWorkspaceOutput() LinkedWorkspaceOutput {
	return i.ToLinkedWorkspaceOutputWithContext(context.Background())
}

func (i *LinkedWorkspace) ToLinkedWorkspaceOutputWithContext(ctx context.Context) LinkedWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedWorkspaceOutput)
}

func (i *LinkedWorkspace) ToOutput(ctx context.Context) pulumix.Output[*LinkedWorkspace] {
	return pulumix.Output[*LinkedWorkspace]{
		OutputState: i.ToLinkedWorkspaceOutputWithContext(ctx).OutputState,
	}
}

type LinkedWorkspaceOutput struct{ *pulumi.OutputState }

func (LinkedWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedWorkspace)(nil)).Elem()
}

func (o LinkedWorkspaceOutput) ToLinkedWorkspaceOutput() LinkedWorkspaceOutput {
	return o
}

func (o LinkedWorkspaceOutput) ToLinkedWorkspaceOutputWithContext(ctx context.Context) LinkedWorkspaceOutput {
	return o
}

func (o LinkedWorkspaceOutput) ToOutput(ctx context.Context) pulumix.Output[*LinkedWorkspace] {
	return pulumix.Output[*LinkedWorkspace]{
		OutputState: o.OutputState,
	}
}

// Friendly name of the linked workspace.
func (o LinkedWorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedWorkspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// LinkedWorkspace specific properties.
func (o LinkedWorkspaceOutput) Properties() LinkedWorkspacePropsResponseOutput {
	return o.ApplyT(func(v *LinkedWorkspace) LinkedWorkspacePropsResponseOutput { return v.Properties }).(LinkedWorkspacePropsResponseOutput)
}

// Resource type of linked workspace.
func (o LinkedWorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedWorkspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedWorkspaceOutput{})
}
