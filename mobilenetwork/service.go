// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mobilenetwork

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Service resource. Must be created in the same location as its parent mobile network.
// Azure REST API version: 2023-06-01. Prior API version in Azure Native 1.x: 2022-04-01-preview.
//
// Other available API versions: 2022-04-01-preview, 2022-11-01, 2023-09-01.
type Service struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The set of data flow policy rules that make up this service.
	PccRules PccRuleConfigurationResponseArrayOutput `pulumi:"pccRules"`
	// The provisioning state of the service resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// A precedence value that is used to decide between services when identifying the QoS values to use for a particular SIM. A lower value means a higher priority. This value should be unique among all services configured in the mobile network.
	ServicePrecedence pulumi.IntOutput `pulumi:"servicePrecedence"`
	// The QoS policy to use for packets matching this service. This can be overridden for particular flows using the ruleQosPolicy field in a PccRuleConfiguration. If this field is null then the UE's SIM policy will define the QoS settings.
	ServiceQosPolicy QosPolicyResponsePtrOutput `pulumi:"serviceQosPolicy"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MobileNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'MobileNetworkName'")
	}
	if args.PccRules == nil {
		return nil, errors.New("invalid value for required argument 'PccRules'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServicePrecedence == nil {
		return nil, errors.New("invalid value for required argument 'ServicePrecedence'")
	}
	if args.ServiceQosPolicy != nil {
		args.ServiceQosPolicy = args.ServiceQosPolicy.ToQosPolicyPtrOutput().ApplyT(func(v *QosPolicy) *QosPolicy { return v.Defaults() }).(QosPolicyPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20221101:Service"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20230601:Service"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20230901:Service"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Service
	err := ctx.RegisterResource("azure-native:mobilenetwork:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:mobilenetwork:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
}

type ServiceState struct {
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the mobile network.
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	// The set of data flow policy rules that make up this service.
	PccRules []PccRuleConfiguration `pulumi:"pccRules"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service. You must not use any of the following reserved strings - `default`, `requested` or `service`
	ServiceName *string `pulumi:"serviceName"`
	// A precedence value that is used to decide between services when identifying the QoS values to use for a particular SIM. A lower value means a higher priority. This value should be unique among all services configured in the mobile network.
	ServicePrecedence int `pulumi:"servicePrecedence"`
	// The QoS policy to use for packets matching this service. This can be overridden for particular flows using the ruleQosPolicy field in a PccRuleConfiguration. If this field is null then the UE's SIM policy will define the QoS settings.
	ServiceQosPolicy *QosPolicy `pulumi:"serviceQosPolicy"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the mobile network.
	MobileNetworkName pulumi.StringInput
	// The set of data flow policy rules that make up this service.
	PccRules PccRuleConfigurationArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the service. You must not use any of the following reserved strings - `default`, `requested` or `service`
	ServiceName pulumi.StringPtrInput
	// A precedence value that is used to decide between services when identifying the QoS values to use for a particular SIM. A lower value means a higher priority. This value should be unique among all services configured in the mobile network.
	ServicePrecedence pulumi.IntInput
	// The QoS policy to use for packets matching this service. This can be overridden for particular flows using the ruleQosPolicy field in a PccRuleConfiguration. If this field is null then the UE's SIM policy will define the QoS settings.
	ServiceQosPolicy QosPolicyPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

// The geo-location where the resource lives
func (o ServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The set of data flow policy rules that make up this service.
func (o ServiceOutput) PccRules() PccRuleConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *Service) PccRuleConfigurationResponseArrayOutput { return v.PccRules }).(PccRuleConfigurationResponseArrayOutput)
}

// The provisioning state of the service resource.
func (o ServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// A precedence value that is used to decide between services when identifying the QoS values to use for a particular SIM. A lower value means a higher priority. This value should be unique among all services configured in the mobile network.
func (o ServiceOutput) ServicePrecedence() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.ServicePrecedence }).(pulumi.IntOutput)
}

// The QoS policy to use for packets matching this service. This can be overridden for particular flows using the ruleQosPolicy field in a PccRuleConfiguration. If this field is null then the UE's SIM policy will define the QoS settings.
func (o ServiceOutput) ServiceQosPolicy() QosPolicyResponsePtrOutput {
	return o.ApplyT(func(v *Service) QosPolicyResponsePtrOutput { return v.ServiceQosPolicy }).(QosPolicyResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Service) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
