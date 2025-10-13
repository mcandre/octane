FROM techknowlogick/xgo:go-1.25.2
RUN apt-get update && \
    apt-get install -y libasound2-dev
