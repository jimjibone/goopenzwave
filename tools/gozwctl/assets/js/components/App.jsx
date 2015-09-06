var React = require('react');
var Nodes = require('./Nodes.jsx');
var NodeStore = require('../stores/NodeStore');
var NodeActions = require('../actions/NodeActions');

var App = React.createClass({
    render() {
        return (
            <div className="pure-g">
                <div className="pure-u-1-1">
                    <h1>gozwave Controller</h1>
                </div>
                <div className="pure-u-1-1">
                    <Nodes />
                </div>
            </div>
        );
    },
});

module.exports = App;
