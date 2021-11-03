from flask import Flask
import urllib.request
app = Flask(__name__)

@app.route('/')
def hello_world():
    return urllib.request.urlopen("test1.test").read()