// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package windowsiot

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The description of the Windows IoT Device Service.
// Azure REST API version: 2019-06-01. Prior API version in Azure Native 1.x: 2019-06-01.
//
// Other available API versions: 2018-02-16-preview.
type Service struct {
	pulumi.CustomResourceState

	// Windows IoT Device Service OEM AAD domain
	AdminDomainName pulumi.StringPtrOutput `pulumi:"adminDomainName"`
	// Windows IoT Device Service ODM AAD domain
	BillingDomainName pulumi.StringPtrOutput `pulumi:"billingDomainName"`
	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The Azure Region where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Windows IoT Device Service notes.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// Windows IoT Device Service device allocation,
	Quantity pulumi.Float64PtrOutput `pulumi:"quantity"`
	// Windows IoT Device Service start date,
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:windowsiot/v20180216preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:windowsiot/v20190601:Service"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Service
	err := ctx.RegisterResource("azure-native:windowsiot:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:windowsiot:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
}

type ServiceState struct {
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Windows IoT Device Service OEM AAD domain
	AdminDomainName *string `pulumi:"adminDomainName"`
	// Windows IoT Device Service ODM AAD domain
	BillingDomainName *string `pulumi:"billingDomainName"`
	// The name of the Windows IoT Device Service.
	DeviceName *string `pulumi:"deviceName"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// Windows IoT Device Service notes.
	Notes *string `pulumi:"notes"`
	// Windows IoT Device Service device allocation,
	Quantity *float64 `pulumi:"quantity"`
	// The name of the resource group that contains the Windows IoT Device Service.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Windows IoT Device Service OEM AAD domain
	AdminDomainName pulumi.StringPtrInput
	// Windows IoT Device Service ODM AAD domain
	BillingDomainName pulumi.StringPtrInput
	// The name of the Windows IoT Device Service.
	DeviceName pulumi.StringPtrInput
	// The Azure Region where the resource lives
	Location pulumi.StringPtrInput
	// Windows IoT Device Service notes.
	Notes pulumi.StringPtrInput
	// Windows IoT Device Service device allocation,
	Quantity pulumi.Float64PtrInput
	// The name of the resource group that contains the Windows IoT Device Service.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

// Windows IoT Device Service OEM AAD domain
func (o ServiceOutput) AdminDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.AdminDomainName }).(pulumi.StringPtrOutput)
}

// Windows IoT Device Service ODM AAD domain
func (o ServiceOutput) BillingDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.BillingDomainName }).(pulumi.StringPtrOutput)
}

// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
func (o ServiceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The Azure Region where the resource lives
func (o ServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Windows IoT Device Service notes.
func (o ServiceOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// Windows IoT Device Service device allocation,
func (o ServiceOutput) Quantity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Service) pulumi.Float64PtrOutput { return v.Quantity }).(pulumi.Float64PtrOutput)
}

// Windows IoT Device Service start date,
func (o ServiceOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

// Resource tags.
func (o ServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o ServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
