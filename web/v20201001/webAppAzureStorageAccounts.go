// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// AzureStorageInfo dictionary resource.
type WebAppAzureStorageAccounts struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure storage accounts.
	Properties AzureStorageInfoValueResponseMapOutput `pulumi:"properties"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppAzureStorageAccounts registers a new resource with the given unique name, arguments, and options.
func NewWebAppAzureStorageAccounts(ctx *pulumi.Context,
	name string, args *WebAppAzureStorageAccountsArgs, opts ...pulumi.ResourceOption) (*WebAppAzureStorageAccounts, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppAzureStorageAccounts"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppAzureStorageAccounts
	err := ctx.RegisterResource("azure-native:web/v20201001:WebAppAzureStorageAccounts", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppAzureStorageAccounts gets an existing WebAppAzureStorageAccounts resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppAzureStorageAccounts(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppAzureStorageAccountsState, opts ...pulumi.ResourceOption) (*WebAppAzureStorageAccounts, error) {
	var resource WebAppAzureStorageAccounts
	err := ctx.ReadResource("azure-native:web/v20201001:WebAppAzureStorageAccounts", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppAzureStorageAccounts resources.
type webAppAzureStorageAccountsState struct {
}

type WebAppAzureStorageAccountsState struct {
}

func (WebAppAzureStorageAccountsState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAzureStorageAccountsState)(nil)).Elem()
}

type webAppAzureStorageAccountsArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Azure storage accounts.
	Properties map[string]AzureStorageInfoValue `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WebAppAzureStorageAccounts resource.
type WebAppAzureStorageAccountsArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Azure storage accounts.
	Properties AzureStorageInfoValueMapInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (WebAppAzureStorageAccountsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAzureStorageAccountsArgs)(nil)).Elem()
}

type WebAppAzureStorageAccountsInput interface {
	pulumi.Input

	ToWebAppAzureStorageAccountsOutput() WebAppAzureStorageAccountsOutput
	ToWebAppAzureStorageAccountsOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsOutput
}

func (*WebAppAzureStorageAccounts) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAzureStorageAccounts)(nil)).Elem()
}

func (i *WebAppAzureStorageAccounts) ToWebAppAzureStorageAccountsOutput() WebAppAzureStorageAccountsOutput {
	return i.ToWebAppAzureStorageAccountsOutputWithContext(context.Background())
}

func (i *WebAppAzureStorageAccounts) ToWebAppAzureStorageAccountsOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppAzureStorageAccountsOutput)
}

func (i *WebAppAzureStorageAccounts) ToOutput(ctx context.Context) pulumix.Output[*WebAppAzureStorageAccounts] {
	return pulumix.Output[*WebAppAzureStorageAccounts]{
		OutputState: i.ToWebAppAzureStorageAccountsOutputWithContext(ctx).OutputState,
	}
}

type WebAppAzureStorageAccountsOutput struct{ *pulumi.OutputState }

func (WebAppAzureStorageAccountsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAzureStorageAccounts)(nil)).Elem()
}

func (o WebAppAzureStorageAccountsOutput) ToWebAppAzureStorageAccountsOutput() WebAppAzureStorageAccountsOutput {
	return o
}

func (o WebAppAzureStorageAccountsOutput) ToWebAppAzureStorageAccountsOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsOutput {
	return o
}

func (o WebAppAzureStorageAccountsOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppAzureStorageAccounts] {
	return pulumix.Output[*WebAppAzureStorageAccounts]{
		OutputState: o.OutputState,
	}
}

// Kind of resource.
func (o WebAppAzureStorageAccountsOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccounts) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppAzureStorageAccountsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccounts) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure storage accounts.
func (o WebAppAzureStorageAccountsOutput) Properties() AzureStorageInfoValueResponseMapOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccounts) AzureStorageInfoValueResponseMapOutput { return v.Properties }).(AzureStorageInfoValueResponseMapOutput)
}

// The system metadata relating to this resource.
func (o WebAppAzureStorageAccountsOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccounts) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o WebAppAzureStorageAccountsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccounts) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppAzureStorageAccountsOutput{})
}
