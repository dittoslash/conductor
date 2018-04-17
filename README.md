# conductor
Web dashboard for controlling cluster computers

![screenshot](https://dittoslash.uk/conductor.png)

## Instructions
1. Install Python 3, [Requests](http://docs.python-requests.org/en/master/), and [Flask](http://flask.pocoo.org/) on the computer you want the webserver on.
2. Download the server/ directory onto the server.
3. Download ConductorClient onto each client computer (the ones being controlled).
4. Make sure the clients have passwordless sudo for reboot and shutdown.
5. Add the IPs and names of each client to the `IPS` and `NAMES` variables in `main.py` on the server.
6. Setup the `stats` file on the clients (see below).
7. ???
8. Profit!


## Setting up the stats file
In each client computer, place a file named stats next to the client executable.
If you don't want rich status data, create the file and leave it empty. As of right now the whole thing'll probably crash if the file's nonexistant.  
Each line should look like this:
`title|ff argument`  
Replace `title` with whatever you want to call the stat.  
Replace `ff` with:
* `cb`: `argument` will be run as a command and the stat will return 'success' or 'failure' depending on the exit status.
* `co`: `argument` will be run as a command and the stat will return its stdout and stderr.
