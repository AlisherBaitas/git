# Git workflow

## Remotes

The repository has two remotes. `origin` points to Gitea (`https://01.tomorrow-school.ai/git/abaitas/git.git`) and `github` points to GitHub (`https://github.com/AlisherBaitas/git`). Every branch is pushed to both of them:

```bash
git push origin <branch>
git push github <branch>
```

## Branching

`main` is the default branch and nothing is committed to it directly. Every task gets its own branch, created with `git switch -c <name>`, and is merged back into `main` when it is done.

| Branch | Purpose |
|---|---|
| `feature` | add `feature/feature.txt` |
| `readme-update` | change the README line used for the merge conflict exercise |
| `docs-workflow` | add this file |
| `analyzer` | add the Go Repo Health Analyzer, its tests and `.gitignore` |
| `docs-final` | fill in the documentation, README and reflection |

## Commit messages

Every commit message follows the convention `type: short description`, where `type` is one of `add`, `update`, `fix`, `delete`, `docs`, `refactor`. The description says what changed and names the file or topic, for example `add: feature.txt on the feature branch`. Each commit holds one logical change.

## Merging

Branches are merged with `git merge --no-ff <branch>`. Without `--no-ff`, git does a fast-forward when `main` has no new commits: it only moves the `main` pointer forward, no merge commit is created and the branch disappears from the history. With `--no-ff`, git always creates a merge commit with two parents, so the graph shows where each branch started and where it was merged:

```
*   f8d341a fix: resolve README merge conflict between main and readme-update
|\
| * ff82ae1 update: README description on readme-update branch
* | 9fad9cc update: README description on main branch
|/
*   3d5ffe6 update: merge feature branch with feature.txt into main
|\
| * 75d589b add: feature.txt on the feature branch
|/
* 2fc9bb2 add: README with project description
```

During the exercise one merge happened as a fast-forward by mistake, because the change on `main` had not been saved. It was undone with `git reset --hard 3d5ffe6`, which moved `main` back before the merge.

## Merge conflict

The conflict was created on purpose. The same line of `README.md` was changed in two different ways: once on `readme-update` and once on `main`. Running `git merge readme-update` on `main` stopped with `CONFLICT (content): Merge conflict in README.md`, and git wrote both versions into the file:

```
<<<<<<< HEAD
Learning repository for git basics. Changed on the main branch.
=======
Learning repository for git basics. Changed on the readme-update branch.
>>>>>>> readme-update
```

- `<<<<<<< HEAD` starts the version from the current branch (`main`).
- `=======` separates the two versions.
- `>>>>>>> readme-update` ends the version from the branch being merged.

To resolve it, the file was edited to keep one final line and all three markers were removed. Then `git add README.md` marked the conflict as resolved and `git commit` finished the merge. `git merge --abort` would cancel the merge and return to the state before it.

## Commands I used

| Command | What it does |
|---|---|
| `git clone <url>` | copy a remote repository to the computer |
| `git config user.name / user.email` | set the author of commits |
| `git status` | show changed, staged and untracked files |
| `git diff` | show changes that are not staged yet |
| `git add <file>` | put changes into the staging area |
| `git commit -m "..."` | save the staged changes as a commit |
| `git log --oneline --graph --all` | show the history of all branches as a graph |
| `git remote add <name> <url>` | connect another remote |
| `git remote -v` | list remotes with their URLs |
| `git push -u <remote> <branch>` | send a branch to a remote and remember the link |
| `git switch -c <branch>` | create a branch and switch to it |
| `git switch <branch>` | switch to an existing branch |
| `git merge --no-ff <branch>` | merge a branch into the current one with a merge commit |
| `git merge --abort` | cancel a merge with conflicts |
| `git reset --hard <commit>` | move the current branch to a commit and reset the files |
| `git reflog` | show every move of HEAD, useful to find lost commits |
