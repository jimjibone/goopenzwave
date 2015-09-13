'use strict';

import React from 'react'
import ConnectionActions from './actions/ConnectionActions'
import NodeActions from './actions/NodeActions'
import App from './components/App.jsx'
import api from './api'

api.addHandler('nodes', NodeActions.updateNodes)
api.addHandler('node-updated', NodeActions.updateNode)
api.onConnect(ConnectionActions.connected)
api.onDisconnect(ConnectionActions.disconnected)
api.connect();

// Build the React app.
React.render(
    React.createElement(App, null),
    document.getElementById('app')
)
