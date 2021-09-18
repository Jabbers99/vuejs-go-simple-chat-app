import Vue from 'vue'
//import Vuex from 'vuex'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import VueNativeSock from 'vue-native-websocket'

Vue.use(VueNativeSock, 'ws://localhost:4500/api/ws', {
    reconnection: true, // (Boolean) whether to reconnect automatically (false)
    reconnectionAttempts: 5, // (Number) number of reconnection attempts before giving up (Infinity),
    reconnectionDelay: 3000, // (Number) how long to initially wait before attempting a new (1000)
})
Vue.config.productionTip = false


const vm = new Vue({
    vuetify,
    render: h => h(App)
})
vm.$mount('#app');

