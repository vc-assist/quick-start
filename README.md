# Quick start

> This repository is meant as a quick start for having a working development copy of VC Assist.

> [!NOTE]
> It is NOT recommended to try and make this a starting point for your *understanding* of the VC Assist codebase, instead take your time going through an individual repo, running/testing it individually, and then come back to this repository.

## Prerequisites

- [Node.js](https://nodejs.org/en)
- [Golang](https://go.dev/)
- `pnpm add -g mprocs`
- An email account and app password provided to `backend/cmd/auth/config.json5`.

## Usage

```sh
go run ./cmd/setup
mprocs -c ./cmd/dev_vcs.yaml

# when you want to update
go run ./cmd/update
```

## Commands

- `go run ./cmd/setup` - sets up dependencies/prereqs for all the repositories.
- `mprocs -c ./cmd/dev_vcs.yaml` - runs a development environment.

## Why manually clone all the repositories with a script instead of using submodules?

Submodules come with a bunch of footguns that can be mostly avoided by treating them as immutable. Therefore we clone each repo by itself so you can make changes to the repo itself instead of its instance as a submodule in another repo.

This doesn't feel the greatest (especially when you're making many changes across many repos and you have to pollute git history with a bunch of potentially broken commits), but it is a lot better than having to deal with submodule footguns or trying to understand a gigantic monorepo all at once.

