# Changelog

## 0.2.2 (2026-04-13)

Full Changelog: [v0.2.1...v0.2.2](https://github.com/oregister/openregister-cli/compare/v0.2.1...v0.2.2)

### Chores

* **internal:** codegen related update ([20c3613](https://github.com/oregister/openregister-cli/commit/20c3613fc46eae020c7670c0883c3877e40cae8a))

## 0.2.1 (2026-04-11)

Full Changelog: [v0.2.0...v0.2.1](https://github.com/oregister/openregister-cli/compare/v0.2.0...v0.2.1)

### Bug Fixes

* **cli:** fix incompatible Go types for flag generated as array of maps ([052f8b5](https://github.com/oregister/openregister-cli/commit/052f8b5214227c5f1937d01864106b3f1e21bbf5))
* fix for failing to drop invalid module replace in link script ([dd8a5f0](https://github.com/oregister/openregister-cli/commit/dd8a5f0e05c4815b82d9a8fc2fe3698f7b7f6890))


### Chores

* **cli:** additional test cases for `ShowJSONIterator` ([7bcaf2a](https://github.com/oregister/openregister-cli/commit/7bcaf2a397d1ece712ec50b906343d554a6d7d6d))
* **cli:** let `--format raw` be used in conjunction with `--transform` ([bd09e98](https://github.com/oregister/openregister-cli/commit/bd09e98ba5abbb045cdf6749b53d6ee17e7f0308))

## 0.2.0 (2026-04-08)

Full Changelog: [v0.1.1...v0.2.0](https://github.com/oregister/openregister-cli/compare/v0.1.1...v0.2.0)

### Features

* allow `-` as value representing stdin to binary-only file parameters in CLIs ([3e1beec](https://github.com/oregister/openregister-cli/commit/3e1beec62d593dd1442851fe647f10f81701da58))
* better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url` ([dc2c6a4](https://github.com/oregister/openregister-cli/commit/dc2c6a465fd7eb8c4e8caf06110a119672605f47))
* binary-only parameters become CLI flags that take filenames only ([b44a04c](https://github.com/oregister/openregister-cli/commit/b44a04c0e4b6d0a640d1e2ed7bc302d10ac3c38f))


### Bug Fixes

* fall back to main branch if linking fails in CI ([d828452](https://github.com/oregister/openregister-cli/commit/d828452ffa84cf983892261ab2483fb791fd5926))
* fix quoting typo ([18b4fac](https://github.com/oregister/openregister-cli/commit/18b4fac5c13bee6c4a47d963478f3801e49cf72c))


### Chores

* mark all CLI-related tests in Go with `t.Parallel()` ([8ad5740](https://github.com/oregister/openregister-cli/commit/8ad57409f88fec5653fb331f2abe64d1bc3b5201))
* modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary ([d0bb945](https://github.com/oregister/openregister-cli/commit/d0bb94526adca83aa5a8e5ebaef5799df865a362))
* switch some CLI Go tests from `os.Chdir` to `t.Chdir` ([ee444ba](https://github.com/oregister/openregister-cli/commit/ee444bad45ab0e7f1240ce5b2905f0fd11cbb340))

## 0.1.1 (2026-04-02)

Full Changelog: [v0.1.0...v0.1.1](https://github.com/oregister/openregister-cli/compare/v0.1.0...v0.1.1)

### Bug Fixes

* handle empty data set using `--format explore` ([309a4c3](https://github.com/oregister/openregister-cli/commit/309a4c37d63ba09846ccb972c95ca71258ef59fa))
* use `RawJSON` when iterating items with `--format explore` in the CLI ([a02e2ae](https://github.com/oregister/openregister-cli/commit/a02e2ae3a5df3da490006ccb6d3e0e48aab103a9))

## 0.1.0 (2026-03-30)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/oregister/openregister-cli/compare/v0.0.1...v0.1.0)

### Features

* add default description for enum CLI flags without an explicit description ([d739a6a](https://github.com/oregister/openregister-cli/commit/d739a6a725f92b32d6759a6cc784b4ffa000397f))
* **api:** monitoring endpoints ([9fa2061](https://github.com/oregister/openregister-cli/commit/9fa2061ef3dd059e0a5233beca9362b2ccf18bfb))
* set CLI flag constant values automatically where `x-stainless-const` is set ([caa77db](https://github.com/oregister/openregister-cli/commit/caa77dba71c08c2b58f1de8280208c8823b9d04d))


### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([b5bf15b](https://github.com/oregister/openregister-cli/commit/b5bf15b033806d50c130e217a80dd7623278197e))
* better support passing client args in any position ([4bf840c](https://github.com/oregister/openregister-cli/commit/4bf840c84255521ae7cd9ddd9d8d1143b4c0b1ad))
* cli no longer hangs when stdin is attached to a pipe with empty input ([60ff5ae](https://github.com/oregister/openregister-cli/commit/60ff5ae85c829f0dff108eefdea2d354d403479d))
* fix for off-by-one error in pagination logic ([5df6d25](https://github.com/oregister/openregister-cli/commit/5df6d2509e2f74497a59e755aa222e03ec76f8df))
* improve linking behavior when developing on a branch not in the Go SDK ([7ee5b5a](https://github.com/oregister/openregister-cli/commit/7ee5b5a9c151a7f2e4dc34e41e4e60792adb35da))
* improved workflow for developing on branches ([97a419f](https://github.com/oregister/openregister-cli/commit/97a419f625391da0a04a5309e3aee3e6599acc36))
* no longer require an API key when building on production repos ([8d1ba9a](https://github.com/oregister/openregister-cli/commit/8d1ba9a56195f0cb30281c1bf8c2bd462e241f24))


### Chores

* **ci:** skip lint on metadata-only changes ([885e009](https://github.com/oregister/openregister-cli/commit/885e0095adbc85a38abb8b69b8ad600f0da121a9))
* configure new SDK language ([a53a041](https://github.com/oregister/openregister-cli/commit/a53a041fdbf9179295fdd7bc852957836ea8bd5a))
* **internal:** tweak CI branches ([2574e45](https://github.com/oregister/openregister-cli/commit/2574e459cbe8085e5b9ac690881ef735821b5a44))
* **internal:** update gitignore ([520c179](https://github.com/oregister/openregister-cli/commit/520c179289ea988eb47da11fe2e425d5300e8003))
* omit full usage information when missing required CLI parameters ([3637626](https://github.com/oregister/openregister-cli/commit/3637626da1c59adbc287e2df5fede611c3caa84f))
* update SDK settings ([bda436e](https://github.com/oregister/openregister-cli/commit/bda436e0c089d4c2d11a1ff04f7c390caa13dd42))
* update SDK settings ([59dd686](https://github.com/oregister/openregister-cli/commit/59dd6862475f792d05b906f9d2aedf1509afd0aa))
