<template>
    <v-container>
        <div id="chatbox">
            <v-progress-circular id="loading" v-if="loading"
            indeterminate
            :size="70"
            ></v-progress-circular>
            <div v-for="(msg, i) in messages" :key="i">
                <div v-if="msg.messageType == 'default'">
                    <v-avatar
                        class="avatar"
                        color="primary"
                        size="56"
                    >
                        <span class="white--text text-h5">{{ msg.initials }}</span>
                    </v-avatar>
                    <div class="chattext-default">
                    <b>{{ msg.username }}</b><br>
                        {{ msg.message }}
                    </div>
                </div>
                <div v-if="msg.messageType == 'info'">
                    <span class="chattext-info">{{msg.message}}</span>
                </div>
            </div>
        </div>

        <v-text-field label="Type a message" v-model="Message" v-on:keyup.enter="onEnter" ref="MessageInput" />
    </v-container>
</template>
<script>
    function isEmptyOrSpaces(str){
        return str === null || str.match(/^ *$/) !== null;
    }
    export default {
        props: {
            Websockets : {}
        },
        data() {
            return {
                loading: true,
                Message: "",
                
                messages: []
            }
        },
        watch: {
            loading: function(newV) {
                this.loading = newV;
            }
        },
        methods: {
            newMessage: function(msg) {
                msg = JSON.parse(msg);
                msg.initials = (msg.username.charAt(0) + msg.username.charAt(msg.username.length/2)).toUpperCase()
                this.messages.push(msg)
            },
            onEnter: function() {
                if (!isEmptyOrSpaces(this.Message)) {
                    console.log("Sent message: ", this.Message)
                    this.$socket.send(this.Message)
                    this.Message = ""
                }
            },
            connect : function(usernameDialog) {
                usernameDialog.dialog = true;
            }
        },
        mounted() {
            this.$options.sockets.onopen = function() {
                this.loading = false;
            }
            this.$options.sockets.onmessage = function(msg) {
                this.newMessage(msg.data);
            }
        }
    }
</script>

<style scoped>
    .avatar {
        float: left;
        margin: 20px 20px 0 20px;
    }
    .chattext-default {
        padding: 20px;
    }
    .chattext-info {
        margin-top: 20px;
        margin-left: 20px;
        display: block;
    }
    #chatbox {
        height: 500px;
        background-color: #f1f1f1;
        border-radius: 5px;
        overflow: auto;
    }
    #loading {
        display: block;
        margin: 65px auto 0 auto;
    }
</style>