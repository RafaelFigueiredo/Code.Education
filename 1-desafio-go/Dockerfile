FROM golang:1.15.8 as build

RUN apt-get update && apt-get install -y upx

WORKDIR /build

COPY main.go .

RUN go build -ldflags="-s -w" -o codeeducation
RUN upx codeeducation

FROM scratch as release
COPY --from=build /build/codeeducation /codeeducation

ENTRYPOINT ["/codeeducation"]
