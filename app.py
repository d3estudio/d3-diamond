import gspread
import json
import os

from flask import Flask
from flask import jsonify
from flask import send_from_directory
from oauth2client.service_account import ServiceAccountCredentials

scope = ['https://spreadsheets.google.com/feeds',
         'https://www.googleapis.com/auth/drive']

creds = ServiceAccountCredentials.from_json_keyfile_name('auth.json', scope)

client = gspread.authorize(creds)
sheet = client.open_by_key(os.environ["SHEET_ID"]).get_worksheet(0)

app = Flask(__name__, static_url_path='/static/')


@app.route('/api/<id>', methods=['GET'])
def api(id):
    for f in sheet.get_all_records():
        if id == f["ID"]:

            f["A"] = float(str(f["A"]).replace(",", "."))
            f["B"] = float(str(f["B"]).replace(",", "."))
            f["C"] = float(str(f["C"]).replace(",", "."))
            f["D"] = float(str(f["D"]).replace(",", "."))
            f["E"] = float(str(f["E"]).replace(",", "."))
            f["F"] = float(str(f["F"]).replace(",", "."))
            f["G"] = float(str(f["G"]).replace(",", "."))
            f["H"] = float(str(f["H"]).replace(",", "."))
            f["I"] = float(str(f["I"]).replace(",", "."))
            f["J"] = float(str(f["J"]).replace(",", "."))

            f["CHANGE_A"] = float(str(f["CHANGE_A"]).replace(",", "."))
            f["CHANGE_B"] = float(str(f["CHANGE_B"]).replace(",", "."))
            f["CHANGE_C"] = float(str(f["CHANGE_C"]).replace(",", "."))
            f["CHANGE_D"] = float(str(f["CHANGE_D"]).replace(",", "."))
            f["CHANGE_E"] = float(str(f["CHANGE_E"]).replace(",", "."))
            f["CHANGE_F"] = float(str(f["CHANGE_F"]).replace(",", "."))
            f["CHANGE_G"] = float(str(f["CHANGE_G"]).replace(",", "."))
            f["CHANGE_H"] = float(str(f["CHANGE_H"]).replace(",", "."))
            f["CHANGE_I"] = float(str(f["CHANGE_I"]).replace(",", "."))
            f["CHANGE_J"] = float(str(f["CHANGE_J"]).replace(",", "."))

            response = app.response_class(
                response=json.dumps(f),
                status=200,
                mimetype='application/json'
            )
            return response

    response = app.response_class(
        response=json.dumps({}),
        status=404,
        mimetype='application/json'
    )
    return response


@app.route('/')
def home():
    return app.send_static_file('index.html')
