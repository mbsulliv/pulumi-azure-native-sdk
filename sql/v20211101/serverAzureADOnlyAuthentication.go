// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Active Directory only authentication.
type ServerAzureADOnlyAuthentication struct {
	pulumi.CustomResourceState

	// Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication pulumi.BoolOutput `pulumi:"azureADOnlyAuthentication"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServerAzureADOnlyAuthentication registers a new resource with the given unique name, arguments, and options.
func NewServerAzureADOnlyAuthentication(ctx *pulumi.Context,
	name string, args *ServerAzureADOnlyAuthenticationArgs, opts ...pulumi.ResourceOption) (*ServerAzureADOnlyAuthentication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzureADOnlyAuthentication == nil {
		return nil, errors.New("invalid value for required argument 'AzureADOnlyAuthentication'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230801preview:ServerAzureADOnlyAuthentication"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServerAzureADOnlyAuthentication
	err := ctx.RegisterResource("azure-native:sql/v20211101:ServerAzureADOnlyAuthentication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerAzureADOnlyAuthentication gets an existing ServerAzureADOnlyAuthentication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerAzureADOnlyAuthentication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerAzureADOnlyAuthenticationState, opts ...pulumi.ResourceOption) (*ServerAzureADOnlyAuthentication, error) {
	var resource ServerAzureADOnlyAuthentication
	err := ctx.ReadResource("azure-native:sql/v20211101:ServerAzureADOnlyAuthentication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerAzureADOnlyAuthentication resources.
type serverAzureADOnlyAuthenticationState struct {
}

type ServerAzureADOnlyAuthenticationState struct {
}

func (ServerAzureADOnlyAuthenticationState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAzureADOnlyAuthenticationState)(nil)).Elem()
}

type serverAzureADOnlyAuthenticationArgs struct {
	// The name of server azure active directory only authentication.
	AuthenticationName *string `pulumi:"authenticationName"`
	// Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication bool `pulumi:"azureADOnlyAuthentication"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// The set of arguments for constructing a ServerAzureADOnlyAuthentication resource.
type ServerAzureADOnlyAuthenticationArgs struct {
	// The name of server azure active directory only authentication.
	AuthenticationName pulumi.StringPtrInput
	// Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication pulumi.BoolInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
}

func (ServerAzureADOnlyAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAzureADOnlyAuthenticationArgs)(nil)).Elem()
}

type ServerAzureADOnlyAuthenticationInput interface {
	pulumi.Input

	ToServerAzureADOnlyAuthenticationOutput() ServerAzureADOnlyAuthenticationOutput
	ToServerAzureADOnlyAuthenticationOutputWithContext(ctx context.Context) ServerAzureADOnlyAuthenticationOutput
}

func (*ServerAzureADOnlyAuthentication) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAzureADOnlyAuthentication)(nil)).Elem()
}

func (i *ServerAzureADOnlyAuthentication) ToServerAzureADOnlyAuthenticationOutput() ServerAzureADOnlyAuthenticationOutput {
	return i.ToServerAzureADOnlyAuthenticationOutputWithContext(context.Background())
}

func (i *ServerAzureADOnlyAuthentication) ToServerAzureADOnlyAuthenticationOutputWithContext(ctx context.Context) ServerAzureADOnlyAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAzureADOnlyAuthenticationOutput)
}

type ServerAzureADOnlyAuthenticationOutput struct{ *pulumi.OutputState }

func (ServerAzureADOnlyAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAzureADOnlyAuthentication)(nil)).Elem()
}

func (o ServerAzureADOnlyAuthenticationOutput) ToServerAzureADOnlyAuthenticationOutput() ServerAzureADOnlyAuthenticationOutput {
	return o
}

func (o ServerAzureADOnlyAuthenticationOutput) ToServerAzureADOnlyAuthenticationOutputWithContext(ctx context.Context) ServerAzureADOnlyAuthenticationOutput {
	return o
}

// Azure Active Directory only Authentication enabled.
func (o ServerAzureADOnlyAuthenticationOutput) AzureADOnlyAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServerAzureADOnlyAuthentication) pulumi.BoolOutput { return v.AzureADOnlyAuthentication }).(pulumi.BoolOutput)
}

// Resource name.
func (o ServerAzureADOnlyAuthenticationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADOnlyAuthentication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o ServerAzureADOnlyAuthenticationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADOnlyAuthentication) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerAzureADOnlyAuthenticationOutput{})
}
