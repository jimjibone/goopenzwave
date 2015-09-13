var React = require('react');
var Connection = require('./Connection.jsx');
var Nodes = require('./Nodes.jsx');
var NodeStore = require('../stores/NodeStore');
var NodeActions = require('../actions/NodeActions');

var App = React.createClass({
    render() {
        return (
            <div className="pure-g">
                <div className="pure-u-1-1 heading">
                    <h1>goopenzwave</h1>
                </div>
                <Connection />
                <div className="pure-u-1-1">
                    <Nodes />
                </div>
            </div>
        );
    },
});

module.exports = App;
