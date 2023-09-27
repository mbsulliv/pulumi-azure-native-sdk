// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package powerplatform

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of the EnterprisePolicy.
// Azure REST API version: 2020-10-30-preview. Prior API version in Azure Native 1.x: 2020-10-30-preview
type EnterprisePolicy struct {
	pulumi.CustomResourceState

	// The encryption settings for a configuration store.
	Encryption PropertiesResponseEncryptionPtrOutput `pulumi:"encryption"`
	// The health status of the resource.
	HealthStatus pulumi.StringPtrOutput `pulumi:"healthStatus"`
	// The identity of the EnterprisePolicy.
	Identity EnterprisePolicyIdentityResponsePtrOutput `pulumi:"identity"`
	// The kind (type) of Enterprise Policy.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Settings concerning lockbox.
	Lockbox PropertiesResponseLockboxPtrOutput `pulumi:"lockbox"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Settings concerning network injection.
	NetworkInjection PropertiesResponseNetworkInjectionPtrOutput `pulumi:"networkInjection"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The internally assigned unique identifier of the resource.
	SystemId pulumi.StringOutput `pulumi:"systemId"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEnterprisePolicy registers a new resource with the given unique name, arguments, and options.
func NewEnterprisePolicy(ctx *pulumi.Context,
	name string, args *EnterprisePolicyArgs, opts ...pulumi.ResourceOption) (*EnterprisePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:powerplatform/v20201030preview:EnterprisePolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EnterprisePolicy
	err := ctx.RegisterResource("azure-native:powerplatform:EnterprisePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterprisePolicy gets an existing EnterprisePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterprisePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterprisePolicyState, opts ...pulumi.ResourceOption) (*EnterprisePolicy, error) {
	var resource EnterprisePolicy
	err := ctx.ReadResource("azure-native:powerplatform:EnterprisePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterprisePolicy resources.
type enterprisePolicyState struct {
}

type EnterprisePolicyState struct {
}

func (EnterprisePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterprisePolicyState)(nil)).Elem()
}

type enterprisePolicyArgs struct {
	// The encryption settings for a configuration store.
	Encryption *PropertiesEncryption `pulumi:"encryption"`
	// Name of the EnterprisePolicy.
	EnterprisePolicyName *string `pulumi:"enterprisePolicyName"`
	// The health status of the resource.
	HealthStatus *string `pulumi:"healthStatus"`
	// The identity of the EnterprisePolicy.
	Identity *EnterprisePolicyIdentity `pulumi:"identity"`
	// The kind (type) of Enterprise Policy.
	Kind string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Settings concerning lockbox.
	Lockbox *PropertiesLockbox `pulumi:"lockbox"`
	// Settings concerning network injection.
	NetworkInjection *PropertiesNetworkInjection `pulumi:"networkInjection"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EnterprisePolicy resource.
type EnterprisePolicyArgs struct {
	// The encryption settings for a configuration store.
	Encryption PropertiesEncryptionPtrInput
	// Name of the EnterprisePolicy.
	EnterprisePolicyName pulumi.StringPtrInput
	// The health status of the resource.
	HealthStatus pulumi.StringPtrInput
	// The identity of the EnterprisePolicy.
	Identity EnterprisePolicyIdentityPtrInput
	// The kind (type) of Enterprise Policy.
	Kind pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Settings concerning lockbox.
	Lockbox PropertiesLockboxPtrInput
	// Settings concerning network injection.
	NetworkInjection PropertiesNetworkInjectionPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (EnterprisePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterprisePolicyArgs)(nil)).Elem()
}

type EnterprisePolicyInput interface {
	pulumi.Input

	ToEnterprisePolicyOutput() EnterprisePolicyOutput
	ToEnterprisePolicyOutputWithContext(ctx context.Context) EnterprisePolicyOutput
}

func (*EnterprisePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterprisePolicy)(nil)).Elem()
}

func (i *EnterprisePolicy) ToEnterprisePolicyOutput() EnterprisePolicyOutput {
	return i.ToEnterprisePolicyOutputWithContext(context.Background())
}

func (i *EnterprisePolicy) ToEnterprisePolicyOutputWithContext(ctx context.Context) EnterprisePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterprisePolicyOutput)
}

func (i *EnterprisePolicy) ToOutput(ctx context.Context) pulumix.Output[*EnterprisePolicy] {
	return pulumix.Output[*EnterprisePolicy]{
		OutputState: i.ToEnterprisePolicyOutputWithContext(ctx).OutputState,
	}
}

type EnterprisePolicyOutput struct{ *pulumi.OutputState }

func (EnterprisePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterprisePolicy)(nil)).Elem()
}

func (o EnterprisePolicyOutput) ToEnterprisePolicyOutput() EnterprisePolicyOutput {
	return o
}

func (o EnterprisePolicyOutput) ToEnterprisePolicyOutputWithContext(ctx context.Context) EnterprisePolicyOutput {
	return o
}

func (o EnterprisePolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*EnterprisePolicy] {
	return pulumix.Output[*EnterprisePolicy]{
		OutputState: o.OutputState,
	}
}

// The encryption settings for a configuration store.
func (o EnterprisePolicyOutput) Encryption() PropertiesResponseEncryptionPtrOutput {
	return o.ApplyT(func(v *EnterprisePolicy) PropertiesResponseEncryptionPtrOutput { return v.Encryption }).(PropertiesResponseEncryptionPtrOutput)
}

// The health status of the resource.
func (o EnterprisePolicyOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterprisePolicy) pulumi.StringPtrOutput { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

// The identity of the EnterprisePolicy.
func (o EnterprisePolicyOutput) Identity() EnterprisePolicyIdentityResponsePtrOutput {
	return o.ApplyT(func(v *EnterprisePolicy) EnterprisePolicyIdentityResponsePtrOutput { return v.Identity }).(EnterprisePolicyIdentityResponsePtrOutput)
}

// The kind (type) of Enterprise Policy.
func (o EnterprisePolicyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterprisePolicy) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o EnterprisePolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterprisePolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Settings concerning lockbox.
func (o EnterprisePolicyOutput) Lockbox() PropertiesResponseLockboxPtrOutput {
	return o.ApplyT(func(v *EnterprisePolicy) PropertiesResponseLockboxPtrOutput { return v.Lockbox }).(PropertiesResponseLockboxPtrOutput)
}

// The name of the resource
func (o EnterprisePolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterprisePolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Settings concerning network injection.
func (o EnterprisePolicyOutput) NetworkInjection() PropertiesResponseNetworkInjectionPtrOutput {
	return o.ApplyT(func(v *EnterprisePolicy) PropertiesResponseNetworkInjectionPtrOutput { return v.NetworkInjection }).(PropertiesResponseNetworkInjectionPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o EnterprisePolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EnterprisePolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The internally assigned unique identifier of the resource.
func (o EnterprisePolicyOutput) SystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterprisePolicy) pulumi.StringOutput { return v.SystemId }).(pulumi.StringOutput)
}

// Resource tags.
func (o EnterprisePolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EnterprisePolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EnterprisePolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterprisePolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnterprisePolicyOutput{})
}
