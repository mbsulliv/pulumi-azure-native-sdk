// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The CA Certificate resource.
type CaCertificate struct {
	pulumi.CustomResourceState

	// Description for the CA Certificate resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Base64 encoded PEM (Privacy Enhanced Mail) format certificate data.
	EncodedCertificate pulumi.StringPtrOutput `pulumi:"encodedCertificate"`
	// Certificate expiry time in UTC. This is a read-only field.
	ExpiryTimeInUtc pulumi.StringOutput `pulumi:"expiryTimeInUtc"`
	// Certificate issue time in UTC. This is a read-only field.
	IssueTimeInUtc pulumi.StringOutput `pulumi:"issueTimeInUtc"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the CA Certificate resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system metadata relating to the CaCertificate resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCaCertificate registers a new resource with the given unique name, arguments, and options.
func NewCaCertificate(ctx *pulumi.Context,
	name string, args *CaCertificateArgs, opts ...pulumi.ResourceOption) (*CaCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:CaCertificate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CaCertificate
	err := ctx.RegisterResource("azure-native:eventgrid/v20230601preview:CaCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCaCertificate gets an existing CaCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCaCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CaCertificateState, opts ...pulumi.ResourceOption) (*CaCertificate, error) {
	var resource CaCertificate
	err := ctx.ReadResource("azure-native:eventgrid/v20230601preview:CaCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CaCertificate resources.
type caCertificateState struct {
}

type CaCertificateState struct {
}

func (CaCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*caCertificateState)(nil)).Elem()
}

type caCertificateArgs struct {
	// The CA certificate name.
	CaCertificateName *string `pulumi:"caCertificateName"`
	// Description for the CA Certificate resource.
	Description *string `pulumi:"description"`
	// Base64 encoded PEM (Privacy Enhanced Mail) format certificate data.
	EncodedCertificate *string `pulumi:"encodedCertificate"`
	// Name of the namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a CaCertificate resource.
type CaCertificateArgs struct {
	// The CA certificate name.
	CaCertificateName pulumi.StringPtrInput
	// Description for the CA Certificate resource.
	Description pulumi.StringPtrInput
	// Base64 encoded PEM (Privacy Enhanced Mail) format certificate data.
	EncodedCertificate pulumi.StringPtrInput
	// Name of the namespace.
	NamespaceName pulumi.StringInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
}

func (CaCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*caCertificateArgs)(nil)).Elem()
}

type CaCertificateInput interface {
	pulumi.Input

	ToCaCertificateOutput() CaCertificateOutput
	ToCaCertificateOutputWithContext(ctx context.Context) CaCertificateOutput
}

func (*CaCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**CaCertificate)(nil)).Elem()
}

func (i *CaCertificate) ToCaCertificateOutput() CaCertificateOutput {
	return i.ToCaCertificateOutputWithContext(context.Background())
}

func (i *CaCertificate) ToCaCertificateOutputWithContext(ctx context.Context) CaCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaCertificateOutput)
}

func (i *CaCertificate) ToOutput(ctx context.Context) pulumix.Output[*CaCertificate] {
	return pulumix.Output[*CaCertificate]{
		OutputState: i.ToCaCertificateOutputWithContext(ctx).OutputState,
	}
}

type CaCertificateOutput struct{ *pulumi.OutputState }

func (CaCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CaCertificate)(nil)).Elem()
}

func (o CaCertificateOutput) ToCaCertificateOutput() CaCertificateOutput {
	return o
}

func (o CaCertificateOutput) ToCaCertificateOutputWithContext(ctx context.Context) CaCertificateOutput {
	return o
}

func (o CaCertificateOutput) ToOutput(ctx context.Context) pulumix.Output[*CaCertificate] {
	return pulumix.Output[*CaCertificate]{
		OutputState: o.OutputState,
	}
}

// Description for the CA Certificate resource.
func (o CaCertificateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CaCertificate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Base64 encoded PEM (Privacy Enhanced Mail) format certificate data.
func (o CaCertificateOutput) EncodedCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CaCertificate) pulumi.StringPtrOutput { return v.EncodedCertificate }).(pulumi.StringPtrOutput)
}

// Certificate expiry time in UTC. This is a read-only field.
func (o CaCertificateOutput) ExpiryTimeInUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *CaCertificate) pulumi.StringOutput { return v.ExpiryTimeInUtc }).(pulumi.StringOutput)
}

// Certificate issue time in UTC. This is a read-only field.
func (o CaCertificateOutput) IssueTimeInUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *CaCertificate) pulumi.StringOutput { return v.IssueTimeInUtc }).(pulumi.StringOutput)
}

// Name of the resource.
func (o CaCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CaCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the CA Certificate resource.
func (o CaCertificateOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CaCertificate) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system metadata relating to the CaCertificate resource.
func (o CaCertificateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CaCertificate) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the resource.
func (o CaCertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CaCertificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CaCertificateOutput{})
}
