// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package voiceservices

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A Contact resource
// Azure REST API version: 2022-12-01-preview. Prior API version in Azure Native 1.x: 2022-12-01-preview.
type Contact struct {
	pulumi.CustomResourceState

	// Full name of contact
	ContactName pulumi.StringOutput `pulumi:"contactName"`
	// Email address of contact
	Email pulumi.StringOutput `pulumi:"email"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Telephone number of contact
	PhoneNumber pulumi.StringOutput `pulumi:"phoneNumber"`
	// Resource provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Job title of contact
	Role pulumi.StringOutput `pulumi:"role"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewContact registers a new resource with the given unique name, arguments, and options.
func NewContact(ctx *pulumi.Context,
	name string, args *ContactArgs, opts ...pulumi.ResourceOption) (*Contact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CommunicationsGatewayName == nil {
		return nil, errors.New("invalid value for required argument 'CommunicationsGatewayName'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.PhoneNumber == nil {
		return nil, errors.New("invalid value for required argument 'PhoneNumber'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:voiceservices/v20221201preview:Contact"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Contact
	err := ctx.RegisterResource("azure-native:voiceservices:Contact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContact gets an existing Contact resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactState, opts ...pulumi.ResourceOption) (*Contact, error) {
	var resource Contact
	err := ctx.ReadResource("azure-native:voiceservices:Contact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Contact resources.
type contactState struct {
}

type ContactState struct {
}

func (ContactState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactState)(nil)).Elem()
}

type contactArgs struct {
	// Unique identifier for this deployment
	CommunicationsGatewayName string `pulumi:"communicationsGatewayName"`
	// Full name of contact
	ContactName *string `pulumi:"contactName"`
	// Email address of contact
	Email string `pulumi:"email"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Telephone number of contact
	PhoneNumber string `pulumi:"phoneNumber"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Job title of contact
	Role string `pulumi:"role"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Contact resource.
type ContactArgs struct {
	// Unique identifier for this deployment
	CommunicationsGatewayName pulumi.StringInput
	// Full name of contact
	ContactName pulumi.StringPtrInput
	// Email address of contact
	Email pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Telephone number of contact
	PhoneNumber pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Job title of contact
	Role pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactArgs)(nil)).Elem()
}

type ContactInput interface {
	pulumi.Input

	ToContactOutput() ContactOutput
	ToContactOutputWithContext(ctx context.Context) ContactOutput
}

func (*Contact) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (i *Contact) ToContactOutput() ContactOutput {
	return i.ToContactOutputWithContext(context.Background())
}

func (i *Contact) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactOutput)
}

func (i *Contact) ToOutput(ctx context.Context) pulumix.Output[*Contact] {
	return pulumix.Output[*Contact]{
		OutputState: i.ToContactOutputWithContext(ctx).OutputState,
	}
}

type ContactOutput struct{ *pulumi.OutputState }

func (ContactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (o ContactOutput) ToContactOutput() ContactOutput {
	return o
}

func (o ContactOutput) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return o
}

func (o ContactOutput) ToOutput(ctx context.Context) pulumix.Output[*Contact] {
	return pulumix.Output[*Contact]{
		OutputState: o.OutputState,
	}
}

// Full name of contact
func (o ContactOutput) ContactName() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.ContactName }).(pulumi.StringOutput)
}

// Email address of contact
func (o ContactOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o ContactOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ContactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Telephone number of contact
func (o ContactOutput) PhoneNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.PhoneNumber }).(pulumi.StringOutput)
}

// Resource provisioning state.
func (o ContactOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Job title of contact
func (o ContactOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ContactOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Contact) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ContactOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ContactOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContactOutput{})
}
