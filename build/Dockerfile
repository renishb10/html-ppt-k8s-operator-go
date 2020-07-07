FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/html-ppt-k8s-operator-go \
    USER_UID=1001 \
    USER_NAME=html-ppt-k8s-operator-go

# install operator binary
COPY build/_output/bin/html-ppt-k8s-operator-go ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
