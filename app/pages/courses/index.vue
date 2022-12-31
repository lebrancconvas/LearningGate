<template>
  <div>
    <header>
      <div>
        <h1>Course List</h1>
      </div>
    </header>
    <section id="course-section">
      <div>
        <v-card v-for="course in courses" :key="course.id" style="margin: 10px">
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
  export default {
    data() {
      return {
        courses: [
          {
            "id": 1,
            "name": "History of World Economy",
            "description": "Learning the history of world economy, Since the industrial revolution to the present day.",
            "instructor": "Mr.Abdul Malai",
          },
          {
            "id": 2,
            "name": "Capitalist 101",
            "description": "Learning the concept of capitalism and how it works.",
            "instructor": "Mr. Mahid Bully"
          },
          {
            "id": 3,
            "name": "Socialist 101",
            "description": "Learning the concept of socialism and how it works.",
            "instructor": "Ms. Nahila Basutin"
          },
          {
            "id": 4,
            "name": "Social Development",
            "description": "Learning the basic of Social Development in the social.",
            "instructor": "Mr. Alexander S. Zhang"
          },
          {
            "id": 5,
            "name": "Union",
            "description": "Learning the history and the concept of Union for workers",
            "instructor": "Ms. Priscilla Wong"
          }
        ]
      }
    },
    methods: {
      viewCourse(id) {
        const result = this.courses.filter(course => course.id === id);
        console.log(result[0]);
      }
    }
  }
</script>

<style scoped>
  header {
    text-align: center;
  }
</style>
