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
	"unsafe"
)

type callbackType func(int)

const (
	LA_USER      uint = 0
	LA_SYSTEM    uint = 1
	LA_IN_MEMORY uint = 2
)

var floatingLicenseCallbackFunction callbackType

//export floatingLicenseCallbackWrapper
func floatingLicenseCallbackWrapper(status int) {
	if floatingLicenseCallbackFunction != nil {
		floatingLicenseCallbackFunction(status)
	}
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
    * libraryVersion - pointer to a buffer that receives the value of the string
    * length - size of the buffer pointed to by the libraryVersion parameter
    
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
    * name - pointer to a buffer that receives the value of the string
    * length - size of the buffer pointed to by the name parameter

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
    * displayName - pointer to a buffer that receives the value of the string
    * length - size of the buffer pointed to by the displayName parameter

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
    * data - pointer to a buffer that receives the value of the string
    * length - size of the buffer pointed to by the data parameter

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
    FUNCTION: GetHostLicenseMetadata()

    PURPOSE: Get the value of the license metadata field associated with the LexFloatServer license.

    PARAMETERS:
    * key - key of the metadata field whose value you want to get
    * value - pointer to a buffer that receives the value of the string
    * length - size of the buffer pointed to by the value parameter

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
    * allowedUses - pointer to the integer that receives the value
    * totalUses - pointer to the integer that receives the value
    * grossUses - pointer to the integer that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_METER_ATTRIBUTE_NOT_FOUND
*/
func GetHostLicenseMeterAttribute(name string, allowedUses *uint, totalUses *uint, grossUses *uint) int {
	cName := goToCString(name)
	var cAllowedUses C.uint
	var cTotalUses C.uint
	var cGrossUses C.uint
	status := C.GetHostLicenseMeterAttribute(cName, &cAllowedUses, &cTotalUses, &cGrossUses)
	*allowedUses = uint(cAllowedUses)
	*totalUses = uint(cTotalUses)
	*grossUses = uint(cGrossUses)
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
    * value - pointer to a buffer that receives the value of the string
    
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
