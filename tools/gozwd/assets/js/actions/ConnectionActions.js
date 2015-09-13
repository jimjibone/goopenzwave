var alt = require('../alt');

class ConnectionActions {
    constructor() {
    }

    // Connection state.

    connected() {
        this.dispatch();
    }

    disconnected() {
        this.dispatch();
    }
}

module.exports = alt.createActions(ConnectionActions);
