// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// EventGrid Topic
type Topic struct {
	pulumi.CustomResourceState

	// Endpoint for the topic.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Identity information for the resource.
	Identity IdentityInfoResponsePtrOutput `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules InboundIpRuleResponseArrayOutput `pulumi:"inboundIpRules"`
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema pulumi.StringPtrOutput `pulumi:"inputSchema"`
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping JsonInputSchemaMappingResponsePtrOutput `pulumi:"inputSchemaMapping"`
	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Metric resource id for the topic.
	MetricResourceId pulumi.StringOutput `pulumi:"metricResourceId"`
	// Name of the resource.
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Provisioning state of the topic.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.TopicProperties.InboundIpRules" />
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The Sku pricing tier for the topic.
	Sku ResourceSkuResponsePtrOutput `pulumi:"sku"`
	// Tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.InputSchema == nil {
		args.InputSchema = pulumi.StringPtr("EventGridSchema")
	}
	if args.PublicNetworkAccess == nil {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20170615preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20170915preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180101:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180501preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180915preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190101:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190201preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190601:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200101preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200601:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211201:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20230601preview:Topic"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Topic
	err := ctx.RegisterResource("azure-native:eventgrid/v20200401preview:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("azure-native:eventgrid/v20200401preview:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
}

type TopicState struct {
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// Identity information for the resource.
	Identity *IdentityInfo `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRule `pulumi:"inboundIpRules"`
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema *string `pulumi:"inputSchema"`
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping *JsonInputSchemaMapping `pulumi:"inputSchemaMapping"`
	// Location of the resource.
	Location                   *string                     `pulumi:"location"`
	PrivateEndpointConnections []PrivateEndpointConnection `pulumi:"privateEndpointConnections"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.TopicProperties.InboundIpRules" />
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Sku pricing tier for the topic.
	Sku *ResourceSku `pulumi:"sku"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Name of the topic.
	TopicName *string `pulumi:"topicName"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// Identity information for the resource.
	Identity IdentityInfoPtrInput
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules InboundIpRuleArrayInput
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema pulumi.StringPtrInput
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping JsonInputSchemaMappingPtrInput
	// Location of the resource.
	Location                   pulumi.StringPtrInput
	PrivateEndpointConnections PrivateEndpointConnectionArrayInput
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.TopicProperties.InboundIpRules" />
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// The Sku pricing tier for the topic.
	Sku ResourceSkuPtrInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
	// Name of the topic.
	TopicName pulumi.StringPtrInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

func (i *Topic) ToOutput(ctx context.Context) pulumix.Output[*Topic] {
	return pulumix.Output[*Topic]{
		OutputState: i.ToTopicOutputWithContext(ctx).OutputState,
	}
}

type TopicOutput struct{ *pulumi.OutputState }

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

func (o TopicOutput) ToOutput(ctx context.Context) pulumix.Output[*Topic] {
	return pulumix.Output[*Topic]{
		OutputState: o.OutputState,
	}
}

// Endpoint for the topic.
func (o TopicOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Identity information for the resource.
func (o TopicOutput) Identity() IdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v *Topic) IdentityInfoResponsePtrOutput { return v.Identity }).(IdentityInfoResponsePtrOutput)
}

// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
func (o TopicOutput) InboundIpRules() InboundIpRuleResponseArrayOutput {
	return o.ApplyT(func(v *Topic) InboundIpRuleResponseArrayOutput { return v.InboundIpRules }).(InboundIpRuleResponseArrayOutput)
}

// This determines the format that Event Grid should expect for incoming events published to the topic.
func (o TopicOutput) InputSchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.InputSchema }).(pulumi.StringPtrOutput)
}

// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
func (o TopicOutput) InputSchemaMapping() JsonInputSchemaMappingResponsePtrOutput {
	return o.ApplyT(func(v *Topic) JsonInputSchemaMappingResponsePtrOutput { return v.InputSchemaMapping }).(JsonInputSchemaMappingResponsePtrOutput)
}

// Location of the resource.
func (o TopicOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Metric resource id for the topic.
func (o TopicOutput) MetricResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.MetricResourceId }).(pulumi.StringOutput)
}

// Name of the resource.
func (o TopicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TopicOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Topic) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the topic.
func (o TopicOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// This determines if traffic is allowed over public network. By default it is enabled.
// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.TopicProperties.InboundIpRules" />
func (o TopicOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The Sku pricing tier for the topic.
func (o TopicOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v *Topic) ResourceSkuResponsePtrOutput { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// Tags of the resource.
func (o TopicOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the resource.
func (o TopicOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TopicOutput{})
}
