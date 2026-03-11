FROM techknowlogick/xgo:go-1.26.1
RUN apt update && \
    apt install -y libasound2-dev
