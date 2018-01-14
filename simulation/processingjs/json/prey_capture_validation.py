import plotly
import plotly.graph_objs as go
import json


# state variables
# [{"PredatorX":[0,14,16],"PredatorY":[51,0,0],"PreyX":16,"PreyY":0},...] - caught by predator3
# {"PredatorX":[17,16,18],"PredatorY":[51,0,0],"PreyX":16,"PreyY":0} - caught by predator2

PreyX = []
PreyY = []
Predator1X = []
Predator1Y = []
Predator2X = []
Predator2Y = []
Predator3X = []
Predator3Y = []

with open('multi_agent_esp_parallel.json') as data_file:
    data = json.load(data_file)
    for v in data:
        PreyX.append(v['PreyX'])
        PreyY.append(v['PreyY'])
        Predator1X.append(v['PredatorX'][0])
        Predator1Y.append(v['PredatorY'][0])
        Predator2X.append(v['PredatorX'][1])
        Predator2Y.append(v['PredatorY'][1])
        Predator3X.append(v['PredatorX'][2])
        Predator3Y.append(v['PredatorY'][2])

layout = dict(title='',
                          xaxis=dict(title='X coordinates'),
                          yaxis=dict(title='Y coordinates'),
                          font=dict(size=16)
                          )

# prey
prey = go.Scattergl(
    x=PreyX,
    y=PreyY,
    mode='markers',
    name='prey'
)

# predators
predator1 = go.Scattergl(
    x=Predator1X,
    y=Predator1Y,
    mode='markers',
    name='predator1'
)

predator2 = go.Scattergl(
    x=Predator2X,
    y=Predator2Y,
    mode='markers',
    name='predator2'
)

predator3 = go.Scattergl(
    x=Predator3X,
    y=Predator3Y,
    mode='markers',
    name='predator3'
)

# Postion of the prey and predators
plotly.offline.plot({
    "data": [prey, predator1, predator2, predator3],
    "layout": layout,
}, image='png')
