// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetSystemInfoResponse get system info response
// swagger:model getSystemInfoResponse
type GetSystemInfoResponse struct {

	// data
	Data []*GetSystemInfoResponseDataItems0 `json:"data"`
}

// Validate validates this get system info response
func (m *GetSystemInfoResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSystemInfoResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSystemInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSystemInfoResponse) UnmarshalBinary(b []byte) error {
	var res GetSystemInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetSystemInfoResponseDataItems0 get system info response data items0
// swagger:model GetSystemInfoResponseDataItems0
type GetSystemInfoResponseDataItems0 struct {

	// autobackup
	Autobackup bool `json:"autobackup,omitempty"`

	// build
	Build string `json:"build,omitempty"`

	// data retention days
	DataRetentionDays int64 `json:"data_retention_days,omitempty"`

	// data retention time in hours for 5minutes scale
	DataRetentionTimeInHoursFor5minutesScale int64 `json:"data_retention_time_in_hours_for_5minutes_scale,omitempty"`

	// data retention time in hours for daily scale
	DataRetentionTimeInHoursForDailyScale int64 `json:"data_retention_time_in_hours_for_daily_scale,omitempty"`

	// data retention time in hours for hourly scale
	DataRetentionTimeInHoursForHourlyScale int64 `json:"data_retention_time_in_hours_for_hourly_scale,omitempty"`

	// data retention time in hours for monthly scale
	DataRetentionTimeInHoursForMonthlyScale int64 `json:"data_retention_time_in_hours_for_monthly_scale,omitempty"`

	// data retention time in hours for others
	DataRetentionTimeInHoursForOthers int64 `json:"data_retention_time_in_hours_for_others,omitempty"`

	// debug device
	DebugDevice string `json:"debug_device,omitempty"`

	// debug mgmt
	DebugMgmt string `json:"debug_mgmt,omitempty"`

	// debug sdn
	DebugSdn string `json:"debug_sdn,omitempty"`

	// debug system
	DebugSystem string `json:"debug_system,omitempty"`

	// default site device auth password alert
	DefaultSiteDeviceAuthPasswordAlert bool `json:"default_site_device_auth_password_alert,omitempty"`

	// facebook wifi registered
	FacebookWifiRegistered bool `json:"facebook_wifi_registered,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// https port
	HTTPSPort int64 `json:"https_port,omitempty"`

	// image maps use google engine
	ImageMapsUseGoogleEngine bool `json:"image_maps_use_google_engine,omitempty"`

	// inform port
	InformPort int64 `json:"inform_port,omitempty"`

	// ip addrs
	IPAddrs []string `json:"ip_addrs"`

	// live chat
	LiveChat string `json:"live_chat,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// override inform host
	OverrideInformHost bool `json:"override_inform_host,omitempty"`

	// radius disconnect running
	RadiusDisconnectRunning bool `json:"radius_disconnect_running,omitempty"`

	// store enabled
	StoreEnabled string `json:"store_enabled,omitempty"`

	// timezone
	Timezone string `json:"timezone,omitempty"`

	// unifi go enabled
	UnifiGoEnabled bool `json:"unifi_go_enabled,omitempty"`

	// unsupported device count
	UnsupportedDeviceCount int64 `json:"unsupported_device_count,omitempty"`

	// update available
	UpdateAvailable bool `json:"update_available,omitempty"`

	// update downloaded
	UpdateDownloaded bool `json:"update_downloaded,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this get system info response data items0
func (m *GetSystemInfoResponseDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetSystemInfoResponseDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSystemInfoResponseDataItems0) UnmarshalBinary(b []byte) error {
	var res GetSystemInfoResponseDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
