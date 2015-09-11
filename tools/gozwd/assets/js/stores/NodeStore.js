var alt = require('../alt');
var NodeActions = require('../actions/NodeActions');
var api = require('../api');

class NodeStore {
    constructor() {
        // this.api = api;
        this.nodes = [];
        this.errorMessage = null;

        this.bindListeners({
            handleFetchNodes: NodeActions.FETCH_NODES,
            handleUpdateNodes: NodeActions.UPDATE_NODES,
            handleSendNodeOnOff: NodeActions.SEND_NODE_ON_OFF,
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

    handleSendNodeOnOff(node_info_id) {
        // console.log('NodeStore::handleSendNodeOnOff:', node_info_id);
        api.send('toggle-node', {
            node_info_id: node_info_id
        });
    }

    handleSendNode(node) {
        // console.log('NodeStore::handleSendNode:', node);
        api.send('set-node', node);
    }

    handleUpdateNode(node) {
        // Update the existing node if it exists, otherwise add it.
        console.log('NodeStore::handleUpdateNode:', node);
        var add = true;
        for (var i = 0, len = this.nodes.length; i < len; ++i) {
            if (this.nodes[i].node_info_id === node.node_info_id) {
                // This is an update to an existing node.
                add = false;
                this.nodes[i] = node;
                break;
            }
        }

        if (add) {
            // We didn't find the node in the list, so add it.
            this.nodes.push(node);
        }
    }

    handleUpdateFailed(errorMessage) {
        this.errorMessage = errorMessage;
    }
}

module.exports = alt.createStore(NodeStore, 'NodeStore');
