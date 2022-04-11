/* LexFloatClient.h */
#pragma once

#include <stdint.h>
#include <stddef.h>
#include "LexFloatStatusCodes.h"

#ifdef _WIN32
    /*
    Make sure you're using the MSVC or Intel compilers on Windows.
    */
    #include <windows.h>

	#ifdef LEXFLOATCLIENT_EXPORTS
        #ifdef LEXFLOATCLIENT_STATIC
            #define LEXFLOATCLIENT_API extern "C"
        #else
            #define LEXFLOATCLIENT_API extern "C" __declspec(dllexport)
        #endif
	#else
		#ifdef __cplusplus
            #ifdef LEXFLOATCLIENT_STATIC
                #define LEXFLOATCLIENT_API extern "C"
            #else
                #define LEXFLOATCLIENT_API extern "C" __declspec(dllimport)
            #endif
        #else
            #ifdef LEXFLOATCLIENT_STATIC
                #define LEXFLOATCLIENT_API
            #else
                #define LEXFLOATCLIENT_API __declspec(dllimport)
            #endif
        #endif
	#endif

    #if defined(USE_STDCALL_DLL) && !defined(LEXFLOATCLIENT_STATIC)
        #define LF_CC __stdcall
    #else
        #define LF_CC __cdecl
    #endif
    typedef const wchar_t* CSTRTYPE;
	typedef wchar_t* STRTYPE;
#else
    #define LF_CC
    #if __GNUC__ >= 4
        #ifdef __cplusplus
            #define LEXFLOATCLIENT_API extern "C" __attribute__((visibility("default")))
        #else
            #define LEXFLOATCLIENT_API __attribute__((visibility("default")))
        #endif
    #else
        #ifdef __cplusplus
            #define LEXFLOATCLIENT_API extern "C"
        #else
            #define LEXFLOATCLIENT_API
        #endif
    #endif
	typedef const char* CSTRTYPE;
	typedef char* STRTYPE;
#endif
typedef void (LF_CC *CallbackType)(uint32_t);

/*
    FUNCTION: SetHostProductId()

    PURPOSE: Sets the product id of your application.

    PARAMETERS:
    * productId - the unique product id of your application as mentioned
      on the product page in the dashboard.

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID
*/
LEXFLOATCLIENT_API int LF_CC SetHostProductId(CSTRTYPE productId);

/*
    FUNCTION: SetHostUrl()

    PURPOSE: Sets the network address of the LexFloatServer.

    The url format should be: http://[ip or hostname]:[port]

    PARAMETERS:
    * hostUrl - url string having the correct format

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_HOST_URL
*/
LEXFLOATCLIENT_API int LF_CC SetHostUrl(CSTRTYPE hostUrl);

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
LEXFLOATCLIENT_API int LF_CC SetFloatingLicenseCallback(CallbackType callback);

/*
    FUNCTION: SetFloatingClientMetadata()

    PURPOSE: Sets the floating client metadata.

    The  metadata appears along with the license details of the license
    in LexFloatServer dashboard.

    PARAMETERS:
    * key - string of maximum length 256 characters with utf-8 encoding.
    * value - string of maximum length 256 characters with utf-8 encoding.

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_METADATA_KEY_LENGTH,
    LF_E_METADATA_VALUE_LENGTH, LF_E_ACTIVATION_METADATA_LIMIT
*/
LEXFLOATCLIENT_API int LF_CC SetFloatingClientMetadata(CSTRTYPE key, CSTRTYPE value);

/*
    FUNCTION: GetHostProductVersionName()

    PURPOSE: Gets the product version name.

    PARAMETERS:
    * name - pointer to a buffer that receives the value of the string
    * length - size of the buffer pointed to by the name parameter

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_PRODUCT_VERSION_NOT_LINKED, LF_E_BUFFER_SIZE
*/
LEXFLOATCLIENT_API int LF_CC GetHostProductVersionName(STRTYPE name, uint32_t length);

