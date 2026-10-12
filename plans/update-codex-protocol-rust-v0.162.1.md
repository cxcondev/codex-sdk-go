# Update Codex protocol to rust-v0.162.1

This living ExecPlan follows `PLANS.md`. Keep Progress, discoveries, decisions, and outcomes current.

## Purpose / Big Picture


Update the Go SDK from Codex 0.160.0 to stable 0.162.1 so callers can use current attachment-owner lookup, prediction notifications, and turn attribution. Preserve intentional wire compatibility and publish SDK v0.162.1 through protected GitHub checks and trusted end-to-end tests.

## Tracker Mapping


Workflow: `WORKFLOW.md`. Linear project: P-COLIN-6 (codex-sdk). Issue: COLIN-225, In Progress, owns all three milestones. The completed 0.160.0 plan remains historical.

## Progress


- [x] (2026-10-12 UTC) Verified clean main at origin/main, 6d56fc295d74d6133eb5d8d618c24c49e33363a9, instructions, latest SDK v0.160.0, and upstream highest exact stable rust-v0.162.1.
- [x] (2026-10-12 UTC) Validated required planner proposal against scripts, schemas, and live protections; created COLIN-225 and codex/update-protocol-v0.162.1.
- [x] (2026-10-12 UTC) Ran updater and printed union guard; all 447 canonical exporter entries exactly match the independently reviewed target source inventory.
- [x] (2026-10-12 UTC) Reviewed printed opaque inventory: 104 entries, exactly the reproduced 105-entry baseline minus the now-lossless CodexErrorInfo interface.
- [x] (2026-10-12 UTC) Repaired compatibility and metadata; two byte-identical generations, formatting, fixtures, vet, unit/race tests, and Staticcheck passed.
- [x] (2026-10-12 UTC) Verified credential-free archive/lifecycle and credential-isolation race tests against checksum-verified official 0.162.1 CLI.
- [x] (2026-10-12 UTC) Corrected Go patch pins and passed the authoritative complete Go 1.26.9 gate, including determinism and zero reachable vulnerabilities. Go 1.27.2 native vet/unit/race/vulnerability checks and supported-toolchain Staticcheck/actionlint also pass.
- [x] (2026-10-12 UTC) Completed independent QA, architecture, and security reviews. QA/security found no actionable issues; fixed architecture's stale plan statement and orphaned documentation finding with a failing-then-passing regression.
- [x] (2026-10-12 UTC) Completed regeneration and the authoritative gate after review fixes: two identical generations, all local checks, no vulnerabilities.
- [ ] Stage named reviewed paths, commit, open protected PR, pass hosted checks, merge, and verify automatic release and immutable tag.

## Surprises & Discoveries


Installed Go is 1.27.1 and Cargo 1.94.0; upstream requests Rust 1.95.0. Check actual compilation before changing tools. The source checkout `/Users/pme/src/openai/codex` is clean and must remain unchanged. Git metadata writes require sandbox escalation. Never source or print `.envrc`.

Govulncheck found nine reachable standard-library advisories in Go 1.26.8, all fixed in Go 1.26.9: GO-2026-6617, 6613, 6612, 6611, 6610, 6608, 6607, 6605, and 6603. Determinism and all earlier gates passed; this is a toolchain gate failure, not a protocol regression. Official go.dev download metadata confirms stable Go 1.26.9 and 1.27.2, both verified locally through GOTOOLCHAIN.

## Decision Log


Use upstream rust-v0.162.1 at 092d3acd6bec3e3a14bdc7e7a2810ab628ab759d and SDK 0.162.1 because it exceeds latest immutable v0.160.0. Recheck before publishing. The direct skill invocation authorizes implementation and protected publication. No additional plan approval is required. Use the existing flat plans/ convention and temporary logs outside the repository.

Goal mutation origin must remain optional; absence never implies user authorization. Add nullable parentTurnId/rootTurnId to handwritten TurnStartParams and rootTurnId to Turn. Keep future error strings/objects and arbitrary additional_tools payloads lossless. Preserve Go aliases where identity matters; document removed namespaceTools rather than sending unsupported fields. Review every canonical union delta before accepting the aggregate hash and every opaque declaration before its digest.

Approve union digest 0a713f8e5cdc656a268e0a295111528b7cf7787cf8bf9e61c286f396c0bb310b after reviewing all 20 added/13 removed entries and proving all 447 printed canonical entries exactly equal target source. Changes are nullable extension/reasoning/origin refs, partial_answer, future string/object errors, additional_tools, prediction completed/failed, and enclosing request/notification containers. Existing raw union wrappers handle tools and predictions with required-member validation. CodexErrorInfo becomes a manual raw JSON alias to prevent demonstrated precision loss; no new generic union fallback is approved.

