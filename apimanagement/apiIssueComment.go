// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Issue Comment Contract details.
// Azure REST API version: 2022-08-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2022-09-01-preview, 2023-03-01-preview.
type ApiIssueComment struct {
	pulumi.CustomResourceState

	// Date and time when the comment was created.
	CreatedDate pulumi.StringPtrOutput `pulumi:"createdDate"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Comment text.
	Text pulumi.StringOutput `pulumi:"text"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// A resource identifier for the user who left the comment.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewApiIssueComment registers a new resource with the given unique name, arguments, and options.
func NewApiIssueComment(ctx *pulumi.Context,
	name string, args *ApiIssueCommentArgs, opts ...pulumi.ResourceOption) (*ApiIssueComment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.IssueId == nil {
		return nil, errors.New("invalid value for required argument 'IssueId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Text == nil {
		return nil, errors.New("invalid value for required argument 'Text'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:ApiIssueComment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApiIssueComment
	err := ctx.RegisterResource("azure-native:apimanagement:ApiIssueComment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiIssueComment gets an existing ApiIssueComment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiIssueComment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiIssueCommentState, opts ...pulumi.ResourceOption) (*ApiIssueComment, error) {
	var resource ApiIssueComment
	err := ctx.ReadResource("azure-native:apimanagement:ApiIssueComment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiIssueComment resources.
type apiIssueCommentState struct {
}

type ApiIssueCommentState struct {
}

func (ApiIssueCommentState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIssueCommentState)(nil)).Elem()
}

type apiIssueCommentArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId string `pulumi:"apiId"`
	// Comment identifier within an Issue. Must be unique in the current Issue.
	CommentId *string `pulumi:"commentId"`
	// Date and time when the comment was created.
	CreatedDate *string `pulumi:"createdDate"`
	// Issue identifier. Must be unique in the current API Management service instance.
	IssueId string `pulumi:"issueId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Comment text.
	Text string `pulumi:"text"`
	// A resource identifier for the user who left the comment.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a ApiIssueComment resource.
type ApiIssueCommentArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId pulumi.StringInput
	// Comment identifier within an Issue. Must be unique in the current Issue.
	CommentId pulumi.StringPtrInput
	// Date and time when the comment was created.
	CreatedDate pulumi.StringPtrInput
	// Issue identifier. Must be unique in the current API Management service instance.
	IssueId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Comment text.
	Text pulumi.StringInput
	// A resource identifier for the user who left the comment.
	UserId pulumi.StringInput
}

func (ApiIssueCommentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIssueCommentArgs)(nil)).Elem()
}

type ApiIssueCommentInput interface {
	pulumi.Input

	ToApiIssueCommentOutput() ApiIssueCommentOutput
	ToApiIssueCommentOutputWithContext(ctx context.Context) ApiIssueCommentOutput
}

func (*ApiIssueComment) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIssueComment)(nil)).Elem()
}

func (i *ApiIssueComment) ToApiIssueCommentOutput() ApiIssueCommentOutput {
	return i.ToApiIssueCommentOutputWithContext(context.Background())
}

func (i *ApiIssueComment) ToApiIssueCommentOutputWithContext(ctx context.Context) ApiIssueCommentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIssueCommentOutput)
}

func (i *ApiIssueComment) ToOutput(ctx context.Context) pulumix.Output[*ApiIssueComment] {
	return pulumix.Output[*ApiIssueComment]{
		OutputState: i.ToApiIssueCommentOutputWithContext(ctx).OutputState,
	}
}

type ApiIssueCommentOutput struct{ *pulumi.OutputState }

func (ApiIssueCommentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIssueComment)(nil)).Elem()
}

func (o ApiIssueCommentOutput) ToApiIssueCommentOutput() ApiIssueCommentOutput {
	return o
}

func (o ApiIssueCommentOutput) ToApiIssueCommentOutputWithContext(ctx context.Context) ApiIssueCommentOutput {
	return o
}

func (o ApiIssueCommentOutput) ToOutput(ctx context.Context) pulumix.Output[*ApiIssueComment] {
	return pulumix.Output[*ApiIssueComment]{
		OutputState: o.OutputState,
	}
}

// Date and time when the comment was created.
func (o ApiIssueCommentOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiIssueComment) pulumi.StringPtrOutput { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ApiIssueCommentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssueComment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Comment text.
func (o ApiIssueCommentOutput) Text() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssueComment) pulumi.StringOutput { return v.Text }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ApiIssueCommentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssueComment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A resource identifier for the user who left the comment.
func (o ApiIssueCommentOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssueComment) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiIssueCommentOutput{})
}
