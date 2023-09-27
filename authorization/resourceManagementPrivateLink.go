// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure REST API version: 2020-05-01. Prior API version in Azure Native 1.x: 2020-05-01
type ResourceManagementPrivateLink struct {
	pulumi.CustomResourceState

	// the region of the rmpl
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The rmpl Name.
	Name       pulumi.StringOutput                                            `pulumi:"name"`
	Properties ResourceManagementPrivateLinkEndpointConnectionsResponseOutput `pulumi:"properties"`
	// The operation type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewResourceManagementPrivateLink registers a new resource with the given unique name, arguments, and options.
func NewResourceManagementPrivateLink(ctx *pulumi.Context,
	name string, args *ResourceManagementPrivateLinkArgs, opts ...pulumi.ResourceOption) (*ResourceManagementPrivateLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization/v20200501:ResourceManagementPrivateLink"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ResourceManagementPrivateLink
	err := ctx.RegisterResource("azure-native:authorization:ResourceManagementPrivateLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceManagementPrivateLink gets an existing ResourceManagementPrivateLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceManagementPrivateLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceManagementPrivateLinkState, opts ...pulumi.ResourceOption) (*ResourceManagementPrivateLink, error) {
	var resource ResourceManagementPrivateLink
	err := ctx.ReadResource("azure-native:authorization:ResourceManagementPrivateLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceManagementPrivateLink resources.
type resourceManagementPrivateLinkState struct {
}

type ResourceManagementPrivateLinkState struct {
}

func (ResourceManagementPrivateLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceManagementPrivateLinkState)(nil)).Elem()
}

type resourceManagementPrivateLinkArgs struct {
	// the region to create private link association.
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource management private link.
	RmplName *string `pulumi:"rmplName"`
}

// The set of arguments for constructing a ResourceManagementPrivateLink resource.
type ResourceManagementPrivateLinkArgs struct {
	// the region to create private link association.
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the resource management private link.
	RmplName pulumi.StringPtrInput
}

func (ResourceManagementPrivateLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceManagementPrivateLinkArgs)(nil)).Elem()
}

type ResourceManagementPrivateLinkInput interface {
	pulumi.Input

	ToResourceManagementPrivateLinkOutput() ResourceManagementPrivateLinkOutput
	ToResourceManagementPrivateLinkOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkOutput
}

func (*ResourceManagementPrivateLink) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceManagementPrivateLink)(nil)).Elem()
}

func (i *ResourceManagementPrivateLink) ToResourceManagementPrivateLinkOutput() ResourceManagementPrivateLinkOutput {
	return i.ToResourceManagementPrivateLinkOutputWithContext(context.Background())
}

func (i *ResourceManagementPrivateLink) ToResourceManagementPrivateLinkOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceManagementPrivateLinkOutput)
}

func (i *ResourceManagementPrivateLink) ToOutput(ctx context.Context) pulumix.Output[*ResourceManagementPrivateLink] {
	return pulumix.Output[*ResourceManagementPrivateLink]{
		OutputState: i.ToResourceManagementPrivateLinkOutputWithContext(ctx).OutputState,
	}
}

type ResourceManagementPrivateLinkOutput struct{ *pulumi.OutputState }

func (ResourceManagementPrivateLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceManagementPrivateLink)(nil)).Elem()
}

func (o ResourceManagementPrivateLinkOutput) ToResourceManagementPrivateLinkOutput() ResourceManagementPrivateLinkOutput {
	return o
}

func (o ResourceManagementPrivateLinkOutput) ToResourceManagementPrivateLinkOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkOutput {
	return o
}

func (o ResourceManagementPrivateLinkOutput) ToOutput(ctx context.Context) pulumix.Output[*ResourceManagementPrivateLink] {
	return pulumix.Output[*ResourceManagementPrivateLink]{
		OutputState: o.OutputState,
	}
}

// the region of the rmpl
func (o ResourceManagementPrivateLinkOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceManagementPrivateLink) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The rmpl Name.
func (o ResourceManagementPrivateLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceManagementPrivateLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceManagementPrivateLinkOutput) Properties() ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
	return o.ApplyT(func(v *ResourceManagementPrivateLink) ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
		return v.Properties
	}).(ResourceManagementPrivateLinkEndpointConnectionsResponseOutput)
}

// The operation type.
func (o ResourceManagementPrivateLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceManagementPrivateLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceManagementPrivateLinkOutput{})
}
