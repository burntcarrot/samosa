<div align="center">
    <img src = "assets/samosa.svg" width=400>
</div>

---

# Samosa

Samosa helps developers prioritize what needs to be tested.

> ⚠️ Do not chase coverage metrics. When a measure becomes a target, it ceases to be a good measure.

Samosa isn't a tool for developers who want to chase 100% coverage; instead, it provides a good way to prioritize which functions should be tested first.

## Usage

Run Samosa:

```
samosa -f <path to coverage file>
```

Samosa returns the list of functions sorted by the impact associated with covering it:

```
File                                                    Function                       Impact   Uncovered Lines Start Line      End Line
...ub.com/deepsourcelabs/cli/utils/remote_resolver.go   ResolveRemote                  6.18     32              14              71
...com/deepsourcelabs/cli/command/issues/list/list.go   Run                            4.44     23              76              118
...com/deepsourcelabs/cli/command/issues/list/list.go   getIssuesData                  4.44     23              122             172
github.com/deepsourcelabs/cli/config/config.go          WriteFile                      3.28     17              105             136
...com/deepsourcelabs/cli/command/issues/list/list.go   NewCmdIssuesList               2.32     12              35              73
...ithub.com/deepsourcelabs/cli/utils/fetch_remote.go   ListRemotes                    2.32     12              74              99
...urcelabs/cli/utils/fetch_analyzers_transformers.go   GetAnalyzersAndTransformersData1.93     10              33              49
...urcelabs/cli/utils/fetch_analyzers_transformers.go   parseSDKResponse               1.93     10              53              80
...com/deepsourcelabs/cli/command/issues/list/list.go   showIssues                     1.93     10              176             193
```

## Status

Samosa is in development. New features will be added soon.
