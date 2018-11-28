from 192.168.138.180:9000/centos/centos:latest

ARG DAPPER_HOST_ARCH
ENV HOST_ARCH=${DAPPER_HOST_ARCH} ARCH=${DAPPER_HOST_ARCH}

RUN yum install -y gcc ca-certificates git wget curl vim less file && \
    rm -f /bin/sh && ln -s /bin/bash /bin/sh


ADD ccr /opt/

CMD ["/opt/ccr","-alsologtostderr=true"] 
