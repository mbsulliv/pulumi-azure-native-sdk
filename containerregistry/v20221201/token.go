// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An object that represents a token for a container registry.
type Token struct {
	pulumi.CustomResourceState

	// The creation date of scope map.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The credentials that can be used for authenticating the token.
	Credentials TokenCredentialsPropertiesResponsePtrOutput `pulumi:"credentials"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource ID of the scope map to which the token will be associated with.
	ScopeMapId pulumi.StringPtrOutput `pulumi:"scopeMapId"`
	// The status of the token example enabled or disabled.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewToken registers a new resource with the given unique name, arguments, and options.
func NewToken(ctx *pulumi.Context,
	name string, args *TokenArgs, opts ...pulumi.ResourceOption) (*Token, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20190501preview:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20201101preview:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20211201preview:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20220201preview:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230101preview:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230601preview:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230701:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230801preview:Token"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20231101preview:Token"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Token
	err := ctx.RegisterResource("azure-native:containerregistry/v20221201:Token", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetToken gets an existing Token resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TokenState, opts ...pulumi.ResourceOption) (*Token, error) {
	var resource Token
	err := ctx.ReadResource("azure-native:containerregistry/v20221201:Token", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Token resources.
type tokenState struct {
}

type TokenState struct {
}

func (TokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenState)(nil)).Elem()
}

type tokenArgs struct {
	// The credentials that can be used for authenticating the token.
	Credentials *TokenCredentialsProperties `pulumi:"credentials"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource ID of the scope map to which the token will be associated with.
	ScopeMapId *string `pulumi:"scopeMapId"`
	// The status of the token example enabled or disabled.
	Status *string `pulumi:"status"`
	// The name of the token.
	TokenName *string `pulumi:"tokenName"`
}

// The set of arguments for constructing a Token resource.
type TokenArgs struct {
	// The credentials that can be used for authenticating the token.
	Credentials TokenCredentialsPropertiesPtrInput
	// The name of the container registry.
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The resource ID of the scope map to which the token will be associated with.
	ScopeMapId pulumi.StringPtrInput
	// The status of the token example enabled or disabled.
	Status pulumi.StringPtrInput
	// The name of the token.
	TokenName pulumi.StringPtrInput
}

func (TokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenArgs)(nil)).Elem()
}

type TokenInput interface {
	pulumi.Input

	ToTokenOutput() TokenOutput
	ToTokenOutputWithContext(ctx context.Context) TokenOutput
}

func (*Token) ElementType() reflect.Type {
	return reflect.TypeOf((**Token)(nil)).Elem()
}

func (i *Token) ToTokenOutput() TokenOutput {
	return i.ToTokenOutputWithContext(context.Background())
}

func (i *Token) ToTokenOutputWithContext(ctx context.Context) TokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenOutput)
}

type TokenOutput struct{ *pulumi.OutputState }

func (TokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Token)(nil)).Elem()
}

func (o TokenOutput) ToTokenOutput() TokenOutput {
	return o
}

func (o TokenOutput) ToTokenOutputWithContext(ctx context.Context) TokenOutput {
	return o
}

// The creation date of scope map.
func (o TokenOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Token) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The credentials that can be used for authenticating the token.
func (o TokenOutput) Credentials() TokenCredentialsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Token) TokenCredentialsPropertiesResponsePtrOutput { return v.Credentials }).(TokenCredentialsPropertiesResponsePtrOutput)
}

// The name of the resource.
func (o TokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Token) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o TokenOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Token) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource ID of the scope map to which the token will be associated with.
func (o TokenOutput) ScopeMapId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.StringPtrOutput { return v.ScopeMapId }).(pulumi.StringPtrOutput)
}

// The status of the token example enabled or disabled.
func (o TokenOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o TokenOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Token) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o TokenOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Token) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TokenOutput{})
}
