FROM techknowlogick/xgo:go-1.24.x
RUN apt-get update && \
    apt-get install -y libasound2-dev
