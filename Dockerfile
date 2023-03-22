FROM crazymax/xgo:1.19
RUN apt-get update && \
    apt-get install -y libasound2-dev
