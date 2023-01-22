<template>
  <div>
    <header>
      <div>
        <h1>Learning Gate</h1>
      </div>
      <div>
        <h2>Platform for Learning Social Topics to improve yourself and society</h2>
      </div>
    </header>
    <section id="action-section">
      <div>
        <h3>Register Now before you'll miss the great opportunity :)</h3>
      </div>
      <div>
        <v-btn color="primary" @click="toRegister">Register</v-btn>
        <v-btn color="success" @click="toLogin">Login</v-btn>
      </div>
    </section>
  </div>
</template>

<script>
  export default {
    methods: {
      toRegister() {
        this.$router.push({ path: '/register' });
      },
      toLogin() {
        this.$router.push({ path: '/login' });
      }
    }
  }
</script>

<style scoped>
  * {
    text-align: center;
  }

  #action-section {
    margin: 20px;
  }
</style>
