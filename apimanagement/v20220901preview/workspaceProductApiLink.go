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

// Product-API link details.
type WorkspaceProductApiLink struct {
	pulumi.CustomResourceState

	// Full resource Id of an API.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceProductApiLink registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceProductApiLink(ctx *pulumi.Context,
	name string, args *WorkspaceProductApiLinkArgs, opts ...pulumi.ResourceOption) (*WorkspaceProductApiLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
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
			Type: pulumi.String("azure-native:apimanagement:WorkspaceProductApiLink"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:WorkspaceProductApiLink"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceProductApiLink
	err := ctx.RegisterResource("azure-native:apimanagement/v20220901preview:WorkspaceProductApiLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceProductApiLink gets an existing WorkspaceProductApiLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceProductApiLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceProductApiLinkState, opts ...pulumi.ResourceOption) (*WorkspaceProductApiLink, error) {
	var resource WorkspaceProductApiLink
	err := ctx.ReadResource("azure-native:apimanagement/v20220901preview:WorkspaceProductApiLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceProductApiLink resources.
type workspaceProductApiLinkState struct {
}

type WorkspaceProductApiLinkState struct {
}

func (WorkspaceProductApiLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceProductApiLinkState)(nil)).Elem()
}

type workspaceProductApiLinkArgs struct {
	// Full resource Id of an API.
	ApiId string `pulumi:"apiId"`
	// Product-API link identifier. Must be unique in the current API Management service instance.
	ApiLinkId *string `pulumi:"apiLinkId"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a WorkspaceProductApiLink resource.
type WorkspaceProductApiLinkArgs struct {
	// Full resource Id of an API.
	ApiId pulumi.StringInput
	// Product-API link identifier. Must be unique in the current API Management service instance.
	ApiLinkId pulumi.StringPtrInput
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput
}

func (WorkspaceProductApiLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceProductApiLinkArgs)(nil)).Elem()
}

type WorkspaceProductApiLinkInput interface {
	pulumi.Input

	ToWorkspaceProductApiLinkOutput() WorkspaceProductApiLinkOutput
	ToWorkspaceProductApiLinkOutputWithContext(ctx context.Context) WorkspaceProductApiLinkOutput
}

func (*WorkspaceProductApiLink) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceProductApiLink)(nil)).Elem()
}

func (i *WorkspaceProductApiLink) ToWorkspaceProductApiLinkOutput() WorkspaceProductApiLinkOutput {
	return i.ToWorkspaceProductApiLinkOutputWithContext(context.Background())
}

func (i *WorkspaceProductApiLink) ToWorkspaceProductApiLinkOutputWithContext(ctx context.Context) WorkspaceProductApiLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceProductApiLinkOutput)
}

func (i *WorkspaceProductApiLink) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceProductApiLink] {
	return pulumix.Output[*WorkspaceProductApiLink]{
		OutputState: i.ToWorkspaceProductApiLinkOutputWithContext(ctx).OutputState,
	}
}

type WorkspaceProductApiLinkOutput struct{ *pulumi.OutputState }

func (WorkspaceProductApiLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceProductApiLink)(nil)).Elem()
}

func (o WorkspaceProductApiLinkOutput) ToWorkspaceProductApiLinkOutput() WorkspaceProductApiLinkOutput {
	return o
}

func (o WorkspaceProductApiLinkOutput) ToWorkspaceProductApiLinkOutputWithContext(ctx context.Context) WorkspaceProductApiLinkOutput {
	return o
}

func (o WorkspaceProductApiLinkOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceProductApiLink] {
	return pulumix.Output[*WorkspaceProductApiLink]{
		OutputState: o.OutputState,
	}
}

// Full resource Id of an API.
func (o WorkspaceProductApiLinkOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceProductApiLink) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// The name of the resource
func (o WorkspaceProductApiLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceProductApiLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceProductApiLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceProductApiLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceProductApiLinkOutput{})
}
