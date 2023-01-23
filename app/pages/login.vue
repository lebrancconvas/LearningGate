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
  import axios from 'axios';
  export default {
    data() {
      return {
        loginForm: {
          username: "",
          password: ""
        },
        results: []
      }
    },
    mounted() {
      const url = `http://localhost:8002/api/v1/users`;

      axios.get(url)
        .then(res => {
          this.results = res.data;
          console.log(this.results);
        })
        .catch(err => {
          console.error(err);
          console.log(`[ERROR] Fetching User Data error.`);
        })

    },
    methods: {
      login() {
        const username = this.results.filter(user => user.Username === this.loginForm.username);
        const password = this.results.filter(user => user.Password === this.loginForm.password);

        const checkUsernameExist = username[0] !== undefined;
        const checkPasswordExist = password[0] !== undefined;

        if(checkUsernameExist && checkPasswordExist) {
          const isMatchLogin = username[0] === password[0];
          if(isMatchLogin) {
            alert("Login Success");
            this.$store.commit("LOGIN");
          } else {
            alert("Login error: Password not match.");
          }
        } else {
          alert("Login error: Either Username or Password cannot find in the database.");
        }
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
