import json
import os
import requests

from flask import Flask
from flask import jsonify
from flask import send_from_directory
from requests.auth import HTTPBasicAuth

app = Flask(__name__, static_url_path='/static/')


@app.route('/api/<id>', methods=['GET'])
def api(id):
    r = requests.get(
        url=os.environ["API_URL"],
        params={
            "ignore_case": "false",
            "ID": id,  # asd97asdhasd87
        },
        auth=HTTPBasicAuth(os.environ["API_KEY"], os.environ["API_SEC"])
    )
    response = app.response_class(
        response=json.dumps(r.json()),
        status=200,
        mimetype='application/json'
    )
    return response


@app.route('/')
def home():
    return app.send_static_file('index.html')
