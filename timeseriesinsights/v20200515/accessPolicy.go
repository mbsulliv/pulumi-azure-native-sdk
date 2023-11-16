// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200515

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An access policy is used to grant users and applications access to the environment. Roles are assigned to service principals in Azure Active Directory. These roles define the actions the principal can perform through the Time Series Insights data plane APIs.
type AccessPolicy struct {
	pulumi.CustomResourceState

	// An description of the access policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The objectId of the principal in Azure Active Directory.
	PrincipalObjectId pulumi.StringPtrOutput `pulumi:"principalObjectId"`
	// The list of roles the principal is assigned on the environment.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewAccessPolicy(ctx *pulumi.Context,
	name string, args *AccessPolicyArgs, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:timeseriesinsights:AccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20170228preview:AccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20171115:AccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20180815preview:AccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210331preview:AccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210630preview:AccessPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AccessPolicy
	err := ctx.RegisterResource("azure-native:timeseriesinsights/v20200515:AccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPolicy gets an existing AccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPolicyState, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	var resource AccessPolicy
	err := ctx.ReadResource("azure-native:timeseriesinsights/v20200515:AccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPolicy resources.
type accessPolicyState struct {
}

type AccessPolicyState struct {
}

func (AccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyState)(nil)).Elem()
}

type accessPolicyArgs struct {
	// Name of the access policy.
	AccessPolicyName *string `pulumi:"accessPolicyName"`
	// An description of the access policy.
	Description *string `pulumi:"description"`
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName string `pulumi:"environmentName"`
	// The objectId of the principal in Azure Active Directory.
	PrincipalObjectId *string `pulumi:"principalObjectId"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of roles the principal is assigned on the environment.
	Roles []string `pulumi:"roles"`
}

// The set of arguments for constructing a AccessPolicy resource.
type AccessPolicyArgs struct {
	// Name of the access policy.
	AccessPolicyName pulumi.StringPtrInput
	// An description of the access policy.
	Description pulumi.StringPtrInput
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName pulumi.StringInput
	// The objectId of the principal in Azure Active Directory.
	PrincipalObjectId pulumi.StringPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// The list of roles the principal is assigned on the environment.
	Roles pulumi.StringArrayInput
}

func (AccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyArgs)(nil)).Elem()
}

type AccessPolicyInput interface {
	pulumi.Input

	ToAccessPolicyOutput() AccessPolicyOutput
	ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput
}

func (*AccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicy)(nil)).Elem()
}

func (i *AccessPolicy) ToAccessPolicyOutput() AccessPolicyOutput {
	return i.ToAccessPolicyOutputWithContext(context.Background())
}

func (i *AccessPolicy) ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyOutput)
}

type AccessPolicyOutput struct{ *pulumi.OutputState }

func (AccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicy)(nil)).Elem()
}

func (o AccessPolicyOutput) ToAccessPolicyOutput() AccessPolicyOutput {
	return o
}

func (o AccessPolicyOutput) ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput {
	return o
}

// An description of the access policy.
func (o AccessPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Resource name
func (o AccessPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The objectId of the principal in Azure Active Directory.
func (o AccessPolicyOutput) PrincipalObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringPtrOutput { return v.PrincipalObjectId }).(pulumi.StringPtrOutput)
}

// The list of roles the principal is assigned on the environment.
func (o AccessPolicyOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// Resource type
func (o AccessPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessPolicyOutput{})
}
