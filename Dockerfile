#FROM --platform=linux/amd64 golang:1.19.2-alpine as builder
#RUN apk add make gcc musl-dev binutils-gold
#
#WORKDIR  /myproject-plugin
#COPY ./mobiquityPlugin/ ./
#RUN  go build -buildmode=plugin -o test-plugin.so .
FROM devopsfaith/krakend:2.6.2

WORKDIR /etc/krakend/
#FROM devopsfaith/krakend:2.6.2 as stage
#
ARG ENV
ENV ENV_NAME $ENV

COPY ./config/krakend ./
# Save temporary file to /tmp to avoid permission errors

RUN FC_ENABLE=1 \
    FC_OUT=/tmp/krakend.json \
    FC_SETTINGS="/etc/krakend/settings/$ENV_NAME" \
    FC_PARTIALS="/etc/krakend/partials" \
    FC_TEMPLATES="/etc/krakend/templates" \
    krakend check -d -t -c ./krakend.tmpl


##
#### The linting needs the final krakend.json file
RUN krakend check -c /tmp/krakend.json --lint
CMD ["sleep", "infinity"]
#
#RUN rm -r ./*
#RUN cp ./tmp/krakend.json ./etc/krakend/

#COPY --from=stage /tmp/krakend.json .
#COPY ./config/krakend/krakend.json ./
#COPY --from=builder /myproject-plugin/test-plugin.so /etc/krakend/plugins/
