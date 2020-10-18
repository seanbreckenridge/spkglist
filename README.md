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
go get gitlab.com/seanbreckenridge/spkglist/cmd/spkglist
```

```
A simple package list parser
Pass one or more files to read from as arguments
If none provided, reads from STDIN.
  -delimiter string
    	delimiter to print between results (default "\n")
  -json
    	Print results as a JSON array
  -print0
    	use the null character as the delimiter
  -split
    	split on all whitespace, instead of newlines

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

```conf
go  # just a package name

package name # assumes you want 'package name', you can pass the -split flag otherwise

# this would be an error, not allowed as a 'bare' character
😀

# you can quote any line with backticks, to include any charater, including a '#'
`github.com/user/emoji_package😀#master`  # install from github

# python/node-esque packages work fine as bare words
requests==2.24.0
@elm-tooling/elm-language-server
serve
```

Output:

```
line 6:0 token recognition error at: '😀'
go
package name
github.com/user/emoji_package😀#master
requests==2.24.0
@elm-tooling/elm-language-server
serve
```

