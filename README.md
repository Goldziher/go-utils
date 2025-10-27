# Go Utils

<div align="center">

[![Go Report Card](https://goreportcard.com/badge/github.com/Goldziher/go-utils)](https://goreportcard.com/report/github.com/Goldziher/go-utils)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Goldziher_go-utils&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=Goldziher_go-utils)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=Goldziher_go-utils&metric=coverage)](https://sonarcloud.io/summary/new_code?id=Goldziher_go-utils)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=Goldziher_go-utils&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=Goldziher_go-utils)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=Goldziher_go-utils&metric=bugs)](https://sonarcloud.io/summary/new_code?id=Goldziher_go-utils)

[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=Goldziher_go-utils&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=Goldziher_go-utils)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=Goldziher_go-utils&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=Goldziher_go-utils)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=Goldziher_go-utils&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=Goldziher_go-utils)

</div>

Go Utils is a comprehensive utility library for **Go 1.21+** that provides functional programming patterns and utilities inspired by JavaScript and Python. Built with generics for type safety, it's designed to be **complementary** with modern standard library packages like `slices`, `maps`, and `cmp`.

## Design Philosophy

- **Go 1.21+ required**: Requires Go 1.21 or later for generics and modern stdlib integration
- **Complementary to stdlib**: Works alongside modern Go stdlib packages (`slices`, `maps`, `cmp`)
  - Use `slices.Index`, `slices.Contains`, `slices.Clone`, `slices.Concat` for basic operations
  - Use this library for functional patterns (Map, Filter, Reduce), LINQ-style operations (GroupBy, Partition), and high-value utilities not in stdlib
- **Zero-value focused**: Functions accept zero values and handle nil slices/maps gracefully
- **Immutability by default**: Functions don't mutate inputs unless explicitly named (e.g., Remove creates new slices)
- **100% test coverage**: Every function has comprehensive test coverage

## Packages

- **sliceutils**: Functional operations (Map, Filter, Reduce, Find) and LINQ-style utilities (GroupBy, Partition, DistinctBy)
- **maputils**: Map transformations (Keys, Values, Filter, Merge, Invert)
- **structutils**: Reflection-based struct utilities (ToMap, ForEach, FieldNames with tag support)
- **stringutils**: String manipulation (Stringify with options, case conversion, padding, truncation)
- **dateutils**: Time/date helpers (overlap detection, business days, month boundaries, age calculation)
- **urlutils**: URL parsing and query string builders from maps/structs
- **mathutils**: Generic math operations (Clamp, InRange, Gcd, Lcm, IsPrime)
- **ptrutils**: Pointer utilities (ToPtr, Deref with defaults)
- **excutils**: Exception-style error handling (Panic, Try, Must)

Read more in the [documentation](https://goldziher.github.io/go-utils/).

## Installation

```shell
go get -u github.com/Goldziher/go-utils
```

## Contribution

This library is open to contributions. Please consult the [Contribution Guide](CONTRIBUTING.md).
