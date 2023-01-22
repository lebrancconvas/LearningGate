CREATE TABLE "Course" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL
);

CREATE TABLE "Title" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL
);

CREATE TABLE "Instructor" (
  "id" bigserial PRIMARY KEY,
  "title_id" bigint,
  "firstname" varchar NOT NULL,
  "surname" varchar NOT NULL,
  "description" varchar NOT NULL,
  "course_id" bigint
);

CREATE TABLE "Lesson" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "content" varchar NOT NULL,
  "course_id" bigint
);

ALTER TABLE "Lesson" ADD FOREIGN KEY ("course_id") REFERENCES "Course" ("id");

ALTER TABLE "Instructor" ADD FOREIGN KEY ("title_id") REFERENCES "Title" ("id");

CREATE TABLE "Course_Instructor" (
  "Course_id" bigserial,
  "Instructor_course_id" bigint,
  PRIMARY KEY ("Course_id", "Instructor_course_id")
);

ALTER TABLE "Course_Instructor" ADD FOREIGN KEY ("Course_id") REFERENCES "Course" ("id");

ALTER TABLE "Course_Instructor" ADD FOREIGN KEY ("Instructor_course_id") REFERENCES "Instructor" ("course_id");

