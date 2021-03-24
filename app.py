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
        scores = get("http://api:8000/user/2/scores", headers={"Authorization":token}).json()['data']
        info = get("http://api:8000/score-type", headers={"Authorization":token}).json()['data']

        indices = {}

        i = 0
        l = "abcdefghijklmnopqrstuvwxyz"
        scores_formated = {}
        for j in info:
            indices[l[i]] = j
            for s in scores:
                if s['name'].lower() == j['name'].lower():
                    scores_formated[s['name'].lower()] = s
            i += 1

        return render("index.html", user = user, indices = indices, scores = scores_formated)
    return redirect("/login")
