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

// Identity Provider details.
type IdentityProvider struct {
	pulumi.CustomResourceState

	// List of Allowed Tenants when configuring Azure Active Directory login.
	AllowedTenants pulumi.StringArrayOutput `pulumi:"allowedTenants"`
	// OpenID Connect discovery endpoint hostname for AAD or AAD B2C.
	Authority pulumi.StringPtrOutput `pulumi:"authority"`
	// Client Id of the Application in the external Identity Provider. It is App ID for Facebook login, Client ID for Google login, App ID for Microsoft.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The client library to be used in the developer portal. Only applies to AAD and AAD B2C Identity Provider.
	ClientLibrary pulumi.StringPtrOutput `pulumi:"clientLibrary"`
	// Client secret of the Application in external Identity Provider, used to authenticate login request. For example, it is App Secret for Facebook login, API Key for Google login, Public Key for Microsoft. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Password Reset Policy Name. Only applies to AAD B2C Identity Provider.
	PasswordResetPolicyName pulumi.StringPtrOutput `pulumi:"passwordResetPolicyName"`
	// Profile Editing Policy Name. Only applies to AAD B2C Identity Provider.
	ProfileEditingPolicyName pulumi.StringPtrOutput `pulumi:"profileEditingPolicyName"`
	// Signin Policy Name. Only applies to AAD B2C Identity Provider.
	SigninPolicyName pulumi.StringPtrOutput `pulumi:"signinPolicyName"`
	// The TenantId to use instead of Common when logging into Active Directory
	SigninTenant pulumi.StringPtrOutput `pulumi:"signinTenant"`
	// Signup Policy Name. Only applies to AAD B2C Identity Provider.
	SignupPolicyName pulumi.StringPtrOutput `pulumi:"signupPolicyName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIdentityProvider registers a new resource with the given unique name, arguments, and options.
func NewIdentityProvider(ctx *pulumi.Context,
	name string, args *IdentityProviderArgs, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:IdentityProvider"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IdentityProvider
	err := ctx.RegisterResource("azure-native:apimanagement/v20220901preview:IdentityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityProvider gets an existing IdentityProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityProviderState, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	var resource IdentityProvider
	err := ctx.ReadResource("azure-native:apimanagement/v20220901preview:IdentityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityProvider resources.
type identityProviderState struct {
}

type IdentityProviderState struct {
}

func (IdentityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderState)(nil)).Elem()
}

type identityProviderArgs struct {
	// List of Allowed Tenants when configuring Azure Active Directory login.
	AllowedTenants []string `pulumi:"allowedTenants"`
	// OpenID Connect discovery endpoint hostname for AAD or AAD B2C.
	Authority *string `pulumi:"authority"`
	// Client Id of the Application in the external Identity Provider. It is App ID for Facebook login, Client ID for Google login, App ID for Microsoft.
	ClientId string `pulumi:"clientId"`
	// The client library to be used in the developer portal. Only applies to AAD and AAD B2C Identity Provider.
	ClientLibrary *string `pulumi:"clientLibrary"`
	// Client secret of the Application in external Identity Provider, used to authenticate login request. For example, it is App Secret for Facebook login, API Key for Google login, Public Key for Microsoft. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	ClientSecret string `pulumi:"clientSecret"`
	// Identity Provider Type identifier.
	IdentityProviderName *string `pulumi:"identityProviderName"`
	// Password Reset Policy Name. Only applies to AAD B2C Identity Provider.
	PasswordResetPolicyName *string `pulumi:"passwordResetPolicyName"`
	// Profile Editing Policy Name. Only applies to AAD B2C Identity Provider.
	ProfileEditingPolicyName *string `pulumi:"profileEditingPolicyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Signin Policy Name. Only applies to AAD B2C Identity Provider.
	SigninPolicyName *string `pulumi:"signinPolicyName"`
	// The TenantId to use instead of Common when logging into Active Directory
	SigninTenant *string `pulumi:"signinTenant"`
	// Signup Policy Name. Only applies to AAD B2C Identity Provider.
	SignupPolicyName *string `pulumi:"signupPolicyName"`
	// Identity Provider Type identifier.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a IdentityProvider resource.
type IdentityProviderArgs struct {
	// List of Allowed Tenants when configuring Azure Active Directory login.
	AllowedTenants pulumi.StringArrayInput
	// OpenID Connect discovery endpoint hostname for AAD or AAD B2C.
	Authority pulumi.StringPtrInput
	// Client Id of the Application in the external Identity Provider. It is App ID for Facebook login, Client ID for Google login, App ID for Microsoft.
	ClientId pulumi.StringInput
	// The client library to be used in the developer portal. Only applies to AAD and AAD B2C Identity Provider.
	ClientLibrary pulumi.StringPtrInput
	// Client secret of the Application in external Identity Provider, used to authenticate login request. For example, it is App Secret for Facebook login, API Key for Google login, Public Key for Microsoft. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	ClientSecret pulumi.StringInput
	// Identity Provider Type identifier.
	IdentityProviderName pulumi.StringPtrInput
	// Password Reset Policy Name. Only applies to AAD B2C Identity Provider.
	PasswordResetPolicyName pulumi.StringPtrInput
	// Profile Editing Policy Name. Only applies to AAD B2C Identity Provider.
	ProfileEditingPolicyName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Signin Policy Name. Only applies to AAD B2C Identity Provider.
	SigninPolicyName pulumi.StringPtrInput
	// The TenantId to use instead of Common when logging into Active Directory
	SigninTenant pulumi.StringPtrInput
	// Signup Policy Name. Only applies to AAD B2C Identity Provider.
	SignupPolicyName pulumi.StringPtrInput
	// Identity Provider Type identifier.
	Type pulumi.StringPtrInput
}

func (IdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderArgs)(nil)).Elem()
}

