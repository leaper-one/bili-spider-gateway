FROM alpine
  
WORKDIR /build

COPY bh .

CMD ["./bh"]
