<script>

    import Chatbox from './components/Chatbox';
    import UsernameDialog from './components/UsernameDialog';
    
    
    export default {
        name: 'App',

        components: {
            UsernameDialog,
            Chatbox,
        },
        mounted() {
            this.$refs.Chatbox.connect(this.$refs.UsernameDialog)
            this.$options.sockets.onmessage = (data) => console.log(data)
            this.$options.sockets.onopen = (data) => console.log(data)

        }
    };
</script>
<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
    </v-app-bar>

    <v-main>
      <UsernameDialog ref="UsernameDialog" />
      <Chatbox ref="Chatbox" />
    </v-main>
  </v-app>
</template>



<style>
html {
    overflow-y: auto;
}
</style>