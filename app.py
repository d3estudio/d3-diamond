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

def load_scores(user, token):
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
                if not "value" in s:
                    s["value"] = 0
                if not s['name'].lower() in scores_formated:
                    scores_formated[s['name'].lower()] = s
                    scores_formated[s['name'].lower()]["total"] = 1
                else:
                    scores_formated[s['name'].lower()]["value"] += s["value"]
                    scores_formated[s['name'].lower()]["total"] += 1

        i += 1
    return [scores_formated, info, indices]

def log(*args, **kwargs):
    print(*args, **kwargs, file=open("app.log", "a"))

def gen_csv(indices, scores):
    # 'ID,NAME,DATE,A,B,C,D,E,F,G,H,I,J,CHANGE_A,CHANGE_B,CHANGE_C,CHANGE_D,CHANGE_E,CHANGE_F,CHANGE_G,CHANGE_H,CHANGE_I,CHANGE_J\n'
    csv = ',,,'

    letters = "abcdefghijklmnopqrstuvwxyz"
    for l in letters:
        if l in indices:
            i = indices[l]
            csv += "\"%.1f\"," % scores[i["name"]]["value"]
    for l in letters:
        if l in indices:
            i = indices[l]
            if 'last' in scores[i["name"]]:
                csv += str(scores[i["name"]]["last"]) + ','
            else:
                csv += "0,"
    if csv[-1] == ',':
        csv = csv[:-1]

    log(csv)
    return csv

@app.route('/dates', methods=['GET'])
def date_filters():
    token = r.cookies.get("d3diamond_token")
    response = post("http://api:8000/verify", headers={"Authorization": token})
    if response.status_code == 200 and response.json()['type'].lower() == 'sucess':
        user = response.json()['data']

        dates = []
        try:
            pre_dates = get("http://api:8000/user/{}/dates".format(user["id"]), headers={"Authorization": token}).json()['data']
            for date in pre_dates:
                if "id" in date:
                    del date["id"]

                if "class" in date:
                    del date["class"]

                if date not in dates:
                    d = dt.fromisoformat(date["date"][:10])

                    date["date"] = f"{str(d.month).rjust(2).replace(' ', '0')}/{str(d.year).rjust(4).replace(' ', '0')}"
                    date["url"] =  f"{str(d.year).rjust(4).replace(' ', '0')}{str(d.month).rjust(2).replace(' ', '0')}"
                    dates.append(date)
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
        scores, info, indices = load_scores(user, token)

        return render("index.html", user = user, indices = indices, scores = scores, csv = gen_csv(indices, scores))
    return redirect("/login")


@app.route("/<date>", methods = ["GET"])
def filter_by_date(date):
    date = dt.fromisoformat(date[:4]+"-"+date[4:]+"-01")
    dateend = dt(day=date.day, month=date.month+1, year=date.year)

    token = r.cookies.get("d3diamond_token")
    response = post("http://api:8000/verify", headers={"Authorization":token})
    if response.status_code == 200:
        user = response.json()['data']
        scores, info, indices = load_scores(user, token)

        return render("index.html", user = user, indices = indices, scores = scores, csv = gen_csv(indices, scores))
    return redirect("/login")
