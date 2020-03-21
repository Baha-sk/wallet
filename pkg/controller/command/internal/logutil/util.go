/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package logutil

import (
	"github.com/sirupsen/logrus"
)

// LogError is a utility function to log error messages.
func LogError(logger *logrus.Logger, command, action, errMsg string, data ...string) {
	logger.Errorf("command=[%s] action=[%s] %s errMsg=[%s]", command, action, data, errMsg)
}

// LogDebug is a utility function to log debug messages.
func LogDebug(logger *logrus.Logger, command, action, msg string, data ...string) {
	logger.Debugf("command=[%s] action=[%s] %s msg=[%s]", command, action, data, msg)
}

// LogInfo is a utility function to log info messages.
func LogInfo(logger *logrus.Logger, command, action, msg string, data ...string) {
	logger.Infof("command=[%s] action=[%s] %s msg=[%s]", command, action, data, msg)
}
