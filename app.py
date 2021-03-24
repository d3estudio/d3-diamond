from datetime import date as dt
from requests import get, post
from flask import \
    Flask, \
    jsonify, \
    redirect, \
    request as r, \
    make_response as mkres, \
    render_template as render

app = Flask(__name__, static_url_path='/static/')
log = open("app.log", "a")

@app.route('/dates', methods=['GET'])
def date_filters():
    token = r.cookies.get("d3diamond_token")
    response = post("http://api:8000/verify", headers={"Authorization":token})
    if response.status_code == 200:
        user = response.json()['data']
        try:
            dates = get("http://api:8000/user/{}/dates".format(user["id"])).json()['data']
            for date in dates:
                dates[date]["url"] = ""
        except:
            dates = None

        return render('dates.html', user = user, dates = dates)
    return redirect("/login")

@app.route('/logout', methods=['GET', 'POST'])
def logout():
    token = r.cookies.get("d3diamond_token")
    post("http://api:8000/logout", headers={"Authorization": token})

    response = mkres(redirect("/"))
    response.set_cookie("d3diamond_token", "", expires=0)
    return response

@app.route('/login', methods=['GET', 'POST'])
def login():
    if r.method == 'GET':
        return render("login.html")

    email = r.form.get("email")
    passw = r.form.get("password")

    response = post("http://api:8000/login", json = {
        "email": email,
        "pass": passw
    })

    if response.status_code == 200:
        token = response.json()['data']['token']

        response = mkres(redirect("/"))
        response.set_cookie("d3diamond_token", token)

        return response

    return render("login.html", error=["Falha no login", response.json()['message']])

@app.route('/', methods=['GET'])
def index():
    token = r.cookies.get("d3diamond_token")
    response = post("http://api:8000/verify", headers={"Authorization":token})
    if response.status_code == 200:
        user = response.json()['data']
        try:
            scores = get("http://api:8000/user/{}/scores".format(user["id"]), headers={
                "Authorization": token
            }).json()['data']
        except:
            scores = []
        info = get("http://api:8000/score-type", headers={"Authorization":token}).json()['data']

        indices = {}

        i = 0
        l = "abcdefghijklmnopqrstuvwxyz"
        scores_formated = {}
        for j in info:
            indices[l[i]] = j
            for s in scores:
                if s['name'].lower() == j['name'].lower():
                    if not s['name'].lower() in scores_formated:
                        scores_formated[s['name'].lower()] = s
                        scores_formated[s['name'].lower()]["total"] = 1
                    else:
                        if "value" in scores_formated[s['name'].lower()]:
                            scores_formated[s['name'].lower()]["value"] += s["value"]
                            scores_formated[s['name'].lower()]["total"] += 1
            i += 1

        return render("index.html", user = user, indices = indices, scores = scores_formated)
    return redirect("/login")

@app.route("/<date>", methods = ["GET"])
def filter_by_date(date):
    date = dt.fromisoformat(date[:4]+"-"+date[4:]+"-01")
    dateend = dt(day=date.day, month=date.month+1, year=date.year)

    token = r.cookies.get("d3diamond_token")
    response = post("http://api:8000/verify", headers={"Authorization":token})
    if response.status_code == 200:
        user = response.json()['data']

        try:
            scores = get("http://api:8000/user/{}/scores".format(user["id"]), headers={
                "Authorization": token,
                "Query": "created_at BETWEEN '{}' AND '{}'".format(
                    date.isoformat(), dateend.isoformat()
                )
            }).json()['data']
        except:
            scores = []

        info = get("http://api:8000/score-type", headers={"Authorization":token}).json()['data']

        indices = {}

        i = 0
        l = "abcdefghijklmnopqrstuvwxyz"
        scores_formated = {}
        for j in info:
            indices[l[i]] = j
            for s in scores:
                if s['name'].lower() == j['name'].lower():
                    if not s['name'].lower() in scores_formated:
                        scores_formated[s['name'].lower()] = s
                        scores_formated[s['name'].lower()]["total"] = 1
                    else:
                        if "value" in scores_formated[s['name'].lower()]:
                            scores_formated[s['name'].lower()]["value"] += s["value"]
                            scores_formated[s['name'].lower()]["total"] += 1
            i += 1

        return render("index.html", user = user, indices = indices, scores = scores_formated)
    return redirect("/login")
