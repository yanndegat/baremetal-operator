package bmc

import (
	"net/url"
)

func init() {
	registerFactory("ovh", newOVHAccessDetails)
}

func newOVHAccessDetails(parsedURL *url.URL, disableCertificateVerification bool) (AccessDetails, error) {
	return &ovhAccessDetails{
		bmcType:    parsedURL.Scheme,
		serverName: parsedURL.Hostname(),
	}, nil
}

type ovhAccessDetails struct {
	bmcType    string
	serverName string
}

func (a *ovhAccessDetails) Type() string {
	return a.bmcType
}

func (a *ovhAccessDetails) NeedsMAC() bool {
	return false
}

func (a *ovhAccessDetails) Driver() string {
	return "ovhapi"
}

func (a *ovhAccessDetails) DisableCertificateVerification() bool {
	return true
}

// DriverInfo returns a data structure to pass as the DriverInfo
// parameter when creating a node in Ironic. The structure is
// pre-populated with the access information, and the caller is
// expected to add any other information that might be needed (such as
// the kernel and ramdisk locations).
func (a *ovhAccessDetails) DriverInfo(bmcCreds Credentials) map[string]interface{} {
	result := map[string]interface{}{
		"server_name": a.serverName,
	}
	return result
}

func (a *ovhAccessDetails) BootInterface() string {
	return ""
}

func (a *ovhAccessDetails) ManagementInterface() string {
	return "ovhapi"
}

func (a *ovhAccessDetails) PowerInterface() string {
	return "ovhapi"
}

func (a *ovhAccessDetails) RAIDInterface() string {
	return ""
}

func (a *ovhAccessDetails) VendorInterface() string {
	return ""
}
