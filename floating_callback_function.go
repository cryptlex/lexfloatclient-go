// Copyright 2021 Cryptlex, LLC. All rights reserved.

package lexfloatclient

/*
#include <stdio.h>

// The gateway functions
void floatingLicenseCallbackCgoGateway(int status)
{
	void floatingLicenseCallbackWrapper(int);
	floatingLicenseCallbackWrapper(status);
}
*/
import "C"
