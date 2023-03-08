#
#	for file in files:
# 		if file.svg isn't in the repo && file.mmd has been changed:
# 			mermaid.generate(file)
#

for file in "$@"; do
    if !(test -f uml/gen/${file}.svg) && (git diff --exit-code --quiet uml/src/${file}.mmd)
    then
        npx -p @mermaid-js/mermaid-cli mmdc -i uml/src/${file}.mmd -o uml/gen/${file}.svg --cssFile uml/src/style.css
    fi
done
