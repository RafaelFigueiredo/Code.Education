FROM golang:1.15.8 as build

WORKDIR /build

COPY main.go .

RUN go build -o codeeducation


FROM scratch as release
COPY --from=build /build/codeeducation /codeeducation

ENTRYPOINT ["/codeeducation"]
