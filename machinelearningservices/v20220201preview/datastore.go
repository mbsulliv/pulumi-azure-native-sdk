// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
type Datastore struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	DatastoreDetails pulumi.AnyOutput `pulumi:"datastoreDetails"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatastore registers a new resource with the given unique name, arguments, and options.
func NewDatastore(ctx *pulumi.Context,
	name string, args *DatastoreArgs, opts ...pulumi.ResourceOption) (*Datastore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatastoreDetails == nil {
		return nil, errors.New("invalid value for required argument 'DatastoreDetails'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:Datastore"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Datastore
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220201preview:Datastore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatastore gets an existing Datastore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatastore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatastoreState, opts ...pulumi.ResourceOption) (*Datastore, error) {
	var resource Datastore
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220201preview:Datastore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Datastore resources.
type datastoreState struct {
}

type DatastoreState struct {
}

func (DatastoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreState)(nil)).Elem()
}

type datastoreArgs struct {
	// [Required] Additional attributes of the entity.
	DatastoreDetails interface{} `pulumi:"datastoreDetails"`
	// Datastore name.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Flag to skip validation.
	SkipValidation *bool `pulumi:"skipValidation"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Datastore resource.
type DatastoreArgs struct {
	// [Required] Additional attributes of the entity.
	DatastoreDetails pulumi.Input
	// Datastore name.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Flag to skip validation.
	SkipValidation pulumi.BoolPtrInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (DatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreArgs)(nil)).Elem()
}

type DatastoreInput interface {
	pulumi.Input

	ToDatastoreOutput() DatastoreOutput
	ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput
}

func (*Datastore) ElementType() reflect.Type {
	return reflect.TypeOf((**Datastore)(nil)).Elem()
}

func (i *Datastore) ToDatastoreOutput() DatastoreOutput {
	return i.ToDatastoreOutputWithContext(context.Background())
}

func (i *Datastore) ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastoreOutput)
}

func (i *Datastore) ToOutput(ctx context.Context) pulumix.Output[*Datastore] {
	return pulumix.Output[*Datastore]{
		OutputState: i.ToDatastoreOutputWithContext(ctx).OutputState,
	}
}

type DatastoreOutput struct{ *pulumi.OutputState }

func (DatastoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Datastore)(nil)).Elem()
}

func (o DatastoreOutput) ToDatastoreOutput() DatastoreOutput {
	return o
}

func (o DatastoreOutput) ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput {
	return o
}

func (o DatastoreOutput) ToOutput(ctx context.Context) pulumix.Output[*Datastore] {
	return pulumix.Output[*Datastore]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o DatastoreOutput) DatastoreDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v *Datastore) pulumi.AnyOutput { return v.DatastoreDetails }).(pulumi.AnyOutput)
}

// The name of the resource
func (o DatastoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DatastoreOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Datastore) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DatastoreOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatastoreOutput{})
}
