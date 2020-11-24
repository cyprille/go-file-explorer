install:
	@if [ -f ".env" ]; then\
		echo "\033[33mSkipping copy:\033[0m The file \033[36m.env\033[0m \033[33malready exists\033[0m";\
	else\
		cp .env.dist .env ||:;\
		echo "\033[32mSuccessful:\033[0m Created \033[36m.env\033[0m from \033[36m.env.dist\033[0m";\
	fi

	go get -u ./...
