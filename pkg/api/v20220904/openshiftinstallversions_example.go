package v20220904

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

// ExampleInstallOpenShiftVersions returns an example
// InstallOpenShiftVersions object i.e []string that the RP would return to an end-user.
func ExampleInstallOpenShiftVersionsResponse() interface{} {
	return &InstallOpenShiftVersions{"4.10.20"}
}
