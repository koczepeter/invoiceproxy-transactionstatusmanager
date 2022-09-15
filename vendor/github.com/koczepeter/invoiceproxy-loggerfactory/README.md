# invoiceproxy-loggerfactory

### RSYSLOG
docker pull peterkocze/invoiceproxy-syslog:1.0.3
docker run  --cap-add SYSLOG --restart always -v /Users/peterkocze/Repositories/go/src/koczepeter/invoiceproxy-loggerfactory/test/log:/var/log -p 514:514 -p 514:514/udp --name rsyslog mysyslog

