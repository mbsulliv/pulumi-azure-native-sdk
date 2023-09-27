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

// Tag-operation link details.
type TagOperationLink struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Full resource Id of an API operation.
	OperationId pulumi.StringOutput `pulumi:"operationId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTagOperationLink registers a new resource with the given unique name, arguments, and options.
func NewTagOperationLink(ctx *pulumi.Context,
	name string, args *TagOperationLinkArgs, opts ...pulumi.ResourceOption) (*TagOperationLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OperationId == nil {
		return nil, errors.New("invalid value for required argument 'OperationId'")
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:TagOperationLink"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:TagOperationLink"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TagOperationLink
	err := ctx.RegisterResource("azure-native:apimanagement/v20230301preview:TagOperationLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagOperationLink gets an existing TagOperationLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagOperationLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagOperationLinkState, opts ...pulumi.ResourceOption) (*TagOperationLink, error) {
	var resource TagOperationLink
	err := ctx.ReadResource("azure-native:apimanagement/v20230301preview:TagOperationLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagOperationLink resources.
type tagOperationLinkState struct {
}

type TagOperationLinkState struct {
}

func (TagOperationLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagOperationLinkState)(nil)).Elem()
}

type tagOperationLinkArgs struct {
	// Full resource Id of an API operation.
	OperationId string `pulumi:"operationId"`
	// Tag-operation link identifier. Must be unique in the current API Management service instance.
	OperationLinkId *string `pulumi:"operationLinkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
}

// The set of arguments for constructing a TagOperationLink resource.
type TagOperationLinkArgs struct {
	// Full resource Id of an API operation.
	OperationId pulumi.StringInput
	// Tag-operation link identifier. Must be unique in the current API Management service instance.
	OperationLinkId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringInput
}

func (TagOperationLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagOperationLinkArgs)(nil)).Elem()
}

type TagOperationLinkInput interface {
	pulumi.Input

	ToTagOperationLinkOutput() TagOperationLinkOutput
	ToTagOperationLinkOutputWithContext(ctx context.Context) TagOperationLinkOutput
}

func (*TagOperationLink) ElementType() reflect.Type {
	return reflect.TypeOf((**TagOperationLink)(nil)).Elem()
}

func (i *TagOperationLink) ToTagOperationLinkOutput() TagOperationLinkOutput {
	return i.ToTagOperationLinkOutputWithContext(context.Background())
}

func (i *TagOperationLink) ToTagOperationLinkOutputWithContext(ctx context.Context) TagOperationLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOperationLinkOutput)
}

func (i *TagOperationLink) ToOutput(ctx context.Context) pulumix.Output[*TagOperationLink] {
	return pulumix.Output[*TagOperationLink]{
		OutputState: i.ToTagOperationLinkOutputWithContext(ctx).OutputState,
	}
}

type TagOperationLinkOutput struct{ *pulumi.OutputState }

func (TagOperationLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagOperationLink)(nil)).Elem()
}

func (o TagOperationLinkOutput) ToTagOperationLinkOutput() TagOperationLinkOutput {
	return o
}

func (o TagOperationLinkOutput) ToTagOperationLinkOutputWithContext(ctx context.Context) TagOperationLinkOutput {
	return o
}

func (o TagOperationLinkOutput) ToOutput(ctx context.Context) pulumix.Output[*TagOperationLink] {
	return pulumix.Output[*TagOperationLink]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o TagOperationLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagOperationLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Full resource Id of an API operation.
func (o TagOperationLinkOutput) OperationId() pulumi.StringOutput {
	return o.ApplyT(func(v *TagOperationLink) pulumi.StringOutput { return v.OperationId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o TagOperationLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TagOperationLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TagOperationLinkOutput{})
}
