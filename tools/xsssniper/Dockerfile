FROM python:2.7-slim

WORKDIR /usr/src/app

COPY requirement.txt src run.sh ./

RUN pip install --no-cache-dir mechanize lxml

# CMD [ "python", "./xsssniper.py" ]
