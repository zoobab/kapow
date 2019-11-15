.PHONY: sync test

all: test

sync:
	pipenv sync --dev

test:
	KAPOW_DATAAPI_URL=http://localhost:8081 pipenv run make -C ../spec/test test

fix:
	KAPOW_DATAAPI_URL=http://localhost:8081 pipenv run make -C ../spec/test fix
