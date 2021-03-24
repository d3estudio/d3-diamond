from requests import get, post
from flask import \
    Flask, \
    jsonify, \
    redirect, \
    request as r, \
    render_template as render

app = Flask(__name__, static_url_path='/static/')

@app.login('/login', methods=['GET', 'POST'])
def login():
    pass

@app.route('/', methods=['GET'])
def index():
    token = r
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

        print(indices, file=open("PORRA.log", "a"))
        return render("index.html", user = user, indices = indices, scores = scores_formated)
    return redirect("/login")
