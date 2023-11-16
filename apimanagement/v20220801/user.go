// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// User details.
type User struct {
	pulumi.CustomResourceState

	// Email address.
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// First name.
	FirstName pulumi.StringPtrOutput `pulumi:"firstName"`
	// Collection of groups user is part of.
	Groups GroupContractPropertiesResponseArrayOutput `pulumi:"groups"`
	// Collection of user identities.
	Identities UserIdentityContractResponseArrayOutput `pulumi:"identities"`
	// Last name.
	LastName pulumi.StringPtrOutput `pulumi:"lastName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional note about a user set by the administrator.
	Note pulumi.StringPtrOutput `pulumi:"note"`
	// Date of user registration. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	RegistrationDate pulumi.StringPtrOutput `pulumi:"registrationDate"`
	// Account state. Specifies whether the user is active or not. Blocked users are unable to sign into the developer portal or call any APIs of subscribed products. Default state is Active.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.FirstName == nil {
		return nil, errors.New("invalid value for required argument 'FirstName'")
	}
	if args.LastName == nil {
		return nil, errors.New("invalid value for required argument 'LastName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.State == nil {
		args.State = pulumi.StringPtr("active")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:User"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource User
	err := ctx.RegisterResource("azure-native:apimanagement/v20220801:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("azure-native:apimanagement/v20220801:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
}

type UserState struct {
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// Determines the type of application which send the create user request. Default is legacy portal.
	AppType *string `pulumi:"appType"`
	// Determines the type of confirmation e-mail that will be sent to the newly created user.
	Confirmation *string `pulumi:"confirmation"`
	// Email address. Must not be empty and must be unique within the service instance.
	Email string `pulumi:"email"`
	// First name.
	FirstName string `pulumi:"firstName"`
	// Collection of user identities.
	Identities []UserIdentityContract `pulumi:"identities"`
	// Last name.
	LastName string `pulumi:"lastName"`
	// Optional note about a user set by the administrator.
	Note *string `pulumi:"note"`
	// Send an Email notification to the User.
	Notify *bool `pulumi:"notify"`
	// User Password. If no value is provided, a default password is generated.
	Password *string `pulumi:"password"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Account state. Specifies whether the user is active or not. Blocked users are unable to sign into the developer portal or call any APIs of subscribed products. Default state is Active.
	State *string `pulumi:"state"`
	// User identifier. Must be unique in the current API Management service instance.
	UserId *string `pulumi:"userId"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Determines the type of application which send the create user request. Default is legacy portal.
	AppType pulumi.StringPtrInput
	// Determines the type of confirmation e-mail that will be sent to the newly created user.
	Confirmation pulumi.StringPtrInput
	// Email address. Must not be empty and must be unique within the service instance.
	Email pulumi.StringInput
	// First name.
	FirstName pulumi.StringInput
	// Collection of user identities.
	Identities UserIdentityContractArrayInput
	// Last name.
	LastName pulumi.StringInput
	// Optional note about a user set by the administrator.
	Note pulumi.StringPtrInput
	// Send an Email notification to the User.
	Notify pulumi.BoolPtrInput
	// User Password. If no value is provided, a default password is generated.
	Password pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Account state. Specifies whether the user is active or not. Blocked users are unable to sign into the developer portal or call any APIs of subscribed products. Default state is Active.
	State pulumi.StringPtrInput
	// User identifier. Must be unique in the current API Management service instance.
	UserId pulumi.StringPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// Email address.
func (o UserOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

// First name.
func (o UserOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.FirstName }).(pulumi.StringPtrOutput)
}

// Collection of groups user is part of.
func (o UserOutput) Groups() GroupContractPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *User) GroupContractPropertiesResponseArrayOutput { return v.Groups }).(GroupContractPropertiesResponseArrayOutput)
}

// Collection of user identities.
func (o UserOutput) Identities() UserIdentityContractResponseArrayOutput {
	return o.ApplyT(func(v *User) UserIdentityContractResponseArrayOutput { return v.Identities }).(UserIdentityContractResponseArrayOutput)
}

// Last name.
func (o UserOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.LastName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o UserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional note about a user set by the administrator.
func (o UserOutput) Note() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Note }).(pulumi.StringPtrOutput)
}

// Date of user registration. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o UserOutput) RegistrationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.RegistrationDate }).(pulumi.StringPtrOutput)
}

// Account state. Specifies whether the user is active or not. Blocked users are unable to sign into the developer portal or call any APIs of subscribed products. Default state is Active.
func (o UserOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o UserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UserOutput{})
}
