// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Active Directory administrator.
// Azure REST API version: 2021-11-01. Prior API version in Azure Native 1.x: 2020-11-01-preview.
//
// Other available API versions: 2014-04-01, 2022-11-01-preview, 2023-02-01-preview.
type ServerAzureADAdministrator struct {
	pulumi.CustomResourceState

	// Type of the sever administrator.
	AdministratorType pulumi.StringOutput `pulumi:"administratorType"`
	// Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication pulumi.BoolOutput `pulumi:"azureADOnlyAuthentication"`
	// Login name of the server administrator.
	Login pulumi.StringOutput `pulumi:"login"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// SID (object ID) of the server administrator.
	Sid pulumi.StringOutput `pulumi:"sid"`
	// Tenant ID of the administrator.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql/v20140401:ServerAzureADAdministrator"),
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
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServerAzureADAdministrator
	err := ctx.RegisterResource("azure-native:sql:ServerAzureADAdministrator", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:sql:ServerAzureADAdministrator", name, id, state, &resource, opts...)
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
	// The name of server active directory administrator.
	AdministratorName *string `pulumi:"administratorName"`
	// Type of the sever administrator.
	AdministratorType string `pulumi:"administratorType"`
	// Login name of the server administrator.
	Login string `pulumi:"login"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// SID (object ID) of the server administrator.
	Sid string `pulumi:"sid"`
	// Tenant ID of the administrator.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a ServerAzureADAdministrator resource.
type ServerAzureADAdministratorArgs struct {
	// The name of server active directory administrator.
	AdministratorName pulumi.StringPtrInput
	// Type of the sever administrator.
	AdministratorType pulumi.StringInput
	// Login name of the server administrator.
	Login pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// SID (object ID) of the server administrator.
	Sid pulumi.StringInput
	// Tenant ID of the administrator.
	TenantId pulumi.StringPtrInput
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

func (i *ServerAzureADAdministrator) ToOutput(ctx context.Context) pulumix.Output[*ServerAzureADAdministrator] {
	return pulumix.Output[*ServerAzureADAdministrator]{
		OutputState: i.ToServerAzureADAdministratorOutputWithContext(ctx).OutputState,
	}
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

func (o ServerAzureADAdministratorOutput) ToOutput(ctx context.Context) pulumix.Output[*ServerAzureADAdministrator] {
	return pulumix.Output[*ServerAzureADAdministrator]{
		OutputState: o.OutputState,
	}
}

// Type of the sever administrator.
func (o ServerAzureADAdministratorOutput) AdministratorType() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADAdministrator) pulumi.StringOutput { return v.AdministratorType }).(pulumi.StringOutput)
}

// Azure Active Directory only Authentication enabled.
func (o ServerAzureADAdministratorOutput) AzureADOnlyAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServerAzureADAdministrator) pulumi.BoolOutput { return v.AzureADOnlyAuthentication }).(pulumi.BoolOutput)
}

// Login name of the server administrator.
func (o ServerAzureADAdministratorOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADAdministrator) pulumi.StringOutput { return v.Login }).(pulumi.StringOutput)
}

// Resource name.
func (o ServerAzureADAdministratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADAdministrator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SID (object ID) of the server administrator.
func (o ServerAzureADAdministratorOutput) Sid() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADAdministrator) pulumi.StringOutput { return v.Sid }).(pulumi.StringOutput)
}

// Tenant ID of the administrator.
func (o ServerAzureADAdministratorOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerAzureADAdministrator) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o ServerAzureADAdministratorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADAdministrator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerAzureADAdministratorOutput{})
}
