// Copyright 2021 Cryptlex LLC. All rights reserved.

package lexfloatclient

// int enumeration from lexfloatclient/int.h int =4
const (
	// Success code.
    LF_OK int = 0

    // Failure code.
    LF_FAIL int = 1

    // The product id is incorrect.
    LF_E_PRODUCT_ID int = 40

    // Invalid or missing callback function.
    LF_E_CALLBACK int = 41

    // Missing or invalid server url.
    LF_E_HOST_URL int = 42

    // Ensure system date and time settings are correct.
    LF_E_TIME int = 43

    // Failed to connect to the server due to network error.
    LF_E_INET int = 44

    // License has not been leased yet.
    LF_E_NO_LICENSE int = 45

    // License has already been leased.
    LF_E_LICENSE_EXISTS int = 46

    // License does not exist on server or has already expired. This
	// happens when the request to refresh the license is delayed.
    LF_E_LICENSE_NOT_FOUND int = 47

    // License lease has expired due to network error. This
	// happens when the request to refresh the license fails due to
	// network error.
    LF_E_LICENSE_EXPIRED_INET int = 48

    // The server has reached it's allowed limit of floating licenses.
    LF_E_LICENSE_LIMIT_REACHED int = 49

    // The buffer size was smaller than required.
    LF_E_BUFFER_SIZE int = 50

    // The metadata key does not exist.
    LF_E_METADATA_KEY_NOT_FOUND int = 51

    // Metadata key length is more than 256 characters.
    LF_E_METADATA_KEY_LENGTH int = 52

    // Metadata value length is more than 256 characters.
    LF_E_METADATA_VALUE_LENGTH int = 53

    // The floating client has reached it's metadata fields limit.
    LF_E_FLOATING_CLIENT_METADATA_LIMIT int = 54

    // The meter attribute does not exist.
    LF_E_METER_ATTRIBUTE_NOT_FOUND int = 55

    // The meter attribute has reached it's usage limit.
    LF_E_METER_ATTRIBUTE_USES_LIMIT_REACHED int = 56

    // IP address is not allowed.
    LF_E_IP int = 60

    // Client error.
    LF_E_CLIENT int = 70

    // Server error.
    LF_E_SERVER int = 71

    // System time on server has been tampered with. Ensure
	// your date and time settings are correct on the server machine.
    LF_E_SERVER_TIME_MODIFIED int = 72

    // The server has not been activated using a license key.
    LF_E_SERVER_LICENSE_NOT_ACTIVATED int = 73

    // The server license has expired.
    LF_E_SERVER_LICENSE_EXPIRED int = 74

    // The server license has been suspended.
    LF_E_SERVER_LICENSE_SUSPENDED int = 75

    // The grace period for server license is over.
    LF_E_SERVER_LICENSE_GRACE_PERIOD_OVER int = 76
)
