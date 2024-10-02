FROM techknowlogick/xgo:go-1.23.x
RUN apt-get update && \
    apt-get install -y libasound2-dev
