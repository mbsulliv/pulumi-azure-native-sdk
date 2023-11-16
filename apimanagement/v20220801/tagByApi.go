// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Tag Contract details.
type TagByApi struct {
	pulumi.CustomResourceState

	// Tag name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTagByApi registers a new resource with the given unique name, arguments, and options.
func NewTagByApi(ctx *pulumi.Context,
	name string, args *TagByApiArgs, opts ...pulumi.ResourceOption) (*TagByApi, error) {
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:TagByApi"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TagByApi
	err := ctx.RegisterResource("azure-native:apimanagement/v20220801:TagByApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagByApi gets an existing TagByApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagByApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagByApiState, opts ...pulumi.ResourceOption) (*TagByApi, error) {
	var resource TagByApi
	err := ctx.ReadResource("azure-native:apimanagement/v20220801:TagByApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagByApi resources.
type tagByApiState struct {
}

type TagByApiState struct {
}

func (TagByApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByApiState)(nil)).Elem()
}

type tagByApiArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId *string `pulumi:"tagId"`
}

// The set of arguments for constructing a TagByApi resource.
type TagByApiArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringPtrInput
}

func (TagByApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByApiArgs)(nil)).Elem()
}

type TagByApiInput interface {
	pulumi.Input

	ToTagByApiOutput() TagByApiOutput
	ToTagByApiOutputWithContext(ctx context.Context) TagByApiOutput
}

func (*TagByApi) ElementType() reflect.Type {
	return reflect.TypeOf((**TagByApi)(nil)).Elem()
}

func (i *TagByApi) ToTagByApiOutput() TagByApiOutput {
	return i.ToTagByApiOutputWithContext(context.Background())
}

func (i *TagByApi) ToTagByApiOutputWithContext(ctx context.Context) TagByApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagByApiOutput)
}

type TagByApiOutput struct{ *pulumi.OutputState }

func (TagByApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagByApi)(nil)).Elem()
}

func (o TagByApiOutput) ToTagByApiOutput() TagByApiOutput {
	return o
}

func (o TagByApiOutput) ToTagByApiOutputWithContext(ctx context.Context) TagByApiOutput {
	return o
}

// Tag name.
func (o TagByApiOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *TagByApi) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The name of the resource
func (o TagByApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagByApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o TagByApiOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TagByApi) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TagByApiOutput{})
}
