// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Class representing a database script.
// Azure REST API version: 2022-12-29. Prior API version in Azure Native 1.x: 2021-01-01.
//
// Other available API versions: 2021-08-27, 2023-05-02, 2023-08-15.
type Script struct {
	pulumi.CustomResourceState

	// Flag that indicates whether to continue if one of the command fails.
	ContinueOnErrors pulumi.BoolPtrOutput `pulumi:"continueOnErrors"`
	// A unique string. If changed the script will be applied again.
	ForceUpdateTag pulumi.StringPtrOutput `pulumi:"forceUpdateTag"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The url to the KQL script blob file. Must not be used together with scriptContent property
	ScriptUrl pulumi.StringPtrOutput `pulumi:"scriptUrl"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewScript registers a new resource with the given unique name, arguments, and options.
func NewScript(ctx *pulumi.Context,
	name string, args *ScriptArgs, opts ...pulumi.ResourceOption) (*Script, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ContinueOnErrors == nil {
		args.ContinueOnErrors = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto/v20210101:Script"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:Script"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:Script"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:Script"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221111:Script"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221229:Script"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230502:Script"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230815:Script"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Script
	err := ctx.RegisterResource("azure-native:kusto:Script", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScript gets an existing Script resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScriptState, opts ...pulumi.ResourceOption) (*Script, error) {
	var resource Script
	err := ctx.ReadResource("azure-native:kusto:Script", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Script resources.
type scriptState struct {
}

type ScriptState struct {
}

func (ScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptState)(nil)).Elem()
}

type scriptArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// Flag that indicates whether to continue if one of the command fails.
	ContinueOnErrors *bool `pulumi:"continueOnErrors"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// A unique string. If changed the script will be applied again.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The script content. This property should be used when the script is provide inline and not through file in a SA. Must not be used together with scriptUrl and scriptUrlSasToken properties.
	ScriptContent *string `pulumi:"scriptContent"`
	// The name of the Kusto database script.
	ScriptName *string `pulumi:"scriptName"`
	// The url to the KQL script blob file. Must not be used together with scriptContent property
	ScriptUrl *string `pulumi:"scriptUrl"`
	// The SaS token that provide read access to the file which contain the script. Must be provided when using scriptUrl property.
	ScriptUrlSasToken *string `pulumi:"scriptUrlSasToken"`
}

// The set of arguments for constructing a Script resource.
type ScriptArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// Flag that indicates whether to continue if one of the command fails.
	ContinueOnErrors pulumi.BoolPtrInput
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringInput
	// A unique string. If changed the script will be applied again.
	ForceUpdateTag pulumi.StringPtrInput
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput
	// The script content. This property should be used when the script is provide inline and not through file in a SA. Must not be used together with scriptUrl and scriptUrlSasToken properties.
	ScriptContent pulumi.StringPtrInput
	// The name of the Kusto database script.
	ScriptName pulumi.StringPtrInput
	// The url to the KQL script blob file. Must not be used together with scriptContent property
	ScriptUrl pulumi.StringPtrInput
	// The SaS token that provide read access to the file which contain the script. Must be provided when using scriptUrl property.
	ScriptUrlSasToken pulumi.StringPtrInput
}

func (ScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptArgs)(nil)).Elem()
}

type ScriptInput interface {
	pulumi.Input

	ToScriptOutput() ScriptOutput
	ToScriptOutputWithContext(ctx context.Context) ScriptOutput
}

func (*Script) ElementType() reflect.Type {
	return reflect.TypeOf((**Script)(nil)).Elem()
}

func (i *Script) ToScriptOutput() ScriptOutput {
	return i.ToScriptOutputWithContext(context.Background())
}

func (i *Script) ToScriptOutputWithContext(ctx context.Context) ScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptOutput)
}

func (i *Script) ToOutput(ctx context.Context) pulumix.Output[*Script] {
	return pulumix.Output[*Script]{
		OutputState: i.ToScriptOutputWithContext(ctx).OutputState,
	}
}

type ScriptOutput struct{ *pulumi.OutputState }

func (ScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Script)(nil)).Elem()
}

func (o ScriptOutput) ToScriptOutput() ScriptOutput {
	return o
}

func (o ScriptOutput) ToScriptOutputWithContext(ctx context.Context) ScriptOutput {
	return o
}

func (o ScriptOutput) ToOutput(ctx context.Context) pulumix.Output[*Script] {
	return pulumix.Output[*Script]{
		OutputState: o.OutputState,
	}
}

// Flag that indicates whether to continue if one of the command fails.
func (o ScriptOutput) ContinueOnErrors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Script) pulumi.BoolPtrOutput { return v.ContinueOnErrors }).(pulumi.BoolPtrOutput)
}

// A unique string. If changed the script will be applied again.
func (o ScriptOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Script) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ScriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Script) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o ScriptOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Script) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The url to the KQL script blob file. Must not be used together with scriptContent property
func (o ScriptOutput) ScriptUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Script) pulumi.StringPtrOutput { return v.ScriptUrl }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ScriptOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Script) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ScriptOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Script) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScriptOutput{})
}
