// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package solutions

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Information about managed application.
// Azure REST API version: 2021-07-01. Prior API version in Azure Native 1.x: 2019-07-01
type Application struct {
	pulumi.CustomResourceState

	// The fully qualified path of managed application definition Id.
	ApplicationDefinitionId pulumi.StringPtrOutput `pulumi:"applicationDefinitionId"`
	// The collection of managed application artifacts.
	Artifacts ApplicationArtifactResponseArrayOutput `pulumi:"artifacts"`
	// The  read-only authorizations property that is retrieved from the application package.
	Authorizations ApplicationAuthorizationResponseArrayOutput `pulumi:"authorizations"`
	// The managed application billing details.
	BillingDetails ApplicationBillingDetailsDefinitionResponseOutput `pulumi:"billingDetails"`
	// The client entity that created the JIT request.
	CreatedBy ApplicationClientDetailsResponseOutput `pulumi:"createdBy"`
	// The read-only customer support property that is retrieved from the application package.
	CustomerSupport ApplicationPackageContactResponseOutput `pulumi:"customerSupport"`
	// The identity of the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The managed application Jit access policy.
	JitAccessPolicy ApplicationJitAccessPolicyResponsePtrOutput `pulumi:"jitAccessPolicy"`
	// The kind of the managed application. Allowed values are MarketPlace and ServiceCatalog.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// ID of the resource that manages this resource.
	ManagedBy pulumi.StringPtrOutput `pulumi:"managedBy"`
	// The managed resource group Id.
	ManagedResourceGroupId pulumi.StringPtrOutput `pulumi:"managedResourceGroupId"`
	// The managed application management mode.
	ManagementMode pulumi.StringOutput `pulumi:"managementMode"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Name and value pairs that define the managed application outputs.
	Outputs pulumi.AnyOutput `pulumi:"outputs"`
	// Name and value pairs that define the managed application parameters. It can be a JObject or a well formed JSON string.
	Parameters pulumi.AnyOutput `pulumi:"parameters"`
	// The plan information.
	Plan PlanResponsePtrOutput `pulumi:"plan"`
	// The managed application provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The publisher tenant Id.
	PublisherTenantId pulumi.StringOutput `pulumi:"publisherTenantId"`
	// The SKU of the resource.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The read-only support URLs property that is retrieved from the application package.
	SupportUrls ApplicationPackageSupportUrlsResponseOutput `pulumi:"supportUrls"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The client entity that last updated the JIT request.
	UpdatedBy ApplicationClientDetailsResponseOutput `pulumi:"updatedBy"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:solutions/v20160901preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20170901:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20171201:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180201:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180301:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180601:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180901preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20190701:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20200821preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20210201preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20210701:Application"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Application
	err := ctx.RegisterResource("azure-native:solutions:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azure-native:solutions:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
}

type ApplicationState struct {
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// The fully qualified path of managed application definition Id.
	ApplicationDefinitionId *string `pulumi:"applicationDefinitionId"`
	// The name of the managed application.
	ApplicationName *string `pulumi:"applicationName"`
	// The identity of the resource.
	Identity *Identity `pulumi:"identity"`
	// The managed application Jit access policy.
	JitAccessPolicy *ApplicationJitAccessPolicy `pulumi:"jitAccessPolicy"`
	// The kind of the managed application. Allowed values are MarketPlace and ServiceCatalog.
	Kind string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// ID of the resource that manages this resource.
	ManagedBy *string `pulumi:"managedBy"`
	// The managed resource group Id.
	ManagedResourceGroupId *string `pulumi:"managedResourceGroupId"`
	// Name and value pairs that define the managed application parameters. It can be a JObject or a well formed JSON string.
	Parameters interface{} `pulumi:"parameters"`
	// The plan information.
	Plan *Plan `pulumi:"plan"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the resource.
	Sku *Sku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The fully qualified path of managed application definition Id.
	ApplicationDefinitionId pulumi.StringPtrInput
	// The name of the managed application.
	ApplicationName pulumi.StringPtrInput
	// The identity of the resource.
	Identity IdentityPtrInput
	// The managed application Jit access policy.
	JitAccessPolicy ApplicationJitAccessPolicyPtrInput
	// The kind of the managed application. Allowed values are MarketPlace and ServiceCatalog.
	Kind pulumi.StringInput
	// Resource location
	Location pulumi.StringPtrInput
	// ID of the resource that manages this resource.
	ManagedBy pulumi.StringPtrInput
	// The managed resource group Id.
	ManagedResourceGroupId pulumi.StringPtrInput
	// Name and value pairs that define the managed application parameters. It can be a JObject or a well formed JSON string.
	Parameters pulumi.Input
	// The plan information.
	Plan PlanPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The SKU of the resource.
	Sku SkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

func (i *Application) ToOutput(ctx context.Context) pulumix.Output[*Application] {
	return pulumix.Output[*Application]{
		OutputState: i.ToApplicationOutputWithContext(ctx).OutputState,
	}
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToOutput(ctx context.Context) pulumix.Output[*Application] {
	return pulumix.Output[*Application]{
		OutputState: o.OutputState,
	}
}

// The fully qualified path of managed application definition Id.
func (o ApplicationOutput) ApplicationDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.ApplicationDefinitionId }).(pulumi.StringPtrOutput)
}

