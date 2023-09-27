// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Workspace active directory administrator
// Azure REST API version: 2021-06-01. Prior API version in Azure Native 1.x: 2021-03-01
//
// Note: SQL AAD Admin is configured automatically during workspace creation and assigned to the current user. One can't add more admins with this resource unless you manually delete the current SQL AAD Admin.
type WorkspaceSqlAadAdmin struct {
	pulumi.CustomResourceState

	// Workspace active directory administrator type
	AdministratorType pulumi.StringPtrOutput `pulumi:"administratorType"`
	// Login of the workspace active directory administrator
	Login pulumi.StringPtrOutput `pulumi:"login"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Object ID of the workspace active directory administrator
	Sid pulumi.StringPtrOutput `pulumi:"sid"`
	// Tenant ID of the workspace active directory administrator
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceSqlAadAdmin registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceSqlAadAdmin(ctx *pulumi.Context,
	name string, args *WorkspaceSqlAadAdminArgs, opts ...pulumi.ResourceOption) (*WorkspaceSqlAadAdmin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:WorkspaceSqlAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:WorkspaceSqlAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:WorkspaceSqlAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:WorkspaceSqlAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:WorkspaceSqlAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:WorkspaceSqlAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:WorkspaceSqlAadAdmin"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceSqlAadAdmin
	err := ctx.RegisterResource("azure-native:synapse:WorkspaceSqlAadAdmin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceSqlAadAdmin gets an existing WorkspaceSqlAadAdmin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceSqlAadAdmin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceSqlAadAdminState, opts ...pulumi.ResourceOption) (*WorkspaceSqlAadAdmin, error) {
	var resource WorkspaceSqlAadAdmin
	err := ctx.ReadResource("azure-native:synapse:WorkspaceSqlAadAdmin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceSqlAadAdmin resources.
type workspaceSqlAadAdminState struct {
}

type WorkspaceSqlAadAdminState struct {
}

func (WorkspaceSqlAadAdminState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceSqlAadAdminState)(nil)).Elem()
}

type workspaceSqlAadAdminArgs struct {
	// Workspace active directory administrator type
	AdministratorType *string `pulumi:"administratorType"`
	// Login of the workspace active directory administrator
	Login *string `pulumi:"login"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Object ID of the workspace active directory administrator
	Sid *string `pulumi:"sid"`
	// Tenant ID of the workspace active directory administrator
	TenantId *string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a WorkspaceSqlAadAdmin resource.
type WorkspaceSqlAadAdminArgs struct {
	// Workspace active directory administrator type
	AdministratorType pulumi.StringPtrInput
	// Login of the workspace active directory administrator
	Login pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Object ID of the workspace active directory administrator
	Sid pulumi.StringPtrInput
	// Tenant ID of the workspace active directory administrator
	TenantId pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (WorkspaceSqlAadAdminArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceSqlAadAdminArgs)(nil)).Elem()
}

type WorkspaceSqlAadAdminInput interface {
	pulumi.Input

	ToWorkspaceSqlAadAdminOutput() WorkspaceSqlAadAdminOutput
	ToWorkspaceSqlAadAdminOutputWithContext(ctx context.Context) WorkspaceSqlAadAdminOutput
}

func (*WorkspaceSqlAadAdmin) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSqlAadAdmin)(nil)).Elem()
}

func (i *WorkspaceSqlAadAdmin) ToWorkspaceSqlAadAdminOutput() WorkspaceSqlAadAdminOutput {
	return i.ToWorkspaceSqlAadAdminOutputWithContext(context.Background())
}

func (i *WorkspaceSqlAadAdmin) ToWorkspaceSqlAadAdminOutputWithContext(ctx context.Context) WorkspaceSqlAadAdminOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSqlAadAdminOutput)
}

func (i *WorkspaceSqlAadAdmin) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceSqlAadAdmin] {
	return pulumix.Output[*WorkspaceSqlAadAdmin]{
		OutputState: i.ToWorkspaceSqlAadAdminOutputWithContext(ctx).OutputState,
	}
}

type WorkspaceSqlAadAdminOutput struct{ *pulumi.OutputState }

func (WorkspaceSqlAadAdminOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSqlAadAdmin)(nil)).Elem()
}

func (o WorkspaceSqlAadAdminOutput) ToWorkspaceSqlAadAdminOutput() WorkspaceSqlAadAdminOutput {
	return o
}

func (o WorkspaceSqlAadAdminOutput) ToWorkspaceSqlAadAdminOutputWithContext(ctx context.Context) WorkspaceSqlAadAdminOutput {
	return o
}

func (o WorkspaceSqlAadAdminOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceSqlAadAdmin] {
	return pulumix.Output[*WorkspaceSqlAadAdmin]{
		OutputState: o.OutputState,
	}
}

// Workspace active directory administrator type
func (o WorkspaceSqlAadAdminOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceSqlAadAdmin) pulumi.StringPtrOutput { return v.AdministratorType }).(pulumi.StringPtrOutput)
}

// Login of the workspace active directory administrator
func (o WorkspaceSqlAadAdminOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceSqlAadAdmin) pulumi.StringPtrOutput { return v.Login }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o WorkspaceSqlAadAdminOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceSqlAadAdmin) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Object ID of the workspace active directory administrator
func (o WorkspaceSqlAadAdminOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceSqlAadAdmin) pulumi.StringPtrOutput { return v.Sid }).(pulumi.StringPtrOutput)
}

// Tenant ID of the workspace active directory administrator
func (o WorkspaceSqlAadAdminOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceSqlAadAdmin) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceSqlAadAdminOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceSqlAadAdmin) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceSqlAadAdminOutput{})
}
