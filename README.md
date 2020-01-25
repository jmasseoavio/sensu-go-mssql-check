[![Bonsai Asset Badge](https://img.shields.io/badge/CHANGEME-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/CHANGEME/CHANGEME)[![TravisCI Build Status](https://travis-ci.org/jmasseoavio/sensu-go-mssql-check.svg?branch=master)
](https://travis-ci.org/jmasseoavio/sensu-go-mssql-check)

# Sensu Go MSSQL Check

-  [Overview](#overview)

-  [Usage examples](#usage-examples)

-  [Configuration](#configuration)

-  [Asset registration](#asset-registration)

-  [Asset configuration](#asset-configuration)

-  [Resource(CHANGEME to: check,filter,mutator,handler) configuration](#resource-configuration)

-  [Functionality](#functionality)

-  [Installation from source and contributing](#installation-from-source-and-contributing)

  

## Overview

  

This is a Golang native check for MSSQL. Using [denisenkom/go-mssqldb][3].

This is pretty simple, it was to help diagnose a problem a customer was seeing with API calls that relied on an MSSQL database.

It just returns 0/1/2 and outputs the times in InfluxDB metrics format.

  

## Configuration

  

All configuration is via arguments or environment variables. -1 means those checks are disabled. If no checks are enabled, this is just a metrics check. All times are in nanos. So you probably want to specify warntime/crittime as 1000000000 for 1 second.

  

```

The Sensu Go mssql check

  

Usage:

sensu-go-mssql-check [flags]

  

Flags:

-C, --connstring string MSSQL Connection String (MSSQL_CHECK_CONNSTRING)

-c, --crittime int MSSQL Query Critical Time (MSSQL_CHECK_CRITTIME) (default -1)

-r, --desiredrows int MSSQL Query Desired Rows (MSSQL_CHECK_DESIREDROWS) (default -1)

-h, --help help for sensu-go-mssql-check

-q, --querystring string MSSQL Query String (MSSQL_CHECK_QUERYSTRING) (default "select 1")

-w, --warntime int MSSQL Query Warning Time (MSSQL_CHECK_WARNTIME) (default -1)

```

  

### Asset Registration

  

Assets are the best way to make use of this plugin. If you're not using an asset, please consider doing so! If you're using sensuctl 5.13 or later, you can use the following command to add the asset:

  

`sensuctl asset add jmasseoavio/sensu-go-mssql-check:VERSION`

  

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index](https://bonsai.sensu.io/assets/jmasseoavio/sensu-go-mssql-check).

  

### Asset configuration

  

TODO: Provide an example asset manifest

  

```yml

---

type: Asset

api_version: core/v2

metadata:

name: CHANGEME

spec:

url: https://CHANGEME

sha512: CHANGEME

```

  

### Resource(check) configuration

  

Example Sensu Go definition:

  

```yml

---

api_version: core/v2

type: check

metadata:

namespace: default

name: sensu-go-mssql-check

spec:

"...": "..."

  

```

  

### Functionality

  

Should run fine. :)

  

## Installation from source and contributing

  

The preferred way of installing and deploying this plugin is to use it as an [asset]. If you would like to compile and install from source or contribute to it, download the latest version from [releases][1] or create an executable from this source.

From the local path of the sensu-go-mssql-check repository:

```

go build -o /usr/local/bin/sensu-go-mssql-check main.go

```

For more information about contributing to this plugin, see https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md

[1]: https://github.com/jmasseoavio/sensu-go-mssql-check/releases

[2]: #asset-registration

[3]: https://github.com/denisenkom/go-mssqldb