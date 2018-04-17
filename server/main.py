#!/usr/bin/python3.4
from flask import Flask, render_template #Web-server library
import requests #talk to target's ConductorClient instances
import os #run commands

#Replace items in these lists with the IPS of the targets, and names, respectively
IPS = ["192.192.0.10", "192.192.0.20"]
NAMES = ["Odin", "Kali"]

app = Flask(__name__) #initiate web-server

@app.route("/") #define /
def conduct():
	results = []
	for ip in IPS: #populate results table with pings
		try:
			result = requests.get("http://" + ip + ":8880", timeout=3).json()
			result['Checks'] = result['Checks'].items()
			results.append((True, result))
		except requests.exceptions.ConnectionError:
			results.append((False, {}))
	return render_template('conductor.html', iterator=zip(NAMES, IPS, results, range(0, len(IPS))))

@app.route("/reboot-<int:id>")
def reboot(id):
	requests.get("http://" + IPS[id] + ":8880/reboot", timeout=1)
	#Reboot the target system.
	#Requires passwordless sudo for reboot.
	#This function might error, because CC basically just dies immediately upon this being activated and won't return any response
	return "nice"

@app.route("/shutdown")
def shutdown():
	for ip, l in zip(IPS, LOGINS):
		requests.get("http://" + ip + ":8880/shutdown",timeout=1)
@app.route("/shutdownp")
def shutdownp():
	os.system("sudo shutdown -h now")

app.run(host="0.0.0.0", port=5000, debug=True)
#It's probably better to run Flask in a different way than its built-in server.
#Given the environment this will be run in, it's not worth implementing yet.
