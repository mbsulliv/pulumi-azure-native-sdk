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

// Tag Contract details.
type Tag struct {
	pulumi.CustomResourceState

	// Tag name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTag registers a new resource with the given unique name, arguments, and options.
func NewTag(ctx *pulumi.Context,
	name string, args *TagArgs, opts ...pulumi.ResourceOption) (*Tag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:Tag"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Tag
	err := ctx.RegisterResource("azure-native:apimanagement/v20230301preview:Tag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTag gets an existing Tag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagState, opts ...pulumi.ResourceOption) (*Tag, error) {
	var resource Tag
	err := ctx.ReadResource("azure-native:apimanagement/v20230301preview:Tag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tag resources.
type tagState struct {
}

type TagState struct {
}

func (TagState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagState)(nil)).Elem()
}

type tagArgs struct {
	// Tag name.
	DisplayName string `pulumi:"displayName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId *string `pulumi:"tagId"`
}

// The set of arguments for constructing a Tag resource.
type TagArgs struct {
	// Tag name.
	DisplayName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringPtrInput
}

func (TagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagArgs)(nil)).Elem()
}

type TagInput interface {
	pulumi.Input

	ToTagOutput() TagOutput
	ToTagOutputWithContext(ctx context.Context) TagOutput
}

func (*Tag) ElementType() reflect.Type {
	return reflect.TypeOf((**Tag)(nil)).Elem()
}

func (i *Tag) ToTagOutput() TagOutput {
	return i.ToTagOutputWithContext(context.Background())
}

func (i *Tag) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOutput)
}

func (i *Tag) ToOutput(ctx context.Context) pulumix.Output[*Tag] {
	return pulumix.Output[*Tag]{
		OutputState: i.ToTagOutputWithContext(ctx).OutputState,
	}
}

type TagOutput struct{ *pulumi.OutputState }

func (TagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tag)(nil)).Elem()
}

func (o TagOutput) ToTagOutput() TagOutput {
	return o
}

func (o TagOutput) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return o
}

func (o TagOutput) ToOutput(ctx context.Context) pulumix.Output[*Tag] {
	return pulumix.Output[*Tag]{
		OutputState: o.OutputState,
	}
}

// Tag name.
func (o TagOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Tag) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The name of the resource
func (o TagOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Tag) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o TagOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Tag) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TagOutput{})
}
