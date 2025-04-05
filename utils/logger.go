// Logging utility
package utils

import (
    "log"
    "os"
)

var (
    LogFile, _ = os.OpenFile("chat.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    Logger     = log.New(LogFile, "", log.LstdFlags)
)

