var React = require('react/addons');
var NodeActions = require('../actions/NodeActions');

var Node = React.createClass({
    getInitialState() {
        return {
            node: {},
        }
    },

    componentWillMount() {
        this.setState({
            node: React.addons.update(this.state.node, {$set: this.props.node}),
        })
    },

    componentWillReceiveProps(nextProps) {
        this.setState({
            node: React.addons.update(this.state.node, {$set: nextProps.node}),
        });
    },

    render() {
        var title = 'Node ' + this.props.node.node_id;
        if (this.props.node.node_name) {
            title = this.props.node.node_name;
        }

        var changed = this._getChanged(this.state.node);

        return (
            <div key={this.props.node.node_info_id} className="pure-u-1-1 node">
                <div className="pure-g">
                    <div className="pure-u-1-2">
                        <h3>{title}</h3>
                        <form className="pure-form pure-form-aligned" onSubmit={this.handleSubmit}>
                            <fieldset>
                                <div className="pure-control-group">
                                    <label htmlFor="name">Name</label>
                                    <input type="text" value={this.state.node.node_name} onChange={this.handleNameChange} />
                                </div>
                                <p>Location: {this.props.node.location}</p>
                                <p>Type: {this.props.node.node_type}</p>
                                <p>Manufacturer: {this.props.node.manufacturer_name}</p>
                                <p>Product Name: {this.props.node.product_name}</p>
                                <p>Node ID: {this.props.node.node_id}</p>
                                <p>Home ID: {this.props.node.home_id}</p>
                                <p>Basic Type: {this.props.node.basic_type}</p>
                                <p>Generic Type: {this.props.node.generic_type}</p>
                                <p>Specific Type: {this.props.node.specific_type}</p>
                                <div className="pure-controls">
                                    <button type="submit" className="pure-button pure-button-primary" disabled={!changed}>Submit</button>
                                </div>
                            </fieldset>
                        </form>
                    </div>
                    <div className="pure-u-1-2">
                        <h3>Values</h3>
                        {this.props.node.values.map((value) => {
                            return (
                                <ul key={value.ID}>
                                    <li>command_class_id: {value.command_class_id}</li>
                                    <li>genre: {value.genre}</li>
                                    <li>help: {value.help}</li>
                                    <li>label: {value.label}</li>
                                    <li>max: {value.max}</li>
                                    <li>min: {value.min}</li>
                                    <li>node_id: {value.node_id}</li>
                                    <li>polled: {value.polled}</li>
                                    <li>read_only: {value.read_only}</li>
                                    <li>write_only: {value.write_only}</li>
                                    <li>set: {value.set}</li>
                                    <li>string: {value.string}</li>
                                    <li>type: {value.type}</li>
                                    <li>units: {value.units}</li>
                                    <li>value_id: {value.id}</li>
                                </ul>
                            );
                        })}
                    </div>
                </div>
            </div>
        );
    },

    handleNameChange(event) {
        this.setState({ node: React.addons.update(this.state.node, {
            node_name: {$set: event.target.value.substr(0, 16)}}
        )});
    },

    handleSubmit(event) {
        event.preventDefault()
        console.log('Node::handleSubmit');
        console.log('Node::handleSubmit: state.node:', this.state.node);
        console.log('Node::handleSubmit: props.node:', this.props.node);

        // This should only be called if the state.node has changed, but check
        // again anyway.
        if (this._getChanged(this.state.node)) {
            // Prepare the changes and send the change request to the server.
            NodeActions.sendNode(this.state.node);
        }
    },

    _getChanged(newNode) {
        // Just compare the new state to the props and determine if the
        // new state differs.
        var changed = false;
        if (newNode.node_name !== this.props.node.node_name) {
            changed = true;
        }
        return changed;
    },
});

module.exports = Node;
