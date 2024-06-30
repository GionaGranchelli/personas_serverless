<template>
  <div class="about">
    <h2>About the Persona Management System</h2>
    <p>
      The Persona Management System is designed to streamline the management of personal information. It was developed
      as a
      serverless application using AWS services, which provides high availability, scalability, and reduced operational
      overhead.
    </p>
    <div>&nbsp;</div>
    <p>
      Our system includes the following core components:
    </p>
    <div>&nbsp;</div>
    <ul>
      <li><strong>AWS Lambda:</strong> Functions to handle CRUD operations and event processing.</li>
      <li><strong>Amazon API Gateway:</strong> Provides a secure and scalable entry point for API requests.</li>
      <li><strong>Amazon DynamoDB:</strong> A NoSQL database for storing persona data.</li>
      <li><strong>Amazon EventBridge:</strong> Manages event-driven workflows, ensuring that updates and deletions
        trigger
        appropriate actions.
      </li>
      <li><strong>Amazon S3 and CloudFront:</strong> Hosts the Vue.js frontend application and delivers it efficiently
        to users.
      </li>
    </ul>
    <div>&nbsp;</div>
    <p>
      This project is a demonstration of modern web application development using serverless architecture. It aims to
      provide
      a robust and flexible system that can be easily adapted and extended for various use cases.
    </p>

  </div>
</template>

<style>

.about {
  padding: 40px;
  margin-top: 3rem;
  min-height: 100vh;
  box-sizing: border-box;
  background: rgba(0, 0, 0, .5);
  box-shadow: 0 15px 25px rgba(0, 0, 0, .6);
  border-radius: 10px;
  max-width: 800px;
}
h2 {
  margin: 0 0 30px;
  padding-top: 40px;
  color: #fff;
  text-align: center;
}
</style>
<script setup>
</script>