// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Policy fragment contract details.
type WorkspacePolicyFragment struct {
	pulumi.CustomResourceState

	// Policy fragment description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Format of the policy fragment content.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Contents of the policy fragment.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewWorkspacePolicyFragment registers a new resource with the given unique name, arguments, and options.
func NewWorkspacePolicyFragment(ctx *pulumi.Context,
	name string, args *WorkspacePolicyFragmentArgs, opts ...pulumi.ResourceOption) (*WorkspacePolicyFragment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	if args.Format == nil {
		args.Format = pulumi.StringPtr("xml")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:WorkspacePolicyFragment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:WorkspacePolicyFragment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspacePolicyFragment
	err := ctx.RegisterResource("azure-native:apimanagement/v20220901preview:WorkspacePolicyFragment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspacePolicyFragment gets an existing WorkspacePolicyFragment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspacePolicyFragment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspacePolicyFragmentState, opts ...pulumi.ResourceOption) (*WorkspacePolicyFragment, error) {
	var resource WorkspacePolicyFragment
	err := ctx.ReadResource("azure-native:apimanagement/v20220901preview:WorkspacePolicyFragment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspacePolicyFragment resources.
type workspacePolicyFragmentState struct {
}

type WorkspacePolicyFragmentState struct {
}

func (WorkspacePolicyFragmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspacePolicyFragmentState)(nil)).Elem()
}

type workspacePolicyFragmentArgs struct {
	// Policy fragment description.
	Description *string `pulumi:"description"`
	// Format of the policy fragment content.
	Format *string `pulumi:"format"`
	// A resource identifier.
	Id *string `pulumi:"id"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Contents of the policy fragment.
	Value string `pulumi:"value"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a WorkspacePolicyFragment resource.
type WorkspacePolicyFragmentArgs struct {
	// Policy fragment description.
	Description pulumi.StringPtrInput
	// Format of the policy fragment content.
	Format pulumi.StringPtrInput
	// A resource identifier.
	Id pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Contents of the policy fragment.
	Value pulumi.StringInput
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput
}

func (WorkspacePolicyFragmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspacePolicyFragmentArgs)(nil)).Elem()
}

type WorkspacePolicyFragmentInput interface {
	pulumi.Input

	ToWorkspacePolicyFragmentOutput() WorkspacePolicyFragmentOutput
	ToWorkspacePolicyFragmentOutputWithContext(ctx context.Context) WorkspacePolicyFragmentOutput
}

func (*WorkspacePolicyFragment) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacePolicyFragment)(nil)).Elem()
}

func (i *WorkspacePolicyFragment) ToWorkspacePolicyFragmentOutput() WorkspacePolicyFragmentOutput {
	return i.ToWorkspacePolicyFragmentOutputWithContext(context.Background())
}

func (i *WorkspacePolicyFragment) ToWorkspacePolicyFragmentOutputWithContext(ctx context.Context) WorkspacePolicyFragmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacePolicyFragmentOutput)
}

func (i *WorkspacePolicyFragment) ToOutput(ctx context.Context) pulumix.Output[*WorkspacePolicyFragment] {
	return pulumix.Output[*WorkspacePolicyFragment]{
		OutputState: i.ToWorkspacePolicyFragmentOutputWithContext(ctx).OutputState,
	}
}

type WorkspacePolicyFragmentOutput struct{ *pulumi.OutputState }

func (WorkspacePolicyFragmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacePolicyFragment)(nil)).Elem()
}

func (o WorkspacePolicyFragmentOutput) ToWorkspacePolicyFragmentOutput() WorkspacePolicyFragmentOutput {
	return o
}

func (o WorkspacePolicyFragmentOutput) ToWorkspacePolicyFragmentOutputWithContext(ctx context.Context) WorkspacePolicyFragmentOutput {
	return o
}

func (o WorkspacePolicyFragmentOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkspacePolicyFragment] {
	return pulumix.Output[*WorkspacePolicyFragment]{
		OutputState: o.OutputState,
	}
}

// Policy fragment description.
func (o WorkspacePolicyFragmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspacePolicyFragment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Format of the policy fragment content.
func (o WorkspacePolicyFragmentOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspacePolicyFragment) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o WorkspacePolicyFragmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacePolicyFragment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspacePolicyFragmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacePolicyFragment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Contents of the policy fragment.
func (o WorkspacePolicyFragmentOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacePolicyFragment) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspacePolicyFragmentOutput{})
}
