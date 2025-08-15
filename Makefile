all:

cmds.%:
	@make -C cmds/plotwave $*
	@make -C cmds/examples $*

test:
	@go test -C wave
	@go test -C sound
	@make cmds.test

clean:
	@find -name "output.*" | xargs rm -f
	@make cmds.clean