Approve opaque digest aa9797066f9ef641746535a494de20b16e172ee23352cb1dcd8d1bf398e4a743 after reproducing the old 105-entry digest with the generator's AST algorithm and comparing the complete printed 104-entry target. Exactly type CodexErrorInfo interface{} was removed; no opaque type or field was added.

Live protection requires strict Coverage report, Go 1.26, and Go 1.27 checks, zero approving reviews, and resolved conversations. Both e2e and release environments permit protected branches with no required reviewers. An active immutable-tag ruleset exists. Verify detailed tag rules before publication. No settings change is currently needed.

Keep supported Go minor families and stable hosted check names; update only fixed patch pins to 1.26.9/1.27.2 in go.mod, tests.yml, README.md, APP.md, and WORKFLOW.md. Validated planner amendment and official metadata justify this security correction within COLIN-225 to satisfy the authorized gate. Historical plans remain unchanged.

Go 1.27.2 passed native vet, unit/race tests, and govulncheck, but Staticcheck 0.8.1 (also the latest module release) cannot decode its export-data version 5. Validated planner amendment: keep pinned analyzers unchanged and run Staticcheck through go run with GOTOOLCHAIN=go1.26.9 in both matrix jobs. This analyzes all SDK source with its patched minimum toolchain; all other matrix checks remain native. No diagnostic is ignored or static gate skipped. Workflow and plan document this temporary analyzer limitation.

## Outcomes & Retrospective


Protocol generation, compatibility regressions, real-CLI lifecycle verification, and the complete patched-toolchain local gate passed. Independent reviews are complete and their two cleanup findings are fixed; the final authoritative regeneration/gate passed again. Protected publication remains.

Initial regressions reproduced missing goal origin, parent/root turn attribution, and unrecognized additional_tools. A future error containing 9007199254740993 also round-tripped incorrectly as 9007199254740992 through the existing interface representation. CodexErrorInfo now aliases json.RawMessage and the intentional low-level source change is documented. Manual ancestry fields pass exact source-schema coverage; both inventory digests are reviewed and approved.

## Context and Orientation


`gen.go` invokes `internal/codegen/main.go`, compiling an exact upstream source exporter. Generated JSON wire types live in `protocol/*_gen.go`; generated JSON-RPC calls and notification parsing live in `rpc/*_gen.go`. JSON-RPC is the request/notification transport to the local Codex app-server. Handwritten wrappers live in `protocol/manual_types.go` and SDK adapters at repository root. Fix generated declarations through the generator, never by hand. Every exported generated declaration needs exact-name GoDoc.

`.github/codex/version` and `.github/codex/checksums.txt` select the verified CLI used by real tests. `.github/sdk-version` selects the independent SDK version. A protected main merge changing it starts `.github/workflows/release.yml`, which validates the candidate, runs Quality and trusted E2E, and alone creates the annotated immutable tag.

## Plan of Work


Milestone 1 runs the bundled updater from the feature branch. If an inventory guard fails, rerun with printed inventory flags and compare every added, removed, or changed canonical JSON union against exported target and committed baseline schemas. Record typed handling or explicit opaque decisions, then approve digests in internal/codegen/main.go. A union is a schema with oneOf/anyOf alternatives; canonical JSON and its hash identify exactly what was reviewed.

Milestone 2 inspects the complete generated diff and adds necessary manual fields and wire regressions. Review partial_answer message phase, future CodexErrorInfo, ThreadPredictionResult completed/failed variants, attachment-owner RPC, arbitrary tools, browser extension headers, skill paths, OAuth loginId, reviewTarget, and subagent model/reasoning metadata. Update official four-platform CLI digests and SDK 0.162.1. Run meaningful regression tests before fixes where feasible, manual source-schema coverage, focused protocol/RPC tests, then authoritative deterministic generation and all gates. Keep high-level API changes limited to actual compatibility needs.

Milestone 3 runs QA, architecture, and security reviews after local verification; resolve actionable findings. Stage only reviewed named paths, commit with COLIN-225, push the feature branch, open and attach a PR to main. Watch required checks and merge through protection. Monitor automatic Release through E2E and publication; confirm the remote annotated tag peels to the exact gated merged commit. Record final evidence and close Linear, using a protected documentation closeout if needed.

## Concrete Steps


Run from `/Users/pme/src/pmenglund/codex-sdk-go`. Initially this plan is the only intentional dirty path. Reclassify paths before each allow-dirty rerun.

    env -u CODEX_REPO_REF CODEX_REPO_ROOT=/Users/pme/src/openai/codex GOTOOLCHAIN=go1.26.9 CARGO_TARGET_DIR=/private/tmp/codex-01621-cargo .codex/skills/update-codex-protocol/scripts/update_codex_protocol.sh --allow-dirty

