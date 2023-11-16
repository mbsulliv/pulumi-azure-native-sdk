// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Administrator.
type AzureADAdministrator struct {
	pulumi.CustomResourceState

	// Type of the sever administrator.
	AdministratorType pulumi.StringPtrOutput `pulumi:"administratorType"`
	// The resource id of the identity used for AAD Authentication.
	IdentityResourceId pulumi.StringPtrOutput `pulumi:"identityResourceId"`
	// Login name of the server administrator.
	Login pulumi.StringPtrOutput `pulumi:"login"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// SID (object ID) of the server administrator.
	Sid pulumi.StringPtrOutput `pulumi:"sid"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tenant ID of the administrator.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAzureADAdministrator registers a new resource with the given unique name, arguments, and options.
func NewAzureADAdministrator(ctx *pulumi.Context,
	name string, args *AzureADAdministratorArgs, opts ...pulumi.ResourceOption) (*AzureADAdministrator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dbformysql:AzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20211201preview:AzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20230601preview:AzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20230630:AzureADAdministrator"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AzureADAdministrator
	err := ctx.RegisterResource("azure-native:dbformysql/v20220101:AzureADAdministrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureADAdministrator gets an existing AzureADAdministrator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureADAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureADAdministratorState, opts ...pulumi.ResourceOption) (*AzureADAdministrator, error) {
	var resource AzureADAdministrator
	err := ctx.ReadResource("azure-native:dbformysql/v20220101:AzureADAdministrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureADAdministrator resources.
type azureADAdministratorState struct {
}

type AzureADAdministratorState struct {
}

func (AzureADAdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureADAdministratorState)(nil)).Elem()
}

type azureADAdministratorArgs struct {
	// The name of the Azure AD Administrator.
	AdministratorName *string `pulumi:"administratorName"`
	// Type of the sever administrator.
	AdministratorType *string `pulumi:"administratorType"`
	// The resource id of the identity used for AAD Authentication.
	IdentityResourceId *string `pulumi:"identityResourceId"`
	// Login name of the server administrator.
	Login *string `pulumi:"login"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// SID (object ID) of the server administrator.
	Sid *string `pulumi:"sid"`
	// Tenant ID of the administrator.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a AzureADAdministrator resource.
type AzureADAdministratorArgs struct {
	// The name of the Azure AD Administrator.
	AdministratorName pulumi.StringPtrInput
	// Type of the sever administrator.
	AdministratorType pulumi.StringPtrInput
	// The resource id of the identity used for AAD Authentication.
	IdentityResourceId pulumi.StringPtrInput
	// Login name of the server administrator.
	Login pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// SID (object ID) of the server administrator.
	Sid pulumi.StringPtrInput
	// Tenant ID of the administrator.
	TenantId pulumi.StringPtrInput
}

func (AzureADAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureADAdministratorArgs)(nil)).Elem()
}

type AzureADAdministratorInput interface {
	pulumi.Input

	ToAzureADAdministratorOutput() AzureADAdministratorOutput
	ToAzureADAdministratorOutputWithContext(ctx context.Context) AzureADAdministratorOutput
}

func (*AzureADAdministrator) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureADAdministrator)(nil)).Elem()
}

func (i *AzureADAdministrator) ToAzureADAdministratorOutput() AzureADAdministratorOutput {
	return i.ToAzureADAdministratorOutputWithContext(context.Background())
}

func (i *AzureADAdministrator) ToAzureADAdministratorOutputWithContext(ctx context.Context) AzureADAdministratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureADAdministratorOutput)
}

type AzureADAdministratorOutput struct{ *pulumi.OutputState }

func (AzureADAdministratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureADAdministrator)(nil)).Elem()
}

func (o AzureADAdministratorOutput) ToAzureADAdministratorOutput() AzureADAdministratorOutput {
	return o
}

func (o AzureADAdministratorOutput) ToAzureADAdministratorOutputWithContext(ctx context.Context) AzureADAdministratorOutput {
	return o
}

// Type of the sever administrator.
func (o AzureADAdministratorOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADAdministrator) pulumi.StringPtrOutput { return v.AdministratorType }).(pulumi.StringPtrOutput)
}

// The resource id of the identity used for AAD Authentication.
func (o AzureADAdministratorOutput) IdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADAdministrator) pulumi.StringPtrOutput { return v.IdentityResourceId }).(pulumi.StringPtrOutput)
}

// Login name of the server administrator.
func (o AzureADAdministratorOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADAdministrator) pulumi.StringPtrOutput { return v.Login }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o AzureADAdministratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureADAdministrator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SID (object ID) of the server administrator.
func (o AzureADAdministratorOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADAdministrator) pulumi.StringPtrOutput { return v.Sid }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o AzureADAdministratorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AzureADAdministrator) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Tenant ID of the administrator.
func (o AzureADAdministratorOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADAdministrator) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AzureADAdministratorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureADAdministrator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureADAdministratorOutput{})
}
