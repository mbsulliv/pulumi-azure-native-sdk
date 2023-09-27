// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Customer subscription.
type CustomerSubscription struct {
	pulumi.CustomResourceState

	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tenant Id.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// Type of Resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCustomerSubscription registers a new resource with the given unique name, arguments, and options.
func NewCustomerSubscription(ctx *pulumi.Context,
	name string, args *CustomerSubscriptionArgs, opts ...pulumi.ResourceOption) (*CustomerSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistrationName == nil {
		return nil, errors.New("invalid value for required argument 'RegistrationName'")
	}
	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestack:CustomerSubscription"),
		},
		{
			Type: pulumi.String("azure-native:azurestack/v20170601:CustomerSubscription"),
		},
		{
			Type: pulumi.String("azure-native:azurestack/v20220601:CustomerSubscription"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CustomerSubscription
	err := ctx.RegisterResource("azure-native:azurestack/v20200601preview:CustomerSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomerSubscription gets an existing CustomerSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomerSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerSubscriptionState, opts ...pulumi.ResourceOption) (*CustomerSubscription, error) {
	var resource CustomerSubscription
	err := ctx.ReadResource("azure-native:azurestack/v20200601preview:CustomerSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomerSubscription resources.
type customerSubscriptionState struct {
}

type CustomerSubscriptionState struct {
}

func (CustomerSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerSubscriptionState)(nil)).Elem()
}

type customerSubscriptionArgs struct {
	// Name of the product.
	CustomerSubscriptionName *string `pulumi:"customerSubscriptionName"`
	// Name of the Azure Stack registration.
	RegistrationName string `pulumi:"registrationName"`
	// Name of the resource group.
	ResourceGroup string `pulumi:"resourceGroup"`
	// Tenant Id.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a CustomerSubscription resource.
type CustomerSubscriptionArgs struct {
	// Name of the product.
	CustomerSubscriptionName pulumi.StringPtrInput
	// Name of the Azure Stack registration.
	RegistrationName pulumi.StringInput
	// Name of the resource group.
	ResourceGroup pulumi.StringInput
	// Tenant Id.
	TenantId pulumi.StringPtrInput
}

func (CustomerSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerSubscriptionArgs)(nil)).Elem()
}

type CustomerSubscriptionInput interface {
	pulumi.Input

	ToCustomerSubscriptionOutput() CustomerSubscriptionOutput
	ToCustomerSubscriptionOutputWithContext(ctx context.Context) CustomerSubscriptionOutput
}

func (*CustomerSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerSubscription)(nil)).Elem()
}

func (i *CustomerSubscription) ToCustomerSubscriptionOutput() CustomerSubscriptionOutput {
	return i.ToCustomerSubscriptionOutputWithContext(context.Background())
}

func (i *CustomerSubscription) ToCustomerSubscriptionOutputWithContext(ctx context.Context) CustomerSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSubscriptionOutput)
}

func (i *CustomerSubscription) ToOutput(ctx context.Context) pulumix.Output[*CustomerSubscription] {
	return pulumix.Output[*CustomerSubscription]{
		OutputState: i.ToCustomerSubscriptionOutputWithContext(ctx).OutputState,
	}
}

type CustomerSubscriptionOutput struct{ *pulumi.OutputState }

func (CustomerSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerSubscription)(nil)).Elem()
}

func (o CustomerSubscriptionOutput) ToCustomerSubscriptionOutput() CustomerSubscriptionOutput {
	return o
}

func (o CustomerSubscriptionOutput) ToCustomerSubscriptionOutputWithContext(ctx context.Context) CustomerSubscriptionOutput {
	return o
}

func (o CustomerSubscriptionOutput) ToOutput(ctx context.Context) pulumix.Output[*CustomerSubscription] {
	return pulumix.Output[*CustomerSubscription]{
		OutputState: o.OutputState,
	}
}

// The entity tag used for optimistic concurrency when modifying the resource.
func (o CustomerSubscriptionOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerSubscription) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o CustomerSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o CustomerSubscriptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CustomerSubscription) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Tenant Id.
func (o CustomerSubscriptionOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerSubscription) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// Type of Resource.
func (o CustomerSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomerSubscriptionOutput{})
}
