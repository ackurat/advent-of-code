day = $(shell date +'%-d')
year = $(shell date +'%-Y')

new:
	@echo "Creating new file structure for day" $(day)"..."

	@if [ $(day) -lt 10 ] ; then \
  		mkdir -p $(year)/day0$(day); \
  		cp template $(year)/day0$(day)/main.go; \
		touch $(year)/day0$(day)/example.txt; \
  	else \
  		mkdir -p $(year)/day$(day); \
		cp template $(year)/day$(day)/main.go; \
		touch $(year)/day0$(day)/example.txt; \
    fi
	$(shell go run input/main.go)
	@echo "Files successfully created.. happy hacking :)"
	@git add $(year)/
