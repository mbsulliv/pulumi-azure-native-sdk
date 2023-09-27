// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220808

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of the credential.
type Credential struct {
	pulumi.CustomResourceState

	// Gets the creation time.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets the last modified time.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets the user name of the credential.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewCredential registers a new resource with the given unique name, arguments, and options.
func NewCredential(ctx *pulumi.Context,
	name string, args *CredentialArgs, opts ...pulumi.ResourceOption) (*Credential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:Credential"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:Credential"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Credential"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:Credential"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Credential
	err := ctx.RegisterResource("azure-native:automation/v20220808:Credential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCredential gets an existing Credential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CredentialState, opts ...pulumi.ResourceOption) (*Credential, error) {
	var resource Credential
	err := ctx.ReadResource("azure-native:automation/v20220808:Credential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Credential resources.
type credentialState struct {
}

type CredentialState struct {
}

func (CredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialState)(nil)).Elem()
}

type credentialArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The parameters supplied to the create or update credential operation.
	CredentialName *string `pulumi:"credentialName"`
	// Gets or sets the description of the credential.
	Description *string `pulumi:"description"`
	// Gets or sets the name of the credential.
	Name string `pulumi:"name"`
	// Gets or sets the password of the credential.
	Password string `pulumi:"password"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the user name of the credential.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a Credential resource.
type CredentialArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// The parameters supplied to the create or update credential operation.
	CredentialName pulumi.StringPtrInput
	// Gets or sets the description of the credential.
	Description pulumi.StringPtrInput
	// Gets or sets the name of the credential.
	Name pulumi.StringInput
	// Gets or sets the password of the credential.
	Password pulumi.StringInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the user name of the credential.
	UserName pulumi.StringInput
}

func (CredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialArgs)(nil)).Elem()
}

type CredentialInput interface {
	pulumi.Input

	ToCredentialOutput() CredentialOutput
	ToCredentialOutputWithContext(ctx context.Context) CredentialOutput
}

func (*Credential) ElementType() reflect.Type {
	return reflect.TypeOf((**Credential)(nil)).Elem()
}

func (i *Credential) ToCredentialOutput() CredentialOutput {
	return i.ToCredentialOutputWithContext(context.Background())
}

func (i *Credential) ToCredentialOutputWithContext(ctx context.Context) CredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialOutput)
}

func (i *Credential) ToOutput(ctx context.Context) pulumix.Output[*Credential] {
	return pulumix.Output[*Credential]{
		OutputState: i.ToCredentialOutputWithContext(ctx).OutputState,
	}
}

type CredentialOutput struct{ *pulumi.OutputState }

func (CredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Credential)(nil)).Elem()
}

func (o CredentialOutput) ToCredentialOutput() CredentialOutput {
	return o
}

func (o CredentialOutput) ToCredentialOutputWithContext(ctx context.Context) CredentialOutput {
	return o
}

func (o CredentialOutput) ToOutput(ctx context.Context) pulumix.Output[*Credential] {
	return pulumix.Output[*Credential]{
		OutputState: o.OutputState,
	}
}

// Gets the creation time.
func (o CredentialOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Gets or sets the description.
func (o CredentialOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets the last modified time.
func (o CredentialOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// The name of the resource
func (o CredentialOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource.
func (o CredentialOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets the user name of the credential.
func (o CredentialOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CredentialOutput{})
}
