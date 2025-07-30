// Copyright 2025 Cryptlex, LLC. All rights reserved.

package lexfloatclient

/*
#cgo linux,!arm64 LDFLAGS: -Wl,-Bstatic -L${SRCDIR}/libs/linux_amd64 -lLexFloatClient -Wl,-Bdynamic -lm -lstdc++ -lpthread
#cgo linux,arm64 LDFLAGS: -Wl,-Bstatic -L${SRCDIR}/libs/linux_arm64 -lLexFloatClient -Wl,-Bdynamic -lm -lstdc++ -lpthread
#cgo darwin LDFLAGS: -L${SRCDIR}/libs/darwin_universal -lLexFloatClient -lc++ -framework CoreFoundation -framework SystemConfiguration -framework Security
#cgo windows LDFLAGS: -L${SRCDIR}/libs/windows_amd64 -lLexFloatClient
#include "lexfloatclient/LexFloatClient.h"
#include <stdlib.h>
void floatingLicenseCallbackCgoGateway(int status);
*/
import "C"
import (
	"encoding/json"
	"strings"
	"unsafe"
)


type HostConfig struct {
	MaxOfflineLeaseDuration int `json:"maxOfflineLeaseDuration"`
}

type HostFeatureEntitlement struct {
	FeatureName          string `json:"featureName"`
	FeatureDisplayName   string `json:"featureDisplayName"`
	Value                string `json:"value"`
}

type callbackType func(int)

const (
    LF_USER      uint = 10
    LF_ALL_USERS uint = 11
)

var floatingLicenseCallbackFunction callbackType

//export floatingLicenseCallbackWrapper
func floatingLicenseCallbackWrapper(status int) {
	if floatingLicenseCallbackFunction != nil {
		floatingLicenseCallbackFunction(status)
	}
}

// SetPermissionFlag sets the permission flag.
//
// This function must be called on every start of your program after SetHostProductId()
// function in case the application allows borrowing of licenses or system wide activation.
//
// Parameters:
// - flags: depending on your application's requirements, choose one of the following values:
//   - LF_USER: This flag indicates that the application does not require admin or root permissions to run.
//   - LF_ALL_USERS: This flag is specifically designed for Windows and should be used 
//     for system-wide activations.
//
// Returns: LF_OK, LF_E_PRODUCT_ID
func SetPermissionFlag(flags uint) int {
	cFlags := (C.uint)(flags)
	status := C.SetPermissionFlag(cFlags)
	return int(status)
}

// SetHostProductId sets the product id of your application.
//
// Parameters:
// - productId: the unique product id of your application as mentioned
//   on the product page in the dashboard.
//
// Returns: LF_OK, LF_E_PRODUCT_ID
func SetHostProductId(productId string) int {
	cProductId := goToCString(productId)
	status := C.SetHostProductId(cProductId)
	freeCString(cProductId)
	return int(status)
}

// SetHostUrl sets the network address of the LexFloatServer.
//
// The url format should be: http://[ip or hostname]:[port]
//
// Parameters:
// - hostUrl: the url string having the correct format
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_HOST_URL
func SetHostUrl(hostUrl string) int {
	cHostUrl := goToCString(hostUrl)
	status := C.SetHostUrl(cHostUrl)
	freeCString(cHostUrl)
	return int(status)
}

// SetFloatingLicenseCallback sets the renew license callback function.
//
// Whenever the license lease is about to expire, a renew request is sent to the
// server. When the request completes, the license callback function
// gets invoked with one of the following status codes:
//
// LF_OK, LF_E_INET, LF_E_LICENSE_EXPIRED_INET, LF_E_LICENSE_NOT_FOUND, LF_E_CLIENT, LF_E_IP,
// LF_E_SERVER, LF_E_TIME, LF_E_SERVER_LICENSE_NOT_ACTIVATED,LF_E_SERVER_TIME_MODIFIED,
// LF_E_SERVER_LICENSE_SUSPENDED, LF_E_SERVER_LICENSE_EXPIRED, LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER
// 
// Returns: LF_OK, LF_E_PRODUCT_ID
func SetFloatingLicenseCallback(callbackFunction func(int)) int {
	status := C.SetFloatingLicenseCallback((C.CallbackType)(unsafe.Pointer(C.floatingLicenseCallbackCgoGateway)))
	floatingLicenseCallbackFunction = callbackFunction
	return int(status)
}

