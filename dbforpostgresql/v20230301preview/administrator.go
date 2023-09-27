// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents an Active Directory administrator.
type Administrator struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The objectId of the Active Directory administrator.
	ObjectId pulumi.StringPtrOutput `pulumi:"objectId"`
	// Active Directory administrator principal name.
	PrincipalName pulumi.StringPtrOutput `pulumi:"principalName"`
	// The principal type used to represent the type of Active Directory Administrator.
	PrincipalType pulumi.StringPtrOutput `pulumi:"principalType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tenantId of the Active Directory administrator.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAdministrator registers a new resource with the given unique name, arguments, and options.
func NewAdministrator(ctx *pulumi.Context,
	name string, args *AdministratorArgs, opts ...pulumi.ResourceOption) (*Administrator, error) {
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
			Type: pulumi.String("azure-native:dbforpostgresql:Administrator"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20220308preview:Administrator"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20221201:Administrator"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Administrator
	err := ctx.RegisterResource("azure-native:dbforpostgresql/v20230301preview:Administrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdministrator gets an existing Administrator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdministratorState, opts ...pulumi.ResourceOption) (*Administrator, error) {
	var resource Administrator
	err := ctx.ReadResource("azure-native:dbforpostgresql/v20230301preview:Administrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Administrator resources.
type administratorState struct {
}

type AdministratorState struct {
}

func (AdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*administratorState)(nil)).Elem()
}

type administratorArgs struct {
	// Guid of the objectId for the administrator.
	ObjectId *string `pulumi:"objectId"`
	// Active Directory administrator principal name.
	PrincipalName *string `pulumi:"principalName"`
	// The principal type used to represent the type of Active Directory Administrator.
	PrincipalType *string `pulumi:"principalType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The tenantId of the Active Directory administrator.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a Administrator resource.
type AdministratorArgs struct {
	// Guid of the objectId for the administrator.
	ObjectId pulumi.StringPtrInput
	// Active Directory administrator principal name.
	PrincipalName pulumi.StringPtrInput
	// The principal type used to represent the type of Active Directory Administrator.
	PrincipalType pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The tenantId of the Active Directory administrator.
	TenantId pulumi.StringPtrInput
}

func (AdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*administratorArgs)(nil)).Elem()
}

type AdministratorInput interface {
	pulumi.Input

	ToAdministratorOutput() AdministratorOutput
	ToAdministratorOutputWithContext(ctx context.Context) AdministratorOutput
}

func (*Administrator) ElementType() reflect.Type {
	return reflect.TypeOf((**Administrator)(nil)).Elem()
}

func (i *Administrator) ToAdministratorOutput() AdministratorOutput {
	return i.ToAdministratorOutputWithContext(context.Background())
}

func (i *Administrator) ToAdministratorOutputWithContext(ctx context.Context) AdministratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdministratorOutput)
}

func (i *Administrator) ToOutput(ctx context.Context) pulumix.Output[*Administrator] {
	return pulumix.Output[*Administrator]{
		OutputState: i.ToAdministratorOutputWithContext(ctx).OutputState,
	}
}

type AdministratorOutput struct{ *pulumi.OutputState }

func (AdministratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Administrator)(nil)).Elem()
}

func (o AdministratorOutput) ToAdministratorOutput() AdministratorOutput {
	return o
}

func (o AdministratorOutput) ToAdministratorOutputWithContext(ctx context.Context) AdministratorOutput {
	return o
}

func (o AdministratorOutput) ToOutput(ctx context.Context) pulumix.Output[*Administrator] {
	return pulumix.Output[*Administrator]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o AdministratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Administrator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The objectId of the Active Directory administrator.
func (o AdministratorOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Administrator) pulumi.StringPtrOutput { return v.ObjectId }).(pulumi.StringPtrOutput)
}

// Active Directory administrator principal name.
func (o AdministratorOutput) PrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Administrator) pulumi.StringPtrOutput { return v.PrincipalName }).(pulumi.StringPtrOutput)
}

// The principal type used to represent the type of Active Directory Administrator.
func (o AdministratorOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Administrator) pulumi.StringPtrOutput { return v.PrincipalType }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AdministratorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Administrator) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenantId of the Active Directory administrator.
func (o AdministratorOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Administrator) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AdministratorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Administrator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AdministratorOutput{})
}
