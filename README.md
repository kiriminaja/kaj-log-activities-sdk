# KAJ Log Activities SDK
## Install
```
go get github.com/kiriminaja/kaj-log-activities-sdk
```
### Example
```
package main

import (
	"github.com/kiriminaja/kaj-log-activities-sdk/cmd/app"
	kaj_log_activities_sdk "github.com/kiriminaja/kaj-log-activities-sdk/sdk"
)

func main() {
	client := kaj_log_activities_sdk.NewClient(kaj_log_activities_sdk.Config{
		URL: "http://localhost:8081",
	})
    client.Packages.Search("$PACKAGE_ID")
}
```

## Command Line
### Show Help
```
# kaj-log -h

Usage:
   [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  event       Event 
  help        Help about any command
  package     package [command]
  payment     payment [command]
  withdrawal  withdrawal [command]

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.

```
### Packages Command
```
# kaj-log withdrawal search -s 11429683 -d WITHDRAWAL_2024082203595068904804
Source             WithdrawalID  Ref Number      
kaj-prd-core-srvc  11429683      OID-7951969212  
Event Tracking ID                  Event Name    
WITHDRAWAL_2024082203595068904804  disbursement 
```