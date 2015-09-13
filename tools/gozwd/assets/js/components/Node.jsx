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
        if (this.props.node.location) {
            title += ' - ' + this.props.node.location;
        }

        var changed = this._isChanged(this.state.node);

        var values = [];
        for (var value in this.state.node.values) {
            if (this.state.node.values.hasOwnProperty(value)) {
                values.push(<Value stringid={value} value={this.state.node.values[value]} onChange={this.handleValueChange} onButton={this.handleValueButton} />);
            }
        }
        // {this.state.node.values.map((value) => {
        //     return (<Value value={value} onChange={this.handleValueChange} />);
        // })}

        // <button className="pure-button pure-button-primary" onClick={this.handleOnOff}>On/Off</button>
        return (
            <div key={this.props.node.node_info_id} className="pure-u-1-1 node">
                <div className="pure-g">
                    <div className="pure-u-1-1">
                        <h3>{title}</h3>
                    </div>
                    <div className="pure-u-1-1 pure-u-md-1-3">
                        <form className="pure-form pure-form-stacked" onSubmit={this.handleSubmit}>
                            <fieldset>
                                <div className="pure-u-1-2">
                                    <label>Name</label>
                                    <input type="text" className="pure-u-23-24" value={this.state.node.node_name} onChange={this.handleNameChange} />
                                </div>
                                <div className="pure-u-1-2">
                                    <label>Location</label>
                                    <input type="text" className="pure-u-23-24" value={this.state.node.location} onChange={this.handleLocationChange} />
                                </div>
                                <div className="pure-controls">
                                    <button type="submit" className="pure-button pure-button-primary" disabled={!changed}>Update</button>
                                </div>
                            </fieldset>
                        </form>
                    </div>
                    <div className="pure-u-1-1 pure-u-md-1-3">
                        <form className="pure-form pure-form-stacked">
                            <fieldset>
                                <div className="pure-u-1-2">
                                    <label>Type</label>
                                    <input type="text" className="pure-u-23-24" value={this.state.node.node_type} readOnly />
                                </div>
                                <div className="pure-u-1-2">
                                    <label>Manufacturer</label>
                                    <input type="text" className="pure-u-23-24" value={this.state.node.manufacturer_name} readOnly />
                                </div>
                                <div className="pure-u-1-2">
                                    <label>Product Name</label>
                                    <input type="text" className="pure-u-23-24" value={this.state.node.product_name} readOnly />
                                </div>
                                <div className="pure-u-1-2">
                                    <label>Node ID</label>
                                    <input type="text" className="pure-u-23-24" value={this.state.node.node_id} readOnly />
                                </div>
                            </fieldset>
                        </form>
                    </div>
                    <div className="pure-u-1-1 pure-u-md-1-3">
                        <form className="pure-form pure-form-stacked">
                            <fieldset>
                                <div className="pure-u-1-2">
                                    <label>Home ID</label>
                                    <input type="text" className="pure-u-23-24" value={this.state.node.home_id} readOnly />
                                </div>
                                <div className="pure-u-1-2">
                                    <label>Basic Type</label>
                                    <input type="text" className="pure-u-23-24" value={this.state.node.basic_type} readOnly />
                                </div>
                                <div className="pure-u-1-2">
                                    <label>Generic Type</label>
                                    <input type="text" className="pure-u-23-24" value={this.state.node.generic_type} readOnly />
                                </div>
                                <div className="pure-u-1-2">
                                    <label>Specific Type</label>
                                    <input type="text" className="pure-u-23-24" value={this.state.node.specific_type} readOnly />
                                </div>
                            </fieldset>
                        </form>
                    </div>
                    <div className="pure-u-1-1 pure-u-md-1-1">
                        <h3>Values</h3>
                        {values}
                    </div>
                </div>
            </div>
        );
    },

    handleOnOff(event) {
        NodeActions.sendNodeOnOff(this.state.node.node_info_id);
    },

    handleNameChange(event) {
        this.setState({ node: React.addons.update(this.state.node, {
            node_name: {$set: event.target.value.substr(0, 16)}}
        )});
    },

    handleLocationChange(event) {
        this.setState({ node: React.addons.update(this.state.node, {
            location: {$set: event.target.value.substr(0, 16)}}
        )});
    },

    handleValueChange(value_stringid, value_string) {
        // console.log('Node::handleValueChange:', value_stringid, value_string);
        this.setState({ node: React.addons.update(this.state.node, {
            values: {[value_stringid]: {string: {$set: value_string}}}
        })});
    },

    handleValueButton(value_stringid) {
        // console.log('Node::handleValueButton:', value_stringid);
        this.setState({ node: React.addons.update(this.state.node, {
            values: {[value_stringid]: {button_press: {$set: true}}}
        })});
    },

    handleSubmit(event) {
        event.preventDefault()
        // console.log('Node::handleSubmit');
        // console.log('Node::handleSubmit: state.node:', this.state.node);
        // console.log('Node::handleSubmit: props.node:', this.props.node);

        // This should only be called if the state.node has changed, but check
        // again anyway.
        if (this._isChanged(this.state.node)) {
            // Prepare the changes and send the change request to the server.
            NodeActions.sendNode(this.state.node);
        }
    },

    _isChanged(newNode) {
        // Just compare the new state to the props and determine if the
        // new state differs.
        var changed = false;
        if (newNode.node_name !== this.props.node.node_name) {
            changed = true;
        }
        else if (newNode.location !== this.props.node.location) {
            changed = true;
        }
        else {
            for (var value in this.props.node.values) {
                if (this.props.node.values.hasOwnProperty(value)) {
                    if (newNode.values.hasOwnProperty(value)) {
                        var currentValue = this.props.node.values[value];
                        var newValue = newNode.values[value];

                        if (newValue.string !== currentValue.string) {
                            changed = true;
                        } else if (newValue.button_press == true) {
                            changed = true;
                        }
                    }
                }
            }
        }
        return changed;
    },
});

