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
          <v-text-field v-model="registerForm.displayname" label="Display Name"></v-text-field>
          <v-text-field v-model="registerForm.password" type="password" label="Password" required></v-text-field>
          <v-text-field v-model="confirmPassword" type="password" label="Confirm Password" required></v-text-field>
          <v-text-field v-model="registerForm.email" type="email" label="E-Mail" required></v-text-field>
          <div id="action">
            <v-btn color="error" @click="reset">Reset</v-btn>
            <v-btn color="primary" @click="register">Register</v-btn>
          </div>
        </v-form>
      </div>
      <div id="alert">
        <v-overlay :value="overlay">
          <v-alert v-model="alertFillAll" color="red" transition="scale-transition">
            Please filled in all the field.
            <v-btn color="error" @click="closeFillAllAlert">
              Close
            </v-btn>
          </v-alert>
          <v-alert v-model="alertPasswordNotMatch" color="red" transition="scale-transition">
            Password doesn't match.
            <v-btn color="error" @click="closePasswordNotMatchAlert">
              Close
            </v-btn>
          </v-alert>
        </v-overlay>
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
          displayname: "",
          password: "",
          email: ""
        },
        confirmPassword: "",
        alertFillAll: false,
        alertPasswordNotMatch: false,
        overlay: false
      }
    },
    methods: {
      register() {
        const allFieldCheck = this.registerForm.username === "" || this.registerForm.password === "" || this.registerForm.email === "";
        const emptyDisplayName = this.registerForm.displayname === "";
        if(allFieldCheck) {
          console.log(`Please filled in all the field.`);
          this.overlay = true;
          this.alertFillAll = true;
          return;
        }
        if(this.registerForm.password !== this.confirmPassword) {
          console.log(`password don't match.`);
          this.overlay = true;
          this.alertPasswordNotMatch = true;
        } else {
          if(emptyDisplayName) {
            this.registerForm.displayname = this.registerForm.username;
          }
          console.log(this.registerForm);
        }
      },
      reset() {
        this.registerForm.username = "";
        this.registerForm.password = "";
        this.registerForm.email = "";
        this.registerForm.displayname = "";
        this.confirmPassword = "";
      },
      closeFillAllAlert() {
        this.alertFillAll = false;
        this.overlay = false;
      },
      closePasswordNotMatchAlert() {
        this.alertPasswordNotMatch = false;
        this.overlay = false;
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
