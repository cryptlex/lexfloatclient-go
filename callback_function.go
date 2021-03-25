// Copyright 2021 Cryptlex, LLC. All rights reserved.

package lexfloatclient

/*
#include <stdio.h>

// The gateway functions
void licenseCallbackCgoGateway(int status)
{
	void licenseCallbackWrapper(int);
	licenseCallbackWrapper(status);
}
*/
import "C"
