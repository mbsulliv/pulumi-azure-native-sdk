// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Issue Attachment Contract details.
// Azure REST API version: 2022-08-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2022-09-01-preview, 2023-03-01-preview.
type ApiIssueAttachment struct {
	pulumi.CustomResourceState

	// An HTTP link or Base64-encoded binary data.
	Content pulumi.StringOutput `pulumi:"content"`
	// Either 'link' if content is provided via an HTTP link or the MIME type of the Base64-encoded binary data provided in the 'content' property.
	ContentFormat pulumi.StringOutput `pulumi:"contentFormat"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Filename by which the binary data will be saved.
	Title pulumi.StringOutput `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApiIssueAttachment registers a new resource with the given unique name, arguments, and options.
func NewApiIssueAttachment(ctx *pulumi.Context,
	name string, args *ApiIssueAttachmentArgs, opts ...pulumi.ResourceOption) (*ApiIssueAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.ContentFormat == nil {
		return nil, errors.New("invalid value for required argument 'ContentFormat'")
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
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:ApiIssueAttachment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApiIssueAttachment
	err := ctx.RegisterResource("azure-native:apimanagement:ApiIssueAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiIssueAttachment gets an existing ApiIssueAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiIssueAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiIssueAttachmentState, opts ...pulumi.ResourceOption) (*ApiIssueAttachment, error) {
	var resource ApiIssueAttachment
	err := ctx.ReadResource("azure-native:apimanagement:ApiIssueAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiIssueAttachment resources.
type apiIssueAttachmentState struct {
}

type ApiIssueAttachmentState struct {
}

func (ApiIssueAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIssueAttachmentState)(nil)).Elem()
}

type apiIssueAttachmentArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId string `pulumi:"apiId"`
	// Attachment identifier within an Issue. Must be unique in the current Issue.
	AttachmentId *string `pulumi:"attachmentId"`
	// An HTTP link or Base64-encoded binary data.
	Content string `pulumi:"content"`
	// Either 'link' if content is provided via an HTTP link or the MIME type of the Base64-encoded binary data provided in the 'content' property.
	ContentFormat string `pulumi:"contentFormat"`
	// Issue identifier. Must be unique in the current API Management service instance.
	IssueId string `pulumi:"issueId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Filename by which the binary data will be saved.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a ApiIssueAttachment resource.
type ApiIssueAttachmentArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId pulumi.StringInput
	// Attachment identifier within an Issue. Must be unique in the current Issue.
	AttachmentId pulumi.StringPtrInput
	// An HTTP link or Base64-encoded binary data.
	Content pulumi.StringInput
	// Either 'link' if content is provided via an HTTP link or the MIME type of the Base64-encoded binary data provided in the 'content' property.
	ContentFormat pulumi.StringInput
	// Issue identifier. Must be unique in the current API Management service instance.
	IssueId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Filename by which the binary data will be saved.
	Title pulumi.StringInput
}

func (ApiIssueAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIssueAttachmentArgs)(nil)).Elem()
}

type ApiIssueAttachmentInput interface {
	pulumi.Input

	ToApiIssueAttachmentOutput() ApiIssueAttachmentOutput
	ToApiIssueAttachmentOutputWithContext(ctx context.Context) ApiIssueAttachmentOutput
}

func (*ApiIssueAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIssueAttachment)(nil)).Elem()
}

func (i *ApiIssueAttachment) ToApiIssueAttachmentOutput() ApiIssueAttachmentOutput {
	return i.ToApiIssueAttachmentOutputWithContext(context.Background())
}

func (i *ApiIssueAttachment) ToApiIssueAttachmentOutputWithContext(ctx context.Context) ApiIssueAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIssueAttachmentOutput)
}

type ApiIssueAttachmentOutput struct{ *pulumi.OutputState }

func (ApiIssueAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIssueAttachment)(nil)).Elem()
}

func (o ApiIssueAttachmentOutput) ToApiIssueAttachmentOutput() ApiIssueAttachmentOutput {
	return o
}

func (o ApiIssueAttachmentOutput) ToApiIssueAttachmentOutputWithContext(ctx context.Context) ApiIssueAttachmentOutput {
	return o
}

// An HTTP link or Base64-encoded binary data.
func (o ApiIssueAttachmentOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssueAttachment) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// Either 'link' if content is provided via an HTTP link or the MIME type of the Base64-encoded binary data provided in the 'content' property.
func (o ApiIssueAttachmentOutput) ContentFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssueAttachment) pulumi.StringOutput { return v.ContentFormat }).(pulumi.StringOutput)
}

// The name of the resource
func (o ApiIssueAttachmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssueAttachment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Filename by which the binary data will be saved.
func (o ApiIssueAttachmentOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssueAttachment) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ApiIssueAttachmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssueAttachment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiIssueAttachmentOutput{})
}
