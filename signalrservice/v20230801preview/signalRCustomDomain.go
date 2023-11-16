// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A custom domain
type SignalRCustomDomain struct {
	pulumi.CustomResourceState

	// Reference to a resource.
	CustomCertificate ResourceReferenceResponseOutput `pulumi:"customCertificate"`
	// The custom domain name.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSignalRCustomDomain registers a new resource with the given unique name, arguments, and options.
func NewSignalRCustomDomain(ctx *pulumi.Context,
	name string, args *SignalRCustomDomainArgs, opts ...pulumi.ResourceOption) (*SignalRCustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CustomCertificate == nil {
		return nil, errors.New("invalid value for required argument 'CustomCertificate'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:signalrservice:SignalRCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220201:SignalRCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220801preview:SignalRCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230201:SignalRCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230301preview:SignalRCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230601preview:SignalRCustomDomain"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SignalRCustomDomain
	err := ctx.RegisterResource("azure-native:signalrservice/v20230801preview:SignalRCustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSignalRCustomDomain gets an existing SignalRCustomDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSignalRCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignalRCustomDomainState, opts ...pulumi.ResourceOption) (*SignalRCustomDomain, error) {
	var resource SignalRCustomDomain
	err := ctx.ReadResource("azure-native:signalrservice/v20230801preview:SignalRCustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SignalRCustomDomain resources.
type signalRCustomDomainState struct {
}

type SignalRCustomDomainState struct {
}

func (SignalRCustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRCustomDomainState)(nil)).Elem()
}

type signalRCustomDomainArgs struct {
	// Reference to a resource.
	CustomCertificate ResourceReference `pulumi:"customCertificate"`
	// The custom domain name.
	DomainName string `pulumi:"domainName"`
	// Custom domain name.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a SignalRCustomDomain resource.
type SignalRCustomDomainArgs struct {
	// Reference to a resource.
	CustomCertificate ResourceReferenceInput
	// The custom domain name.
	DomainName pulumi.StringInput
	// Custom domain name.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the resource.
	ResourceName pulumi.StringInput
}

func (SignalRCustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRCustomDomainArgs)(nil)).Elem()
}

type SignalRCustomDomainInput interface {
	pulumi.Input

	ToSignalRCustomDomainOutput() SignalRCustomDomainOutput
	ToSignalRCustomDomainOutputWithContext(ctx context.Context) SignalRCustomDomainOutput
}

func (*SignalRCustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCustomDomain)(nil)).Elem()
}

func (i *SignalRCustomDomain) ToSignalRCustomDomainOutput() SignalRCustomDomainOutput {
	return i.ToSignalRCustomDomainOutputWithContext(context.Background())
}

func (i *SignalRCustomDomain) ToSignalRCustomDomainOutputWithContext(ctx context.Context) SignalRCustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCustomDomainOutput)
}

type SignalRCustomDomainOutput struct{ *pulumi.OutputState }

func (SignalRCustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCustomDomain)(nil)).Elem()
}

func (o SignalRCustomDomainOutput) ToSignalRCustomDomainOutput() SignalRCustomDomainOutput {
	return o
}

func (o SignalRCustomDomainOutput) ToSignalRCustomDomainOutputWithContext(ctx context.Context) SignalRCustomDomainOutput {
	return o
}

// Reference to a resource.
func (o SignalRCustomDomainOutput) CustomCertificate() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v *SignalRCustomDomain) ResourceReferenceResponseOutput { return v.CustomCertificate }).(ResourceReferenceResponseOutput)
}

// The custom domain name.
func (o SignalRCustomDomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomDomain) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The name of the resource
func (o SignalRCustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o SignalRCustomDomainOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomDomain) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SignalRCustomDomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SignalRCustomDomain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SignalRCustomDomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomDomain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SignalRCustomDomainOutput{})
}
