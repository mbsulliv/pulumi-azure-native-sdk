// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

// Show costs accumulated over time.
type AccumulatedType string

const (
	AccumulatedTypeTrue  = AccumulatedType("true")
	AccumulatedTypeFalse = AccumulatedType("false")
)

// Chart type of the main view in Cost Analysis. Required.
type ChartType string

const (
	ChartTypeArea          = ChartType("Area")
	ChartTypeLine          = ChartType("Line")
	ChartTypeStackedColumn = ChartType("StackedColumn")
	ChartTypeGroupedColumn = ChartType("GroupedColumn")
	ChartTypeTable         = ChartType("Table")
)

// Days of Week.
type DaysOfWeek string

const (
	DaysOfWeekMonday    = DaysOfWeek("Monday")
	DaysOfWeekTuesday   = DaysOfWeek("Tuesday")
	DaysOfWeekWednesday = DaysOfWeek("Wednesday")
	DaysOfWeekThursday  = DaysOfWeek("Thursday")
	DaysOfWeekFriday    = DaysOfWeek("Friday")
	DaysOfWeekSaturday  = DaysOfWeek("Saturday")
	DaysOfWeekSunday    = DaysOfWeek("Sunday")
)

// The type of the export. Note that 'Usage' is equivalent to 'ActualCost' and is applicable to exports that do not yet provide data for charges or amortization for service reservations.
type ExportType string

const (
	ExportTypeUsage         = ExportType("Usage")
	ExportTypeActualCost    = ExportType("ActualCost")
	ExportTypeAmortizedCost = ExportType("AmortizedCost")
)

// Destination of the view data. Currently only CSV format is supported.
type FileFormat string

const (
	FileFormatCsv = FileFormat("Csv")
)

// The format of the export being delivered. Currently only 'Csv' is supported.
type FormatType string

const (
	FormatTypeCsv = FormatType("Csv")
)

// The name of the aggregation function to use.
type FunctionType string

const (
	FunctionTypeSum = FunctionType("Sum")
)

// The granularity of rows in the export. Currently only 'Daily' is supported.
type GranularityType string

const (
	GranularityTypeDaily = GranularityType("Daily")
)

// KPI type (Forecast, Budget).
type KpiTypeType string

const (
	KpiTypeTypeForecast = KpiTypeType("Forecast")
	KpiTypeTypeBudget   = KpiTypeType("Budget")
)

// Metric to use when displaying costs.
type MetricType string

const (
	MetricTypeActualCost    = MetricType("ActualCost")
	MetricTypeAmortizedCost = MetricType("AmortizedCost")
	MetricTypeAHUB          = MetricType("AHUB")
)

// The operator to use for comparison.
type OperatorType string

const (
	OperatorTypeIn       = OperatorType("In")
	OperatorTypeContains = OperatorType("Contains")
)

// Data type to show in view.
type PivotTypeType string

const (
	PivotTypeTypeDimension = PivotTypeType("Dimension")
	PivotTypeTypeTagKey    = PivotTypeType("TagKey")
)

// Has type of the column to group.
type QueryColumnType string

const (
	// The tag associated with the cost data.
	QueryColumnTypeTagKey = QueryColumnType("TagKey")
	// The dimension of cost data.
	QueryColumnTypeDimension = QueryColumnType("Dimension")
)

// The schedule recurrence.
type RecurrenceType string

const (
	RecurrenceTypeDaily    = RecurrenceType("Daily")
	RecurrenceTypeWeekly   = RecurrenceType("Weekly")
	RecurrenceTypeMonthly  = RecurrenceType("Monthly")
	RecurrenceTypeAnnually = RecurrenceType("Annually")
)

// Direction of sort.
type ReportConfigSortingType string

const (
	ReportConfigSortingTypeAscending  = ReportConfigSortingType("Ascending")
	ReportConfigSortingTypeDescending = ReportConfigSortingType("Descending")
)

// The granularity of rows in the report.
type ReportGranularityType string

const (
	ReportGranularityTypeDaily   = ReportGranularityType("Daily")
	ReportGranularityTypeMonthly = ReportGranularityType("Monthly")
)

// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
type ReportTimeframeType string

const (
	ReportTimeframeTypeWeekToDate  = ReportTimeframeType("WeekToDate")
	ReportTimeframeTypeMonthToDate = ReportTimeframeType("MonthToDate")
	ReportTimeframeTypeYearToDate  = ReportTimeframeType("YearToDate")
	ReportTimeframeTypeCustom      = ReportTimeframeType("Custom")
)

// The type of the report. Usage represents actual usage, forecast represents forecasted data and UsageAndForecast represents both usage and forecasted data. Actual usage and forecasted data can be differentiated based on dates.
type ReportType string

const (
	ReportTypeUsage = ReportType("Usage")
)

// Frequency of the schedule.
type ScheduleFrequency string

const (
	// Cost analysis data will be emailed every day.
	ScheduleFrequencyDaily = ScheduleFrequency("Daily")
	// Cost analysis data will be emailed every week.
	ScheduleFrequencyWeekly = ScheduleFrequency("Weekly")
	// Cost analysis data will be emailed every month.
	ScheduleFrequencyMonthly = ScheduleFrequency("Monthly")
)

// Kind of the scheduled action.
type ScheduledActionKind string

const (
	// Cost analysis data will be emailed.
	ScheduledActionKindEmail = ScheduledActionKind("Email")
	// Cost anomaly information will be emailed. Available only on subscription scope at daily frequency. If no anomaly is detected on the resource, an email won't be sent.
	ScheduledActionKindInsightAlert = ScheduledActionKind("InsightAlert")
)

// Status of the scheduled action.
type ScheduledActionStatus string

const (
	// Scheduled action is saved but will not be run.
	ScheduledActionStatusDisabled = ScheduledActionStatus("Disabled")
	// Scheduled action is saved and will be run.
	ScheduledActionStatusEnabled = ScheduledActionStatus("Enabled")
	// Scheduled action is expired.
	ScheduledActionStatusExpired = ScheduledActionStatus("Expired")
)

// The status of the export's schedule. If 'Inactive', the export's schedule is paused.
type StatusType string

const (
	StatusTypeActive   = StatusType("Active")
	StatusTypeInactive = StatusType("Inactive")
)

// The time frame for pulling data for the export. If custom, then a specific time period must be provided.
type TimeframeType string

const (
	TimeframeTypeMonthToDate         = TimeframeType("MonthToDate")
	TimeframeTypeBillingMonthToDate  = TimeframeType("BillingMonthToDate")
	TimeframeTypeTheLastMonth        = TimeframeType("TheLastMonth")
	TimeframeTypeTheLastBillingMonth = TimeframeType("TheLastBillingMonth")
	TimeframeTypeWeekToDate          = TimeframeType("WeekToDate")
	TimeframeTypeCustom              = TimeframeType("Custom")
)

// Weeks of month.
type WeeksOfMonth string

const (
	WeeksOfMonthFirst  = WeeksOfMonth("First")
	WeeksOfMonthSecond = WeeksOfMonth("Second")
	WeeksOfMonthThird  = WeeksOfMonth("Third")
	WeeksOfMonthFourth = WeeksOfMonth("Fourth")
	WeeksOfMonthLast   = WeeksOfMonth("Last")
)

func init() {
}