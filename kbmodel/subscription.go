// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Subscription subscription
// swagger:model Subscription
type Subscription struct {

	// account Id
	// Format: uuid
	AccountID strfmt.UUID `json:"accountId,omitempty"`

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// bill cycle day local
	BillCycleDayLocal int32 `json:"billCycleDayLocal,omitempty"`

	// billing end date
	// Format: date-time
	BillingEndDate strfmt.DateTime `json:"billingEndDate,omitempty"`

	// billing period
	// Required: true
	// Enum: [DAILY WEEKLY BIWEEKLY THIRTY_DAYS THIRTY_ONE_DAYS SIXTY_DAYS NINETY_DAYS MONTHLY BIMESTRIAL QUARTERLY TRIANNUAL BIANNUAL ANNUAL SESQUIENNIAL BIENNIAL TRIENNIAL NO_BILLING_PERIOD]
	BillingPeriod *SubscriptionBillingPeriodEnum `json:"billingPeriod"`

	// billing start date
	// Format: date-time
	BillingStartDate strfmt.DateTime `json:"billingStartDate,omitempty"`

	// bundle external key
	BundleExternalKey string `json:"bundleExternalKey,omitempty"`

	// bundle Id
	// Format: uuid
	BundleID strfmt.UUID `json:"bundleId,omitempty"`

	// cancelled date
	// Format: date-time
	CancelledDate strfmt.DateTime `json:"cancelledDate,omitempty"`

	// charged through date
	// Format: date
	ChargedThroughDate strfmt.Date `json:"chargedThroughDate,omitempty"`

	// events
	Events []*EventSubscription `json:"events"`

	// external key
	ExternalKey string `json:"externalKey,omitempty"`

	// phase type
	// Enum: [TRIAL DISCOUNT FIXEDTERM EVERGREEN]
	PhaseType SubscriptionPhaseTypeEnum `json:"phaseType,omitempty"`

	// plan name
	// Required: true
	PlanName *string `json:"planName"`

	// price list
	// Required: true
	PriceList *string `json:"priceList"`

	// price overrides
	PriceOverrides []*PhasePrice `json:"priceOverrides"`

	// prices
	Prices []*PhasePrice `json:"prices"`

	// product category
	// Enum: [BASE ADD_ON STANDALONE]
	ProductCategory SubscriptionProductCategoryEnum `json:"productCategory,omitempty"`

	// product name
	// Required: true
	ProductName *string `json:"productName"`

	// quantity
	Quantity int32 `json:"quantity,omitempty"`

	// source type
	// Enum: [NATIVE MIGRATED TRANSFERRED]
	SourceType SubscriptionSourceTypeEnum `json:"sourceType,omitempty"`

	// start date
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// state
	// Enum: [PENDING ACTIVE BLOCKED CANCELLED EXPIRED]
	State SubscriptionStateEnum `json:"state,omitempty"`

	// subscription Id
	// Format: uuid
	SubscriptionID strfmt.UUID `json:"subscriptionId,omitempty"`
}

