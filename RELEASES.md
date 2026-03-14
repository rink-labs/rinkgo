# Release Notes

## Pending (v1.14.2)

## [v1.14.1](https://github.com/ava-labs/avalanchego/releases/tag/v1.14.1)

This version is backwards compatible to [v1.14.0](https://github.com/ava-labs/avalanchego/releases/tag/v1.14.0). It is optional, but encouraged.

The plugin version is unchanged at `44` and is compatible with version `v1.14.0`.

### Config

- Added:
  - `--system-tracker-disk-required-available-space-percentage`
  - `--system-tracker-disk-warning-available-space-percentage`
- Deprecated:
  - `--system-tracker-disk-required-available-space`
  - `--system-tracker-disk-warning-threshold-available-space`

### EVM

- Removed `avax.version` API
- Removed `customethclient` package in favor of `ethclient` package and temporary type registrations (`WithTempRegisteredLibEVMExtras`)
  - Removed blockHook extension in `ethclient` package.

### What's Changed

The changelog is omitted, as the Coreth and Subnet-EVM repositories were grafted into the repository.

**Full Changelog**: https://github.com/ava-labs/avalanchego/compare/v1.14.0...v1.14.1

## [v1.14.0](https://github.com/ava-labs/avalanchego/releases/tag/v1.14.0)

This release schedules the activation of the following Avalanche Community Proposals (ACPs):
- [ACP-181](https://github.com/avalanche-foundation/ACPs/blob/main/ACPs/181-p-chain-epoched-views/README.md) P-Chain Epoched Views
- [ACP-204](https://github.com/avalanche-foundation/ACPs/blob/main/ACPs/204-precompile-secp256r1/README.md) Precompile for secp256r1 Curve Support
- [ACP-226](https://github.com/avalanche-foundation/ACPs/blob/main/ACPs/226-dynamic-minimum-block-times/README.md) Dynamic Minimum Block Times

The ACPs in this upgrade go into effect at 11 AM ET (4 PM UTC) on Wednesday, November 19th, 2025 on Mainnet.

**All Mainnet nodes must upgrade before 11 AM ET, November 19th 2025.**

The plugin version is updated to `44`; all plugins must update to be compatible.

### LibEVM Hook Registration

This release includes changes for how EVM modifications are registered with libevm. For consumers of avalanchego as a library, manual registration of libevm callbacks is now required.

### APIs

- Added support for specifying multiple `Avalanche-Api-Route` headers for more complex routing.
- Added proposervm gRPC, connectrpc, and jsonrpc APIs for `GetProposedHeight` and `GetCurrentEpoch`.
  - The gRPC and connectrpc APIs are routed by adding a second `Avalanche-Api-Route` header with the value `proposervm`.
  - The jsonrpc APIs are added to all chains with the base endpoint `/proposervm`.
- Added platformvm `platform.GetAllValidatorsAt` API.

### Configs

- Changed default `proposerMinBlockDelay` for L1s (other than Primary Network) to 0.

### Fixes

- Improved bootstrapping ETA predictions.

**Full Changelog**: https://github.com/ava-labs/avalanchego/compare/v1.13.5...v1.14.0

## [v1.13.5](https://github.com/ava-labs/avalanchego/releases/tag/v1.13.5)

This version is backwards compatible to [v1.13.0](https://github.com/ava-labs/avalanchego/releases/tag/v1.13.0). It is optional, but encouraged.

The plugin version is unchanged at `43` and is compatible with version `v1.13.4`.

### Configs

- Replaced `pull-gossip-throttling-limit` with `pull-gossip-requests-per-validator` in the X-Chain and P-Chain configs

### Fixes

- Fixed duplicate C-Chain eth gossip registration
- Fixed various C-Chain atomic mempool edge cases

### What's Changed

- Separate re-execution job params for PR from schedule by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4151
- chore: fix a typo in gossip,go by @Galoretka in https://github.com/ava-labs/avalanchego/pull/4154
- Add @joshua-kim as CODEOWNER to testing-related packages by @JuanLeon2 in https://github.com/ava-labs/avalanchego/pull/4118
- refactor(load): remove context from test interface by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/4157
- Move C-Chain benchmark to custom action and add ARC + GH runner triggers by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4165
- Use EmptyVoteMetadata in Simplex Proto Messages by @samliok in https://github.com/ava-labs/avalanchego/pull/4174
- feat(load): add token test by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/4171
- Fix typo in comment - PChainHeight context by @yacovm in https://github.com/ava-labs/avalanchego/pull/4176
- chore: fix function name by @yinwenyu6 in https://github.com/ava-labs/avalanchego/pull/4178
- chore: fix typo by @kks-code in https://github.com/ava-labs/avalanchego/pull/4179
- Migrate predicate package from evm repos by @JonathanOppenheimer in https://github.com/ava-labs/avalanchego/pull/4147
- feat: add eviction callback in LRU cache by @DracoLi in https://github.com/ava-labs/avalanchego/pull/4088
- `config/config.md:` Added Env Variable representation of flags + improved UI design by @navillanueva in https://github.com/ava-labs/avalanchego/pull/4110
- Storage Component For Simplex by @samliok in https://github.com/ava-labs/avalanchego/pull/4122
- Block Database by @DracoLi in https://github.com/ava-labs/avalanchego/pull/4027
- Change cache path to tmp included in gitignore by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4183
- Add optional step to archive post-reexecution state to S3 by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4172
- Rename height field to numBlocks by @samliok in https://github.com/ava-labs/avalanchego/pull/4184
- refactor: replace []byte(fmt.Sprintf) with fmt.Appendf by @queryfast in https://github.com/ava-labs/avalanchego/pull/4161
- Make Draco the codeowner of the blockdb by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/4187
- Add redundant import alias linting by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/4191
- Add config option for AWS S3 read only credential duration by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4192
- fix: blockdb file eviction race issue by @DracoLi in https://github.com/ava-labs/avalanchego/pull/4186
- Count throttled requests as hits by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/4199
- Rename Engine Types by @samliok in https://github.com/ava-labs/avalanchego/pull/4193
- Add s5cmd progress bar by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/4204
- Update block + validator + pgo checkpoints to 2025-08-23 by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/4205
- Remove buf lint action by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/4189
- Add ability to create zstd compressor with compression level by @DracoLi in https://github.com/ava-labs/avalanchego/pull/4203
- Dynamically update mempool gossip request rate limit by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/4162
- Add support for passing config and predefined configs to VM re-execution tests by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4180

### New Contributors

- @Galoretka made their first contribution in https://github.com/ava-labs/avalanchego/pull/4154
- @JuanLeon2 made their first contribution in https://github.com/ava-labs/avalanchego/pull/4118
- @yinwenyu6 made their first contribution in https://github.com/ava-labs/avalanchego/pull/4178
- @kks-code made their first contribution in https://github.com/ava-labs/avalanchego/pull/4179
- @navillanueva made their first contribution in https://github.com/ava-labs/avalanchego/pull/4110
- @queryfast made their first contribution in https://github.com/ava-labs/avalanchego/pull/4161

**Full Changelog**: https://github.com/ava-labs/avalanchego/compare/v1.13.4...v1.13.5

## [v1.13.4](https://github.com/ava-labs/avalanchego/releases/tag/v1.13.4)

This version is backwards compatible to [v1.13.0](https://github.com/ava-labs/avalanchego/releases/tag/v1.13.0). It is optional, but encouraged.

The plugin version is updated to `43` all plugins must update to be compatible.

### Configs

- Added `staking-rpc-signer-endpoint` flag.

### What's Changed

- Bump golang.org/x/oauth2 from 0.21.0 to 0.27.0 by @dependabot[bot] in https://github.com/ava-labs/avalanchego/pull/4100
- [tmpnet] Enable externally accessible URIs for kube-hosted nodes by @maru-ava in https://github.com/ava-labs/avalanchego/pull/4016
- [antithesis] Enable reuse of banff e2e test for antithesis testing by @marun in https://github.com/ava-labs/avalanchego/pull/3554
- feat(load2): add contract tests by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/4071
- Add test to re-execute specified range of mainnet C-Chain blocks by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4019
- [tmpnet] Enable installation of chaos mesh to local kind cluster by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3674
- Update release notes and bump version to `v1.13.4` by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/4113
- Update reexecute c-chain range cronjob to run daily by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4115
- Remove Stale References of the toEngine Channel by @samliok in https://github.com/ava-labs/avalanchego/pull/4101
- Add step to push benchmark results to gh-pages by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4103
- ci: remove load 1.0 by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/4106
- Simplex QuorumCertificate and BLS aggregator by @samliok in https://github.com/ava-labs/avalanchego/pull/4091
- Update codeowners of reexecution changes by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4116
- refactor: use maps.Copy for cleaner map handling by @jishudashu in https://github.com/ava-labs/avalanchego/pull/4119
- refactor: remove load 1.0  by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/4112
- Enable Cubist Signer integration by @geoff-vball in https://github.com/ava-labs/avalanchego/pull/3965
- Split action benchmark comparison and push to gh-pages by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4130
- Remove external-data-json-path from benchmark push step by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4134
- With golangci-lint v2.2.2 using http.NewRequest is discouraged by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/4136
- Add runner input to run c-chain reexecution benchmark on arbitrary target by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4121
- Remove flaky dial throttler tests by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4139
- Remove gitignore line that ignores the `database/dbtest` package by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/4142
- chore: Update header year to 2025 by @JonathanOppenheimer in https://github.com/ava-labs/avalanchego/pull/4140
- feat(load): add trie stress test by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/4137
- Parameterize values in transfer tests by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/4144
- uplift: Add combined metrics package from evm repositories by @JonathanOppenheimer in https://github.com/ava-labs/avalanchego/pull/4135
- docs: load by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/4132
- chore: fix minor typo in comment by @lechpzn in https://github.com/ava-labs/avalanchego/pull/4150
- fix metrics tests by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/4146
- Update coreth to v0.15.3-rc.5 by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/4153

### New Contributors

- @jishudashu made their first contribution in https://github.com/ava-labs/avalanchego/pull/4119
- @lechpzn made their first contribution in https://github.com/ava-labs/avalanchego/pull/4150

**Full Changelog**: https://github.com/ava-labs/avalanchego/compare/v1.13.3...v1.13.4

## [v1.13.3](https://github.com/ava-labs/avalanchego/releases/tag/v1.13.3)

This version is backwards compatible to [v1.13.0](https://github.com/ava-labs/avalanchego/releases/tag/v1.13.0). It is optional, but encouraged.

The plugin version is updated to `42` all plugins must update to be compatible.

**This release removes the support for running on Windows. Any users running on Windows should utilize Windows Subsystem for Linux.**

### APIs

- Converted HTTP2 support to route by the `Avalanche-Api-Route` header.
- Added `avalanche_snowman_blks_built_time_skew` metric.
- Added p2p `bloomfilter_hit_rate` metrics for the P-chain, X-chain, C-chain atomic, and C-chain mempools.

### Fixes

- Fixed HTTP2 update request handling.
- Fixed plugin handling of special errors during HTTP request handling.
- Fixed plugin forwarding of HTTP headers.
- Fixed C-chain shared memory UTXO parsing by the Primary Network Wallet.

### What's Changed

- chore: move windows from tier 3 to unsupported by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4031
- [tmpnet] Run kube load test under a service account to validate RBAC by @maru-ava in https://github.com/ava-labs/avalanchego/pull/4030
- Add block timeskew metric by @yacovm in https://github.com/ava-labs/avalanchego/pull/4034
- refactor: use maps.Copy for cleaner map handling by @gopherorg in https://github.com/ava-labs/avalanchego/pull/4042
- Fix bug in rpcchainvm where `io.EOF` is not propagated correctly by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/4035
- chore: fix some minor issues in the comments by @alongdate in https://github.com/ava-labs/avalanchego/pull/4045
- Update golangci-lint to v2  by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/4037
- Simplex VerifiedBlock and BlockDeserializer by @samliok in https://github.com/ava-labs/avalanchego/pull/4009
- Load Framework 2 by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/3996
- Rename test_util -> util_test  by @samliok in https://github.com/ava-labs/avalanchego/pull/4054
- Replace toEngine channel with a dedicated subscription API by @yacovm in https://github.com/ava-labs/avalanchego/pull/3999
- Ensure that headers added by the server are propagated by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/4060
- Update coreth to v0.15.3-rc.0 by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/4062
- Add default validator set function to snowtest.context by @JonathanOppenheimer in https://github.com/ava-labs/avalanchego/pull/4065
- Fix flake in TestNotifierReSubscribeAtPrefChange by @yacovm in https://github.com/ava-labs/avalanchego/pull/4067
- Support header-based routing by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/4036
- Simplify json-rpc client handling by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/4079
- Update coreth to v0.15.3-rc.1 by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/4080
- feat: add default flag options by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/4044
- chore: fix some function names in comment by @shangchenglumetro in https://github.com/ava-labs/avalanchego/pull/4053
- chore: fix some inconsistent comments by @jingchanglu in https://github.com/ava-labs/avalanchego/pull/4068
- refactor: use the built-in max/min to simplify the code by @socialsister in https://github.com/ava-labs/avalanchego/pull/4087
- Comm Component for Simplex by @samliok in https://github.com/ava-labs/avalanchego/pull/3998
- ci(load-tests): add GitHub Actions workflow for load testing on k8s by @Elvis339 in https://github.com/ava-labs/avalanchego/pull/4051
- Update run monitored tmpnet cmd to make dashboard param configurable by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4089
- Add WritePrometheusServiceDiscoveryConfigFile helper by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/4072
- [nix] Update to 25.05 by @maru-ava in https://github.com/ava-labs/avalanchego/pull/4050
- chore: replace polling with subscription-based acceptance by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/4059
- Add Verification + Tracking Logic To Simplex Blocks by @samliok in https://github.com/ava-labs/avalanchego/pull/4052
- Fix wrong pkg import for Primary Wallet by @0xJohnnyGault in https://github.com/ava-labs/avalanchego/pull/4096
- [tmpnet] Source tmpnet defaults from a configmap by @maru-ava in https://github.com/ava-labs/avalanchego/pull/4057
- Throttle platformVM block building until normal ops by @yacovm in https://github.com/ava-labs/avalanchego/pull/4097
- fix: add throttle to error retries by @alarso16 in https://github.com/ava-labs/avalanchego/pull/4098
- Measure bloom filter AppGossip hit rate by @yacovm in https://github.com/ava-labs/avalanchego/pull/4085

### New Contributors

- @gopherorg made their first contribution in https://github.com/ava-labs/avalanchego/pull/4042
- @alongdate made their first contribution in https://github.com/ava-labs/avalanchego/pull/4045
- @JonathanOppenheimer made their first contribution in https://github.com/ava-labs/avalanchego/pull/4065
- @shangchenglumetro made their first contribution in https://github.com/ava-labs/avalanchego/pull/4053
- @jingchanglu made their first contribution in https://github.com/ava-labs/avalanchego/pull/4068
- @0xJohnnyGault made their first contribution in https://github.com/ava-labs/avalanchego/pull/4096

**Full Changelog**: https://github.com/ava-labs/avalanchego/compare/v1.13.2...v1.13.3

## [v1.13.2](https://github.com/ava-labs/avalanchego/releases/tag/v1.13.2)

This version is backwards compatible to [v1.13.0](https://github.com/ava-labs/avalanchego/releases/tag/v1.13.0). It is optional, but encouraged.

The plugin version is updated to `41` all plugins must update to be compatible.

### APIs

- Added initial support for HTTP2 connections into VMs
- Removed native support for gzip compression of HTTP requests

### Fixes

- Fixed message timeout handling on L1s configured with `validatorOnly=true`
- Fixed segfault on ARM64 when profiling is enabled

### What's Changed

- chore(tests/load): C-chain load testing by @qdm12 in https://github.com/ava-labs/avalanchego/pull/3914
- Improve comments on message Ops by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3987
- Remove requestID expectation for UnrequestedOps by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3989
- Add Simplex Messages To p2p.proto by @samliok in https://github.com/ava-labs/avalanchego/pull/3976
- [tmpnet] Enable exclusive scheduling by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3988
- [testing] Add local kube support for load tests by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/3986
- Add a label to XSVM tests by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3991
- [tmpnet] Avoid port forwarding when running in a kube cluster by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3997
- Add support for VM HTTP2 handlers by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3294
- Move HTTP2 routing information into headers by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/4001
- Clarify field names for simplex p2p messages by @samliok in https://github.com/ava-labs/avalanchego/pull/4002
- Remove gzip middleware from API server by @mpignatelli12 in https://github.com/ava-labs/avalanchego/pull/4005
- fix: allow for load tests to run in-cluster by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/4003
- chore: update libevm version by @alarso16 in https://github.com/ava-labs/avalanchego/pull/4006
- Add support for XSVM grpc server reflection by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/4010
- fix: update log commands for stopping collectors by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/4011
- [tmpnet] Enure node config is saved on restart for both runtimes by @maru-ava in https://github.com/ava-labs/avalanchego/pull/4015
- BLS Components for Simplex by @samliok in https://github.com/ava-labs/avalanchego/pull/3993
- Capitalize secrets references by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/4018
- build: update `libevm` to `v1.13.14-0.3.0.rc.1` by @alarso16 in https://github.com/ava-labs/avalanchego/pull/4023
- Allow internal messages from disallowed nodeIDs by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/4024
- optimize historical range by @rrazvan1 in https://github.com/ava-labs/avalanchego/pull/3658

### New Contributors

- @mpignatelli12 made their first contribution in https://github.com/ava-labs/avalanchego/pull/4005
- @alarso16 made their first contribution in https://github.com/ava-labs/avalanchego/pull/4006

**Full Changelog**: https://github.com/ava-labs/avalanchego/compare/v1.13.1...v1.13.2

## [v1.13.1](https://github.com/ava-labs/avalanchego/releases/tag/v1.13.1)

This version is backwards compatible to [v1.13.0](https://github.com/ava-labs/avalanchego/releases/tag/v1.13.0). It is optional, but encouraged.

The plugin version is updated to `40` all plugins must update to be compatible.

### APIs

- Removed `avm.getAddressTxs` api
- Added L1 validators to `platformvm.GetCurrentValidators` client implementation

### Configs

- Removed `--tracing-enabled` and added `disabled` as an option to `--tracing-exporter-type`
- Removed AVM indexer configs
  - `index-transactions`
  - `index-allow-incomplete`

### What's Changed

- Export tmpnet functions for CLI interface by @felipemadero in https://github.com/ava-labs/avalanchego/pull/3727
- Use IPv4 addresses if possible by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3812
- [ci] Source shellcheck from nix instead of installing via script by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3811
- [ci] Use SHAs instead of tags for 3rd-party github actions by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3822
- [tmpnet] Add collector log path to readiness check log output by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3823
- Add context for errors in proposervm `repairAcceptedChainByHeight` by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3818
- [ci] Enable run-monitored-tmpnet-cmd to use a remote flake file by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3820
- Add context to errors in `avm` by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3821
- Remove block reindexing after Etna by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3813
- Update protobuf dependencies to the same version as nix packages by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3828
- update: platformvm config doc by @ashucoder9 in https://github.com/ava-labs/avalanchego/pull/3694
- Expose merkledb defaults by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3748
- [tmpnet] Add table of contents to README by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3837
- fix(metrics): fix c-chain metrics not reporting by @darioush in https://github.com/ava-labs/avalanchego/pull/3835
- [tmpnet] Update script to run instead of install by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3830
- [ci] Update to use commit SHAs for non-floating tags by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3834
- [ci] Configure action/setup-go to read golang version from go.mod by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3825
- [nix] Install protobuf codegen binaries in the dev shell by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3829
- [tmpnet] Remove obsolete readme content by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3827
- [docs] Document requirement to install modern bash on macos by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3841
- refactor: use the built-in max/min to simplify the code by @evenevent in https://github.com/ava-labs/avalanchego/pull/3844
- Update BLST to v0.3.14 to support Go 1.24 by @yacovm in https://github.com/ava-labs/avalanchego/pull/3846
- chore: allow individuals to extend `direnv` config by @ARR4N in https://github.com/ava-labs/avalanchego/pull/3847
- [tmpnet] s/Network.ChainConfigs/Network.PrimaryChainConfigs/ by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3854
- Remove plugins/ from .gitignore by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3862
- fix: record wrong nil `err` by @tinyfoxy in https://github.com/ava-labs/avalanchego/pull/3851
- [tmpnet] Provide genesis, subnet and chain config via content flags by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3857
- Make sure inner state summary accept is called by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/3831
- avoid tmpnet to create empty genesis on disk by @felipemadero in https://github.com/ava-labs/avalanchego/pull/3868
- Remove GetAddressTxs by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3872
- [tmpnet] Fixed faulty error handling on bootstrap failure by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3873
- Set `subnets.Config.ConsensusParameters` to serialize with omitempty by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3874
- [tmpnet] Update rpc version check to tolerate usage of `go run` by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3869
- Update coreth to v0.15.1-rc.0 by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3875
- [tmpnet] Ensure Node has a reference to Network by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3870
- [tooling] Simplify avalanchego build script by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3861
- Simplify P-Chain block has changes check by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3880
- [tmpnet] Switch back to using maps for subnet config by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3877
- [tmpnet] Refactor runtime configuration in preparation for kube by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3867
- [tooling] Add scripts that build+run tools and put them in the path by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3878
- [tooling] Add support for the Task (go-task) task runner by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3863
- [Docs] Fix links and make paths absolute by @martineckardt in https://github.com/ava-labs/avalanchego/pull/3885
- Remove unused constant checkIndexedFrequency by @yacovm in https://github.com/ava-labs/avalanchego/pull/3887
- [tmpnet] Unify start network flag usage between e2e and tmpnetctl by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3871
- [tmpnet] Avoid serializing the node data directory by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3881
- [tmpnet] Rename NodeProcess to ProcessRuntime by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3890
- wrap db in initDatabase with corruptable db by @ceyonur in https://github.com/ava-labs/avalanchego/pull/3892
- [tmpnet] Switch FlagsMap from map[string]any to map[string]string by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3884
- [tmpnet] Ensure tmpnet methods always have a logger by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3893
- [tmpnet] Ensure all node runtime methods accept a context by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3894
- [tmpnet] Move WaitForHealthy from a function to a tmpnet.Node method by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3896
- Grant marun ownership of tooling configuration by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3895
- Bump golang.org/x/net from 0.36.0 to 0.38.0 by @dependabot in https://github.com/ava-labs/avalanchego/pull/3889
- Fix typos by @omahs in https://github.com/ava-labs/avalanchego/pull/3908
- Fix typos and add missing hyphens in README files by @Dimitrolito in https://github.com/ava-labs/avalanchego/pull/3583
- Fix typos in iterator.go by @Marcofann in https://github.com/ava-labs/avalanchego/pull/3809
- Use EstimateBaseFee in e2e tests by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3782
- docs: fix flag name to `--proposervm-min-block-delay` by @ceyonur in https://github.com/ava-labs/avalanchego/pull/3911
- fix rpcchainvm handling for arbitrary length http body by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3910
- Add git to nix packages by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3912
- refactor: use the built-in max/min to simplify the code by @careworry in https://github.com/ava-labs/avalanchego/pull/3913
- Update codeowners by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3915
- [tmpnet] Delegate writing of the flag file to the runtime by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3897
- [tmpnet] Move monitoring label handling to node by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3898
- Move database creation to database factory package by @ceyonur in https://github.com/ava-labs/avalanchego/pull/3899
- Update owner of CODEOWNERS by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3920
- update merkledb codeowners by @rrazvan1 in https://github.com/ava-labs/avalanchego/pull/3919
- chore: fix some comments by @standstaff in https://github.com/ava-labs/avalanchego/pull/3584
- Support UnmarshalJSON for `ExporterType` by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/3565
- Fix change/range proofs + simplify the code by @rrazvan1 in https://github.com/ava-labs/avalanchego/pull/3688
- [ci] Fix windows build job by reverting to use build script by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3921
- update: service.md for callouts by @ashucoder9 in https://github.com/ava-labs/avalanchego/pull/3832
- Reintroduce P-chain block reindexing by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3883
- Bump bufbuild/buf-action from 1.1.0 to 1.1.1 by @dependabot in https://github.com/ava-labs/avalanchego/pull/3855
- Bump github/codeql-action from 3.28.13 to 3.28.16 by @dependabot in https://github.com/ava-labs/avalanchego/pull/3916
- Ensure HTTP headers are propagated through the rpcchainvm by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3917
- Document acp-118 message verification by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3925
- Add P-Chain state test by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3924
- Document ACP-77 handling of 0 weight requests by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3926
- Refactor cache implementations by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3239
- Fix `proposerMinBlockDelay` location in config doc for L1s by @federiconardelli7 in https://github.com/ava-labs/avalanchego/pull/3819
- Close stale issues and PRs by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3906
- Update wallet to report tx processing duration via event handlers by @marun in https://github.com/ava-labs/avalanchego/pull/3560
- Pretty print logged durations in E2E by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3930
- Reenable the upgrade test by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3929
- Remove RequestBuildBlock on P-Chain Mempool by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3705
- Add test for proposervm BuildBlock after bootstrapping by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/2876
- Add logging to corruptabledb closure by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3938
- Fix broken Buf documentation link in proto README by @GarmashAlex in https://github.com/ava-labs/avalanchego/pull/3936
- refactor: replace []byte(fmt.Sprintf) with fmt.Appendf by @findnature in https://github.com/ava-labs/avalanchego/pull/3932
- Add linkspector CI action by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3939
- Update minimum golang version to v1.23.9 by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3940
- fix: validate allocations locked amount in genesis to prevent panic by @DracoLi in https://github.com/ava-labs/avalanchego/pull/3941
- [tmpnet] Enable runtime-specific restart behavior by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3882
- [tooling] Misc direnv changes by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3944
- Fully populate test context by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3943
- Use libevm instead of coreth by @ceyonur in https://github.com/ava-labs/avalanchego/pull/3918
- [tmpnet] Define reusable flags for configuring kubernetes client access by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3945
- Fix flaky bootstrapping test by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3955
- [tmpnet] Separate start of prometheus and promtail collectors by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3947
- Add L1 validators to getCurrentValidators response by @ceyonur in https://github.com/ava-labs/avalanchego/pull/3843
- refactor: use slices.Contains to simplify code by @yetyear in https://github.com/ava-labs/avalanchego/pull/3952
- Update proposervm summary to roll forward only by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/3950
- Remove dead code by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3966
- Add Granite to the `upgrade.Config` by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3964
- refactor genesis building logic in avm and platformvm by @DracoLi in https://github.com/ava-labs/avalanchego/pull/3949
- [ci] Update dependabot to only propose security updates for github actions by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3969
- Bump bufbuild/buf-action from 1.1.1 to 1.1.4 by @dependabot in https://github.com/ava-labs/avalanchego/pull/3971
- Bump github/codeql-action from 3.28.16 to 3.28.18 by @dependabot in https://github.com/ava-labs/avalanchego/pull/3970
- Add Load Framework by @RodrigoVillar in https://github.com/ava-labs/avalanchego/pull/3942
- adds config.json for C-Chain during antithesis - json logs by @aleksandarknezevic in https://github.com/ava-labs/avalanchego/pull/3968
- [tmpnet] Ensure GetNodeURIs returns locally-accessible URIs to ensure kube compatibility by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3973
- refactor: use slices.Contains to simplify code by @pullmerge in https://github.com/ava-labs/avalanchego/pull/3974
- chore(deps): use coreth with geth-aligned ethclient package by @qdm12 in https://github.com/ava-labs/avalanchego/pull/3977
- Remove dead “Turtle’s Way HTTP/gRPC” link by @gap-editor in https://github.com/ava-labs/avalanchego/pull/3978
- Small cleanup in `App` and `Node` by @geoff-vball in https://github.com/ava-labs/avalanchego/pull/3962
- make GOPROXY overridable in constants.sh by @siphonelee in https://github.com/ava-labs/avalanchego/pull/3979
- Disallow slow sorting by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3981
- [tmpnet] Enable deployment to kube by @marun in https://github.com/ava-labs/avalanchego/pull/3615
- [tmpnet] Enable monitoring of nodes running in kube by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3794
- Bump coreth to include fix for large tx handling by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/3984
- Align minCompatibleTime settings across TestNetwork and Network by @michaelkaplan13 in https://github.com/ava-labs/avalanchego/pull/3842
- Remove dead “Turtle’s Way HTTP/gRPC” link by @gap-editor in https://github.com/ava-labs/avalanchego/pull/3983

### New Contributors

- @evenevent made their first contribution in https://github.com/ava-labs/avalanchego/pull/3844
- @tinyfoxy made their first contribution in https://github.com/ava-labs/avalanchego/pull/3851
- @Dimitrolito made their first contribution in https://github.com/ava-labs/avalanchego/pull/3583
- @Marcofann made their first contribution in https://github.com/ava-labs/avalanchego/pull/3809
- @careworry made their first contribution in https://github.com/ava-labs/avalanchego/pull/3913
- @standstaff made their first contribution in https://github.com/ava-labs/avalanchego/pull/3584
- @RodrigoVillar made their first contribution in https://github.com/ava-labs/avalanchego/pull/3565
- @federiconardelli7 made their first contribution in https://github.com/ava-labs/avalanchego/pull/3819
- @GarmashAlex made their first contribution in https://github.com/ava-labs/avalanchego/pull/3936
- @findnature made their first contribution in https://github.com/ava-labs/avalanchego/pull/3932
- @yetyear made their first contribution in https://github.com/ava-labs/avalanchego/pull/3952
- @aleksandarknezevic made their first contribution in https://github.com/ava-labs/avalanchego/pull/3968
- @pullmerge made their first contribution in https://github.com/ava-labs/avalanchego/pull/3974
- @gap-editor made their first contribution in https://github.com/ava-labs/avalanchego/pull/3978
- @geoff-vball made their first contribution in https://github.com/ava-labs/avalanchego/pull/3962
- @siphonelee made their first contribution in https://github.com/ava-labs/avalanchego/pull/3979

**Full Changelog**: https://github.com/ava-labs/avalanchego/compare/v1.13.0...v1.13.1

## [v1.13.0](https://github.com/ava-labs/avalanchego/releases/tag/v1.13.0)

This upgrade consists of the following Avalanche Community Proposal (ACP):
- [ACP-176](https://github.com/avalanche-foundation/ACPs/blob/main/ACPs/176-dynamic-evm-gas-limit-and-price-discovery-updates/README.md) Dynamic EVM Gas Limits and Price Discovery Updates

The ACP in this upgrade goes into effect at 11 AM ET (3 PM UTC) on Tuesday, April 8th, 2025 on Mainnet.

**All Fortuna supporting Mainnet nodes should upgrade before 11 AM ET, April 8th 2025.**

The plugin version is unchanged at `39` and is compatible with version `v1.12.2`.

### APIs

- Added ProposerVM block timestamp metrics: `avalanche_proposervm_last_accepted_timestamp`
- Added network health check to alert if a primary network validator has no ingress connections. Runs a configurable time after startup or 10 minutes by default.

### Configs

- Added:
  - `--proposervm-min-block-delay`
  - `--network-no-ingress-connections-grace-period` to configure how long after startup it is expected for a Mainnet validator to have received an ingress connection.

### What's Changed

- Implement traversal based early termination by @yacovm in https://github.com/ava-labs/avalanchego/pull/3337
- Add networked-signer tests by @richardpringle in https://github.com/ava-labs/avalanchego/pull/3613
- Remove unused code by @yacovm in https://github.com/ava-labs/avalanchego/pull/3682
- chore(proposervm): timestamp metrics for block acceptance by @ARR4N in https://github.com/ava-labs/avalanchego/pull/3680
- remove dependency of validator.State from BitSetSignature.Verify by @tsachiherman in https://github.com/ava-labs/avalanchego/pull/3679
- [docker] Optimize build time by copying and downloading deps first by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3683
- fix: broken link in README.md by @DeVikingMark in https://github.com/ava-labs/avalanchego/pull/3681
- [docker] Silence remaining InvalidDefaultArgInFrom warnings by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3684
- [tmpnet] Update subnet configuration in README by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3686
- testing: improve e2e test bootstrapping by @tsachiherman in https://github.com/ava-labs/avalanchego/pull/3690
- [tmpnet] Update URI and StakingAddress usage in support of kube by @marun in https://github.com/ava-labs/avalanchego/pull/3665
- [tmpnet] Re-enable reuse of dynamically allocated API ports by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3697
- fix spelling issues  by @futreall in https://github.com/ava-labs/avalanchego/pull/3700
- fix: correct typos in parser.go and tmpnet documentation by @avorylli in https://github.com/ava-labs/avalanchego/pull/3712
- fix: typos in documentation files by @maximevtush in https://github.com/ava-labs/avalanchego/pull/3710
- Fail fast in tests if avalancheGo executable isn't an absolute path by @yacovm in https://github.com/ava-labs/avalanchego/pull/3707
- [ci] Fix metrics link annotation emitted for e2e and upgrade jobs by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3713
- Remove Mock Mempool by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3687
- cleanup(tmpnet): resolve chainconfig post-etna TODO by @darioush in https://github.com/ava-labs/avalanchego/pull/3720
- Update to go 1.23.6 by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3722
- docs: Fix grammatical errors and improve clarity in documentation and comments by @VolodymyrBg in https://github.com/ava-labs/avalanchego/pull/3716
- [testing] Replace script-based tool installation with nix by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3691
- Remove unnecessary function by @richardpringle in https://github.com/ava-labs/avalanchego/pull/3723
- bump coreth to master by @darioush in https://github.com/ava-labs/avalanchego/pull/3724
- Make `bls.Signer` api fallible by @richardpringle in https://github.com/ava-labs/avalanchego/pull/3696
- Add comment to seemingly dead code by @richardpringle in https://github.com/ava-labs/avalanchego/pull/3721
- Bump coreth by @richardpringle in https://github.com/ava-labs/avalanchego/pull/3728
- Comment on the need for `CGO_ENABLED=1` to support cross-compilation by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3735
- [ci] Update to golangci-lint version compatible with go 1.23 by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3739
- [testing] Provide more logging context for SynchronizedBeforeSuite by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3741
- refactor: export PeerSample by @Elvis339 in https://github.com/ava-labs/avalanchego/pull/3745
- [antithesis] Set AVAGO_PLUGIN_DIR for VM images by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3751
- [ci] Drop support for Ubuntu 20.04 by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3737
- Fix spelling errors in `majority.go`, `minority.go`, `compressor.go`, and `logger.go` by @tomasandroil in https://github.com/ava-labs/avalanchego/pull/3738
- [ci] Simplify tmpnet monitoring action by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3736
- Remove apostrophe from Dockerfile comments by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/3706
- fix error 404 link README.md by @futreall in https://github.com/ava-labs/avalanchego/pull/3750
- fix: typos in documentation files by @leopardracer in https://github.com/ava-labs/avalanchego/pull/3733
- Add With and WithOptions receivers to the Logger interface by @iansuvak in https://github.com/ava-labs/avalanchego/pull/3729
- [tmpnet] Minimize duration of tx acceptance for e2e testing by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3685
- Remove unused `ForceCreateChain` function from `testManager` by @strmfos in https://github.com/ava-labs/avalanchego/pull/3755
- chore: make function comments match function names by @rustco in https://github.com/ava-labs/avalanchego/pull/3757
- L1 validator eviction block validity by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3758
- Add ACP-176 e2e tests by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3749
- [tmpnet] Deploy collectors with golang to simplify cross-repo use by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3692
- fix spelling issues config_test.go by @futreall in https://github.com/ava-labs/avalanchego/pull/3760
- [tmpnet] Add check for collection of logs and metrics to custom github action by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3740
- Deprecate the `snow.Context.Lock` by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3762
- Name F-Upgrade Fortuna by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3761
- Add CodecID to ICM README by @iansuvak in https://github.com/ava-labs/avalanchego/pull/3759
- Update CODEOWNERS s/marun/maru-ava/ by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3768
- Support caller-defined namespaces in merkledb by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3747
- Fix empty standard block check by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3775
- Enable empty standard block check by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3776
- Restrict ProposerVM P-chain height advancement by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3777
- Add canoto serialization support to the block context by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/3709
- Remove support for AVM tx checksums by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3774
- [ci] Disable monitoring for jobs of PRs of fork branches by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3781
- Update db_test.go by @sky-coderay in https://github.com/ava-labs/avalanchego/pull/3765
- chore: fix some function names in comment by @tcpdumppy in https://github.com/ava-labs/avalanchego/pull/3773
- Implement ACP-118 Aggregator by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3394
- Print git commit version upon startup by @yacovm in https://github.com/ava-labs/avalanchego/pull/3771
- chore(nix): add darwin.apple_sdk.frameworks.Security by @darioush in https://github.com/ava-labs/avalanchego/pull/3769
- Add comment to SendConfig documenting how to broadcast by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/3783
- refactor: use a more straightforward return value by @fuyangpengqi in https://github.com/ava-labs/avalanchego/pull/3726
- [ci] Stop emitting grafana link as an annotation by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3767
- Remove Etna activation banner by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3789
- Upgrade canoto to v0.13.3 by @tsachiherman in https://github.com/ava-labs/avalanchego/pull/3790
- Healthcheck for zero ingress connection count by @yacovm in https://github.com/ava-labs/avalanchego/pull/3719
- Timely halt Avalanche engine by @yacovm in https://github.com/ava-labs/avalanchego/pull/3792
- [docker] Update all images to debian12/bookworm by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3798
- [ci] Move monitoring check from github action to code by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3766
- [tmpnet] Start kind cluster with golang by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3780
- Fix flake TestIngressConnCount by @yacovm in https://github.com/ava-labs/avalanchego/pull/3799
- [tmpnet] Rename tmpnetctl main package from cmd to tmpnetctl by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3787
- Improve check-clean CI script by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3800
- Bump golang.org/x/net from 0.33.0 to 0.36.0 by @dependabot in https://github.com/ava-labs/avalanchego/pull/3793
- Fix ACP-118 Aggregator Test Flake by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3801
- Canoto v0.15.0 upgrade by @tsachiherman in https://github.com/ava-labs/avalanchego/pull/3805
- [tmpnet] Fix README example for tmpnetctl script by @maru-ava in https://github.com/ava-labs/avalanchego/pull/3807
- Update coreth to v0.15.0-rc.0 by @darioush in https://github.com/ava-labs/avalanchego/pull/3808

### New Contributors

- @maru-ava made their first contribution in https://github.com/ava-labs/avalanchego/pull/3683
- @DeVikingMark made their first contribution in https://github.com/ava-labs/avalanchego/pull/3681
- @futreall made their first contribution in https://github.com/ava-labs/avalanchego/pull/3700
- @avorylli made their first contribution in https://github.com/ava-labs/avalanchego/pull/3712
- @maximevtush made their first contribution in https://github.com/ava-labs/avalanchego/pull/3710
- @VolodymyrBg made their first contribution in https://github.com/ava-labs/avalanchego/pull/3716
- @Elvis339 made their first contribution in https://github.com/ava-labs/avalanchego/pull/3745
- @tomasandroil made their first contribution in https://github.com/ava-labs/avalanchego/pull/3738
- @leopardracer made their first contribution in https://github.com/ava-labs/avalanchego/pull/3733
- @strmfos made their first contribution in https://github.com/ava-labs/avalanchego/pull/3755
- @rustco made their first contribution in https://github.com/ava-labs/avalanchego/pull/3757
- @sky-coderay made their first contribution in https://github.com/ava-labs/avalanchego/pull/3765
- @tcpdumppy made their first contribution in https://github.com/ava-labs/avalanchego/pull/3773
- @fuyangpengqi made their first contribution in https://github.com/ava-labs/avalanchego/pull/3726

**Full Changelog**: https://github.com/ava-labs/avalanchego/compare/v1.12.2...v1.13.0

## [v1.12.2](https://github.com/ava-labs/avalanchego/releases/tag/v1.12.2)

This version is backwards compatible to [v1.12.0](https://github.com/ava-labs/avalanchego/releases/tag/v1.12.0). It is optional, but encouraged.

The plugin version is updated to `39` all plugins must update to be compatible.

**This release removes the support for the long deprecated Keystore API. Any users still relying on the keystore API will not be able to update to this version, or any later versions, of Avalanchego until their dependency on the keystore API has been removed.**

### APIs

- Deprecated:
  - `info.GetTxFee`
- Added:
  - `avm.GetTxFee`
  - `platform.getValidatorFeeConfig`
  - `platform.getValidatorFeeState`
  - `validationID` field to `platform.getL1Validator` results
  - L1 validators to `platform.getCurrentValidators`
- Removed:
  - `StakeAmount` field from `platform.getCurrentValidators` results
  - `keystore.createUser`
  - `keystore.deleteUser`
  - `keystore.listUsers`
  - `keystore.importUser`
  - `keystore.exportUser`
  - `avm.createAddress`
  - `avm.createFixedCapAsset`
  - `avm.createNFTAsset`
  - `avm.createVariableCapAsset`
  - `avm.export`
  - `avm.exportKey`
  - `avm.import`
  - `avm.importKey`
  - `avm.listAddresses`
  - `avm.mint`
  - `avm.mintNFT`
  - `avm.send`
  - `avm.sendMultiple`
  - `avm.sendNFT`
  - `wallet.send`
  - `wallet.sendMultiple`
  - `platform.exportKey`
  - `platform.listAddresses`

### Configs

- Removed static fee config flags
  - `--create-subnet-tx-fee`
  - `--transform-subnet-tx-fee`
  - `--create-blockchain-tx-fee`
  - `--add-primary-network-validator-fee`
  - `--add-primary-network-delegator-fee`
  - `--add-subnet-validator-fee`
  - `--add-subnet-delegator-fee`
- Removed `--api-keystore-enabled`

### What's Changed

- [testing] Always use the go.mod version of ginkgo by @marun in https://github.com/ava-labs/avalanchego/pull/3618
- Bump antithesishq/antithesis-trigger-action from 0.5 to 0.6 by @dependabot in https://github.com/ava-labs/avalanchego/pull/3620
- Fix typos in document files by @taozui472 in https://github.com/ava-labs/avalanchego/pull/3622
- [ci] Always use the specified go version by @marun in https://github.com/ava-labs/avalanchego/pull/3616
- Fix: quotation mark by @jasmyhigh in https://github.com/ava-labs/avalanchego/pull/3623
- refactor: move node configs to config/node by @darioush in https://github.com/ava-labs/avalanchego/pull/3600
- Update e2e tests and CI jobs for post-etna by @marun in https://github.com/ava-labs/avalanchego/pull/3614
- Replace AWM terminology in ReadMe with ICM  by @meaghanfitzgerald in https://github.com/ava-labs/avalanchego/pull/3595
- fix: grammatical mistakes by @crStiv in https://github.com/ava-labs/avalanchego/pull/3625
- [testing] Update golangci-lint to latest version by @marun in https://github.com/ava-labs/avalanchego/pull/3617
- partial sync default info by @meaghanfitzgerald in https://github.com/ava-labs/avalanchego/pull/3602
- Update stale comment on commitToDB by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/3627
- coreth atomic pkg dependency by @ceyonur in https://github.com/ava-labs/avalanchego/pull/3588
- Mark Meag as the owner of README files by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3635
- Update x/net to v0.33.0 by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3636
- Fix codeowners to simplify PR review by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3637
- Add BLS healthcheck to communicate incorrect BLS key configuration by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3638
- chore(all): mocks generation improved by @qdm12 in https://github.com/ava-labs/avalanchego/pull/3628
- fix LRU sized cache: consistent size at element removal by @rrazvan1 in https://github.com/ava-labs/avalanchego/pull/3634
- Index API and AvalancheGo Configs Docs Fix by @meaghanfitzgerald in https://github.com/ava-labs/avalanchego/pull/3632
- [ci] Migrate from buf-*-action to buf-action by @marun in https://github.com/ava-labs/avalanchego/pull/3639
- chore(ci): define Github labels as code with a workflow by @qdm12 in https://github.com/ava-labs/avalanchego/pull/3629
- feat(github): add "needs Go upgrade" label by @qdm12 in https://github.com/ava-labs/avalanchego/pull/3642
- [ci] Fix post-merge protobuf lint job breakage by @marun in https://github.com/ava-labs/avalanchego/pull/3644
- merkledb visualisations v1 (change proofs and range proofs) by @rrazvan1 in https://github.com/ava-labs/avalanchego/pull/3643
- Remove Static Fee Config by @samliok in https://github.com/ava-labs/avalanchego/pull/3610
- fix(ci): trigger labels workflow on push to master not main by @qdm12 in https://github.com/ava-labs/avalanchego/pull/3646
- X-Chain API fix by @meaghanfitzgerald in https://github.com/ava-labs/avalanchego/pull/3654
- [ci] Rename {PROMETHEUS,LOKI}_ID to {PROMETHEUS,LOKI}_USERNAME by @marun in https://github.com/ava-labs/avalanchego/pull/3652
- chore: replaced faulty link by @Radovenchyk in https://github.com/ava-labs/avalanchego/pull/3649
- Add L1 validator fees API by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3647
- Reintroduce the deprecated `info.getTxFee` API by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3656
- remove x-chain api obsolete metadata  by @meaghanfitzgerald in https://github.com/ava-labs/avalanchego/pull/3655
- Remove the Keystore API by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3657
- Add F Upgrade Scaffolding. Post-Etna Cleanup by @michaelkaplan13 in https://github.com/ava-labs/avalanchego/pull/3672
- [testing] Fix instructions for triggering antithesis test runs by @marun in https://github.com/ava-labs/avalanchego/pull/3664
- [testing] Ensure run_prometheus.sh uses a writeable storage path by @marun in https://github.com/ava-labs/avalanchego/pull/3662
- Make snowman use snowflake directly instead of snowball by @yacovm in https://github.com/ava-labs/avalanchego/pull/3403
- chore: fix some typos by @chuangjinglu in https://github.com/ava-labs/avalanchego/pull/3670
- Bump antithesishq/antithesis-trigger-action from 0.6 to 0.7 by @dependabot in https://github.com/ava-labs/avalanchego/pull/3667
- [ci] Use go env {GOOS,GOARCH} for os and arch detection by @marun in https://github.com/ava-labs/avalanchego/pull/3661
- Silence docker InvalidDefaultArgInFrom warnings by @marun in https://github.com/ava-labs/avalanchego/pull/3659
- add L1 support to getCurrentValidators API by @ceyonur in https://github.com/ava-labs/avalanchego/pull/3564
- [docker] Enable image builds from git worktrees by @marun in https://github.com/ava-labs/avalanchego/pull/3660
- [tmpnet] Set an explicit `instance` label for logs and metrics by @marun in https://github.com/ava-labs/avalanchego/pull/3650
- [docker] Switch to kube-compatible plugin path for images by @marun in https://github.com/ava-labs/avalanchego/pull/3653
- [testing] Support direnv to simplify usage of test tooling by @marun in https://github.com/ava-labs/avalanchego/pull/3651

### New Contributors

- @taozui472 made their first contribution in https://github.com/ava-labs/avalanchego/pull/3622
- @jasmyhigh made their first contribution in https://github.com/ava-labs/avalanchego/pull/3623
- @crStiv made their first contribution in https://github.com/ava-labs/avalanchego/pull/3625
- @qdm12 made their first contribution in https://github.com/ava-labs/avalanchego/pull/3628
- @rrazvan1 made their first contribution in https://github.com/ava-labs/avalanchego/pull/3634
- @Radovenchyk made their first contribution in https://github.com/ava-labs/avalanchego/pull/3649
- @chuangjinglu made their first contribution in https://github.com/ava-labs/avalanchego/pull/3670

**Full Changelog**: https://github.com/ava-labs/avalanchego/compare/v1.12.1...v1.12.2

## [v1.12.1](https://github.com/ava-labs/avalanchego/releases/tag/v1.12.1)

This version is backwards compatible to [v1.12.0](https://github.com/ava-labs/avalanchego/releases/tag/v1.12.0). It is optional, but encouraged.

The plugin version is unchanged at `38` and is compatible with version `v1.12.0`.

### Configs

- Added PebbleDB option `sync` which defaults to `true`

### Fixes

- Fixed P-chain mempool verification to disallow transactions that exceed the available chain capacity

### What's Changed

- Expose test network cfg by @cam-schultz in https://github.com/ava-labs/avalanchego/pull/3573
- encapsulate signer by @richardpringle in https://github.com/ava-labs/avalanchego/pull/3576
- use pebble nosync by default by @ceyonur in https://github.com/ava-labs/avalanchego/pull/3581
- fix: feeState API call in docs by @ashucoder9 in https://github.com/ava-labs/avalanchego/pull/3596
- Format Service.MD by @samliok in https://github.com/ava-labs/avalanchego/pull/3599
- Verify tx gas isn't too large in VerifyTx by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3604
- Add already implemented merkledb.View to MerkleDB interface by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/3593
- Add tempdir in for chain ctx data dir by @aaronbuchwald in https://github.com/ava-labs/avalanchego/pull/3594
- Improve block building and verification logging by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3605
- Bump golang.org/x/crypto from 0.26.0 to 0.31.0 by @dependabot in https://github.com/ava-labs/avalanchego/pull/3608

**Full Changelog**: https://github.com/ava-labs/avalanchego/compare/v1.12.0...v1.12.1

## [v1.12.0](https://github.com/ava-labs/avalanchego/releases/tag/v1.12.0)

This upgrade consists of the following Avalanche Community Proposals (ACPs):
- [ACP-77](https://github.com/avalanche-foundation/ACPs/blob/main/ACPs/77-reinventing-subnets/README.md) Reinventing Subnets
- [ACP-103](https://github.com/avalanche-foundation/ACPs/blob/main/ACPs/103-dynamic-fees/README.md) Add Dynamic Fees to the P-Chain
- [ACP-118](https://github.com/avalanche-foundation/ACPs/blob/main/ACPs/118-warp-signature-request/README.md) Warp Signature Interface Standard
- [ACP-125](https://github.com/avalanche-foundation/ACPs/blob/main/ACPs/125-basefee-reduction/README.md) Reduce C-Chain minimum base fee from 25 nAVAX to 1 nAVAX
- [ACP-131](https://github.com/avalanche-foundation/ACPs/blob/main/ACPs/131-cancun-eips/README.md) Activate Cancun EIPs on C-Chain and Subnet-EVM chains
- [ACP-151](https://github.com/avalanche-foundation/ACPs/blob/main/ACPs/151-use-current-block-pchain-height-as-context/README.md) Use current block P-Chain height as context for state verification

The changes in the upgrade go into effect at 12 AM ET (5 PM UTC) on Monday, December 16th, 2024 on Mainnet.

**All Etna supporting Mainnet nodes should upgrade before 12 AM ET, December 16th 2024.**

The plugin version is unchanged at `38` and is compatible with version `v1.11.13`.

### APIs

- Allowed `platform.issueTx` to be called, for non-ImportTx transactions, while partial syncing

### What's Changed

- Fix SubnetToL1ConversionData typo by @cam-schultz in https://github.com/ava-labs/avalanchego/pull/3555
- Refactor `logging.Format` to expose constants by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3561
- [testing] Switch to logging with zap by @marun in https://github.com/ava-labs/avalanchego/pull/3557
- Use JSON logs during Antithesis runs by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3562
- Antithesis: Skip checks if tx confirmation fails by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3563
- chore: fix some function names in comment by @wanxiangchwng in https://github.com/ava-labs/avalanchego/pull/3566
- update api docs by @ashucoder9 in https://github.com/ava-labs/avalanchego/pull/3558
- Remove unused wallet interface by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3568
- Remove required fields from config by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3569
- Remove redundant field in platformVM/network's Network by @yacovm in https://github.com/ava-labs/avalanchego/pull/3571
- Remove observedSubnetUptime from Info Docs by @samliok in https://github.com/ava-labs/avalanchego/pull/3575
- Allow issuing transactions when using partial-sync by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3570
- Add partial-sync support to the wallet by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3567

### New Contributors

- @wanxiangchwng made their first contribution in https://github.com/ava-labs/avalanchego/pull/3566
- @ashucoder9 made their first contribution in https://github.com/ava-labs/avalanchego/pull/3558

**Full Changelog**: https://github.com/ava-labs/avalanchego/compare/v1.11.13...v1.12.0

## [v1.11.13](https://github.com/ava-labs/avalanchego/releases/tag/v1.11.13)

This version is backwards compatible to [v1.11.0](https://github.com/ava-labs/avalanchego/releases/tag/v1.11.0). It is optional, but encouraged.

The plugin version is updated to `38` all plugins must update to be compatible.

### APIs

- Added `platform.getL1Validator`
- Added `platform.getProposedHeight`
- Updated `platform.getValidatorsAt` to accept `"proposed"` as valid `height` input

### Configs

- Added P-chain configs
  - `"l1-weights-cache-size"`
  - `"l1-inactive-validators-cache-size"`
  - `"l1-subnet-id-node-id-cache-size"`

### Fixes

- Fixed metrics initialization in the RPCChainVM. This could cause crashes during startup if metrics was requested during VM initialization.
- Fixed compilations on macos 14.7 and higher
- Fixed avalanchego wallet usage with ledger v0.8.4
- Fixed missing `NodeIDs` argument in the `info.peers` client implementation
- Fixed `getSubnetID` state tracing

### What's Changed

- [testing] Double image build timeout for bootstrap monitor e2e by @marun in https://github.com/ava-labs/avalanchego/pull/3468
- [antithesis] Double the duration of sanity checks by @marun in https://github.com/ava-labs/avalanchego/pull/3475
- Properly initialize metrics in rpcchainVM by @yacovm in https://github.com/ava-labs/avalanchego/pull/3477
- Cleanup editorconfig by @dhrubabasu in https://github.com/ava-labs/avalanchego/pull/3473
- Update avalanche ledger go package by @sukantoraymond in https://github.com/ava-labs/avalanchego/pull/3456
- [testing] Enable config of log format for bootstrap monitor by @marun in https://github.com/ava-labs/avalanchego/pull/3467
- cache signatures only in acp118 handler by @ceyonur in https://github.com/ava-labs/avalanchego/pull/3474
- Introduce and use `database.WithDefault` by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3478
- Evict recentlyAccepted blocks based on wall-clock time  by @iansuvak in https://github.com/ava-labs/avalanchego/pull/3460
- fix improper use of FailNow in testing by @tsachiherman in https://github.com/ava-labs/avalanchego/pull/3479
- [ACP 151] Use current block's P-Chain height as context for verifying state of the inner block by @iansuvak in https://github.com/ava-labs/avalanchego/pull/3459
- [tmpnet] Add --start-network to support hypersdk MODE=run by @marun in https://github.com/ava-labs/avalanchego/pull/3465
- [e2e] Check network health after bootstrap checks by @marun in https://github.com/ava-labs/avalanchego/pull/3466
- ACP-77: Add ConversionID to state by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3481
- Make bootstrapping handle its own timeouts by @yacovm in https://github.com/ava-labs/avalanchego/pull/3410
- Wrap `TestDiffExpiry` sub-tests in `t.Run` by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3483
- Move RPC metrics registration after its client's initialization by @yacovm in https://github.com/ava-labs/avalanchego/pull/3488
- database: add applicable dbtests for linkeddb by @darioush in https://github.com/ava-labs/avalanchego/pull/3486
- Add SoV Excess to P-chain state by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3482
- Remove deprecated X-chain pubsub server by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3490
- Update SoV struct to align with latest ACP-77 spec by @StephenButtolph in https://github.com/ava-labs/avalanchego/pull/3492
- Register VM and snowman metrics after chain creation by @yacovm in https://github.com/ava-labs/avalanchego/pull/3489
- Skip Flaky Test by @joshua-kim in https://github.com/ava-labs/avalanchego/pull/3495
- Add request to u