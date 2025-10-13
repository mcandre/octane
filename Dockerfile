FROM techknowlogick/xgo:go-1.25.3
RUN apt-get update && \
    apt-get install -y libasound2-dev
