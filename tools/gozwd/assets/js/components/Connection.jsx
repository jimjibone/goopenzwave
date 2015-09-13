var React = require('react');
var ConnectionStore = require('../stores/ConnectionStore');

var Connection = React.createClass({
    getInitialState() {
        return ConnectionStore.getState();
    },

    componentDidMount() {
        ConnectionStore.listen(this.onChange);
    },

    componentWillUnmount() {
        ConnectionStore.unlisten(this.onChange);
    },

    onChange(state) {
        this.setState(state);
    },

    render() {
        if (this.state.connected == false) {
            return (
                <div className="pure-g connection disconnected">
                    <div className="pure-u-1-1">
                        <h2>Connecting...</h2>
                    </div>
                </div>
            );
        }

        return (
            <div className="pure-g connection connected">
            </div>
        );
    },
});

module.exports = Connection;
