// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230331

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A class representing a CommunicationService resource.
type CommunicationService struct {
	pulumi.CustomResourceState

	// The location where the communication service stores its data at rest.
	DataLocation pulumi.StringOutput `pulumi:"dataLocation"`
	// FQDN of the CommunicationService instance.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The immutable resource Id of the communication service.
	ImmutableResourceId pulumi.StringOutput `pulumi:"immutableResourceId"`
	// List of email Domain resource Ids.
	LinkedDomains pulumi.StringArrayOutput `pulumi:"linkedDomains"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource ID of an Azure Notification Hub linked to this resource.
	NotificationHubId pulumi.StringOutput `pulumi:"notificationHubId"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the CommunicationService resource. Probably you need the same or higher version of client SDKs.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewCommunicationService registers a new resource with the given unique name, arguments, and options.
func NewCommunicationService(ctx *pulumi.Context,
	name string, args *CommunicationServiceArgs, opts ...pulumi.ResourceOption) (*CommunicationService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataLocation == nil {
		return nil, errors.New("invalid value for required argument 'DataLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:communication:CommunicationService"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20200820:CommunicationService"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20200820preview:CommunicationService"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20211001preview:CommunicationService"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20220701preview:CommunicationService"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20230301preview:CommunicationService"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20230401:CommunicationService"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20230401preview:CommunicationService"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20230601preview:CommunicationService"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CommunicationService
	err := ctx.RegisterResource("azure-native:communication/v20230331:CommunicationService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommunicationService gets an existing CommunicationService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommunicationService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommunicationServiceState, opts ...pulumi.ResourceOption) (*CommunicationService, error) {
	var resource CommunicationService
	err := ctx.ReadResource("azure-native:communication/v20230331:CommunicationService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CommunicationService resources.
type communicationServiceState struct {
}

type CommunicationServiceState struct {
}

func (CommunicationServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*communicationServiceState)(nil)).Elem()
}

type communicationServiceArgs struct {
	// The name of the CommunicationService resource.
	CommunicationServiceName *string `pulumi:"communicationServiceName"`
	// The location where the communication service stores its data at rest.
	DataLocation string `pulumi:"dataLocation"`
	// List of email Domain resource Ids.
	LinkedDomains []string `pulumi:"linkedDomains"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CommunicationService resource.
type CommunicationServiceArgs struct {
	// The name of the CommunicationService resource.
	CommunicationServiceName pulumi.StringPtrInput
	// The location where the communication service stores its data at rest.
	DataLocation pulumi.StringInput
	// List of email Domain resource Ids.
	LinkedDomains pulumi.StringArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (CommunicationServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*communicationServiceArgs)(nil)).Elem()
}

type CommunicationServiceInput interface {
	pulumi.Input

	ToCommunicationServiceOutput() CommunicationServiceOutput
	ToCommunicationServiceOutputWithContext(ctx context.Context) CommunicationServiceOutput
}

func (*CommunicationService) ElementType() reflect.Type {
	return reflect.TypeOf((**CommunicationService)(nil)).Elem()
}

func (i *CommunicationService) ToCommunicationServiceOutput() CommunicationServiceOutput {
	return i.ToCommunicationServiceOutputWithContext(context.Background())
}

func (i *CommunicationService) ToCommunicationServiceOutputWithContext(ctx context.Context) CommunicationServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommunicationServiceOutput)
}

type CommunicationServiceOutput struct{ *pulumi.OutputState }

func (CommunicationServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommunicationService)(nil)).Elem()
}

func (o CommunicationServiceOutput) ToCommunicationServiceOutput() CommunicationServiceOutput {
	return o
}

func (o CommunicationServiceOutput) ToCommunicationServiceOutputWithContext(ctx context.Context) CommunicationServiceOutput {
	return o
}

// The location where the communication service stores its data at rest.
func (o CommunicationServiceOutput) DataLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationService) pulumi.StringOutput { return v.DataLocation }).(pulumi.StringOutput)
}

// FQDN of the CommunicationService instance.
func (o CommunicationServiceOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationService) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// The immutable resource Id of the communication service.
func (o CommunicationServiceOutput) ImmutableResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationService) pulumi.StringOutput { return v.ImmutableResourceId }).(pulumi.StringOutput)
}

// List of email Domain resource Ids.
func (o CommunicationServiceOutput) LinkedDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CommunicationService) pulumi.StringArrayOutput { return v.LinkedDomains }).(pulumi.StringArrayOutput)
}

// The geo-location where the resource lives
func (o CommunicationServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationService) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o CommunicationServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource ID of an Azure Notification Hub linked to this resource.
func (o CommunicationServiceOutput) NotificationHubId() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationService) pulumi.StringOutput { return v.NotificationHubId }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o CommunicationServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o CommunicationServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CommunicationService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o CommunicationServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CommunicationService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CommunicationServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Version of the CommunicationService resource. Probably you need the same or higher version of client SDKs.
func (o CommunicationServiceOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationService) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CommunicationServiceOutput{})
}
