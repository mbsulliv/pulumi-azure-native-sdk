// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Traffic Collector resource.
type AzureTrafficCollector struct {
	pulumi.CustomResourceState

	// Collector Policies for Azure Traffic Collector.
	CollectorPolicies ResourceReferenceResponseArrayOutput `pulumi:"collectorPolicies"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the application rule collection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData TrackedResourceResponseSystemDataOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The virtualHub to which the Azure Traffic Collector belongs.
	VirtualHub ResourceReferenceResponsePtrOutput `pulumi:"virtualHub"`
}

// NewAzureTrafficCollector registers a new resource with the given unique name, arguments, and options.
func NewAzureTrafficCollector(ctx *pulumi.Context,
	name string, args *AzureTrafficCollectorArgs, opts ...pulumi.ResourceOption) (*AzureTrafficCollector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkfunction:AzureTrafficCollector"),
		},
		{
			Type: pulumi.String("azure-native:networkfunction/v20210901preview:AzureTrafficCollector"),
		},
		{
			Type: pulumi.String("azure-native:networkfunction/v20220501:AzureTrafficCollector"),
		},
		{
			Type: pulumi.String("azure-native:networkfunction/v20220801:AzureTrafficCollector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AzureTrafficCollector
	err := ctx.RegisterResource("azure-native:networkfunction/v20221101:AzureTrafficCollector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureTrafficCollector gets an existing AzureTrafficCollector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureTrafficCollector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureTrafficCollectorState, opts ...pulumi.ResourceOption) (*AzureTrafficCollector, error) {
	var resource AzureTrafficCollector
	err := ctx.ReadResource("azure-native:networkfunction/v20221101:AzureTrafficCollector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureTrafficCollector resources.
type azureTrafficCollectorState struct {
}

type AzureTrafficCollectorState struct {
}

func (AzureTrafficCollectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureTrafficCollectorState)(nil)).Elem()
}

type azureTrafficCollectorArgs struct {
	// Azure Traffic Collector name
	AzureTrafficCollectorName *string `pulumi:"azureTrafficCollectorName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AzureTrafficCollector resource.
type AzureTrafficCollectorArgs struct {
	// Azure Traffic Collector name
	AzureTrafficCollectorName pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (AzureTrafficCollectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureTrafficCollectorArgs)(nil)).Elem()
}

type AzureTrafficCollectorInput interface {
	pulumi.Input

	ToAzureTrafficCollectorOutput() AzureTrafficCollectorOutput
	ToAzureTrafficCollectorOutputWithContext(ctx context.Context) AzureTrafficCollectorOutput
}

func (*AzureTrafficCollector) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureTrafficCollector)(nil)).Elem()
}

func (i *AzureTrafficCollector) ToAzureTrafficCollectorOutput() AzureTrafficCollectorOutput {
	return i.ToAzureTrafficCollectorOutputWithContext(context.Background())
}

func (i *AzureTrafficCollector) ToAzureTrafficCollectorOutputWithContext(ctx context.Context) AzureTrafficCollectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTrafficCollectorOutput)
}

type AzureTrafficCollectorOutput struct{ *pulumi.OutputState }

func (AzureTrafficCollectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureTrafficCollector)(nil)).Elem()
}

func (o AzureTrafficCollectorOutput) ToAzureTrafficCollectorOutput() AzureTrafficCollectorOutput {
	return o
}

func (o AzureTrafficCollectorOutput) ToAzureTrafficCollectorOutputWithContext(ctx context.Context) AzureTrafficCollectorOutput {
	return o
}

// Collector Policies for Azure Traffic Collector.
func (o AzureTrafficCollectorOutput) CollectorPolicies() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) ResourceReferenceResponseArrayOutput { return v.CollectorPolicies }).(ResourceReferenceResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o AzureTrafficCollectorOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource location.
func (o AzureTrafficCollectorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o AzureTrafficCollectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the application rule collection resource.
func (o AzureTrafficCollectorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o AzureTrafficCollectorOutput) SystemData() TrackedResourceResponseSystemDataOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) TrackedResourceResponseSystemDataOutput { return v.SystemData }).(TrackedResourceResponseSystemDataOutput)
}

// Resource tags.
func (o AzureTrafficCollectorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o AzureTrafficCollectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The virtualHub to which the Azure Traffic Collector belongs.
func (o AzureTrafficCollectorOutput) VirtualHub() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) ResourceReferenceResponsePtrOutput { return v.VirtualHub }).(ResourceReferenceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureTrafficCollectorOutput{})
}