// SetFloatingClientMetadata sets the floating client metadata.
//
// The metadata appears along with the license details of the license
// in LexFloatServer dashboard.
//
// Parameters:
// - key: string of maximum length 256 characters with utf-8 encoding.
// - value: string of maximum length 4096 characters with utf-8 encoding.
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_METADATA_KEY_LENGTH,
// LF_E_METADATA_VALUE_LENGTH, LF_E_ACTIVATION_METADATA_LIMIT
func SetFloatingClientMetadata(key string, value string) int {
	cKey := goToCString(key)
	cValue := goToCString(value)
	status := C.SetFloatingClientMetadata(cKey, cValue)
	freeCString(cKey)
	freeCString(cValue)
	return int(status)
}

// GetFloatingClientLibraryVersion gets the version of this library.
//
// Parameters:
// - libraryVersion: pointer to a string that receives the value
//
// Returns: LF_OK, LF_E_BUFFER_SIZE
func GetFloatingClientLibraryVersion(libraryVersion *string) int {
	var cLibraryVersion = getCArray()
	status := C.GetFloatingClientLibraryVersion(&cLibraryVersion[0], maxCArrayLength)
	*libraryVersion = ctoGoString(&cLibraryVersion[0])
	return int(status)
}

// GetHostProductVersionName gets the product version name.
//
// Deprecated: This function is deprecated. Use GetHostLicenseEntitlementSetName() instead.
//
// Parameters:
// - name: pointer to a string that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_PRODUCT_VERSION_NOT_LINKED, LF_E_BUFFER_SIZE
func GetHostProductVersionName(name *string) int {
	var cName = getCArray()
	status := C.GetHostProductVersionName(&cName[0], maxCArrayLength)
	*name = ctoGoString(&cName[0])
	return int(status)
}

// GetHostProductVersionDisplayName gets the product version display name.
//
// Deprecated: This function is deprecated. Use GetHostLicenseEntitlementSetDisplayName() instead.
//
// Parameters:
// - displayName: pointer to a string that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_PRODUCT_VERSION_NOT_LINKED, LF_E_BUFFER_SIZE
func GetHostProductVersionDisplayName(displayName *string) int {
	var cDisplayName = getCArray()
	status := C.GetHostProductVersionDisplayName(&cDisplayName[0], maxCArrayLength)
	*displayName = ctoGoString(&cDisplayName[0])
	return int(status)
}

// GetHostProductVersionFeatureFlag gets the product version feature flag.
//
// Deprecated: This function is deprecated. Use GetHostFeatureEntitlement() instead.
//
// Parameters:
// - name: name of the feature flag
// - enabled: pointer to the integer that receives the value - 0 or 1
// - data: pointer to a string that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_PRODUCT_VERSION_NOT_LINKED, LF_E_FEATURE_FLAG_NOT_FOUND, LF_E_BUFFER_SIZE
func GetHostProductVersionFeatureFlag(name string, enabled *bool, data *string) int {
    cName := goToCString(name)
    var cEnabled C.uint
    var cData = getCArray()
    status := C.GetHostProductVersionFeatureFlag(cName, &cEnabled, &cData[0], maxCArrayLength)
    freeCString(cName)
    *enabled = cEnabled > 0
    *data = ctoGoString(&cData[0])
    return int(status)
}

// GetHostLicenseEntitlementSetName gets the name of the entitlement set associated with the LexFloatServer license.
//
// Parameters:
// - name: pointer to a string that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE, LF_E_ENTITLEMENT_SET_NOT_LINKED
func GetHostLicenseEntitlementSetName(name *string) int {
    var cName = getCArray()
    status := C.GetHostLicenseEntitlementSetName(&cName[0], maxCArrayLength)
    *name = ctoGoString(&cName[0])
    return int(status)
}

// GetHostLicenseEntitlementSetDisplayName gets the display name of the entitlement set associated with the LexFloatServer license.
//
// Parameters:
// - displayName: pointer to a string that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE, LF_E_ENTITLEMENT_SET_NOT_LINKED
func GetHostLicenseEntitlementSetDisplayName(displayName *string) int {
    var cDisplayName = getCArray()
    status := C.GetHostLicenseEntitlementSetDisplayName(&cDisplayName[0], maxCArrayLength)
    *displayName = ctoGoString(&cDisplayName[0])
    return int(status)
}

