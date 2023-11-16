// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package avs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ExpressRoute Circuit Authorization
// Azure REST API version: 2022-05-01. Prior API version in Azure Native 1.x: 2020-03-20.
//
// Other available API versions: 2023-03-01.
type Authorization struct {
	pulumi.CustomResourceState

	// The ID of the ExpressRoute Circuit Authorization
	ExpressRouteAuthorizationId pulumi.StringOutput `pulumi:"expressRouteAuthorizationId"`
	// The key of the ExpressRoute Circuit Authorization
	ExpressRouteAuthorizationKey pulumi.StringOutput `pulumi:"expressRouteAuthorizationKey"`
	// The ID of the ExpressRoute Circuit
	ExpressRouteId pulumi.StringPtrOutput `pulumi:"expressRouteId"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The state of the  ExpressRoute Circuit Authorization provisioning
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAuthorization registers a new resource with the given unique name, arguments, and options.
func NewAuthorization(ctx *pulumi.Context,
	name string, args *AuthorizationArgs, opts ...pulumi.ResourceOption) (*Authorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs/v20200320:Authorization"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:Authorization"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:Authorization"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:Authorization"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:Authorization"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:Authorization"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230301:Authorization"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Authorization
	err := ctx.RegisterResource("azure-native:avs:Authorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorization gets an existing Authorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationState, opts ...pulumi.ResourceOption) (*Authorization, error) {
	var resource Authorization
	err := ctx.ReadResource("azure-native:avs:Authorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Authorization resources.
type authorizationState struct {
}

type AuthorizationState struct {
}

func (AuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationState)(nil)).Elem()
}

type authorizationArgs struct {
	// Name of the ExpressRoute Circuit Authorization in the private cloud
	AuthorizationName *string `pulumi:"authorizationName"`
	// The ID of the ExpressRoute Circuit
	ExpressRouteId *string `pulumi:"expressRouteId"`
	// The name of the private cloud.
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Authorization resource.
type AuthorizationArgs struct {
	// Name of the ExpressRoute Circuit Authorization in the private cloud
	AuthorizationName pulumi.StringPtrInput
	// The ID of the ExpressRoute Circuit
	ExpressRouteId pulumi.StringPtrInput
	// The name of the private cloud.
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (AuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationArgs)(nil)).Elem()
}

type AuthorizationInput interface {
	pulumi.Input

	ToAuthorizationOutput() AuthorizationOutput
	ToAuthorizationOutputWithContext(ctx context.Context) AuthorizationOutput
}

func (*Authorization) ElementType() reflect.Type {
	return reflect.TypeOf((**Authorization)(nil)).Elem()
}

func (i *Authorization) ToAuthorizationOutput() AuthorizationOutput {
	return i.ToAuthorizationOutputWithContext(context.Background())
}

func (i *Authorization) ToAuthorizationOutputWithContext(ctx context.Context) AuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationOutput)
}

type AuthorizationOutput struct{ *pulumi.OutputState }

func (AuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Authorization)(nil)).Elem()
}

func (o AuthorizationOutput) ToAuthorizationOutput() AuthorizationOutput {
	return o
}

func (o AuthorizationOutput) ToAuthorizationOutputWithContext(ctx context.Context) AuthorizationOutput {
	return o
}

// The ID of the ExpressRoute Circuit Authorization
func (o AuthorizationOutput) ExpressRouteAuthorizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Authorization) pulumi.StringOutput { return v.ExpressRouteAuthorizationId }).(pulumi.StringOutput)
}

// The key of the ExpressRoute Circuit Authorization
func (o AuthorizationOutput) ExpressRouteAuthorizationKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Authorization) pulumi.StringOutput { return v.ExpressRouteAuthorizationKey }).(pulumi.StringOutput)
}

// The ID of the ExpressRoute Circuit
func (o AuthorizationOutput) ExpressRouteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Authorization) pulumi.StringPtrOutput { return v.ExpressRouteId }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o AuthorizationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Authorization) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The state of the  ExpressRoute Circuit Authorization provisioning
func (o AuthorizationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Authorization) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o AuthorizationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Authorization) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationOutput{})
}
