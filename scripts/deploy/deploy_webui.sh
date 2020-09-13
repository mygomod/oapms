#!/bin/bash
rm -rf dist.tar.gz
tar zcvf dist.tar.gz dist

ssh root@${2} "cd /home/www/server/${1}/webui && mv dist.tar.gz distbak.tar.gz"
scp dist.tar.gz root@${2}:/home/www/server/${1}/webui/

ssh root@${2} "cd /home/www/server/${1}/webui/ && \
	    tar xvf dist.tar.gz && \
	    chown www:www -R /home/www/server/${1}/webui/"
