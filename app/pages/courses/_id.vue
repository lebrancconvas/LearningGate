<template>
  <div>
    <header>
      <div>
        <h1>{{ courseData.name }}</h1>
      </div>
    </header>
    <section id="course-detail-section">
      <div id="course-detail-heading">
        <h3>Course Detail</h3>
      </div>
      <div id="course-detail-canvas">
        <v-card>
          <v-card-title>
            Title: {{ courseData.name }}
          </v-card-title>
          <v-card-text>
            {{ courseData.description }}
          </v-card-text>
          <v-card-text>
            Instructor: {{ courseData.instructor }}
          </v-card-text>
          <v-card-text>
            <h2>Lesson</h2>
          </v-card-text>
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
        courseData: {},
        id: ""
      }
    },
    created() {
      const url = `http://localhost:9002/courses`;

      axios.get(url)
        .then(res => {
          this.courseData = res.data.courses;
          this.courseData = this.courseData.filter(course => course.id === parseInt(this.$route.params.id))[0];
        })
        .catch(err => {
          console.error(err);
          console.log(`[ERROR] Fetching Data error.`);
        })
    }
  }
</script>

<style scoped>
  header {
    text-align: center;
  }

  #course-detail-heading {
    margin: 10px;
  }
</style>