Add CODEX_PRINT_UNION_INVENTORY=1 or CODEX_PRINT_OPAQUE_INTERFACE_INVENTORY=1 when guards stop generation. Store logs under /private/tmp/codex-01621-review. Source baseline and target JSON are already extracted there for comparison.

    go test ./internal/codegen ./protocol ./rpc .
    CODEX_SCHEMA_TEST_DIR=/private/tmp/codex-01621-review/target/codex-rs/app-server-protocol/schema/json go test ./internal/codegen
    .github/scripts/validate-codex-metadata.sh
    .github/scripts/sdk-release-tag.sh
    .github/scripts/validate-release-tag.sh "$(.github/scripts/sdk-release-tag.sh)"
    git diff --check

The final updater must report two identical generations and complete local gate success, including formatting, metadata/installer fixtures, vet, unit/race tests, Staticcheck 0.8.1, and govulncheck 1.3.0. Verify generated headers all name 092d3acd6bec3e3a14bdc7e7a2810ab628ab759d.

## Validation and Acceptance


Wire regressions demonstrate omission and populated attribution/provenance fields, raw future error/tool payload preservation including integers above 2^53, known/future prediction variants, and new RPC routing. Existing user flows remain passing. All local gates and hosted checks pass, then trusted E2E and publication succeed. Remote v0.162.1 must be annotated and peel to the Release candidate on protected main.

## Idempotence and Recovery


Recheck status before mutations and preserve concurrent edits. Never stage everything, push main, manually publish or move tags, bypass failed gates, or approve pending environment jobs. If a newer stable tag appears before accepted generation, revise target, branch, filename, and all plan references. Fix failed Release through an exact-candidate retry or protected corrective merge retaining the unpublished version; dispatch with its exact SHA. After publication use a higher immutable corrective version and retract where appropriate.

## Artifacts and Notes


Official release API reports SHA-256 digests b293a0fcdd871900a44c0a4f40a1aa401b31e02854f790ebdb8bd5d1177fd1fe (Darwin arm64), d68e39d8f62639f33f7d00656554cb9c3ef7809e664a3827e62ab3ace3c21b21 (Darwin x86_64), bf3ca2719f1143b1440fbd85c149379e696c782c50ad7991f8bfcec60b2600c1 (Linux arm64 musl), and 86f268d81b898f3e144c5ecff2ad2fda2c5802e042d6195b72538fe0fa125057 (Linux x86_64 musl).

## Interfaces and Dependencies


Retain existing Go packages, generator dependencies, transport boundaries, and raw-preserving union conventions. Rust is exporter build tooling. Add no runtime dependency. Determine exact exported APIs from generated target schemas.

Revision note: Created after validated planner advice and live tracker/protection checks, before generation.

Revision note: Recorded failing wire regressions and successful manual source-schema coverage. Source extraction reproduces the approved 440-entry baseline digest exactly; target has 447 entries, 20 additions and 13 removals, digest 0a713f8e5cdc656a268e0a295111528b7cf7787cf8bf9e61c286f396c0bb310b. Review covered nullable browser extension, reasoning effort and goal-origin refs, MessagePhase partial_answer, open errors, additional_tools, prediction completed/failed, and enclosing request/notification containers. Printed exporter confirmation remains required before approval.

Revision note: Printed exporter exactly matched all 447 target entries. Recorded union approval before updating its digest. Installed Cargo 1.94.0 built the exact 0.162.1 exporter successfully without modifying upstream or installing tools.

Revision note: Reviewed and recorded the reduced opaque inventory before digest approval; no new opaque fallback was introduced. All generator tests against target source schemas pass.

Revision note: Recorded the govulncheck failure and validated planner amendment; updated only patch toolchain pins, retaining check names, analyzer versions, dependencies, and release protections. Final validation uses Go 1.26.9 and 1.27.2.

Revision note: Recorded the current Staticcheck export-format limitation and validated the minimal CI correction; native Go 1.27.2 runtime and vulnerability checks pass.

Revision note: Authoritative updater passed two identical generations and every local gate on Go 1.26.9, with no vulnerabilities found. Log: /private/tmp/codex-01621-review/patched-final-gate.log. All generated headers name the exact target commit; go.sum is unchanged. Work is ready for independent review before staging.

Revision note: QA and security reviews found no actionable issues. Architecture requested a current inventory-approval statement and removal of orphaned generated documentation; corrected the plan and added a failing regression proving the filter retains attached comments. The filter now removes contiguous generated GoDoc with a removed manual declaration while retaining unrelated comments and documentation. Regeneration and final gate rerun remain before staging.

Revision note: Final updater after review fixes passed in /private/tmp/codex-01621-review/review-final-gate.log. Reviewed the regenerated comment removals: all belong to filtered declarations; remaining declarations and their documentation are preserved. No inventory or protocol wire change was introduced by the cleanup.

Revision note: Architecture confirmed both fixes and no remaining findings. The changed generator also passed uncached Go 1.27.2 race tests. Release/tag/version-transition/content-manifest fixtures passed before named-path staging.
