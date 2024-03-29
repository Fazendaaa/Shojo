FROM golang:1.16.14-alpine3.15 as base

ENV CGO_ENABLED=0
WORKDIR /go/src/github.com/Fazendaaa/Shojo

COPY go.* .
COPY main.go .

COPY cmd cmd/
COPY controllers controllers/
COPY pkg pkg/
COPY test test/

RUN [ "go", "test", "./..." ]
RUN [ "go", "build" ]
RUN [ "mv", "Shojo", "/usr/bin/shojo" ]

FROM alpine:3.15 as runner
LABEL author="fazenda"
LABEL project="shojo-latex"

COPY --from=base /usr/bin/shojo /usr/bin/shojo

RUN [ "apk", "add", "--no-cache", \
  "texlive-xetex", \
  "texmf-dist-latexextra", \
  "texlive-luatex" \
]
RUN [ "apk", "add", "--no-cache", \
  "--repository=http://dl-cdn.alpinelinux.org/alpine/edge/community", \
  "texlive", \
  "texlive-full" \
]

# https://stackoverflow.com/q/55312675/7092954
ENV PATH=/usr/local/texlive/bin/x86_64-linuxmusl:$PATH

COPY texlive-profile.txt /tmp/

RUN apk --no-cache add \
  xz \
  wget \
  perl \
  tar \
  fontconfig-dev && \
  wget http://mirror.ctan.org/systems/texlive/tlnet/install-tl-unx.tar.gz && \
  mkdir /tmp/install-tl && \
  tar -xzf install-tl-unx.tar.gz -C /tmp/install-tl --strip-components=1 && \
  /tmp/install-tl/install-tl --profile=/tmp/texlive-profile.txt && \
  tlmgr update --self
RUN rm -rf install-tl-unx.tar.gz

ENTRYPOINT [ "shojo" ]
