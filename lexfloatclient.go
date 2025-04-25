// Copyright 2021 Cryptlex, LLC. All rights reserved.

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
/*
    FUNCTION: SetPermissionFlag()

    PURPOSE: Sets the permission flag.

    This function must be called on every start of your program after SetHostProductId()
    function in case the application allows borrowing of licenses or system wide activation.

    PARAMETERS:
    * flags - depending on your application's requirements, choose one of 
      the following values: LF_USER, LF_ALL_USERS.

      - LF_USER: This flag indicates that the application does not require
        admin or root permissions to run.

      - LF_ALL_USERS: This flag is specifically designed for Windows and should be used 
        for system-wide activations.

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID
*/
func SetPermissionFlag(flags uint) int {
	cFlags := (C.uint)(flags)
	status := C.SetPermissionFlag(cFlags)
	return int(status)
}

/*
    FUNCTION: SetHostProductId()

    PURPOSE: Sets the product id of your application.

    PARAMETERS:
    * productId - the unique product id of your application as mentioned
      on the product page in the dashboard.

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID
*/
func SetHostProductId(productId string) int {
	cProductId := goToCString(productId)
	status := C.SetHostProductId(cProductId)
	freeCString(cProductId)
	return int(status)
}

/*
    FUNCTION: SetHostUrl()

    PURPOSE: Sets the network address of the LexFloatServer.

    The url format should be: http://[ip or hostname]:[port]

    PARAMETERS:
    * hostUrl - url string having the correct format

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_HOST_URL
*/
func SetHostUrl(hostUrl string) int {
	cHostUrl := goToCString(hostUrl)
	status := C.SetHostUrl(cHostUrl)
	freeCString(cHostUrl)
	return int(status)
}

/*
    FUNCTION: SetFloatingLicenseCallback()

    PURPOSE: Sets the renew license callback function.

    Whenever the license lease is about to expire, a renew request is sent to the
    server. When the request completes, the license callback function
    gets invoked with one of the following status codes:

    LF_OK, LF_E_INET, LF_E_LICENSE_EXPIRED_INET, LF_E_LICENSE_NOT_FOUND, LF_E_CLIENT, LF_E_IP,
    LF_E_SERVER, LF_E_TIME, LF_E_SERVER_LICENSE_NOT_ACTIVATED,LF_E_SERVER_TIME_MODIFIED,
    LF_E_SERVER_LICENSE_SUSPENDED, LF_E_SERVER_LICENSE_EXPIRED, LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER

    PARAMETERS:
    * callback - name of the callback function

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID
*/
func SetFloatingLicenseCallback(callbackFunction func(int)) int {
	status := C.SetFloatingLicenseCallback((C.CallbackType)(unsafe.Pointer(C.floatingLicenseCallbackCgoGateway)))
	floatingLicenseCallbackFunction = callbackFunction
	return int(status)
}

/*
    FUNCTION: SetFloatingClientMetadata()

    PURPOSE: Sets the floating client metadata.

    The  metadata appears along with the license details of the license
    in LexFloatServer dashboard.

    PARAMETERS:
    * key - string of maximum length 256 characters with utf-8 encoding.
    * value - string of maximum length 4096 characters with utf-8 encoding.

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_METADATA_KEY_LENGTH,
    LF_E_METADATA_VALUE_LENGTH, LF_E_ACTIVATION_METADATA_LIMIT
*/
func SetFloatingClientMetadata(key string, value string) int {
	cKey := goToCString(key)
	cValue := goToCString(value)
	status := C.SetFloatingClientMetadata(cKey, cValue)
	freeCString(cKey)
	freeCString(cValue)
	return int(status)
}

/*
    FUNCTION: GetFloatingClientLibraryVersion()

    PURPOSE: Gets the version of this library.

    PARAMETERS:
    * libraryVersion - pointer to a string that receives the value

    RETURN CODES: LF_OK, LF_E_BUFFER_SIZE
*/
func GetFloatingClientLibraryVersion(libraryVersion *string) int {
	var cLibraryVersion = getCArray()
	status := C.GetFloatingClientLibraryVersion(&cLibraryVersion[0], maxCArrayLength)
	*libraryVersion = ctoGoString(&cLibraryVersion[0])
	return int(status)
}

/*
    FUNCTION: GetHostProductVersionName()

    PURPOSE: Gets the product version name.

    PARAMETERS:
    * name - pointer to a string that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_PRODUCT_VERSION_NOT_LINKED, LF_E_BUFFER_SIZE
*/
func GetHostProductVersionName(name *string) int {
	var cName = getCArray()
	status := C.GetHostProductVersionName(&cName[0], maxCArrayLength)
	*name = ctoGoString(&cName[0])
	return int(status)
}

