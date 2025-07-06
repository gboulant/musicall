all:

test:
	@go test -C wave
	@go test -C sound
	@make -C cmds/graphwave test

clean:
	@make -C cmds/graphwave clean