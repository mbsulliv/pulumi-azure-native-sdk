// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A hub setting
type WebPubSubHub struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of a hub.
	Properties WebPubSubHubPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebPubSubHub registers a new resource with the given unique name, arguments, and options.
func NewWebPubSubHub(ctx *pulumi.Context,
	name string, args *WebPubSubHubArgs, opts ...pulumi.ResourceOption) (*WebPubSubHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	args.Properties = args.Properties.ToWebPubSubHubPropertiesOutput().ApplyT(func(v WebPubSubHubProperties) WebPubSubHubProperties { return *v.Defaults() }).(WebPubSubHubPropertiesOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:webpubsub:WebPubSubHub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20211001:WebPubSubHub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20220801preview:WebPubSubHub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230301preview:WebPubSubHub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230601preview:WebPubSubHub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230801preview:WebPubSubHub"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebPubSubHub
	err := ctx.RegisterResource("azure-native:webpubsub/v20230201:WebPubSubHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebPubSubHub gets an existing WebPubSubHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebPubSubHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebPubSubHubState, opts ...pulumi.ResourceOption) (*WebPubSubHub, error) {
	var resource WebPubSubHub
	err := ctx.ReadResource("azure-native:webpubsub/v20230201:WebPubSubHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebPubSubHub resources.
type webPubSubHubState struct {
}

type WebPubSubHubState struct {
}

func (WebPubSubHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubHubState)(nil)).Elem()
}

type webPubSubHubArgs struct {
	// The hub name.
	HubName *string `pulumi:"hubName"`
	// Properties of a hub.
	Properties WebPubSubHubProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a WebPubSubHub resource.
type WebPubSubHubArgs struct {
	// The hub name.
	HubName pulumi.StringPtrInput
	// Properties of a hub.
	Properties WebPubSubHubPropertiesInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the resource.
	ResourceName pulumi.StringInput
}

func (WebPubSubHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubHubArgs)(nil)).Elem()
}

type WebPubSubHubInput interface {
	pulumi.Input

	ToWebPubSubHubOutput() WebPubSubHubOutput
	ToWebPubSubHubOutputWithContext(ctx context.Context) WebPubSubHubOutput
}

func (*WebPubSubHub) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubHub)(nil)).Elem()
}

func (i *WebPubSubHub) ToWebPubSubHubOutput() WebPubSubHubOutput {
	return i.ToWebPubSubHubOutputWithContext(context.Background())
}

func (i *WebPubSubHub) ToWebPubSubHubOutputWithContext(ctx context.Context) WebPubSubHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubHubOutput)
}

func (i *WebPubSubHub) ToOutput(ctx context.Context) pulumix.Output[*WebPubSubHub] {
	return pulumix.Output[*WebPubSubHub]{
		OutputState: i.ToWebPubSubHubOutputWithContext(ctx).OutputState,
	}
}

type WebPubSubHubOutput struct{ *pulumi.OutputState }

func (WebPubSubHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubHub)(nil)).Elem()
}

func (o WebPubSubHubOutput) ToWebPubSubHubOutput() WebPubSubHubOutput {
	return o
}

func (o WebPubSubHubOutput) ToWebPubSubHubOutputWithContext(ctx context.Context) WebPubSubHubOutput {
	return o
}

func (o WebPubSubHubOutput) ToOutput(ctx context.Context) pulumix.Output[*WebPubSubHub] {
	return pulumix.Output[*WebPubSubHub]{
		OutputState: o.OutputState,
	}
}

// The name of the resource.
func (o WebPubSubHubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubHub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of a hub.
func (o WebPubSubHubOutput) Properties() WebPubSubHubPropertiesResponseOutput {
	return o.ApplyT(func(v *WebPubSubHub) WebPubSubHubPropertiesResponseOutput { return v.Properties }).(WebPubSubHubPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o WebPubSubHubOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebPubSubHub) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
func (o WebPubSubHubOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubHub) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebPubSubHubOutput{})
}
