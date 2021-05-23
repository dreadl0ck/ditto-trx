# DittoTRX

[![Go Report Card](https://goreportcard.com/badge/github.com/dreadl0ck/maltego)](https://goreportcard.com/report/github.com/dreadl0ck/maltego)
[![License](https://img.shields.io/badge/license-GPL-green)](https://raw.githubusercontent.com/dreadl0ck/ditto-trx/master/LICENSE)

A [Maltego](https://www.maltego.com) transform server that implements a transform set to handle queries to the [Ditto](https://github.com/evilsocket/ditto) IDN homograph attacks and detection tool,
as well as local transformations for working with the resulting entities.
You can read more about it in my accompanying blogpost [](https://dreadl0ck.net/posts/ditto-trx).

## Remote Transforms

- SimilarDomains
- RegisteredDomains
- LiveDomains
- AvailableDomains
- LiveDomainsTLD

## Local Transforms

- LookupAddr
- ToDomainNames
- ToRegistrarNames
- ToNameServers
- ToCreationDate

## Compile from source

    go build

## Docker Containers

    docker pull dreadl0ck/ditto-trx

## Usage Examples

Check the examples folder and unit tests!

## Maltego Configuration

Import the _dittotrx.mtz_ file into maltego, to install the transforms and entities. 

```
$ tree dittotrx
dittotrx
├── Entities
│   └── dittotrx.IDNDomain.entity
├── EntityCategories
│   └── dittotrx.category
├── Icons
│   └── dittotrx
│       ├── domain_black.svg
│       ├── domain_black.xml
│       ├── domain_black24.svg
│       ├── domain_black32.svg
│       ├── domain_black48.svg
│       └── domain_black96.svg
├── Servers
│   └── Local.tas
├── TransformRepositories
│   └── Local
│       ├── dittotrx.LookupAddr.transform
│       ├── dittotrx.LookupAddr.transformsettings
│       ├── dittotrx.ToCreationDate.transform
│       ├── dittotrx.ToCreationDate.transformsettings
│       ├── dittotrx.ToDomainNames.transform
│       ├── dittotrx.ToDomainNames.transformsettings
│       ├── dittotrx.ToIPAddresses.transform
│       ├── dittotrx.ToIPAddresses.transformsettings
│       ├── dittotrx.ToNameServers.transform
│       ├── dittotrx.ToNameServers.transformsettings
│       ├── dittotrx.ToRegistrarNames.transform
│       ├── dittotrx.ToRegistrarNames.transformsettings
│       ├── dittotrx.VisitDomain.transform
│       └── dittotrx.VisitDomain.transformsettings
├── TransformSets
│   └── DittoTRX.set
└── version.properties
```

## Code Stats

    $ cloc *.go
           2 text files.
           2 unique files.                              
           0 files ignored.
    
    github.com/AlDanial/cloc v 1.84  T=0.01 s (227.0 files/s, 23263.9 lines/s)
    -------------------------------------------------------------------------------
    Language                     files          blank        comment           code
    -------------------------------------------------------------------------------
    Go                               2             39             39            127
    -------------------------------------------------------------------------------
    SUM:                             2             39             39            127
    -------------------------------------------------------------------------------

## License

GPLv3
