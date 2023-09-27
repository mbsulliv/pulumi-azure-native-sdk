// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Authorization Provider contract.
type AuthorizationProvider struct {
	pulumi.CustomResourceState

	// Authorization Provider name. Must be 1 to 300 characters long.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Identity provider name. Must be 1 to 300 characters long.
	IdentityProvider pulumi.StringPtrOutput `pulumi:"identityProvider"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// OAuth2 settings
	Oauth2 AuthorizationProviderOAuth2SettingsResponsePtrOutput `pulumi:"oauth2"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAuthorizationProvider registers a new resource with the given unique name, arguments, and options.
func NewAuthorizationProvider(ctx *pulumi.Context,
	name string, args *AuthorizationProviderArgs, opts ...pulumi.ResourceOption) (*AuthorizationProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:AuthorizationProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:AuthorizationProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:AuthorizationProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:AuthorizationProvider"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AuthorizationProvider
	err := ctx.RegisterResource("azure-native:apimanagement/v20220901preview:AuthorizationProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizationProvider gets an existing AuthorizationProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizationProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationProviderState, opts ...pulumi.ResourceOption) (*AuthorizationProvider, error) {
	var resource AuthorizationProvider
	err := ctx.ReadResource("azure-native:apimanagement/v20220901preview:AuthorizationProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthorizationProvider resources.
type authorizationProviderState struct {
}

type AuthorizationProviderState struct {
}

func (AuthorizationProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationProviderState)(nil)).Elem()
}

type authorizationProviderArgs struct {
	// Identifier of the authorization provider.
	AuthorizationProviderId *string `pulumi:"authorizationProviderId"`
	// Authorization Provider name. Must be 1 to 300 characters long.
	DisplayName *string `pulumi:"displayName"`
	// Identity provider name. Must be 1 to 300 characters long.
	IdentityProvider *string `pulumi:"identityProvider"`
	// OAuth2 settings
	Oauth2 *AuthorizationProviderOAuth2Settings `pulumi:"oauth2"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a AuthorizationProvider resource.
type AuthorizationProviderArgs struct {
	// Identifier of the authorization provider.
	AuthorizationProviderId pulumi.StringPtrInput
	// Authorization Provider name. Must be 1 to 300 characters long.
	DisplayName pulumi.StringPtrInput
	// Identity provider name. Must be 1 to 300 characters long.
	IdentityProvider pulumi.StringPtrInput
	// OAuth2 settings
	Oauth2 AuthorizationProviderOAuth2SettingsPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (AuthorizationProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationProviderArgs)(nil)).Elem()
}

type AuthorizationProviderInput interface {
	pulumi.Input

	ToAuthorizationProviderOutput() AuthorizationProviderOutput
	ToAuthorizationProviderOutputWithContext(ctx context.Context) AuthorizationProviderOutput
}

func (*AuthorizationProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationProvider)(nil)).Elem()
}

func (i *AuthorizationProvider) ToAuthorizationProviderOutput() AuthorizationProviderOutput {
	return i.ToAuthorizationProviderOutputWithContext(context.Background())
}

func (i *AuthorizationProvider) ToAuthorizationProviderOutputWithContext(ctx context.Context) AuthorizationProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationProviderOutput)
}

func (i *AuthorizationProvider) ToOutput(ctx context.Context) pulumix.Output[*AuthorizationProvider] {
	return pulumix.Output[*AuthorizationProvider]{
		OutputState: i.ToAuthorizationProviderOutputWithContext(ctx).OutputState,
	}
}

type AuthorizationProviderOutput struct{ *pulumi.OutputState }

func (AuthorizationProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationProvider)(nil)).Elem()
}

func (o AuthorizationProviderOutput) ToAuthorizationProviderOutput() AuthorizationProviderOutput {
	return o
}

func (o AuthorizationProviderOutput) ToAuthorizationProviderOutputWithContext(ctx context.Context) AuthorizationProviderOutput {
	return o
}

func (o AuthorizationProviderOutput) ToOutput(ctx context.Context) pulumix.Output[*AuthorizationProvider] {
	return pulumix.Output[*AuthorizationProvider]{
		OutputState: o.OutputState,
	}
}

// Authorization Provider name. Must be 1 to 300 characters long.
func (o AuthorizationProviderOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationProvider) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Identity provider name. Must be 1 to 300 characters long.
func (o AuthorizationProviderOutput) IdentityProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationProvider) pulumi.StringPtrOutput { return v.IdentityProvider }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o AuthorizationProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// OAuth2 settings
func (o AuthorizationProviderOutput) Oauth2() AuthorizationProviderOAuth2SettingsResponsePtrOutput {
	return o.ApplyT(func(v *AuthorizationProvider) AuthorizationProviderOAuth2SettingsResponsePtrOutput { return v.Oauth2 }).(AuthorizationProviderOAuth2SettingsResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AuthorizationProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationProvider) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationProviderOutput{})
}
