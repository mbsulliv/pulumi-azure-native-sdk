// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Cloud Connector resource.
type CloudConnector struct {
	pulumi.CustomResourceState

	// Account identifier of the remote cloud.
	AccountId pulumi.StringPtrOutput `pulumi:"accountId"`
	// The cloud connector type.
	CloudType pulumi.StringPtrOutput `pulumi:"cloudType"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the cloud connector resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCloudConnector registers a new resource with the given unique name, arguments, and options.
func NewCloudConnector(ctx *pulumi.Context,
	name string, args *CloudConnectorArgs, opts ...pulumi.ResourceOption) (*CloudConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcloud:CloudConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CloudConnector
	err := ctx.RegisterResource("azure-native:hybridcloud/v20230101preview:CloudConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudConnector gets an existing CloudConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudConnectorState, opts ...pulumi.ResourceOption) (*CloudConnector, error) {
	var resource CloudConnector
	err := ctx.ReadResource("azure-native:hybridcloud/v20230101preview:CloudConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudConnector resources.
type cloudConnectorState struct {
}

type CloudConnectorState struct {
}

func (CloudConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudConnectorState)(nil)).Elem()
}

type cloudConnectorArgs struct {
	// Account identifier of the remote cloud.
	AccountId *string `pulumi:"accountId"`
	// The name of the cloud connector resource
	CloudConnectorName *string `pulumi:"cloudConnectorName"`
	// The cloud connector type.
	CloudType *string `pulumi:"cloudType"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CloudConnector resource.
type CloudConnectorArgs struct {
	// Account identifier of the remote cloud.
	AccountId pulumi.StringPtrInput
	// The name of the cloud connector resource
	CloudConnectorName pulumi.StringPtrInput
	// The cloud connector type.
	CloudType pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (CloudConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudConnectorArgs)(nil)).Elem()
}

type CloudConnectorInput interface {
	pulumi.Input

	ToCloudConnectorOutput() CloudConnectorOutput
	ToCloudConnectorOutputWithContext(ctx context.Context) CloudConnectorOutput
}

func (*CloudConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudConnector)(nil)).Elem()
}

func (i *CloudConnector) ToCloudConnectorOutput() CloudConnectorOutput {
	return i.ToCloudConnectorOutputWithContext(context.Background())
}

func (i *CloudConnector) ToCloudConnectorOutputWithContext(ctx context.Context) CloudConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudConnectorOutput)
}

func (i *CloudConnector) ToOutput(ctx context.Context) pulumix.Output[*CloudConnector] {
	return pulumix.Output[*CloudConnector]{
		OutputState: i.ToCloudConnectorOutputWithContext(ctx).OutputState,
	}
}

type CloudConnectorOutput struct{ *pulumi.OutputState }

func (CloudConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudConnector)(nil)).Elem()
}

func (o CloudConnectorOutput) ToCloudConnectorOutput() CloudConnectorOutput {
	return o
}

func (o CloudConnectorOutput) ToCloudConnectorOutputWithContext(ctx context.Context) CloudConnectorOutput {
	return o
}

func (o CloudConnectorOutput) ToOutput(ctx context.Context) pulumix.Output[*CloudConnector] {
	return pulumix.Output[*CloudConnector]{
		OutputState: o.OutputState,
	}
}

// Account identifier of the remote cloud.
func (o CloudConnectorOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudConnector) pulumi.StringPtrOutput { return v.AccountId }).(pulumi.StringPtrOutput)
}

// The cloud connector type.
func (o CloudConnectorOutput) CloudType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudConnector) pulumi.StringPtrOutput { return v.CloudType }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o CloudConnectorOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudConnector) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o CloudConnectorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudConnector) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o CloudConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the cloud connector resource.
func (o CloudConnectorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudConnector) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o CloudConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CloudConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o CloudConnectorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloudConnector) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CloudConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudConnectorOutput{})
}
