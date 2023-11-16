// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotoperationsmq

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// MQ broker/listener resource
// Azure REST API version: 2023-10-04-preview.
type BrokerListener struct {
	pulumi.CustomResourceState

	// The flag for enabling Authentication rules on Listener Port.
	AuthenticationEnabled pulumi.BoolPtrOutput `pulumi:"authenticationEnabled"`
	// The flag for enabling Authorization policies on Listener Port. false - AllowAll, true - Use Authorization resource rules if present.
	AuthorizationEnabled pulumi.BoolPtrOutput `pulumi:"authorizationEnabled"`
	// The k8s cr/resource reference of mq/broker.
	BrokerRef pulumi.StringOutput `pulumi:"brokerRef"`
	// Extended Location
	ExtendedLocation ExtendedLocationPropertyResponseOutput `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The node port to use on the Host node.
	NodePort pulumi.IntPtrOutput `pulumi:"nodePort"`
	// The port to start Listening for connections on.
	Port pulumi.IntOutput `pulumi:"port"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The service name to expose Listener port on.
	ServiceName pulumi.StringPtrOutput `pulumi:"serviceName"`
	// The Kubernetes Service type to deploy for Listener.
	ServiceType pulumi.StringPtrOutput `pulumi:"serviceType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Defines configuration of a TLS server certificate. NOTE Enum - Only one TLS Cert method is supported
	Tls TlsCertMethodResponsePtrOutput `pulumi:"tls"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBrokerListener registers a new resource with the given unique name, arguments, and options.
func NewBrokerListener(ctx *pulumi.Context,
	name string, args *BrokerListenerArgs, opts ...pulumi.ResourceOption) (*BrokerListener, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BrokerName == nil {
		return nil, errors.New("invalid value for required argument 'BrokerName'")
	}
	if args.BrokerRef == nil {
		return nil, errors.New("invalid value for required argument 'BrokerRef'")
	}
	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.MqName == nil {
		return nil, errors.New("invalid value for required argument 'MqName'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.AuthenticationEnabled == nil {
		args.AuthenticationEnabled = pulumi.BoolPtr(false)
	}
	if args.AuthorizationEnabled == nil {
		args.AuthorizationEnabled = pulumi.BoolPtr(false)
	}
	if args.ServiceName == nil {
		args.ServiceName = pulumi.StringPtr("aio-mq-dmqtt-frontend")
	}
	if args.ServiceType == nil {
		args.ServiceType = pulumi.StringPtr("clusterIp")
	}
	if args.Tls != nil {
		args.Tls = args.Tls.ToTlsCertMethodPtrOutput().ApplyT(func(v *TlsCertMethod) *TlsCertMethod { return v.Defaults() }).(TlsCertMethodPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:iotoperationsmq/v20231004preview:BrokerListener"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BrokerListener
	err := ctx.RegisterResource("azure-native:iotoperationsmq:BrokerListener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBrokerListener gets an existing BrokerListener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBrokerListener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BrokerListenerState, opts ...pulumi.ResourceOption) (*BrokerListener, error) {
	var resource BrokerListener
	err := ctx.ReadResource("azure-native:iotoperationsmq:BrokerListener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BrokerListener resources.
type brokerListenerState struct {
}

type BrokerListenerState struct {
}

func (BrokerListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*brokerListenerState)(nil)).Elem()
}

type brokerListenerArgs struct {
	// The flag for enabling Authentication rules on Listener Port.
	AuthenticationEnabled *bool `pulumi:"authenticationEnabled"`
	// The flag for enabling Authorization policies on Listener Port. false - AllowAll, true - Use Authorization resource rules if present.
	AuthorizationEnabled *bool `pulumi:"authorizationEnabled"`
	// Name of MQ broker resource
	BrokerName string `pulumi:"brokerName"`
	// The k8s cr/resource reference of mq/broker.
	BrokerRef string `pulumi:"brokerRef"`
	// Extended Location
	ExtendedLocation ExtendedLocationProperty `pulumi:"extendedLocation"`
	// Name of MQ broker/listener resource
	ListenerName *string `pulumi:"listenerName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of MQ resource
	MqName string `pulumi:"mqName"`
	// The node port to use on the Host node.
	NodePort *int `pulumi:"nodePort"`
	// The port to start Listening for connections on.
	Port int `pulumi:"port"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The service name to expose Listener port on.
	ServiceName *string `pulumi:"serviceName"`
	// The Kubernetes Service type to deploy for Listener.
	ServiceType *string `pulumi:"serviceType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Defines configuration of a TLS server certificate. NOTE Enum - Only one TLS Cert method is supported
	Tls *TlsCertMethod `pulumi:"tls"`
}

// The set of arguments for constructing a BrokerListener resource.
type BrokerListenerArgs struct {
	// The flag for enabling Authentication rules on Listener Port.
	AuthenticationEnabled pulumi.BoolPtrInput
	// The flag for enabling Authorization policies on Listener Port. false - AllowAll, true - Use Authorization resource rules if present.
	AuthorizationEnabled pulumi.BoolPtrInput
	// Name of MQ broker resource
	BrokerName pulumi.StringInput
	// The k8s cr/resource reference of mq/broker.
	BrokerRef pulumi.StringInput
	// Extended Location
	ExtendedLocation ExtendedLocationPropertyInput
	// Name of MQ broker/listener resource
	ListenerName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of MQ resource
	MqName pulumi.StringInput
	// The node port to use on the Host node.
	NodePort pulumi.IntPtrInput
	// The port to start Listening for connections on.
	Port pulumi.IntInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The service name to expose Listener port on.
	ServiceName pulumi.StringPtrInput
	// The Kubernetes Service type to deploy for Listener.
	ServiceType pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Defines configuration of a TLS server certificate. NOTE Enum - Only one TLS Cert method is supported
	Tls TlsCertMethodPtrInput
}

func (BrokerListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*brokerListenerArgs)(nil)).Elem()
}

type BrokerListenerInput interface {
	pulumi.Input

	ToBrokerListenerOutput() BrokerListenerOutput
	ToBrokerListenerOutputWithContext(ctx context.Context) BrokerListenerOutput
}

func (*BrokerListener) ElementType() reflect.Type {
	return reflect.TypeOf((**BrokerListener)(nil)).Elem()
}

func (i *BrokerListener) ToBrokerListenerOutput() BrokerListenerOutput {
	return i.ToBrokerListenerOutputWithContext(context.Background())
}

func (i *BrokerListener) ToBrokerListenerOutputWithContext(ctx context.Context) BrokerListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BrokerListenerOutput)
}

type BrokerListenerOutput struct{ *pulumi.OutputState }

func (BrokerListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BrokerListener)(nil)).Elem()
}

func (o BrokerListenerOutput) ToBrokerListenerOutput() BrokerListenerOutput {
	return o
}

func (o BrokerListenerOutput) ToBrokerListenerOutputWithContext(ctx context.Context) BrokerListenerOutput {
	return o
}

// The flag for enabling Authentication rules on Listener Port.
func (o BrokerListenerOutput) AuthenticationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BrokerListener) pulumi.BoolPtrOutput { return v.AuthenticationEnabled }).(pulumi.BoolPtrOutput)
}

// The flag for enabling Authorization policies on Listener Port. false - AllowAll, true - Use Authorization resource rules if present.
func (o BrokerListenerOutput) AuthorizationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BrokerListener) pulumi.BoolPtrOutput { return v.AuthorizationEnabled }).(pulumi.BoolPtrOutput)
}

// The k8s cr/resource reference of mq/broker.
func (o BrokerListenerOutput) BrokerRef() pulumi.StringOutput {
	return o.ApplyT(func(v *BrokerListener) pulumi.StringOutput { return v.BrokerRef }).(pulumi.StringOutput)
}

// Extended Location
func (o BrokerListenerOutput) ExtendedLocation() ExtendedLocationPropertyResponseOutput {
	return o.ApplyT(func(v *BrokerListener) ExtendedLocationPropertyResponseOutput { return v.ExtendedLocation }).(ExtendedLocationPropertyResponseOutput)
}

// The geo-location where the resource lives
func (o BrokerListenerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BrokerListener) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o BrokerListenerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BrokerListener) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The node port to use on the Host node.
func (o BrokerListenerOutput) NodePort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BrokerListener) pulumi.IntPtrOutput { return v.NodePort }).(pulumi.IntPtrOutput)
}

// The port to start Listening for connections on.
func (o BrokerListenerOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *BrokerListener) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The status of the last operation.
func (o BrokerListenerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *BrokerListener) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The service name to expose Listener port on.
func (o BrokerListenerOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BrokerListener) pulumi.StringPtrOutput { return v.ServiceName }).(pulumi.StringPtrOutput)
}

// The Kubernetes Service type to deploy for Listener.
func (o BrokerListenerOutput) ServiceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BrokerListener) pulumi.StringPtrOutput { return v.ServiceType }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o BrokerListenerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BrokerListener) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o BrokerListenerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BrokerListener) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Defines configuration of a TLS server certificate. NOTE Enum - Only one TLS Cert method is supported
func (o BrokerListenerOutput) Tls() TlsCertMethodResponsePtrOutput {
	return o.ApplyT(func(v *BrokerListener) TlsCertMethodResponsePtrOutput { return v.Tls }).(TlsCertMethodResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o BrokerListenerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BrokerListener) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BrokerListenerOutput{})
}
