// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The description of Dicom Service
type DicomService struct {
	pulumi.CustomResourceState

	// Dicom Service authentication configuration.
	AuthenticationConfiguration DicomServiceAuthenticationConfigurationResponsePtrOutput `pulumi:"authenticationConfiguration"`
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ServiceManagedIdentityResponseIdentityPtrOutput `pulumi:"identity"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of private endpoint connections that are set up for this resource.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess pulumi.StringOutput `pulumi:"publicNetworkAccess"`
	// The url of the Dicom Services.
	ServiceUrl pulumi.StringOutput `pulumi:"serviceUrl"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDicomService registers a new resource with the given unique name, arguments, and options.
func NewDicomService(ctx *pulumi.Context,
	name string, args *DicomServiceArgs, opts ...pulumi.ResourceOption) (*DicomService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:healthcareapis:DicomService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20210601preview:DicomService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220131preview:DicomService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220515:DicomService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220601:DicomService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20221001preview:DicomService"),
		},
	})
	opts = append(opts, aliases)
	var resource DicomService
	err := ctx.RegisterResource("azure-native:healthcareapis/v20211101:DicomService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDicomService gets an existing DicomService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDicomService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DicomServiceState, opts ...pulumi.ResourceOption) (*DicomService, error) {
	var resource DicomService
	err := ctx.ReadResource("azure-native:healthcareapis/v20211101:DicomService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DicomService resources.
type dicomServiceState struct {
}

type DicomServiceState struct {
}

func (DicomServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomServiceState)(nil)).Elem()
}

type dicomServiceArgs struct {
	// The name of DICOM Service resource.
	DicomServiceName *string `pulumi:"dicomServiceName"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityIdentity `pulumi:"identity"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of workspace resource.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a DicomService resource.
type DicomServiceArgs struct {
	// The name of DICOM Service resource.
	DicomServiceName pulumi.StringPtrInput
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ServiceManagedIdentityIdentityPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of workspace resource.
	WorkspaceName pulumi.StringInput
}

func (DicomServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomServiceArgs)(nil)).Elem()
}

type DicomServiceInput interface {
	pulumi.Input

	ToDicomServiceOutput() DicomServiceOutput
	ToDicomServiceOutputWithContext(ctx context.Context) DicomServiceOutput
}

func (*DicomService) ElementType() reflect.Type {
	return reflect.TypeOf((**DicomService)(nil)).Elem()
}

func (i *DicomService) ToDicomServiceOutput() DicomServiceOutput {
	return i.ToDicomServiceOutputWithContext(context.Background())
}

func (i *DicomService) ToDicomServiceOutputWithContext(ctx context.Context) DicomServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DicomServiceOutput)
}

type DicomServiceOutput struct{ *pulumi.OutputState }

func (DicomServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DicomService)(nil)).Elem()
}

func (o DicomServiceOutput) ToDicomServiceOutput() DicomServiceOutput {
	return o
}

func (o DicomServiceOutput) ToDicomServiceOutputWithContext(ctx context.Context) DicomServiceOutput {
	return o
}

// Dicom Service authentication configuration.
func (o DicomServiceOutput) AuthenticationConfiguration() DicomServiceAuthenticationConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *DicomService) DicomServiceAuthenticationConfigurationResponsePtrOutput {
		return v.AuthenticationConfiguration
	}).(DicomServiceAuthenticationConfigurationResponsePtrOutput)
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o DicomServiceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Setting indicating whether the service has a managed identity associated with it.
func (o DicomServiceOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v *DicomService) ServiceManagedIdentityResponseIdentityPtrOutput { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

// The resource location.
func (o DicomServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o DicomServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of private endpoint connections that are set up for this resource.
func (o DicomServiceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *DicomService) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// The provisioning state.
func (o DicomServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
func (o DicomServiceOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringOutput { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

// The url of the Dicom Services.
func (o DicomServiceOutput) ServiceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringOutput { return v.ServiceUrl }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o DicomServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DicomService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o DicomServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o DicomServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DicomServiceOutput{})
}
