day = $(shell date +'%-d')
year = $(shell date +'%-Y')

new:
	@echo "Creating new file structure for day" $(day)"..."

	@if [ $(day) -lt 10 ] ; then \
  		mkdir -p $(year)/day-0$(day); \
  		cp template $(year)/day-0$(day)/day0$(day).go; \
  	else \
  		mkdir -p $(year)/day-$(day); \
		cp template $(year)/day-$(day)/day$(day).go; \
    fi
	$(shell go run input/main.go)
	@echo "Files successfully created.. happy hacking :)"
	@git add $(year)/
