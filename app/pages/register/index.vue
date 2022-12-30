<template>
  <div>
    <header>
      <div>
        <h1>Register</h1>
      </div>
    </header>
    <section id="register-section">
      <div>
        <v-form lazy-validation>
          <v-text-field v-model="registerForm.username" label="Username" required></v-text-field>
          <v-text-field v-model="registerForm.password" label="Password" required></v-text-field>
          <v-text-field v-model="confirmPassword" label="Confirm Password" required></v-text-field>
          <v-text-field v-model="registerForm.email" label="E-Mail" required></v-text-field>
          <div id="action">
            <v-btn color="error" @click="reset">Reset</v-btn>
            <v-btn color="primary" @click="register">Register</v-btn>
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
        registerForm: {
          username: "",
          password: "",
          email: ""
        },
        confirmPassword: ""
      }
    },
    methods: {
      register() {
        const allFieldCheck = this.registerForm.username === "" || this.registerForm.password === "" || this.registerForm.email === "";
        if(allFieldCheck) {
          console.log(`Please filled in all the field.`);
          return;
        }
        if(this.registerForm.password !== this.confirmPassword) {
          console.log(`password don't match.`)
        } else {
          console.log(this.registerForm);
        }
      },
      reset() {
        this.registerForm.username = "";
        this.registerForm.password = "";
        this.registerForm.email = "";
        this.confirmPassword = "";
      }
    }
  }
</script>

<style scoped>
  * {
    text-align: center;
  }

  #register-section {
    width: 50%;
    margin: 0 auto;
  }
</style>
