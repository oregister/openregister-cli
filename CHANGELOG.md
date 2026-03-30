# Changelog

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
