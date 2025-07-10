all:

test:
	@go test -C wave
	@go test -C sound
	@make -C cmds/plotwave test
	@make -C cmds/examples test

clean:
	@make -C cmds/plotwave clean
	@make -C cmds/examples clean
