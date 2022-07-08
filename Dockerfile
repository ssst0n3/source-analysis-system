FROM golang:1.17-alpine AS backend
#ENV GOPROXY="https://proxy.golang.org"
ENV GOPROXY="https://goproxy.io,direct"
COPY backend /backend
WORKDIR /backend
RUN GO111MODULE="on" GOPROXY=$GOPROXY go mod tidy
RUN GO111MODULE="on" GOPROXY=$GOPROXY CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"
RUN sed -i "s@https://dl-cdn.alpinelinux.org/@https://mirrors.huaweicloud.com/@g" /etc/apk/repositories
RUN apk update && apk add upx
RUN upx source-analysis-system

FROM node:14-slim AS frontend
ENV NPM_REGISTRY https://mirrors.huaweicloud.com/repository/npm/
ENV NPM_REGISTRY https://registry.npm.taobao.org
COPY frontend /frontend
WORKDIR /frontend
RUN npm config set registry $NPM_REGISTRY && \
npm cache clean -f && \
npm install
RUN npm run-script build

FROM alpine
RUN mkdir -p /app
COPY --from=backend /backend/source-analysis-system /app/
COPY --from=frontend /frontend/dist /app/dist
WORKDIR /app
CMD ./source-analysis-system
#CMD ls -lah db/