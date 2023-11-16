// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Hybrid Connection contract. This is used to configure a Hybrid Connection.
type WebAppHybridConnection struct {
	pulumi.CustomResourceState

	// The hostname of the endpoint.
	Hostname pulumi.StringPtrOutput `pulumi:"hostname"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The port of the endpoint.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The ARM URI to the Service Bus relay.
	RelayArmUri pulumi.StringPtrOutput `pulumi:"relayArmUri"`
	// The name of the Service Bus relay.
	RelayName pulumi.StringPtrOutput `pulumi:"relayName"`
	// The name of the Service Bus key which has Send permissions. This is used to authenticate to Service Bus.
	SendKeyName pulumi.StringPtrOutput `pulumi:"sendKeyName"`
	// The value of the Service Bus key. This is used to authenticate to Service Bus. In ARM this key will not be returned
	// normally, use the POST /listKeys API instead.
	SendKeyValue pulumi.StringPtrOutput `pulumi:"sendKeyValue"`
	// The name of the Service Bus namespace.
	ServiceBusNamespace pulumi.StringPtrOutput `pulumi:"serviceBusNamespace"`
	// The suffix for the service bus endpoint. By default this is .servicebus.windows.net
	ServiceBusSuffix pulumi.StringPtrOutput `pulumi:"serviceBusSuffix"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppHybridConnection registers a new resource with the given unique name, arguments, and options.
func NewWebAppHybridConnection(ctx *pulumi.Context,
	name string, args *WebAppHybridConnectionArgs, opts ...pulumi.ResourceOption) (*WebAppHybridConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppHybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:WebAppHybridConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppHybridConnection
	err := ctx.RegisterResource("azure-native:web/v20201001:WebAppHybridConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppHybridConnection gets an existing WebAppHybridConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppHybridConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppHybridConnectionState, opts ...pulumi.ResourceOption) (*WebAppHybridConnection, error) {
	var resource WebAppHybridConnection
	err := ctx.ReadResource("azure-native:web/v20201001:WebAppHybridConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppHybridConnection resources.
type webAppHybridConnectionState struct {
}

type WebAppHybridConnectionState struct {
}

func (WebAppHybridConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppHybridConnectionState)(nil)).Elem()
}

type webAppHybridConnectionArgs struct {
	// The hostname of the endpoint.
	Hostname *string `pulumi:"hostname"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// The name of the web app.
	Name string `pulumi:"name"`
	// The namespace for this hybrid connection.
	NamespaceName string `pulumi:"namespaceName"`
	// The port of the endpoint.
	Port *int `pulumi:"port"`
	// The ARM URI to the Service Bus relay.
	RelayArmUri *string `pulumi:"relayArmUri"`
	// The name of the Service Bus relay.
	RelayName *string `pulumi:"relayName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service Bus key which has Send permissions. This is used to authenticate to Service Bus.
	SendKeyName *string `pulumi:"sendKeyName"`
	// The value of the Service Bus key. This is used to authenticate to Service Bus. In ARM this key will not be returned
	// normally, use the POST /listKeys API instead.
	SendKeyValue *string `pulumi:"sendKeyValue"`
	// The name of the Service Bus namespace.
	ServiceBusNamespace *string `pulumi:"serviceBusNamespace"`
	// The suffix for the service bus endpoint. By default this is .servicebus.windows.net
	ServiceBusSuffix *string `pulumi:"serviceBusSuffix"`
}

// The set of arguments for constructing a WebAppHybridConnection resource.
type WebAppHybridConnectionArgs struct {
	// The hostname of the endpoint.
	Hostname pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// The name of the web app.
	Name pulumi.StringInput
	// The namespace for this hybrid connection.
	NamespaceName pulumi.StringInput
	// The port of the endpoint.
	Port pulumi.IntPtrInput
	// The ARM URI to the Service Bus relay.
	RelayArmUri pulumi.StringPtrInput
	// The name of the Service Bus relay.
	RelayName pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// The name of the Service Bus key which has Send permissions. This is used to authenticate to Service Bus.
	SendKeyName pulumi.StringPtrInput
	// The value of the Service Bus key. This is used to authenticate to Service Bus. In ARM this key will not be returned
	// normally, use the POST /listKeys API instead.
	SendKeyValue pulumi.StringPtrInput
	// The name of the Service Bus namespace.
	ServiceBusNamespace pulumi.StringPtrInput
	// The suffix for the service bus endpoint. By default this is .servicebus.windows.net
	ServiceBusSuffix pulumi.StringPtrInput
}

func (WebAppHybridConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppHybridConnectionArgs)(nil)).Elem()
}

type WebAppHybridConnectionInput interface {
	pulumi.Input

	ToWebAppHybridConnectionOutput() WebAppHybridConnectionOutput
	ToWebAppHybridConnectionOutputWithContext(ctx context.Context) WebAppHybridConnectionOutput
}

func (*WebAppHybridConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppHybridConnection)(nil)).Elem()
}

func (i *WebAppHybridConnection) ToWebAppHybridConnectionOutput() WebAppHybridConnectionOutput {
	return i.ToWebAppHybridConnectionOutputWithContext(context.Background())
}

func (i *WebAppHybridConnection) ToWebAppHybridConnectionOutputWithContext(ctx context.Context) WebAppHybridConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppHybridConnectionOutput)
}

type WebAppHybridConnectionOutput struct{ *pulumi.OutputState }

func (WebAppHybridConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppHybridConnection)(nil)).Elem()
}

func (o WebAppHybridConnectionOutput) ToWebAppHybridConnectionOutput() WebAppHybridConnectionOutput {
	return o
}

func (o WebAppHybridConnectionOutput) ToWebAppHybridConnectionOutputWithContext(ctx context.Context) WebAppHybridConnectionOutput {
	return o
}

// The hostname of the endpoint.
func (o WebAppHybridConnectionOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.Hostname }).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o WebAppHybridConnectionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppHybridConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The port of the endpoint.
func (o WebAppHybridConnectionOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// The ARM URI to the Service Bus relay.
func (o WebAppHybridConnectionOutput) RelayArmUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.RelayArmUri }).(pulumi.StringPtrOutput)
}

// The name of the Service Bus relay.
func (o WebAppHybridConnectionOutput) RelayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.RelayName }).(pulumi.StringPtrOutput)
}

// The name of the Service Bus key which has Send permissions. This is used to authenticate to Service Bus.
func (o WebAppHybridConnectionOutput) SendKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.SendKeyName }).(pulumi.StringPtrOutput)
}

// The value of the Service Bus key. This is used to authenticate to Service Bus. In ARM this key will not be returned
// normally, use the POST /listKeys API instead.
func (o WebAppHybridConnectionOutput) SendKeyValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.SendKeyValue }).(pulumi.StringPtrOutput)
}

// The name of the Service Bus namespace.
func (o WebAppHybridConnectionOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

// The suffix for the service bus endpoint. By default this is .servicebus.windows.net
func (o WebAppHybridConnectionOutput) ServiceBusSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringPtrOutput { return v.ServiceBusSuffix }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o WebAppHybridConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o WebAppHybridConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHybridConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppHybridConnectionOutput{})
}