var Value = React.createClass({
    handleChange(event) {
        this.props.onChange(this.props.stringid, event.target.value);
    },
    handleButton(event) {
        event.preventDefault();
        this.props.onButton(this.props.stringid);
    },

    render() {
        var disabled = this.props.value.read_only;

        return (
            <div key={this.props.stringid} className="pure-u-1-1 value">
                <div className="pure-g">
                    <div className="pure-u-1-1">
                        <form className="pure-form pure-form-stacked">
                            <fieldset>
                                <legend>{this.props.value.label}</legend>
                                {this.props.value.read_only ? <p>Read Only</p> : null}
                                {this.props.value.write_only ? <p>Write Only</p> : null}
                                <div className="pure-g">
                                    <div className="pure-u-1 pure-u-md-1-2 pure-u-lg-1-4">
                                        <label>Units</label>
                                        <input type="text" value={this.props.value.units} readOnly/>
                                    </div>
                                    <div className="pure-u-1 pure-u-md-1-2 pure-u-lg-1-4">
                                        <label>Value</label>
                                        <input type="text" value={this.props.value.string} onChange={this.handleChange} disabled={disabled}/>
                                    </div>
                                    <div className="pure-u-1 pure-u-md-1-2 pure-u-lg-1-4">
                                        <label>Genre</label>
                                        <input type="text" value={this.props.value.genre} readOnly/>
                                    </div>
                                    <div className="pure-u-1 pure-u-md-1-2 pure-u-lg-1-4">
                                        <label>Type</label>
                                        <input type="text" value={this.props.value.type} readOnly/>
                                    </div>
                                    {this.props.value.write_only ?
                                        <div className="pure-u-1 pure-u-md-1-2 pure-u-lg-1-4">
                                            <button className="pure-button" onClick={this.handleButton}>Press Button</button>
                                        </div>
                                        : null
                                    }
                                </div>
                            </fieldset>
                        </form>
                    </div>
                </div>
            </div>
        );

        // <ul key={value.ID}>
        //     <li>command_class_id: {value.command_class_id}</li>
        //     <li>genre: {value.genre}</li>
        //     <li>help: {value.help}</li>
        //     <li>label: {value.label}</li>
        //     <li>max: {value.max}</li>
        //     <li>min: {value.min}</li>
        //     <li>node_id: {value.node_id}</li>
        //     <li>polled: {value.polled}</li>
        //     <li>read_only: {value.read_only}</li>
        //     <li>write_only: {value.write_only}</li>
        //     <li>set: {value.set}</li>
        //     <li>string: {value.string}</li>
        //     <li>type: {value.type}</li>
        //     <li>units: {value.units}</li>
        //     <li>value_id: {value.id}</li>
        // </ul>
    },
});

module.exports = Node;
