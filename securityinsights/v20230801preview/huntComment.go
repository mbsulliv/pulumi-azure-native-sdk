// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Hunt Comment in Azure Security Insights
type HuntComment struct {
	pulumi.CustomResourceState

	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The message for the comment
	Message pulumi.StringOutput `pulumi:"message"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHuntComment registers a new resource with the given unique name, arguments, and options.
func NewHuntComment(ctx *pulumi.Context,
	name string, args *HuntCommentArgs, opts ...pulumi.ResourceOption) (*HuntComment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HuntId == nil {
		return nil, errors.New("invalid value for required argument 'HuntId'")
	}
	if args.Message == nil {
		return nil, errors.New("invalid value for required argument 'Message'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:HuntComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:HuntComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:HuntComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:HuntComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:HuntComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:HuntComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:HuntComment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HuntComment
	err := ctx.RegisterResource("azure-native:securityinsights/v20230801preview:HuntComment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHuntComment gets an existing HuntComment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHuntComment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HuntCommentState, opts ...pulumi.ResourceOption) (*HuntComment, error) {
	var resource HuntComment
	err := ctx.ReadResource("azure-native:securityinsights/v20230801preview:HuntComment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HuntComment resources.
type huntCommentState struct {
}

type HuntCommentState struct {
}

func (HuntCommentState) ElementType() reflect.Type {
	return reflect.TypeOf((*huntCommentState)(nil)).Elem()
}

type huntCommentArgs struct {
	// The hunt comment id (GUID)
	HuntCommentId *string `pulumi:"huntCommentId"`
	// The hunt id (GUID)
	HuntId string `pulumi:"huntId"`
	// The message for the comment
	Message string `pulumi:"message"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a HuntComment resource.
type HuntCommentArgs struct {
	// The hunt comment id (GUID)
	HuntCommentId pulumi.StringPtrInput
	// The hunt id (GUID)
	HuntId pulumi.StringInput
	// The message for the comment
	Message pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (HuntCommentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*huntCommentArgs)(nil)).Elem()
}

type HuntCommentInput interface {
	pulumi.Input

	ToHuntCommentOutput() HuntCommentOutput
	ToHuntCommentOutputWithContext(ctx context.Context) HuntCommentOutput
}

func (*HuntComment) ElementType() reflect.Type {
	return reflect.TypeOf((**HuntComment)(nil)).Elem()
}

func (i *HuntComment) ToHuntCommentOutput() HuntCommentOutput {
	return i.ToHuntCommentOutputWithContext(context.Background())
}

func (i *HuntComment) ToHuntCommentOutputWithContext(ctx context.Context) HuntCommentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HuntCommentOutput)
}

type HuntCommentOutput struct{ *pulumi.OutputState }

func (HuntCommentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HuntComment)(nil)).Elem()
}

func (o HuntCommentOutput) ToHuntCommentOutput() HuntCommentOutput {
	return o
}

func (o HuntCommentOutput) ToHuntCommentOutputWithContext(ctx context.Context) HuntCommentOutput {
	return o
}

// Etag of the azure resource
func (o HuntCommentOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HuntComment) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The message for the comment
func (o HuntCommentOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v *HuntComment) pulumi.StringOutput { return v.Message }).(pulumi.StringOutput)
}

// The name of the resource
func (o HuntCommentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HuntComment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o HuntCommentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HuntComment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o HuntCommentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HuntComment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HuntCommentOutput{})
}
