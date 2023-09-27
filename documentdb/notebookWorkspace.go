// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package documentdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A notebook workspace resource
// Azure REST API version: 2023-04-15. Prior API version in Azure Native 1.x: 2021-03-15
type NotebookWorkspace struct {
	pulumi.CustomResourceState

	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the endpoint of Notebook server.
	NotebookServerEndpoint pulumi.StringOutput `pulumi:"notebookServerEndpoint"`
	// Status of the notebook workspace. Possible values are: Creating, Online, Deleting, Failed, Updating.
	Status pulumi.StringOutput `pulumi:"status"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNotebookWorkspace registers a new resource with the given unique name, arguments, and options.
func NewNotebookWorkspace(ctx *pulumi.Context,
	name string, args *NotebookWorkspaceArgs, opts ...pulumi.ResourceOption) (*NotebookWorkspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230301preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230415:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915preview:NotebookWorkspace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NotebookWorkspace
	err := ctx.RegisterResource("azure-native:documentdb:NotebookWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebookWorkspace gets an existing NotebookWorkspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebookWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebookWorkspaceState, opts ...pulumi.ResourceOption) (*NotebookWorkspace, error) {
	var resource NotebookWorkspace
	err := ctx.ReadResource("azure-native:documentdb:NotebookWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebookWorkspace resources.
type notebookWorkspaceState struct {
}

type NotebookWorkspaceState struct {
}

func (NotebookWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookWorkspaceState)(nil)).Elem()
}

type notebookWorkspaceArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The name of the notebook workspace resource.
	NotebookWorkspaceName *string `pulumi:"notebookWorkspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a NotebookWorkspace resource.
type NotebookWorkspaceArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// The name of the notebook workspace resource.
	NotebookWorkspaceName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (NotebookWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookWorkspaceArgs)(nil)).Elem()
}

type NotebookWorkspaceInput interface {
	pulumi.Input

	ToNotebookWorkspaceOutput() NotebookWorkspaceOutput
	ToNotebookWorkspaceOutputWithContext(ctx context.Context) NotebookWorkspaceOutput
}

func (*NotebookWorkspace) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookWorkspace)(nil)).Elem()
}

func (i *NotebookWorkspace) ToNotebookWorkspaceOutput() NotebookWorkspaceOutput {
	return i.ToNotebookWorkspaceOutputWithContext(context.Background())
}

func (i *NotebookWorkspace) ToNotebookWorkspaceOutputWithContext(ctx context.Context) NotebookWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookWorkspaceOutput)
}

func (i *NotebookWorkspace) ToOutput(ctx context.Context) pulumix.Output[*NotebookWorkspace] {
	return pulumix.Output[*NotebookWorkspace]{
		OutputState: i.ToNotebookWorkspaceOutputWithContext(ctx).OutputState,
	}
}

type NotebookWorkspaceOutput struct{ *pulumi.OutputState }

func (NotebookWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookWorkspace)(nil)).Elem()
}

func (o NotebookWorkspaceOutput) ToNotebookWorkspaceOutput() NotebookWorkspaceOutput {
	return o
}

func (o NotebookWorkspaceOutput) ToNotebookWorkspaceOutputWithContext(ctx context.Context) NotebookWorkspaceOutput {
	return o
}

func (o NotebookWorkspaceOutput) ToOutput(ctx context.Context) pulumix.Output[*NotebookWorkspace] {
	return pulumix.Output[*NotebookWorkspace]{
		OutputState: o.OutputState,
	}
}

// The name of the database account.
func (o NotebookWorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookWorkspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the endpoint of Notebook server.
func (o NotebookWorkspaceOutput) NotebookServerEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookWorkspace) pulumi.StringOutput { return v.NotebookServerEndpoint }).(pulumi.StringOutput)
}

// Status of the notebook workspace. Possible values are: Creating, Online, Deleting, Failed, Updating.
func (o NotebookWorkspaceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookWorkspace) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The type of Azure resource.
func (o NotebookWorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookWorkspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NotebookWorkspaceOutput{})
}