// Validate validates this subscription
func (m *Subscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBundleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCancelledDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargedThroughDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhaseType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceOverrides(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subscription) validateAccountID(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := validate.FormatOf("accountId", "body", "uuid", m.AccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateAuditLogs(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditLogs); i++ {
		if swag.IsZero(m.AuditLogs[i]) { // not required
			continue
		}

		if m.AuditLogs[i] != nil {
			if err := m.AuditLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Subscription) validateBillingEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("billingEndDate", "body", "date-time", m.BillingEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var subscriptionTypeBillingPeriodPropEnum []interface{}

func init() {
	var res []SubscriptionBillingPeriodEnum
	if err := json.Unmarshal([]byte(`["DAILY","WEEKLY","BIWEEKLY","THIRTY_DAYS","THIRTY_ONE_DAYS","SIXTY_DAYS","NINETY_DAYS","MONTHLY","BIMESTRIAL","QUARTERLY","TRIANNUAL","BIANNUAL","ANNUAL","SESQUIENNIAL","BIENNIAL","TRIENNIAL","NO_BILLING_PERIOD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionTypeBillingPeriodPropEnum = append(subscriptionTypeBillingPeriodPropEnum, v)
	}
}

type SubscriptionBillingPeriodEnum string

const (

	// SubscriptionBillingPeriodDAILY captures enum value "DAILY"
	SubscriptionBillingPeriodDAILY SubscriptionBillingPeriodEnum = "DAILY"

	// SubscriptionBillingPeriodWEEKLY captures enum value "WEEKLY"
	SubscriptionBillingPeriodWEEKLY SubscriptionBillingPeriodEnum = "WEEKLY"

	// SubscriptionBillingPeriodBIWEEKLY captures enum value "BIWEEKLY"
	SubscriptionBillingPeriodBIWEEKLY SubscriptionBillingPeriodEnum = "BIWEEKLY"

	// SubscriptionBillingPeriodTHIRTYDAYS captures enum value "THIRTY_DAYS"
	SubscriptionBillingPeriodTHIRTYDAYS SubscriptionBillingPeriodEnum = "THIRTY_DAYS"

	// SubscriptionBillingPeriodTHIRTYONEDAYS captures enum value "THIRTY_ONE_DAYS"
	SubscriptionBillingPeriodTHIRTYONEDAYS SubscriptionBillingPeriodEnum = "THIRTY_ONE_DAYS"

	// SubscriptionBillingPeriodSIXTYDAYS captures enum value "SIXTY_DAYS"
	SubscriptionBillingPeriodSIXTYDAYS SubscriptionBillingPeriodEnum = "SIXTY_DAYS"

	// SubscriptionBillingPeriodNINETYDAYS captures enum value "NINETY_DAYS"
	SubscriptionBillingPeriodNINETYDAYS SubscriptionBillingPeriodEnum = "NINETY_DAYS"

	// SubscriptionBillingPeriodMONTHLY captures enum value "MONTHLY"
	SubscriptionBillingPeriodMONTHLY SubscriptionBillingPeriodEnum = "MONTHLY"

	// SubscriptionBillingPeriodBIMESTRIAL captures enum value "BIMESTRIAL"
	SubscriptionBillingPeriodBIMESTRIAL SubscriptionBillingPeriodEnum = "BIMESTRIAL"

	// SubscriptionBillingPeriodQUARTERLY captures enum value "QUARTERLY"
	SubscriptionBillingPeriodQUARTERLY SubscriptionBillingPeriodEnum = "QUARTERLY"

	// SubscriptionBillingPeriodTRIANNUAL captures enum value "TRIANNUAL"
	SubscriptionBillingPeriodTRIANNUAL SubscriptionBillingPeriodEnum = "TRIANNUAL"

	// SubscriptionBillingPeriodBIANNUAL captures enum value "BIANNUAL"
	SubscriptionBillingPeriodBIANNUAL SubscriptionBillingPeriodEnum = "BIANNUAL"

	// SubscriptionBillingPeriodANNUAL captures enum value "ANNUAL"
	SubscriptionBillingPeriodANNUAL SubscriptionBillingPeriodEnum = "ANNUAL"

	// SubscriptionBillingPeriodSESQUIENNIAL captures enum value "SESQUIENNIAL"
	SubscriptionBillingPeriodSESQUIENNIAL SubscriptionBillingPeriodEnum = "SESQUIENNIAL"

	// SubscriptionBillingPeriodBIENNIAL captures enum value "BIENNIAL"
	SubscriptionBillingPeriodBIENNIAL SubscriptionBillingPeriodEnum = "BIENNIAL"

	// SubscriptionBillingPeriodTRIENNIAL captures enum value "TRIENNIAL"
	SubscriptionBillingPeriodTRIENNIAL SubscriptionBillingPeriodEnum = "TRIENNIAL"

	// SubscriptionBillingPeriodNOBILLINGPERIOD captures enum value "NO_BILLING_PERIOD"
	SubscriptionBillingPeriodNOBILLINGPERIOD SubscriptionBillingPeriodEnum = "NO_BILLING_PERIOD"
)

var SubscriptionBillingPeriodEnumValues = []string{
	"DAILY",
	"WEEKLY",
	"BIWEEKLY",
	"THIRTY_DAYS",
	"THIRTY_ONE_DAYS",
	"SIXTY_DAYS",
	"NINETY_DAYS",
	"MONTHLY",
	"BIMESTRIAL",
	"QUARTERLY",
	"TRIANNUAL",
	"BIANNUAL",
	"ANNUAL",
	"SESQUIENNIAL",
	"BIENNIAL",
	"TRIENNIAL",
	"NO_BILLING_PERIOD",
}

func (e SubscriptionBillingPeriodEnum) IsValid() bool {
	for _, v := range SubscriptionBillingPeriodEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *Subscription) validateBillingPeriodEnum(path, location string, value SubscriptionBillingPeriodEnum) error {
	if err := validate.Enum(path, location, value, subscriptionTypeBillingPeriodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subscription) validateBillingPeriod(formats strfmt.Registry) error {

	if err := validate.Required("billingPeriod", "body", m.BillingPeriod); err != nil {
		return err
	}

	// value enum
	if err := m.validateBillingPeriodEnum("billingPeriod", "body", *m.BillingPeriod); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateBillingStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("billingStartDate", "body", "date-time", m.BillingStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateBundleID(formats strfmt.Registry) error {

	if swag.IsZero(m.BundleID) { // not required
		return nil
	}

	if err := validate.FormatOf("bundleId", "body", "uuid", m.BundleID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateCancelledDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CancelledDate) { // not required
		return nil
	}

	if err := validate.FormatOf("cancelledDate", "body", "date-time", m.CancelledDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateChargedThroughDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargedThroughDate) { // not required
		return nil
	}

	if err := validate.FormatOf("chargedThroughDate", "body", "date", m.ChargedThroughDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {
		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {
			if err := m.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var subscriptionTypePhaseTypePropEnum []interface{}

func init() {
	var res []SubscriptionPhaseTypeEnum
	if err := json.Unmarshal([]byte(`["TRIAL","DISCOUNT","FIXEDTERM","EVERGREEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionTypePhaseTypePropEnum = append(subscriptionTypePhaseTypePropEnum, v)
	}
}

type SubscriptionPhaseTypeEnum string

const (

	// SubscriptionPhaseTypeTRIAL captures enum value "TRIAL"
	SubscriptionPhaseTypeTRIAL SubscriptionPhaseTypeEnum = "TRIAL"

	// SubscriptionPhaseTypeDISCOUNT captures enum value "DISCOUNT"
	SubscriptionPhaseTypeDISCOUNT SubscriptionPhaseTypeEnum = "DISCOUNT"

	// SubscriptionPhaseTypeFIXEDTERM captures enum value "FIXEDTERM"
	SubscriptionPhaseTypeFIXEDTERM SubscriptionPhaseTypeEnum = "FIXEDTERM"

	// SubscriptionPhaseTypeEVERGREEN captures enum value "EVERGREEN"
	SubscriptionPhaseTypeEVERGREEN SubscriptionPhaseTypeEnum = "EVERGREEN"
)

var SubscriptionPhaseTypeEnumValues = []string{
	"TRIAL",
	"DISCOUNT",
	"FIXEDTERM",
	"EVERGREEN",
}

func (e SubscriptionPhaseTypeEnum) IsValid() bool {
	for _, v := range SubscriptionPhaseTypeEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *Subscription) validatePhaseTypeEnum(path, location string, value SubscriptionPhaseTypeEnum) error {
	if err := validate.Enum(path, location, value, subscriptionTypePhaseTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subscription) validatePhaseType(formats strfmt.Registry) error {

	if swag.IsZero(m.PhaseType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePhaseTypeEnum("phaseType", "body", m.PhaseType); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validatePlanName(formats strfmt.Registry) error {

	if err := validate.Required("planName", "body", m.PlanName); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validatePriceList(formats strfmt.Registry) error {

	if err := validate.Required("priceList", "body", m.PriceList); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validatePriceOverrides(formats strfmt.Registry) error {

	if swag.IsZero(m.PriceOverrides) { // not required
		return nil
	}

	for i := 0; i < len(m.PriceOverrides); i++ {
		if swag.IsZero(m.PriceOverrides[i]) { // not required
			continue
		}

		if m.PriceOverrides[i] != nil {
			if err := m.PriceOverrides[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("priceOverrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Subscription) validatePrices(formats strfmt.Registry) error {

	if swag.IsZero(m.Prices) { // not required
		return nil
	}

	for i := 0; i < len(m.Prices); i++ {
		if swag.IsZero(m.Prices[i]) { // not required
			continue
		}

		if m.Prices[i] != nil {
			if err := m.Prices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var subscriptionTypeProductCategoryPropEnum []interface{}

func init() {
	var res []SubscriptionProductCategoryEnum
	if err := json.Unmarshal([]byte(`["BASE","ADD_ON","STANDALONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionTypeProductCategoryPropEnum = append(subscriptionTypeProductCategoryPropEnum, v)
	}
}

type SubscriptionProductCategoryEnum string

const (

	// SubscriptionProductCategoryBASE captures enum value "BASE"
	SubscriptionProductCategoryBASE SubscriptionProductCategoryEnum = "BASE"

	// SubscriptionProductCategoryADDON captures enum value "ADD_ON"
	SubscriptionProductCategoryADDON SubscriptionProductCategoryEnum = "ADD_ON"

	// SubscriptionProductCategorySTANDALONE captures enum value "STANDALONE"
	SubscriptionProductCategorySTANDALONE SubscriptionProductCategoryEnum = "STANDALONE"
)

var SubscriptionProductCategoryEnumValues = []string{
	"BASE",
	"ADD_ON",
	"STANDALONE",
}

func (e SubscriptionProductCategoryEnum) IsValid() bool {
	for _, v := range SubscriptionProductCategoryEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *Subscription) validateProductCategoryEnum(path, location string, value SubscriptionProductCategoryEnum) error {
	if err := validate.Enum(path, location, value, subscriptionTypeProductCategoryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subscription) validateProductCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductCategory) { // not required
		return nil
	}

	// value enum
	if err := m.validateProductCategoryEnum("productCategory", "body", m.ProductCategory); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateProductName(formats strfmt.Registry) error {

	if err := validate.Required("productName", "body", m.ProductName); err != nil {
		return err
	}

	return nil
}

var subscriptionTypeSourceTypePropEnum []interface{}

func init() {
	var res []SubscriptionSourceTypeEnum
	if err := json.Unmarshal([]byte(`["NATIVE","MIGRATED","TRANSFERRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionTypeSourceTypePropEnum = append(subscriptionTypeSourceTypePropEnum, v)
	}
}

type SubscriptionSourceTypeEnum string

const (

	// SubscriptionSourceTypeNATIVE captures enum value "NATIVE"
	SubscriptionSourceTypeNATIVE SubscriptionSourceTypeEnum = "NATIVE"

	// SubscriptionSourceTypeMIGRATED captures enum value "MIGRATED"
	SubscriptionSourceTypeMIGRATED SubscriptionSourceTypeEnum = "MIGRATED"

	// SubscriptionSourceTypeTRANSFERRED captures enum value "TRANSFERRED"
	SubscriptionSourceTypeTRANSFERRED SubscriptionSourceTypeEnum = "TRANSFERRED"
)

var SubscriptionSourceTypeEnumValues = []string{
	"NATIVE",
	"MIGRATED",
	"TRANSFERRED",
}

func (e SubscriptionSourceTypeEnum) IsValid() bool {
	for _, v := range SubscriptionSourceTypeEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *Subscription) validateSourceTypeEnum(path, location string, value SubscriptionSourceTypeEnum) error {
	if err := validate.Enum(path, location, value, subscriptionTypeSourceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subscription) validateSourceType(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSourceTypeEnum("sourceType", "body", m.SourceType); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var subscriptionTypeStatePropEnum []interface{}

func init() {
	var res []SubscriptionStateEnum
	if err := json.Unmarshal([]byte(`["PENDING","ACTIVE","BLOCKED","CANCELLED","EXPIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionTypeStatePropEnum = append(subscriptionTypeStatePropEnum, v)
	}
}

type SubscriptionStateEnum string

const (

	// SubscriptionStatePENDING captures enum value "PENDING"
	SubscriptionStatePENDING SubscriptionStateEnum = "PENDING"

	// SubscriptionStateACTIVE captures enum value "ACTIVE"
	SubscriptionStateACTIVE SubscriptionStateEnum = "ACTIVE"

	// SubscriptionStateBLOCKED captures enum value "BLOCKED"
	SubscriptionStateBLOCKED SubscriptionStateEnum = "BLOCKED"

	// SubscriptionStateCANCELLED captures enum value "CANCELLED"
	SubscriptionStateCANCELLED SubscriptionStateEnum = "CANCELLED"

	// SubscriptionStateEXPIRED captures enum value "EXPIRED"
	SubscriptionStateEXPIRED SubscriptionStateEnum = "EXPIRED"
)

var SubscriptionStateEnumValues = []string{
	"PENDING",
	"ACTIVE",
	"BLOCKED",
	"CANCELLED",
	"EXPIRED",
}

func (e SubscriptionStateEnum) IsValid() bool {
	for _, v := range SubscriptionStateEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *Subscription) validateStateEnum(path, location string, value SubscriptionStateEnum) error {
	if err := validate.Enum(path, location, value, subscriptionTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subscription) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateSubscriptionID(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionID) { // not required
		return nil
	}

	if err := validate.FormatOf("subscriptionId", "body", "uuid", m.SubscriptionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subscription) UnmarshalBinary(b []byte) error {
	var res Subscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
