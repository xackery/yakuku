NAME := yakuku
SHELL := /bin/bash
VERSION ?= 0.0.8

run-sql-%:
	make build
	mkdir -p bin/out
	cd bin && ./${NAME} sql $* $*.yaml out/$*.sql	

run-inject-%:
	make build
	mkdir -p bin/out
	cd bin && ./${NAME} inject out/$*.sql

run-yaml-%:
	make build
	mkdir -p bin/out
	@#cd bin && ./${NAME} yaml $* $*.yaml 9979 9990 9991 32601 32601
	@cd bin && ./${NAME} yaml $* $*.yaml 6378 9979 9989 9990 9991 9997 9998 9999 15026 15041 15093 15200 15202 15288 15343 15372 18017 18018 18019 18203 18204 18205 18206 18207 18363 18364 18365 18366 18367 18431 18432 18433 18434 18551 18552 18553 18554 18702 18703 18704 18705 18706 18707 18708 18709 18710 18711 18712 18713 18714 18715 18716 18717 18718 18719 18720 18721 18723 18724 18725 18726 18727 18728 18729 18731 18732 18733 18734 18735 18736 18737 18738 18739 18740 18741 18742 18743 18744 18745 18746 18747 18748 18751 18752 18753 18754 18755 18756 18757 18758 18759 18760 18761 18762 18765 18766 18767 18768 18769 18770 18771 18772 18773 18774 18775 18776 18777 18778 18779 18780 18781 18782 18783 18784 18785 18786 18787 18788 18789 18790 18791 18792 18845 18846 18847 18848 18849 18850 18851 18852 18853 18854 18855 18856 18857 21779 32601 36000 36001 36002 36003 36004 51634 51635 51636 51637 55623 59892
	@#cd bin && ./${NAME} yaml $* $*.yaml 21


# compilies the project to bin/
build:
	mkdir -p bin
	go build -o bin/${NAME} main.go
# runs all build commands
build-all: build-linux build-windows
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${NAME} -ldflags="-X main.Version=${VERSION} -s -w" main.go
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/${NAME}.exe -ldflags="-X main.Version=${VERSION} -s -w" main.go
dump-item-%:
	cd bin && ./${NAME} import item $*
show-tables: copy-data
	source .env && cd bin && docker run --rm \
	-v ${PWD}:/src \
    imega/mysql-client \
    mysql --host=$$DB_HOST --user=$$DB_USER --password=$$DB_PASS --database=$$DB_NAME \
	--execute='show tables;'
# inject tries to insert the sql files from bin/ to the db
inject: copy-data
	docker run -it --rm \
	-v ${PWD}/bin:/src \
    imega/mysql-client \
	/bin/sh -c 'source /src/.env && mysql --host=$$DB_HOST --user=$$DB_USER --password=$$DB_PASS $$DB_NAME < /src/aa.sql'

# CICD triggers this
set-version-%:
	@echo "VERSION=${VERSION}.$*" >> $$GITHUB_ENV

	