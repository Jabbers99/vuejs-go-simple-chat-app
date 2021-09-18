module.exports = {
    Connect : function( chatbox, usernameDialog ) {
        console.log("Starting connection to WebSocket Server")

        this.connection = new WebSocket("ws://localhost:4500/api/ws")

        this.connection.onmessage = function(event) {
            chatbox.NewMessage(event.data)
            
        }

        this.connection.onopen = function() {
            chatbox.Loading = false;
            usernameDialog.dialog = true;
            console.log("Successfully connected to the echo websocket server...",)
        }
    },
    Message: function(message) {
        this.connection.send(message)
    }
}

