// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Security Assignment on a resource group over a given scope
type Assignment struct {
	pulumi.CustomResourceState

	// Additional data about the assignment
	AdditionalData AssignmentPropertiesResponseAdditionalDataPtrOutput `pulumi:"additionalData"`
	// Component item with key as applied to this standard assignment over the given scope
	AssignedComponent AssignedComponentItemResponsePtrOutput `pulumi:"assignedComponent"`
	// Standard item with key as applied to this standard assignment over the given scope
	AssignedStandard AssignedStandardItemResponsePtrOutput `pulumi:"assignedStandard"`
	// description of the standardAssignment
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// display name of the standardAssignment
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// expected effect of this assignment (Disable/Exempt/etc)
	Effect pulumi.StringPtrOutput `pulumi:"effect"`
	// Entity tag is used for comparing two or more entities from the same requested resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Expiration date of this assignment as a full ISO date
	ExpiresOn pulumi.StringPtrOutput `pulumi:"expiresOn"`
	// Kind of the resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Location where the resource is stored
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The assignment metadata. Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Scope to which the standardAssignment applies - can be a subscription path or a resource group under that subscription
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// A list of key value pairs that describe the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAssignment registers a new resource with the given unique name, arguments, and options.
func NewAssignment(ctx *pulumi.Context,
	name string, args *AssignmentArgs, opts ...pulumi.ResourceOption) (*Assignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security:Assignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Assignment
	err := ctx.RegisterResource("azure-native:security/v20210801preview:Assignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssignment gets an existing Assignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssignmentState, opts ...pulumi.ResourceOption) (*Assignment, error) {
	var resource Assignment
	err := ctx.ReadResource("azure-native:security/v20210801preview:Assignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Assignment resources.
type assignmentState struct {
}

type AssignmentState struct {
}

func (AssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentState)(nil)).Elem()
}

type assignmentArgs struct {
	// Additional data about the assignment
	AdditionalData *AssignmentPropertiesAdditionalData `pulumi:"additionalData"`
	// Component item with key as applied to this standard assignment over the given scope
	AssignedComponent *AssignedComponentItem `pulumi:"assignedComponent"`
	// Standard item with key as applied to this standard assignment over the given scope
	AssignedStandard *AssignedStandardItem `pulumi:"assignedStandard"`
	// The security assignment key - unique key for the standard assignment
	AssignmentId *string `pulumi:"assignmentId"`
	// description of the standardAssignment
	Description *string `pulumi:"description"`
	// display name of the standardAssignment
	DisplayName *string `pulumi:"displayName"`
	// expected effect of this assignment (Disable/Exempt/etc)
	Effect *string `pulumi:"effect"`
	// Expiration date of this assignment as a full ISO date
	ExpiresOn *string `pulumi:"expiresOn"`
	// Kind of the resource
	Kind *string `pulumi:"kind"`
	// Location where the resource is stored
	Location *string `pulumi:"location"`
	// The assignment metadata. Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata interface{} `pulumi:"metadata"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Scope to which the standardAssignment applies - can be a subscription path or a resource group under that subscription
	Scope *string `pulumi:"scope"`
	// A list of key value pairs that describe the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Assignment resource.
type AssignmentArgs struct {
	// Additional data about the assignment
	AdditionalData AssignmentPropertiesAdditionalDataPtrInput
	// Component item with key as applied to this standard assignment over the given scope
	AssignedComponent AssignedComponentItemPtrInput
	// Standard item with key as applied to this standard assignment over the given scope
	AssignedStandard AssignedStandardItemPtrInput
	// The security assignment key - unique key for the standard assignment
	AssignmentId pulumi.StringPtrInput
	// description of the standardAssignment
	Description pulumi.StringPtrInput
	// display name of the standardAssignment
	DisplayName pulumi.StringPtrInput
	// expected effect of this assignment (Disable/Exempt/etc)
	Effect pulumi.StringPtrInput
	// Expiration date of this assignment as a full ISO date
	ExpiresOn pulumi.StringPtrInput
	// Kind of the resource
	Kind pulumi.StringPtrInput
	// Location where the resource is stored
	Location pulumi.StringPtrInput
	// The assignment metadata. Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata pulumi.Input
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Scope to which the standardAssignment applies - can be a subscription path or a resource group under that subscription
	Scope pulumi.StringPtrInput
	// A list of key value pairs that describe the resource.
	Tags pulumi.StringMapInput
}

func (AssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentArgs)(nil)).Elem()
}

type AssignmentInput interface {
	pulumi.Input

	ToAssignmentOutput() AssignmentOutput
	ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput
}

func (*Assignment) ElementType() reflect.Type {
	return reflect.TypeOf((**Assignment)(nil)).Elem()
}

func (i *Assignment) ToAssignmentOutput() AssignmentOutput {
	return i.ToAssignmentOutputWithContext(context.Background())
}

func (i *Assignment) ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentOutput)
}

func (i *Assignment) ToOutput(ctx context.Context) pulumix.Output[*Assignment] {
	return pulumix.Output[*Assignment]{
		OutputState: i.ToAssignmentOutputWithContext(ctx).OutputState,
	}
}

type AssignmentOutput struct{ *pulumi.OutputState }

func (AssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Assignment)(nil)).Elem()
}

func (o AssignmentOutput) ToAssignmentOutput() AssignmentOutput {
	return o
}

func (o AssignmentOutput) ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput {
	return o
}

func (o AssignmentOutput) ToOutput(ctx context.Context) pulumix.Output[*Assignment] {
	return pulumix.Output[*Assignment]{
		OutputState: o.OutputState,
	}
}

// Additional data about the assignment
func (o AssignmentOutput) AdditionalData() AssignmentPropertiesResponseAdditionalDataPtrOutput {
	return o.ApplyT(func(v *Assignment) AssignmentPropertiesResponseAdditionalDataPtrOutput { return v.AdditionalData }).(AssignmentPropertiesResponseAdditionalDataPtrOutput)
}

// Component item with key as applied to this standard assignment over the given scope
func (o AssignmentOutput) AssignedComponent() AssignedComponentItemResponsePtrOutput {
	return o.ApplyT(func(v *Assignment) AssignedComponentItemResponsePtrOutput { return v.AssignedComponent }).(AssignedComponentItemResponsePtrOutput)
}

// Standard item with key as applied to this standard assignment over the given scope
func (o AssignmentOutput) AssignedStandard() AssignedStandardItemResponsePtrOutput {
	return o.ApplyT(func(v *Assignment) AssignedStandardItemResponsePtrOutput { return v.AssignedStandard }).(AssignedStandardItemResponsePtrOutput)
}

// description of the standardAssignment
func (o AssignmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// display name of the standardAssignment
func (o AssignmentOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// expected effect of this assignment (Disable/Exempt/etc)
func (o AssignmentOutput) Effect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.Effect }).(pulumi.StringPtrOutput)
}

// Entity tag is used for comparing two or more entities from the same requested resource.
func (o AssignmentOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Expiration date of this assignment as a full ISO date
func (o AssignmentOutput) ExpiresOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.ExpiresOn }).(pulumi.StringPtrOutput)
}

// Kind of the resource
func (o AssignmentOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Location where the resource is stored
func (o AssignmentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The assignment metadata. Metadata is an open ended object and is typically a collection of key value pairs.
func (o AssignmentOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *Assignment) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

// Resource name
func (o AssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Scope to which the standardAssignment applies - can be a subscription path or a resource group under that subscription
func (o AssignmentOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Assignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// A list of key value pairs that describe the resource.
func (o AssignmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o AssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssignmentOutput{})
}
