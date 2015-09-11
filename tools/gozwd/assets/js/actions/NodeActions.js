var alt = require('../alt');

class NodeActions {
    constructor() {
    }

    // All nodes.

    fetchNodes() {
        this.dispatch();
    }

    updateNodes(nodes) {
        this.dispatch(nodes);
    }

    // Individual node.

    // fetchNode(home_id, node_id) {...}

    sendNodeOnOff(node_info_id) {
        this.dispatch(node_info_id);
    }

    sendNode(node) {
        this.dispatch(node);
    }

    updateNode(node) {
        this.dispatch(node);
    }

    updateFailed(errorMessage) {
        this.dispatch(errorMessage);
    }
}

module.exports = alt.createActions(NodeActions);