// GetHostFeatureEntitlements gets the feature entitlements associated with the LexFloatServer license.
//
// Parameters:
// - hostFeatureEntitlements: pointer to an array of HostFeatureEntitlement structs that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE
func GetHostFeatureEntitlements(hostFeatureEntitlements *[]HostFeatureEntitlement) int {
    var cHostFeatureEntitlements = getCArray()
    hostFeatureEntitlementsJson := ""
    status := C.GetHostFeatureEntitlementsInternal(&cHostFeatureEntitlements[0], maxCArrayLength)
    hostFeatureEntitlementsJson = strings.TrimRight(ctoGoString(&cHostFeatureEntitlements[0]), "\x00")
    if hostFeatureEntitlementsJson != "" {
        json.Unmarshal([]byte(hostFeatureEntitlementsJson), hostFeatureEntitlements)
    }
    return int(status)
}

// GetHostFeatureEntitlement gets the value of the feature entitlement field associated with the LexFloatServer license.
//
// Parameters:
// - name: name of the feature
// - hostFeatureEntitlement: pointer to the HostFeatureEntitlement struct that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE, LF_E_FEATURE_ENTITLEMENT_NOT_FOUND
func GetHostFeatureEntitlement(name string, hostFeatureEntitlement *HostFeatureEntitlement) int {
	cName := goToCString(name)
	var cHostFeatureEntitlement = getCArray()
	status := C.GetHostFeatureEntitlementInternal(cName, &cHostFeatureEntitlement[0], maxCArrayLength)
	freeCString(cName)
	hostFeatureEntitlementJson := strings.TrimRight(ctoGoString(&cHostFeatureEntitlement[0]), "\x00")
	if hostFeatureEntitlementJson != "" {
		json.Unmarshal([]byte(hostFeatureEntitlementJson), hostFeatureEntitlement)
	}
	return int(status)
}

// GetHostProductMetadata gets the value of the field associated with the product-metadata key.
//
// Parameters:
// - key: key of the metadata field whose value you want to get
// - value: pointer to a string that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE, LF_E_METADATA_KEY_NOT_FOUND
func GetHostProductMetadata(key string, value *string) int {
	cKey := goToCString(key)
	var cValue = getCArray()
	status := C.GetHostProductMetadata(cKey, &cValue[0], maxCArrayLength)
	*value = ctoGoString(&cValue[0])
	freeCString(cKey)
	return int(status)
}

// GetHostLicenseMetadata gets the value of the license metadata field associated with the LexFloatServer license.
//
// Parameters:
// - key: key of the metadata field whose value you want to get
// - value: pointer to a string that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE, LF_E_METADATA_KEY_NOT_FOUND
func GetHostLicenseMetadata(key string, value *string) int {
	cKey := goToCString(key)
	var cValue = getCArray()
	status := C.GetHostLicenseMetadata(cKey, &cValue[0], maxCArrayLength)
	*value = ctoGoString(&cValue[0])
	freeCString(cKey)
	return int(status)
}

// GetHostLicenseMeterAttribute gets the license meter attribute allowed uses and total uses associated with the LexFloatServer license.
//
// Parameters:
// - name: name of the meter attribute
// - allowedUses: pointer to the integer that receives the value. A value of -1 indicates unlimited allowed uses.
// - totalUses: pointer to the integer that receives the value
// - grossUses: pointer to the integer that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_METER_ATTRIBUTE_NOT_FOUND
func GetHostLicenseMeterAttribute(name string, allowedUses *int64, totalUses *uint64, grossUses *uint64) int {
	cName := goToCString(name)
	var cAllowedUses C.int64_t
	var cTotalUses C.uint64_t
	var cGrossUses C.uint64_t
	status := C.GetHostLicenseMeterAttribute(cName, &cAllowedUses, &cTotalUses, &cGrossUses)
	*allowedUses = int64(cAllowedUses)
	*totalUses = uint64(cTotalUses)
	*grossUses = uint64(cGrossUses)
	freeCString(cName)
	return int(status)
}

// GetHostLicenseExpiryDate gets the license expiry date timestamp of the LexFloatServer license.
//
// Parameters:
// - expiryDate: pointer to the integer that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE
func GetHostLicenseExpiryDate(expiryDate *uint) int {
	var cExpiryDate C.uint
	status := C.GetHostLicenseExpiryDate(&cExpiryDate)
	*expiryDate = uint(cExpiryDate)
	return int(status)
}

