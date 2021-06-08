FROM python:2-alpine
RUN apk add --no-cache git
RUN git clone --depth 1 https://github.com/sqlmapproject/sqlmap.git
WORKDIR /sqlmap
COPY run.sh .
#ENTRYPOINT ["python", "sqlmap.py"]
