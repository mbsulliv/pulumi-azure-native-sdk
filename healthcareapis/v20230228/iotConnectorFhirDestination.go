// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230228

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IoT Connector FHIR destination definition.
type IotConnectorFhirDestination struct {
	pulumi.CustomResourceState

	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// FHIR Mappings
	FhirMapping IotMappingPropertiesResponseOutput `pulumi:"fhirMapping"`
	// Fully qualified resource id of the FHIR service to connect to.
	FhirServiceResourceId pulumi.StringOutput `pulumi:"fhirServiceResourceId"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Determines how resource identity is resolved on the destination.
	ResourceIdentityResolutionType pulumi.StringOutput `pulumi:"resourceIdentityResolutionType"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIotConnectorFhirDestination registers a new resource with the given unique name, arguments, and options.
func NewIotConnectorFhirDestination(ctx *pulumi.Context,
	name string, args *IotConnectorFhirDestinationArgs, opts ...pulumi.ResourceOption) (*IotConnectorFhirDestination, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FhirMapping == nil {
		return nil, errors.New("invalid value for required argument 'FhirMapping'")
	}
	if args.FhirServiceResourceId == nil {
		return nil, errors.New("invalid value for required argument 'FhirServiceResourceId'")
	}
	if args.IotConnectorName == nil {
		return nil, errors.New("invalid value for required argument 'IotConnectorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceIdentityResolutionType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceIdentityResolutionType'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:healthcareapis:IotConnectorFhirDestination"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20210601preview:IotConnectorFhirDestination"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20211101:IotConnectorFhirDestination"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220131preview:IotConnectorFhirDestination"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220515:IotConnectorFhirDestination"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220601:IotConnectorFhirDestination"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20221001preview:IotConnectorFhirDestination"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20221201:IotConnectorFhirDestination"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20230906:IotConnectorFhirDestination"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IotConnectorFhirDestination
	err := ctx.RegisterResource("azure-native:healthcareapis/v20230228:IotConnectorFhirDestination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotConnectorFhirDestination gets an existing IotConnectorFhirDestination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotConnectorFhirDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotConnectorFhirDestinationState, opts ...pulumi.ResourceOption) (*IotConnectorFhirDestination, error) {
	var resource IotConnectorFhirDestination
	err := ctx.ReadResource("azure-native:healthcareapis/v20230228:IotConnectorFhirDestination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotConnectorFhirDestination resources.
type iotConnectorFhirDestinationState struct {
}

type IotConnectorFhirDestinationState struct {
}

func (IotConnectorFhirDestinationState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotConnectorFhirDestinationState)(nil)).Elem()
}

type iotConnectorFhirDestinationArgs struct {
	// The name of IoT Connector FHIR destination resource.
	FhirDestinationName *string `pulumi:"fhirDestinationName"`
	// FHIR Mappings
	FhirMapping IotMappingProperties `pulumi:"fhirMapping"`
	// Fully qualified resource id of the FHIR service to connect to.
	FhirServiceResourceId string `pulumi:"fhirServiceResourceId"`
	// The name of IoT Connector resource.
	IotConnectorName string `pulumi:"iotConnectorName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Determines how resource identity is resolved on the destination.
	ResourceIdentityResolutionType string `pulumi:"resourceIdentityResolutionType"`
	// The name of workspace resource.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a IotConnectorFhirDestination resource.
type IotConnectorFhirDestinationArgs struct {
	// The name of IoT Connector FHIR destination resource.
	FhirDestinationName pulumi.StringPtrInput
	// FHIR Mappings
	FhirMapping IotMappingPropertiesInput
	// Fully qualified resource id of the FHIR service to connect to.
	FhirServiceResourceId pulumi.StringInput
	// The name of IoT Connector resource.
	IotConnectorName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput
	// Determines how resource identity is resolved on the destination.
	ResourceIdentityResolutionType pulumi.StringInput
	// The name of workspace resource.
	WorkspaceName pulumi.StringInput
}

func (IotConnectorFhirDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotConnectorFhirDestinationArgs)(nil)).Elem()
}

type IotConnectorFhirDestinationInput interface {
	pulumi.Input

	ToIotConnectorFhirDestinationOutput() IotConnectorFhirDestinationOutput
	ToIotConnectorFhirDestinationOutputWithContext(ctx context.Context) IotConnectorFhirDestinationOutput
}

func (*IotConnectorFhirDestination) ElementType() reflect.Type {
	return reflect.TypeOf((**IotConnectorFhirDestination)(nil)).Elem()
}

func (i *IotConnectorFhirDestination) ToIotConnectorFhirDestinationOutput() IotConnectorFhirDestinationOutput {
	return i.ToIotConnectorFhirDestinationOutputWithContext(context.Background())
}

func (i *IotConnectorFhirDestination) ToIotConnectorFhirDestinationOutputWithContext(ctx context.Context) IotConnectorFhirDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotConnectorFhirDestinationOutput)
}

type IotConnectorFhirDestinationOutput struct{ *pulumi.OutputState }

func (IotConnectorFhirDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotConnectorFhirDestination)(nil)).Elem()
}

func (o IotConnectorFhirDestinationOutput) ToIotConnectorFhirDestinationOutput() IotConnectorFhirDestinationOutput {
	return o
}

func (o IotConnectorFhirDestinationOutput) ToIotConnectorFhirDestinationOutputWithContext(ctx context.Context) IotConnectorFhirDestinationOutput {
	return o
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o IotConnectorFhirDestinationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// FHIR Mappings
func (o IotConnectorFhirDestinationOutput) FhirMapping() IotMappingPropertiesResponseOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) IotMappingPropertiesResponseOutput { return v.FhirMapping }).(IotMappingPropertiesResponseOutput)
}

// Fully qualified resource id of the FHIR service to connect to.
func (o IotConnectorFhirDestinationOutput) FhirServiceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) pulumi.StringOutput { return v.FhirServiceResourceId }).(pulumi.StringOutput)
}

// The resource location.
func (o IotConnectorFhirDestinationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o IotConnectorFhirDestinationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Determines how resource identity is resolved on the destination.
func (o IotConnectorFhirDestinationOutput) ResourceIdentityResolutionType() pulumi.StringOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) pulumi.StringOutput { return v.ResourceIdentityResolutionType }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o IotConnectorFhirDestinationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource type.
func (o IotConnectorFhirDestinationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IotConnectorFhirDestinationOutput{})
}
