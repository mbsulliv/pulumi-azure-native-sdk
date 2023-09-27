// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210630

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A Database Migration Service resource
type Service struct {
	pulumi.CustomResourceState

	// HTTP strong entity tag value. Ignored if submitted
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The resource kind. Only 'vm' (the default) is supported.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource's provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The public key of the service, used to encrypt secrets sent to the service
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
	// Service SKU
	Sku ServiceSkuResponsePtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The ID of the Microsoft.Network/networkInterfaces resource which the service have
	VirtualNicId pulumi.StringPtrOutput `pulumi:"virtualNicId"`
	// The ID of the Microsoft.Network/virtualNetworks/subnets resource to which the service should be joined
	VirtualSubnetId pulumi.StringOutput `pulumi:"virtualSubnetId"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.VirtualSubnetId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualSubnetId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datamigration:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20171115preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180315preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180331preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180419:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180715preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20211030preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220130preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220330preview:Service"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Service
	err := ctx.RegisterResource("azure-native:datamigration/v20210630:Service", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:datamigration/v20210630:Service", name, id, state, &resource, opts...)
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
	// Name of the resource group
	GroupName string `pulumi:"groupName"`
	// The resource kind. Only 'vm' (the default) is supported.
	Kind *string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The public key of the service, used to encrypt secrets sent to the service
	PublicKey *string `pulumi:"publicKey"`
	// Name of the service
	ServiceName *string `pulumi:"serviceName"`
	// Service SKU
	Sku *ServiceSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Microsoft.Network/networkInterfaces resource which the service have
	VirtualNicId *string `pulumi:"virtualNicId"`
	// The ID of the Microsoft.Network/virtualNetworks/subnets resource to which the service should be joined
	VirtualSubnetId string `pulumi:"virtualSubnetId"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Name of the resource group
	GroupName pulumi.StringInput
	// The resource kind. Only 'vm' (the default) is supported.
	Kind pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The public key of the service, used to encrypt secrets sent to the service
	PublicKey pulumi.StringPtrInput
	// Name of the service
	ServiceName pulumi.StringPtrInput
	// Service SKU
	Sku ServiceSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The ID of the Microsoft.Network/networkInterfaces resource which the service have
	VirtualNicId pulumi.StringPtrInput
	// The ID of the Microsoft.Network/virtualNetworks/subnets resource to which the service should be joined
	VirtualSubnetId pulumi.StringInput
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

func (i *Service) ToOutput(ctx context.Context) pulumix.Output[*Service] {
	return pulumix.Output[*Service]{
		OutputState: i.ToServiceOutputWithContext(ctx).OutputState,
	}
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

func (o ServiceOutput) ToOutput(ctx context.Context) pulumix.Output[*Service] {
	return pulumix.Output[*Service]{
		OutputState: o.OutputState,
	}
}

// HTTP strong entity tag value. Ignored if submitted
func (o ServiceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The resource kind. Only 'vm' (the default) is supported.
func (o ServiceOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o ServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource's provisioning state
func (o ServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The public key of the service, used to encrypt secrets sent to the service
func (o ServiceOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

// Service SKU
func (o ServiceOutput) Sku() ServiceSkuResponsePtrOutput {
	return o.ApplyT(func(v *Service) ServiceSkuResponsePtrOutput { return v.Sku }).(ServiceSkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Service) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The ID of the Microsoft.Network/networkInterfaces resource which the service have
func (o ServiceOutput) VirtualNicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.VirtualNicId }).(pulumi.StringPtrOutput)
}

// The ID of the Microsoft.Network/virtualNetworks/subnets resource to which the service should be joined
func (o ServiceOutput) VirtualSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.VirtualSubnetId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
