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

// Azure Active Directory only authentication.
// Azure REST API version: 2021-11-01. Prior API version in Azure Native 1.x: 2020-11-01-preview
type ManagedInstanceAzureADOnlyAuthentication struct {
	pulumi.CustomResourceState

	// Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication pulumi.BoolOutput `pulumi:"azureADOnlyAuthentication"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedInstanceAzureADOnlyAuthentication registers a new resource with the given unique name, arguments, and options.
func NewManagedInstanceAzureADOnlyAuthentication(ctx *pulumi.Context,
	name string, args *ManagedInstanceAzureADOnlyAuthenticationArgs, opts ...pulumi.ResourceOption) (*ManagedInstanceAzureADOnlyAuthentication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzureADOnlyAuthentication == nil {
		return nil, errors.New("invalid value for required argument 'AzureADOnlyAuthentication'")
	}
	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:ManagedInstanceAzureADOnlyAuthentication"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedInstanceAzureADOnlyAuthentication
	err := ctx.RegisterResource("azure-native:sql:ManagedInstanceAzureADOnlyAuthentication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedInstanceAzureADOnlyAuthentication gets an existing ManagedInstanceAzureADOnlyAuthentication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedInstanceAzureADOnlyAuthentication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstanceAzureADOnlyAuthenticationState, opts ...pulumi.ResourceOption) (*ManagedInstanceAzureADOnlyAuthentication, error) {
	var resource ManagedInstanceAzureADOnlyAuthentication
	err := ctx.ReadResource("azure-native:sql:ManagedInstanceAzureADOnlyAuthentication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedInstanceAzureADOnlyAuthentication resources.
type managedInstanceAzureADOnlyAuthenticationState struct {
}

type ManagedInstanceAzureADOnlyAuthenticationState struct {
}

func (ManagedInstanceAzureADOnlyAuthenticationState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceAzureADOnlyAuthenticationState)(nil)).Elem()
}

type managedInstanceAzureADOnlyAuthenticationArgs struct {
	// The name of server azure active directory only authentication.
	AuthenticationName *string `pulumi:"authenticationName"`
	// Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication bool `pulumi:"azureADOnlyAuthentication"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagedInstanceAzureADOnlyAuthentication resource.
type ManagedInstanceAzureADOnlyAuthenticationArgs struct {
	// The name of server azure active directory only authentication.
	AuthenticationName pulumi.StringPtrInput
	// Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication pulumi.BoolInput
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
}

func (ManagedInstanceAzureADOnlyAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceAzureADOnlyAuthenticationArgs)(nil)).Elem()
}

type ManagedInstanceAzureADOnlyAuthenticationInput interface {
	pulumi.Input

	ToManagedInstanceAzureADOnlyAuthenticationOutput() ManagedInstanceAzureADOnlyAuthenticationOutput
	ToManagedInstanceAzureADOnlyAuthenticationOutputWithContext(ctx context.Context) ManagedInstanceAzureADOnlyAuthenticationOutput
}

func (*ManagedInstanceAzureADOnlyAuthentication) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceAzureADOnlyAuthentication)(nil)).Elem()
}

func (i *ManagedInstanceAzureADOnlyAuthentication) ToManagedInstanceAzureADOnlyAuthenticationOutput() ManagedInstanceAzureADOnlyAuthenticationOutput {
	return i.ToManagedInstanceAzureADOnlyAuthenticationOutputWithContext(context.Background())
}

func (i *ManagedInstanceAzureADOnlyAuthentication) ToManagedInstanceAzureADOnlyAuthenticationOutputWithContext(ctx context.Context) ManagedInstanceAzureADOnlyAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceAzureADOnlyAuthenticationOutput)
}

func (i *ManagedInstanceAzureADOnlyAuthentication) ToOutput(ctx context.Context) pulumix.Output[*ManagedInstanceAzureADOnlyAuthentication] {
	return pulumix.Output[*ManagedInstanceAzureADOnlyAuthentication]{
		OutputState: i.ToManagedInstanceAzureADOnlyAuthenticationOutputWithContext(ctx).OutputState,
	}
}

type ManagedInstanceAzureADOnlyAuthenticationOutput struct{ *pulumi.OutputState }

func (ManagedInstanceAzureADOnlyAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceAzureADOnlyAuthentication)(nil)).Elem()
}

func (o ManagedInstanceAzureADOnlyAuthenticationOutput) ToManagedInstanceAzureADOnlyAuthenticationOutput() ManagedInstanceAzureADOnlyAuthenticationOutput {
	return o
}

func (o ManagedInstanceAzureADOnlyAuthenticationOutput) ToManagedInstanceAzureADOnlyAuthenticationOutputWithContext(ctx context.Context) ManagedInstanceAzureADOnlyAuthenticationOutput {
	return o
}

func (o ManagedInstanceAzureADOnlyAuthenticationOutput) ToOutput(ctx context.Context) pulumix.Output[*ManagedInstanceAzureADOnlyAuthentication] {
	return pulumix.Output[*ManagedInstanceAzureADOnlyAuthentication]{
		OutputState: o.OutputState,
	}
}

// Azure Active Directory only Authentication enabled.
func (o ManagedInstanceAzureADOnlyAuthenticationOutput) AzureADOnlyAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v *ManagedInstanceAzureADOnlyAuthentication) pulumi.BoolOutput {
		return v.AzureADOnlyAuthentication
	}).(pulumi.BoolOutput)
}

// Resource name.
func (o ManagedInstanceAzureADOnlyAuthenticationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceAzureADOnlyAuthentication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o ManagedInstanceAzureADOnlyAuthenticationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceAzureADOnlyAuthentication) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedInstanceAzureADOnlyAuthenticationOutput{})
}
