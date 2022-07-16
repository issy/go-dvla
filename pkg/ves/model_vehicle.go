/*
Vehicle Enquiry API

Interface specification for the DVLA Vehicle Enquiry API

API version: 1.2.0
Contact: DvlaAPIAccess@dvla.gov.uk
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ves

import (
	"encoding/json"
)

// Vehicle struct for Vehicle
type Vehicle struct {
	// Registration number of the vehicle
	RegistrationNumber string `json:"registrationNumber"`
	// Tax status of the vehicle
	TaxStatus *string `json:"taxStatus,omitempty"`
	// Date of tax liablity, Used in calculating licence information presented to user
	TaxDueDate *string `json:"taxDueDate,omitempty"`
	// Additional Rate of Tax End Date, format: YYYY-MM-DD
	ArtEndDate *string `json:"artEndDate,omitempty"`
	// MOT Status of the vehicle
	MotStatus *string `json:"motStatus,omitempty"`
	// Mot Expiry Date
	MotExpiryDate *string `json:"motExpiryDate,omitempty"`
	// Vehicle make
	Make *string `json:"make,omitempty"`
	// Month of First DVLA Registration
	MonthOfFirstDvlaRegistration *string `json:"monthOfFirstDvlaRegistration,omitempty"`
	// Month of First Registration
	MonthOfFirstRegistration *string `json:"monthOfFirstRegistration,omitempty"`
	// Year of Manufacture
	YearOfManufacture *int32 `json:"yearOfManufacture,omitempty"`
	// Engine capacity in cubic centimetres
	EngineCapacity *int32 `json:"engineCapacity,omitempty"`
	// Carbon Dioxide emissions in grams per kilometre
	Co2Emissions *int32 `json:"co2Emissions,omitempty"`
	// Fuel type (Method of Propulsion)
	FuelType *string `json:"fuelType,omitempty"`
	// True only if vehicle has been export marked
	MarkedForExport *bool `json:"markedForExport,omitempty"`
	// Vehicle colour
	Colour *string `json:"colour,omitempty"`
	// Vehicle Type Approval Category
	TypeApproval *string `json:"typeApproval,omitempty"`
	// Vehicle wheel plan
	Wheelplan *string `json:"wheelplan,omitempty"`
	// Revenue weight in kilograms
	RevenueWeight *int32 `json:"revenueWeight,omitempty"`
	// Real Driving Emissions value
	RealDrivingEmissions *string `json:"realDrivingEmissions,omitempty"`
	// Date of last V5C issued
	DateOfLastV5CIssued *string `json:"dateOfLastV5CIssued,omitempty"`
	// Euro Status (Dealer / Customer Provided (new vehicles))
	EuroStatus *string `json:"euroStatus,omitempty"`
	// Automated Vehicle (AV)
	AutomatedVehicle *bool `json:"automatedVehicle,omitempty"`
}

// NewVehicle instantiates a new Vehicle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicle(registrationNumber string) *Vehicle {
	this := Vehicle{}
	this.RegistrationNumber = registrationNumber
	return &this
}

// NewVehicleWithDefaults instantiates a new Vehicle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleWithDefaults() *Vehicle {
	this := Vehicle{}
	return &this
}

// GetRegistrationNumber returns the RegistrationNumber field value
func (o *Vehicle) GetRegistrationNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistrationNumber
}

// GetRegistrationNumberOk returns a tuple with the RegistrationNumber field value
// and a boolean to check if the value has been set.
func (o *Vehicle) GetRegistrationNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistrationNumber, true
}

// SetRegistrationNumber sets field value
func (o *Vehicle) SetRegistrationNumber(v string) {
	o.RegistrationNumber = v
}

// GetTaxStatus returns the TaxStatus field value if set, zero value otherwise.
func (o *Vehicle) GetTaxStatus() string {
	if o == nil || o.TaxStatus == nil {
		var ret string
		return ret
	}
	return *o.TaxStatus
}

// GetTaxStatusOk returns a tuple with the TaxStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetTaxStatusOk() (*string, bool) {
	if o == nil || o.TaxStatus == nil {
		return nil, false
	}
	return o.TaxStatus, true
}

// HasTaxStatus returns a boolean if a field has been set.
func (o *Vehicle) HasTaxStatus() bool {
	if o != nil && o.TaxStatus != nil {
		return true
	}

	return false
}

// SetTaxStatus gets a reference to the given string and assigns it to the TaxStatus field.
func (o *Vehicle) SetTaxStatus(v string) {
	o.TaxStatus = &v
}

// GetTaxDueDate returns the TaxDueDate field value if set, zero value otherwise.
func (o *Vehicle) GetTaxDueDate() string {
	if o == nil || o.TaxDueDate == nil {
		var ret string
		return ret
	}
	return *o.TaxDueDate
}

// GetTaxDueDateOk returns a tuple with the TaxDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetTaxDueDateOk() (*string, bool) {
	if o == nil || o.TaxDueDate == nil {
		return nil, false
	}
	return o.TaxDueDate, true
}

// HasTaxDueDate returns a boolean if a field has been set.
func (o *Vehicle) HasTaxDueDate() bool {
	if o != nil && o.TaxDueDate != nil {
		return true
	}

	return false
}

// SetTaxDueDate gets a reference to the given string and assigns it to the TaxDueDate field.
func (o *Vehicle) SetTaxDueDate(v string) {
	o.TaxDueDate = &v
}

// GetArtEndDate returns the ArtEndDate field value if set, zero value otherwise.
func (o *Vehicle) GetArtEndDate() string {
	if o == nil || o.ArtEndDate == nil {
		var ret string
		return ret
	}
	return *o.ArtEndDate
}

// GetArtEndDateOk returns a tuple with the ArtEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetArtEndDateOk() (*string, bool) {
	if o == nil || o.ArtEndDate == nil {
		return nil, false
	}
	return o.ArtEndDate, true
}

// HasArtEndDate returns a boolean if a field has been set.
func (o *Vehicle) HasArtEndDate() bool {
	if o != nil && o.ArtEndDate != nil {
		return true
	}

	return false
}

// SetArtEndDate gets a reference to the given string and assigns it to the ArtEndDate field.
func (o *Vehicle) SetArtEndDate(v string) {
	o.ArtEndDate = &v
}

// GetMotStatus returns the MotStatus field value if set, zero value otherwise.
func (o *Vehicle) GetMotStatus() string {
	if o == nil || o.MotStatus == nil {
		var ret string
		return ret
	}
	return *o.MotStatus
}

// GetMotStatusOk returns a tuple with the MotStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetMotStatusOk() (*string, bool) {
	if o == nil || o.MotStatus == nil {
		return nil, false
	}
	return o.MotStatus, true
}

// HasMotStatus returns a boolean if a field has been set.
func (o *Vehicle) HasMotStatus() bool {
	if o != nil && o.MotStatus != nil {
		return true
	}

	return false
}

// SetMotStatus gets a reference to the given string and assigns it to the MotStatus field.
func (o *Vehicle) SetMotStatus(v string) {
	o.MotStatus = &v
}

// GetMotExpiryDate returns the MotExpiryDate field value if set, zero value otherwise.
func (o *Vehicle) GetMotExpiryDate() string {
	if o == nil || o.MotExpiryDate == nil {
		var ret string
		return ret
	}
	return *o.MotExpiryDate
}

// GetMotExpiryDateOk returns a tuple with the MotExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetMotExpiryDateOk() (*string, bool) {
	if o == nil || o.MotExpiryDate == nil {
		return nil, false
	}
	return o.MotExpiryDate, true
}

// HasMotExpiryDate returns a boolean if a field has been set.
func (o *Vehicle) HasMotExpiryDate() bool {
	if o != nil && o.MotExpiryDate != nil {
		return true
	}

	return false
}

// SetMotExpiryDate gets a reference to the given string and assigns it to the MotExpiryDate field.
func (o *Vehicle) SetMotExpiryDate(v string) {
	o.MotExpiryDate = &v
}

// GetMake returns the Make field value if set, zero value otherwise.
func (o *Vehicle) GetMake() string {
	if o == nil || o.Make == nil {
		var ret string
		return ret
	}
	return *o.Make
}

// GetMakeOk returns a tuple with the Make field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetMakeOk() (*string, bool) {
	if o == nil || o.Make == nil {
		return nil, false
	}
	return o.Make, true
}

// HasMake returns a boolean if a field has been set.
func (o *Vehicle) HasMake() bool {
	if o != nil && o.Make != nil {
		return true
	}

	return false
}

// SetMake gets a reference to the given string and assigns it to the Make field.
func (o *Vehicle) SetMake(v string) {
	o.Make = &v
}

// GetMonthOfFirstDvlaRegistration returns the MonthOfFirstDvlaRegistration field value if set, zero value otherwise.
func (o *Vehicle) GetMonthOfFirstDvlaRegistration() string {
	if o == nil || o.MonthOfFirstDvlaRegistration == nil {
		var ret string
		return ret
	}
	return *o.MonthOfFirstDvlaRegistration
}

// GetMonthOfFirstDvlaRegistrationOk returns a tuple with the MonthOfFirstDvlaRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetMonthOfFirstDvlaRegistrationOk() (*string, bool) {
	if o == nil || o.MonthOfFirstDvlaRegistration == nil {
		return nil, false
	}
	return o.MonthOfFirstDvlaRegistration, true
}

// HasMonthOfFirstDvlaRegistration returns a boolean if a field has been set.
func (o *Vehicle) HasMonthOfFirstDvlaRegistration() bool {
	if o != nil && o.MonthOfFirstDvlaRegistration != nil {
		return true
	}

	return false
}

// SetMonthOfFirstDvlaRegistration gets a reference to the given string and assigns it to the MonthOfFirstDvlaRegistration field.
func (o *Vehicle) SetMonthOfFirstDvlaRegistration(v string) {
	o.MonthOfFirstDvlaRegistration = &v
}

// GetMonthOfFirstRegistration returns the MonthOfFirstRegistration field value if set, zero value otherwise.
func (o *Vehicle) GetMonthOfFirstRegistration() string {
	if o == nil || o.MonthOfFirstRegistration == nil {
		var ret string
		return ret
	}
	return *o.MonthOfFirstRegistration
}

// GetMonthOfFirstRegistrationOk returns a tuple with the MonthOfFirstRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetMonthOfFirstRegistrationOk() (*string, bool) {
	if o == nil || o.MonthOfFirstRegistration == nil {
		return nil, false
	}
	return o.MonthOfFirstRegistration, true
}

// HasMonthOfFirstRegistration returns a boolean if a field has been set.
func (o *Vehicle) HasMonthOfFirstRegistration() bool {
	if o != nil && o.MonthOfFirstRegistration != nil {
		return true
	}

	return false
}

// SetMonthOfFirstRegistration gets a reference to the given string and assigns it to the MonthOfFirstRegistration field.
func (o *Vehicle) SetMonthOfFirstRegistration(v string) {
	o.MonthOfFirstRegistration = &v
}

// GetYearOfManufacture returns the YearOfManufacture field value if set, zero value otherwise.
func (o *Vehicle) GetYearOfManufacture() int32 {
	if o == nil || o.YearOfManufacture == nil {
		var ret int32
		return ret
	}
	return *o.YearOfManufacture
}

// GetYearOfManufactureOk returns a tuple with the YearOfManufacture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetYearOfManufactureOk() (*int32, bool) {
	if o == nil || o.YearOfManufacture == nil {
		return nil, false
	}
	return o.YearOfManufacture, true
}

// HasYearOfManufacture returns a boolean if a field has been set.
func (o *Vehicle) HasYearOfManufacture() bool {
	if o != nil && o.YearOfManufacture != nil {
		return true
	}

	return false
}

// SetYearOfManufacture gets a reference to the given int32 and assigns it to the YearOfManufacture field.
func (o *Vehicle) SetYearOfManufacture(v int32) {
	o.YearOfManufacture = &v
}

// GetEngineCapacity returns the EngineCapacity field value if set, zero value otherwise.
func (o *Vehicle) GetEngineCapacity() int32 {
	if o == nil || o.EngineCapacity == nil {
		var ret int32
		return ret
	}
	return *o.EngineCapacity
}

// GetEngineCapacityOk returns a tuple with the EngineCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetEngineCapacityOk() (*int32, bool) {
	if o == nil || o.EngineCapacity == nil {
		return nil, false
	}
	return o.EngineCapacity, true
}

// HasEngineCapacity returns a boolean if a field has been set.
func (o *Vehicle) HasEngineCapacity() bool {
	if o != nil && o.EngineCapacity != nil {
		return true
	}

	return false
}

// SetEngineCapacity gets a reference to the given int32 and assigns it to the EngineCapacity field.
func (o *Vehicle) SetEngineCapacity(v int32) {
	o.EngineCapacity = &v
}

// GetCo2Emissions returns the Co2Emissions field value if set, zero value otherwise.
func (o *Vehicle) GetCo2Emissions() int32 {
	if o == nil || o.Co2Emissions == nil {
		var ret int32
		return ret
	}
	return *o.Co2Emissions
}

// GetCo2EmissionsOk returns a tuple with the Co2Emissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetCo2EmissionsOk() (*int32, bool) {
	if o == nil || o.Co2Emissions == nil {
		return nil, false
	}
	return o.Co2Emissions, true
}

// HasCo2Emissions returns a boolean if a field has been set.
func (o *Vehicle) HasCo2Emissions() bool {
	if o != nil && o.Co2Emissions != nil {
		return true
	}

	return false
}

// SetCo2Emissions gets a reference to the given int32 and assigns it to the Co2Emissions field.
func (o *Vehicle) SetCo2Emissions(v int32) {
	o.Co2Emissions = &v
}

// GetFuelType returns the FuelType field value if set, zero value otherwise.
func (o *Vehicle) GetFuelType() string {
	if o == nil || o.FuelType == nil {
		var ret string
		return ret
	}
	return *o.FuelType
}

// GetFuelTypeOk returns a tuple with the FuelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetFuelTypeOk() (*string, bool) {
	if o == nil || o.FuelType == nil {
		return nil, false
	}
	return o.FuelType, true
}

// HasFuelType returns a boolean if a field has been set.
func (o *Vehicle) HasFuelType() bool {
	if o != nil && o.FuelType != nil {
		return true
	}

	return false
}

// SetFuelType gets a reference to the given string and assigns it to the FuelType field.
func (o *Vehicle) SetFuelType(v string) {
	o.FuelType = &v
}

// GetMarkedForExport returns the MarkedForExport field value if set, zero value otherwise.
func (o *Vehicle) GetMarkedForExport() bool {
	if o == nil || o.MarkedForExport == nil {
		var ret bool
		return ret
	}
	return *o.MarkedForExport
}

// GetMarkedForExportOk returns a tuple with the MarkedForExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetMarkedForExportOk() (*bool, bool) {
	if o == nil || o.MarkedForExport == nil {
		return nil, false
	}
	return o.MarkedForExport, true
}

// HasMarkedForExport returns a boolean if a field has been set.
func (o *Vehicle) HasMarkedForExport() bool {
	if o != nil && o.MarkedForExport != nil {
		return true
	}

	return false
}

// SetMarkedForExport gets a reference to the given bool and assigns it to the MarkedForExport field.
func (o *Vehicle) SetMarkedForExport(v bool) {
	o.MarkedForExport = &v
}

// GetColour returns the Colour field value if set, zero value otherwise.
func (o *Vehicle) GetColour() string {
	if o == nil || o.Colour == nil {
		var ret string
		return ret
	}
	return *o.Colour
}

// GetColourOk returns a tuple with the Colour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetColourOk() (*string, bool) {
	if o == nil || o.Colour == nil {
		return nil, false
	}
	return o.Colour, true
}

// HasColour returns a boolean if a field has been set.
func (o *Vehicle) HasColour() bool {
	if o != nil && o.Colour != nil {
		return true
	}

	return false
}

// SetColour gets a reference to the given string and assigns it to the Colour field.
func (o *Vehicle) SetColour(v string) {
	o.Colour = &v
}

// GetTypeApproval returns the TypeApproval field value if set, zero value otherwise.
func (o *Vehicle) GetTypeApproval() string {
	if o == nil || o.TypeApproval == nil {
		var ret string
		return ret
	}
	return *o.TypeApproval
}

// GetTypeApprovalOk returns a tuple with the TypeApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetTypeApprovalOk() (*string, bool) {
	if o == nil || o.TypeApproval == nil {
		return nil, false
	}
	return o.TypeApproval, true
}

// HasTypeApproval returns a boolean if a field has been set.
func (o *Vehicle) HasTypeApproval() bool {
	if o != nil && o.TypeApproval != nil {
		return true
	}

	return false
}

// SetTypeApproval gets a reference to the given string and assigns it to the TypeApproval field.
func (o *Vehicle) SetTypeApproval(v string) {
	o.TypeApproval = &v
}

// GetWheelplan returns the Wheelplan field value if set, zero value otherwise.
func (o *Vehicle) GetWheelplan() string {
	if o == nil || o.Wheelplan == nil {
		var ret string
		return ret
	}
	return *o.Wheelplan
}

// GetWheelplanOk returns a tuple with the Wheelplan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetWheelplanOk() (*string, bool) {
	if o == nil || o.Wheelplan == nil {
		return nil, false
	}
	return o.Wheelplan, true
}

// HasWheelplan returns a boolean if a field has been set.
func (o *Vehicle) HasWheelplan() bool {
	if o != nil && o.Wheelplan != nil {
		return true
	}

	return false
}

// SetWheelplan gets a reference to the given string and assigns it to the Wheelplan field.
func (o *Vehicle) SetWheelplan(v string) {
	o.Wheelplan = &v
}

// GetRevenueWeight returns the RevenueWeight field value if set, zero value otherwise.
func (o *Vehicle) GetRevenueWeight() int32 {
	if o == nil || o.RevenueWeight == nil {
		var ret int32
		return ret
	}
	return *o.RevenueWeight
}

// GetRevenueWeightOk returns a tuple with the RevenueWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetRevenueWeightOk() (*int32, bool) {
	if o == nil || o.RevenueWeight == nil {
		return nil, false
	}
	return o.RevenueWeight, true
}

// HasRevenueWeight returns a boolean if a field has been set.
func (o *Vehicle) HasRevenueWeight() bool {
	if o != nil && o.RevenueWeight != nil {
		return true
	}

	return false
}

// SetRevenueWeight gets a reference to the given int32 and assigns it to the RevenueWeight field.
func (o *Vehicle) SetRevenueWeight(v int32) {
	o.RevenueWeight = &v
}

// GetRealDrivingEmissions returns the RealDrivingEmissions field value if set, zero value otherwise.
func (o *Vehicle) GetRealDrivingEmissions() string {
	if o == nil || o.RealDrivingEmissions == nil {
		var ret string
		return ret
	}
	return *o.RealDrivingEmissions
}

// GetRealDrivingEmissionsOk returns a tuple with the RealDrivingEmissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetRealDrivingEmissionsOk() (*string, bool) {
	if o == nil || o.RealDrivingEmissions == nil {
		return nil, false
	}
	return o.RealDrivingEmissions, true
}

// HasRealDrivingEmissions returns a boolean if a field has been set.
func (o *Vehicle) HasRealDrivingEmissions() bool {
	if o != nil && o.RealDrivingEmissions != nil {
		return true
	}

	return false
}

// SetRealDrivingEmissions gets a reference to the given string and assigns it to the RealDrivingEmissions field.
func (o *Vehicle) SetRealDrivingEmissions(v string) {
	o.RealDrivingEmissions = &v
}

// GetDateOfLastV5CIssued returns the DateOfLastV5CIssued field value if set, zero value otherwise.
func (o *Vehicle) GetDateOfLastV5CIssued() string {
	if o == nil || o.DateOfLastV5CIssued == nil {
		var ret string
		return ret
	}
	return *o.DateOfLastV5CIssued
}

// GetDateOfLastV5CIssuedOk returns a tuple with the DateOfLastV5CIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetDateOfLastV5CIssuedOk() (*string, bool) {
	if o == nil || o.DateOfLastV5CIssued == nil {
		return nil, false
	}
	return o.DateOfLastV5CIssued, true
}

// HasDateOfLastV5CIssued returns a boolean if a field has been set.
func (o *Vehicle) HasDateOfLastV5CIssued() bool {
	if o != nil && o.DateOfLastV5CIssued != nil {
		return true
	}

	return false
}

// SetDateOfLastV5CIssued gets a reference to the given string and assigns it to the DateOfLastV5CIssued field.
func (o *Vehicle) SetDateOfLastV5CIssued(v string) {
	o.DateOfLastV5CIssued = &v
}

// GetEuroStatus returns the EuroStatus field value if set, zero value otherwise.
func (o *Vehicle) GetEuroStatus() string {
	if o == nil || o.EuroStatus == nil {
		var ret string
		return ret
	}
	return *o.EuroStatus
}

// GetEuroStatusOk returns a tuple with the EuroStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetEuroStatusOk() (*string, bool) {
	if o == nil || o.EuroStatus == nil {
		return nil, false
	}
	return o.EuroStatus, true
}

// HasEuroStatus returns a boolean if a field has been set.
func (o *Vehicle) HasEuroStatus() bool {
	if o != nil && o.EuroStatus != nil {
		return true
	}

	return false
}

// SetEuroStatus gets a reference to the given string and assigns it to the EuroStatus field.
func (o *Vehicle) SetEuroStatus(v string) {
	o.EuroStatus = &v
}

// GetAutomatedVehicle returns the AutomatedVehicle field value if set, zero value otherwise.
func (o *Vehicle) GetAutomatedVehicle() bool {
	if o == nil || o.AutomatedVehicle == nil {
		var ret bool
		return ret
	}
	return *o.AutomatedVehicle
}

// GetAutomatedVehicleOk returns a tuple with the AutomatedVehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetAutomatedVehicleOk() (*bool, bool) {
	if o == nil || o.AutomatedVehicle == nil {
		return nil, false
	}
	return o.AutomatedVehicle, true
}

// HasAutomatedVehicle returns a boolean if a field has been set.
func (o *Vehicle) HasAutomatedVehicle() bool {
	if o != nil && o.AutomatedVehicle != nil {
		return true
	}

	return false
}

// SetAutomatedVehicle gets a reference to the given bool and assigns it to the AutomatedVehicle field.
func (o *Vehicle) SetAutomatedVehicle(v bool) {
	o.AutomatedVehicle = &v
}

func (o Vehicle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["registrationNumber"] = o.RegistrationNumber
	}
	if o.TaxStatus != nil {
		toSerialize["taxStatus"] = o.TaxStatus
	}
	if o.TaxDueDate != nil {
		toSerialize["taxDueDate"] = o.TaxDueDate
	}
	if o.ArtEndDate != nil {
		toSerialize["artEndDate"] = o.ArtEndDate
	}
	if o.MotStatus != nil {
		toSerialize["motStatus"] = o.MotStatus
	}
	if o.MotExpiryDate != nil {
		toSerialize["motExpiryDate"] = o.MotExpiryDate
	}
	if o.Make != nil {
		toSerialize["make"] = o.Make
	}
	if o.MonthOfFirstDvlaRegistration != nil {
		toSerialize["monthOfFirstDvlaRegistration"] = o.MonthOfFirstDvlaRegistration
	}
	if o.MonthOfFirstRegistration != nil {
		toSerialize["monthOfFirstRegistration"] = o.MonthOfFirstRegistration
	}
	if o.YearOfManufacture != nil {
		toSerialize["yearOfManufacture"] = o.YearOfManufacture
	}
	if o.EngineCapacity != nil {
		toSerialize["engineCapacity"] = o.EngineCapacity
	}
	if o.Co2Emissions != nil {
		toSerialize["co2Emissions"] = o.Co2Emissions
	}
	if o.FuelType != nil {
		toSerialize["fuelType"] = o.FuelType
	}
	if o.MarkedForExport != nil {
		toSerialize["markedForExport"] = o.MarkedForExport
	}
	if o.Colour != nil {
		toSerialize["colour"] = o.Colour
	}
	if o.TypeApproval != nil {
		toSerialize["typeApproval"] = o.TypeApproval
	}
	if o.Wheelplan != nil {
		toSerialize["wheelplan"] = o.Wheelplan
	}
	if o.RevenueWeight != nil {
		toSerialize["revenueWeight"] = o.RevenueWeight
	}
	if o.RealDrivingEmissions != nil {
		toSerialize["realDrivingEmissions"] = o.RealDrivingEmissions
	}
	if o.DateOfLastV5CIssued != nil {
		toSerialize["dateOfLastV5CIssued"] = o.DateOfLastV5CIssued
	}
	if o.EuroStatus != nil {
		toSerialize["euroStatus"] = o.EuroStatus
	}
	if o.AutomatedVehicle != nil {
		toSerialize["automatedVehicle"] = o.AutomatedVehicle
	}
	return json.Marshal(toSerialize)
}

type NullableVehicle struct {
	value *Vehicle
	isSet bool
}

func (v NullableVehicle) Get() *Vehicle {
	return v.value
}

func (v *NullableVehicle) Set(val *Vehicle) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicle) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicle(val *Vehicle) *NullableVehicle {
	return &NullableVehicle{value: val, isSet: true}
}

func (v NullableVehicle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

