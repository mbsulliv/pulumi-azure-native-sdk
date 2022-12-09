// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configuration settings for the Azure ContainerApp Service Authentication / Authorization feature.
type ContainerAppsAuthConfig struct {
	pulumi.CustomResourceState

	// The configuration settings that determines the validation flow of users using  Service Authentication/Authorization.
	GlobalValidation GlobalValidationResponsePtrOutput `pulumi:"globalValidation"`
	// The configuration settings of the HTTP requests for authentication and authorization requests made against ContainerApp Service Authentication/Authorization.
	HttpSettings HttpSettingsResponsePtrOutput `pulumi:"httpSettings"`
	// The configuration settings of each of the identity providers used to configure ContainerApp Service Authentication/Authorization.
	IdentityProviders IdentityProvidersResponsePtrOutput `pulumi:"identityProviders"`
	// The configuration settings of the login flow of users using ContainerApp Service Authentication/Authorization.
	Login LoginResponsePtrOutput `pulumi:"login"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The configuration settings of the platform of ContainerApp Service Authentication/Authorization.
	Platform AuthPlatformResponsePtrOutput `pulumi:"platform"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewContainerAppsAuthConfig registers a new resource with the given unique name, arguments, and options.
func NewContainerAppsAuthConfig(ctx *pulumi.Context,
	name string, args *ContainerAppsAuthConfigArgs, opts ...pulumi.ResourceOption) (*ContainerAppsAuthConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerAppName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerAppName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app:ContainerAppsAuthConfig"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220101preview:ContainerAppsAuthConfig"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220301:ContainerAppsAuthConfig"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220601preview:ContainerAppsAuthConfig"),
		},
	})
	opts = append(opts, aliases)
	var resource ContainerAppsAuthConfig
	err := ctx.RegisterResource("azure-native:app/v20221001:ContainerAppsAuthConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerAppsAuthConfig gets an existing ContainerAppsAuthConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerAppsAuthConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerAppsAuthConfigState, opts ...pulumi.ResourceOption) (*ContainerAppsAuthConfig, error) {
	var resource ContainerAppsAuthConfig
	err := ctx.ReadResource("azure-native:app/v20221001:ContainerAppsAuthConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerAppsAuthConfig resources.
type containerAppsAuthConfigState struct {
}

type ContainerAppsAuthConfigState struct {
}

func (ContainerAppsAuthConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppsAuthConfigState)(nil)).Elem()
}

type containerAppsAuthConfigArgs struct {
	// Name of the Container App AuthConfig.
	AuthConfigName *string `pulumi:"authConfigName"`
	// Name of the Container App.
	ContainerAppName string `pulumi:"containerAppName"`
	// The configuration settings that determines the validation flow of users using  Service Authentication/Authorization.
	GlobalValidation *GlobalValidation `pulumi:"globalValidation"`
	// The configuration settings of the HTTP requests for authentication and authorization requests made against ContainerApp Service Authentication/Authorization.
	HttpSettings *HttpSettings `pulumi:"httpSettings"`
	// The configuration settings of each of the identity providers used to configure ContainerApp Service Authentication/Authorization.
	IdentityProviders *IdentityProviders `pulumi:"identityProviders"`
	// The configuration settings of the login flow of users using ContainerApp Service Authentication/Authorization.
	Login *Login `pulumi:"login"`
	// The configuration settings of the platform of ContainerApp Service Authentication/Authorization.
	Platform *AuthPlatform `pulumi:"platform"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ContainerAppsAuthConfig resource.
type ContainerAppsAuthConfigArgs struct {
	// Name of the Container App AuthConfig.
	AuthConfigName pulumi.StringPtrInput
	// Name of the Container App.
	ContainerAppName pulumi.StringInput
	// The configuration settings that determines the validation flow of users using  Service Authentication/Authorization.
	GlobalValidation GlobalValidationPtrInput
	// The configuration settings of the HTTP requests for authentication and authorization requests made against ContainerApp Service Authentication/Authorization.
	HttpSettings HttpSettingsPtrInput
	// The configuration settings of each of the identity providers used to configure ContainerApp Service Authentication/Authorization.
	IdentityProviders IdentityProvidersPtrInput
	// The configuration settings of the login flow of users using ContainerApp Service Authentication/Authorization.
	Login LoginPtrInput
	// The configuration settings of the platform of ContainerApp Service Authentication/Authorization.
	Platform AuthPlatformPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ContainerAppsAuthConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppsAuthConfigArgs)(nil)).Elem()
}

type ContainerAppsAuthConfigInput interface {
	pulumi.Input

	ToContainerAppsAuthConfigOutput() ContainerAppsAuthConfigOutput
	ToContainerAppsAuthConfigOutputWithContext(ctx context.Context) ContainerAppsAuthConfigOutput
}

func (*ContainerAppsAuthConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppsAuthConfig)(nil)).Elem()
}

func (i *ContainerAppsAuthConfig) ToContainerAppsAuthConfigOutput() ContainerAppsAuthConfigOutput {
	return i.ToContainerAppsAuthConfigOutputWithContext(context.Background())
}

func (i *ContainerAppsAuthConfig) ToContainerAppsAuthConfigOutputWithContext(ctx context.Context) ContainerAppsAuthConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppsAuthConfigOutput)
}

type ContainerAppsAuthConfigOutput struct{ *pulumi.OutputState }

func (ContainerAppsAuthConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppsAuthConfig)(nil)).Elem()
}

func (o ContainerAppsAuthConfigOutput) ToContainerAppsAuthConfigOutput() ContainerAppsAuthConfigOutput {
	return o
}

func (o ContainerAppsAuthConfigOutput) ToContainerAppsAuthConfigOutputWithContext(ctx context.Context) ContainerAppsAuthConfigOutput {
	return o
}

// The configuration settings that determines the validation flow of users using  Service Authentication/Authorization.
func (o ContainerAppsAuthConfigOutput) GlobalValidation() GlobalValidationResponsePtrOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) GlobalValidationResponsePtrOutput { return v.GlobalValidation }).(GlobalValidationResponsePtrOutput)
}

// The configuration settings of the HTTP requests for authentication and authorization requests made against ContainerApp Service Authentication/Authorization.
func (o ContainerAppsAuthConfigOutput) HttpSettings() HttpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) HttpSettingsResponsePtrOutput { return v.HttpSettings }).(HttpSettingsResponsePtrOutput)
}

// The configuration settings of each of the identity providers used to configure ContainerApp Service Authentication/Authorization.
func (o ContainerAppsAuthConfigOutput) IdentityProviders() IdentityProvidersResponsePtrOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) IdentityProvidersResponsePtrOutput { return v.IdentityProviders }).(IdentityProvidersResponsePtrOutput)
}

// The configuration settings of the login flow of users using ContainerApp Service Authentication/Authorization.
func (o ContainerAppsAuthConfigOutput) Login() LoginResponsePtrOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) LoginResponsePtrOutput { return v.Login }).(LoginResponsePtrOutput)
}

// The name of the resource
func (o ContainerAppsAuthConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The configuration settings of the platform of ContainerApp Service Authentication/Authorization.
func (o ContainerAppsAuthConfigOutput) Platform() AuthPlatformResponsePtrOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) AuthPlatformResponsePtrOutput { return v.Platform }).(AuthPlatformResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ContainerAppsAuthConfigOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ContainerAppsAuthConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerAppsAuthConfigOutput{})
}
