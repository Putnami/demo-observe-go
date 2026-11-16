# AGENTS.md

<!-- putnami:guidance v2 begin sha256:477e24673041f499163a45fab57c93e8ef5b9f7080f1c774cb7373e07d8d4149 -->
This is a Putnami workspace.
- Build, test, lint, install, add dependencies, and generate with `./putnamiw` when present, otherwise `putnami`. Never call npm, bun, go, pip, or cargo directly for those.
- While iterating, select the projects you changed with `--projects <a>,<b>`; unlike `--impacted`, it skips every project that depends on them. Before declaring work complete, run `putnami lint,test,build --impacted --enforce-coverage` once, when feasible.
- Discover code through the CLI, then grep: `putnami projects list|describe|owner|deps|why` for the project graph, `putnami context map|pack` to orient on the workspace or a project, and `putnami <verb> --impacted --dry-run` for what a change affects. Use grep for literal text.
- Skills and agents that an extension's agent content installs under `.agents/`, `.claude/` and `.codex/` are managed by Putnami. To customize one, copy it under another name; edited managed files are preserved and reported, never overwritten.
- Read `.agents/constraints.md` when it exists.
<!-- putnami:guidance v2 end -->
