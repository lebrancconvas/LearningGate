<template>
  <div>
    <header>
      <div>
        <h1>Login</h1>
      </div>
    </header>
    <section id="login-section">
      <div>
        <v-form lazy-validation>
          <v-text-field v-model="loginForm.username" label="Username" required></v-text-field>
          <v-text-field v-model="loginForm.password" type="password" label="Password" required></v-text-field>
          <div id="action">
            <v-btn color="primary" @click="login">Login</v-btn>
          </div>
        </v-form>
      </div>
    </section>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        loginForm: {
          username: "",
          password: ""
        }
      }
    },
    methods: {
      login() {
        console.log(this.loginForm);
      }
    }
  }
</script>

<style scoped>
  * {
    text-align: center;
  }

  #login-section {
    width: 50%;
    margin: 0 auto;
  }
</style>
