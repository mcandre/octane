FROM karalabe/xgo-latest
RUN apt-get update && \
    apt-get install -y libasound2-dev