// GetFloatingClientMeterAttributeUses gets the meter attribute uses consumed by the floating client.
//
// Parameters:
// - name: name of the meter attribute
// - uses: pointer to the integer that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_METER_ATTRIBUTE_NOT_FOUND
func GetFloatingClientMeterAttributeUses(name string, uses *uint) int {
	cName := goToCString(name)
	var cUses C.uint
	status := C.GetFloatingClientMeterAttributeUses(cName, &cUses)
	*uses = uint(cUses)
	freeCString(cName)
	return int(status)
}

// GetFloatingClientMetadata gets the value of the floating client metadata.
//
// Parameters:
// - key: key of the metadata field whose value you want to retrieve
// - value: pointer to a string that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE, LF_E_METADATA_KEY_NOT_FOUND
func GetFloatingClientMetadata(key string, value *string) int {
	cKey := goToCString(key)
	var cValue = getCArray()
	status := C.GetFloatingClientMetadata(cKey, &cValue[0], maxCArrayLength)
	*value = ctoGoString(&cValue[0])
	freeCString(cKey)
	return int(status)

}

// RequestFloatingLicense sends the request to lease the license from the LexFloatServer.
//
// Returns: LF_OK, LF_FAIL, LF_E_PRODUCT_ID, LF_E_LICENSE_EXISTS, LF_E_HOST_URL,
// LF_E_CALLBACK, LF_E_LICENSE_LIMIT_REACHED, LF_E_INET, LF_E_TIME, LF_E_CLIENT, LF_E_IP, LF_E_SERVER,
// LF_E_SERVER_LICENSE_NOT_ACTIVATED, LF_E_SERVER_TIME_MODIFIED, LF_E_SERVER_LICENSE_SUSPENDED,
// LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER, LF_E_SERVER_LICENSE_EXPIRED
func RequestFloatingLicense() int {
	status := C.RequestFloatingLicense()
	return int(status)
}

// GetFloatingClientLeaseExpiryDate gets the lease expiry date timestamp of the floating client.
//
// Parameters:
// - leaseExpiryDate: pointer to the integer that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE
func GetFloatingClientLeaseExpiryDate(leaseExpiryDate *uint) int {
	var cLeaseExpiryDate C.uint
	status := C.GetFloatingClientLeaseExpiryDate(&cLeaseExpiryDate)
	*leaseExpiryDate = uint(cLeaseExpiryDate)
	return int(status)
}

// DropFloatingLicense sends the request to the LexFloatServer to free the license.
//
// Call this function before you exit your application to prevent zombie licenses.
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_HOST_URL, LF_E_CALLBACK,
// LF_E_INET, LF_E_LICENSE_NOT_FOUND, LF_E_CLIENT, LF_E_IP, LF_E_SERVER,
// LF_E_SERVER_LICENSE_NOT_ACTIVATED, LF_E_SERVER_TIME_MODIFIED, LF_E_SERVER_LICENSE_SUSPENDED,
// LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER, LF_E_SERVER_LICENSE_EXPIRED
func DropFloatingLicense() int {
	status := C.DropFloatingLicense()
	return int(status)
}

// HasFloatingLicense checks whether any license has been leased or not. If yes,
// it returns LF_OK.
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE
func HasFloatingLicense() int {
	status := C.HasFloatingLicense()
	return int(status)
}

// GetHostConfig gets the host configuration.
//
// This function sends a network request to LexFloatServer to get the configuration details.
// Parameters:
// - hostConfig: pointer to the HostConfig struct that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_HOST_URL, LF_E_BUFFER_SIZE, LF_E_INET, LF_E_CLIENT, LF_E_IP, LF_E_SERVER
func GetHostConfig(hostConfig *HostConfig) int {
	var cHostConfig = getCArray()
	hostConfigJson := ""
	status := C.GetHostConfigInternal(&cHostConfig[0], maxCArrayLength)
	hostConfigJson = strings.TrimRight(ctoGoString(&cHostConfig[0]), "\x00")
	if hostConfigJson != "" {
		config := []byte(hostConfigJson)
		json.Unmarshal(config, hostConfig)
	}
	return int(status)
}

