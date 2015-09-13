var alt = require('../alt');
var ConnectionActions = require('../actions/ConnectionActions');
var api = require('../api');

class ConnectionStore {
    constructor() {
        this.connected = api.connected;

        this.bindListeners({
            handleConnected: ConnectionActions.CONNECTED,
            handleDisconnected: ConnectionActions.DISCONNECTED,
        });
    }

    handleConnected() {
        this.connected = true;
    }

    handleDisconnected() {
        this.connected = false;
    }
}

module.exports = alt.createStore(ConnectionStore, 'ConnectionStore');
