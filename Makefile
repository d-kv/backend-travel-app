install-tools:
	npm install


mmd_src := add_poi get_route remove_poi find_pois update_poi

.SILENT: generate-uml
generate-uml:
	for file in $(mmd_src); do \
		npx -p @mermaid-js/mermaid-cli mmdc -i uml/src/$$file.mmd -o uml/gen/$$file.svg --cssFile uml/src/style.css;\
	done
