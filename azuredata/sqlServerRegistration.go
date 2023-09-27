// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredata

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A SQL server registration.
// Azure REST API version: 2019-07-24-preview. Prior API version in Azure Native 1.x: 2019-07-24-preview
type SqlServerRegistration struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional Properties as JSON string
	PropertyBag pulumi.StringPtrOutput `pulumi:"propertyBag"`
	// Resource Group Name
	ResourceGroup pulumi.StringPtrOutput `pulumi:"resourceGroup"`
	// Subscription Id
	SubscriptionId pulumi.StringPtrOutput `pulumi:"subscriptionId"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlServerRegistration registers a new resource with the given unique name, arguments, and options.
func NewSqlServerRegistration(ctx *pulumi.Context,
	name string, args *SqlServerRegistrationArgs, opts ...pulumi.ResourceOption) (*SqlServerRegistration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azuredata/v20170301preview:SqlServerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:azuredata/v20190724preview:SqlServerRegistration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SqlServerRegistration
	err := ctx.RegisterResource("azure-native:azuredata:SqlServerRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlServerRegistration gets an existing SqlServerRegistration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlServerRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlServerRegistrationState, opts ...pulumi.ResourceOption) (*SqlServerRegistration, error) {
	var resource SqlServerRegistration
	err := ctx.ReadResource("azure-native:azuredata:SqlServerRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlServerRegistration resources.
type sqlServerRegistrationState struct {
}

type SqlServerRegistrationState struct {
}

func (SqlServerRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerRegistrationState)(nil)).Elem()
}

type sqlServerRegistrationArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Optional Properties as JSON string
	PropertyBag *string `pulumi:"propertyBag"`
	// Resource Group Name
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the SQL Server registration.
	SqlServerRegistrationName *string `pulumi:"sqlServerRegistrationName"`
	// Subscription Id
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SqlServerRegistration resource.
type SqlServerRegistrationArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Optional Properties as JSON string
	PropertyBag pulumi.StringPtrInput
	// Resource Group Name
	ResourceGroup pulumi.StringPtrInput
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Name of the SQL Server registration.
	SqlServerRegistrationName pulumi.StringPtrInput
	// Subscription Id
	SubscriptionId pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SqlServerRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerRegistrationArgs)(nil)).Elem()
}

type SqlServerRegistrationInput interface {
	pulumi.Input

	ToSqlServerRegistrationOutput() SqlServerRegistrationOutput
	ToSqlServerRegistrationOutputWithContext(ctx context.Context) SqlServerRegistrationOutput
}

func (*SqlServerRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerRegistration)(nil)).Elem()
}

func (i *SqlServerRegistration) ToSqlServerRegistrationOutput() SqlServerRegistrationOutput {
	return i.ToSqlServerRegistrationOutputWithContext(context.Background())
}

func (i *SqlServerRegistration) ToSqlServerRegistrationOutputWithContext(ctx context.Context) SqlServerRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerRegistrationOutput)
}

func (i *SqlServerRegistration) ToOutput(ctx context.Context) pulumix.Output[*SqlServerRegistration] {
	return pulumix.Output[*SqlServerRegistration]{
		OutputState: i.ToSqlServerRegistrationOutputWithContext(ctx).OutputState,
	}
}

type SqlServerRegistrationOutput struct{ *pulumi.OutputState }

func (SqlServerRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerRegistration)(nil)).Elem()
}

func (o SqlServerRegistrationOutput) ToSqlServerRegistrationOutput() SqlServerRegistrationOutput {
	return o
}

func (o SqlServerRegistrationOutput) ToSqlServerRegistrationOutputWithContext(ctx context.Context) SqlServerRegistrationOutput {
	return o
}

func (o SqlServerRegistrationOutput) ToOutput(ctx context.Context) pulumix.Output[*SqlServerRegistration] {
	return pulumix.Output[*SqlServerRegistration]{
		OutputState: o.OutputState,
	}
}

// The geo-location where the resource lives
func (o SqlServerRegistrationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServerRegistration) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SqlServerRegistrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServerRegistration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional Properties as JSON string
func (o SqlServerRegistrationOutput) PropertyBag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerRegistration) pulumi.StringPtrOutput { return v.PropertyBag }).(pulumi.StringPtrOutput)
}

// Resource Group Name
func (o SqlServerRegistrationOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerRegistration) pulumi.StringPtrOutput { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

// Subscription Id
func (o SqlServerRegistrationOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerRegistration) pulumi.StringPtrOutput { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

// Read only system data
func (o SqlServerRegistrationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlServerRegistration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SqlServerRegistrationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlServerRegistration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o SqlServerRegistrationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServerRegistration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlServerRegistrationOutput{})
}
