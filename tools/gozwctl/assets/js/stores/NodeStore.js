var alt = require('../alt');
var NodeActions = require('../actions/NodeActions');
var api = require('../api');

class NodeStore {
    constructor() {
        // this.api = api;
        this.nodes = {};
        this.errorMessage = null;

        this.bindListeners({
            handleFetchNodes: NodeActions.FETCH_NODES,
            handleUpdateNodes: NodeActions.UPDATE_NODES,
            handleSendNode: NodeActions.SEND_NODE,
            handleUpdateNode: NodeActions.UPDATE_NODE,
            handleUpdateFailed: NodeActions.UPDATE_FAILED,
        });

        api.addHandler('nodes', NodeActions.updateNodes);
        api.addHandler('node-updated', NodeActions.updateNode);
        api.connect();
    }

    handleFetchNodes() {
        this.nodes = {};
        api.send('get-nodes');
    }

    handleUpdateNodes(nodes) {
        console.log('NodeStore::handleUpdateNodes:', nodes);
        this.nodes = nodes;
        this.errorMessage = null;
    }

    handleSendNode(node) {
        console.log('NodeStore::handleSendNode:', node);
        api.send('set-node', node);
    }

    handleUpdateNode(node) {
        // Update the existing node if it exists, otherwise add it.
        console.log('NodeStore::handleUpdateNode:', node);
    }

    handleUpdateFailed(errorMessage) {
        this.errorMessage = errorMessage;
    }
}

module.exports = alt.createStore(NodeStore, 'NodeStore');
