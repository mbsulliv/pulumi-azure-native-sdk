// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deploymentmanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The resource representation of a service topology.
// Azure REST API version: 2019-11-01-preview. Prior API version in Azure Native 1.x: 2019-11-01-preview.
type ServiceTopology struct {
	pulumi.CustomResourceState

	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceId pulumi.StringPtrOutput `pulumi:"artifactSourceId"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceTopology registers a new resource with the given unique name, arguments, and options.
func NewServiceTopology(ctx *pulumi.Context,
	name string, args *ServiceTopologyArgs, opts ...pulumi.ResourceOption) (*ServiceTopology, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:deploymentmanager/v20180901preview:ServiceTopology"),
		},
		{
			Type: pulumi.String("azure-native:deploymentmanager/v20191101preview:ServiceTopology"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServiceTopology
	err := ctx.RegisterResource("azure-native:deploymentmanager:ServiceTopology", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceTopology gets an existing ServiceTopology resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceTopology(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceTopologyState, opts ...pulumi.ResourceOption) (*ServiceTopology, error) {
	var resource ServiceTopology
	err := ctx.ReadResource("azure-native:deploymentmanager:ServiceTopology", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceTopology resources.
type serviceTopologyState struct {
}

type ServiceTopologyState struct {
}

func (ServiceTopologyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTopologyState)(nil)).Elem()
}

type serviceTopologyArgs struct {
	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceId *string `pulumi:"artifactSourceId"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service topology .
	ServiceTopologyName *string `pulumi:"serviceTopologyName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ServiceTopology resource.
type ServiceTopologyArgs struct {
	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceId pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the service topology .
	ServiceTopologyName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ServiceTopologyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTopologyArgs)(nil)).Elem()
}

type ServiceTopologyInput interface {
	pulumi.Input

	ToServiceTopologyOutput() ServiceTopologyOutput
	ToServiceTopologyOutputWithContext(ctx context.Context) ServiceTopologyOutput
}

func (*ServiceTopology) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTopology)(nil)).Elem()
}

func (i *ServiceTopology) ToServiceTopologyOutput() ServiceTopologyOutput {
	return i.ToServiceTopologyOutputWithContext(context.Background())
}

func (i *ServiceTopology) ToServiceTopologyOutputWithContext(ctx context.Context) ServiceTopologyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTopologyOutput)
}

func (i *ServiceTopology) ToOutput(ctx context.Context) pulumix.Output[*ServiceTopology] {
	return pulumix.Output[*ServiceTopology]{
		OutputState: i.ToServiceTopologyOutputWithContext(ctx).OutputState,
	}
}

type ServiceTopologyOutput struct{ *pulumi.OutputState }

func (ServiceTopologyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTopology)(nil)).Elem()
}

func (o ServiceTopologyOutput) ToServiceTopologyOutput() ServiceTopologyOutput {
	return o
}

func (o ServiceTopologyOutput) ToServiceTopologyOutputWithContext(ctx context.Context) ServiceTopologyOutput {
	return o
}

func (o ServiceTopologyOutput) ToOutput(ctx context.Context) pulumix.Output[*ServiceTopology] {
	return pulumix.Output[*ServiceTopology]{
		OutputState: o.OutputState,
	}
}

// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
func (o ServiceTopologyOutput) ArtifactSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceTopology) pulumi.StringPtrOutput { return v.ArtifactSourceId }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o ServiceTopologyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceTopology) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ServiceTopologyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceTopology) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource tags.
func (o ServiceTopologyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceTopology) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ServiceTopologyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceTopology) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceTopologyOutput{})
}
