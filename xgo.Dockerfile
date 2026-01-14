FROM techknowlogick/xgo:go-1.25.4
RUN apt update && \
    apt install -y libasound2-dev
