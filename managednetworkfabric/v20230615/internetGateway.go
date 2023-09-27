// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230615

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The Internet Gateway resource definition.
type InternetGateway struct {
	pulumi.CustomResourceState

	// Switch configuration description.
	Annotation pulumi.StringPtrOutput `pulumi:"annotation"`
	// ARM Resource ID of the Internet Gateway Rule.
	InternetGatewayRuleId pulumi.StringPtrOutput `pulumi:"internetGatewayRuleId"`
	// IPv4 Address of Internet Gateway.
	Ipv4Address pulumi.StringOutput `pulumi:"ipv4Address"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// ARM Resource ID of the Network Fabric Controller.
	NetworkFabricControllerId pulumi.StringOutput `pulumi:"networkFabricControllerId"`
	// Port number of Internet Gateway.
	Port pulumi.IntOutput `pulumi:"port"`
	// Provisioning state of resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewInternetGateway registers a new resource with the given unique name, arguments, and options.
func NewInternetGateway(ctx *pulumi.Context,
	name string, args *InternetGatewayArgs, opts ...pulumi.ResourceOption) (*InternetGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkFabricControllerId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkFabricControllerId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetworkfabric:InternetGateway"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource InternetGateway
	err := ctx.RegisterResource("azure-native:managednetworkfabric/v20230615:InternetGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInternetGateway gets an existing InternetGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInternetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InternetGatewayState, opts ...pulumi.ResourceOption) (*InternetGateway, error) {
	var resource InternetGateway
	err := ctx.ReadResource("azure-native:managednetworkfabric/v20230615:InternetGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InternetGateway resources.
type internetGatewayState struct {
}

type InternetGatewayState struct {
}

func (InternetGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*internetGatewayState)(nil)).Elem()
}

type internetGatewayArgs struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Name of the Internet Gateway.
	InternetGatewayName *string `pulumi:"internetGatewayName"`
	// ARM Resource ID of the Internet Gateway Rule.
	InternetGatewayRuleId *string `pulumi:"internetGatewayRuleId"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// ARM Resource ID of the Network Fabric Controller.
	NetworkFabricControllerId string `pulumi:"networkFabricControllerId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gateway Type of the resource.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a InternetGateway resource.
type InternetGatewayArgs struct {
	// Switch configuration description.
	Annotation pulumi.StringPtrInput
	// Name of the Internet Gateway.
	InternetGatewayName pulumi.StringPtrInput
	// ARM Resource ID of the Internet Gateway Rule.
	InternetGatewayRuleId pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// ARM Resource ID of the Network Fabric Controller.
	NetworkFabricControllerId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Gateway Type of the resource.
	Type pulumi.StringInput
}

func (InternetGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*internetGatewayArgs)(nil)).Elem()
}

type InternetGatewayInput interface {
	pulumi.Input

	ToInternetGatewayOutput() InternetGatewayOutput
	ToInternetGatewayOutputWithContext(ctx context.Context) InternetGatewayOutput
}

func (*InternetGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**InternetGateway)(nil)).Elem()
}

func (i *InternetGateway) ToInternetGatewayOutput() InternetGatewayOutput {
	return i.ToInternetGatewayOutputWithContext(context.Background())
}

func (i *InternetGateway) ToInternetGatewayOutputWithContext(ctx context.Context) InternetGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetGatewayOutput)
}

func (i *InternetGateway) ToOutput(ctx context.Context) pulumix.Output[*InternetGateway] {
	return pulumix.Output[*InternetGateway]{
		OutputState: i.ToInternetGatewayOutputWithContext(ctx).OutputState,
	}
}

type InternetGatewayOutput struct{ *pulumi.OutputState }

func (InternetGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InternetGateway)(nil)).Elem()
}

func (o InternetGatewayOutput) ToInternetGatewayOutput() InternetGatewayOutput {
	return o
}

func (o InternetGatewayOutput) ToInternetGatewayOutputWithContext(ctx context.Context) InternetGatewayOutput {
	return o
}

func (o InternetGatewayOutput) ToOutput(ctx context.Context) pulumix.Output[*InternetGateway] {
	return pulumix.Output[*InternetGateway]{
		OutputState: o.OutputState,
	}
}

// Switch configuration description.
func (o InternetGatewayOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InternetGateway) pulumi.StringPtrOutput { return v.Annotation }).(pulumi.StringPtrOutput)
}

// ARM Resource ID of the Internet Gateway Rule.
func (o InternetGatewayOutput) InternetGatewayRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InternetGateway) pulumi.StringPtrOutput { return v.InternetGatewayRuleId }).(pulumi.StringPtrOutput)
}

// IPv4 Address of Internet Gateway.
func (o InternetGatewayOutput) Ipv4Address() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetGateway) pulumi.StringOutput { return v.Ipv4Address }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o InternetGatewayOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetGateway) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o InternetGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ARM Resource ID of the Network Fabric Controller.
func (o InternetGatewayOutput) NetworkFabricControllerId() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetGateway) pulumi.StringOutput { return v.NetworkFabricControllerId }).(pulumi.StringOutput)
}

// Port number of Internet Gateway.
func (o InternetGatewayOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *InternetGateway) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Provisioning state of resource.
func (o InternetGatewayOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetGateway) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o InternetGatewayOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *InternetGateway) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o InternetGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InternetGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o InternetGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InternetGatewayOutput{})
}
