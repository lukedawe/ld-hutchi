FROM python:3.13

WORKDIR /populate

ADD populate.py dogs.json ./

RUN pip install python-dotenv requests

CMD [ "python", "./populate.py" ]