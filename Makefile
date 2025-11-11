all:

demos.%:
	@make -C demos/d01.minimall $*
	@make -C demos/d02.waveplot $*
	@make -C demos/d03.wavechart $*
	@make -C demos/d04.wavesound $*
	@make -C demos/d05.wavefft $*	
	@make -C demos/d06.letter2sound $*
	@make -C demos/d07.calibration $*
	@make -C demos/d10.playguitar $*
	@make -C demos/d11.guitarneck $*

pkg.%:
	@make -C wave $*
	@make -C sound $*
	@make -C music $*
	@make -C guitar $*

test: pkg.test demos.test
clean: pkg.clean demos.clean
