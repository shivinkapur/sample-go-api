# Question

**************
Initial Schema
**************

User
====

* Email address (Text)

Question
========

* Title (Text)
* Content (Text)
* Creator (User)

******************
Initial API Design
******************

The API endpoints we want to create today include:

- Create a question
- List all questions (including the question IDs)

In the interest of time, hardcode users rather than creating endpoints for users.

For example, the API (if we choose a REST-like design) might be:

- POST /questions (for creating a question)
- GET /questions (for listing all questions)

If you prefer to use something other than REST-like, describe your choice.