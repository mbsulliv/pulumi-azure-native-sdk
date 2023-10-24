// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type WorkspaceConnection struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name       pulumi.StringOutput `pulumi:"name"`
	Properties pulumi.AnyOutput    `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceConnection registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceConnection(ctx *pulumi.Context,
	name string, args *WorkspaceConnectionArgs, opts ...pulumi.ResourceOption) (*WorkspaceConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200601:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200801:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210101:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210701:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220101preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:WorkspaceConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceConnection
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230801preview:WorkspaceConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceConnection gets an existing WorkspaceConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceConnectionState, opts ...pulumi.ResourceOption) (*WorkspaceConnection, error) {
	var resource WorkspaceConnection
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230801preview:WorkspaceConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceConnection resources.
type workspaceConnectionState struct {
}

type WorkspaceConnectionState struct {
}

func (WorkspaceConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceConnectionState)(nil)).Elem()
}

type workspaceConnectionArgs struct {
	// Friendly name of the workspace connection
	ConnectionName *string     `pulumi:"connectionName"`
	Properties     interface{} `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure Machine Learning Workspace Name
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a WorkspaceConnection resource.
type WorkspaceConnectionArgs struct {
	// Friendly name of the workspace connection
	ConnectionName pulumi.StringPtrInput
	Properties     pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Azure Machine Learning Workspace Name
	WorkspaceName pulumi.StringInput
}

func (WorkspaceConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceConnectionArgs)(nil)).Elem()
}

type WorkspaceConnectionInput interface {
	pulumi.Input

	ToWorkspaceConnectionOutput() WorkspaceConnectionOutput
	ToWorkspaceConnectionOutputWithContext(ctx context.Context) WorkspaceConnectionOutput
}

func (*WorkspaceConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceConnection)(nil)).Elem()
}

func (i *WorkspaceConnection) ToWorkspaceConnectionOutput() WorkspaceConnectionOutput {
	return i.ToWorkspaceConnectionOutputWithContext(context.Background())
}

func (i *WorkspaceConnection) ToWorkspaceConnectionOutputWithContext(ctx context.Context) WorkspaceConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceConnectionOutput)
}

func (i *WorkspaceConnection) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceConnection] {
	return pulumix.Output[*WorkspaceConnection]{
		OutputState: i.ToWorkspaceConnectionOutputWithContext(ctx).OutputState,
	}
}

type WorkspaceConnectionOutput struct{ *pulumi.OutputState }

func (WorkspaceConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceConnection)(nil)).Elem()
}

func (o WorkspaceConnectionOutput) ToWorkspaceConnectionOutput() WorkspaceConnectionOutput {
	return o
}

func (o WorkspaceConnectionOutput) ToWorkspaceConnectionOutputWithContext(ctx context.Context) WorkspaceConnectionOutput {
	return o
}

func (o WorkspaceConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceConnection] {
	return pulumix.Output[*WorkspaceConnection]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o WorkspaceConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkspaceConnectionOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *WorkspaceConnection) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WorkspaceConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WorkspaceConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceConnectionOutput{})
}
