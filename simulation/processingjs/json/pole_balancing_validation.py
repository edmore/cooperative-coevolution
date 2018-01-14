import plotly
import plotly.graph_objs as go
import json

# state variables
X = []
Theta1 = []
Theta2 = []

with open('sane_parallel.json') as data_file:
    data = json.load(data_file)
    for v in data:
        # print(v['X'], v['Theta1'], v['Theta2'])
        X.append(v['X'] * 100)  # cm
        Theta1.append(v['Theta1'] * 57.295)  # degrees
        Theta2.append(v['Theta2'] * 57.295)  # degrees

position_of_cart_layout = dict(title='',
                               xaxis=dict(title='Time steps'),
                               yaxis=dict(title='Position of the cart (cm)'),
                               font=dict(size=16)
                               )

pole_angles_layout = dict(title='',
                          xaxis=dict(title='Time steps'),
                          yaxis=dict(title='Pole angles (degrees)'),
                          font=dict(size=16)
                          )

cart_position = go.Scattergl(
    y=X,
    mode='lines',
)

long_pole = go.Scattergl(
    y=Theta1,
    mode='lines',
    name='long pole angle'
)

short_pole = go.Scattergl(
    y=Theta2,
    mode='lines',
    name='short pole angle'
)

# Postion of cart
plotly.offline.plot({
    "data": [cart_position],
    "layout": position_of_cart_layout,
}, image='png')


# Pole angles (theta1 and theta2)
plotly.offline.plot({
    "data": [long_pole, short_pole],
    "layout": pole_angles_layout,
}, image='png')