type IdentityProviderInput interface {
	pulumi.Input

	ToIdentityProviderOutput() IdentityProviderOutput
	ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput
}

func (*IdentityProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvider)(nil)).Elem()
}

func (i *IdentityProvider) ToIdentityProviderOutput() IdentityProviderOutput {
	return i.ToIdentityProviderOutputWithContext(context.Background())
}

func (i *IdentityProvider) ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderOutput)
}

func (i *IdentityProvider) ToOutput(ctx context.Context) pulumix.Output[*IdentityProvider] {
	return pulumix.Output[*IdentityProvider]{
		OutputState: i.ToIdentityProviderOutputWithContext(ctx).OutputState,
	}
}

type IdentityProviderOutput struct{ *pulumi.OutputState }

func (IdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvider)(nil)).Elem()
}

func (o IdentityProviderOutput) ToIdentityProviderOutput() IdentityProviderOutput {
	return o
}

func (o IdentityProviderOutput) ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput {
	return o
}

func (o IdentityProviderOutput) ToOutput(ctx context.Context) pulumix.Output[*IdentityProvider] {
	return pulumix.Output[*IdentityProvider]{
		OutputState: o.OutputState,
	}
}

// List of Allowed Tenants when configuring Azure Active Directory login.
func (o IdentityProviderOutput) AllowedTenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringArrayOutput { return v.AllowedTenants }).(pulumi.StringArrayOutput)
}

// OpenID Connect discovery endpoint hostname for AAD or AAD B2C.
func (o IdentityProviderOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.Authority }).(pulumi.StringPtrOutput)
}

// Client Id of the Application in the external Identity Provider. It is App ID for Facebook login, Client ID for Google login, App ID for Microsoft.
func (o IdentityProviderOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// The client library to be used in the developer portal. Only applies to AAD and AAD B2C Identity Provider.
func (o IdentityProviderOutput) ClientLibrary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.ClientLibrary }).(pulumi.StringPtrOutput)
}

// Client secret of the Application in external Identity Provider, used to authenticate login request. For example, it is App Secret for Facebook login, API Key for Google login, Public Key for Microsoft. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
func (o IdentityProviderOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o IdentityProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password Reset Policy Name. Only applies to AAD B2C Identity Provider.
func (o IdentityProviderOutput) PasswordResetPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.PasswordResetPolicyName }).(pulumi.StringPtrOutput)
}

// Profile Editing Policy Name. Only applies to AAD B2C Identity Provider.
func (o IdentityProviderOutput) ProfileEditingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.ProfileEditingPolicyName }).(pulumi.StringPtrOutput)
}

// Signin Policy Name. Only applies to AAD B2C Identity Provider.
func (o IdentityProviderOutput) SigninPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.SigninPolicyName }).(pulumi.StringPtrOutput)
}

// The TenantId to use instead of Common when logging into Active Directory
func (o IdentityProviderOutput) SigninTenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.SigninTenant }).(pulumi.StringPtrOutput)
}

// Signup Policy Name. Only applies to AAD B2C Identity Provider.
func (o IdentityProviderOutput) SignupPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.SignupPolicyName }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o IdentityProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityProviderOutput{})
}
