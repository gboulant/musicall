all:

cmds.%:
	@make -C cmds/minimall $*
	@make -C cmds/plotwave $*
	@make -C cmds/wavechart $*
	@make -C cmds/examples $*
	@make -C cmds/letter2sound $*
	@make -C cmds/calibration $*
	@make -C cmds/playguitar $*

pkg.%:
	@make -C wave $*
	@make -C sound $*
	@make -C music $*
	@make -C guitar $*

test: pkg.test cmds.test
clean: pkg.clean cmds.clean
