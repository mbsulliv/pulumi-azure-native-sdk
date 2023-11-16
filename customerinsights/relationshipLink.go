// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package customerinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The relationship link resource format.
// Azure REST API version: 2017-04-26. Prior API version in Azure Native 1.x: 2017-04-26.
//
// Other available API versions: 2017-01-01.
type RelationshipLink struct {
	pulumi.CustomResourceState

	// Localized descriptions for the Relationship Link.
	Description pulumi.StringMapOutput `pulumi:"description"`
	// Localized display name for the Relationship Link.
	DisplayName pulumi.StringMapOutput `pulumi:"displayName"`
	// The InteractionType associated with the Relationship Link.
	InteractionType pulumi.StringOutput `pulumi:"interactionType"`
	// The name of the Relationship Link.
	LinkName pulumi.StringOutput `pulumi:"linkName"`
	// The mappings between Interaction and Relationship fields.
	Mappings RelationshipLinkFieldMappingResponseArrayOutput `pulumi:"mappings"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The property references for the Profile of the Relationship.
	ProfilePropertyReferences ParticipantProfilePropertyReferenceResponseArrayOutput `pulumi:"profilePropertyReferences"`
	// Provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The property references for the Related Profile of the Relationship.
	RelatedProfilePropertyReferences ParticipantProfilePropertyReferenceResponseArrayOutput `pulumi:"relatedProfilePropertyReferences"`
	// The relationship guid id.
	RelationshipGuidId pulumi.StringOutput `pulumi:"relationshipGuidId"`
	// The Relationship associated with the Link.
	RelationshipName pulumi.StringOutput `pulumi:"relationshipName"`
	// The hub name.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRelationshipLink registers a new resource with the given unique name, arguments, and options.
func NewRelationshipLink(ctx *pulumi.Context,
	name string, args *RelationshipLinkArgs, opts ...pulumi.ResourceOption) (*RelationshipLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.InteractionType == nil {
		return nil, errors.New("invalid value for required argument 'InteractionType'")
	}
	if args.ProfilePropertyReferences == nil {
		return nil, errors.New("invalid value for required argument 'ProfilePropertyReferences'")
	}
	if args.RelatedProfilePropertyReferences == nil {
		return nil, errors.New("invalid value for required argument 'RelatedProfilePropertyReferences'")
	}
	if args.RelationshipName == nil {
		return nil, errors.New("invalid value for required argument 'RelationshipName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customerinsights/v20170101:RelationshipLink"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:RelationshipLink"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RelationshipLink
	err := ctx.RegisterResource("azure-native:customerinsights:RelationshipLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRelationshipLink gets an existing RelationshipLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRelationshipLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RelationshipLinkState, opts ...pulumi.ResourceOption) (*RelationshipLink, error) {
	var resource RelationshipLink
	err := ctx.ReadResource("azure-native:customerinsights:RelationshipLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RelationshipLink resources.
type relationshipLinkState struct {
}

type RelationshipLinkState struct {
}

func (RelationshipLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*relationshipLinkState)(nil)).Elem()
}

type relationshipLinkArgs struct {
	// Localized descriptions for the Relationship Link.
	Description map[string]string `pulumi:"description"`
	// Localized display name for the Relationship Link.
	DisplayName map[string]string `pulumi:"displayName"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The InteractionType associated with the Relationship Link.
	InteractionType string `pulumi:"interactionType"`
	// The mappings between Interaction and Relationship fields.
	Mappings []RelationshipLinkFieldMapping `pulumi:"mappings"`
	// The property references for the Profile of the Relationship.
	ProfilePropertyReferences []ParticipantProfilePropertyReference `pulumi:"profilePropertyReferences"`
	// The property references for the Related Profile of the Relationship.
	RelatedProfilePropertyReferences []ParticipantProfilePropertyReference `pulumi:"relatedProfilePropertyReferences"`
	// The name of the relationship link.
	RelationshipLinkName *string `pulumi:"relationshipLinkName"`
	// The Relationship associated with the Link.
	RelationshipName string `pulumi:"relationshipName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a RelationshipLink resource.
type RelationshipLinkArgs struct {
	// Localized descriptions for the Relationship Link.
	Description pulumi.StringMapInput
	// Localized display name for the Relationship Link.
	DisplayName pulumi.StringMapInput
	// The name of the hub.
	HubName pulumi.StringInput
	// The InteractionType associated with the Relationship Link.
	InteractionType pulumi.StringInput
	// The mappings between Interaction and Relationship fields.
	Mappings RelationshipLinkFieldMappingArrayInput
	// The property references for the Profile of the Relationship.
	ProfilePropertyReferences ParticipantProfilePropertyReferenceArrayInput
	// The property references for the Related Profile of the Relationship.
	RelatedProfilePropertyReferences ParticipantProfilePropertyReferenceArrayInput
	// The name of the relationship link.
	RelationshipLinkName pulumi.StringPtrInput
	// The Relationship associated with the Link.
	RelationshipName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (RelationshipLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*relationshipLinkArgs)(nil)).Elem()
}

type RelationshipLinkInput interface {
	pulumi.Input

	ToRelationshipLinkOutput() RelationshipLinkOutput
	ToRelationshipLinkOutputWithContext(ctx context.Context) RelationshipLinkOutput
}

func (*RelationshipLink) ElementType() reflect.Type {
	return reflect.TypeOf((**RelationshipLink)(nil)).Elem()
}

func (i *RelationshipLink) ToRelationshipLinkOutput() RelationshipLinkOutput {
	return i.ToRelationshipLinkOutputWithContext(context.Background())
}

func (i *RelationshipLink) ToRelationshipLinkOutputWithContext(ctx context.Context) RelationshipLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipLinkOutput)
}

type RelationshipLinkOutput struct{ *pulumi.OutputState }

func (RelationshipLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RelationshipLink)(nil)).Elem()
}

func (o RelationshipLinkOutput) ToRelationshipLinkOutput() RelationshipLinkOutput {
	return o
}

func (o RelationshipLinkOutput) ToRelationshipLinkOutputWithContext(ctx context.Context) RelationshipLinkOutput {
	return o
}

// Localized descriptions for the Relationship Link.
func (o RelationshipLinkOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RelationshipLink) pulumi.StringMapOutput { return v.Description }).(pulumi.StringMapOutput)
}

// Localized display name for the Relationship Link.
func (o RelationshipLinkOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RelationshipLink) pulumi.StringMapOutput { return v.DisplayName }).(pulumi.StringMapOutput)
}

// The InteractionType associated with the Relationship Link.
func (o RelationshipLinkOutput) InteractionType() pulumi.StringOutput {
	return o.ApplyT(func(v *RelationshipLink) pulumi.StringOutput { return v.InteractionType }).(pulumi.StringOutput)
}

// The name of the Relationship Link.
func (o RelationshipLinkOutput) LinkName() pulumi.StringOutput {
	return o.ApplyT(func(v *RelationshipLink) pulumi.StringOutput { return v.LinkName }).(pulumi.StringOutput)
}

// The mappings between Interaction and Relationship fields.
func (o RelationshipLinkOutput) Mappings() RelationshipLinkFieldMappingResponseArrayOutput {
	return o.ApplyT(func(v *RelationshipLink) RelationshipLinkFieldMappingResponseArrayOutput { return v.Mappings }).(RelationshipLinkFieldMappingResponseArrayOutput)
}

// Resource name.
func (o RelationshipLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RelationshipLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The property references for the Profile of the Relationship.
func (o RelationshipLinkOutput) ProfilePropertyReferences() ParticipantProfilePropertyReferenceResponseArrayOutput {
	return o.ApplyT(func(v *RelationshipLink) ParticipantProfilePropertyReferenceResponseArrayOutput {
		return v.ProfilePropertyReferences
	}).(ParticipantProfilePropertyReferenceResponseArrayOutput)
}

// Provisioning state.
func (o RelationshipLinkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RelationshipLink) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The property references for the Related Profile of the Relationship.
func (o RelationshipLinkOutput) RelatedProfilePropertyReferences() ParticipantProfilePropertyReferenceResponseArrayOutput {
	return o.ApplyT(func(v *RelationshipLink) ParticipantProfilePropertyReferenceResponseArrayOutput {
		return v.RelatedProfilePropertyReferences
	}).(ParticipantProfilePropertyReferenceResponseArrayOutput)
}

// The relationship guid id.
func (o RelationshipLinkOutput) RelationshipGuidId() pulumi.StringOutput {
	return o.ApplyT(func(v *RelationshipLink) pulumi.StringOutput { return v.RelationshipGuidId }).(pulumi.StringOutput)
}

// The Relationship associated with the Link.
func (o RelationshipLinkOutput) RelationshipName() pulumi.StringOutput {
	return o.ApplyT(func(v *RelationshipLink) pulumi.StringOutput { return v.RelationshipName }).(pulumi.StringOutput)
}

// The hub name.
func (o RelationshipLinkOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *RelationshipLink) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type.
func (o RelationshipLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RelationshipLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RelationshipLinkOutput{})
}
