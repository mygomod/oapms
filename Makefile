MUSES_SYSTEM:=github.com/mygomod/muses/pkg/system
APPNAME:=oapms
APPPATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
APPOUT:=$(APPPATH)/appgo/$(APPNAME)


ant:
	@cd $(APPPATH)/webui && npm start

# 执行go指令
bee:
	@bee pro migration --sqlmode=down --sqlpath=scripts/sql
	@bee pro gen

prod.deploy.webui:
	@cd $(APPPATH)/webui && npm run build
	@cd $(APPPATH)/webui && $(APPPATH)/scripts/deploy/deploy_webui.sh $(APPNAME) root@xx-01

prod.deploy.go:
	@cd $(APPPATH) && go build
	@cd $(APPPATH) && $(APPPATH)/scripts/deploy/deploy_go.sh $(APPNAME) root@xx-01

prepub.deploy.webui:
	@cd $(APPPATH)/webui && npm run build
	@cd $(APPPATH)/webui && $(APPPATH)/scripts/deploy/deploy_webui.sh $(APPNAME) askuy@askuy

prepub.deploy.go:
	@cd $(APPPATH) && go build
	@cd $(APPPATH) && $(APPPATH)/scripts/deploy/deploy_go.sh $(APPNAME) askuy@askuy

go:
	@cd $(APPPATH) && go run main.go start --conf=conf/conf-dev.toml
clear:
	@bee pro migration --sqlmode=down --sqlpath=scripts/sql
install:
	@cd $(APPPATH) && go run main.go install --conf=conf/conf-dev.toml --mode=install
mock:
	@cd $(APPPATH) && go run main.go install --conf=conf/conf-dev.toml --mode=mock
debug: clear install mock
