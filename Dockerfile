FROM techknowlogick/xgo:go-1.27.0
RUN apt-get update && \
    apt-get install -y libasound2-dev
