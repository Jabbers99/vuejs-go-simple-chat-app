<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
    </v-app-bar>

    <v-main>
      <UsernameDialog/>
      <Chatbox />
    </v-main>
  </v-app>
</template>

<script>
import Chatbox from './components/Chatbox';
import UsernameDialog from './components/UsernameDialog';

export default {
  name: 'App',

  components: {
    UsernameDialog,
    Chatbox,
  },

  data: () => ({
    //
  }),
};
</script>

<style>
html {
    overflow-y: auto;
}
</style>