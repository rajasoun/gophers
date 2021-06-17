package controller

const S string = " \"Everyday Git in twenty commands or so\" \n \"git help everyday\" \n\n \"Show helpful guides that come with Git\" \n \"git help -g\" \n\n \"Search change by content\" \n \"git log -S'<a term in the source>'\" \n\n \"Show changes over time for specific file\" \n \"git log -p <file_name>\" \n\n \"Remove sensitive data from history, after a push\" \n \"git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch <path-to-your-file>' --prune-empty --tag-name-filter cat -- --all && git push origin --force --all\" \n\n \"Sync with remote, overwrite local changes\" \n \"git fetch origin && git reset --hard origin/master && git clean -f -d\" \n\n \"List of all files till a commit\" \n \"git ls-tree --name-only -r <commit-ish>\" \n\n \"Git reset first commit\" \n \"git update-ref -d HEAD\" \n\n \"Reset: preserve uncommitted local changes\" \n \"git reset --keep <commit>\" \n\n \"List all the conflicted files\" \n \"git diff --name-only --diff-filter=U\" \n\n \"List of all files changed in a commit\" \n \"git diff-tree --no-commit-id --name-only -r <commit-ish>\" \n\n \"Unstaged changes since last commit\" \n \"git diff\" \n\n \"Changes staged for commit\" \n \"git diff --cached\" \n\n \"Show both staged and unstaged changes\" \n \"git diff HEAD\" \n\n \"List all branches that are already merged into master\" \n \"git branch --merged master\" \n\n \"Quickly switch to the previous branch\" \n \"git checkout -\" \n\n \"Remove branches that have already been merged with master\" \n \"git branch --merged master | grep -v '^\\\\*' | xargs -n 1 git branch -d\" \n\n \"List all branches and their upstreams, as well as last commit on branch\" \n \"git branch -vv\" \n\n \"Track upstream branch\" \n \"git branch -u origin/mybranch\" \n\n \"Delete local branch\" \n \"git branch -d <local_branchname>\" \n\n \"Delete remote branch\" \n \"git push origin --delete <remote_branchname>\" \n\n \"Create local tag\" \n \"git tag <tag-name>\" \n\n \"Delete local tag\" \n \"git tag -d <tag-name>\" \n\n \"Delete remote tag\" \n \"git push origin :refs/tags/<tag-name>\" \n\n \"Undo local changes with the last content in head\" \n \"git checkout -- <file_name>\" \n\n \"Revert: Undo a commit by creating a new commit\" \n \"git revert <commit-ish>\" \n\n \"Reset: Discard commits, advised for private branch\" \n \"git reset <commit-ish>\" \n\n \"Reword the previous commit message\" \n \"git commit -v --amend\" \n\n \"See commit history for just the current branch\" \n \"git cherry -v master\" \n\n \"Amend author.\" \n \"git commit --amend --author='Author Name <email@address.com>'\" \n\n \"Reset author, after author has been changed in the global config.\" \n \"git commit --amend --reset-author --no-edit\" \n\n \"Changing a remote's URL\" \n \"git remote set-url origin <URL>\" \n\n \"Get list of all remote references\" \n \"git remote\" \n\n \"Get list of all local and remote branches\" \n \"git branch -a\" \n\n \"Get only remote branches\" \n \"git branch -r\" \n\n \"Stage parts of a changed file, instead of the entire file\" \n \"git add -p\" \n\n \"Get git bash completion\" \n \"curl -L http://git.io/vfhol > ~/.git-completion.bash && echo '[ -f ~/.git-completion.bash ] && . ~/.git-completion.bash' >> ~/.bashrc\" \n\n \"What changed since two weeks?\" \n \"git log --no-merges --raw --since='2 weeks ago'\" \n\n \"See all commits made since forking from master\" \n \"git log --no-merges --stat --reverse master..\" \n\n \"Pick commits across branches using cherry-pick\" \n \"git checkout <branch-name> && git cherry-pick <commit-ish>\" \n\n \"Find out branches containing commit-hash\" \n \"git branch -a --contains <commit-ish>\" \n\n \"Git Aliases\" \n \"git config --global alias.<handle> <command> \\ngit config --global alias.st status\" \n\n \"Saving current state of tracked files without commiting\" \n \"git stash\" \n\n \"Saving current state of unstaged changes to tracked files\" \n \"git stash -k\" \n\n \"Saving current state including untracked files\" \n \"git stash -u\" \n\n \"Saving current state with message\" \n \"git stash push -m <message>\" \n\n \"Saving current state of all files (ignored, untracked, and tracked)\" \n \"git stash -a\" \n\n \"Show list of all saved stashes\" \n \"git stash list\" \n\n \"Show the contents of any stash in patch form\" \n \"git stash show -p <stash@{n}>\" \n\n \"Apply any stash without deleting from the stashed list\" \n \"git stash apply <stash@{n}>\" \n\n \"Apply last stashed state and delete it from stashed list\" \n \"git stash pop\" \n\n \"Delete all stored stashes\" \n \"git stash clear\" \n\n \"Grab a single file from a stash\" \n \"git checkout <stash@{n}> -- <file_path>\" \n\n \"Show all tracked files\" \n \"git ls-files -t\" \n\n \"Show all untracked files\" \n \"git ls-files --others\" \n\n \"Show all ignored files\" \n \"git ls-files --others -i --exclude-standard\" \n\n \"Create new working tree from a repository (git 2.5)\" \n \"git worktree add -b <branch-name> <path> <start-point>\" \n\n \"Create new working tree from HEAD state\" \n \"git worktree add --detach <path> HEAD\" \n\n \"Untrack files without deleting\" \n \"git rm --cached <file_path>\" \n\n \"Before deleting untracked files/directory, do a dry run to get the list of these files/directories\" \n \"git clean -n\" \n\n \"Forcefully remove untracked files\" \n \"git clean -f\" \n\n \"Forcefully remove untracked directory\" \n \"git clean -f -d\" \n\n \"Update all the submodules\" \n \"git submodule foreach git pull\" \n\n \"Show all commits in the current branch yet to be merged to master\" \n \"git cherry -v master\" \n\n \"Rename a branch\" \n \"git branch -m <new-branch-name>\" \n\n \"Rebases 'feature' to 'master' and merges it in to master \" \n \"git rebase master feature && git checkout master && git merge -\" \n\n \"Archive the master branch\" \n \"git archive master --format=zip --output=master.zip\" \n\n \"Modify previous commit without modifying the commit message\" \n \"git add --all && git commit --amend --no-edit\" \n\n \"Prunes references to remove branches that have been deleted in the remote.\" \n \"git fetch -p\" \n\n \"Delete local branches that has been squash and merged in the remote.\" \n \"git branch -vv | grep ': gone]' | awk '{print $1}' | xargs git branch -D\" \n\n \"Retrieve the commit hash of the initial revision.\" \n \" git rev-list --reverse HEAD | head -1\" \n\n \"Visualize the version tree.\" \n \"git log --pretty=oneline --graph --decorate --all\" \n\n \"Visualize the tree including commits that are only referenced from reflogs\" \n \"git log --graph --decorate --oneline $(git rev-list --walk-reflogs --all)\" \n\n \"Deploying git tracked subfolder to gh-pages\" \n \"git subtree push --prefix subfolder_name origin gh-pages\" \n\n \"Adding a project to repo using subtree\" \n \"git subtree add --prefix=<directory_name>/<project_name> --squash git@github.com:<username>/<project_name>.git master\" \n\n \"Get latest changes in your repo for a linked project using subtree\" \n \"git subtree pull --prefix=<directory_name>/<project_name> --squash git@github.com:<username>/<project_name>.git master\" \n\n \"Export a branch with history to a file.\" \n \"git bundle create <file> <branch-name>\" \n\n \"Import from a bundle\" \n \"git clone repo.bundle <repo-dir> -b <branch-name>\" \n\n \"Get the name of current branch.\" \n \"git rev-parse --abbrev-ref HEAD\" \n\n \"Ignore one file on commit (e.g. Changelog).\" \n \"git update-index --assume-unchanged Changelog; git commit -a; git update-index --no-assume-unchanged Changelog\" \n\n \"Stash changes before rebasing\" \n \"git rebase --autostash\" \n\n \"Fetch pull request by ID to a local branch\" \n \"git fetch origin pull/<id>/head:<branch-name>\" \n\n \"Show the most recent tag on the current branch.\" \n \"git describe --tags --abbrev=0\" \n\n \"Show inline word diff.\" \n \"git diff --word-diff\" \n\n \"Show changes using common diff tools.\" \n \"git difftool [-t <tool>] <commit1> <commit2> <path>\" \n\n \"Don’t consider changes for tracked file.\" \n \"git update-index --assume-unchanged <file_name>\" \n\n \"Undo assume-unchanged.\" \n \"git update-index --no-assume-unchanged <file_name>\" \n\n \"Clean the files from .gitignore.\" \n \"git clean -X -f\" \n\n \"Restore deleted file.\" \n \"git checkout <deleting_commit> -- <file_path>\" \n\n \"Restore file to a specific commit-hash\" \n \"git checkout <commit-ish> -- <file_path>\" \n\n \"Always rebase instead of merge on pull.\" \n \"git config --global pull.rebase true\" \n\n \"List all the alias and configs.\" \n \"git config --list\" \n\n \"Make git case sensitive.\" \n \"git config --global core.ignorecase false\" \n\n \"Add custom editors.\" \n \"git config --global core.editor '$EDITOR'\" \n\n \"Auto correct typos.\" \n \"git config --global help.autocorrect 1\" \n\n \"Check if the change was a part of a release.\" \n \"git name-rev --name-only <SHA-1>\" \n\n \"Dry run. (any command that supports dry-run flag should do.)\" \n \"git clean -fd --dry-run\" \n\n \"Marks your commit as a fix of a previous commit.\" \n \"git commit --fixup <SHA-1>\" \n\n \"Squash fixup commits normal commits.\" \n \"git rebase -i --autosquash\" \n\n \"Skip staging area during commit.\" \n \"git commit --only <file_path>\" \n\n \"Interactive staging.\" \n \"git add -i\" \n\n \"List ignored files.\" \n \"git check-ignore *\" \n\n \"Status of ignored files.\" \n \"git status --ignored\" \n\n \"Commits in Branch1 that are not in Branch2\" \n \"git log Branch1 ^Branch2\" \n\n \"List n last commits\" \n \"git log -<n>\" \n\n \"Reuse recorded resolution, record and reuse previous conflicts resolutions.\" \n \"git config --global rerere.enabled 1\" \n\n \"Open all conflicted files in an editor.\" \n \"git diff --name-only | uniq | xargs $EDITOR\" \n\n \"Count unpacked number of objects and their disk consumption.\" \n \"git count-objects --human-readable\" \n\n \"Prune all unreachable objects from the object database.\" \n \"git gc --prune=now --aggressive\" \n\n \"Instantly browse your working repository in gitweb.\" \n \"git instaweb [--local] [--httpd=<httpd>] [--port=<port>] [--browser=<browser>]\" \n\n \"View the GPG signatures in the commit log\" \n \"git log --show-signature\" \n\n \"Remove entry in the global config.\" \n \"git config --global --unset <entry-name>\" \n\n \"Checkout a new branch without any history\" \n \"git checkout --orphan <branch_name>\" \n\n \"Extract file from another branch.\" \n \"git show <branch_name>:<file_name>\" \n\n \"List only the root and merge commits.\" \n \"git log --first-parent\" \n\n \"Change previous two commits with an interactive rebase.\" \n \"git rebase --interactive HEAD~2\" \n\n \"List all branch is WIP\" \n \"git checkout master && git branch --no-merged\" \n\n \"Find guilty with binary search\" \n \"git bisect start                    # Search start \\ngit bisect bad                      # Set point to bad commit \\ngit bisect good v2.6.13-rc2         # Set point to good commit|tag \\ngit bisect bad                      # Say current state is bad \\ngit bisect good                     # Say current state is good \\ngit bisect reset                    # Finish search \\n\" \n\n \"Bypass pre-commit and commit-msg githooks\" \n \"git commit --no-verify\" \n\n \"List commits and changes to a specific file (even through renaming)\" \n \"git log --follow -p -- <file_path>\" \n\n \"Clone a single branch\" \n \"git clone -b <branch-name> --single-branch https://github.com/user/repo.git\" \n\n \"Create and switch new branch\" \n \"git checkout -b <branch-name>\" \n\n \"Ignore file mode changes on commits\" \n \"git config core.fileMode false\" \n\n \"Turn off git colored terminal output\" \n \"git config --global color.ui false\" \n\n \"Specific color settings\" \n \"git config --global <specific command e.g branch, diff> <true, false or always>\" \n\n \"Show all local branches ordered by recent commits\" \n \"git for-each-ref --sort=-committerdate --format='%(refname:short)' refs/heads/\" \n\n \"Find lines matching the pattern (regex or string) in tracked files\" \n \"git grep --heading --line-number 'foo bar'\" \n\n \"Clone a shallow copy of a repository\" \n \"git clone https://github.com/user/repo.git --depth 1\" \n\n \"Search Commit log across all branches for given text\" \n \"git log --all --grep='<given-text>'\" \n\n \"Get first commit in a branch (from master)\" \n \"git log --oneline master..<branch-name> | tail -1\" \n\n \"Unstaging Staged file\" \n \"git reset HEAD <file-name>\" \n\n \"Force push to Remote Repository\" \n \"git push -f <remote-name> <branch-name>\" \n\n \"Adding Remote name\" \n \"git remote add <remote-nickname> <remote-url>\" \n\n \"List all currently configured remotes\" \n \"git remote -v\" \n\n \"Show the author, time and last revision made to each line of a given file\" \n \"git blame <file-name>\" \n\n \"Group commits by authors and title\" \n \"git shortlog\" \n\n \"Forced push but still ensure you don't overwrite other's work\" \n \"git push --force-with-lease <remote-name> <branch-name>\" \n\n \"Show how many lines does an author contribute\" \n \"git log --author='_Your_Name_Here_' --pretty=tformat: --numstat | gawk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf \\\"added lines: %s removed lines: %s total lines: %s\\n\\\", add, subs, loc }' -\" \n\n \"Revert: Reverting an entire merge\" \n \"git revert -m 1 <commit-ish>\" \n\n \"Number of commits in a branch\" \n \"git rev-list --count <branch-name>\" \n\n \"Alias: git undo\" \n \"git config --global alias.undo '!f() { git reset --hard $(git rev-parse --abbrev-ref HEAD)@{${1-1}}; }; f'\" \n\n \"Add object notes\" \n \"git notes add -m 'Note on the previous commit....'\" \n\n \"Show all the git-notes\" \n \"git log --show-notes='*'\" \n\n \"Apply commit from another repository\" \n \"git --git-dir=<source-dir>/.git format-patch -k -1 --stdout <SHA1> | git am -3 -k\" \n\n \"Specific fetch reference\" \n \"git fetch origin master:refs/remotes/origin/mymaster\" \n\n \"Find common ancestor of two branches\" \n \"git merge-base <branch-name> <other-branch-name>\" \n\n \"List unpushed git commits\" \n \"git log --branches --not --remotes\" \n\n \"Add everything, but whitespace changes\" \n \"git diff --ignore-all-space | git apply --cached\" \n\n \"Edit [local/global] git config\" \n \"git config [--global] --edit\" \n\n \"blame on certain range\" \n \"git blame -L <start>,<end>\" \n\n \"Show a Git logical variable.\" \n \"git var -l | <variable>\" \n\n \"Preformatted patch file.\" \n \"git format-patch -M upstream..topic\" \n\n \"Get the repo name.\" \n \"git rev-parse --show-toplevel\" \n\n \"logs between date range\" \n \"git log --since='FEB 1 2017' --until='FEB 14 2017'\" \n\n \"Exclude author from logs\" \n \"git log --perl-regexp --author='^((?!excluded-author-regex).*)$'\" \n\n \"Generates a summary of pending changes\" \n \"git request-pull v1.0 https://git.ko.xz/project master:for-linus\" \n\n \"List references in a remote repository\" \n \"git ls-remote git://git.kernel.org/pub/scm/git/git.git\" \n\n \"Backup untracked files.\" \n \"git ls-files --others -i --exclude-standard | xargs zip untracked.zip\" \n\n \"List all git aliases\" \n \"git config -l | grep alias | sed 's/^alias\\\\.//g'\" \n\n \"Show git status short\" \n \"git status --short --branch\" \n\n \"Checkout a commit prior to a day ago\" \n \"git checkout master@{yesterday}\" \n\n \"Push the current branch to the same name on the remote repository\" \n \"git push origin HEAD\" \n\n \"Push a new local branch to remote repository and track\" \n \"git push -u origin <branch_name>\" \n\n \"Change a branch base\" \n \"git rebase --onto <new_base> <old_base>\" \n\n \"Use SSH instead of HTTPs for remotes\" \n \"git config --global url.'git@github.com:'.insteadOf 'https://github.com/'\" \n\n \"Update a submodule to the latest commit\" \n \"cd <path-to-submodule>\\ngit pull origin <branch>\\ncd <root-of-your-main-project>\\ngit add <path-to-submodule>\\ngit commit -m \\\"submodule updated\\\"\" \n\n \"Prevent auto replacing LF with CRLF\" \n \"git config --global core.autocrlf false\" \n\n"
