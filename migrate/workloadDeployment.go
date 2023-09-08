// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Workload deployment model.
// Azure REST API version: 2022-05-01-preview.
type WorkloadDeployment struct {
	pulumi.CustomResourceState

	// Gets or sets the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Workload deployment model properties.
	Properties WorkloadDeploymentModelPropertiesResponseOutput `pulumi:"properties"`
	SystemData WorkloadDeploymentModelResponseSystemDataOutput `pulumi:"systemData"`
	// Gets or sets the resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets or sets the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkloadDeployment registers a new resource with the given unique name, arguments, and options.
func NewWorkloadDeployment(ctx *pulumi.Context,
	name string, args *WorkloadDeploymentArgs, opts ...pulumi.ResourceOption) (*WorkloadDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ModernizeProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ModernizeProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20220501preview:WorkloadDeployment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkloadDeployment
	err := ctx.RegisterResource("azure-native:migrate:WorkloadDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadDeployment gets an existing WorkloadDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadDeploymentState, opts ...pulumi.ResourceOption) (*WorkloadDeployment, error) {
	var resource WorkloadDeployment
	err := ctx.ReadResource("azure-native:migrate:WorkloadDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadDeployment resources.
type workloadDeploymentState struct {
}

type WorkloadDeploymentState struct {
}

func (WorkloadDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadDeploymentState)(nil)).Elem()
}

type workloadDeploymentArgs struct {
	// ModernizeProject name.
	ModernizeProjectName string `pulumi:"modernizeProjectName"`
	// Workload deployment model properties.
	Properties *WorkloadDeploymentModelProperties `pulumi:"properties"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure Subscription Id in which project was created.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Gets or sets the resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Workload deployment name.
	WorkloadDeploymentName *string `pulumi:"workloadDeploymentName"`
}

// The set of arguments for constructing a WorkloadDeployment resource.
type WorkloadDeploymentArgs struct {
	// ModernizeProject name.
	ModernizeProjectName pulumi.StringInput
	// Workload deployment model properties.
	Properties WorkloadDeploymentModelPropertiesPtrInput
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput
	// Azure Subscription Id in which project was created.
	SubscriptionId pulumi.StringPtrInput
	// Gets or sets the resource tags.
	Tags pulumi.StringMapInput
	// Workload deployment name.
	WorkloadDeploymentName pulumi.StringPtrInput
}

func (WorkloadDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadDeploymentArgs)(nil)).Elem()
}

type WorkloadDeploymentInput interface {
	pulumi.Input

	ToWorkloadDeploymentOutput() WorkloadDeploymentOutput
	ToWorkloadDeploymentOutputWithContext(ctx context.Context) WorkloadDeploymentOutput
}

func (*WorkloadDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadDeployment)(nil)).Elem()
}

func (i *WorkloadDeployment) ToWorkloadDeploymentOutput() WorkloadDeploymentOutput {
	return i.ToWorkloadDeploymentOutputWithContext(context.Background())
}

func (i *WorkloadDeployment) ToWorkloadDeploymentOutputWithContext(ctx context.Context) WorkloadDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadDeploymentOutput)
}

type WorkloadDeploymentOutput struct{ *pulumi.OutputState }

func (WorkloadDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadDeployment)(nil)).Elem()
}

func (o WorkloadDeploymentOutput) ToWorkloadDeploymentOutput() WorkloadDeploymentOutput {
	return o
}

func (o WorkloadDeploymentOutput) ToWorkloadDeploymentOutputWithContext(ctx context.Context) WorkloadDeploymentOutput {
	return o
}

// Gets or sets the name of the resource.
func (o WorkloadDeploymentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadDeployment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Workload deployment model properties.
func (o WorkloadDeploymentOutput) Properties() WorkloadDeploymentModelPropertiesResponseOutput {
	return o.ApplyT(func(v *WorkloadDeployment) WorkloadDeploymentModelPropertiesResponseOutput { return v.Properties }).(WorkloadDeploymentModelPropertiesResponseOutput)
}

func (o WorkloadDeploymentOutput) SystemData() WorkloadDeploymentModelResponseSystemDataOutput {
	return o.ApplyT(func(v *WorkloadDeployment) WorkloadDeploymentModelResponseSystemDataOutput { return v.SystemData }).(WorkloadDeploymentModelResponseSystemDataOutput)
}

// Gets or sets the resource tags.
func (o WorkloadDeploymentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WorkloadDeployment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o WorkloadDeploymentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadDeployment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadDeploymentOutput{})
}