/*
    FUNCTION: GetHostProductVersionDisplayName()

    PURPOSE: Gets the product version display name.

    PARAMETERS:
    * displayName - pointer to a string that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_PRODUCT_VERSION_NOT_LINKED, LF_E_BUFFER_SIZE
*/
func GetHostProductVersionDisplayName(displayName *string) int {
	var cDisplayName = getCArray()
	status := C.GetHostProductVersionDisplayName(&cDisplayName[0], maxCArrayLength)
	*displayName = ctoGoString(&cDisplayName[0])
	return int(status)
}

/*
    FUNCTION: GetHostProductVersionFeatureFlag()

    PURPOSE: Gets the product version feature flag.

    PARAMETERS:
    * name - name of the feature flag
    * enabled - pointer to the integer that receives the value - 0 or 1
    * data - pointer to a string that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_PRODUCT_VERSION_NOT_LINKED, LF_E_FEATURE_FLAG_NOT_FOUND, LF_E_BUFFER_SIZE
*/
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

/*
    FUNCTION: GetHostLicenseEntitlementSetName()

    PURPOSE: Gets the name of the entitlement set associated with the LexFloatServer license.

    PARAMETERS:
    * name - pointer to a string that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE, LF_E_ENTITLEMENT_SET_NOT_LINKED
*/
func GetHostLicenseEntitlementSetName(name *string) int {
    var cName = getCArray()
    status := C.GetHostLicenseEntitlementSetName(&cName[0], maxCArrayLength)
    *name = ctoGoString(&cName[0])
    return int(status)
}

/*
    FUNCTION: GetHostLicenseEntitlementSetDisplayName()

    PURPOSE: Gets the display name of the entitlement set associated with the LexFloatServer license.

    PARAMETERS:
    * displayName - pointer to a string that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE, LF_E_ENTITLEMENT_SET_NOT_LINKED
*/
func GetHostLicenseEntitlementSetDisplayName(displayName *string) int {
    var cDisplayName = getCArray()
    status := C.GetHostLicenseEntitlementSetDisplayName(&cDisplayName[0], maxCArrayLength)
    *displayName = ctoGoString(&cDisplayName[0])
    return int(status)
}

/*
    FUNCTION: GetHostFeatureEntitlements()

    PURPOSE: Gets the feature entitlements associated with the LexFloatServer license.

    PARAMETERS:
    * hostFeatureEntitlements - pointer to an array of HostFeatureEntitlement structs that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE
*/
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

/*
    FUNCTION: GetHostFeatureEntitlement()

    PURPOSE: Get the value of the feature entitlement field associated with the LexFloatServer license.

    PARAMETERS:
    * name - name of the feature
    * hostFeatureEntitlement - pointer to the HostFeatureEntitlement struct that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE, LF_E_FEATURE_ENTITLEMENT_NOT_FOUND
*/
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
/*
    FUNCTION: GetHostLicenseMetadata()

    PURPOSE: Get the value of the license metadata field associated with the LexFloatServer license.

    PARAMETERS:
    * key - key of the metadata field whose value you want to get
    * value - pointer to a string that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE,
    LF_E_METADATA_KEY_NOT_FOUND
*/
func GetHostLicenseMetadata(key string, value *string) int {
	cKey := goToCString(key)
	var cValue = getCArray()
	status := C.GetHostLicenseMetadata(cKey, &cValue[0], maxCArrayLength)
	*value = ctoGoString(&cValue[0])
	freeCString(cKey)
	return int(status)
}

/*
    FUNCTION: GetHostLicenseMeterAttribute()

    PURPOSE: Gets the license meter attribute allowed uses and total uses associated with the LexFloatServer license.

    PARAMETERS:
    * name - name of the meter attribute
    * allowedUses - pointer to the integer that receives the value. A value of -1 indicates unlimited allowed uses.
    * totalUses - pointer to the integer that receives the value
    * grossUses - pointer to the integer that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_METER_ATTRIBUTE_NOT_FOUND
*/
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

/*
    FUNCTION: GetHostLicenseExpiryDate()

    PURPOSE: Gets the license expiry date timestamp of the LexFloatServer license.

    PARAMETERS:
    * expiryDate - pointer to the integer that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE
*/
func GetHostLicenseExpiryDate(expiryDate *uint) int {
	var cExpiryDate C.uint
	status := C.GetHostLicenseExpiryDate(&cExpiryDate)
	*expiryDate = uint(cExpiryDate)
	return int(status)
}

