// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20140401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An server Active Directory Administrator.
type ServerAzureADAdministrator struct {
	pulumi.CustomResourceState

	// The type of administrator.
	AdministratorType pulumi.StringOutput `pulumi:"administratorType"`
	// The server administrator login value.
	Login pulumi.StringOutput `pulumi:"login"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The server administrator Sid (Secure ID).
	Sid pulumi.StringOutput `pulumi:"sid"`
	// The server Active Directory Administrator tenant id.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServerAzureADAdministrator registers a new resource with the given unique name, arguments, and options.
func NewServerAzureADAdministrator(ctx *pulumi.Context,
	name string, args *ServerAzureADAdministratorArgs, opts ...pulumi.ResourceOption) (*ServerAzureADAdministrator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdministratorType == nil {
		return nil, errors.New("invalid value for required argument 'AdministratorType'")
	}
	if args.Login == nil {
		return nil, errors.New("invalid value for required argument 'Login'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.Sid == nil {
		return nil, errors.New("invalid value for required argument 'Sid'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:ServerAzureADAdministrator"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServerAzureADAdministrator
	err := ctx.RegisterResource("azure-native:sql/v20140401:ServerAzureADAdministrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerAzureADAdministrator gets an existing ServerAzureADAdministrator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerAzureADAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerAzureADAdministratorState, opts ...pulumi.ResourceOption) (*ServerAzureADAdministrator, error) {
	var resource ServerAzureADAdministrator
	err := ctx.ReadResource("azure-native:sql/v20140401:ServerAzureADAdministrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerAzureADAdministrator resources.
type serverAzureADAdministratorState struct {
}

type ServerAzureADAdministratorState struct {
}

func (ServerAzureADAdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAzureADAdministratorState)(nil)).Elem()
}

type serverAzureADAdministratorArgs struct {
	// Name of the server administrator resource.
	AdministratorName *string `pulumi:"administratorName"`
	// The type of administrator.
	AdministratorType string `pulumi:"administratorType"`
	// The server administrator login value.
	Login string `pulumi:"login"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The server administrator Sid (Secure ID).
	Sid string `pulumi:"sid"`
	// The server Active Directory Administrator tenant id.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a ServerAzureADAdministrator resource.
type ServerAzureADAdministratorArgs struct {
	// Name of the server administrator resource.
	AdministratorName pulumi.StringPtrInput
	// The type of administrator.
	AdministratorType pulumi.StringInput
	// The server administrator login value.
	Login pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The server administrator Sid (Secure ID).
	Sid pulumi.StringInput
	// The server Active Directory Administrator tenant id.
	TenantId pulumi.StringInput
}

func (ServerAzureADAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAzureADAdministratorArgs)(nil)).Elem()
}

type ServerAzureADAdministratorInput interface {
	pulumi.Input

	ToServerAzureADAdministratorOutput() ServerAzureADAdministratorOutput
	ToServerAzureADAdministratorOutputWithContext(ctx context.Context) ServerAzureADAdministratorOutput
}

func (*ServerAzureADAdministrator) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAzureADAdministrator)(nil)).Elem()
}

func (i *ServerAzureADAdministrator) ToServerAzureADAdministratorOutput() ServerAzureADAdministratorOutput {
	return i.ToServerAzureADAdministratorOutputWithContext(context.Background())
}

func (i *ServerAzureADAdministrator) ToServerAzureADAdministratorOutputWithContext(ctx context.Context) ServerAzureADAdministratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAzureADAdministratorOutput)
}

type ServerAzureADAdministratorOutput struct{ *pulumi.OutputState }

func (ServerAzureADAdministratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAzureADAdministrator)(nil)).Elem()
}

func (o ServerAzureADAdministratorOutput) ToServerAzureADAdministratorOutput() ServerAzureADAdministratorOutput {
	return o
}

func (o ServerAzureADAdministratorOutput) ToServerAzureADAdministratorOutputWithContext(ctx context.Context) ServerAzureADAdministratorOutput {
	return o
}

// The type of administrator.
func (o ServerAzureADAdministratorOutput) AdministratorType() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADAdministrator) pulumi.StringOutput { return v.AdministratorType }).(pulumi.StringOutput)
}

// The server administrator login value.
func (o ServerAzureADAdministratorOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADAdministrator) pulumi.StringOutput { return v.Login }).(pulumi.StringOutput)
}

// Resource name.
func (o ServerAzureADAdministratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADAdministrator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The server administrator Sid (Secure ID).
func (o ServerAzureADAdministratorOutput) Sid() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADAdministrator) pulumi.StringOutput { return v.Sid }).(pulumi.StringOutput)
}

// The server Active Directory Administrator tenant id.
func (o ServerAzureADAdministratorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADAdministrator) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type.
func (o ServerAzureADAdministratorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADAdministrator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerAzureADAdministratorOutput{})
}
