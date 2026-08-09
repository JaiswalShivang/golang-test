# flake-agent Go demo repo

A minimal Go repo demonstrating flake-agent end to end against a real `go test -json` pipeline, using Ollama Cloud for analysis.

## Setup

1. Replace `YOUR-GITHUB-USERNAME` in both workflow files with your actual GitHub username/org.
2. Create an API key at https://ollama.com/settings/keys and add it as `OLLAMA_API_KEY` in this repo's Settings → Secrets and variables → Actions.
3. Open a PR. `TestFlakyNetworkCall` fails roughly 30% of the time by design.

## Note on Go test caching

`go test` caches results per package when source has not changed. Always use `-count=1` to force a real re-run, already included in this workflow's `Run tests` and `Rerun failing test locally` steps.

## What to look for

- Actions tab: the `Flake triage` step should run and end with `PR comment posted: True`.
- The PR: a comment starting with `### Flake triage: ...`.
- Issues tab: a tracking Issue once the flake count threshold is reached.
- The `flake-data` branch: confirms history persistence is working across runs.
