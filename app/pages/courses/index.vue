<template>
  <div>
    <header>
      <div>
        <h1>Course List</h1>
      </div>
    </header>
    <section id="course-section">
      <div>
        <v-card v-for="course in results" :key="course.id" style="margin: 10px">
          <v-card-title>
            <h2>{{ course.name }}</h2>
          </v-card-title>
          <v-card-text>
            <h3>{{ course.description }}</h3>
            <p>Instructor: {{ course.instructor }}</p>
          </v-card-text>
          <v-card-actions>
            <v-btn color="primary" @click="viewCourse(course.id)">View Course</v-btn>
          </v-card-actions>
        </v-card>
      </div>
    </section>
  </div>
</template>

<script>
  import axios from 'axios';
  export default {
    data() {
      return {
        results: []
      }
    },
    created() {
      const url = `http://localhost:9002/courses`;

      axios.get(url)
        .then(response => {
          this.results = response.data.courses;
        })
        .catch(error => {
          console.error(error);
          console.log(`[ERROR] Fetching Data Error`);
        })
    },
    methods: {
      viewCourse(id) {
        this.$router.push({ path: `courses/${id}` });
      }
    }
  }
</script>

<style scoped>
  header {
    text-align: center;
  }
</style>
