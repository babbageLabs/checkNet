# Introduction

This is a simple CLI application to collect statistics about the device internet uptime.
The information collected is stored in a sql lite DB and the data can be consumed for presentation

## Prerequisites

- windows PC with [powershell](https://learn.microsoft.com/en-us/powershell/)
- Appropriate permissions to write to disk (or network) and to check network status
- [SQLite](https://www.sqlite.org/])

## How it works

The application uses the interface already exposed by the OS. The text responses from 
the commands is collected and stored in the configured [SQLite](https://www.sqlite.org/]) db.
For example in windows we would run

```ps
netsh wlan show interfaces
```

## Analytics

The toolkit will also run a simple web server to show a simplified dashboard to consume the data

NB: This feature is WIP

## Available commands

```ps
checkNet # see the current network status
checkNet m # start monitoring
checkNet sm # stop monitoring
checkNet im # check monitoring status
```