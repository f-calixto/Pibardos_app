service_structure:
	mkdir -p $(service)_service/pkg/$(service) && touch $(service)_service/pkg/$(service)/repository.go && touch $(service)_service/pkg/$(service)/service.go && touch $(service)_service/pkg/$(service)/transport.go && touch $(service)_service/pkg/$(service)/middleware.go
	mkdir $(service)_service/cmd && touch $(service)_service/cmd/main.go
	mkdir $(service)_service/pkg/errors
	mkdir $(service)_service/pkg/log
	mkdir $(service)_service/pkg/utils
	touch $(service)_service/Dockerfile
	cd $(service)_service && go mod init github.com/coding-kiko/$(service)_service
