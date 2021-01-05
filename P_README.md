## spkglist

A simple package list parser.

This isn't for any language/tool in particular, its meant to clean up hand written package files for any tool, especially while creating scripts to bootstrap installation processes.

For example, if you had a list of system packages you wanted to install:

`packages.txt`:

```conf
# packages I want to install with my package manager
java  # ....
go
  python3
```

This removes comments and cleans up whitespace/newlines, so that that doesn't have to be done in the shell.

`$ spkglist packages.txt`

```conf
java
go
python3
```

### Install

```
go get github.com/seanbreckenridge/spkglist/cmd/spkglist
```

```
>>>PMARK
perl -E 'print "`"x3, "\n"'
go build ./cmd/spkglist/
./spkglist -help
perl -E 'print "`"x3, "\n"'
```

### Common Usage

To pass the parsed information to a command, one could do:

```
spkglist -print0 packages.txt | xargs -0 sudo apt install
```

Some `bash` that might be used to prompt the user to install packages using this:

```bash
#!/bin/bash
# for complications with prompting while looping, see
# https://stackoverflow.com/q/6883363/9348376
# https://github.com/koalaman/shellcheck/wiki/SC2013
# http://mywiki.wooledge.org/DontReadLinesWithFor
{ spkglist packages.txt | while IFS= read -r pkg; do
		# prompt user
		printf "install '%s' [Y/n]" "$pkg"
    read -u 3 -n 1
		if [[ $REPLY =~ [Nn] ]]; then
			# skip to next package if user answered no
			printf '\n'
			continue
		else
			printf "\ninstalling '%s'\n" "$pkg"
		fi
done; } 3<&0
```

## Specification

```
>>>PMARK
perl -E 'print "`"x3, "conf", "\n"'
cat ./examples/spec.txt
perl -E 'print "`"x3, "\n"'
```

Output:

```
>>>PMARK
perl -E 'print "`"x3, "\n"'
go build ./cmd/spkglist/
./spkglist ./examples/spec.txt 2>/dev/null
perl -E 'print "\n", "`"x3, "\n"'
```

---

For more examples, you can see my usage in my system bootstrap script [here](https://github.com/seanbreckenridge/dotfiles/blob/44ce8023b0cc517bd014d45b7051e38aa2d7c463/.config/yadm/bootstrap#L119-L202), corresponding package lists [here](https://github.com/seanbreckenridge/dotfiles/tree/8fba977a478b7b3307109e4e72974600a6beed16/.config/yadm/package_lists).
