// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Product-group link details.
type WorkspaceProductGroupLink struct {
	pulumi.CustomResourceState

	// Full resource Id of a group.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceProductGroupLink registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceProductGroupLink(ctx *pulumi.Context,
	name string, args *WorkspaceProductGroupLinkArgs, opts ...pulumi.ResourceOption) (*WorkspaceProductGroupLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:WorkspaceProductGroupLink"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:WorkspaceProductGroupLink"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceProductGroupLink
	err := ctx.RegisterResource("azure-native:apimanagement/v20220901preview:WorkspaceProductGroupLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceProductGroupLink gets an existing WorkspaceProductGroupLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceProductGroupLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceProductGroupLinkState, opts ...pulumi.ResourceOption) (*WorkspaceProductGroupLink, error) {
	var resource WorkspaceProductGroupLink
	err := ctx.ReadResource("azure-native:apimanagement/v20220901preview:WorkspaceProductGroupLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceProductGroupLink resources.
type workspaceProductGroupLinkState struct {
}

type WorkspaceProductGroupLinkState struct {
}

func (WorkspaceProductGroupLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceProductGroupLinkState)(nil)).Elem()
}

type workspaceProductGroupLinkArgs struct {
	// Full resource Id of a group.
	GroupId string `pulumi:"groupId"`
	// Product-Group link identifier. Must be unique in the current API Management service instance.
	GroupLinkId *string `pulumi:"groupLinkId"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a WorkspaceProductGroupLink resource.
type WorkspaceProductGroupLinkArgs struct {
	// Full resource Id of a group.
	GroupId pulumi.StringInput
	// Product-Group link identifier. Must be unique in the current API Management service instance.
	GroupLinkId pulumi.StringPtrInput
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput
}

func (WorkspaceProductGroupLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceProductGroupLinkArgs)(nil)).Elem()
}

type WorkspaceProductGroupLinkInput interface {
	pulumi.Input

	ToWorkspaceProductGroupLinkOutput() WorkspaceProductGroupLinkOutput
	ToWorkspaceProductGroupLinkOutputWithContext(ctx context.Context) WorkspaceProductGroupLinkOutput
}

func (*WorkspaceProductGroupLink) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceProductGroupLink)(nil)).Elem()
}

func (i *WorkspaceProductGroupLink) ToWorkspaceProductGroupLinkOutput() WorkspaceProductGroupLinkOutput {
	return i.ToWorkspaceProductGroupLinkOutputWithContext(context.Background())
}

func (i *WorkspaceProductGroupLink) ToWorkspaceProductGroupLinkOutputWithContext(ctx context.Context) WorkspaceProductGroupLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceProductGroupLinkOutput)
}

type WorkspaceProductGroupLinkOutput struct{ *pulumi.OutputState }

func (WorkspaceProductGroupLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceProductGroupLink)(nil)).Elem()
}

func (o WorkspaceProductGroupLinkOutput) ToWorkspaceProductGroupLinkOutput() WorkspaceProductGroupLinkOutput {
	return o
}

func (o WorkspaceProductGroupLinkOutput) ToWorkspaceProductGroupLinkOutputWithContext(ctx context.Context) WorkspaceProductGroupLinkOutput {
	return o
}

// Full resource Id of a group.
func (o WorkspaceProductGroupLinkOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceProductGroupLink) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The name of the resource
func (o WorkspaceProductGroupLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceProductGroupLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceProductGroupLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceProductGroupLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceProductGroupLinkOutput{})
}