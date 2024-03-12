// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure SQL job agent.
type JobAgent struct {
	pulumi.CustomResourceState

	// Resource ID of the database to store job metadata in.
	DatabaseId pulumi.StringOutput `pulumi:"databaseId"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name and tier of the SKU.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The state of the job agent.
	State pulumi.StringOutput `pulumi:"state"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJobAgent registers a new resource with the given unique name, arguments, and options.
func NewJobAgent(ctx *pulumi.Context,
	name string, args *JobAgentArgs, opts ...pulumi.ResourceOption) (*JobAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseId == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230801preview:JobAgent"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource JobAgent
	err := ctx.RegisterResource("azure-native:sql/v20211101:JobAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobAgent gets an existing JobAgent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobAgentState, opts ...pulumi.ResourceOption) (*JobAgent, error) {
	var resource JobAgent
	err := ctx.ReadResource("azure-native:sql/v20211101:JobAgent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobAgent resources.
type jobAgentState struct {
}

type JobAgentState struct {
}

func (JobAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobAgentState)(nil)).Elem()
}

type jobAgentArgs struct {
	// Resource ID of the database to store job metadata in.
	DatabaseId string `pulumi:"databaseId"`
	// The name of the job agent to be created or updated.
	JobAgentName *string `pulumi:"jobAgentName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name and tier of the SKU.
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a JobAgent resource.
type JobAgentArgs struct {
	// Resource ID of the database to store job metadata in.
	DatabaseId pulumi.StringInput
	// The name of the job agent to be created or updated.
	JobAgentName pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The name and tier of the SKU.
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (JobAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobAgentArgs)(nil)).Elem()
}

type JobAgentInput interface {
	pulumi.Input

	ToJobAgentOutput() JobAgentOutput
	ToJobAgentOutputWithContext(ctx context.Context) JobAgentOutput
}

func (*JobAgent) ElementType() reflect.Type {
	return reflect.TypeOf((**JobAgent)(nil)).Elem()
}

func (i *JobAgent) ToJobAgentOutput() JobAgentOutput {
	return i.ToJobAgentOutputWithContext(context.Background())
}

func (i *JobAgent) ToJobAgentOutputWithContext(ctx context.Context) JobAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobAgentOutput)
}

type JobAgentOutput struct{ *pulumi.OutputState }

func (JobAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobAgent)(nil)).Elem()
}

func (o JobAgentOutput) ToJobAgentOutput() JobAgentOutput {
	return o
}

func (o JobAgentOutput) ToJobAgentOutputWithContext(ctx context.Context) JobAgentOutput {
	return o
}

// Resource ID of the database to store job metadata in.
func (o JobAgentOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobAgent) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

// Resource location.
func (o JobAgentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *JobAgent) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o JobAgentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobAgent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name and tier of the SKU.
func (o JobAgentOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *JobAgent) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The state of the job agent.
func (o JobAgentOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *JobAgent) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Resource tags.
func (o JobAgentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *JobAgent) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o JobAgentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JobAgent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobAgentOutput{})
}
