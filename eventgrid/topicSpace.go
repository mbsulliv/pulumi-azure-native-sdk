// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Topic space resource.
// Azure REST API version: 2023-06-01-preview.
//
// Other available API versions: 2023-12-15-preview.
type TopicSpace struct {
	pulumi.CustomResourceState

	// Description for the Topic Space resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the TopicSpace resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system metadata relating to the TopicSpace resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The topic filters in the topic space.
	// Example: "topicTemplates": [
	//               "devices/foo/bar",
	//               "devices/topic1/+",
	//               "devices/${principal.name}/${principal.attributes.keyName}" ].
	TopicTemplates pulumi.StringArrayOutput `pulumi:"topicTemplates"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTopicSpace registers a new resource with the given unique name, arguments, and options.
func NewTopicSpace(ctx *pulumi.Context,
	name string, args *TopicSpaceArgs, opts ...pulumi.ResourceOption) (*TopicSpace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid/v20230601preview:TopicSpace"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20231215preview:TopicSpace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TopicSpace
	err := ctx.RegisterResource("azure-native:eventgrid:TopicSpace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicSpace gets an existing TopicSpace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicSpace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicSpaceState, opts ...pulumi.ResourceOption) (*TopicSpace, error) {
	var resource TopicSpace
	err := ctx.ReadResource("azure-native:eventgrid:TopicSpace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicSpace resources.
type topicSpaceState struct {
}

type TopicSpaceState struct {
}

func (TopicSpaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicSpaceState)(nil)).Elem()
}

type topicSpaceArgs struct {
	// Description for the Topic Space resource.
	Description *string `pulumi:"description"`
	// Name of the namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The topic space name.
	TopicSpaceName *string `pulumi:"topicSpaceName"`
	// The topic filters in the topic space.
	// Example: "topicTemplates": [
	//               "devices/foo/bar",
	//               "devices/topic1/+",
	//               "devices/${principal.name}/${principal.attributes.keyName}" ].
	TopicTemplates []string `pulumi:"topicTemplates"`
}

// The set of arguments for constructing a TopicSpace resource.
type TopicSpaceArgs struct {
	// Description for the Topic Space resource.
	Description pulumi.StringPtrInput
	// Name of the namespace.
	NamespaceName pulumi.StringInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// The topic space name.
	TopicSpaceName pulumi.StringPtrInput
	// The topic filters in the topic space.
	// Example: "topicTemplates": [
	//               "devices/foo/bar",
	//               "devices/topic1/+",
	//               "devices/${principal.name}/${principal.attributes.keyName}" ].
	TopicTemplates pulumi.StringArrayInput
}

func (TopicSpaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicSpaceArgs)(nil)).Elem()
}

type TopicSpaceInput interface {
	pulumi.Input

	ToTopicSpaceOutput() TopicSpaceOutput
	ToTopicSpaceOutputWithContext(ctx context.Context) TopicSpaceOutput
}

func (*TopicSpace) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicSpace)(nil)).Elem()
}

func (i *TopicSpace) ToTopicSpaceOutput() TopicSpaceOutput {
	return i.ToTopicSpaceOutputWithContext(context.Background())
}

func (i *TopicSpace) ToTopicSpaceOutputWithContext(ctx context.Context) TopicSpaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSpaceOutput)
}

type TopicSpaceOutput struct{ *pulumi.OutputState }

func (TopicSpaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicSpace)(nil)).Elem()
}

func (o TopicSpaceOutput) ToTopicSpaceOutput() TopicSpaceOutput {
	return o
}

func (o TopicSpaceOutput) ToTopicSpaceOutputWithContext(ctx context.Context) TopicSpaceOutput {
	return o
}

// Description for the Topic Space resource.
func (o TopicSpaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TopicSpace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o TopicSpaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TopicSpace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the TopicSpace resource.
func (o TopicSpaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *TopicSpace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system metadata relating to the TopicSpace resource.
func (o TopicSpaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TopicSpace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The topic filters in the topic space.
// Example: "topicTemplates": [
//
//	"devices/foo/bar",
//	"devices/topic1/+",
//	"devices/${principal.name}/${principal.attributes.keyName}" ].
func (o TopicSpaceOutput) TopicTemplates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TopicSpace) pulumi.StringArrayOutput { return v.TopicTemplates }).(pulumi.StringArrayOutput)
}

// Type of the resource.
func (o TopicSpaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TopicSpace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TopicSpaceOutput{})
}
