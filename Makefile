MUSES_SYSTEM:=github.com/mygomod/muses/pkg/system
APPNAME:=oapms
APPPATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
APPOUT:=$(APPPATH)/appgo/$(APPNAME)


ant:
	@cd $(APPPATH)/webui && npm start

# 执行go指令
go:
	@cd $(APPPATH) && go run main.go start --conf=conf/conf.toml

bee:
	@bee pro migration --sqlmode=down --sqlpath=scripts/sql
	@bee pro gen

install:
	@bee pro migration --sqlmode=down --sqlpath=scripts/sql
	@bee pro migration --sqlmode=up --sqlpath=scripts/sql
	@cd $(APPPATH) && go run main.go install

prod.deploy.webui:
	@cd $(APPPATH)/webui && npm run build
	@cd $(APPPATH)/webui && $(APPPATH)/scripts/deploy/deploy_webui.sh $(APPNAME) xx-01

prod.deploy.go:
	@cd $(APPPATH) && go build
	@cd $(APPPATH) && $(APPPATH)/scripts/deploy/deploy_go.sh $(APPNAME) xx-01
