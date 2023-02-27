install-tools:
	npm install


mmd_src := add_poi get_poi remove_poi search_poi update_poi
md_src := legend

.SILENT: generate-uml
generate-uml:
	for file in $(mmd_src); do \
		npx -p @mermaid-js/mermaid-cli mmdc -i uml/src/$$file.mmd -o uml/gen/$$file.svg --cssFile uml/src/style.css;\
	done
	npx -p @mermaid-js/mermaid-cli mmdc -i uml/src/legend.md -o uml/gen/legend.md --cssFile uml/src/style.css;\
