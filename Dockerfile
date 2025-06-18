FROM techknowlogick/xgo:go-1.24.4
RUN apt-get update && \
    apt-get install -y libasound2-dev
