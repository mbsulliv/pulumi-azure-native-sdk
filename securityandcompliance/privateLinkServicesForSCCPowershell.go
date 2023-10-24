// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityandcompliance

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The description of the service.
// Azure REST API version: 2021-03-08. Prior API version in Azure Native 1.x: 2021-03-08.
type PrivateLinkServicesForSCCPowershell struct {
	pulumi.CustomResourceState

	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ServicesResourceResponseIdentityPtrOutput `pulumi:"identity"`
	// The kind of the service.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The common properties of a service.
	Properties ServicesPropertiesResponseOutput `pulumi:"properties"`
	// Required property for system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateLinkServicesForSCCPowershell registers a new resource with the given unique name, arguments, and options.
func NewPrivateLinkServicesForSCCPowershell(ctx *pulumi.Context,
	name string, args *PrivateLinkServicesForSCCPowershellArgs, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForSCCPowershell, error) {
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
			Type: pulumi.String("azure-native:securityandcompliance:privateLinkServicesForSCCPowershell"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance/v20210111:PrivateLinkServicesForSCCPowershell"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance/v20210111:privateLinkServicesForSCCPowershell"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance/v20210308:PrivateLinkServicesForSCCPowershell"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance/v20210308:privateLinkServicesForSCCPowershell"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateLinkServicesForSCCPowershell
	err := ctx.RegisterResource("azure-native:securityandcompliance:PrivateLinkServicesForSCCPowershell", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateLinkServicesForSCCPowershell gets an existing PrivateLinkServicesForSCCPowershell resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateLinkServicesForSCCPowershell(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServicesForSCCPowershellState, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForSCCPowershell, error) {
	var resource PrivateLinkServicesForSCCPowershell
	err := ctx.ReadResource("azure-native:securityandcompliance:PrivateLinkServicesForSCCPowershell", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateLinkServicesForSCCPowershell resources.
type privateLinkServicesForSCCPowershellState struct {
}

type PrivateLinkServicesForSCCPowershellState struct {
}

func (PrivateLinkServicesForSCCPowershellState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForSCCPowershellState)(nil)).Elem()
}

type privateLinkServicesForSCCPowershellArgs struct {
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServicesResourceIdentity `pulumi:"identity"`
	// The kind of the service.
	Kind Kind `pulumi:"kind"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The common properties of a service.
	Properties *ServicesProperties `pulumi:"properties"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName *string `pulumi:"resourceName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PrivateLinkServicesForSCCPowershell resource.
type PrivateLinkServicesForSCCPowershellArgs struct {
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ServicesResourceIdentityPtrInput
	// The kind of the service.
	Kind KindInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The common properties of a service.
	Properties ServicesPropertiesPtrInput
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput
	// The name of the service instance.
	ResourceName pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (PrivateLinkServicesForSCCPowershellArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForSCCPowershellArgs)(nil)).Elem()
}

type PrivateLinkServicesForSCCPowershellInput interface {
	pulumi.Input

	ToPrivateLinkServicesForSCCPowershellOutput() PrivateLinkServicesForSCCPowershellOutput
	ToPrivateLinkServicesForSCCPowershellOutputWithContext(ctx context.Context) PrivateLinkServicesForSCCPowershellOutput
}

func (*PrivateLinkServicesForSCCPowershell) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForSCCPowershell)(nil)).Elem()
}

func (i *PrivateLinkServicesForSCCPowershell) ToPrivateLinkServicesForSCCPowershellOutput() PrivateLinkServicesForSCCPowershellOutput {
	return i.ToPrivateLinkServicesForSCCPowershellOutputWithContext(context.Background())
}

func (i *PrivateLinkServicesForSCCPowershell) ToPrivateLinkServicesForSCCPowershellOutputWithContext(ctx context.Context) PrivateLinkServicesForSCCPowershellOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServicesForSCCPowershellOutput)
}

func (i *PrivateLinkServicesForSCCPowershell) ToOutput(ctx context.Context) pulumix.Output[*PrivateLinkServicesForSCCPowershell] {
	return pulumix.Output[*PrivateLinkServicesForSCCPowershell]{
		OutputState: i.ToPrivateLinkServicesForSCCPowershellOutputWithContext(ctx).OutputState,
	}
}

type PrivateLinkServicesForSCCPowershellOutput struct{ *pulumi.OutputState }

func (PrivateLinkServicesForSCCPowershellOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForSCCPowershell)(nil)).Elem()
}

func (o PrivateLinkServicesForSCCPowershellOutput) ToPrivateLinkServicesForSCCPowershellOutput() PrivateLinkServicesForSCCPowershellOutput {
	return o
}

func (o PrivateLinkServicesForSCCPowershellOutput) ToPrivateLinkServicesForSCCPowershellOutputWithContext(ctx context.Context) PrivateLinkServicesForSCCPowershellOutput {
	return o
}

func (o PrivateLinkServicesForSCCPowershellOutput) ToOutput(ctx context.Context) pulumix.Output[*PrivateLinkServicesForSCCPowershell] {
	return pulumix.Output[*PrivateLinkServicesForSCCPowershell]{
		OutputState: o.OutputState,
	}
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o PrivateLinkServicesForSCCPowershellOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Setting indicating whether the service has a managed identity associated with it.
func (o PrivateLinkServicesForSCCPowershellOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) ServicesResourceResponseIdentityPtrOutput {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

// The kind of the service.
func (o PrivateLinkServicesForSCCPowershellOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The resource location.
func (o PrivateLinkServicesForSCCPowershellOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name.
func (o PrivateLinkServicesForSCCPowershellOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The common properties of a service.
func (o PrivateLinkServicesForSCCPowershellOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) ServicesPropertiesResponseOutput { return v.Properties }).(ServicesPropertiesResponseOutput)
}

// Required property for system data
func (o PrivateLinkServicesForSCCPowershellOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource tags.
func (o PrivateLinkServicesForSCCPowershellOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o PrivateLinkServicesForSCCPowershellOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServicesForSCCPowershellOutput{})
}
