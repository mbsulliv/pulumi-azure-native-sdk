// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231004preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// MQ mqttBridgeConnector resource
type MqttBridgeConnector struct {
	pulumi.CustomResourceState

	// The number of instances to deploy for a bridge rollout.
	BridgeInstances pulumi.IntPtrOutput `pulumi:"bridgeInstances"`
	// The client id prefix of the dynamically generated client ids.
	ClientIdPrefix pulumi.StringPtrOutput `pulumi:"clientIdPrefix"`
	// Extended Location
	ExtendedLocation ExtendedLocationPropertyResponseOutput `pulumi:"extendedLocation"`
	// The details of MqttBridge Docker Image.
	Image ContainerImageResponseOutput `pulumi:"image"`
	// The details for connecting with Local Broker.
	LocalBrokerConnection LocalBrokerConnectionSpecResponsePtrOutput `pulumi:"localBrokerConnection"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The log level of the Bridge Connector instances.
	LogLevel pulumi.StringPtrOutput `pulumi:"logLevel"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Node Tolerations for the Bridge Connector pods.
	NodeTolerations NodeTolerationsResponsePtrOutput `pulumi:"nodeTolerations"`
	// The protocol to use for connecting with Brokers.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The details for connecting with Remote Broker.
	RemoteBrokerConnection MqttBridgeRemoteBrokerConnectionSpecResponseOutput `pulumi:"remoteBrokerConnection"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMqttBridgeConnector registers a new resource with the given unique name, arguments, and options.
func NewMqttBridgeConnector(ctx *pulumi.Context,
	name string, args *MqttBridgeConnectorArgs, opts ...pulumi.ResourceOption) (*MqttBridgeConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.Image == nil {
		return nil, errors.New("invalid value for required argument 'Image'")
	}
	if args.MqName == nil {
		return nil, errors.New("invalid value for required argument 'MqName'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.RemoteBrokerConnection == nil {
		return nil, errors.New("invalid value for required argument 'RemoteBrokerConnection'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.LocalBrokerConnection != nil {
		args.LocalBrokerConnection = args.LocalBrokerConnection.ToLocalBrokerConnectionSpecPtrOutput().ApplyT(func(v *LocalBrokerConnectionSpec) *LocalBrokerConnectionSpec { return v.Defaults() }).(LocalBrokerConnectionSpecPtrOutput)
	}
	args.RemoteBrokerConnection = args.RemoteBrokerConnection.ToMqttBridgeRemoteBrokerConnectionSpecOutput().ApplyT(func(v MqttBridgeRemoteBrokerConnectionSpec) MqttBridgeRemoteBrokerConnectionSpec {
		return *v.Defaults()
	}).(MqttBridgeRemoteBrokerConnectionSpecOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:iotoperationsmq:MqttBridgeConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MqttBridgeConnector
	err := ctx.RegisterResource("azure-native:iotoperationsmq/v20231004preview:MqttBridgeConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMqttBridgeConnector gets an existing MqttBridgeConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMqttBridgeConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MqttBridgeConnectorState, opts ...pulumi.ResourceOption) (*MqttBridgeConnector, error) {
	var resource MqttBridgeConnector
	err := ctx.ReadResource("azure-native:iotoperationsmq/v20231004preview:MqttBridgeConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MqttBridgeConnector resources.
type mqttBridgeConnectorState struct {
}

type MqttBridgeConnectorState struct {
}

func (MqttBridgeConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*mqttBridgeConnectorState)(nil)).Elem()
}

type mqttBridgeConnectorArgs struct {
	// The number of instances to deploy for a bridge rollout.
	BridgeInstances *int `pulumi:"bridgeInstances"`
	// The client id prefix of the dynamically generated client ids.
	ClientIdPrefix *string `pulumi:"clientIdPrefix"`
	// Extended Location
	ExtendedLocation ExtendedLocationProperty `pulumi:"extendedLocation"`
	// The details of MqttBridge Docker Image.
	Image ContainerImage `pulumi:"image"`
	// The details for connecting with Local Broker.
	LocalBrokerConnection *LocalBrokerConnectionSpec `pulumi:"localBrokerConnection"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The log level of the Bridge Connector instances.
	LogLevel *string `pulumi:"logLevel"`
	// Name of MQ resource
	MqName string `pulumi:"mqName"`
	// Name of MQ mqttBridgeConnector resource
	MqttBridgeConnectorName *string `pulumi:"mqttBridgeConnectorName"`
	// The Node Tolerations for the Bridge Connector pods.
	NodeTolerations *NodeTolerations `pulumi:"nodeTolerations"`
	// The protocol to use for connecting with Brokers.
	Protocol string `pulumi:"protocol"`
	// The details for connecting with Remote Broker.
	RemoteBrokerConnection MqttBridgeRemoteBrokerConnectionSpec `pulumi:"remoteBrokerConnection"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MqttBridgeConnector resource.
type MqttBridgeConnectorArgs struct {
	// The number of instances to deploy for a bridge rollout.
	BridgeInstances pulumi.IntPtrInput
	// The client id prefix of the dynamically generated client ids.
	ClientIdPrefix pulumi.StringPtrInput
	// Extended Location
	ExtendedLocation ExtendedLocationPropertyInput
	// The details of MqttBridge Docker Image.
	Image ContainerImageInput
	// The details for connecting with Local Broker.
	LocalBrokerConnection LocalBrokerConnectionSpecPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The log level of the Bridge Connector instances.
	LogLevel pulumi.StringPtrInput
	// Name of MQ resource
	MqName pulumi.StringInput
	// Name of MQ mqttBridgeConnector resource
	MqttBridgeConnectorName pulumi.StringPtrInput
	// The Node Tolerations for the Bridge Connector pods.
	NodeTolerations NodeTolerationsPtrInput
	// The protocol to use for connecting with Brokers.
	Protocol pulumi.StringInput
	// The details for connecting with Remote Broker.
	RemoteBrokerConnection MqttBridgeRemoteBrokerConnectionSpecInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (MqttBridgeConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mqttBridgeConnectorArgs)(nil)).Elem()
}

type MqttBridgeConnectorInput interface {
	pulumi.Input

	ToMqttBridgeConnectorOutput() MqttBridgeConnectorOutput
	ToMqttBridgeConnectorOutputWithContext(ctx context.Context) MqttBridgeConnectorOutput
}

func (*MqttBridgeConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**MqttBridgeConnector)(nil)).Elem()
}

func (i *MqttBridgeConnector) ToMqttBridgeConnectorOutput() MqttBridgeConnectorOutput {
	return i.ToMqttBridgeConnectorOutputWithContext(context.Background())
}

func (i *MqttBridgeConnector) ToMqttBridgeConnectorOutputWithContext(ctx context.Context) MqttBridgeConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MqttBridgeConnectorOutput)
}

type MqttBridgeConnectorOutput struct{ *pulumi.OutputState }

func (MqttBridgeConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MqttBridgeConnector)(nil)).Elem()
}

func (o MqttBridgeConnectorOutput) ToMqttBridgeConnectorOutput() MqttBridgeConnectorOutput {
	return o
}

func (o MqttBridgeConnectorOutput) ToMqttBridgeConnectorOutputWithContext(ctx context.Context) MqttBridgeConnectorOutput {
	return o
}

// The number of instances to deploy for a bridge rollout.
func (o MqttBridgeConnectorOutput) BridgeInstances() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) pulumi.IntPtrOutput { return v.BridgeInstances }).(pulumi.IntPtrOutput)
}

// The client id prefix of the dynamically generated client ids.
func (o MqttBridgeConnectorOutput) ClientIdPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) pulumi.StringPtrOutput { return v.ClientIdPrefix }).(pulumi.StringPtrOutput)
}

// Extended Location
func (o MqttBridgeConnectorOutput) ExtendedLocation() ExtendedLocationPropertyResponseOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) ExtendedLocationPropertyResponseOutput { return v.ExtendedLocation }).(ExtendedLocationPropertyResponseOutput)
}

// The details of MqttBridge Docker Image.
func (o MqttBridgeConnectorOutput) Image() ContainerImageResponseOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) ContainerImageResponseOutput { return v.Image }).(ContainerImageResponseOutput)
}

// The details for connecting with Local Broker.
func (o MqttBridgeConnectorOutput) LocalBrokerConnection() LocalBrokerConnectionSpecResponsePtrOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) LocalBrokerConnectionSpecResponsePtrOutput {
		return v.LocalBrokerConnection
	}).(LocalBrokerConnectionSpecResponsePtrOutput)
}

// The geo-location where the resource lives
func (o MqttBridgeConnectorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The log level of the Bridge Connector instances.
func (o MqttBridgeConnectorOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) pulumi.StringPtrOutput { return v.LogLevel }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o MqttBridgeConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Node Tolerations for the Bridge Connector pods.
func (o MqttBridgeConnectorOutput) NodeTolerations() NodeTolerationsResponsePtrOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) NodeTolerationsResponsePtrOutput { return v.NodeTolerations }).(NodeTolerationsResponsePtrOutput)
}

// The protocol to use for connecting with Brokers.
func (o MqttBridgeConnectorOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o MqttBridgeConnectorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The details for connecting with Remote Broker.
func (o MqttBridgeConnectorOutput) RemoteBrokerConnection() MqttBridgeRemoteBrokerConnectionSpecResponseOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) MqttBridgeRemoteBrokerConnectionSpecResponseOutput {
		return v.RemoteBrokerConnection
	}).(MqttBridgeRemoteBrokerConnectionSpecResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MqttBridgeConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o MqttBridgeConnectorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MqttBridgeConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MqttBridgeConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MqttBridgeConnectorOutput{})
}
