// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents an incident comment
type IncidentComment struct {
	pulumi.CustomResourceState

	// Describes the client that created the comment
	Author ClientInfoResponseOutput `pulumi:"author"`
	// The time the comment was created
	CreatedTimeUtc pulumi.StringOutput `pulumi:"createdTimeUtc"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The time the comment was updated
	LastModifiedTimeUtc pulumi.StringOutput `pulumi:"lastModifiedTimeUtc"`
	// The comment message
	Message pulumi.StringOutput `pulumi:"message"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIncidentComment registers a new resource with the given unique name, arguments, and options.
func NewIncidentComment(ctx *pulumi.Context,
	name string, args *IncidentCommentArgs, opts ...pulumi.ResourceOption) (*IncidentComment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IncidentId == nil {
		return nil, errors.New("invalid value for required argument 'IncidentId'")
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
			Type: pulumi.String("azure-native:securityinsights:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210401:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:IncidentComment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:IncidentComment"),
		},
	})
	opts = append(opts, aliases)
	var resource IncidentComment
	err := ctx.RegisterResource("azure-native:securityinsights/v20210901preview:IncidentComment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIncidentComment gets an existing IncidentComment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIncidentComment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IncidentCommentState, opts ...pulumi.ResourceOption) (*IncidentComment, error) {
	var resource IncidentComment
	err := ctx.ReadResource("azure-native:securityinsights/v20210901preview:IncidentComment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IncidentComment resources.
type incidentCommentState struct {
}

type IncidentCommentState struct {
}

func (IncidentCommentState) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentCommentState)(nil)).Elem()
}

type incidentCommentArgs struct {
	// Incident comment ID
	IncidentCommentId *string `pulumi:"incidentCommentId"`
	// Incident ID
	IncidentId string `pulumi:"incidentId"`
	// The comment message
	Message string `pulumi:"message"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a IncidentComment resource.
type IncidentCommentArgs struct {
	// Incident comment ID
	IncidentCommentId pulumi.StringPtrInput
	// Incident ID
	IncidentId pulumi.StringInput
	// The comment message
	Message pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (IncidentCommentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentCommentArgs)(nil)).Elem()
}

type IncidentCommentInput interface {
	pulumi.Input

	ToIncidentCommentOutput() IncidentCommentOutput
	ToIncidentCommentOutputWithContext(ctx context.Context) IncidentCommentOutput
}

func (*IncidentComment) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentComment)(nil)).Elem()
}

func (i *IncidentComment) ToIncidentCommentOutput() IncidentCommentOutput {
	return i.ToIncidentCommentOutputWithContext(context.Background())
}

func (i *IncidentComment) ToIncidentCommentOutputWithContext(ctx context.Context) IncidentCommentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentCommentOutput)
}

type IncidentCommentOutput struct{ *pulumi.OutputState }

func (IncidentCommentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentComment)(nil)).Elem()
}

func (o IncidentCommentOutput) ToIncidentCommentOutput() IncidentCommentOutput {
	return o
}

func (o IncidentCommentOutput) ToIncidentCommentOutputWithContext(ctx context.Context) IncidentCommentOutput {
	return o
}

// Describes the client that created the comment
func (o IncidentCommentOutput) Author() ClientInfoResponseOutput {
	return o.ApplyT(func(v *IncidentComment) ClientInfoResponseOutput { return v.Author }).(ClientInfoResponseOutput)
}

// The time the comment was created
func (o IncidentCommentOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentComment) pulumi.StringOutput { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o IncidentCommentOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentComment) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The time the comment was updated
func (o IncidentCommentOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentComment) pulumi.StringOutput { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

// The comment message
func (o IncidentCommentOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentComment) pulumi.StringOutput { return v.Message }).(pulumi.StringOutput)
}

// The name of the resource
func (o IncidentCommentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentComment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o IncidentCommentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IncidentComment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o IncidentCommentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentComment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IncidentCommentOutput{})
}
