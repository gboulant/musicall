all:

cmds.%:
	@make -C cmds/plotwave $*
	@make -C cmds/wavechart $*
	@make -C cmds/examples $*

pkg.%:
	@make -C wave $*
	@make -C sound $*
	@make -C music $*

test: pkg.test cmds.test
clean: pkg.clean cmds.clean