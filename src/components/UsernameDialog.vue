<template>
  <div class="text-center">
    <v-dialog
      v-model="dialog"
      width="500"
    >
      <v-card>
        <v-card-title class="text-h5 grey lighten-2">Pick a username</v-card-title>
        <v-card-text>
            <v-text-field 
                v-model="username"
                label="Username" 
                ref="username"
                :rules="[() => !!username || 'This field is required']"
            />
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="primary"
            text
            @click="changeUsername"
          >Choose</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>

  export default {
    data () {
        return {
            dialog: false,
            username: "",
            formHasErrors: false,
        }
    },
    computed: {
        form () {
            return {
                username: null,
            }
        }
    },
    methods: {
        changeUsername: function() {
            if (this.username !== "" && this.username.length < 32 && this.username.length > 4) {
                this.dialog = false;
                this.$socket.send(this.username);
            }
        }
    }
  }
</script>

