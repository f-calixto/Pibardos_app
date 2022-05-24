service_structure:
	mkdir -p $(service)_service/pkg/$(service) && touch $(service)_service/pkg/$(service)/repository.go && touch $(service)_service/pkg/$(service)/service.go && touch $(service)_service/pkg/$(service)/handlers.go && touch $(service)_service/pkg/$(service)/middleware.go && touch $(service)_service/pkg/$(service)/models.go && touch $(service)_service/pkg/$(service)/router.go
	cp group_service/main.go $(service)_service/main.go
	mkdir $(service)_service/pkg/errors
	cp -r group_service/pkg/log $(service)_service/pkg/log
	mkdir $(service)_service/pkg/utils
	touch $(service)_service/Dockerfile
	cd $(service)_service && go mod init github.com/coding-kiko/$(service)_service
