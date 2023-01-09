// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A CommunicationsGateway resource
type CommunicationsGateway struct {
	pulumi.CustomResourceState

	// Details of API bridge functionality, if required
	ApiBridge pulumi.AnyOutput `pulumi:"apiBridge"`
	// Voice codecs to support
	Codecs pulumi.StringArrayOutput `pulumi:"codecs"`
	// How to connect back to the operator network, e.g. MAPS
	Connectivity pulumi.StringOutput `pulumi:"connectivity"`
	// How to handle 911 calls
	E911Type pulumi.StringOutput `pulumi:"e911Type"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// What platforms to support
	Platforms pulumi.StringArrayOutput `pulumi:"platforms"`
	// Resource provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The regions in which to deploy the resources needed for Teams Calling
	ServiceLocations ServiceRegionPropertiesResponseArrayOutput `pulumi:"serviceLocations"`
	// The current status of the deployment.
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCommunicationsGateway registers a new resource with the given unique name, arguments, and options.
func NewCommunicationsGateway(ctx *pulumi.Context,
	name string, args *CommunicationsGatewayArgs, opts ...pulumi.ResourceOption) (*CommunicationsGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Codecs == nil {
		return nil, errors.New("invalid value for required argument 'Codecs'")
	}
	if args.Connectivity == nil {
		return nil, errors.New("invalid value for required argument 'Connectivity'")
	}
	if args.E911Type == nil {
		return nil, errors.New("invalid value for required argument 'E911Type'")
	}
	if args.Platforms == nil {
		return nil, errors.New("invalid value for required argument 'Platforms'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceLocations == nil {
		return nil, errors.New("invalid value for required argument 'ServiceLocations'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:voiceservices:CommunicationsGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource CommunicationsGateway
	err := ctx.RegisterResource("azure-native:voiceservices/v20221201preview:CommunicationsGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommunicationsGateway gets an existing CommunicationsGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommunicationsGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommunicationsGatewayState, opts ...pulumi.ResourceOption) (*CommunicationsGateway, error) {
	var resource CommunicationsGateway
	err := ctx.ReadResource("azure-native:voiceservices/v20221201preview:CommunicationsGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CommunicationsGateway resources.
type communicationsGatewayState struct {
}

type CommunicationsGatewayState struct {
}

func (CommunicationsGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*communicationsGatewayState)(nil)).Elem()
}

type communicationsGatewayArgs struct {
	// Details of API bridge functionality, if required
	ApiBridge interface{} `pulumi:"apiBridge"`
	// Voice codecs to support
	Codecs []string `pulumi:"codecs"`
	// Unique identifier for this deployment
	CommunicationsGatewayName *string `pulumi:"communicationsGatewayName"`
	// How to connect back to the operator network, e.g. MAPS
	Connectivity string `pulumi:"connectivity"`
	// How to handle 911 calls
	E911Type string `pulumi:"e911Type"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// What platforms to support
	Platforms []string `pulumi:"platforms"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The regions in which to deploy the resources needed for Teams Calling
	ServiceLocations []ServiceRegionProperties `pulumi:"serviceLocations"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CommunicationsGateway resource.
type CommunicationsGatewayArgs struct {
	// Details of API bridge functionality, if required
	ApiBridge pulumi.Input
	// Voice codecs to support
	Codecs pulumi.StringArrayInput
	// Unique identifier for this deployment
	CommunicationsGatewayName pulumi.StringPtrInput
	// How to connect back to the operator network, e.g. MAPS
	Connectivity pulumi.StringInput
	// How to handle 911 calls
	E911Type pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// What platforms to support
	Platforms pulumi.StringArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The regions in which to deploy the resources needed for Teams Calling
	ServiceLocations ServiceRegionPropertiesArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (CommunicationsGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*communicationsGatewayArgs)(nil)).Elem()
}

type CommunicationsGatewayInput interface {
	pulumi.Input

	ToCommunicationsGatewayOutput() CommunicationsGatewayOutput
	ToCommunicationsGatewayOutputWithContext(ctx context.Context) CommunicationsGatewayOutput
}

func (*CommunicationsGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**CommunicationsGateway)(nil)).Elem()
}

func (i *CommunicationsGateway) ToCommunicationsGatewayOutput() CommunicationsGatewayOutput {
	return i.ToCommunicationsGatewayOutputWithContext(context.Background())
}

func (i *CommunicationsGateway) ToCommunicationsGatewayOutputWithContext(ctx context.Context) CommunicationsGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommunicationsGatewayOutput)
}

type CommunicationsGatewayOutput struct{ *pulumi.OutputState }

func (CommunicationsGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommunicationsGateway)(nil)).Elem()
}

func (o CommunicationsGatewayOutput) ToCommunicationsGatewayOutput() CommunicationsGatewayOutput {
	return o
}

func (o CommunicationsGatewayOutput) ToCommunicationsGatewayOutputWithContext(ctx context.Context) CommunicationsGatewayOutput {
	return o
}

// Details of API bridge functionality, if required
func (o CommunicationsGatewayOutput) ApiBridge() pulumi.AnyOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.AnyOutput { return v.ApiBridge }).(pulumi.AnyOutput)
}

// Voice codecs to support
func (o CommunicationsGatewayOutput) Codecs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringArrayOutput { return v.Codecs }).(pulumi.StringArrayOutput)
}

// How to connect back to the operator network, e.g. MAPS
func (o CommunicationsGatewayOutput) Connectivity() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringOutput { return v.Connectivity }).(pulumi.StringOutput)
}

// How to handle 911 calls
func (o CommunicationsGatewayOutput) E911Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringOutput { return v.E911Type }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o CommunicationsGatewayOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o CommunicationsGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// What platforms to support
func (o CommunicationsGatewayOutput) Platforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringArrayOutput { return v.Platforms }).(pulumi.StringArrayOutput)
}

// Resource provisioning state.
func (o CommunicationsGatewayOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The regions in which to deploy the resources needed for Teams Calling
func (o CommunicationsGatewayOutput) ServiceLocations() ServiceRegionPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *CommunicationsGateway) ServiceRegionPropertiesResponseArrayOutput { return v.ServiceLocations }).(ServiceRegionPropertiesResponseArrayOutput)
}

// The current status of the deployment.
func (o CommunicationsGatewayOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o CommunicationsGatewayOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CommunicationsGateway) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o CommunicationsGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CommunicationsGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CommunicationsGatewayOutput{})
}