/*
    FUNCTION: GetFloatingClientMeterAttributeUses()

    PURPOSE: Gets the meter attribute uses consumed by the floating client.

    PARAMETERS:
    * name - name of the meter attribute
    * uses - pointer to the integer that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_METER_ATTRIBUTE_NOT_FOUND
*/
func GetFloatingClientMeterAttributeUses(name string, uses *uint) int {
	cName := goToCString(name)
	var cUses C.uint
	status := C.GetFloatingClientMeterAttributeUses(cName, &cUses)
	*uses = uint(cUses)
	freeCString(cName)
	return int(status)
}

/*
    FUNCTION: GetFloatingClientMetadata()

    PURPOSE: Gets the value of the floating client metadata.

    PARAMETERS:
    * key - key of the metadata field whose value you want to retrieve
    * value - pointer to a string that receives the value
    
    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE,
    LF_E_METADATA_KEY_NOT_FOUND
*/
func GetFloatingClientMetadata(key string, value *string) int {
	cKey := goToCString(key)
	var cValue = getCArray()
	status := C.GetFloatingClientMetadata(cKey, &cValue[0], maxCArrayLength)
	*value = ctoGoString(&cValue[0])
	freeCString(cKey)
	return int(status)

}

/*
    FUNCTION: RequestFloatingLicense()

    PURPOSE: Sends the request to lease the license from the LexFloatServer.

    RETURN CODES: LF_OK, LF_FAIL, LF_E_PRODUCT_ID, LF_E_LICENSE_EXISTS, LF_E_HOST_URL,
    LF_E_CALLBACK, LF_E_LICENSE_LIMIT_REACHED, LF_E_INET, LF_E_TIME, LF_E_CLIENT, LF_E_IP, LF_E_SERVER,
    LF_E_SERVER_LICENSE_NOT_ACTIVATED, LF_E_SERVER_TIME_MODIFIED, LF_E_SERVER_LICENSE_SUSPENDED,
    LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER, LF_E_SERVER_LICENSE_EXPIRED
*/
func RequestFloatingLicense() int {
	status := C.RequestFloatingLicense()
	return int(status)
}

/*
    FUNCTION: GetFloatingClientLeaseExpiryDate()

    PURPOSE: Gets the lease expiry date timestamp of the floating client.

    PARAMETERS:
    * leaseExpiryDate - pointer to the integer that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE
*/
func GetFloatingClientLeaseExpiryDate(leaseExpiryDate *uint) int {
	var cLeaseExpiryDate C.uint
	status := C.GetFloatingClientLeaseExpiryDate(&cLeaseExpiryDate)
	*leaseExpiryDate = uint(cLeaseExpiryDate)
	return int(status)
}

/*
    FUNCTION: DropFloatingLicense()

    PURPOSE: Sends the request to the LexFloatServer to free the license.

    Call this function before you exit your application to prevent zombie licenses.

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_HOST_URL, LF_E_CALLBACK,
    LF_E_INET, LF_E_LICENSE_NOT_FOUND, LF_E_CLIENT, LF_E_IP, LF_E_SERVER,
    LF_E_SERVER_LICENSE_NOT_ACTIVATED, LF_E_SERVER_TIME_MODIFIED, LF_E_SERVER_LICENSE_SUSPENDED,
    LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER, LF_E_SERVER_LICENSE_EXPIRED
*/
func DropFloatingLicense() int {
	status := C.DropFloatingLicense()
	return int(status)
}

/*
    FUNCTION: HasFloatingLicense()

    PURPOSE: Checks whether any license has been leased or not. If yes,
    it retuns LF_OK.

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE
*/
func HasFloatingLicense() int {
	status := C.HasFloatingLicense()
	return int(status)
}

/*
	FUNCTION: GetHostConfig()

	PURPOSE: Gets the host configuration.

	This function sends a network request to LexFloatServer to get the configuration details.

	PARAMETERS:
	* hostConfig - pointer to the HostConfig struct that receives the value

	RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_HOST_URL, LF_E_BUFFER_SIZE
	LF_E_INET, LF_E_CLIENT, LF_E_IP, LF_E_SERVER
*/
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

/*

   FUNCTION: GetFloatinglicenseMode()

   PURPOSE: Gets the mode of the floating license (online or offline).

   PARAMETERS:
   * mode - pointer to a string that receives the value

   RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_BUFFER_SIZE
*/
func GetFloatingLicenseMode(mode *string) int {
	var cMode = getCArray()
	status := C.GetFloatingLicenseMode(&cMode[0], maxCArrayLength)
	*mode= ctoGoString(&cMode[0])
	return int(status)
}

