all:

test:
	@go test -C wave
	@go test -C sound
	@make -C cmds/graphwave test
	@make -C cmds/examples test

clean:
	@make -C cmds/graphwave clean
	@make -C cmds/examples clean