// The collection of managed application artifacts.
func (o ApplicationOutput) Artifacts() ApplicationArtifactResponseArrayOutput {
	return o.ApplyT(func(v *Application) ApplicationArtifactResponseArrayOutput { return v.Artifacts }).(ApplicationArtifactResponseArrayOutput)
}

// The  read-only authorizations property that is retrieved from the application package.
func (o ApplicationOutput) Authorizations() ApplicationAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v *Application) ApplicationAuthorizationResponseArrayOutput { return v.Authorizations }).(ApplicationAuthorizationResponseArrayOutput)
}

// The managed application billing details.
func (o ApplicationOutput) BillingDetails() ApplicationBillingDetailsDefinitionResponseOutput {
	return o.ApplyT(func(v *Application) ApplicationBillingDetailsDefinitionResponseOutput { return v.BillingDetails }).(ApplicationBillingDetailsDefinitionResponseOutput)
}

// The client entity that created the JIT request.
func (o ApplicationOutput) CreatedBy() ApplicationClientDetailsResponseOutput {
	return o.ApplyT(func(v *Application) ApplicationClientDetailsResponseOutput { return v.CreatedBy }).(ApplicationClientDetailsResponseOutput)
}

// The read-only customer support property that is retrieved from the application package.
func (o ApplicationOutput) CustomerSupport() ApplicationPackageContactResponseOutput {
	return o.ApplyT(func(v *Application) ApplicationPackageContactResponseOutput { return v.CustomerSupport }).(ApplicationPackageContactResponseOutput)
}

// The identity of the resource.
func (o ApplicationOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Application) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The managed application Jit access policy.
func (o ApplicationOutput) JitAccessPolicy() ApplicationJitAccessPolicyResponsePtrOutput {
	return o.ApplyT(func(v *Application) ApplicationJitAccessPolicyResponsePtrOutput { return v.JitAccessPolicy }).(ApplicationJitAccessPolicyResponsePtrOutput)
}

// The kind of the managed application. Allowed values are MarketPlace and ServiceCatalog.
func (o ApplicationOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource location
func (o ApplicationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// ID of the resource that manages this resource.
func (o ApplicationOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// The managed resource group Id.
func (o ApplicationOutput) ManagedResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.ManagedResourceGroupId }).(pulumi.StringPtrOutput)
}

// The managed application management mode.
func (o ApplicationOutput) ManagementMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ManagementMode }).(pulumi.StringOutput)
}

// Resource name
func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name and value pairs that define the managed application outputs.
func (o ApplicationOutput) Outputs() pulumi.AnyOutput {
	return o.ApplyT(func(v *Application) pulumi.AnyOutput { return v.Outputs }).(pulumi.AnyOutput)
}

// Name and value pairs that define the managed application parameters. It can be a JObject or a well formed JSON string.
func (o ApplicationOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *Application) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

// The plan information.
func (o ApplicationOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v *Application) PlanResponsePtrOutput { return v.Plan }).(PlanResponsePtrOutput)
}

// The managed application provisioning state.
func (o ApplicationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The publisher tenant Id.
func (o ApplicationOutput) PublisherTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.PublisherTenantId }).(pulumi.StringOutput)
}

// The SKU of the resource.
func (o ApplicationOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Application) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The read-only support URLs property that is retrieved from the application package.
func (o ApplicationOutput) SupportUrls() ApplicationPackageSupportUrlsResponseOutput {
	return o.ApplyT(func(v *Application) ApplicationPackageSupportUrlsResponseOutput { return v.SupportUrls }).(ApplicationPackageSupportUrlsResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ApplicationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Application) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o ApplicationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Application) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o ApplicationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The client entity that last updated the JIT request.
func (o ApplicationOutput) UpdatedBy() ApplicationClientDetailsResponseOutput {
	return o.ApplyT(func(v *Application) ApplicationClientDetailsResponseOutput { return v.UpdatedBy }).(ApplicationClientDetailsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
}