// GetFloatingLicenseMode gets the mode of the floating license (online or offline).
//
// Parameters:
// - mode: pointer to a string that receives the value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE
func GetFloatingLicenseMode(mode *string) int {
	var cMode = getCArray()
	status := C.GetFloatingLicenseMode(&cMode[0], maxCArrayLength)
	*mode= ctoGoString(&cMode[0])
	return int(status)
}

// RequestOfflineFloatingLicense sends the request to lease the license from the LexFloatServer for offline usage.
//
// The maximum value of lease duration is configured in the config.yml of LexFloatServer 
//
// Parameters:
// - leaseDuration: value of the lease duration.
//
// Returns: LF_OK, LF_FAIL, LF_E_PRODUCT_ID, LF_E_LICENSE_EXISTS, LF_E_HOST_URL,
// LF_E_LICENSE_LIMIT_REACHED, LF_E_INET, LF_E_TIME, LF_E_CLIENT, LF_E_IP, LF_E_SERVER,
// LF_E_SERVER_LICENSE_NOT_ACTIVATED, LF_E_SERVER_TIME_MODIFIED, LF_E_SERVER_LICENSE_SUSPENDED,
// LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER, LF_E_SERVER_LICENSE_EXPIRED, LF_E_WMIC, LF_E_SYSTEM_PERMISSION
func RequestOfflineFloatingLicense(leaseDuration uint) int {
    cLeaseDuration := (C.uint)(leaseDuration)
    status := C.RequestOfflineFloatingLicense(cLeaseDuration)
    return int(status)
}

// IncrementFloatingClientMeterAttributeUses increments the meter attribute uses of the floating client.
//
// Parameters:
// - name: name of the meter attribute
// - increment: the increment value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_HOST_URL, LF_E_METER_ATTRIBUTE_NOT_FOUND,
// LF_E_INET, LF_E_LICENSE_NOT_FOUND, LF_E_CLIENT, LF_E_IP, LF_E_SERVER, LF_E_METER_ATTRIBUTE_USES_LIMIT_REACHED,
// LF_E_SERVER_LICENSE_NOT_ACTIVATED, LF_E_SERVER_TIME_MODIFIED, LF_E_SERVER_LICENSE_SUSPENDED,
// LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER, LF_E_SERVER_LICENSE_EXPIRED
func IncrementFloatingClientMeterAttributeUses(name string, increment uint) int {
	cName := goToCString(name)
	cIncrement := (C.uint)(increment)
	status := C.IncrementFloatingClientMeterAttributeUses(cName, cIncrement)
	freeCString(cName)
	return int(status)
}
    
// DecrementFloatingClientMeterAttributeUses decrements the meter attribute uses of the floating client.
//
// Parameters:
// - name: name of the meter attribute
// - decrement: the decrement value
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_HOST_URL, LF_E_METER_ATTRIBUTE_NOT_FOUND,
// LF_E_INET, LF_E_LICENSE_NOT_FOUND, LF_E_CLIENT, LF_E_IP, LF_E_SERVER, LF_E_METER_ATTRIBUTE_USES_LIMIT_REACHED,
// LF_E_SERVER_LICENSE_NOT_ACTIVATED, LF_E_SERVER_TIME_MODIFIED, LF_E_SERVER_LICENSE_SUSPENDED,
// LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER, LF_E_SERVER_LICENSE_EXPIRED
func DecrementFloatingClientMeterAttributeUses(name string, decrement uint) int {
	cName := goToCString(name)
	cDecrement := (C.uint)(decrement)
	status := C.DecrementFloatingClientMeterAttributeUses(cName, cDecrement)
	freeCString(cName)
	return int(status)
}

// ResetFloatingClientMeterAttributeUses resets the meter attribute uses consumed by the floating client.
//
// Parameters:
// - name: name of the meter attribute
//
// Returns: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_HOST_URL, LF_E_METER_ATTRIBUTE_NOT_FOUND,
// LF_E_INET, LF_E_LICENSE_NOT_FOUND, LF_E_CLIENT, LF_E_IP, LF_E_SERVER, LF_E_SERVER_LICENSE_NOT_ACTIVATED,
// LF_E_SERVER_TIME_MODIFIED, LF_E_SERVER_LICENSE_SUSPENDED, LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER, LF_E_SERVER_LICENSE_EXPIRED
func ResetFloatingClientMeterAttributeUses(name string) int {
	cName := goToCString(name)
	status := C.ResetFloatingClientMeterAttributeUses(cName)
	freeCString(cName)
	return int(status)
}
