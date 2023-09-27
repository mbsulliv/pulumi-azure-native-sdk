// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Tag-API link details.
type WorkspaceTagApiLink struct {
	pulumi.CustomResourceState

	// Full resource Id of an API.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceTagApiLink registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceTagApiLink(ctx *pulumi.Context,
	name string, args *WorkspaceTagApiLinkArgs, opts ...pulumi.ResourceOption) (*WorkspaceTagApiLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.TagId == nil {
		return nil, errors.New("invalid value for required argument 'TagId'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:WorkspaceTagApiLink"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:WorkspaceTagApiLink"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceTagApiLink
	err := ctx.RegisterResource("azure-native:apimanagement/v20230301preview:WorkspaceTagApiLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceTagApiLink gets an existing WorkspaceTagApiLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceTagApiLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceTagApiLinkState, opts ...pulumi.ResourceOption) (*WorkspaceTagApiLink, error) {
	var resource WorkspaceTagApiLink
	err := ctx.ReadResource("azure-native:apimanagement/v20230301preview:WorkspaceTagApiLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceTagApiLink resources.
type workspaceTagApiLinkState struct {
}

type WorkspaceTagApiLinkState struct {
}

func (WorkspaceTagApiLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceTagApiLinkState)(nil)).Elem()
}

type workspaceTagApiLinkArgs struct {
	// Full resource Id of an API.
	ApiId string `pulumi:"apiId"`
	// Tag-API link identifier. Must be unique in the current API Management service instance.
	ApiLinkId *string `pulumi:"apiLinkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a WorkspaceTagApiLink resource.
type WorkspaceTagApiLinkArgs struct {
	// Full resource Id of an API.
	ApiId pulumi.StringInput
	// Tag-API link identifier. Must be unique in the current API Management service instance.
	ApiLinkId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringInput
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput
}

func (WorkspaceTagApiLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceTagApiLinkArgs)(nil)).Elem()
}

type WorkspaceTagApiLinkInput interface {
	pulumi.Input

	ToWorkspaceTagApiLinkOutput() WorkspaceTagApiLinkOutput
	ToWorkspaceTagApiLinkOutputWithContext(ctx context.Context) WorkspaceTagApiLinkOutput
}

func (*WorkspaceTagApiLink) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceTagApiLink)(nil)).Elem()
}

func (i *WorkspaceTagApiLink) ToWorkspaceTagApiLinkOutput() WorkspaceTagApiLinkOutput {
	return i.ToWorkspaceTagApiLinkOutputWithContext(context.Background())
}

func (i *WorkspaceTagApiLink) ToWorkspaceTagApiLinkOutputWithContext(ctx context.Context) WorkspaceTagApiLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceTagApiLinkOutput)
}

func (i *WorkspaceTagApiLink) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceTagApiLink] {
	return pulumix.Output[*WorkspaceTagApiLink]{
		OutputState: i.ToWorkspaceTagApiLinkOutputWithContext(ctx).OutputState,
	}
}

type WorkspaceTagApiLinkOutput struct{ *pulumi.OutputState }

func (WorkspaceTagApiLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceTagApiLink)(nil)).Elem()
}

func (o WorkspaceTagApiLinkOutput) ToWorkspaceTagApiLinkOutput() WorkspaceTagApiLinkOutput {
	return o
}

func (o WorkspaceTagApiLinkOutput) ToWorkspaceTagApiLinkOutputWithContext(ctx context.Context) WorkspaceTagApiLinkOutput {
	return o
}

func (o WorkspaceTagApiLinkOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceTagApiLink] {
	return pulumix.Output[*WorkspaceTagApiLink]{
		OutputState: o.OutputState,
	}
}

// Full resource Id of an API.
func (o WorkspaceTagApiLinkOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceTagApiLink) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// The name of the resource
func (o WorkspaceTagApiLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceTagApiLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceTagApiLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceTagApiLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceTagApiLinkOutput{})
}
