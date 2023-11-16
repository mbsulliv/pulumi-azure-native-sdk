// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Managed certificates used for Custom Domain bindings of Container Apps in a Managed Environment
type ManagedCertificate struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Certificate resource specific properties
	Properties ManagedCertificateResponsePropertiesOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedCertificate registers a new resource with the given unique name, arguments, and options.
func NewManagedCertificate(ctx *pulumi.Context,
	name string, args *ManagedCertificateArgs, opts ...pulumi.ResourceOption) (*ManagedCertificate, error) {
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
			Type: pulumi.String("azure-native:app:ManagedCertificate"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221101preview:ManagedCertificate"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230401preview:ManagedCertificate"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230502preview:ManagedCertificate"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230801preview:ManagedCertificate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedCertificate
	err := ctx.RegisterResource("azure-native:app/v20230501:ManagedCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedCertificate gets an existing ManagedCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedCertificateState, opts ...pulumi.ResourceOption) (*ManagedCertificate, error) {
	var resource ManagedCertificate
	err := ctx.ReadResource("azure-native:app/v20230501:ManagedCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedCertificate resources.
type managedCertificateState struct {
}

type ManagedCertificateState struct {
}

func (ManagedCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedCertificateState)(nil)).Elem()
}

type managedCertificateArgs struct {
	// Name of the Managed Environment.
	EnvironmentName string `pulumi:"environmentName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of the Managed Certificate.
	ManagedCertificateName *string `pulumi:"managedCertificateName"`
	// Certificate resource specific properties
	Properties *ManagedCertificateProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ManagedCertificate resource.
type ManagedCertificateArgs struct {
	// Name of the Managed Environment.
	EnvironmentName pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of the Managed Certificate.
	ManagedCertificateName pulumi.StringPtrInput
	// Certificate resource specific properties
	Properties ManagedCertificatePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ManagedCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedCertificateArgs)(nil)).Elem()
}

type ManagedCertificateInput interface {
	pulumi.Input

	ToManagedCertificateOutput() ManagedCertificateOutput
	ToManagedCertificateOutputWithContext(ctx context.Context) ManagedCertificateOutput
}

func (*ManagedCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCertificate)(nil)).Elem()
}

func (i *ManagedCertificate) ToManagedCertificateOutput() ManagedCertificateOutput {
	return i.ToManagedCertificateOutputWithContext(context.Background())
}

func (i *ManagedCertificate) ToManagedCertificateOutputWithContext(ctx context.Context) ManagedCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedCertificateOutput)
}

type ManagedCertificateOutput struct{ *pulumi.OutputState }

func (ManagedCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCertificate)(nil)).Elem()
}

func (o ManagedCertificateOutput) ToManagedCertificateOutput() ManagedCertificateOutput {
	return o
}

func (o ManagedCertificateOutput) ToManagedCertificateOutputWithContext(ctx context.Context) ManagedCertificateOutput {
	return o
}

// The geo-location where the resource lives
func (o ManagedCertificateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCertificate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ManagedCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Certificate resource specific properties
func (o ManagedCertificateOutput) Properties() ManagedCertificateResponsePropertiesOutput {
	return o.ApplyT(func(v *ManagedCertificate) ManagedCertificateResponsePropertiesOutput { return v.Properties }).(ManagedCertificateResponsePropertiesOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ManagedCertificateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedCertificate) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ManagedCertificateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedCertificate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ManagedCertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCertificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedCertificateOutput{})
}
