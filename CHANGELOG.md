# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.9.1] - 2025-02-13
### Added
- Restored deprecated helpers removed in 1.9.0 (`sliceutils.FindIndexOf`, `Includes`, `Merge`, `Insert`, `Copy`; `dateutils.BeforeOrEqual`, `AfterOrEqual`) with guidance to stdlib replacements.
- Reinstated documentation pages and mkdocs navigation for the deprecated helpers; added maputils Copy doc note for `maps.Clone`.

### Fixed
- Refined `stringutils.Stringify` for typed nil pointers and clearer flow; resolved lint issues.
- Corrected typo in `dateutils.Ceil` comment (“Subtract”).

### Notes
- Release tagged as `v1.9.1` and published.

## [1.9.0] - 2025-02-XX
### Removed
- Removed several helpers now provided by the Go standard library (`sliceutils.FindIndexOf`, `Includes`, `Merge`, `Insert`, `Copy`; `dateutils.BeforeOrEqual`, `AfterOrEqual`).

### Added
- New utilities across sliceutils, maputils, dateutils, ptrutils, mathutils, excutils, urlutils, structutils (see docs for details).

[1.9.1]: https://github.com/Goldziher/go-utils/releases/tag/v1.9.1
[1.9.0]: https://github.com/Goldziher/go-utils/releases/tag/v1.9.0
