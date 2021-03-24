FROM python:3.8.2-alpine
LABEL maintainer John <jotadev@me.com>

RUN pip install Flask==1.1.2
RUN pip install gspread==3.6.0
RUN pip install oauth2client==4.1.3
RUN pip install pep8==1.7.1
RUN pip install requests==2.23.0

RUN mkdir /app
ADD . /app

EXPOSE 5000

WORKDIR /app

CMD flask run --host=0.0.0.0 > app.log 2> app.err.log