/*
    FUNCTION: GetHostProductVersionDisplayName()

    PURPOSE: Gets the product version display name.

    PARAMETERS:
    * displayName - pointer to a buffer that receives the value of the string
    * length - size of the buffer pointed to by the displayName parameter

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_PRODUCT_VERSION_NOT_LINKED, LF_E_BUFFER_SIZE
*/
LEXFLOATCLIENT_API int LF_CC GetHostProductVersionDisplayName(STRTYPE displayName, uint32_t length);

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
LEXFLOATCLIENT_API int LF_CC GetHostProductVersionFeatureFlag(CSTRTYPE name, uint32_t *enabled, STRTYPE data, uint32_t length);

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
LEXFLOATCLIENT_API int LF_CC GetHostLicenseMetadata(CSTRTYPE key, STRTYPE value, uint32_t length);

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
LEXFLOATCLIENT_API int LF_CC GetHostLicenseMeterAttribute(CSTRTYPE name, uint32_t *allowedUses, uint32_t *totalUses, uint32_t *grossUses);

/*
    FUNCTION: GetHostLicenseExpiryDate()

    PURPOSE: Gets the license expiry date timestamp of the LexFloatServer license.

    PARAMETERS:
    * expiryDate - pointer to the integer that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE
*/
LEXFLOATCLIENT_API int LF_CC GetHostLicenseExpiryDate(uint32_t *expiryDate);

/*
    FUNCTION: GetFloatingClientMeterAttributeUses()

    PURPOSE: Gets the meter attribute uses consumed by the floating client.

    PARAMETERS:
    * name - name of the meter attribute
    * allowedUses - pointer to the integer that receives the value

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_METER_ATTRIBUTE_NOT_FOUND
*/
LEXFLOATCLIENT_API int LF_CC GetFloatingClientMeterAttributeUses(CSTRTYPE name, uint32_t *uses);

/*
    FUNCTION: RequestFloatingLicense()

    PURPOSE: Sends the request to lease the license from the LexFloatServer.

    RETURN CODES: LF_OK, LF_FAIL, LF_E_PRODUCT_ID, LF_E_LICENSE_EXISTS, LF_E_HOST_URL,
    LF_E_CALLBACK, LF_E_LICENSE_LIMIT_REACHED, LF_E_INET, LF_E_TIME, LF_E_CLIENT, LF_E_IP, LF_E_SERVER,
    LF_E_SERVER_LICENSE_NOT_ACTIVATED, LF_E_SERVER_TIME_MODIFIED, LF_E_SERVER_LICENSE_SUSPENDED,
    LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER, LF_E_SERVER_LICENSE_EXPIRED
*/
LEXFLOATCLIENT_API int LF_CC RequestFloatingLicense();

/*
    FUNCTION: DropFloatingLicense()

    PURPOSE: Sends the request to the LexFloatServer to free the license.

    Call this function before you exit your application to prevent zombie licenses.

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE, LF_E_HOST_URL, LF_E_CALLBACK,
    LF_E_INET, LF_E_LICENSE_NOT_FOUND, LF_E_CLIENT, LF_E_IP, LF_E_SERVER,
    LF_E_SERVER_LICENSE_NOT_ACTIVATED, LF_E_SERVER_TIME_MODIFIED, LF_E_SERVER_LICENSE_SUSPENDED,
    LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER, LF_E_SERVER_LICENSE_EXPIRED
*/
LEXFLOATCLIENT_API int LF_CC DropFloatingLicense();

/*
    FUNCTION: HasFloatingLicense()

    PURPOSE: Checks whether any license has been leased or not. If yes,
    it retuns LF_OK.

    RETURN CODES: LF_OK, LF_E_PRODUCT_ID, LF_E_NO_LICENSE
*/
LEXFLOATCLIENT_API int LF_CC HasFloatingLicense();

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
LEXFLOATCLIENT_API int LF_CC IncrementFloatingClientMeterAttributeUses(CSTRTYPE name, uint32_t increment);

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
LEXFLOATCLIENT_API int LF_CC DecrementFloatingClientMeterAttributeUses(CSTRTYPE name, uint32_t decrement);

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
LEXFLOATCLIENT_API int LF_CC ResetFloatingClientMeterAttributeUses(CSTRTYPE name);