/*
    FUNCTION: RequestOfflineFloatingLicense()

    PURPOSE: Sends the request to lease the license from the LexFloatServer for offline usage.

    The maximum value of lease duration is configured in the config.yml of LexFloatServer 

    PARAMETERS:
    * leaseDuration - value of the lease duration.

    RETURN CODES: LF_OK, LF_FAIL, LF_E_PRODUCT_ID, LF_E_LICENSE_EXISTS, LF_E_HOST_URL,
    LF_E_LICENSE_LIMIT_REACHED, LF_E_INET, LF_E_TIME, LF_E_CLIENT, LF_E_IP, LF_E_SERVER,
    LF_E_SERVER_LICENSE_NOT_ACTIVATED, LF_E_SERVER_TIME_MODIFIED, LF_E_SERVER_LICENSE_SUSPENDED,
    LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER, LF_E_SERVER_LICENSE_EXPIRED, LF_E_WMIC, LF_E_SYSTEM_PERMISSION
*/
func RequestOfflineFloatingLicense(leaseDuration uint) int {
    cLeaseDuration := (C.uint)(leaseDuration)
    status := C.RequestOfflineFloatingLicense(cLeaseDuration)
    return int(status)
}

/*
    FUNCTION: IncrementFloatingClientMeterAttributeUses()

    PURPOSE: Increments the meter attribute uses of the floating client.

    PARAMETERS:
    * name - name of the meter attribute
    * increment - the increment value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_HOST_URL, LF_E_METER_ATTRIBUTE_NOT_FOUND,
    LF_E_INET, LF_E_LICENSE_NOT_FOUND, LF_E_CLIENT, LF_E_IP, LF_E_SERVER, LF_E_METER_ATTRIBUTE_USES_LIMIT_REACHED,
    LF_E_SERVER_LICENSE_NOT_ACTIVATED, LF_E_SERVER_TIME_MODIFIED, LF_E_SERVER_LICENSE_SUSPENDED,
    LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER, LF_E_SERVER_LICENSE_EXPIRED
*/
func IncrementFloatingClientMeterAttributeUses(name string, increment uint) int {
	cName := goToCString(name)
	cIncrement := (C.uint)(increment)
	status := C.IncrementFloatingClientMeterAttributeUses(cName, cIncrement)
	freeCString(cName)
	return int(status)
}

/*
    FUNCTION: DecrementFloatingClientMeterAttributeUses()

    PURPOSE: Decrements the meter attribute uses of the floating client.

    PARAMETERS:
    * name - name of the meter attribute
    * decrement - the decrement value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_HOST_URL, LF_E_METER_ATTRIBUTE_NOT_FOUND,
    LF_E_INET, LF_E_LICENSE_NOT_FOUND, LF_E_CLIENT, LF_E_IP, LF_E_SERVER,
    LF_E_SERVER_LICENSE_NOT_ACTIVATED, LF_E_SERVER_TIME_MODIFIED, LF_E_SERVER_LICENSE_SUSPENDED,
    LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER, LF_E_SERVER_LICENSE_EXPIRED

    NOTE: If the decrement is more than the current uses, it resets the uses to 0.
*/
func DecrementFloatingClientMeterAttributeUses(name string, decrement uint) int {
	cName := goToCString(name)
	cDecrement := (C.uint)(decrement)
	status := C.DecrementFloatingClientMeterAttributeUses(cName, cDecrement)
	freeCString(cName)
	return int(status)
}

/*
    FUNCTION: ResetFloatingClientMeterAttributeUses()

    PURPOSE: Resets the meter attribute uses consumed by the floating client.

    PARAMETERS:
    * name - name of the meter attribute

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_HOST_URL, LF_E_METER_ATTRIBUTE_NOT_FOUND,
    LF_E_INET, LF_E_LICENSE_NOT_FOUND, LF_E_CLIENT, LF_E_IP, LF_E_SERVER,
    LF_E_SERVER_LICENSE_NOT_ACTIVATED, LF_E_SERVER_TIME_MODIFIED, LF_E_SERVER_LICENSE_SUSPENDED,
    LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER, LF_E_SERVER_LICENSE_EXPIRED
*/
func ResetFloatingClientMeterAttributeUses(name string) int {
	cName := goToCString(name)
	status := C.ResetFloatingClientMeterAttributeUses(cName)
	freeCString(cName)
	return int(status)
}
