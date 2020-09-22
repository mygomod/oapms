#!/bin/bash
rm -rf bin.tar.gz
tar zcvf bin.tar.gz conf ${1} scripts data go.mod .beegopro.timestamp
echo "参数1：$1"
echo "参数2：$2"
ssh ${2} "cd /home/www/server/${1} && mv bin.tar.gz binbak.tar.gz"
scp bin.tar.gz ${2}:/home/www/server/${1}/

ssh ${2} "cd /home/www/server/${1} && \
	    tar xvf bin.tar.gz && \
	    chown www:www -R /home/www/server/${1} && \
	    supervisorctl restart ${1}"
