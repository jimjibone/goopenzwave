var API = {
    socket: null,
    connected: false,
    reconnection: null,
    buffer: [],
    handlers: {},

    connect: function() {
        // Check that the user's browser supports WebSockets.
        if (!window['WebSocket']) {
            // TODO: warn user that their browser does not support websockets.
            console.log('WARNING: this browser does not support websockets');
        } else {
            this.socket =  new WebSocket('ws://'+window.location.host+'/ws')

            var self = this;

            this.socket.onopen = function(event) {
                // console.log('API::onopen:', event)
                self.connected = true;
                for (var i = 0; i < self.buffer.length; i++) {
                    var message = self.buffer[i];
                    self.send(message.topic, message.payload);
                }
            }

            this.socket.onclose = function(event) {
                // console.log('API::onclose:', event)
                self.connected = false;

                // Attempt to reconnect.
                setTimeout(function() {
                    self.connect();
                }, 5000);
            }

            this.socket.onmessage = function(event) {
                // console.log('API::onmessage:', event)
                var message = JSON.parse(event.data);
                var handled = false;

                for (var h in self.handlers) {
                    var handler = self.handlers[h];
                    if (handler.topic === message.topic) {
                        handler.handler(message.payload);
                        handled = true;
                        break;
                    }
                }

                if (handled === false) {
                    console.log('API::onmessage: received unhandled message:', message);
                    // console.log('API::onmessage: handlers:', self.handlers);
                }
            }
        }
    },

    send: function(topic, payload) {
        var message = { topic: topic };
        if (payload) {
            message.payload = payload;
        }

        // Either send if the WebSocket has connected or buffer.
        if (this.connected) {
            // console.log('API::send: sending:', message);
            var data = JSON.stringify(message);
            this.socket.send(data);
        } else {
            // console.log('API::send: buffering: ', message);
            this.buffer.push(message);
        }
    },

    addHandler(topic, handler) {
        this.handlers[topic] = {
            topic: topic,
            handler: handler
        };
    }
}

module.exports = API;
